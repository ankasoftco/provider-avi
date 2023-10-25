/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"avi_actiongroupconfig":                                  config.IdentifierFromProvider,
	"avi_albservicesconfig":                                 config.IdentifierFromProvider,
	"avi_albservicesfileupload":                      config.IdentifierFromProvider,
	"avi_albservicesjob":                      config.IdentifierFromProvider,
	"avi_alertconfig":                      	   config.IdentifierFromProvider,
	"avi_alertemailconfig":                       config.IdentifierFromProvider,
	"avi_alertscriptconfig":                                 config.IdentifierFromProvider,
	"avi_alertsyslogconfig":                                 config.IdentifierFromProvider,
	"avi_analyticsprofile":                        config.IdentifierFromProvider,
	"avi_applicationpersistenceprofile":                                    config.IdentifierFromProvider,
	"avi_applicationprofile":                             config.IdentifierFromProvider,
	"avi_authmappingprofile":            config.IdentifierFromProvider,
	"avi_authprofile":                              config.IdentifierFromProvider,
	"avi_autoscalelaunchconfig":                    config.IdentifierFromProvider,
	"avi_availabilityzone":                   config.IdentifierFromProvider,
	"avi_backup":                   config.IdentifierFromProvider,
	"avi_backupconfiguration":                  config.IdentifierFromProvider,
	"avi_botconfigconsolidator":                             config.IdentifierFromProvider,
	"avi_botdetectionpolicy":                                     config.IdentifierFromProvider,
	"avi_botipreputationtypemapping":                                  config.IdentifierFromProvider,
	"avi_botmapping":                                config.IdentifierFromProvider,
	"avi_certificatemanagementprofile":                      config.IdentifierFromProvider,
	"avi_cloud":                         config.IdentifierFromProvider,
	"avi_cloudconnectoruser":                         config.IdentifierFromProvider,
	"avi_cloudproperties":               config.IdentifierFromProvider,
	"avi_cluster":                           config.IdentifierFromProvider,
	"avi_clusterclouddetails":                       config.IdentifierFromProvider,
	"avi_controllerportalregistration":                                    config.IdentifierFromProvider,
	"avi_controllerproperties":                                 config.IdentifierFromProvider,
	"avi_controllersite":                                config.IdentifierFromProvider,
	"avi_dnspolicy":                               config.IdentifierFromProvider,
	"avi_dynamicdnsrecord":                 config.IdentifierFromProvider,
	"avi_errorpagebody":                     config.IdentifierFromProvider,
	"avi_errorpageprofile":                               config.IdentifierFromProvider,
	"avi_federationcheckpoint":                     config.IdentifierFromProvider,
	"avi_fileobject":                                      config.IdentifierFromProvider,
	"avi_fileservice":                                       config.IdentifierFromProvider,
	"avi_geodb":                                     config.IdentifierFromProvider,
	"avi_gslb":                       config.IdentifierFromProvider,
	"avi_gslbgeodbprofile":                                    config.IdentifierFromProvider,
	"avi_gslbservice":                                     config.IdentifierFromProvider,
	"avi_hardwaresecuritymodulegroup":                        config.IdentifierFromProvider,
	"avi_healthmonitor":                       config.IdentifierFromProvider,
	"avi_httppolicyset":                                     config.IdentifierFromProvider,
	"avi_icapprofile":                                  config.IdentifierFromProvider,
	"avi_image":                           config.IdentifierFromProvider,
	"avi_inventoryfaultconfig":                                     config.IdentifierFromProvider,
	"avi_ipaddrgroup":                               config.IdentifierFromProvider,
	"avi_ipamdnsproviderprofile":                        config.IdentifierFromProvider,
	"avi_ipreputationdb":                       config.IdentifierFromProvider,
	"avi_jwtprofile":                                 config.IdentifierFromProvider,
	"avi_jwtserverprofile":                                     config.IdentifierFromProvider,
	"avi_l4policyset":                    config.IdentifierFromProvider,
	"avi_labelgroup":                         config.IdentifierFromProvider,
	"avi_licenseledgerdetails":                                config.IdentifierFromProvider,
	"avi_licensestatus":                                config.IdentifierFromProvider,
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
