package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	log "github.com/dhf0820/vslog"
	vsLog "github.com/dhf0820/vslog"

	//"github.com/sirupsen/logrus"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
	URL          string
	Database     *mongo.Database
	Session      mongo.Session
	Collection   *mongo.Collection
}

var DB MongoDB
var mongoClient *mongo.Client

// var DbConnector *DataConnector
var insertResult *mongo.InsertOneResult

func OpenDBUrl(dbURL string) *MongoDB {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	startTime := time.Now()
	uri := dbURL
	VLog("INFO", "Opening database: "+uri)
	//uri := dbURL + databaseName
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	// error = client.Connect(ctx)

	// //Checking the connection
	// error = client.Ping(context.TODO(), nil)
	// fmt.Println("Database connected")

	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if mongoClient, err = mongo.Connect(ctx, opts); err != nil {
		vsLog.Error("mongo.Connect error: " + err.Error())
		return nil
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect:62 %v\n", err)
		vsLog.Errorf("Database did not connect err: " + err.Error())
		return nil
	}

	VLog("INFO", "using Company: "+os.Getenv("COMPANY"))
	DB.Client = mongoClient
	DB.DatabaseName = os.Getenv("COMPANY") //DbConnector.Database  //databaseName
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	vsLog.Info("Database: " + DB.DatabaseName + " connected")
	//fmt.Printf("Client: %s\n", spew.Sdump(client))

	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	vsLog.Debug3(fmt.Sprintf("took %d ms", time.Since(startTime).Milliseconds()))
	return &DB
}

func OpenDB() *MongoDB {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	//startTime := time.Now()
	//DbConnector = svcConfig.DataConnector
	//dbURL := DbConnector.Server
	dbURL := DBUrl() //os.Getenv("CORE_DB")
	uri := dbURL
	vsLog.Info("Opening MongoDB connection: " + dbURL)
	vsLog.Info("Using CORE_DB: " + dbURL)
	vsLog.Info("Opening database: " + uri + " using Company: " + os.Getenv("COMPANY"))
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	clientOptions.SetMaxPoolSize(5)

	//fmt.Printf("vsmongo:110 -- Using new connect routine from atlas\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		vsLog.Error("Mongo.Connect error: " + err.Error())
		panic(err.Error())
	}
	if mongoClient == nil {
		vsLog.Error("Mongo.Connect error: mongoClient is nil")
		panic("mongoClient is nil")
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect: %s\n", err.Error())
		VsLog("ERROR", "Database did not connect by ping:"+err.Error())
		err = pingDb(5)
		VsLog("INFO", "Direct Ping failed error:"+err.Error())
		return nil
	}
	DB.Client = mongoClient
	DB.DatabaseName = os.Getenv("COMPANY") //DbConnector.Database  //databaseName
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	VsLog("INFO", "Using Database: "+DB.DatabaseName)
	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	VsLog("INFO", " Database "+DB.DatabaseName+" Connected")
	return &DB
}

func DBUrl() string {
	coreDB := os.Getenv("CORE_DB")
	vsLog.Debug2("CORE_DB Env value: " + coreDB)
	return coreDB
}

// ConnectToDB starts a new database connection and returns a reference to it
func ConnectToDB() (*MongoDB, error) {
	url := DBUrl()
	if url == "" {
		log.Fatal("coreDB is not defined. Should contain the name of the actual Database to use\n")
	}
	databaseName := os.Getenv("COMPANY")
	vsLog.Debug2("Using DB: " + databaseName)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(url)

	options.SetMaxPoolSize(DbPoolSize())
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	DB.Client = client
	DB.DatabaseName = databaseName
	DB.Database = client.Database(DB.DatabaseName)
	DB.URL = url
	return &DB, nil
}

func DbPoolSize() uint64 {
	var poolSize uint64
	poolSizeString := GetDbField("poolsize")
	if poolSizeString == "" {
		poolSizeString = "100"
	}
	poolSizeInt64, err := strconv.ParseInt(poolSizeString, 10, 64)
	if err == nil {
		poolSize = uint64(poolSizeInt64)
	} else {
		poolSize = 100
	}
	return poolSize
}
func Current(dbName string) (*MongoDB, error) {
	if DB.Client != nil {
		if dbName == DB.DatabaseName {
			//vsLog.Info("CurrentDB: " + DB.DatabaseName + " exists, use it")
			return &DB, nil
		}
	}
	vsLog.Debug4("Switching to dbName: " + dbName + " ConnectToDB() being called")
	_, err := ConnectToDB()
	//client, err := Open("")
	return &DB, err
}

func (db *MongoDB) Close() error {
	err := db.Client.Disconnect(context.TODO())
	return err
}

func GetCollection(collection string) (*mongo.Collection, error) {
	if collection == "" {
		collection = CollectionName()
		vsLog.Warn("Using default Collection: " + collection)
	}
	dbName := os.Getenv("SYSTEMDBNAME")

	if dbName == "" {
		vsLog.Debug2("SYSTEMDBNAME not set, using default: test")
		dbName = "test"
	}
	vsLog.Debug2("Using DB: " + dbName)
	db, err := Current(dbName) //"mongodb://admin:Sacj0nhat1@cat.vertisoft.com:27017")
	if err != nil {
		log.Fatal("Current DB returned error: " + err.Error())
		//return nil, err
	}
	vsLog.Debug2("Changed to Collection: " + collection + " in database: " + DB.DatabaseName)
	client := db.Client
	coll := client.Database(DB.DatabaseName).Collection(collection)

	//vsLog.Debug2("Changed to Collection: " + collection + " in database: " + DB.DatabaseName)
	return coll, nil
}

func CollectionName() string {
	return "srv_config"
}

func GetDbField(key string) string {
	return ""
	// //LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
	// flds := mod.DataConnector.Fields
	// for _, fld := range flds {
	// 	switch {
	// 	case fld.Name == key:
	// 		return fld.Value
	// 	}
	// }
	// return ""

}

// IsDup returns whether err informs of a duplicate key error because
// a primary key index or a secondary unique index already has an entry
// with the given value.
func IsDup(err error) bool {
	if wes, ok := err.(mongo.WriteException); ok {
		for i := range wes.WriteErrors {
			if wes.WriteErrors[i].Code == 11000 || wes.WriteErrors[i].Code == 11001 || wes.WriteErrors[i].Code == 12582 || wes.WriteErrors[i].Code == 16460 {
				return true
			}
		}
	}
	return false
}

func pingDb(count int) error {
	for i := 0; i < count; i++ {
		cmd := exec.Command("ping", "-c 5", "35.167.154.12")
		if cmd.Err != nil {
			vsLog.Error("Ping error: " + cmd.Err.Error())
		} else {
			vsLog.Debug3("Ping successful")
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Ping failed after %d tries", count))
}
