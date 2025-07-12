package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	//"io"

	//"net/http"
	"strings"

	//"github.com/davecgh/go-spew/spew"

	common "github.com/dhf0820/uc_common"
	vsLog "github.com/dhf0820/vslog"

	//"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

/*
ServiceConfig is the basic configuration for this all services available to the current service
what it can talk to and who can talk to it
*/
// type ServiceConfig struct {
// 	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	Name    string             `json:"name"`
// 	Version string             `json:"version"`
// 	//Messaging  		Messaging					// move to endpoints
// 	DataConnector    *DataConnector        `json:"dataconnector"`
// 	Services         []*ServiceScope       `json:"services" bson:"services"`
// 	MyEndPoints      []*EndPoint `json:"myendpoints"`
// 	ServiceEndPoints []*EndPoint `json:"serviceendpoints" bson:"service_endpoints"`
// 	ConnectInfo      []*KVData   `json:"connect_info" bson:"connect_info"`
// }

// type RunTimeConfig struct {
// 	CoreDB        string
// 	CoreDataBase  string
// 	ServiceName   string
// 	ConfigVersion string
// 	Company       string
// 	RefreshSecret string
// 	AccessSecret  string
// 	TokenDuration int
// 	ListenPort    string
// 	CfgString     string
// 	Port          string
// }

// type ServiceScope struct {
// 	Name  string `json:"name" bson:"name"`
// 	Scope string `json:"scope" bson:"scope"` // min, norm, max
// }

// type DataConnector struct {
// 	Server     string              `json:"server"`
// 	User       string              `json:"user"`
// 	Password   string              `json:"password"`
// 	Database   string              `json:"database"`
// 	Collection string              `json:"collection"`
// 	Fields     []*KVData `json:"fields"`
// }

// type EndPoint struct { // Replaces BaseUrl
// 	Name        string //internal name
// 	Label       string `json:"label"`
// 	Scope       string `json:"scope,omitempty" bson:"scope"`
// 	Protocol    string `json:"protocol" bson:"protocol"` // grpc or amqp
// 	Address     string `json:"address" bson:"address"`   //How do I get to this service
// 	Port        string `json:"port"`
// 	Credentials string `json:"credentials" bson:"credentials"`
// 	CertName    string `json:"certname" bson:"cert_name"`
// 	TLSMode     string `json:"tlsmode" bson:"tls_mode"`
// 	DeployMode  string `json:"deploymode" bson:"deploy_mode"`
// }

