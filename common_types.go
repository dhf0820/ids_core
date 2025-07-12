package main

import (
	"encoding/json"

	//"github.com/dgrijalva/jwt-go"
	jwt "github.com/dhf0820/golangJWT"
	//"github.com/dhf0820/uc_common"
	//"github.com/gorilla/schema"
	"time"

	fhir "github.com/dhf0820/fhir4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Used to set the patient that will be linked to any Resource Saved.

type AccessDetails struct {
	SessionId  string
	AccessUuid string
	UserId     string
	Token      *jwt.Token
	//ExpiresAt   int64
}

type ActivePatient struct {
	SysCfgId string `json:"sysCfgId" bson:"SysCfgId"`
	//Patient        fhir.Patient `json:"patient" bson:"patient"` // fhir record of the current active Patient
	PatientSummary string     `json:"patientSummary" bson:"patientSummary"`
	PatientId      string     `json:"patientId" bson:"patientId"`
	PatientName    string     `json:"patientName" bson:"patientName"`
	PatientDOB     string     `json:"patientDOB" bson:"patientDOB"`
	PatientMRN     string     `json:"patientMRN" bson:"patientMRN"`
	CreatedAt      *time.Time `json:"createdAt" bson:"createdAt"`
	//MRNSystem      string       `json:"mrnSystem`
}

type Attachment struct {
	Id          *string `bson:"id,omitempty" json:"id,omitempty"`
	ContentType *string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Language    *string `bson:"language,omitempty" json:"language,omitempty"`
	Data        *string `bson:"data,omitempty" json:"data,omitempty"`
	Url         *string `bson:"url,omitempty" json:"url,omitempty"`
	Title       *string `bson:"title,omitempty" json:"title,omitempty"`
}

type AuthSession struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Status         int                `json:"status" bson:"status"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserName       string             `json:"user_name" bson:"user_name"`
	FullName       string             `json:"fullName" bson:"fullName"`
	JWToken        string             `json:"jwToken" bson:"jwToken"`
	CurrentPatId   string             `json:"currentPatId" bson:"currentPatId"` //Keeps the current patient. If changes, start a new session, Delete old
	ExpiresAt      *time.Time         `json:"expiresAt" bson:"expiresAt"`
	CreatedAt      *time.Time         `json:"createdAt" bson:"createdAt"`
	LastAccessedAt *time.Time         `json:"lastAccessedAt" bson:"lastAccessedAt"`
	// May not want to include this in what gets returned to the user on login
	Connections []SessionConnection `json:"connections" bson:"connections"`
}
type BaseResMetaData struct {
	BaseResID    primitive.ObjectID `bson:"baseResID" json:"baseResID"`
	BaseResId    string             `bson:"baseResId" json:"baseResId"`
	CategoryText string             `bson:"categoryText" json:"categoryText"`
	TypeText     string             `bson:"typeText" json:"typeText"`
	Date         string             `bson:"date" json:"date"`
	PDFUrl       string             `bson:"pdfUrl" json:"pdfUrl"`
}

type BaseResCache struct {
	ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	QueryeID  primitive.ObjectID     `json:"queryID" bson:"queryID"`
	Meta      BaseResMetaData        `json:"metaData" bson:"metaData"`
	BaseRes   fhir.DocumentReference `json:"docRef" bson:"docRef"`
	CreatedAt time.Time              `json:"createdAt" bson:"createdAt"`
}

type BasicResource struct {
	Id string `json:"id"`
	//Text         fhir.Narrative `json:"text"`
	ResourceType string `json:"resourceType"`
}

type BundleCacheResponse struct {
	QueryId string `bson:"queryId" json:"queryId"`
}

type CacheBundle struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`          //This cache entry
	QueryID   primitive.ObjectID `bson:"queryID" json:"queryID"` //Id combining all pages for the query generated on first bundle response
	PageId    int                `bson:"pageId" json:"pageId"`   //Page of CacheBundle
	Header    *CacheHeader       `bson:"header" json:"header"`
	Bundle    *fhir.Bundle       `bson:"bundle" json:"bundle"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type CacheHeader struct {
	QueryID   primitive.ObjectID `bson:"queryID" json:"queryID"` //Id combining all pages for the query generated on first bundle response
	SystemCfg *SystemConfig      `bson:"systemCfg" json:"systemCfg,omitempty"`
	SystemId  string             `bson:"SystemId" json:"SystemId"`
	//CacheUrl            string  `bson:"cacheBase" json:"cacheBase,omitempty"`
	CacheUrl            string `bson:"cacheUrl" json:"cacheUrl"`
	GetResourceCacheUrl string `bson:"getResourceCacheUrl" json:"getResourceCacheUrl,omitempty"`
	GetBundleCacheUrl   string `bson:"getBundleCacheUrl" json:"getBundleCacheUrl,omitempty"`
	CacheStatusUrl      string `bson:"cacheStatusUrl" json:"cacheStatusUrl"`
	CachePageUrl        string `bson:"cachePageUrl" json:"cachePageUrl"`
	//Identifiers          []*KVData          `json:"identifiers" bson:"identifiers,omitempty"`
	UserId       string     `bson:"userId" json:"userId"`       //Global userId in ChartArchive
	PatientId    string     `bson:"patientId" json:"patientId"` //Patient this page of resources belongs to (not used caching patients)
	ResourceId   string     `bson:"resourceId" json:"resourceId,omitempty"`
	PageId       int        `bson:"pageId" json:"pageId"` //Part of this pages cache
	ResourceType string     `bson:"resourceType" json:"resourceType"`
	CreatedAt    *time.Time `bson:"createdAt" json:"createdAt"` //When was this cache created
	Query        string     `bson:"query" json:"query"`
	PageCount    int        `bson:"pageCount" json:"pageCount"`
	Status       string     `bson:"status" json:"status"`
}

type CachePageStatus struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"` //same as ResourceCache QueryId
	QueryId      primitive.ObjectID `bson:"queryId" json:"queryId"`
	Status       string             `bson:"status" json:"status"`
	TotalInCache int64              `bson:"totalInCache" json:"totalInCache"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	Query        string             `bson:"query" json:"query"`
}

type CacheResource interface {
	GetResourceType() string
	GetResourceId() string
	GetPatientId() string
}

type CacheSavePayload struct {
	ResourceType string       `bson:"resourceType" json:"resourceType"`
	PageNum      int          `bson:"pageNum" json:"pageNum"`
	LastPage     bool         `bson:"lastPage" json:"lastPage"`
	QueryId      string       `bson:"queryId" json:"queryId"`
	Query        string       `bson:"query" json:"query"`
	Option       string       `bson:"option" json:"option"` // "ALL", "BUNDLE", "ENTRIES"
	Bundle       *fhir.Bundle `bson:"bundle" json:"bundle"`
}

type CacheStatus struct {
	ID primitive.ObjectID `bson:"_id" json:"id"` //same as ResourceCache QueryId
	//QueryId      string             `bson:"queryId" json:"queryId"`
	Status string `bson:"status" json:"status"`
	//BundleSize   int       `bson:"bundleSize,omitempty" json:"bundleSize,omitempty"`
	PageSize       int       `bson:"pageSize" json:"pageSize"`
	PageCount      int       `bson:"pageCount" json:"pageCount"` // number of pages in query results, last may not be fill size
	TotalInCache   int       `bson:"totalInCache" json:"totalInCache"`
	TotalResources int64     `bson:"totalResources" json:"totalResources"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	Query          string    `bson:"query,omitEmpty" json:"query",omitEmpty`
}
type ConditionCache struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	QueryID   primitive.ObjectID `json:"queryID" bson:"queryID"`
	Meta      ConditionMetaData  `json:"metaData" bson:"metaData"`
	Procedure fhir.Procedure     `json:"procedure" bson:"procedure"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type ConditionMetaData struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id             string             `bson:"id" json:"id"`
	CategoryText   string             `bson:"categoryText" json:"categoryText"`
	CategoryCode   string             `bson:"code" json:"code"`
	CategorySystem string             `bson:"categorySystem" json:"categorySystem"`
	CodeText       string             `bson:"codeText" json:"codeText"`
	CodeSystem     string             `bson:"codeSystem" json:"codeSystem"`
	CodeCode       string             `bson:"codeCode" json:"codeCode"`
	Issued         string             `bson:"issued" json:"issued"`
	Patient        string             `bson:"patient" json:"patient"`
	//TypeText     string             `bson:"typeText" json:"typeText"`
	DocumentType string `bson:"documentType" json:"documentType"`
	NameText     string `bson:"nameText" json:"nameText"`
	ServiceDate  string `bson:"serviceDate" json:"serviceDate"`
	//HTMLImage    string 			  `bson:"htmlImage,omitempty" json:"htmlImage,omitempty"`
	//XMLImage     string 			  `bson:"xmlImage" json:"xmlImage,omitempty"`
	//PDFImage     string  			  `bson:"pdfImage" json:"pdfImage"`
	Encounter string `bson:"encounter" json:"encounter"`
	Related   string `bson:"related,omitempty" json:"related,omitempty"`
	Status    string `bson:"status" json:"status"`
}

type ConnectInfo struct {
	//ID    string `json:"id" bson:"id,omitempty"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Connector struct {
	Id   string   `bson:"id" json:"id"`
	Name string   `bson:"name" json:"name"`
	URL  string   `bson:"url" json:"url"`
	Data []KVData `bson:"data" json:"data"`
}

