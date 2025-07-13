package main

import (
	"context"
	"fmt"

	// <<<<<<< Updated upstream

	"github.com/Pallinder/go-randomdata"
	log "github.com/dhf0820/vslog"

	//"github.com/gopherjs/gopherjs/compiler/filter"//

	"github.com/davecgh/go-spew/spew"
	"os"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	//"os"
	//log "github.com/sirupsen/logrus"
	// =======
	// 	"github.com/Pallinder/go-randomdata"
	// 	"os"
	// 	"time"
	// 	//"os"
	// 	//"github.com/davecgh/go-spew/spew"

	// 	. "github.com/smartystreets/goconvey/convey"
	// >>>>>>> Stashed changes
	"testing"
	//log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbUrl = "mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority"

func TestCreateIndexTime(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.core_test")
	fmt.Printf("\n\nTestOpenDB\n")
	os.Setenv("COMPANY", "test")
	dbName := os.Getenv("COMPANY")
	fmt.Printf("env DB Name : %s\n", dbName)
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		InitCore("uc_core", "test", "test")
		db := OpenDB()
		So(db, ShouldNotBeNil)
		fmt.Printf("CurrentDataBase: %s\n", db.DatabaseName)

		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		//DB := *db.Database
		err = CreateTimeIndex(*db, "TimeTest", 30)

	})
}

func CreateTimeIndex(db MongoDB, collectionName string, seconds int) error {
	ctx := context.TODO()
	//client := db.Client
	col, err := GetCollection(collectionName)
	if err != nil {
		return err
	}
	model := mongo.IndexModel{
		Keys:    bson.M{"createdAt": 1},
		Options: options.Index().SetExpireAfterSeconds(15),
	}
	ind, err := col.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(ind)

	// insert some datas each seconds
	for i := 0; i < 5; i++ {
		name := randomdata.SillyName()
		res, err := col.InsertOne(ctx, NFT{Timestamp: time.Now(), CreatedAt: time.Now(), Name: name})
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Inserted", name, "with id", res.InsertedID)
		time.Sleep(1 * time.Second)
	}
	return nil
}

type NFT struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Name      string             `bson:"name,omitempty"`
}

func TestOpenDB(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.core_test")
	fmt.Printf("\n\nTestOpenDB\n")
	os.Setenv("COMPANY", "test")
	dbName := os.Getenv("COMPANY")
	fmt.Printf("env DB Name : %s\n", dbName)
	os.Setenv("CORE_DB", dbUrl)
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		InitCore("ids_core", "test", "test")
		db := OpenDB()
		So(db, ShouldNotBeNil)
		fmt.Printf("CurrentDataBase: %s\n", db.DatabaseName)

		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		// conf, err := InitCore("test", "test")

		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)
		// mongo := OpenDB()
		// So(mongo, ShouldNotBeNil)
		// c, err := GetCollection("configs")
		// So(err, ShouldBeNil)
		// So(c, ShouldNotBeNil)
	})
}

func TestOpenDBURL(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.core_test")
	fmt.Printf("\n\nTestOpenDB\n")
	os.Setenv("COMPANY", "test")
	dbName := os.Getenv("COMPANY")
	fmt.Printf("env DB Name : %s\n", dbName)
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		url := "mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority"
		//url := "mongodb+srv://idsUser:Sacj0nhat1@ids.fmhdg.mongodb.net/test?retryWrites=true&w=majority"
		db := OpenDBUrl(url)
		So(db, ShouldNotBeNil)

		fmt.Printf("DB Is open so see if can query\n")
		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		fmt.Printf("Config returned: %v\n", spew.Sdump(cfg))
	})
}

func TestConnectToDB(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.core_test")
	fmt.Printf("\n\nTestConnectToDB\n")
	os.Setenv("COMPANY", "test")
	dbName := os.Getenv("COMPANY")
	fmt.Printf("env DB Name : %s\n", dbName)
	os.Setenv("CORE_DB", dbUrl)
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		conf, err := InitCore("uc_core", "test", "test")

		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		mongo, err := ConnectToDB()
		So(mongo, ShouldNotBeNil)
		So(err, ShouldBeNil)
		c, err := GetCollection("configs")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
	})
}

// func TestCreateIndexTime(t *testing.T) {
// 	//t.Parallel()
// 	godotenv.Load(".env.core_test")
// 	vsLog.SetDebuglevel("DEBUG3")
// 	fmt.Printf("\n\nTestCreateIndexTime\n")
// 	os.Setenv("COMPANY", "test")
// 	dbName := os.Getenv("COMPANY")
// 	fmt.Printf("env DB Name : %s\n", dbName)
// 	Convey("Subject: Open the mongo DB", t, func() {
// 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// 		vsLog.Info("--- Convey")
// 		InitCore("uc_core", "test", "test")
// 		vsLog.Debug3("calling OpenDb")
// 		db := OpenDB()
// 		So(db, ShouldNotBeNil)
// 		vsLog.Debug3("CurrentDataBase: " + db.DatabaseName)
// 		vsLog.Debug3("DBURL returned: " + DBUrl())
// 		// cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
// 		// So(err, ShouldBeNil)
// 		// So(cfg, ShouldNotBeNil)
// 		//DB := *db.Database
// 		err := CreateTimeIndex(*db, "TimeTest", 60)
// 		So(err, ShouldBeNil)
// 	})
// }

