package main

import (

	//INTERNAL
	//"crypto/rsa"
	"fmt"
	//"runtime/debug"
	"net/http"
	"os"

	//VERTISOFT
	"github.com/dhf0820/golangJWT"
	vsLog "github.com/dhf0820/vslog"

	//common "github.com/dhf0820/uc_common"
	//common "github.com/dhf0820/uc_common"
	//jwt "github.com/dhf0820/golangJWT"

	//EXTERNAL

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	gh "github.com/gorilla/handlers"

	//gm "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	// "strings"
	//"google.golang.org/grpc/credentials"
)

var CurrentResourceConfigs = []*fhir.ResourceConfig{}
var SysSummary = &SystemSummary{}
var CurrentSystem = &SystemConfig{}
var CurrentService = &ServiceConfig{}
var CurrentUser = &User{}
var JWTPayload = &golangJWT.UcPayload{}
var Version = "250712.0"
var CodeVersion = "250712.0"
var RunEnv = &RunTimeConfig{}
var ServiceName string
var Company string
var RunConfig *RunTimeConfig
var err = error(nil)
var Mode string
var Env string
var ResponseType string
var RawQuery string
var CoreURL string
var JwToken string

func main() {
	vsLog.Info("Starting IDS Core")

	envVersion := os.Getenv("CORE_ENV")
	if envVersion == "" {
		vsLog.Info("CORE_ENV not set using default .env")
		envVersion = ".env"
	} else {
		vsLog.Info("CORE_ENV set to: " + envVersion)
	}
	debugLevel := os.Getenv("DEBUG_LEVEL")
	if debugLevel == "" {
		vsLog.Info("DEBUG_LEVEL not set using default DEBUG1")
		debugLevel = "DEBUG1"
	}
	vsLog.Info("Code Version: " + CodeVersion)
	vsLog.Info("Version: " + Version)
	vsLog.Info("Set Debuglevel to " + debugLevel)
	vsLog.SetDebuglevel(debugLevel)
	vsLog.Debug2("Starting Company: " + os.Getenv("COMPANY"))
	// switch os.Getenv("MODE") {
	// case "yawl":
	// 	vsLog.Info("YawlLoading .env.core_yawl")
	// 	if err == nil {
	// 		Env = ".env.core_yawl"
	// 		Mode = "yawl"
	// 		debugLevel := os.Getenv("DEBUG_LEVEL")
	// 		if debugLevel == "" {
	// 			debugLevel = "DEBUG2"
	// 		}
	// 		vsLog.SetDebuglevel(debugLevel)
	// 	}
	// case "local":
	// 	vsLog.Info("Local Loading .env.core")
	// 	err = godotenv.Load("./." + envVersion)
	// 	if err == nil {
	// 		Env = ".env.core_go_test"
	// 		Mode = "go_test"
	// 		debugLevel := os.Getenv("DEBUG_LEVEL")
	// 		if debugLevel == "" {
	// 			debugLevel = "DEBUG2"
	// 		}
	// 		vsLog.Info("Set Debuglevel to " + debugLevel)
	// 		vsLog.SetDebuglevel(debugLevel)
	// 	} else {
	// 		vsLog.Info("core_go_test not found")
	// 	}
	// 	// if err == nil {
	// 	// 	Env = envVersion
	// 	// 	Mode = "local"
	// 	// 	debugLevel := os.Getenv("DEBUG_LEVEL")
	// 	// 	if debugLevel == "" {
	// 	// 		debugLevel = "DEBUG2"
	// 	// 	}
	// 	// 	vsLog.SetDebuglevel(debugLevel)
	// 	// }
	// case "go_test":
	// 	vsLog.Info("runing in go_test mode")
	// 	err = godotenv.Load("./.env.core_go_test")
	// 	if err == nil {
	// 		Env = ".env.core_go_test"
	// 		Mode = "go_test"
	// 		debugLevel := os.Getenv("DEBUG_LEVEL")
	// 		if debugLevel == "" {
	// 			debugLevel = "DEBUG2"
	// 		}
	// 		vsLog.Info("Set Debuglevel to " + debugLevel)
	// 		vsLog.SetDebuglevel(debugLevel)
	// 	} else {
	// 		vsLog.Info("core_go_test not found")
	// 	}
	// case "test":
	// 	vsLog.Info("Default Loading .env.core_test")
	// 	err = godotenv.Load("./.env.core_test")
	// 	if err == nil {
	// 		Env = ".env.core_test"
	// 		Mode = "test"
	// 		debugLevel := os.Getenv("DEBUG_LEVEL")
	// 		if debugLevel == "" {
	// 			debugLevel = "DEBUG3"
	// 		}
	// 		vsLog.Info("Set Debuglevel to " + debugLevel)
	// 	}
	// 	vsLog.SetDebuglevel(debugLevel)
	// default:
	// 	vsLog.Info("Default Loading: " + envVersion)
	// 	err = godotenv.Load(envVersion)
	// 	if err == nil {
	// 		Env = ".env"
	// 		Mode = "default"
	// 		debugLevel := os.Getenv("DEBUG_LEVEL")
	// 		if debugLevel == "" {
	// 			debugLevel = "DEBUG1"
	// 		}
	// 		vsLog.Info("Set Debuglevel to " + debugLevel)
	// 		vsLog.SetDebuglevel(debugLevel)
	// 	}
	// }

	// if err != nil {
	// 	vsLog.Error("Error loading environment: " + err.Error())
	// 	os.Exit(1)
	// }
	// os.Setenv("CodeVersion", Version)
	RunEnv = SetRunTimeConfig()
	vsLog.Debug2("RunEnv: " + spew.Sdump(RunEnv))
	if RunEnv.ServiceName == "" {
		vsLog.Error("Get environment: .env.core_test err: " + err.Error())
		os.Exit(1)
	}

	vsLog.Debug4("ids_core starting version: " + CodeVersion)
	if RunEnv.ServiceName == "" {
		vsLog.Warn("Environment is not set using: " + os.Getenv("MODE"))
		err := godotenv.Load(".env.core_test")
		if err != nil {
			vsLog.Error("Get environment: .env.core_test  err: " + err.Error())
			os.Exit(1)
		}
	}

	vsLog.Debug1(fmt.Sprintf("ServiceName: %s  ConfigVersion: %s", RunEnv.ServiceName, RunEnv.ConfigVersion))
	// err := godotenv.Load(".env.core")
	// if err != nil {
	// 	fmt.Printf("Main: Error getting environment: %v\n", err)

	// } else {
	Start(RunEnv)
	vsLog.Error("Start returned and should not have")
	// fmt.Printf("\n\n---Start restful handler\n")
	// restEp := service.GetMyEndpoint("restful_core")
	// //restAddress := restEp.Address

	// restAddress := fmt.Sprintf("%s:%s", restEp.Address, restEp.Port)
	// router := h.NewRouter()
	// logrus.Infof("----listening for restful requests at %s", restAddress)
	// mainErr := http.ListenAndServe(restAddress, router)
	// if mainErr != nil {
	// 	logrus.Errorf("Rest Startup error: %v", mainErr)
	// }
	//}
}