type ConnectorConfig struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name          string             `json:"name"`
	Version       string             `json:"version"`
	Label         string             `json:"label" bson:"label"`
	Credentials   string             `json:"credentials" bson:"credentials"`
	Data          []*KVData          `json:"data" bson:"data"`
	Identifiers   []*KVData          `json:"identifiers" bson:"identifiers"`
	CacheUrl      string             `json:"cacheUrl" bson:"cacheUrl"`
	CreatedAt     *time.Time         `json:"createdAt" bson:"created_at"`
	UpdatesAt     *time.Time         `json:"updatedAt" bson:"updated_at"`
	Returns       string             `json:"returns" bson:"returns"`
	HostUrl       string             `json:"hostUrl" bson:"hostUrl"` // Filled in dynamically on each call
	ConnectorType string             `json:"connectorType" bson:"connectorType"`
	AcceptValue   string             `json:"acceptValue" bson:"acceptValue"`
	URL           string             `json:"url" bson:"url"` // Get from ENV.
	//TODO: ConnectorConfig remove uneeded fields
	// DeployMode    string             `json:"deployMode" bson:"deployMode"`  //Put in Env
	// QueueName     string             `json:"queueName" bson:"queueName"` // put in Env
	// Enabled       string             `json:"enabled"  bson:"enabled"` // ?//true/false
	// CertName      string             `json:"certName" bson:"certName"`
	// TlsMode       string             `json:"tlsMode" bson:"tlsMode"`
	// Protocol      string             `json:"protocol" bson:"protocol"` //?
	// Address       string             `json:"address" bson:"address"` //?
}

type ConnectorPayload struct {
	Facility        *Facility        `json:"facility"`
	System          *SystemConfig    `json:"system"`
	ConnectorConfig *ConnectorConfig `json:"connectorConfig"`
	SavePayload     *SavePayload     `json:"savePayload"`
	FhirAuthToken   string           `json:"fhirAuthToken"` // for the Destination System
	DeployMode      string           `json:"deployMode"`    // Container(K8S or Docker) or Local
	AcceptValue     string           `json:"acceptValue"`
	Enabled         bool             `json:"enabled"`
	HostUrl         string           `json:"hostUrl"`
	CoreUrl         string           `json:"coreUrl"`
	BaseCoreUrl     string           `json:"baseCoreUrl"`
	CacheUrl        string           `json:"cacheUrl"`
	Data            []*KVData        `json:"data"`
}

type ConnectorResponse struct {
	Status    int              `json:"status"`
	Message   string           `json:"message"`
	Connector *ConnectorConfig `json:"connector"`
}

type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Customer struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CustomerName string             `json:"customerName" bson:"customer_name"`
	ContactName  string             `json:"contactName" bson:"contact_name"`
	ContactEmail string             `json:"contactEmail" bson:"contact_email"`
	ContactPhone string             `json:"contactPhone" bson:"contact_phone"`
	Code         string             `json:"code"`
	Facility     string             `json:"facility"`
	//UserId 		string 	`json:"user_id"`
}
type DataConnector struct {
	Name       string    `json:"name" bson:"name"`
	DbName     string    `json:"dbName" bson:"dbName"`
	Server     string    `json:"server" bson:"server"`
	User       string    `json:"user" bson:"user"`
	Password   string    `json:"password" bson:"password"`
	Database   string    `json:"database" bson:"database"`
	Collection string    `json:"collection" bson:"collection"`
	Fields     []*KVData `json:"fields" bson:"fields"`
}

type DiagReptCache struct {
	//ID        primitive.ObjectID    `json:"_id" bson:"_id"`
	QueryID   primitive.ObjectID    `json:"queryID" bson:"queryID"`
	Meta      DiagReptMetaData      `json:"metaData" bson:"metaData"`
	DiagRept  fhir.DiagnosticReport `json:"diagRept" bson:"diagRept"`
	CreatedAt time.Time             `json:"createdAt" bson:"createdAt"`
}
type DiagReptMetaData struct {
	ID             primitive.ObjectID `bson:"diagReptID" json:"diagReptID"`
	Id             string             `bson:"id" json:"id"`
	Subject        string             `bson:"subject" json:"subject"`
	CategoryText   string             `bson:"categoryText" json:"categoryText"`
	CategoryCode   string             `bson:"categoryCode" json:"categoryCode"`
	CategorySystem string             `bson:"categorySystem" json:"categorySystem"`
	CodeText       string             `bson:"codeText" json:"codeText"`
	CodeCode       string             `bson:"codeCode" json:"codeCode"`
	CodeSystem     string             `bson:"codeSystem" json:"codeSystem"`
	Issued         string             `bson:"issued" json:"issued"`
	PDFImage       string             `bson:"pdfImage" json:"pdfImage"`
	Status         string             `bson:"status" json:"status"`
	Encounter      string             `bson:"encounter" json:"encounter,omitempty"`
}

type DocRefCache struct {
	ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	QueryID   primitive.ObjectID     `json:"queryID" bson:"queryID"`
	Meta      DocRefMetaData         `json:"metaData" bson:"metaData"`
	DocRef    fhir.DocumentReference `json:"docRef" bson:"docRef"`
	CreatedAt time.Time              `json:"createdAt" bson:"createdAt"`
}
type DocRefMetaData struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id           string             `bson:"id" json:"id"`
	CategoryText string             `bson:"categoryText" json:"categoryText"`
	//TypeText     string             `bson:"typeText" json:"typeText"`
	DocumentType string `bson:"documentType" json:"documentType"`
	NameText     string `bson:"nameText" json:"nameText"`
	ServiceDate  string `bson:"serviceDate" json:"serviceDate"`
	HTMLImage    string `bson:"htmlImage,omitempty" json:"htmlImage,omitempty"`
	XMLImage     string `bson:"xmlImage" json:"xmlImage,omitempty"`
	PDFImage     string `bson:"pdfImage" json:"pdfImage"`
	Encounter    string `bson:"encounter" json:"encounter"`
	Related      string `bson:"related,omitempty" json:"related,omitempty"`
	Status       string `bson:"status" json:"status"`
}