// func CreateTimeIndex(db MongoDB, collectionName string, seconds int) error {
// 	fmt.Printf("\n\n\n")
// 	vsLog.Debug3(fmt.Sprintf("CreateTimeIndex: [%s]	", collectionName))

// 	ctx := context.TODO()
// 	//client := db.Client

// 	col, err := GetCollection(collectionName)
// 	if err != nil {
// 		return err
// 	}
// 	vsLog.Debug3("Collection: " + col.Name())
// 	vsLog.Debug3(fmt.Sprintf("Creating Index for DB: %s collection: %s ", col.Database().Name(), col.Name()))
// 	cursor, err := col.Indexes().List(ctx)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		index := bson.D{}
// 		if err = cursor.Decode(&index); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		fmt.Printf("\n\n")
// 		vsLog.Debug3("found index and drop it: " + spew.Sdump(index))
// 		// idx := index
// 		// vsLog.Debug3("Index name: " + idx["name"].(string))
// 		//col.Indexes().DropOne(ctx, idx)
// 		fmt.Printf("\n\n")
// 	}
// 	_, err = col.Indexes().DropOne(ctx, "createdAt_1")
// 	if err != nil {
// 		vsLog.Debug3("Drop Error: " + err.Error())
// 	}
// 	model := mongo.IndexModel{
// 		Keys:    bson.M{"createdAt": 1},
// 		Options: options.Index().SetExpireAfterSeconds(120),
// 	}
// 	ind, err := col.Indexes().CreateOne(ctx, model)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(ind)
// 	// see if can retrieve the index
// 	cursor, err = col.Indexes().List(ctx)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		index := bson.M{}
// 		if err = cursor.Decode(&index); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		fmt.Printf("\n\n")
// 		vsLog.Debug3("found index: " + spew.Sdump(index))
// 		fmt.Printf("\n\n")
// 	}

// 	// insert some datas each seconds
// 	for i := 0; i < 5; i++ {
// 		name := randomdata.SillyName()
// 		id := primitive.NewObjectID()
// 		res, err := col.InsertOne(ctx, NFT{ID: id, Timestamp: time.Now(), CreatedAt: time.Now(), Name: name})
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}

// 		fmt.Println("Inserted", name, "with id", res.InsertedID)
// 		nft := &NFT{}
// 		result := col.FindOne(ctx, bson.M{"_id": id}).Decode(nft)
// 		if result != nil {
// 			vsLog.Debug3("Find Error: " + result.Error())
// 		}
// 		vsLog.Debug3("FindOne: " + spew.Sdump(nft))
// 		time.Sleep(1 * time.Second)
// 	}
// 	return nil
// }

// type NFT struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	CreatedAt time.Time          `bson:"createdAt,omitempty"`
// 	Timestamp time.Time          `bson:"timestamp,omitempty"`
// 	Name      string             `bson:"name,omitempty"`
// }

// package main

// import (
// 	"context"
// 	"fmt"

// 	// <<<<<<< Updated upstream

// 	"github.com/Pallinder/go-randomdata"
// 	log "github.com/dhf0820/vslog"

// 	//"github.com/gopherjs/gopherjs/compiler/filter"

// 	//"github.com/davecgh/go-spew/spew"
// 	"os"
// 	"time"

// 	. "github.com/smartystreets/goconvey/convey"

// 	//"os"
// 	//log "github.com/sirupsen/logrus"
// 	// =======
// 	// 	"github.com/Pallinder/go-randomdata"
// 	// 	"os"
// 	// 	"time"
// 	// 	//"os"
// 	// 	//"github.com/davecgh/go-spew/spew"

// 	// 	. "github.com/smartystreets/goconvey/convey"
// 	// >>>>>>> Stashed changes
// 	"testing"
// 	//log "github.com/sirupsen/logrus"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func TestCreateIndexTime(t *testing.T) {
// 	//t.Parallel()
// 	godotenv.Load(".env.core_test")
// 	fmt.Printf("\n\nTestOpenDB\n")
// 	os.Setenv("COMPANY", "test")
// 	dbName := os.Getenv("COMPANY")
// 	fmt.Printf("env DB Name : %s\n", dbName)
// 	Convey("Subject: Open the mongo DB", t, func() {
// 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// 		fmt.Printf("\n\n--- Convey")
// 		InitCore("uc_core", "test", "test")
// 		db := OpenDB()
// 		So(db, ShouldNotBeNil)
// 		fmt.Printf("CurrentDataBase: %s\n", db.DatabaseName)

