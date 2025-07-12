package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	// Route{
	// 	"Signup",
	// 	"POST", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/Signup",
	// 	postSignup,
	// },
	// Route{
	// 	"UpdateAccount",
	// 	"POST", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/Extractor",
	// 	updateAccount,
	// },
	// Route{
	// 	"ChangePassword",
	// 	"PUT", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/Extractor",
	// 	changePassword,
	// },
	// Route{
	// 	"GetEmrs",
	// 	"GET", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/GetEmrs",
	// 	getEMRs,
	// },

	// Route{
	// 	"Extractor",
	// 	"POST", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/Extractor",
	// 	pdfExtractor,
	// },

	// Route{
	// 	"AutoDelivery",
	// 	"POST", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/AutoDelivery/{resourceType}/{resourceId}",
	// 	resourceAutoDelivery,
	// },
	// Route{
	// 	"AutoDelivery",
	// 	"POST", // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/AutoDelivery/{resourceType}", // Payload contains the PDF/PS to be delivered
	// 	pdfAutoDelivery,
	// },
	// Route{
	// 	"SaveAutoDeliveryConfig",
	// 	"POST",                                  // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/AutoDeliveryConfig", // Payload contains the AutoDeliveryConfig to be saved
	// 	saveAutoDeliveryConfig,
	// },
	// Route{
	// 	"UpdateAutoDeliveryConfig",
	// 	"PUT",                                   // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/AutoDeliveryConfig", // Payload contains the AutoDeliveryConfig to be updated
	// 	updateAutoDeliveryConfig,
	// },
	// Route{
	// 	"GetAutoDeliveryConfigByResourceType",
	// 	"GET",
	// 	"/system/{SystemId}/AutoDeliveryConfig/{resourceType}", // REturns the requested Auto Delivery Config
	// 	findAutoDeliveryConfig,
	// },
	// Route{
	// 	"QueryAutoDeliveryConfig",
	// 	"Get",                                   // SystemId = Destination SystemId as all GETs/Puts are
	// 	"/system/{SystemId}/AutoDeliveryConfig", // Payload contains the PDF/PS to be delivered
	// 	findAutoDeliveryConfig,
	// },
	// Route{
	// 	"GetAutoDeliveryConfig",
	// 	"GET", //
	// 	"/system/{SystemId}/AutoDeliveryConfig/{resourceType}", // Payload contains the PDF/PS to be delivered
	// 	getAutoDeliveryConfig,
	// },
	// Route{
	// 	"SetupRelease",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/api/rest/v1/SetupRelease",
	// 	"/system/{SystemId}/SetUpRelease",
	// 	startRelease,
	// },
	// Route{
	// 	"StartRelease",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/api/rest/v1/SetupRelease",
	// 	"/system/{SystemId}/StartRelease",
	// 	startRelease,
	// },
	// Route{
	// 	"QueueRelease",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/QueueRelease/{ReleaseId}",
	// 	queueRelease,
	// },

	// Route{
	// 	"GetRelease",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/Release/{releaseId}",
	// 	getRelease,
	// },

	// Route{
	// 	"FindRelease",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/Release",
	// 	findReleases,
	// },
	// Route{
	// 	"UpdateRelease",
	// 	"PUT", // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/Release",
	// 	updateRelease,
	// },
	// Route{
	// 	"Redirect",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	"/api/rest/v1/Redirect",
	// 	redirectHandler,
	// },
	///////////////////////////////////////////// Health Check Routes //////////////////////////////////////////
	Route{
		"HealthCheck",
		"GET",
		"/api/rest/v1/healthcheck",
		HealthCheck,
	},
	Route{
		"Health",
		"GET",
		"/api/rest/v1/health",
		HealthCheckHandler,
	},
	Route{
		"Health",
		"POST",
		"/api/rest/v1/health",
		HealthCheckHandler,
	},
	// Add Check Connectors
	// Route{
	// 	"ConnectorHealth",
	// 	"GET",
	// 	"system/{system_id}/check_connector",
	// 	checkConnectors,
	// },

	//////////////////////////////////////////// Authorization Routes /////////////////////////////
	// Route{
	// 	"ValidateLicense",
	// 	"GET",
	// 	"/client/{clientId}/{product}/ValidateLicense",
	// 	validateLicense,
	// },
	// Route{
	// 	"LoginConfig",
	// 	"POST",
	// 	//"/api/rest/v1/auth/authorize",
	// 	"/api/rest/v1/login",
	// 	PostLogin,
	// },
	// Route{
	// 	"LoginPost",
	// 	"POST",
	// 	//"/api/rest/v1/auth/authorize",
	// 	"/api/rest/v1/authorize",
	// 	PostLogin,
	// },
	// Route{
	// 	"AuthAuthorizePost",
	// 	"POST",
	// 	"/api/rest/v1/auth/authorize",
	// 	//"/api/rest/v1/authorize",
	// 	PostLogin,
	// },
	// Route{
	// 	"FhirCallback",
	// 	"POST",
	// 	//"/api/rest/v1/auth/authorize",
	// 	//"/api/rest/v1/authorize",
	// 	"/api/rest/v1/callB",
	// 	smartCallBack,
	// },
	// Route{
	// 	"RestLogin",
	// 	"GET",
	// 	"/api/rest/v1/login",
	// 	login,
	// },
	// Route{
	// 	"Logout",
	// 	"DELETE",
	// 	"/api/rest/v1/authorize",
	// 	login,
	// },
	// Route{
	// 	"remoteLogin/",
	// 	"POST",                           // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/RemoteLogin", // Body username and password
	// 	remoteLogin,
	// },
	// Route{
	// 	"remoteFhirLogin",
	// 	"POST",                               // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/RemoteFhirLogin", // Body username and password
	// 	remoteFhirLogin,
	// },
	// // Route{
	// // 	"validateSession/",
	// // 	"GET",                                // SystemId = Destination SystemId as all GETs are
	// // 	"/system/{SystemId}/ValidateSession", // the UC token is in authorization of header
	// // 	validateSession,
	// // },
	// Route{
	// 	"GetConfig",
	// 	"GET",
	// 	"/api/rest/v1/config",
	// 	getConfig,
	// },
	// Route{
	// 	"GetFhirSystem",
	// 	"GET",
	// 	"/api/rest/v1/fhirSystem/{fhirSystemId}",
	// 	getFhirSystem,
	// },
	// Route{
	// 	"GetImmunization",
	// 	"GET",
	// 	"/system/{SystemId}/Immunization",
	// 	getImmunization,
	// },
	// Route{
	// 	"FINDImmunization",
	// 	"GET",
	// 	"/system/{SystemId}/Immunization",
	// 	findImmunizations,
	// },
	// Route{
	// 	"GetConnectorConfig",
	// 	"GET",
	// 	"/api/rest/v1/connector",
	// 	getConnectorConfig,
	// 	//getConnectorConfig,
	// },
	// Route{
	// 	"GetSystemSummary",
	// 	"GET",
	// 	"/system/{SystemId}/SystemSummary",
	// 	getSystemSummary,
	// 	//getConnectorConfig,
	// },
	// Route{
	// 	"FindRecipients",
	// 	"GET",
	// 	"/system/{SystemId}/Recipients",
	// 	findRecipients,
	// },
	// Route{
	// 	"GetRecipient",
	// 	"GET",
	// 	"/system/{SystemId}/Recipient/{RecipientId}",
	// 	getRecipient,
	// },
	// Route{
	// 	"SaveRecipient",
	// 	"POST",
	// 	"/system/{SystemId}/Recipient",
	// 	saveRecipient,
	// },
	///////////////////////////////////////////// Resource Routes //////////////////////////////////////////

	////////////////////Patient Routes/////////////////////
	// Route{
	// 	"FindPatients",
	// 	"GET",
	// 	"/system/{SystemId}/Patient",
	// 	findPatients,
	// },
	// Route{
	// 	"GetPatient",
	// 	"GET",
	// 	"/system/{SystemId}/Patient/{resourceId}",
	// 	getPatient,
	// },
	// // Route{
	// // 	"QueuePatient",
	// // 	"POST",                       // /Patient/
	// // 	"/system/{SystemId}/Patient", // Body source system id, Patient id, remote SystemId, MRN
	// // 	savePatient,
	// // },
	// Route{
	// 	"SavePatient",
	// 	"POST",                       // /Patient/
	// 	"/system/{SystemId}/Patient", // Body source system id, Patient id, remote SystemId, MRN
	// 	savePatient,
	// },
	// Route{
	// 	"SaveTemplate",
	// 	"POST",
	// 	"/system/{SystemId}/Template",
	// 	saveTemplate,
	// },
	// Route{
	// 	"UpdateTemplate",
	// 	"PUT",
	// 	"/system/{SystemId}/Template",
	// 	updateTemplate,
	// },
	// Route{
	// 	"GetTemplateByName",
	// 	"GET",
	// 	"/system/{SystemId}/Template/{templateName}",
	// 	getTemplateByName,
	// },
	// Route{
	// 	"GetTemplateById",
	// 	"GET",
	// 	"/system/{SystemId}/TemplateById/{templateId}",
	// 	getTemplateById,
	// },

	// Route{
	// 	"GetPatientCacheStatus",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/PatientCacheStatus",
	// 	getPatientCacheStatus,
	// },
	// Route{
	// 	"GetPatientCachePage",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/PatientCachePage", // /{pageNum}", //"",
	// 	getPatientCachePage,
	// },
	// Route{
	// 	"GetPatientCachePageSize",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/PatientCachePage", ///{pageNum}/PageSize/{pageSize}", //"",
	// 	getPatientCachePage,
	// },
	///////////////////////////////////////////// Cache Management Routes //////////////////////////////////////////
	// Route{
	// 	"GetCachedResourcePage",
	// 	"GET",
	// 	//"/system/{SystemId}/query/{queryId}/CachePage/{page}/PageSize/{pageSize}",
	// 	"/system/{systemId}/query/{queryId}/{resource}/CachedResourcePage",
	// 	getCachedResourcePageForQueryId,
	// },
	// Route{
	// 	"GetResourceCachePage",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/ResourceCachePage/{resource}",
	// 	getResourceCachePage,
	// },

	// Route{
	// 	"FinishCacheResources",
	// 	"POST",
	// 	"/system/{SystemId}/FinishCache/{queryId}/{resourceType}",
	// 	finishCacheResources,
	// },
	// Route{
	// 	"GetCachedDocRefPage",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/DocRefCachePage",
	// 	getCachedDocRefPageForQueryId,
	// },
	// Route{
	// 	"GetCacheStatus",
	// 	"GET",
	// 	"/system/{systemID}/query/{queryId}/CacheStatus",
	// 	getCacheStatus,
	// },

	// Route{
	// 	"GetDocRefCacheStatus",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/DocRefCacheStatus",
	// 	getDocRefCacheStatus,
	// },
	// Route{
	// 	"GetDocRefCacheStatus",
	// 	"GET",
	// 	"/system/{systemId}/query/{queryId}/DocumentReferenceCacheStatus",
	// 	getDocRefCacheStatus,
	// },
	// Route{
	// 	"DocRefCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/DocRefCachePage",
	// 	//getCachePage,
	// 	getDocRefCachePageForQueryId,
	// },
	// Route{
	// 	"DocRefCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/DocumentReferenceCachePage",
	// 	//getCachePage,
	// 	getDocRefCachePageForQueryId,
	// },
	//TODO:  getObservationCachePage
	// Route{
	// 	"ObservationCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ObservationCachePage",
	// 	getObservationCachePageForQueryId,
	// },
	// Route{
	// 	"ObservationCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ObservationCacheStatus",
	// 	getObservationCacheStatus,
	// },
	// TODO:  getProcedureCachePage
	// Route{
	// 	"ProcedureCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ProcedureCachePage",
	// 	getProcedureCachePageForQueryId,
	// },
	//TODO:  getProcedureCacheStatus
	// Route{
	// 	"ProcedureCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ProcedureCacheStatus",
	// 	getProcedureCacheStatus,
	// },
	// Route{
	// 	"DiagReptCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/DiagReptCachePage",
	// 	getDiagReptCachePageForQueryId,
	// },
	// Route{
	// 	"DiagReptCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/DiagReptCacheStatus",
	// 	getDiagReptCacheStatus,
	// },
	// Route{
	// 	"CacheBundle",
	// 	"GET",
	// 	"/system/{SystemId}/{queryId}/Bundle",
	// 	getBundlePage,
	// },
	// Route{
	// 	"GetCacheBundle",
	// 	"GET",
	// 	"/system/{SystemId}/Bundle/{queryId}",
	// 	getBundlePage,
	// },
	// Route{
	// 	"CacheStatus",
	// 	"GET",
	// 	"/api/rest/v1/CacheStatus/{query_id}",
	// 	getCacheStatus,
	// },
	// Route{
	// 	"CachePageForQuery",
	// 	"GET",
	// 	"/system/{SystemId}/CachePage",
	// 	getCachePageForQueryId,
	// },
	// Route{
	// 	"PostCacheResources",
	// 	"POST",
	// 	"/{fhirSystem}/Cache/{queryId}/{pageNum}",
	// 	CacheResources,
	// },
	// Cache a Resource
	// Route{
	// 	"CacheResources",
	// 	"POST",
	// 	"/system/{SystemId}/Cache",
	// 	cacheResource,
	// },
	// TODO: setPatientCacheTTL
	// Route{
	// 	"CacheTTL",
	// 	"POST",
	// 	"/system/{SystemId}/CacheTTL",
	// 	setPatientCacheTTL,
	// },
	// Route{
	// 	"FinishCacheResources",
	// 	"POST",
	// 	"/system/{SystemId}/FinishCache",
	// 	finishCacheResources,
	// },
	// Route{
	// 	"CacheResources",
	// 	"GET",
	// 	"/facility/{facilityId}/system/{SystemId}/Cache",
	// 	CacheResources,
	// },
	// Route{
	// 	"GetPatientCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/PatientCacheStatus",
	// 	getPatientCacheStatus,
	// },
	// Route{
	// 	"GetStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ResourceCacheStatus/{ResourceType}",
	// 	getResourceCacheStatus,
	// },

	// Route{
	// 	"GetStatus",
	// 	"GET",
	// 	"/system/{SystemId}/ResourceCacheStatus/{ResourceType}",
	// 	getDRCacheStatus,
	// },
	// Route{
	// 	"GetCachePage",
	// 	"GET",
	// 	"/api/rest/v1/Cache/{queryId}/{pageNum}",
	// 	getCachePage,
	// },
	// Route{
	// 	"GetCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/CachePage",
	// 	//getCachePage,
	// 	getCachePageForQueryId,
	// },

	// Route{
	// 	"GetCachedResourcePage",
	// 	"GET",
	// 	"/system/{SystemId}/{queryId}/CachePage",
	// 	//getCachePage,
	// 	getCachePageForQueryId,
	// },

	///////////////////////////////////////////// LinkResource Routes //////////////////////////////////////////
	// Route{
	// 	"LinkResource/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/Link/{resource}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	linkResource,
	// },

	///////////////////////////////////////////// ROI Routes //////////////////////////////////////////
	Route{
		// "ROIAuthImage/",
		// "POST", // SystemId = Destination SystemId as all GETs are
		// //"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
		// //"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
		// "/system/{destSystemId}/roiAuthImage", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
		// saveRoiImage,
	},
	////////////////////////////////////////// POST Resource Routes /////////////////////////////////////////
	// Route{
	// 	"LogHipaa/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	"/api/rest/v1/LogHipaa",
	// 	logHipaa,
	// },

	// Route{
	// 	"LinkResource/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/Link/{resource}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	linkResource,
	// },

	// Route{
	// 	"SetActivePatient/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/ActivePatient", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	setActivePatient,
	// },
	// Route{
	// 	"PostPatientBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/PatBundleTransaction",
	// 	postPatBundleTransaction,
	// },

	// Route{
	// 	"PostDocRefBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/DocRefBundleTransaction/{queryId}",
	// 	postDocRefBundleTransaction,
	// },
	// Route{
	// 	"PostDiagReptBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/DiagReptBundleTransaction",
	// 	postDiagReptBundleTransaction,
	// },

	//TODO:  PostObservationBundleTransaction
	// Route{
	// 	"PostObservationBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/ObservationBundleTransaction",
	// 	postObservationBundleTransaction,
	// },

	// TODO:
	// Route{
	// 	"PostProcedureBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/ProcedureBundleTransaction",
	// 	postProcedureBundleTransaction,
	// },
	// Route{
	// 	"PostBundleTxn",
	// 	"POST",
	// 	"/system/{SystemId}/BundleTxn",
	// 	postBundleTransaction,
	// },
	// Route{
	// 	"SaveBaseRes/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	"/system/{DestSystemId}/BaseRes", // Body source system id, array of ResourceIds
	// 	saveBaseRes,
	// },
	// Route{
	// 	"SaveDiagRept",
	// 	"POST",                                // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/DiagnosticReport", // Body source system id, array of ResourceIds
	// 	saveDiagnosticRept,
	// },
	// Route{
	// 	"SaveDocRef",
	// 	"POST",                                 // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/DocumentReference", // Body source system id, array of ResourceIds
	// 	saveDocRef,
	// },
	// Route{
	// 	"SaveResource/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/{resource}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	saveResource,
	// },

	// Route{
	// 	"GetSaveQueue/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/SaveQueue", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getSaveQueue,
	// },
	// Route{
	// 	"GetSaveQueue/{ResourceType}",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/SaveQueue/{ResourceType}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getSaveQueue,
	// },
	// Route{
	// 	"GetSaveQueueCount/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/SaveQueueCount", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getSaveQueueCount,
	// },
	// Route{
	// 	"GetSaveQueueSummary/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/SaveQueueSummary", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getSaveQueueSummary,
	// },
	// Route{
	// 	"GetSaveQueueSummary/{ResourceType}",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/SaveQueueSummary/{ResourceType}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getSaveQueueSummary,
	// },
	// Route{
	// 	"GetPendingQueueSummary/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/PendingQueueSummary", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getPendingQueueSummary,
	// },
	// Route{
	// 	"ManagePendingQueue/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/ManagePendingQueue/{entryId}/patientId/{patId}/{command}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	managePendingQueue,
	// },
	// Route{
	// 	"GetPendingQueueSummary/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{SystemId}/PendingQueueSummary", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getPendingQueueSummary,
	// },
	// Route{
	// 	"GetActivePatient/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/ActivePatient", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getActivePatient,
	// },
	// Route{
	// 	"GetResourceSummary/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/ResourceSummary/{resourceType}", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getResourceSummary,
	// },
	// Route{
	// 	"PostBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/BundleTransaction",
	// 	postBundleTransaction,
	// },
	// Route{
	// 	"BundleTransaction",
	// 	"GET",
	// 	"/system/{SystemId}/BundleTransaction",
	// 	cacheResources,
	// },

	// Route{
	// 	"SavediagRept/",
	// 	"POST", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	"/system/{DestSystemId}/DocumentReference", // Body source system id, array of ResourceIds
	// 	savediagRept,
	// },

	////////////////////////////////////////// Find Resource Routes /////////////////////////////////////////
	// Route{
	// 	"FindAllergyIntolerance",
	// 	"GET",
	// 	"/system/{SystemId}/AllergyIntolerance",
	// 	findResources,
	// },
	// Route{
	// 	"FindCarePlan",
	// 	"GET",
	// 	"/system/{SystemId}/CarePlan",
	// 	findCarePlans,
	// },
	// Route{
	// 	"GetCarePlan",
	// 	"GET",
	// 	"/system/{SystemId}/CarePlan/{resourceId}",
	// 	getCarePlan,
	// },
	// Route{
	// 	"CarePlanCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/CarePlanCachePage",
	// 	getCarePlanCachePageForQueryId,
	// },
	// Route{
	// 	"CarePlanCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/CarePlanCacheStatus",
	// 	getCarePlanCacheStatus,
	// },
	// Route{
	// 	"SaveCarePlan",
	// 	"POST",                        // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/CarePlan", // Body source system id, array of ResourceIds
	// 	saveCarePlan,
	// },
	// Route{
	// 	"PostCarePlanBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/CarePlanBundleTransaction",
	// 	postCarePlanBundleTransaction,
	// },
	// Route{
	// 	"FindCareTeam",
	// 	"GET",
	// 	"/system/{SystemId}/CareTeam",
	// 	findResources,
	// },

	// Route{
	// 	"FindConditions",
	// 	"GET",
	// 	"/system/{SystemId}/Condition",
	// 	findConditions,
	// },
	// Route{
	// 	"GetCondition",
	// 	"GET",
	// 	"/system/{SystemId}/Condition/{resourceId}",
	// 	getCondition,
	// },
	// Route{
	// 	"ConditionCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ConditionCachePage",
	// 	getConditionCachePageForQueryId,
	// },
	// Route{
	// 	"ConditionCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ConditionCacheStatus",
	// 	getConditionCacheStatus,
	// },
	// Route{
	// 	"SaveCondition",
	// 	"POST",                         // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/Condition", // Body source system id, array of ResourceIds
	// 	saveCondition,
	// },
	// Route{
	// 	"PostConditionBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/ConditionBundleTransaction",
	// 	postConditionBundleTransaction,
	// },
	// TODO:  findImmunization
	// Route{
	// 	"FindImmunization",
	// 	"GET",
	// 	"/system/{SystemId}/Immunization",
	// 	findImmunizations,
	// },
	//TODO:replace below
	// Route{
	// 	"GetImmunization",
	// 	"GET",
	// 	"/system/{SystemId}/Immunization/{resourceId}",
	// 	getImmunization,
	// },
	// Route{
	// 	"ImmunizationCachePage",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ImmunizationCachePage",
	// 	getImmunizationCachePageForQueryId,
	// },
	// Route{
	// 	"ImmunizationCacheStatus",
	// 	"GET",
	// 	"/system/{SystemId}/query/{queryId}/ImmunizationCacheStatus",
	// 	getImmunizationCacheStatus,
	// },
	// Route{
	// 	"SaveImmunization",
	// 	"POST",                            // SystemId = Destination SystemId as all GETs are
	// 	"/system/{SystemId}/Immunization", // Body source system id, array of ResourceIds
	// 	saveImmunization,
	// },
	// Route{
	// 	"PostImmunizationBundleTransaction",
	// 	"POST",
	// 	"/system/{SystemId}/ImmunizationBundleTransaction",
	// 	postImmunizationBundleTransaction,
	// },
	// Route{
	// 	"GetActivePatient/",
	// 	"GET", // SystemId = Destination SystemId as all GETs are
	// 	//"/system/{SystemId}/srcSystem/{srcSystemId}/Patient/{patientId}",
	// 	//"/system/{srcSystemId}/{resource}/destSystem/{destSystemId}/{srcResourceId}", // Body source system id, array of ResourceIds
	// 	"/system/{destSystemId}/ActivePatient", ///destSystem/{destSystemId}/{srcResourceId}, // Body source system id, array of ResourceIds
	// 	getActivePatient,
	// },

	// Route{
	// 	"FindPatientsConnector",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{connectorId}/Patient",
	// 	findPatients,
	// },
	// Route{

	// 	"GetTagLine",
	// 	"GET",
	// 	"/api/rest/v1/TagLine",
	// 	getTagLine,
	// },
	// Route{
	// 	"GetDocumentReference",
	// 	"GET",
	// 	"/system/{SystemId}/DocumentReference/{documentId}",
	// 	getDocRef,
	// },
	// // Route{
	// // 	"GetDocumentReference",
	// // 	"GET",
	// // 	"/system/{SystemId}/DocumentReference/{documentId}/pdf",
	// // 	getDocRefPDF,
	// // },
	// Route{
	// 	"FindDocumentReferences",
	// 	"GET",
	// 	"/system/{SystemId}/DocumentReference",
	// 	findDocRefs,
	// },
	// Route{
	// 	"FindDocumentReferencesConnector",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{ConnectorId}/DocumentReference",
	// 	findDocRefs,
	// },
	// Route{
	// 	"MigratePatients",
	// 	"POST",
	// 	"/system/{SystemId}/MigratePatients",
	// 	migratePatients,
	// },
	// Route{
	// 	"AddAccessLog",
	// 	"POST",
	// 	"/system/{SystemId}/AddAccessLog",
	// 	addAccessLog,
	// },
	// Route{
	// 	"FindResources",
	// 	"GET",
	// 	"/system/{SystemId}/{Resource}", // Save/Patient?id={id}&remoteId={remoteId}
	// 	findResources,                   //returns a ResourceResponse
	// },

	// Route{
	// 	"GetResource",
	// 	"GET",
	// 	"/system/{SystemId}/{Resource}/{resourceId}", // Save/Patient?id={id}&remoteId={remoteId}
	// 	getResource,
	// },

	// Route{
	// 	"FindResource",
	// 	"GET",
	// 	"/system/{SystemId}/Find/{resource}", // Save/Patient?id={id}&remoteId={remoteId}
	// 	FindResource,
	// },
	// Route{
	// 	"FindBundle",
	// 	"GET",
	// 	"/system/{SystemId}/bundle/{resource}",
	// 	FindResources,
	// },
	// Route{
	// 	"FindAllergyIntolerance",
	// 	"GET",
	// 	"/system/{SystemId}/AllergyIntolerance",
	// 	findResources,
	// },
	// Route{
	// 	"FindCarePlan",
	// 	"GET",
	// 	"/system/{SystemId}/CarePlan",
	// 	findResources,
	// },

	// Route{
	// 	"FindCareTeam",
	// 	"GET",
	// 	"/system/{SystemId}/CareTeam",
	// 	findResources,
	// },

	// Route{
	// 	"FindCoverage",
	// 	"GET",
	// 	"/system/{SystemId}/Coverage",
	// 	findResources,
	// },
	// Route{
	// 	"FindResourceCoverage",
	// 	"GET",
	// 	"/system/{SystemId}/resource/{resource}",
	// 	findResources,
	// },
	// Route{
	// 	"FindDiagnosticReport",
	// 	"GET",
	// 	"/system/{SystemId}/DiagnosticReport",
	// 	findDiagnosticRepts,
	// },
	// Route{
	// 	"FindDiagnosticReportConnector",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{ConnectorId}/DiagnosticReport",
	// 	findDiagnosticRepts,
	// },
	// // TODO: findObservations findObservations
	// Route{
	// 	"FindObservations",
	// 	"GET",
	// 	"/system/{SystemId}/Observation",
	// 	findObservations,
	// },
	// Route{
	// 	"FindObservationsConnector",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{ConnectorId}/Observation",
	// 	findObservations,
	// },

	// Route{
	// 	"FindDocumentReferences",
	// 	"GET",
	// 	"/system/{SystemId}/DocumentReference",
	// 	findResources,
	// },
	// Route{
	// 	"FindEncounters",
	// 	"GET",
	// 	"/system/{SystemId}/Encounter",
	// 	findResources,
	// },
	// Route{
	// 	"FindFamilyMemberHistory",
	// 	"GET",
	// 	"/system/{SystemId}/FamilyMemberHistory",
	// 	findResources,
	// },
	// Route{
	// 	"FindGoal",
	// 	"GET",
	// 	"/system/{SystemId}/Goal",
	// 	findResources,
	// },
	// TODO: findImmunization
	// Route{
	// 	"FindImmunization",
	// 	"GET",
	// 	"/system/{SystemId}/Immunization",
	// 	findImmunizations,
	// },

	// Route{
	// 	"FindMedicationAdminstration",
	// 	"GET",
	// 	"/system/{SystemId}/MedicationAdministration",
	// 	findResources,
	// },
	// Route{
	// 	"FindMedicationOrder",
	// 	"GET",
	// 	"/system/{SystemId}/MedicationOrder",
	// 	findResources,
	// },
	// Route{
	// 	"FindMedicationRequest",
	// 	"GET",
	// 	"/system/{SystemId}/MedicationRequest",
	// 	findResources,
	// },
	// Route{
	// 	"FindNutritionOrder",
	// 	"GET",
	// 	"/system/{SystemId}/NutritionOrder",
	// 	findResources,
	// },
	// TODO: findObservation
	// Route{
	// 	"FindObservation",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{ConnectorId}/Observation",
	// 	findObservations,
	// },
	//TODO: findProcedure
	// Route{
	// 	"FindProcedure",
	// 	"GET",
	// 	"/system/{SystemId}/Procedure",
	// 	findProcedures,
	// },

	// Route{
	// 	"FindDocumentReferences",
	// 	"GET",
	// 	"/system/{SystemId}/DocumentReference",
	// 	findResources,
	// },
	// Route{
	// 	"FindQuestionnaireResponse",
	// 	"GET",
	// 	"/system/{SystemId}/QuestionnaireResponse",
	// 	findResources,
	// },
	// Route{
	// 	"FindServiceRequest",
	// 	"GET",
	// 	"/system/{SystemId}/ServiceRequest",
	// 	findResources,
	// },

	///////////////////////////////////// GET RESOURCE ROUTES //////////////////////////////////////////////////////////////////
	// Route{
	// 	"GetAllergyIntolerance",
	// 	"GET",
	// 	"/system/{SystemId}/AllergyIntolerance/{resourceId}",
	// 	getResource,
	// },
	// Route{
	// 	"GetBinary",
	// 	"GET",
	// 	"/system/{systemId}/Binary/{imageId}",
	// 	getBinary,
	// },
	// Route{
	// 	"GetPDF",
	// 	"GET",
	// 	"/system/{systemId}/PDF/{imageId}",
	// 	getPDF,
	// },
	// Route{
	// 	"GetCondition",
	// 	"GET",
	// 	"/system/{SystemId}/connector/{ConnectorIdCondition}/{resourceId}",
	// 	getResource,
	// },
	// Route{
	// 	"GetDiagnosticReport",
	// 	"GET",
	// 	"/system/{SystemId}/DiagnosticReport/{DiagId}",
	// 	getDiagnosticReport,
	// },
	// Route{
	// 	"GetDocumentReference",
	// 	"GET",
	// 	"/system/{SystemId}/DocumentReference/{resourceId}",
	// 	getResource,
	// },
	// Route{
	// 	"GetEncounter",
	// 	"GET",
	// 	"/system/{SystemId}/Encounter/{resourceId}",
	// 	getResource,
	// },
	// Route{
	// 	"GetObservation",
	// 	"GET",
	// 	"/system/{SystemId}/Observation/{resourceId}",
	// 	getResource,
	// },

	// Route{
	// 	"GetResource",
	// 	"GET",
	// 	"/system/{SystemId}/{Resource}/{resourceId}",
	// 	getResource,
	// },
	// Route{
	// 	"FindResource",
	// 	"GET",
	// 	"/system/{SystemId}/{Resource}",
	// 	findResources,
	// },

	//////////////////////////////////////////////// Admin Routes ////////////////////////////////////
	// Route{
	// 	"Redirect",
	// 	"POST",
	// 	"/redirectapp",
	// 	reDirect,
	// },

	//Route{
	//	"SetLogLevel",
	//	"GET",
	//	"/system/{SystemId}/log_level/{level}",
	//	GetLogLevel,
	//},
	//Route{
	//	"SetLogLevel",
	//	"PUT",
	//	"/system/{SystemId}/log_level/{level}",
	//	SetLogLevel,
	//},

	////////////////////////////////////////////// UnUsed Routes //////////////////////////////////////////////
	// Route{
	// 	"SearchResource",
	// 	"GET",
	// 	"/{fhirSystem}}/api/rest/v1/{resource}",
	// 	searchResource,
	// },
	// Route{
	// 	"GetResource",
	// 	"GET",
	// 	"/system/{SystemId}/api/rest/v1/{resource}/{id}",
	// 	getResource,
	// },

}