type EndPoint struct { // Replaces BaseUrl
	Name        string //internal name
	Label       string `json:"label"`
	Scope       string `json:"scope,omitempty" bson:"scope"`
	Protocol    string `json:"protocol" bson:"protocol"` // grpc or amqp
	Address     string `json:"address" bson:"address"`   //How do I get to this service
	Port        string `json:"port"`
	Credentials string `json:"credentials" bson:"credentials"`
	CertName    string `json:"certname" bson:"cert_name"`
	TLSMode     string `json:"tlsmode" bson:"tls_mode"`
	DeployMode  string `json:"deploymode" bson:"deploy_mode"`
}

// type ResourceGrid struct {
// 	ID           primitive.ObjectID `json:"id" bson:"_id"`
// 	SystemId     string             `json:"SystemId" bson:"SystemId"`
// 	ResourceType string             `json:"resourceType" bson:"resourceType"`
// 	HeaderData   []HeaderData       `json:"headerData" bson:"headerData"`
// 	//FieldExtractors []FieldExtractor   `json:"fieldExtractor" bson:"fieldExtractor"` // How to find each data element
// 	EntryResults []EntryResult `json:"EntryResults" bson:"entryResults"` // one for each entry
// }

// Elements = individual fields in a FHIR record.
// Entry = individual records of the returned resource
// The data for a single element of a single entry in a fhir bundle .eg. patient.Name
type ElementData struct {
	Name      string      `json:"name" bson:"name"`
	FieldType string      `json:"fieldType" bson:"fieldType"`
	Value     interface{} `json:"value" bson:"value"`
	Url       string      `json:"url,omitempty" bson:"url"`
}

// All the selected dataElements from an single entry(line) in a fhir bundle .eg.  index0 patient.Id, index1 patient.Name
type EntryResult struct {
	Elements []ElementData `json:"elements" bson:"elements"` //all the individual data elements in a single bundle entry
}

// ALL of the selected dataElements for  each entry in a fhir bundle
type EntryResults struct {
	Results []EntryResult `json:"results" bson:"results"` // the index is the entry number in the bundle
}

type Facility struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name" bson:"name"`
	DisplayName    string             `json:"displayName" bson:"displayName"`
	Description    string             `json:"description" bson:"description"`
	Lat            string             `json:"lat,omitempty" bson:"lat"`
	Lon            string             `json:"lon,omitempty" bson:"lon"`
	IconLogo       string             `json:"iconLog,omitempty" bson:"iconLogo"`
	Systems        []*SystemSummary   `json:"systems" bson:"systems"`
	Classification string             `json:"classification" bson:"classification"` //Practice or hospital
	CoreUrl        string             `json:"coreUrl" bson:"coreUrl"`
	//Members        []*UserSummary     `json:"members,omitempty" bson:"members"`
}

type FhirConfig struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FacilityId  primitive.ObjectID `json:"facilityId" bson:"facilityId"`
	Name        string             `json:"name"`
	DisplayName string             `json:"displayName" bson:"display_name"`
	URL         string             `json:"url,omitempty" bson:"url"`   // URL of the UniversalCharts handler
	ActualURL   string             `json:"actualURL" bson:"actualURL"` // URL of the fhir server we call
	EndPoints   []*EndPoint        `json:"endPoints,omitempty"`        //All System endpoints this connector needs to talk to
	FhirInfo    []*KVData          `json:"fhirInfo" bson:"fhir_info"`
	Priority    string             `json:"priority,omitempty" bson:"priority,omitempty"`
	Enabled     string             `json:"enabled,omitempty"`                       //true/false
	FhirFields  []*FhirField       `json:"fhirFields,omitempty" bson:"fhir_fields"` // Fhir data extract and display information
	Fields      []*Field           `json:"Fields,omitempty" bson:"fields"`          // GUI information to manage the configuration
	Data        []*KVData          `json:"data"`
	CreatedAt   *time.Time         `json:"createdAt" bson:"created_at"`
	UpdatesAt   *time.Time         `json:"updatedAt" bson:"updated_at"`
}

type FhirExtractor struct {
	FieldName     string         `json:"fieldName" bson:"fieldName"`
	Values        []KVData       `json:"values,omitempty" bson:"values"`
	Value         string         `json:"value,omitempty" bson:"value"`
	Path          []KVData       `json:"path" bson:"path"`
	FieldMatchers []FieldMatcher `json:"fieldMatchers" bson:"fieldMatchers"`
}

type FhirField struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Key         string `json:"key"`                              //oid/system/other identifier type
	Value       string `json:"value"`                            // Value of specific key type
	DisplayName string `json:"displayName" bson:"display_name"`  // FieldName for displaying field. Actual name or column name for table
	Extract     string `json:"extract omitempty" bson:"extract"` //Go ufnction Name or js:methodName
	Required    string `json:"required" bson:"requied"`          // blank = false or "true" = true
}