// 		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
// 		So(err, ShouldBeNil)
// 		So(cfg, ShouldNotBeNil)
// 		//DB := *db.Database
// 		err = CreateTimeIndex(*db, "TimeTest", 30)

// 	})
// }

// func CreateTimeIndex(db MongoDB, collectionName string, seconds int) error {
// 	ctx := context.TODO()
// 	//client := db.Client
// 	col, err := GetCollection(collectionName)
// 	if err != nil {
// 		return err
// 	}
// 	model := mongo.IndexModel{
// 		Keys:    bson.M{"createdAt": 1},
// 		Options: options.Index().SetExpireAfterSeconds(15),
// 	}
// 	ind, err := col.Indexes().CreateOne(ctx, model)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(ind)

// 	// insert some datas each seconds
// 	for i := 0; i < 5; i++ {
// 		name := randomdata.SillyName()
// 		res, err := col.InsertOne(ctx, NFT{Timestamp: time.Now(), CreatedAt: time.Now(), Name: name})
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		fmt.Println("Inserted", name, "with id", res.InsertedID)
// 		time.Sleep(1 * time.Second)
// 	}
// 	return nil
// }

// type NFT struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	CreatedAt time.Time          `bson:"createdAt,omitempty"`
// 	Timestamp time.Time          `bson:"timestamp,omitempty"`
// 	Name      string             `bson:"name,omitempty"`
// }

// func TestOpenDB(t *testing.T) {
// 	//t.Parallel()
// 	godotenv.Load(".env.core_test")
// 	fmt.Printf("\n\nTestOpenDB\n")
// 	os.Setenv("COMPANY", "test")
// 	dbName := os.Getenv("COMPANY")
// 	fmt.Printf("env DB Name : %s\n", dbName)
// 	Convey("Subject: Open the mongo DB", t, func() {
// 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// 		fmt.Printf("\n\n--- Convey")
// 		InitCore("uc_core", "test", "test")
// 		db := OpenDB()
// 		So(db, ShouldNotBeNil)
// 		fmt.Printf("CurrentDataBase: %s\n", db.DatabaseName)

// 		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
// 		So(err, ShouldBeNil)
// 		So(cfg, ShouldNotBeNil)
// 		// conf, err := InitCore("test", "test")

// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// mongo := OpenDB()
// 		// So(mongo, ShouldNotBeNil)
// 		// c, err := GetCollection("configs")
// 		// So(err, ShouldBeNil)
// 		// So(c, ShouldNotBeNil)
// 	})
// }

// func TestOpenDBURL(t *testing.T) {
// 	//t.Parallel()
// 	godotenv.Load(".env.core_test")
// 	fmt.Printf("\n\nTestOpenDB\n")
// 	os.Setenv("COMPANY", "test")
// 	dbName := os.Getenv("COMPANY")
// 	fmt.Printf("env DB Name : %s\n", dbName)
// 	Convey("Subject: Open the mongo DB", t, func() {
// 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// 		fmt.Printf("\n\n--- Convey")
// 		url := "mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority"
// 		//url := "mongodb+srv://idsUser:Sacj0nhat1@ids.fmhdg.mongodb.net/test?retryWrites=true&w=majority"
// 		db := OpenDBUrl(url)
// 		So(db, ShouldNotBeNil)

// 		fmt.Printf("DB Is open so ce if can query\n")
// 		cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
// 		So(err, ShouldBeNil)
// 		So(cfg, ShouldNotBeNil)
// 	})
// }

// func TestConnectToDB(t *testing.T) {
// 	//t.Parallel()
// 	godotenv.Load(".env.core_test")
// 	fmt.Printf("\n\nTestConnectToDB\n")
// 	os.Setenv("COMPANY", "test")
// 	dbName := os.Getenv("COMPANY")
// 	fmt.Printf("env DB Name : %s\n", dbName)
// 	Convey("Subject: Open the mongo DB", t, func() {
// 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// 		fmt.Printf("\n\n--- Convey")
// 		conf, err := InitCore("uc_core", "test", "test")

// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		mongo, err := ConnectToDB()
// 		So(mongo, ShouldNotBeNil)
// 		So(err, ShouldBeNil)
// 		c, err := GetCollection("configs")
// 		So(err, ShouldBeNil)
// 		So(c, ShouldNotBeNil)
// 	})
// }