func Start(runEnv *RunTimeConfig) {
	//ar opts = []grpc.ServerOption{}
	// godotenv.Load("./.env_ids_core")
	// service.ServiceName = os.Getenv("SERVICE_NAME")
	// service.RunEnv = os.Getenv("CONFIG_VERSION")
	company := os.Getenv("COMPANY")
	vsLog.Debug2("Start: " + company)
	//vsLog.Debug2("Start: " + spew.Sdump(runenv))
	CurrentService, err = InitCore(RunEnv.ServiceName, RunEnv.ConfigVersion, RunEnv.Company) //TODO: get the env value from flag
	if err != nil {
		vsLog.Error("InitCore failed: " + err.Error())
		return
	}

	CoreURL = CurrentService.BaseURL
	vsLog.Debug2("CoreURL: " + CoreURL)
	OpenDB()
	cfg := GetConfig()
	CoreURL = cfg.BaseURL
	//ep := GetMyEndpoint("core")
	eps := GetMyEndpoints(cfg)
	for _, ep := range eps {
		if ep.Protocol == "http" {
			// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
			vsLog.Info("Starting restful ids_core service: " + CodeVersion)
			router := NewRouter()

			// // NO CORS Handler
			// restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
			// vsLog.Infof("\n\nids_core Service(no CORS) listening for restful requests at %s\n\n", restAddress)
			// err = http.ListenAndServe(restAddress, router)
			// vsLog.Errorf("This should not happen: Err = %s\n", err.Error())

			// rs/cors
			// //handler := cors.Default().Handler(router)
			// c := cors.New(cors.Options{
			// 	AllowedHeaders:[]string{"X-Requested-With", "Content-Type", "Fhir-System"},
			// 	AllowedOrigins: []string{"*"},

			// 	//AllowedOrigins: []string{"http://demo.universalcharts.com", "http://vuetest.universalcharts.com", "http://vuetest.universalcharts.com:8085"},
			// 	//AllowCredentials: false,
			// 	// Enable Debugging for testing, consider disabling in production
			// 	Debug: true,
			// })
			// //handler := cors.Default().Handler(router)
			// handler := c.Handler(router)
			// err := http.ListenAndServe(":40100", handler)
			// if err != nil {
			// 	vsLog.Errorf("rs/cors startup err: %s", err.Error())
			// 	return
			// }

			//Gorillia CORS WildCard
			credentialsOk := gh.AllowCredentials()
			headersOk := gh.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Fhir-System", "PatientId", "Authentication"})
			methodsOk := gh.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
			//originsOk := gh.AllowedOrigins([]string{"localhost", "192.168.1.142", ".universalcharts.com", ".ihids.com", ".vertisoft.com"})
			originsOk := gh.AllowedOrigins([]string{"http://gui.universalcharts.com", "http://vuetest.universalcharts.com", "http://demo.universalcharts.com", "localhost:8080", "http://localhost:8080", "localhost", "http://build.vertisoft.com"})
			//originsOk := gh.AllowedOrigins([]string{"*"})
			//originsOk := gh.AllowedOrigins([]string{".universalcharts.com", "10.0.0.2"})
			listenPort := RunEnv.ListenPort //os.Getenv("LISTEN_PORT")
			if listenPort != "" {
				vsLog.Info("ListenPort: " + listenPort)
				ep.Port = listenPort
			} else {
				vsLog.Info("ListenPort not set")
			}
			vsLog.Debug1(fmt.Sprintf("Starting Core CodeVersion: %s on port %s with environment [%s]", CodeVersion, ep.Port, Env))
			vsLog.Debug2("Calling ListenAndServe")
			err := http.ListenAndServe(":"+ep.Port, gh.CORS(originsOk, headersOk, methodsOk, credentialsOk)(router))
			if err != nil {
				vsLog.Error("ListenAndServe error: " + err.Error())
			}
			//fmt.Printf("main:166  --   Gorilla CORS stopped\n")
			//router := h.NewRouter()
			// restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
			// fmt.Printf("\n\n$$$ uc_fhir is listening for restful requests at %s\n\n", restAddress)
			// mainErr := http.ListenAndServe(restAddress, router )
			// if mainErr != nil {
			// 	vsLog.Errorf("Rest Startup error: %v", mainErr)
			// }

			// cors := gh.CORS(
			// 	gh.AllowedHeaders([]string{"content-type"}),
			// 	gh.AllowedOrigins([]string{"*"}),
			// 	gh.AllowCredentials(),
			// )
			//router := handlers.NewRouter()

			// r := mux.NewRouter()
			// r.HandleFunc("/api/rest/v1/auth/authorize", h.PostLogin).Methods("POST")
			// r.HandleFunc("/api/rest/v1/auth/login", h.Login).Methods("GET"

			// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)

			//fmt.Printf("main:102 -OriginsAllowed: %v\n", service.GetOriginsAllowed())

			//fmt.Printf("\n\n$$$ Core is listening for restful requests at %s:%s\n\n", "0.0.0.0", ep.Port)

			// //credentialsOk := gh.AllowCredentials()
			// headersOk := gh.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Fhir-System"})
			// methodsOk := gh.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
			// //originsOk := gh.AllowedOrigins([]string{"localhost", "192.168.1.142", ".universalcharts.com", ".ihids.com"})
			// // originsOk := gh.AllowedOrigins([]string{"*"})

			// log.Fatal(http.ListenAndServe(":"+ep.Port, gh.CORS(originsOk, headersOk, methodsOk, credentialsOk)(router)))
			// log.Fatal(http.ListenAndServe(":"+ep.Port, gh.CORS(originsOk, headersOk, methodsOk)(router)))
			// handler := cors.Default().Handler(router)
			// mainErr := http.ListenAndServe(restAddress, handler)

			// fmt.Printf("\n\n$$$ Core is listening for restful requests at %s\n\n", restAddress)
		}
		vsLog.Error("Start restful service returned and should not have")
	}
}