type FhirResults struct {
	ResourceType      string                  `bson:"resourceType" json:"resourceType"`
	ResourceId        string                  `bson:"resourceId" json:"resourceId"`
	Resource          json.RawMessage         `bson:"resource" json:"resource,omitempty"`
	ActivePatient     *ActivePatient          `bson:"activePatient" json:"activePatient,omitempty"`
	Binary            *fhir.Binary            `bson:"binary,omitempty" json:"binary,omitempty"`
	Condition         *fhir.Condition         `bson:"condition,omitempty" json:"condition,omitempty"`
	DiagnosticReport  *fhir.DiagnosticReport  `bson:"diagnosticReport,omitempty" json:"diagnosticReport,omitempty"`
	DocumentReference *fhir.DocumentReference `bson:"documentReference,omitempty" json:"documentReference,omitempty"`
	Encounter         *fhir.Encounter         `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Immunization      *fhir.Immunization      `bson:"immunization,omitempty" json:"immunization,omitempty"`
	Observation       *fhir.Observation       `bson:"observation,omitempty" json:"observation,omitempty"`
	OperationOutcome  *fhir.OperationOutcome  `bson:"operationOutcome,omitempty" json:"operationOutcome,omitempty"`
	Patient           *fhir.Patient           `bson:"patient,omitempty" json:"patient,omitempty"`
	Procedure         *fhir.Procedure         `bson:"procedure,omitempty" json:"procedure,omitempty"`
	ResourceSummary   *ResourceSummary        `bson:"resourceSummary,omitempty" json:"resourceSummary,omitempty"`
	//ReourceDisplay    *ResourceDisplay        `bson:"resourceDisplay" json:"resourceDisplay"`
	Bundle fhir.Bundle `json:"bundle" bson:"bundle"`
}

type Field struct {
	Name         string `json:"name"`
	Label        string `json:"label"`
	Default      string `json:"default"`
	Value        string `json:"value"`
	DisplayValue string `json:"display_value" bson:"display_value"`
	Required     string `json:"required omitempty"`
	UserVisible  string `json:"user_visible omitempty" bson:"user_visible"`
	IsNameValue  string `json:"is_name_value" bson:"is_name_value"`
	Sensitive    string `json:"sensitive"`
}

type FieldMatcher struct {
	FieldName      string            `json:"fieldName" bson:"fientName"`
	MatchQualifier string            `json:"qualifier" bson:"qualifier"` //=
	Values         []KVData          `json:"values" bson:"values"`
	MatchValues    []FieldMatchValue `json:"matchValues" bson:"matchValues"`
	MatchPullValue string            `json:"pullValue" bson:"pullValue"`
}

type FieldMatchValue struct {
	Name string `json:"name" bson:"name"`
}

type FinishCachePayload struct {
	ResourceType string `bson:"resourceType" json:"resourceType"`
	PageNum      int    `bson:"pageNum" json:"pageNum"`
	QueryId      string `bson:"queryId" json:"queryId"`
	Query        string `bson:"query" json:"query"`
	OnPage       int    `bson:"onPage" json:"onPage"`
	PageSize     int    `bson:"pageSize" json:"pagSize"`
}

type FinishCacheResources struct {
	ResourceType string `bson:"resourceType" json:"resourceType"`
	QueryId      string `bson:"queryId" json:"queryId"`
}

type HeaderData struct {
	FieldName string `json:"fieldName" bson:"fieldName"`
	Title     string `json:"title" bson:"title"`
	Order     int    `json:"order,omitEmpty" bson:"order"`
	Width     string `json:"width" bson:"width"`
	Type      string `json:"type" bson:"type"` // Visible, Data
	//FhirElement string `json:"fhirElement" bson:"fhirElement"`
}

//	type UserResponse struct {
//		ID         primitive.ObjectID `json:"id"`
//		Token      string             `json:"token"`
//		Status     int                `json:"status"`
//		Message    string             `json:"message"`
//		UserName   string             `json:"userName"`
//		FullName   string             `json:"fullName"`
//		Role       string             `json:"role"`
//		Locals  []*Practice        `json:"practices"`
//		Remotes []*Facility        `json:"facilities"`
//	}

// type ResourceElement struct {
// 	Results []ResourceElementResult `json:"results" bson:"results"` // one for each data element
// }

type HippaLog struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	UserId          string             `json:"userId" bson:"userId"`
	PatientId       string             `json:"patientId" bson:"patientId"`
	SrcPatientId    string             `json:"srcPatientId" bson:"srcPatientId"`
	DestPatientId   string             `json:"destPatientId" bson:"destPatientId"`
	SrcPatientName  string             `json:"srcPatientName" bson:"srcPatientName"`
	DestPatientName string             `json:"destPatientName" bson:"destPatientName"`
	SrcPatientDOB   string             `json:"srcPatientDOB" bson:"srcPatientDOB"`
	DestPatientDOB  string             `json:"destPatientDOB" bson:"destPatientDOB"`
	ResourceType    string             `json:"resourceType" bson:"resourceType"`
	ResourceId      string             `json:"resourceId" bson:"resourceId"`
	SrcResourceId   string             `json:"srcResourceId" bson:"srcResourceId"`
	DestResourceId  string             `json:"destResourceId" bson:"destResourceId"`
	SystemId        string             `json:"SystemId" bson:"SystemId"`
	SrcSystemId     string             `json:"srcSystemId" bson:"srcSystemId"`
	DestSystemId    string             `json:"destSystemId" bson:"destSystemId"`
	LogType         string             `json:"logType" bson:"logType"` //view, save, over_ride
	LogTime         time.Time          `json:"logTime" bson:"logTime"`
	LogMessage      string             `json:"logMessage" bson:"logMessage"`
}

type ImmunizationCache struct {
	ID           primitive.ObjectID   `json:"_id" bson:"_id"`
	QueryID      primitive.ObjectID   `json:"queryID" bson:"queryID"`
	Meta         ImmunizationMetaData `json:"metaData" bson:"metaData"`
	Immunization fhir.Immunization    `json:"immunization" bson:"immunization"`
	CreatedAt    time.Time            `json:"createdAt" bson:"createdAt"`
}
type ImmunizationMetaData struct {
	ID             primitive.ObjectID `bson:"immunID" json:"immunID"`
	Id             string             `bson:"id" json:"id"`
	Subject        string             `bson:"subject" json:"subject"`
	CategoryText   string             `bson:"categoryText" json:"categoryText"`
	CategoryCode   string             `bson:"categoryCode" json:"categoryCode"`
	CategorySystem string             `bson:"categorySystem" json:"categorySystem"`
	CodeText       string             `bson:"codeText" json:"codeText"`
	CodeCode       string             `bson:"codeCode" json:"codeCode"`
	CodeSystem     string             `bson:"codeSystem" json:"codeSystem"`
	Issued         string             `bson:"issued" json:"issued"`
	PDFImage       string             `bson:"pdfImage" json:"pdfImage"`
	Status         string             `bson:"status" json:"status"`
	Encounter      string             `bson:"encounter" json:"encounter,omitempty"`
}
type InitialResourceResponse struct {
	Status        int                    `json:"status"`
	Message       string                 `json:"message"`
	OpOutcome     *fhir.OperationOutcome `json:"opOutcome,omitempty"`
	OriginalQuery string                 `json:"originalQuery"`
	ResourceType  string                 `json:"resourceType"`
	PageNumber    int                    `json:"pageNumber"`
	TotalPages    int64                  `json:"totalPages,omitempty"`
	CountInPage   int                    `json:"countInPage,omitempty"`
	BundleId      string                 `json:"bundleId,omitempty"`
	QueryId       string                 `json:"queryId"`
	Header        *CacheHeader           `json:"header,omitempty"`
}

type KVData struct {
	Name  string
	Value string
}

// A user logged into UC receives an AuthSession

type LoginFilter struct {
	UserName string `schema:"userName"`
	Password string `schema:"password"`
	//PracticeId string `schema:"practiceId"`
}

type LoginResponse struct {
	ID                   primitive.ObjectID `json:"id"`
	Token                string             `json:"token"`
	Status               int                `json:"status"`
	Message              string             `json:"message"`
	UserName             string             `json:"userName"`
	FullName             string             `json:"fullName"`
	Role                 string             `json:"role"`
	SessionId            string             `json:"sessionId"`
	Remotes              []*Facility        `json:"remotes"` // Remotes have Classification of REMOTE(Read only) or LOCAL (Can save to)
	Locals               []*Facility        `json:"locals"`
	CurrentLocalPatient  *ActivePatient     `json:"currentLocalPatient" bson:"currentLocalPatient"` //Current local patient
	CurrentRemotePatient *ActivePatient     `json:"currentRemotePatient" bson:"currentRemotePatient"`
	Session              AuthSession        `json:"session,omitempty"`
	BaseUrl              string             `json:"baseUrl"`
	ResourceConfig       *ResourceConfig    `json:"resourceConfig,omitempty"`
}

type Narrative struct {
	Id  *string `bson:"id,omitempty" json:"id,omitempty"` // if nil it is generated.
	Div string  `bson:"div" json:"div"`
}

type ObservationCache struct {
	ID          primitive.ObjectID  `json:"_id" bson:"_id"`
	QueryID     primitive.ObjectID  `json:"queryID" bson:"queryID"`
	Meta        ObservationMetaData `json:"metaData" bson:"metaData"`
	Observation fhir.Observation    `json:"observation" bson:"observation"`
	CreatedAt   time.Time           `json:"createdAt" bson:"createdAt"`
}

type ObservationMetaData struct {
	//ID primitive.ObjectID `bson:"_id" json:"_id"`
	Id string `bson:"id" json:"id"`
	//Subject      string `bson:"subject" json:"subject"`
	CategoryText string `bson:"categoryText" json:"categoryText"`
	CodeText     string `bson:"codeText" json:"codeText"`
	Issued       string `bson:"issued" json:"issued"`
	PDFImage     string `bson:"pdfImage" json:"pdfImage"`
	Status       string `bson:"status" json:"status"`
	Encounter    string `bson:"encounter" json:"encounter,omitempty"`
	Value        string `bson:"value" json:"value"`
}

type PatientIndexFields struct {
	PatientID   primitive.ObjectID `json:"patientID" bson:"patientID"`
	PatientName fhir.HumanName     `json:"patientName" bson:"patientName"`
}

type PatientCache struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	QueryID   primitive.ObjectID `json:"queryID" bson:"queryID"`
	Meta      PatientMetaData    `json:"patientMetaData" bson:"patientMetaData"`
	Patient   fhir.Patient       `json:"patient" bson:"patient"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

// A limited information of the patient.
type PatientSummary struct {
	ID         primitive.ObjectID `json:"id" bson:"id"`
	FullName   string             `json:"fullName" bson:"fullName"`
	LastAccess time.Time          `json:"lastAccess" bson:"lastAccess"`
}

type PatientMetaData struct {
	PatId      string `bson:"patId" json:"patId"`
	FamilyName string `bson:"familyName" json:"familyName"`
	GivenName  string `bson:"givenName" json:"givenName"`
	MiddleName string `bson:"middleName" json:"middleName"`
	NameText   string `bson:"nameText" json:"nameText"`
	BirthDate  string `bson:"birthDate" json:"birthDate"`
	Gender     string `bson:"gender" json:"gender"`
	MedRecNum  string `bson:"medRecNum" json:"medRecNum"`
	SSN        string `bson:"ssn" json:"ssn"`
}

type ProcedureCache struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	QueryID   primitive.ObjectID `json:"queryID" bson:"queryID"`
	Meta      ProcedureMetaData  `json:"metaData" bson:"metaData"`
	Procedure fhir.Procedure     `json:"procedure" bson:"procedure"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type ProcedureMetaData struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id             string             `bson:"id" json:"id"`
	CategoryText   string             `bson:"categoryText" json:"categoryText"`
	CategoryCode   string             `bson:"code" json:"code"`
	CategorySystem string             `bson:"categorySystem" json:"categorySystem"`
	CodeText       string             `bson:"codeText" json:"codeText"`
	CodeSystem     string             `bson:"codeSystem" json:"codeSystem"`
	CodeCode       string             `bson:"codeCode" json:"codeCode"`
	Issued         string             `bson:"issued" json:"issued"`
	Patient        string             `bson:"patient" json:"patient"`
	//TypeText     string             `bson:"typeText" json:"typeText"`
	DocumentType string `bson:"documentType" json:"documentType"`
	NameText     string `bson:"nameText" json:"nameText"`
	ServiceDate  string `bson:"serviceDate" json:"serviceDate"`
	//HTMLImage    string `bson:"htmlImage,omitempty" json:"htmlImage,omitempty"`
	//XMLImage     string `bson:"xmlImage" json:"xmlImage,omitempty"`
	//PDFImage     string `bson:"pdfImage" json:"pdfImage"`
	Encounter string `bson:"encounter" json:"encounter"`
	Related   string `bson:"related,omitempty" json:"related,omitempty"`
	Status    string `bson:"status" json:"status"`
}

// type SessionHistory struct {
// 	FacilityId		primitive.ObjectID	`json:"facilityId" bson:"facilityId"`
// 	SystemId		primitive.ObjectID	`json:"SystemId" bson:"SystemId"`
// 	//Last X patients the user has selected
// 	PatientHistory	[]PatientSummary	`json:"patientHistory" bson:"patientHistory"`
// 	Token 			string				`json:"token" bson:"token"`
// }

// A user logged into UC receives an AuthSession

type ReceiptQueue struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	SourceName   string             `json:"sourceName" bson:"sourceName"`
	SourceID     primitive.ObjectID `json:"sourceID" bson:"sourceID"`
	ResourceData *ResourceData      `json:"resourceData" bson:"resourceData"`
}

// resourceCache is the type sent to CacheServer to cache one resource from a GetResource
type ResourceCache struct {
	ID                primitive.ObjectID      `bson:"_id" json:"id"`
	Header            *CacheHeader            `bson:"header" json:"header"`
	ResourceHeader    *ResourceHeader         `bson:"resourceHeader" json:"resourceHeader"`
	ResourceType      string                  `bson:"resourceType" json:"resourceType"`
	Resource          json.RawMessage         `bson:"resource" json:"resource"`
	Patient           *fhir.Patient           `bson:"patient" json:"patient"`
	Condition         *fhir.Condition         `bson:"condition" json:"condition"`
	DocumentReference *fhir.DocumentReference `bson:"docRef" json:"docRef"`
	Encounter         *fhir.Encounter         `bson:"encounter" json:"encounter"`
	Immunization      *fhir.Immunization      `bson:"immunization" json:"immunization"`
	Observation       *fhir.Observation       `bson:"observation" json:"observation"`
	DiagnosticReport  *fhir.DiagnosticReport  `bson:"diagnosticReport" json:"diagnosticReport"`
	CreatedAt         time.Time               `bson:"createdAt" json:"createdAt"`
	//TODO: Addd rest of resources to ResourceCache
}

type ResourceCacheResponse struct {
	CacheId string `bson:"cacheId" json:"cacheId"`
}

type ResourceCacheStatus struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"` //same as ResourceCache QueryId
	QueryId     string             `bson:"queryId" json:"queryId"`
	Status      string             `bson:"status" json:"status"`
	Count       int                `bson:"count" json:"count"` // the last requested number of Resources
	LastUpdated time.Time          `bson:"lastUpdated" json:"lastUpdated"`
	Query       string             `bson:"query" json:"query"`
}

