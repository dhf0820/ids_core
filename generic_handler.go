package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	//"github.com/davecgh/go-spew/spew"

	//logging "github.com/dhf0820/uc_common/logging"

	common "github.com/dhf0820/uc_common"
	vsLog "github.com/dhf0820/vslog"
	"github.com/gorilla/mux"
	//m "github.com/dhf0820/ROIPrint/pkg/model"
)

// ####################################### Structures #######################################
// GenericResponse struct the resultant message being returned
type GenericResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ####################################### Response Functions #######################################
func WriteGenericResponse(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// switch status {
	// case 200:
	// 	w.WriteHeader(http.StatusOK)
	// case 400:
	// 	w.WriteHeader(http.StatusBadRequest)
	// case 401:
	// 	w.WriteHeader(http.StatusUnauthorized)
	// case 403:
	// 	w.WriteHeader(http.StatusForbidden)
	// case 404:
	// 	w.WriteHeader(http.StatusNotFound)
	// case 405:
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// case 500:
	// 	w.WriteHeader(http.StatusInternalServerError)
	// case 501:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// case 503:
	// 	w.WriteHeader(http.StatusServiceUnavailable)
	// }
	resp := GenericResponse{Status: status, Message: message}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		vsLog.Error(fmt.Sprintf("Error marshaling json: %v", err.Error()))
		return err
	}

	return nil
}

// ####################################### Route Handlers #######################################
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	version := fmt.Sprintf("OK: ids_core Version %s Environment: %s", os.Getenv("CodeVersion"), Env)
	WriteGenericResponse(w, 200, version)
	fmt.Println(version)
}

func processStatusReport(r *http.Request) *common.StatusReport {
	statusReport := common.StatusReport{}
	decoder := json.NewDecoder(r.Body).Decode(&statusReport)
	if decoder != nil {
		fmt.Printf("decoder: %s\n", decoder.Error())
	}
	if statusReport.Timestamp.String() == "0001-01-01 00:00:00 +0000 UTC" {
		statusReport.Timestamp = time.Now().UTC()
		statusReport.Nanotime = time.Now().UnixMilli()
	}
	return &statusReport
}

func createStatusReport(w http.ResponseWriter, r *http.Request) {
	statusReport := processStatusReport(r)
	msg := fmt.Sprintf("StatusReport Type: %s,  Status: %s, Time: %s, Comment: %s",
		statusReport.StatusType,
		statusReport.Status,
		statusReport.Timestamp,
		statusReport.Comment)
	WriteGenericResponse(w, 200, msg)
}

// func createLogEntry(w http.ResponseWriter, r *http.Request) {
// 	logMsg := processLogEntry(r)
// 	WriteGenericResponse(w, 200, logMsg.Message)
// }
// func processLogEntry(r *http.Request) *logging.Message {
// 	logMessage := logging.Message{}
// 	decoder := json.NewDecoder(r.Body).Decode(&logMessage)
// 	if decoder != nil {
// 		fmt.Printf("decoder: %s\n", decoder.Error())
// 	}
// 	return &logMessage
// }

func NewLogLevel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vsLog.Debug2(fmt.Sprintf("NewLogLevel: Params: %v", params))
	vsLog.Debug2(fmt.Sprintf("level: %v", params["level"]))
	level := params["level"]
	//m.SetLogLevel(lev
	logLevel := fmt.Sprintf("Log level now %s", level)
	WriteGenericResponse(w, 200, logLevel)
	fmt.Println(logLevel)
}

// func healthcheck(w http.ResponseWriter, r *http.Request) {
// 	session := ApiSession.Copy()
// 	defer session.Close()

// 	status, message := CheckHealth(session)
// 	WriteGenericResponse(w, status, message)
// }

//func HeaderStatus(code int)

// HandleFhirError extracts the acutal error code and message from err. It send the message to
// the genericResponse Writer providing the proper code and message. The result is a usable api message
func HandleFhirError(from string, w http.ResponseWriter, err error) {
	vsLog.Info(fmt.Sprintf("FHIR Error Handler: %v\n", err))
	code, message := extractErrorDetails(err.Error())
	err = WriteGenericResponse(w, code, message)
	if err != nil {
		vsLog.Error(fmt.Sprintf("Error writing FHIR ERROR response: %v", err))
	}
	vsLog.Debug2(fmt.Sprintf("%s failed with code: %d  msg: %s", from, code, message))
	return
}

// HandleError extracts the acutal error code and message from err. It send the message to
// the genericResponse Writer providing the proper code and message. The result is a usable api message
func HandleError(w http.ResponseWriter, from string, err error) {
	//vsLog.Infof("Generic-HandlerError-78: %v", err)
	code, message := extractErrorDetails(err.Error())
	message = fmt.Sprintf("%s", message)
	err = WriteGenericResponse(w, code, message)
	if err != nil {
		vsLog.Error(fmt.Sprintf("Error writing ERROR response: %v", err))
	}

	vsLog.Error(fmt.Sprintf("%s failed with code: %d  msg:%s", from, code, message))
	return
}

func extractErrorDetails(result string) (int, string) {
	s := strings.Split(result, "|")
	var statusCode int
	// if statusCode, err = strconv.ParseInt(s[0], 10, 64); err == nil {
	statusCode, err := strconv.Atoi(s[0])
	if err != nil {
		//vsLog.Warnf("extractErrorDetails error: %v\n", err)
		statusCode = 500
	}
	if len(s) > 1 {
		return statusCode, s[1]
	}

	return statusCode, result
}

// func ValidateSession(from string, w http.ResponseWriter, token string) *m.AuthSession {
// 	as, err := m.ValidateAuth(token)

// 	return as
// }