// func FillEnv() {
// 	CORE_DB := "mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/dev?retryWrites=true&w=majority"
// 	CORE_DATABASE :test
// 	SERVICE_NAME=ids_core
// 	CONFIG_VERSION=go_test
// 	COMPANY=test
// 	REFRESH_SECRET="Debbie loves me more"
// 				 // 12345678901234567890123456789012
// 	ACCESS_SECRET="I am so blessed Debbie loves me!"
// 	TOKEN_DURATION=15
// }

// func GetCurrentUserBaseUrl(system string) (*string, error) {
// 	if CurrentUser.CurrentLocalPatient.SysCfgId == system {
// 		return
// 	}
// 	CurrentSystem.
// }

func GetCurrentUser() (*User, error) {
	userId := JWTPayload.UserId

	//TODO: Remove after test
	//userId = "686e8a4aa00dc6346ce2a65d"
	user, err := GetUserById(userId)
	if err != nil {
		return nil, vsLog.Errorf(err.Error())
	}
	CurrentUser = user
	return CurrentUser, err

	// user, err := GetUserById(JWTPayload.UserId)
	// if err != nil {
	// 	return nil, vsLog.Errorf(err.Error())
	// }
	// user := User{}
	// user.FullName = "Debbie French"

	// CurrentUser = "dhfrench@vertisoft.com"
	// return CurrentUser, err
}