type ResourceConfig struct {
	ID         primitive.ObjectID    `json:"_id" bson:"_id"`
	Name       string                `json:"name" bson:"name"`             // fhir resource Name (Observation, Patient,...)
	Element    string                `json:"element" bson:"element"`       // The element of the patient struct this pertains to
	Filters    []fhir.ResourceFilter `json:"filters" bson:"filters"`       // The filters to apply to the resource
	Extracts   []FhirExtractor       `json:"extracts" bson:"extracts"`     // The data to extract from the resource
	Url        string                `json:"url" bson:"url"`               // Normal url to use in the filter
	FilterType string                `json:"filterType" bson:"filterType"` //URL or Internal
	//SummaryType string             `json:"summaryType" bson:"summaryType"`//For grid display, Not used yet
	//DisplayType string             `json:"displayType" bson:"displayType"`/// For Displaying information not used yet
	Abilities []string `json:"abilities,omitempty" bson:"abilities"`
}

type ResourceData struct {
	ResourceHeader *ResourceHeader `bson:"resourceHeader" json:"resourceHeader"`
	ResourceType   string          `bson:"resourceType" json:"resourceType"`
	ResourceId     string          `bson:"resourceId" json:"resourceId"`
	//ResourceGrid   ResourceGrid    `bson:"resourceGrid" json:"resourceGrid"`
	EntrySummary EntryResult `bson:"entrySummary" json:"entrySummary"`

	Resource          json.RawMessage         `bson:"resource" json:"resource,omitempty"`
	ActivePatient     *ActivePatient          `bson:"activePatient" json:"activePatient,omitempty"`
	Binary            *fhir.Binary            `bson:"binary" json:"binary,omitempty"`
	Condition         *fhir.Condition         `bson:"condition" json:"condition,omitempty"`
	DiagnosticReport  *fhir.DiagnosticReport  `bson:"diagnosticReport" json:"diagnosticReport,omitempty"`
	DocumentReference *fhir.DocumentReference `bson:"documentReference" json:"documentReference,omitempty"`
	Encounter         *fhir.Encounter         `bson:"encounter" json:"encounter,omitempty"`
	Immunization      *fhir.Immunization      `bson:"immunization" json:"immunization,omitempty"`
	Observation       *fhir.Observation       `bson:"observation" json:"observation,omitempty"`
	OperationOutcome  *fhir.OperationOutcome  `json:"operationOutcome,omitempty"`
	Patient           *fhir.Patient           `bson:"patient" json:"patient,omitempty"`
	Procedure         *fhir.Procedure         `bson:"procedure" json:"procedure,omitempty"`
}

type ResourceDisplay struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	SystemId        string             `json:"SystemId" bson:"SystemId"`
	ResourceType    string             `json:"resourceType" bson:"resourceType"`
	FieldExtractors []FhirExtractor    `json:"fieldExtractor" bson:"fieldExtractor"` // How to find each data element
	FieldResults    []ElementData      `json:"fieldResults" bson:"fieldResults"`     // one for each entry
}

type ResourceExtractor struct {
	Element         string `json:"element" bson:"element"`
	IdentifierType  string `json:"identifierType" bson:"identifierType"`
	IdentifierValue string `json:"identifierValue" bson:"identifierValue"`
	System          string `json:"system" bson:"system"`
}