// // func TestCreateIndexTime(t *testing.T) {
// // 	//t.Parallel()
// // 	godotenv.Load(".env.core_test")
// // 	vsLog.SetDebuglevel("DEBUG3")
// // 	fmt.Printf("\n\nTestCreateIndexTime\n")
// // 	os.Setenv("COMPANY", "test")
// // 	dbName := os.Getenv("COMPANY")
// // 	fmt.Printf("env DB Name : %s\n", dbName)
// // 	Convey("Subject: Open the mongo DB", t, func() {
// // 		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
// // 		vsLog.Info("--- Convey")
// // 		InitCore("uc_core", "test", "test")
// // 		vsLog.Debug3("calling OpenDb")
// // 		db := OpenDB()
// // 		So(db, ShouldNotBeNil)
// // 		vsLog.Debug3("CurrentDataBase: " + db.DatabaseName)
// // 		vsLog.Debug3("DBURL returned: " + DBUrl())
// // 		// cfg, err := GetSvcConfig(context.Background(), "test", "test", "test")
// // 		// So(err, ShouldBeNil)
// // 		// So(cfg, ShouldNotBeNil)
// // 		//DB := *db.Database
// // 		err := CreateTimeIndex(*db, "TimeTest", 60)
// // 		So(err, ShouldBeNil)
// // 	})
// // }

// // func CreateTimeIndex(db MongoDB, collectionName string, seconds int) error {
// // 	fmt.Printf("\n\n\n")
// // 	vsLog.Debug3(fmt.Sprintf("CreateTimeIndex: [%s]	", collectionName))

// // 	ctx := context.TODO()
// // 	//client := db.Client

// // 	col, err := GetCollection(collectionName)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	vsLog.Debug3("Collection: " + col.Name())
// // 	vsLog.Debug3(fmt.Sprintf("Creating Index for DB: %s collection: %s ", col.Database().Name(), col.Name()))
// // 	cursor, err := col.Indexes().List(ctx)
// // 	if err != nil {
// // 		log.Fatal(err.Error())
// // 	}
// // 	defer cursor.Close(ctx)
// // 	for cursor.Next(ctx) {
// // 		index := bson.D{}
// // 		if err = cursor.Decode(&index); err != nil {
// // 			log.Fatal(err.Error())
// // 		}
// // 		fmt.Printf("\n\n")
// // 		vsLog.Debug3("found index and drop it: " + spew.Sdump(index))
// // 		// idx := index
// // 		// vsLog.Debug3("Index name: " + idx["name"].(string))
// // 		//col.Indexes().DropOne(ctx, idx)
// // 		fmt.Printf("\n\n")
// // 	}
// // 	_, err = col.Indexes().DropOne(ctx, "createdAt_1")
// // 	if err != nil {
// // 		vsLog.Debug3("Drop Error: " + err.Error())
// // 	}
// // 	model := mongo.IndexModel{
// // 		Keys:    bson.M{"createdAt": 1},
// // 		Options: options.Index().SetExpireAfterSeconds(120),
// // 	}
// // 	ind, err := col.Indexes().CreateOne(ctx, model)
// // 	if err != nil {
// // 		log.Fatal(err.Error())
// // 	}
// // 	fmt.Println(ind)
// // 	// see if can retrieve the index
// // 	cursor, err = col.Indexes().List(ctx)
// // 	if err != nil {
// // 		log.Fatal(err.Error())
// // 	}
// // 	defer cursor.Close(ctx)
// // 	for cursor.Next(ctx) {
// // 		index := bson.M{}
// // 		if err = cursor.Decode(&index); err != nil {
// // 			log.Fatal(err.Error())
// // 		}
// // 		fmt.Printf("\n\n")
// // 		vsLog.Debug3("found index: " + spew.Sdump(index))
// // 		fmt.Printf("\n\n")
// // 	}

// // 	// insert some datas each seconds
// // 	for i := 0; i < 5; i++ {
// // 		name := randomdata.SillyName()
// // 		id := primitive.NewObjectID()
// // 		res, err := col.InsertOne(ctx, NFT{ID: id, Timestamp: time.Now(), CreatedAt: time.Now(), Name: name})
// // 		if err != nil {
// // 			log.Fatal(err.Error())
// // 		}

// // 		fmt.Println("Inserted", name, "with id", res.InsertedID)
// // 		nft := &NFT{}
// // 		result := col.FindOne(ctx, bson.M{"_id": id}).Decode(nft)
// // 		if result != nil {
// // 			vsLog.Debug3("Find Error: " + result.Error())
// // 		}
// // 		vsLog.Debug3("FindOne: " + spew.Sdump(nft))
// // 		time.Sleep(1 * time.Second)
// // 	}
// // 	return nil
// // }

// // type NFT struct {
// // 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// // 	CreatedAt time.Time          `bson:"createdAt,omitempty"`
// // 	Timestamp time.Time          `bson:"timestamp,omitempty"`
// // 	Name      string             `bson:"name,omitempty"`
// // }
