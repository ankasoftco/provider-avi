package aviprovider

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("avi_actiongroupconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ActionGroupConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_albservicesconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ALBServicesConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_albservicesfileupload", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ALBServicesFileUpload"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_albservicesjob", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ALBServicesJob"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_alertconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AlertConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_alertemailconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AlertEmailConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_alertscriptconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AlertScriptConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_alertsyslogconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AlertSyslogConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_analyticsprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AnalyticsProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_applicationpersistenceprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ApplicationPersistenceProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_applicationprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ApplicationProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_authmappingprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AuthMappingProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_authprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AuthProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_autoscalelaunchconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AutoScaleLaunchConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_availabilityzone", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "AvailabilityZone"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_backup", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "Backup"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("avi_backupconfiguration", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BackupConfiguration"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_botconfigconsolidator", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BotConfigConsolidator"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_botdetectionpolicy", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BotDetectionPolicy"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_botipreputationtypemapping", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BotIPReputationTypeMapping"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_botmapping", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BotMapping"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_certificatemanagementprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "CertificateManagementProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_cloud", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "Cloud"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_cloudconnectoruser", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "CloudConnectorUser"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_cloudproperties", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "CloudProperties"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_cluster", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "Cluster"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_clusterclouddetails", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ClusterCloudDetails"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_controllerportalregistration", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ControllerPortalRegistration"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_controllerproperties", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ControllerProperties"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_controllersite", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ControllerSite"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_dnspolicy", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "DnsPolicy"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_dynamicdnsrecord", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "DynamicDnsRecord"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_errorpagebody", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ErrorPageBody"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_errorpageprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "ErrorPageProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_federationcheckpoint", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "FederationCheckpoint"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_fileobject", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "FileObject"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_fileservice", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "FileService"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_geodb", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "GeoDB"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_gslb", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "Gslb"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_gslbgeodbprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "GslbGeoDbProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_gslbservice", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "GslbService"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_hardwaresecuritymodulegroup", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "HardwareSecurityModuleGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_healthmonitor", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "HealthMonitor"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_httppolicyset", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "HTTPPolicySet"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_icapprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "IcapProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_image", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "Image"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_inventoryfaultconfig", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "InventoryFaultConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_ipaddrgroup", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "IpAddrGroup"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("avi_ipamdnsproviderprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "IpamDnsProviderProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_ipreputationdb", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "IPReputationDB"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_jwtprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "JWTProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_jwtserverprofile", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "JWTServerProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_l4policyset", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "BotMaL4PolicySet"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_labelgroup", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "LabelGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_licenseledgerdetails", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "LicenseLedgerDetails"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_licensestatus", func(r *config.Resource) {
		r.ShortGroup = "aviprovider"
		r.Kind = "LicenseStatus"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("avi_memorybalancerrequest", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "MemoryBalancerRequest"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_microservicegroup", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "MicroServiceGroup"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_natpolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "NatPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_network", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Network"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_networkprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "NetworkProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_networksecuritypolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "NetworkSecurityPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_networkservice", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "NetworkService"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_nsxtsegmentruntime", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "NsxtSegmentRuntime"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_objectaccesspolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ObjectAccessPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_pingaccessagent", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PingAccessAgent"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_pkiprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PkiProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_pool", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Pool"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_poolgroup", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PoolGroup"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_poolgroupdeploymentpolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PoolGroupDeploymentPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_portalfileupload", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PortalFileUpload"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_prioritylabels", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "PriorityLabels"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_protocolparser", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ProtocolParser"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_rmcloudopsproto", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "RmCloudOpsProto"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_role", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Role"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_scheduler", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Scheduler"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_securitymanagerdata", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SecurityManagerData"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_securitypolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SecurityPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_seproperties", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SeProperties"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_server", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Server"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_serverautoscalepolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ServerAutoScalePolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_serviceengine", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ServiceEngine"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_serviceenginegroup", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ServiceEngineGroup"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_serviceenginepolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "ServiceEnginePolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_siteversion", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SiteVersion"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_snmptrapprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SnmpTrapProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_sslkeyandcertificate", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SslKeyAndCertificate"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_sslprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SslProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_ssopolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SsoPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_statediffoperation", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "StateDiffOperation"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_statediffsnapshot", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "StateDiffSnapshot"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_stringgroup", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "StringGroup"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_systemconfiguration", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SystemConfiguration"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_systemlimits", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "SystemLimits"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_tenant", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Tenant"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_testsedatastorelevel1", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Testsedatastorelevel1"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_testsedatastorelevel2", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Testsedatastorelevel2"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_testsedatastorelevel3", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Testsedatastorelevel3"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_trafficcloneprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "TrafficCloneProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_upgradestatusinfo", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "UpgradeStatusInfo"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_upgradestatussummary", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "UpgradeStatusSummary"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_user", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "User"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_useraccount", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "UserAccount"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_useraccountprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "UserAccountProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_vcenterserver", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "VcenterServer"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_virtualservice", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "VirtualService"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_vrfcontext", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "VrfContext"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_vsdatascriptset", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "VsDataScriptSet"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_vsgs", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Vsgs"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_vsvip", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Vsvip"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_wafapplicationsignatureprovider", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WafApplicationSignatureProvider"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_wafcrs", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WafCrs"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_wafpolicy", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WafPolicy"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_wafpolicypsmgroup", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WafPolicyPsmGroup"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_wafprofile", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WafProfile"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_webapput", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "WebappUt"
		r.Version = "v1"
	})

	p.AddResourceConfigurator("avi_webhook", func(r *config.Resource) {
		r.ShortGroup = "avi"
		r.Kind = "Webhook"
		r.Version = "v1"
	})

}