// type ResourceFilter struct {
// 	Name           string   `json:"name" bson:"name"`
// 	Display        string   `json:"display" bson:"display"`
// 	Description    string   `json:"description,omitempty" bson:"description"`
// 	FullName       string   `json:"fullName" bson:"fullName"`
// 	Values         []KVData `json:"values,omitempty" bson:"values"` // Array of search possibility instead of user entered. Works for Category in Observation
// 	System         string   `json:"system" bson:"system"`
// 	RetrieveMethod string   `json:"retrieveMethod" bson:"retrieveMethod"` //Direct, OID, Function,
// 	Value          string   `json:"value" bson:"value"`
// }

type ResourceForm struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	SystemId     string             `json:"SystemId" bson:"SystemId"`
	ResourceType string             `json:"resourceType" bson:"resourceType"`
	HeaderData   []HeaderData       `json:"headerData" bson:"headerData"`
	//FieldExtractors []FieldExtractor   `json:"fieldExtractor" bson:"fieldExtractor"` // How to find each data element
	EntryResults []EntryResult `json:"EntryResults" bson:"entryResults"` // one for each entry
}

type ResourceGrid struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	SystemId       string             `json:"SystemId" bson:"SystemId"`
	ResourceType   string             `json:"resourceType" bson:"resourceType"`
	HeaderData     []HeaderData       `json:"headerData" bson:"headerData"`
	FhirExtractors []FhirExtractor    `json:"fhirExtractors" bson:"fhirExtractors"` // How to find each data element
	EntryResults   []EntryResult      `json:"EntryResults" bson:"entryResults"`     // one for each entry
}

type ResourceHeader struct {
	CacheId       primitive.ObjectID `bson:"cacheId" json:"cacheId"`
	DisplayFields []KVData           `bson:"displayFields" json:"displayFields"`
	Attachment    *Attachment        `bson:"attachment" json:"attachment"`
	Narrative     *Narrative         `bson:"narrative" json:"narrative"`
}

type ResourceResponse struct {
	Status              int                      `json:"status"`
	Message             string                   `json:"message"`
	ResourceType        string                   `json:"resourceType"`
	QueryId             string                   `json:"queryId,omitempty"`
	OpOutcome           *fhir.OperationOutcome   `json:"opOutcome,omitempty"`
	OriginalQuery       string                   `json:"originalQuery,omitempty"`
	PageNumber          int                      `json:"pageNumber,omitempty"`
	TotalPages          int64                    `json:"totalPages,omitempty"`
	CountInPage         int                      `json:"countInPage,omitempty"`
	BundleId            string                   `json:"bundleId,omitempty"`
	Header              *CacheHeader             `json:"header,omitempty"`
	Bundle              *fhir.Bundle             `json:"bundle,omitempty"`
	CoreUrl             string                   `json:"coreUrl,omitempty"`
	ResourceId          string                   `json:"resourceId,omitempty"` //Only returned if it is a single Resource not a bundle
	PatientId           string                   `json:"patientId,omitempty"`
	RawResource         json.RawMessage          `json:"rawResource,omitempty"`
	Resource            *ResourceData            `json:"resource,omitempty"`
	Resources           []ResourceData           `json:"resources,omitempty"` // an array of same type of resources.
	GridHeaderData      []HeaderData             `json:"gridHeaderData,omitempty" bson:"gridHeaderData"`
	SummaryType         string                   `json:"summaryType,omitempty"`
	DisplayType         string                   `json:"displayType,omitempty"`
	ActivePatient       *ActivePatient           `json:"activePatient,omitempty"`
	AllergyIntollerance *fhir.AllergyIntolerance `json:"allergyIntollerance,omitempty"`
	Binary              *fhir.Binary             `json:"binary,omitempty"`
	CarePlan            *fhir.CarePlan           `json:"carePlan,omitempty"`
	Condition           *fhir.Condition          `json:"condition,omitempty"`
	Conditions          []*fhir.Condition        `json:"conditions,omitempty"`
	Coverage            *fhir.Coverage           `json:"coverage,omitempty"`
	Coverages           []fhir.Coverage          `json:"coverages,omitempty"`
	DiagnosticReport    *fhir.DiagnosticReport   `json:"diagnosticReport,omitempty"`
	DiagnosticReports   []fhir.DiagnosticReport  `json:"diagnosticReports,omitempty"`
	DocumentReference   *fhir.DocumentReference  `json:"documentReference,omitempty"`
	DocumentReferences  []fhir.DocumentReference `json:"documentReferences,omitempty"`
	Encounter           *fhir.Encounter          `json:"encounter,omitempty"`
	Encounters          []*fhir.Encounter        `json:"encounters,omitempty"`
	Goal                *fhir.Goal               `json:"goal,omitempty"`
	Goals               []fhir.Goal              `json:"goals,omitempty"`
	Immunization        *fhir.Immunization       `json:"immunization,omitempty"`
	Immunizations       []*fhir.Immunization     `json:"immunizations,omitempty"`
	Observation         *fhir.Observation        `json:"observation,omitempty"`
	Observations        []*fhir.Observation      `json:"observations,omitempty"`
	Patient             *fhir.Patient            `json:"patient,omitempty"`
	Patients            []fhir.Patient           `json:"patients,omitempty"`
	Procedures          []fhir.Procedure         `json:"procedures,omitempty"`
	Procedure           *fhir.Procedure          `json:"procedure,omitempty"`

	PatientCache      *PatientCache       `json:"vsPatient,omitempty"`
	PatientCaches     []*PatientCache     `json:"vsPatients,omitempty"`
	DiagReptCache     *DiagReptCache      `json:"vsDiagRept,omitempty"`
	DiagReptCaches    []*DiagReptCache    `json:"vsDiagRepts,omitempty"`
	DocRefCaches      []*DocRefCache      `json:"vsDocRefs,omitempty"`
	DocRefCache       *DocRefCache        `json:"vsDocRef,omitempty"`
	VsProcedures      []*ProcedureCache   `json:"vsProcedures,omitempty"`
	VsProcedure       *ProcedureCache     `json:"vsProcedure,omitempty"`
	ObservationCaches []*ObservationCache `json:"vsObservations,omitempty"`
	ObservationCache  *ObservationCache   `json:"vsObservation,omitempty"`

	MedicationRequest        *fhir.MedicationRequest        `json:"medicationRequest,omitempty"`
	MedicationStatement      *fhir.MedicationStatement      `json:"medicationStatement,omitempty"`
	MedicationDispense       *fhir.MedicationDispense       `json:"medicationDispense,omitempty"`
	MedicationAdministration *fhir.MedicationAdministration `json:"medicationAdministration,omitempty"`
	OperationOutcome         *fhir.OperationOutcome         `json:"operationOutcome,omitempty"`
	Medication               *fhir.Medication               `json:"medication,omitempty"`
	Organization             *fhir.Organization             `json:"organization,omitempty"`
	Practitioner             *fhir.Practitioner             `json:"practitioner,omitempty"`
	PractitionerRole         *fhir.PractitionerRole         `json:"practitionerRole,omitempty"`

	RelatedPerson   *fhir.RelatedPerson   `json:"relatedPerson,omitempty"`
	ResearchStudy   *fhir.ResearchStudy   `json:"researchStudy,omitempty"`
	ResearchSubject *fhir.ResearchSubject `json:"researchSubject,omitempty"`

	Questionnaire         *fhir.Questionnaire         `json:"questionnaire,omitempty"`
	QuestionnaireResponse *fhir.QuestionnaireResponse `json:"questionnaireResponse,omitempty"`
	Device                *fhir.Device                `json:"device,omitempty" bson:"device,omitempty"`
	DeviceRequest         *fhir.DeviceRequest         `json:"deviceRequest,omitempty"`
	DeviceUseStatement    *fhir.DeviceUseStatement    `json:"deviceUseStatement,omitempty"`
	DeviceMetric          *fhir.DeviceMetric          `json:"deviceMetric,omitempty"`
	DeviceDefinition      *fhir.DeviceDefinition      `json:"deviceDefinition,omitempty"`

	//IResource 	 *interface			`json:"iResource,omitempty"`
}

