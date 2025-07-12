package main

import (
	"bytes"

	"github.com/davecgh/go-spew/spew"

	//"encoding/json"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
	"time"

	vsLog "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	//"os"
	//"github.com/davecgh/go-spew/spew"
	//"io/ioutil"
)

var (
	Conf              *ServiceConfig
	ReleaseEp         EndPoint
	DeliverEp         *EndPoint
	DocumentEp        *EndPoint
	ConnectedServices map[string]string
)

func InitCoreFromEnv(envName string) (*ServiceConfig, error) {
	OpenDB()
	var err error
	fname := os.Getenv(envName)
	fmt.Printf("InitCoreFromEnv using: [%s]\n", fname)
	Conf, err = GetSvcConfigFromFile(context.Background(), fname)
	return Conf, err
}

func InitCore(name, version, company string) (*ServiceConfig, error) {
	vsLog.Debug2(fmt.Sprintf("Starting with name: %s, version: %s, company: %s", name, version, company))
	vsLog.Debug2("Setting SYSTEMDBNAME: " + company)
	os.Setenv("SYSTEMDBNAME", company)
	//dbName := os.Getenv("CORE_DB")
	//fmt.Printf("\n--InitCore:47 -- Opening Database: %s  \n", os.Getenv("CORE_DB"))
	startTime := time.Now()
	OpenDB()
	vsLog.Debug3("InitCore OpenDB took " + fmt.Sprintf("%f seconds", time.Since(startTime).Seconds()) + " to open")
	ConnectedServices = make(map[string]string)
	var err error
	Conf, err = GetSvcConfig(context.Background(), name, version, company)
	if err != nil {
		err := fmt.Errorf("core:54 -- GetSvcConfig error %s", err.Error())
		return nil, err
	}
	vsLog.Debug3(fmt.Sprintf("InitCore --  ServiceConfig: %s", spew.Sdump(Conf)))
	setEndPoints()
	return Conf, err
}

func GetConfig() *ServiceConfig {
	//vsLog.Debug2(fmt.Sprintf("GetConfig returning: %s ", spew.Sdump(Conf)))
	return Conf
}

// Services connected to the system remotely
func GetConnectedServiceVersion(service string) string {
	return ConnectedServices[service]
}

func SetConnectedServiceVersion(service, version string) {
	vsLog.Debug2(fmt.Sprintf("--  SetConnectServicesVersion: %s - %s", service, version))
	//ConnectedServices[service] = version
}

func GetOriginsAllowed() []string {
	return Conf.OriginsAllowed
	// allowed := ""
	// for _, oa := range Conf.OriginsAllowed {
	// 	if len(allowed) == 0 {
	// 		allowed = oa.Name
	// 	} else {
	// 		allowed = allowed + "," + oa.Name
	// 	}
	// }
	// return allowed
}

func setEndPoints() {
	// ReleaseEp = GetServiceEndpoint("release")
	// if ReleaseEp == nil {
	// 	vsLog.Errorf("Release EndPoint was not found in configuration")
	// }
	// DeliverEp = GetServiceEndpoint("delivery")
	// if DeliverEp == nil {
	// 	vsLog.Errorf("Delivery EndPoint was not found in configuration")
	// }
}

func GetServiceEndpoint(value string) *EndPoint {
	vsLog.Debug2("Calling GetConfig")
	cfg := GetConfig()
	if cfg == nil {
		vsLog.Errorf("No configuration!!")
		return nil
	}
	//fmt.Printf("GetServiceEndPoint: %s\n", spew.Sdump(cfg))
	endPoints := cfg.ServiceEndPoints
	vsLog.Debug2("Core Endpoints: " + spew.Sdump(endPoints))
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			return ep
		}
	}
	return nil
}

//func GetMyEndpoint(endPoints []*EndPoint, value string) *EndPoint {
//	//endPoints := GetConfig().MyEndPoints
//	fmt.Printf("Core Endpoints: %v", endPoints)
//	for _, ep := range endPoints {
//		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
//		if ep.Name == value {
//			return ep
//		}
//	}
//	return nil
//}

func GetMyEndpoints(cfg *ServiceConfig) []*EndPoint {
	//vsLog.Debug2(fmt.Sprintf("Calling GetMyEndpoints for ServiceConfig: %s", cfg.ID.Hex()))
	endPoints := cfg.MyEndPoints
	//vsLog.Debug2("Core Endpoints: " + spew.Sdump(endPoints))
	return endPoints
}

func GetDB() MongoDB {
	return DB
}

// func (doc *Document) WriteGridFs(imageData []byte) (primitive.ObjectID, error) {
func WriteGridFs(metaData map[string]string, imageData []byte) (primitive.ObjectID, error) {
	startTime := time.Now()
	//mdb := db.DB.Database

	bucket, err := gridfs.NewBucket(
		DB.Database,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		err = fmt.Errorf("WriteGridFS:154  --  Unable to get GridFS Bucket: %s", err)
		fmt.Printf("%s\n", err.Error())
		return primitive.NilObjectID, err
	}
	//client := "carenow"
	//facility := "demo"
	//mrn :=  "011621"
	//docid := primitive.NewObjectID()
	facility := metaData["facility"]
	mrn := metaData["mrn"]
	srcID := metaData["src_id"]
	//metaData := make(map[string]string)
	//metaData["content_type"] = "pdf"
	//metaData["mrn"] = mrn
	//fileName := fmt.Sprintf("%s_%s_%s_%s", doc.Client, doc.Facility, doc.MRN, doc.ID.Hex())
	//fileName := fmt.Sprintf("%s_%s_%s_%s", client, facility, mrn, docid)
	fileName := fmt.Sprintf("%s_%s_%s", facility, srcID, mrn)
	saveImage, err := bucket.OpenUploadStream(
		fileName,
		options.GridFSUpload().SetMetadata(metaData),
	)
	if err != nil {
		err = fmt.Errorf("WriteGridFs:166 - OpenUploadStream failed: %s", err)
		fmt.Println(err)
		return primitive.NilObjectID, err
	}
	defer saveImage.Close()

	fileSize, err := saveImage.Write(imageData)
	if err != nil {
		err = fmt.Errorf("WriteGridFs:174 - Save GridFS failed: %v", err)
		fmt.Println(err)
		return primitive.NilObjectID, err
	}
	fmt.Printf("WriteGridFs:178 - Gridfs Saved %d bytes in %f seconds\n", fileSize, time.Since(startTime).Seconds())
	return saveImage.FileID.(primitive.ObjectID), nil
}

func getGridFsImage(imageID primitive.ObjectID) (*[]byte, error) {
	db := DB.Database
	startTime := time.Now()
	bucket, err := gridfs.NewBucket(
		db,
	)
	if err != nil {
		return nil, vsLog.Errorf("New Bucket failed: " + err.Error())
	}
	var buf bytes.Buffer
	fmt.Printf("Retrieving ID:%s\n", imageID)
	_, err = bucket.DownloadToStream(imageID, &buf)
	if err != nil {
		err = fmt.Errorf("getGridFsImage:218  --  DownloadToStream failed: %s", err.Error())
		return nil, err
	}
	image := buf.Bytes()
	fmt.Printf("getGridFsImage:224  --  getGridfs retrieved in %f seconds", time.Since(startTime).Seconds())
	return &image, nil
}
