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

}