type ResourceSummary struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	SystemId       string             `json:"SystemId" bson:"SystemId"`
	ResourceType   string             `json:"resourceType" bson:"resourceType"`
	HeaderData     []HeaderData       `json:"headerData" bson:"headerData"`
	FhirExtractors []FhirExtractor    `json:"fhirExtractor" bson:"fhirExtractor"` // How to find each data element
	EntryResults   []ElementData      `json:"entryResults" bson:"entryResults"`   // one for each entry
}

// type ResourceGrid struct {
// 	ID           primitive.ObjectID `json:"id" bson:"_id"`
// 	SystemId     string             `json:"SystemId" bson:"SystemId"`
// 	ResourceType string             `json:"resourceType" bson:"resourceType"`
// 	HeaderData   []HeaderData       `json:"headerData" bson:"headerData"`
// 	//FieldExtractors []FieldExtractor   `json:"fieldExtractor" bson:"fieldExtractor"` // How to find each data element
// 	EntryResults []EntryResult `json:"EntryResults" bson:"entryResults"` // one for each entry
// }

type RunTimeConfig struct {
	CoreDB        string
	CoreDataBase  string
	ServiceName   string
	ConfigVersion string
	Company       string
	RefreshSecret string
	AccessSecret  string
	TokenDuration int
	ListenPort    string
	CfgString     string
	Port          string
}

type SavePayload struct {
	ResourceType  string          `json:"resourceType"  bson:"resourceType"`  // Type of ResourceId is (Patient,...)
	DestPatientId string          `json:"destPatientId" bson:"destPatientId"` // PatientId in the Destination(practice) EMR to link the resource to
	DestSystemId  string          `json:"destSystemId" bson:"destSystemId"`   // Destination (local) EMR FhirSystemId
	SrcSystemId   string          `json:"srcSystemId" bson:"srcSystemId"`     // Source (remote) EMR FhirSystemId
	SrcResourceId string          `json:"srcResourceId" bson:"srcResourceId"` // Source (remote) EMR ResourceId
	OtherId       string          `json:"otherId" bson:"otherId"`
	SrcResource   json.RawMessage `json:"srcResource,omitempty" bson:"srcResource,omitempty"` // the fhir information of the of Resource to Save
}

type SaveResponse struct {
	ResourceType string `json:"resourceType"  bson:"resourceType"` // Type of ResourceId is (Patient,...)
	Id           string `json:"Id" bson:"id"`                      //ID of the patient the resource will be attached to. on non patient saves.

	Resource json.RawMessage `json:"Resource" bson:"resource"`
	//Text         string `json:"Text" bson:"text"`
	Mrn string `json:"Mrn" bson:"mrn"` // For patient save only the MRN that should be used for the patient in the local EMR
}

type SearchFilter struct {
	ResourceType string `json:"resourceType" bson:"resourceType"`
}
type SearchSortFields struct {
	Id           primitive.ObjectID `bson:"_id" 		 json:"id"`
	ResoureType  string             `bson:"resourceType" json:"resourceType"`
	SearchFields []string           `bson:"searchFields" json:"searchFields"`
	SortFields   []string           `bson:"sortFields"   json:"sortfFields"`
}
type ServiceConfig struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Customer         Customer           `json:"customer"`
	Name             string             `json:"name"`
	Version          string             `json:"version"`
	DataConnector    *DataConnector     `json:"data_connector"`
	DataConnectors   []*DataConnector   `json:"data_connectors"`
	Services         []*ServiceScope    `json:"services" bson:"services"` //Services used by this Connector
	MyEndPoints      []*EndPoint        `json:"my_endpoints"`
	ServiceEndPoints []*EndPoint        `json:"service_endpoints" bson:"service_endpoints"`
	ConnectInfo      []*KVData          `json:"connect_info" bson:"connect_info"`
	Data             []*KVData          `json:"data" bson:"data"`
	Connectors       []ConnectorConfig  `json:"connectors" bson:"connectors"`
	CallBacks        []*KVData          `json:"call_backs" bson:"call_backs"`
	OriginsAllowed   []string           `json:"originsAllowed" bson:"originsAllowed"`
	BaseURL          string             `json:"baseUrl" bson:"baseUrl"` // First part is the protocol http://localhost:port/api/v1
	MRNSystem        string             `json:"mrnSystem" bson:"mrnSystem"`
}

type ServiceScope struct {
	Name  string `json:"name" bson:"name"`
	Scope string `json:"scope" bson:"scope"` // min, norm, max
}

//Elements = individual fields in a FHIR record.
//Entry = individual records of the returned resource

// type ResourceElement struct {
// 	Results []ResourceElementResult `json:"results" bson:"results"` // one for each data element
// }

//	type Practice struct {
//		ID             primitive.ObjectID `json:"id" bson:"_id"`
//		Name           string             `json:"name" bson:"name"`
//		DisplayName    string             `json:"displayName" bson:"displayName"`
//		Description    string             `json:"description" bson:"description"`
//		Lat            string             `json:"lat,omitempty" bson:"lat"`
//		Lon            string             `json:"lon,omitempty" bson:"lon"`
//		IconLogo       string             `json:"iconLog,omitempty" bson:"iconLogo"`
//		Systems        []*SystemSummary   `json:"systems" bson:"systems"`
//		Classification string             `json:"classification" bson:"classification"` //Practice or hospital
//		Members        []*UserSummary     `json:"members,omitempty" bson:"members"`
//	}

// SessionConnection is a remote EMR The User may connect to
// Would need to include the Facility/System Name
type SessionConnection struct {
	UserId     primitive.ObjectID `json:"userId" bson:"userId"`
	FacilityId primitive.ObjectID `json:"facilityId" bson:"facilityId"`
	SystemId   primitive.ObjectID `json:"SystemId" bson:"SystemId"`
	//Last X patients the user has selected
	PatientHistory []PatientSummary `json:"patientHistory" bson:"patientHistory"`
	Token          string           `json:"token" bson:"token"` //for this connection to Remote

}

type Status struct {
	State      string    `json:"state" bson:"state"` // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
	StatusTime time.Time `json:"status_time" bson:"status_time"`
	Comment    string    `json:"comment"`
}

type StatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type StatusReport struct {
	StatusType string    `json:"status_type" bson:"status_type"` //update, critical
	Status     string    `json:"status" bson:"status"`           // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	Nanotime   int64     `json:"nanotime" bson:"nanotime"`
	Comment    string    `json:"comment" bson:"comment"`
}

type System struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FacilityId      primitive.ObjectID `json:"facilityId" bson:"facilityId"`
	FacilityName    string             `json:"facilityName" bson:"facilityName"`
	FhirVersion     string             `json:"fhirVersion,omitempty" bson:"fhirVersion"`
	Name            string             `json:"name" bson:"name"`
	DisplayName     string             `json:"displayName" bson:"displayName"`
	Abilities       []string           `json:"abilities,omitempty" bson:"abilities"`
	SystemType      string             `json:"systemType,omitempty" bson:"systemType"`
	ConnectorConfig *ConnectorConfig   `json:"connectorConfig" bson:"connectorConfig"`
	CoreUrl         string             `json:"coreUrl" bson:"coreUrl"`
	BaseCoreUrl     string             `json:"baseCoreUrl" bson:"baseCoreUrl"`
	MRNSystem       string             `json:"mrnSystem" bson:"mrnSystem"`
	ResourceConfigs []ResourceConfig   `json:"resourceConfigs" bson:"resourceConfigs"` // Resources on this system
}

