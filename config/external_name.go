/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"avi_actiongroupconfig":               config.IdentifierFromProvider,
	"avi_albservicesconfig":               config.IdentifierFromProvider,
	"avi_albservicesfileupload":           config.IdentifierFromProvider,
	"avi_albservicesjob":                  config.IdentifierFromProvider,
	"avi_alertconfig":                     config.IdentifierFromProvider,
	"avi_alertemailconfig":                config.IdentifierFromProvider,
	"avi_alertscriptconfig":               config.IdentifierFromProvider,
	"avi_alertsyslogconfig":               config.IdentifierFromProvider,
	"avi_analyticsprofile":                config.IdentifierFromProvider,
	"avi_applicationpersistenceprofile":   config.IdentifierFromProvider,
	"avi_applicationprofile":              config.IdentifierFromProvider,
	"avi_authmappingprofile":              config.IdentifierFromProvider,
	"avi_authprofile":                     config.IdentifierFromProvider,
	"avi_autoscalelaunchconfig":           config.IdentifierFromProvider,
	"avi_availabilityzone":                config.IdentifierFromProvider,
	"avi_backup":                          config.IdentifierFromProvider,
	"avi_backupconfiguration":             config.IdentifierFromProvider,
	"avi_botconfigconsolidator":           config.IdentifierFromProvider,
	"avi_botdetectionpolicy":              config.IdentifierFromProvider,
	"avi_botipreputationtypemapping":      config.IdentifierFromProvider,
	"avi_botmapping":                      config.IdentifierFromProvider,
	"avi_certificatemanagementprofile":    config.IdentifierFromProvider,
	"avi_cloud":                           config.IdentifierFromProvider,
	"avi_cloudconnectoruser":              config.IdentifierFromProvider,
	"avi_cloudproperties":                 config.IdentifierFromProvider,
	"avi_cluster":                         config.IdentifierFromProvider,
	"avi_clusterclouddetails":             config.IdentifierFromProvider,
	"avi_controllerportalregistration":    config.IdentifierFromProvider,
	"avi_controllerproperties":            config.IdentifierFromProvider,
	"avi_controllersite":                  config.IdentifierFromProvider,
	"avi_dnspolicy":                       config.IdentifierFromProvider,
	"avi_dynamicdnsrecord":                config.IdentifierFromProvider,
	"avi_errorpagebody":                   config.IdentifierFromProvider,
	"avi_errorpageprofile":                config.IdentifierFromProvider,
	"avi_federationcheckpoint":            config.IdentifierFromProvider,
	"avi_fileobject":                      config.IdentifierFromProvider,
	"avi_fileservice":                     config.IdentifierFromProvider,
	"avi_geodb":                           config.IdentifierFromProvider,
	"avi_gslb":                            config.IdentifierFromProvider,
	"avi_gslbgeodbprofile":                config.IdentifierFromProvider,
	"avi_gslbservice":                     config.IdentifierFromProvider,
	"avi_hardwaresecuritymodulegroup":     config.IdentifierFromProvider,
	"avi_healthmonitor":                   config.IdentifierFromProvider,
	"avi_httppolicyset":                   config.IdentifierFromProvider,
	"avi_icapprofile":                     config.IdentifierFromProvider,
	"avi_image":                           config.IdentifierFromProvider,
	"avi_inventoryfaultconfig":            config.IdentifierFromProvider,
	"avi_ipaddrgroup":                     config.IdentifierFromProvider,
	"avi_ipamdnsproviderprofile":          config.IdentifierFromProvider,
	"avi_ipreputationdb":                  config.IdentifierFromProvider,
	"avi_jwtprofile":                      config.IdentifierFromProvider,
	"avi_jwtserverprofile":                config.IdentifierFromProvider,
	"avi_l4policyset":                     config.IdentifierFromProvider,
	"avi_labelgroup":                      config.IdentifierFromProvider,
	"avi_licenseledgerdetails":            config.IdentifierFromProvider,
	"avi_licensestatus":                   config.IdentifierFromProvider,
	"avi_memorybalancerrequest":           config.IdentifierFromProvider,
	"avi_microservicegroup":               config.IdentifierFromProvider,
	"avi_natpolicy":                       config.IdentifierFromProvider,
	"avi_network":                         config.IdentifierFromProvider,
	"avi_networkprofile":                  config.IdentifierFromProvider,
	"avi_networksecuritypolicy":           config.IdentifierFromProvider,
	"avi_networkservice":                  config.IdentifierFromProvider,
	"avi_nsxtsegmentruntime":              config.IdentifierFromProvider,
	"avi_objectaccesspolicy":              config.IdentifierFromProvider,
	"avi_pingaccessagent":                 config.IdentifierFromProvider,
	"avi_pkiprofile":                      config.IdentifierFromProvider,
	"avi_pool":                            config.IdentifierFromProvider,
	"avi_poolgroup":                       config.IdentifierFromProvider,
	"avi_poolgroupdeploymentpolicy":       config.IdentifierFromProvider,
	"avi_portalfileupload":                config.IdentifierFromProvider,
	"avi_prioritylabels":                  config.IdentifierFromProvider,
	"avi_protocolparser":                  config.IdentifierFromProvider,
	"avi_rmcloudopsproto":                 config.IdentifierFromProvider,
	"avi_role":                            config.IdentifierFromProvider,
	"avi_scheduler":                       config.IdentifierFromProvider,
	"avi_securitymanagerdata":             config.IdentifierFromProvider,
	"avi_securitypolicy":                  config.IdentifierFromProvider,
	"avi_seproperties":                    config.IdentifierFromProvider,
	"avi_server":                          config.IdentifierFromProvider,
	"avi_serverautoscalepolicy":           config.IdentifierFromProvider,
	"avi_serviceengine":                   config.IdentifierFromProvider,
	"avi_serviceenginegroup":              config.IdentifierFromProvider,
	"avi_serviceenginepolicy":             config.IdentifierFromProvider,
	"avi_siteversion":                     config.IdentifierFromProvider,
	"avi_snmptrapprofile":                 config.IdentifierFromProvider,
	"avi_sslkeyandcertificate":            config.IdentifierFromProvider,
	"avi_sslprofile":                      config.IdentifierFromProvider,
	"avi_ssopolicy":                       config.IdentifierFromProvider,
	"avi_statediffoperation":              config.IdentifierFromProvider,
	"avi_statediffsnapshot":               config.IdentifierFromProvider,
	"avi_stringgroup":                     config.IdentifierFromProvider,
	"avi_systemconfiguration":             config.IdentifierFromProvider,
	"avi_systemlimits":                    config.IdentifierFromProvider,
	"avi_tenant":                          config.IdentifierFromProvider,
	"avi_testsedatastorelevel1":           config.IdentifierFromProvider,
	"avi_testsedatastorelevel2":           config.IdentifierFromProvider,
	"avi_testsedatastorelevel3":           config.IdentifierFromProvider,
	"avi_trafficcloneprofile":             config.IdentifierFromProvider,
	"avi_upgradestatusinfo":               config.IdentifierFromProvider,
	"avi_upgradestatussummary":            config.IdentifierFromProvider,
	"avi_user":                            config.IdentifierFromProvider,
	"avi_useraccount":                     config.IdentifierFromProvider,
	"avi_useraccountprofile":              config.IdentifierFromProvider,
	"avi_vcenterserver":                   config.IdentifierFromProvider,
	"avi_virtualservice":                  config.IdentifierFromProvider,
	"avi_vrfcontext":                      config.IdentifierFromProvider,
	"avi_vsdatascriptset":                 config.IdentifierFromProvider,
	"avi_vsgs":                            config.IdentifierFromProvider,
	"avi_vsvip":                           config.IdentifierFromProvider,
	"avi_wafapplicationsignatureprovider": config.IdentifierFromProvider,
	"avi_wafcrs":                          config.IdentifierFromProvider,
	"avi_wafpolicy":                       config.IdentifierFromProvider,
	"avi_wafpolicypsmgroup":               config.IdentifierFromProvider,
	"avi_wafprofile":                      config.IdentifierFromProvider,
	"avi_webhook":                         config.IdentifierFromProvider,
	"avi_webapput":                        config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