// type ConnectInfo struct {
// 	//ID    string `json:"id" bson:"id,omitempty"`
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type KVData struct {
// 	Name  string
// 	Value string
// }

// var ServiceName string
// var RunEnv string
// var Company string
// var RunConfig *RunTimeConfig

func ConfigFromJsonFile(fileName string) (*ServiceConfig, error) {
	cfgFile, err := os.ReadFile(fileName)
	//cfgFile, err := ioutil.ReadFile("/Users/dhf/work/roi/services/core_service/config/config.json")
	if err != nil {
		fmt.Printf("config:99  --  Can not read configuration file: %s\n", fileName)
		return nil, fmt.Errorf("can not read configuration file: %s, err: %v", fileName, err)
	}
	fmt.Printf("config:102  --  Config file read: %s: \n%s \n", fileName, cfgFile)
	conf := &ServiceConfig{}
	err = json.Unmarshal(cfgFile, conf)
	if err != nil {
		err = fmt.Errorf("config:106  --  Unmarshal Config: %s  err: %v", fileName, err)
		return nil, err
	}
	return conf, nil
}

// func CreateSvcConfig(ctx context.Context, data *ServiceConfig) (*ServiceConfig, error) {
// 	collection, err := GetCollection("srv_config")
// 	if err != nil {
// 		return nil, err
// 	}
// 	//timeNow := time.Now()
// 	//data.CreatedAt = &timeNow
// 	//data.UpdatedAt = data.CreatedAt
// 	data.ID = primitive.NewObjectID()
// 	fmt.Printf("config:122  --  ID: %s\n", data.ID.Hex())
// 	fmt.Printf("config:123  --  Inserting: %s\n",.Sdump(data))
// 	res, err := collection.InsertOne(ctx, data)
// 	if err != nil {
// 		vsLog.Errorf("config:124  --  domain.CreateSvrConfig InsertOne failed: %v\n", err)
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("config:129  --  InsertOne failed creating SVRConfig: %v", err),
// 		)
// 	}
// 	oid, ok := res.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		vsLog.Errorf("config:132  --  Invalid ID from insert\n")
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("config:132  --  Environment.Create Cannot convert OID: [%v]", oid),
// 		)
// 	}
// 	//fmt.Printf("New id : %v\n", oid)
// 	data.ID = oid

// 	//fmt.Printf("New Environment: %s\n", Sdump(e))
// 	return data, nil
// }

func GetSvcConfigFromFile(ctx context.Context, fname string) (*ServiceConfig, error) {
	if fname == "" {
		fname = "/Users/dhf/work/roi/services/core_service/config/core_test.json"
		vsLog.Debug3("--  Fname is blank using default: " + fname)

	}
	return ConfigFromJsonFile(fname)
}

func GetSvcConfig(ctx context.Context, svcName, version, company string) (*ServiceConfig, error) {
	// fname := "/Users/dhf/work/roi/services/core_service/config/core_test.json"
	// svcConfig, err := ConfigFromJsonFile(fname)
	vsLog.Debug2(fmt.Sprintf("Name: %s, version: %s, company: %s", svcName, version, company))
	collection, err := GetCollection("srv_config")
	if err != nil {
		return nil, err
	}

	//filter := bson.D{bson.E{Key:"name", Value: svcName},bson.E{Key:"version", Value: version},
	//	bson.E{Key:"customer.name", Value: company}}
	filter := bson.D{bson.E{Key: "name", Value: svcName}, bson.E{Key: "version", Value: version}}
	srvConfig := ServiceConfig{}
	vsLog.Debug2("collectionName: " + collection.Name())
	//fmt.Printf("Collection: %s\n", Sdump(collection))
	vsLog.Debug2(fmt.Sprintf("Calling FindOne SvcConfig: %v", filter))
	err = collection.FindOne(context.Background(), filter).Decode(&srvConfig)
	if err != nil {
		vsLog.Debug2(fmt.Sprintf("FindOne %v err: %s", filter, err.Error()))
		return nil, vsLog.Errorf(fmt.Sprintf("FindOne %v NotFound", filter))
	}
	//vsLog.Debug2("Returning Config: " + spew.Sdump(srvConfig))
	vsLog.Debug2("Calling SetConnectedServiceVersion")
	SetConnectedServiceVersion(svcName, version)
	return &srvConfig, err
}

// func GetSvrConfig(ctx context.Context, svcName, version, company string) (*common.Svr, error) {
// 	// fname := "/Users/dhf/work/roi/services/core_service/config/core_test.json"
// 	// svcConfig, err := ConfigFromJsonFile(fname)
// 	vsLog.Debug2(fmt.Sprintf("Name: %s, version: %s, company: %s", svcName, version, company))
// 	collection, err := GetCollection("srv_config")
// 	if err != nil {
// 		return nil, err
// 	}

// 	//filter := bson.D{bson.E{Key:"name", Value: svcName},bson.E{Key:"version", Value: version},
// 	//	bson.E{Key:"customer.name", Value: company}}
// 	filter := bson.D{bson.E{Key: "name", Value: svcName}, bson.E{Key: "version", Value: version}}
// 	srvConfig := ServiceConfig{}
// 	vsLog.Debug2("collectionName: " + collection.Name())
// 	//fmt.Printf("Collection: %s\n", Sdump(collection))
// 	vsLog.Debug2(fmt.Sprintf("Calling FindOne SvcConfig: %v", filter))
// 	err = collection.FindOne(context.Background(), filter).Decode(&srvConfig)
// 	if err != nil {
// 		vsLog.Debug2(fmt.Sprintf("FindOne %v err: %s", filter, err.Error()))
// 		return nil, vsLog.Errorf(fmt.Sprintf("FindOne %v NotFound", filter))
// 	}
// 	vsLog.Debug2("Returning Config: " + spew.Sdump(srvConfig))
// 	vsLog.Debug2("Calling SetConnectedServiceVersion")
// 	SetConnectedServiceVersion(svcName, version)
// 	return &srvConfig, err
// }

func GetConnectorConfig(nameVersion string) (*common.ConnectorConfig, error) {
	parts := strings.Split(nameVersion, ":")
	name := parts[0]
	if len(parts) == 1 {
		return nil, vsLog.Errorf("--  nameVersion format is name:version")
	}
	version := parts[1]
	vsLog.Debug3(fmt.Sprintf("--   Name: %s, version: %s\n", name, version))
	collection, err := GetCollection("ConnectorConfig")
	if err != nil {
		return nil, err
	}
	//filter := bson.D{bson.E{Key:"name", Value: svcName},bson.E{Key:"version", Value: version},
	//	bson.E{Key:"customer.name", Value: company}}
	filter := bson.D{bson.E{Key: "name", Value: name}, bson.E{Key: "version", Value: version}}
	connConfig := common.ConnectorConfig{}
	vsLog.Debug3(" --  Calling FindOne SvcConfig: " + fmt.Sprint(filter))
	err = collection.FindOne(context.Background(), filter).Decode(&connConfig)
	if err != nil {
		//vsLog.Errorf(fmt.Sprintf("FindOne [%v] NotFound\n", filter))
		return nil, vsLog.Errorf(fmt.Sprintf("FindOne [%v] NotFound\n", filter))
	}
	//fmt.Printf("GetConnectorConfig:209  --  NO ERRORS REPORTED: %s\n", Sdump(connConfig))
	//fmt.Printf("   Returning Config : %s\n", Sdump(svcConfig))
	// SetConnectedServiceVersion(svcName, version)
	return &connConfig, err
}

// func GetConnectorConfig(name, version string) (*ConnectorConfig, error) {
// 	ep := GetServiceEndpoint("delivery")
// 	fmt.Printf("EP: %s\n", spew.Sdump(ep))
// 	url := fmt.Sprintf("%s://%s/connector?name=%s&version=%s", ep.Protocol, ep.Address, name, version)
// 	fmt.Printf("Final URL: %s\n", url)
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		err = fmt.Errorf("Error in req: %s", err)
// 		return nil, err
// 	}
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		err = fmt.Errorf("error in request: %s", err)
// 		return nil, err
// 	}
// 	fmt.Printf("\n183 -###Processsing the response\n")
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	b := string(body)
// 	fmt.Printf("b: %s\n", b)

// 	var response ConnectorResponse
// 	//var rel  releasePkg.Release
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		fmt.Printf("##### UnMarshal err: %v\n", err)
// 	} else {
// 		fmt.Printf("GetRelease UnMarshal: %s\n", Sdump(response))
// 	}
// 	return response.Connector, nil
// 	//fmt.Printf("In GetConnectorConfig-164\n")
// 	//collection, err := GetCollection("connector_config")
// 	//if err != nil {
// 	//	return nil, err
// }

//mailId, _ := primitive.ObjectIDFromHex("6038642f89b12233ed022283")

//connectorConfig := pkg.ConnectorConfig{}

//fmt.Printf("   Now Calling GetConnectorConfig FindOne ConnectorConfig: %v\n", filter)
//err = collection.FindOne(context.Background(), filter).Decode(&connectorConfig)
//if err != nil {
//	return nil, fmt.Errorf("GetConnectorConfig FindOne [%v] NotFound\n", filter)
//}
//fmt.Printf("   NO ERRORS REPORTED\n")
////fmt.Printf("   Returning Config : %s\n", spew.Sdump(svcConfig))
////SetConnectedServiceVersion(svcName, version )
//return nil, nil
//return &connectorConfig, err

//func GetSvcConfigForCustomer(ctx context.Context, svcName, version, customer string) (*mod.ServiceConfig, error) {
//	// fname := "/Users/dhf/work/roi/services/core_service/config/core_test.json"
//	// svcConfig, err := ConfigFromJsonFile(fname)
//	fmt.Printf("In GetSvcConfig-133\n")
//	collection, err := GetCollection("srv_config")
//	if err != nil {
//		return nil, err
//	}
//
//	filter := bson.D{bson.E{Key:"name", Value: svcName},bson.E{Key:"version", Value: version},bson.E{Key:}}
//	svcConfig := mod.ServiceConfig{}
//
//	fmt.Printf("   Now Calling GetSvcConfig FindOne SvcConfig: %v\n", filter)
//	err = collection.FindOne(context.Background(), filter).Decode(&svcConfig)
//	if err != nil {
//		vsLog.Errorf("GetSvcConfig FindOne [%v] NotFound\n", filter)
//		return nil, status.Errorf(
//			codes.NotFound,
//			fmt.Sprintf("GetSvcConfig FindOne [%v] NotFound\n", filter),
//		)
//	}
//	fmt.Printf("   NO ERRORS REPORTED\n")
//	fmt.Printf("   Returning Config : %s\n", spew.Sdump(svcConfig))
//	return &svcConfig, err
//}

// func GetConnectorConfigById(id primitive.ObjectID) (*common.ConnectorConfig, error) {
// 	vsLog.Debug3(" -- Connector Id: " + id.Hex())
// 	collection, err := GetCollection("ConnectorConfig")
// 	if err != nil {
// 		return nil, err
// 	}

//		//filter := bson.D{bson.E{Key:"name", Value: svcName},bson.E{Key:"version", Value: version},
//		//	bson.E{Key:"customer.name", Value: company}}
//		filter := bson.M{"_id": id}
//		connConfig := common.ConnectorConfig{}
//		//fmt.Printf("Collection: %s\n", spew.Sdump(collection))
//		vsLog.Info(fmt.Sprintf("Calling FindOne ConnectorConfig: %v", filter))
//		err = collection.FindOne(context.Background(), filter).Decode(&connConfig)
//		if err != nil {
//			vsLog.Errorf(fmt.Sprintf("FindOne %v NotFound\n", filter))
//			return nil, vsLog.Errorf(fmt.Sprintf("FindOne %v NotFound\n", filter))
//		}
//		//vsLog.Info("Config: " + spew.Sdump(connConfig))
//		return &connConfig, err
//	}

func SetRunTimeConfig() *RunTimeConfig {
	vsLog.Info("Setting RunTimeConfig")
	RunConfig = &RunTimeConfig{}
	RunConfig.CoreDB = os.Getenv("CORE_DB")
	RunConfig.CoreDataBase = os.Getenv("CORE_DATABASE")
	RunConfig.ServiceName = os.Getenv("SERVICE_NAME")
	RunConfig.ConfigVersion = os.Getenv("CONFIG_VERSION")
	vsLog.Debug2("ConfigVersion: " + RunConfig.ConfigVersion)
	RunConfig.Company = os.Getenv("COMPANY")
	vsLog.Debug2("os.Getenv(COMPANY): " + os.Getenv("COMPANY"))
	RunConfig.RefreshSecret = os.Getenv("REFRESH_SECRET")
	RunConfig.ListenPort = os.Getenv("LISTEN_PORT")
	RunConfig.Port = os.Getenv("PORT")
	RunConfig.CfgString = fmt.Sprintf("%s:%s", RunConfig.ServiceName, RunConfig.ConfigVersion)
	os.Setenv("CONFIG_STRING", RunConfig.CfgString)
	os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
	RunConfig.AccessSecret = os.Getenv("ACCESS_SECRET")
	//vsLog.Debug2("RunConfig: " + spew.Sdump(RunConfig))
	return RunConfig
}