type SystemConfig struct {
	ID              primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Name            string                 `json:"name" bson:"name"`
	DisplayName     string                 `json:"displayName" bson:"displayName"`
	SystemType      string                 `json:"systemType" bson:"systemType"`
	Url             string                 `json:"url" bson:"url"`
	ConnectorId     primitive.ObjectID     `json:"connectorId" bson:"connectorId"`
	Identifiers     []*KVData              `json:"identifiers" bson:"identifiers"`
	Data            []*KVData              `json:"data,omitempty" bson:"data,omitempty"`
	ResourceConfigs []*fhir.ResourceConfig `json:"resourceConfigs" bson:"resourceConfigs"` // Resources on this system
	FacilityId      primitive.ObjectID     `json:"facilityId" bson:"facilityId"`
	CoreUrl         string                 `json:"coreUrl" bson:"coreUrl"`
	BaseCoreUrl     string                 `json:"baseCoreUrl" bson:"baseCoreUrl"`
	MRNSystem       string                 `json:"mrnSystem" bson:"mrnSystem"`
	Abilities       []string               `json:"abilities,omitempty" bson:"abilities"`
	ConnectorConfig *ConnectorConfig       `json:"connectorConfig" bson:"connectorConfig"`
	// TimeToLive      int                `json:"timeToLive" bson:"timeToLive"`
	// DbName          string             `json:"dbName" bson:"dbname"` // for cache and actual data  for CA3
}

type SystemSummary struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FacilityId      primitive.ObjectID `json:"facilityId" bson:"facilityId"`
	Name            string             `json:"name" bson:"name"`
	DisplayName     string             `json:"displayName" bson:"displayName"`
	CoreUrl         string             `json:"coreUrl" bson:"coreUrl"`
	Abilities       []string           `json:"abilities,omitempty" bson:"abilities"`
	MRNSystem       string             `json:"mrnSystem" bson:"mrnSystem"`
	SystemType      string             `json:"systemType" bson:"systemType"`
	ResourceConfigs []*ResourceConfig  `json:"resourceConfigs" bson:"resourceConfigs"` // Resources on this system
}

//	type ServiceConfig struct {
//		ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
//		Name    string             `json:"name"`
//		Version string             `json:"version"`
//		//Messaging  		Messaging					// move to endpoints
//		DataConnector    *DataConnector  `json:"dataconnector"`
//		Services         []*ServiceScope `json:"services" bson:"services"`
//		MyEndPoints      []*EndPoint     `json:"myendpoints"`
//		ServiceEndPoints []*EndPoint     `json:"serviceendpoints" bson:"service_endpoints"`
//		ConnectInfo      []*KVData       `json:"connect_info" bson:"connect_info"`
//	}

// type cacheStatus struct {
// }

//SrcResourceId string          `json:"srcResourceId" bson:"srcResourceId"`                 // The ID of the source Resource to save.
//SrcSystemId   string          `json:"srcSystemId" bson:"srcSystemId"`                     // Source(Hospital) EMR FhirSystemId}

//ResourceType string `json:"resourceType"  bson:"resourceType"` // Type of ResourceId is (Patient,...)
//DestPatientId string `json:"destPatientId" bson:"destPatientId"` // PatientId in the Destination(practice) Blank for save patient
//DestSystemId string `json:"destSystemId" bson:"destSystemId"` // Destination (practice) EMR FhirSystemId
//SrcSystemId  string `json:"srcSystemId" bson:"srcSystemId"`   // Source(Hospital) EMR FhirSystemId
//SrcResourceId string `json:"srcResourceId" bson:"srcResourceId"` // The ID of the source Resource to save.
//SrcResource json.RawMessage `json:"srcResource,omitempty" bson:"srcResource,omitempty"` // the fhir information of the of Resource to Save

// type SavePatientPayload struct {
// 	SrcPatientId string          `json:"srcPatientId" bson:"srcPatientId"`
// 	SrcPatient   json.RawMessage `json:"srcPatient,omitempty" bson:"srcPatient,omitempty"` // the fhir information of the of Resource to Save
// 	DestPatientId  string   `json:"destPatientId" bson:"destPatientId"` // PatientId in the Destination(practice) EMR to link the resource to
// 	DestSystemId string          `json:"destSystemId" bson:"destSystemId"`                 // Destination (practice) EMR SystemId
// 	SrcSystemId  string          `json:"srcSystemId" bson:"srcSystemId"`                   // Source(Hospital) EMR SystemId
// }

// type SavePatientResponse struct {
// 	Patient fhir.Patient `json:"Patient"`
// 	Id      string       `json:"Id" bson:"id"`
// }

// type SaveResourcePayload struct {
// 	ResourceIds []string `json:"resourceIds" bson:"resourceIds"`
// }

// type Status struct {
// 	State      string    `json:"state" bson:"state"` // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
// 	StatusTime time.Time `json:"status_time" bson:"status_time"`
// 	Comment    string    `json:"comment"`
// }

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// type User struct {
// 	ID                   primitive.ObjectID `json:"id" bson:"_id"` //UserId Must be Email Address
// 	UserName             string             `json:"user_name" bson:"user_name"`
// 	Password             string             `json:"password" bson:"password"`
// 	FullName             string             `json:"full_name" bson:"full_name"`
// 	Phone                string             `json:"phone" bson:"phone"`
// 	Role                 string             `json:"role" bson:"role"` //Single String role: Provider, Admin, Nurse, Office
// 	LastLogin            *time.Time         `json:"last_login" bson:"last_login"`
// 	LastAttempt          *time.Time         `json:"last_attempt" bson:"last_attempt"`
// 	Attempts             int                `json:"attempts" bson:"attempts"`
// 	Remotes              []*Facility        `json:"remotes" bson:"remotes"` //Medical Remotes user has access to, can include the emr for the Practice
// 	Locals               []*Facility        `json:"locals"  bson:"locals"`  // the users EMR The have Write auth to
// 	PasswordChangedAt    *time.Time         `json:"passwordChangedAt" bson:"password_changed_at"`
// 	CreatedAt            *time.Time         `json:"created_at" bson:"created_at"`
// 	UpdatedAt            *time.Time         `json:"updated_at" bson:"updated_at"`
// 	DeletedAt            *time.Time         `json:"deleted_at" bson:"deleted_at"`
// 	CurrentLocalPatient  *ActivePatient     `json:"currentLocalPatient" bson:"currentLocalPatient"` //Current local patient
// 	CurrentRemotePatient *ActivePatient     `json:"currentRemotePatient" bson:"currentRemotePatient"`
// 	AuthSession          *AuthSession       `json:"authSession" bson:"authSession"` // set to nil on logout.
// 	BaseUrl              string             `json:"baseUrl" bson:"baseUrl"`
// }

// type UserSummary struct {
// 	ID       primitive.ObjectID `json:"id" bson:"_id"`
// 	FullName string             `json:"fullName" bson:"fullName"`
// }

// type VsFhirPatient struct {
// 	OriginalSystem string       `json:"originalSystem" bson:"originalSystem"`
// 	OriginalId     string       `json:"originalId" bson:"originalId"`
// 	NewSystem      string       `json:"newSystem" bson:"newSystem"`
// 	NewSystemId    string       `json:"newSystemId" bson:"newSystemId"`
// 	Patient        fhir.Patient `json:"patient" bson:"patient"`
// }

type VsFhirDocumentReference struct {
	OriginalSystem string                 `json:"originalSystem" bson:"originalSystem"`
	OriginalId     string                 `json:"originalId" bson:"originalId"`
	NewSystem      string                 `json:"newSystem" bson:"newSystem"`
	NewSystemId    string                 `json:"newSystemId" bson:"newSystemId"`
	PDFUrl         string                 `json:"pdfUrl" bson:"pdfUrl"`
	baseReserence  fhir.DocumentReference `json:"documentReference" bson:"documentReference"`
}
