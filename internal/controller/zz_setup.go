/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	actiongroupconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/actiongroupconfig"
	albservicesconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/albservicesconfig"
	albservicesfileupload "github.com/ankasoftco/provider-avi/internal/controller/avi/albservicesfileupload"
	albservicesjob "github.com/ankasoftco/provider-avi/internal/controller/avi/albservicesjob"
	alertconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/alertconfig"
	alertemailconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/alertemailconfig"
	alertscriptconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/alertscriptconfig"
	alertsyslogconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/alertsyslogconfig"
	analyticsprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/analyticsprofile"
	applicationpersistenceprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/applicationpersistenceprofile"
	applicationprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/applicationprofile"
	authmappingprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/authmappingprofile"
	authprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/authprofile"
	autoscalelaunchconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/autoscalelaunchconfig"
	availabilityzone "github.com/ankasoftco/provider-avi/internal/controller/avi/availabilityzone"
	backup "github.com/ankasoftco/provider-avi/internal/controller/avi/backup"
	backupconfiguration "github.com/ankasoftco/provider-avi/internal/controller/avi/backupconfiguration"
	botconfigconsolidator "github.com/ankasoftco/provider-avi/internal/controller/avi/botconfigconsolidator"
	botdetectionpolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/botdetectionpolicy"
	botipreputationtypemapping "github.com/ankasoftco/provider-avi/internal/controller/avi/botipreputationtypemapping"
	botmal4policyset "github.com/ankasoftco/provider-avi/internal/controller/avi/botmal4policyset"
	botmapping "github.com/ankasoftco/provider-avi/internal/controller/avi/botmapping"
	certificatemanagementprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/certificatemanagementprofile"
	cloud "github.com/ankasoftco/provider-avi/internal/controller/avi/cloud"
	cloudconnectoruser "github.com/ankasoftco/provider-avi/internal/controller/avi/cloudconnectoruser"
	cloudproperties "github.com/ankasoftco/provider-avi/internal/controller/avi/cloudproperties"
	cluster "github.com/ankasoftco/provider-avi/internal/controller/avi/cluster"
	clusterclouddetails "github.com/ankasoftco/provider-avi/internal/controller/avi/clusterclouddetails"
	controllerportalregistration "github.com/ankasoftco/provider-avi/internal/controller/avi/controllerportalregistration"
	controllerproperties "github.com/ankasoftco/provider-avi/internal/controller/avi/controllerproperties"
	controllersite "github.com/ankasoftco/provider-avi/internal/controller/avi/controllersite"
	dnspolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/dnspolicy"
	dynamicdnsrecord "github.com/ankasoftco/provider-avi/internal/controller/avi/dynamicdnsrecord"
	errorpagebody "github.com/ankasoftco/provider-avi/internal/controller/avi/errorpagebody"
	errorpageprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/errorpageprofile"
	federationcheckpoint "github.com/ankasoftco/provider-avi/internal/controller/avi/federationcheckpoint"
	fileobject "github.com/ankasoftco/provider-avi/internal/controller/avi/fileobject"
	fileservice "github.com/ankasoftco/provider-avi/internal/controller/avi/fileservice"
	geodb "github.com/ankasoftco/provider-avi/internal/controller/avi/geodb"
	gslb "github.com/ankasoftco/provider-avi/internal/controller/avi/gslb"
	gslbgeodbprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/gslbgeodbprofile"
	gslbservice "github.com/ankasoftco/provider-avi/internal/controller/avi/gslbservice"
	hardwaresecuritymodulegroup "github.com/ankasoftco/provider-avi/internal/controller/avi/hardwaresecuritymodulegroup"
	healthmonitor "github.com/ankasoftco/provider-avi/internal/controller/avi/healthmonitor"
	httppolicyset "github.com/ankasoftco/provider-avi/internal/controller/avi/httppolicyset"
	icapprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/icapprofile"
	image "github.com/ankasoftco/provider-avi/internal/controller/avi/image"
	inventoryfaultconfig "github.com/ankasoftco/provider-avi/internal/controller/avi/inventoryfaultconfig"
	ipaddrgroup "github.com/ankasoftco/provider-avi/internal/controller/avi/ipaddrgroup"
	ipamdnsproviderprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/ipamdnsproviderprofile"
	ipreputationdb "github.com/ankasoftco/provider-avi/internal/controller/avi/ipreputationdb"
	jwtserverprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/jwtserverprofile"
	labelgroup "github.com/ankasoftco/provider-avi/internal/controller/avi/labelgroup"
	licenseledgerdetails "github.com/ankasoftco/provider-avi/internal/controller/avi/licenseledgerdetails"
	licensestatus "github.com/ankasoftco/provider-avi/internal/controller/avi/licensestatus"
	memorybalancerrequest "github.com/ankasoftco/provider-avi/internal/controller/avi/memorybalancerrequest"
	microservicegroup "github.com/ankasoftco/provider-avi/internal/controller/avi/microservicegroup"
	natpolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/natpolicy"
	network "github.com/ankasoftco/provider-avi/internal/controller/avi/network"
	networkprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/networkprofile"
	networksecuritypolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/networksecuritypolicy"
	networkservice "github.com/ankasoftco/provider-avi/internal/controller/avi/networkservice"
	nsxtsegmentruntime "github.com/ankasoftco/provider-avi/internal/controller/avi/nsxtsegmentruntime"
	pingaccessagent "github.com/ankasoftco/provider-avi/internal/controller/avi/pingaccessagent"
	pkiprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/pkiprofile"
	pool "github.com/ankasoftco/provider-avi/internal/controller/avi/pool"
	poolgroup "github.com/ankasoftco/provider-avi/internal/controller/avi/poolgroup"
	poolgroupdeploymentpolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/poolgroupdeploymentpolicy"
	prioritylabels "github.com/ankasoftco/provider-avi/internal/controller/avi/prioritylabels"
	protocolparser "github.com/ankasoftco/provider-avi/internal/controller/avi/protocolparser"
	rmcloudopsproto "github.com/ankasoftco/provider-avi/internal/controller/avi/rmcloudopsproto"
	role "github.com/ankasoftco/provider-avi/internal/controller/avi/role"
	scheduler "github.com/ankasoftco/provider-avi/internal/controller/avi/scheduler"
	securitymanagerdata "github.com/ankasoftco/provider-avi/internal/controller/avi/securitymanagerdata"
	securitypolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/securitypolicy"
	seproperties "github.com/ankasoftco/provider-avi/internal/controller/avi/seproperties"
	server "github.com/ankasoftco/provider-avi/internal/controller/avi/server"
	serverautoscalepolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/serverautoscalepolicy"
	serviceengine "github.com/ankasoftco/provider-avi/internal/controller/avi/serviceengine"
	serviceenginegroup "github.com/ankasoftco/provider-avi/internal/controller/avi/serviceenginegroup"
	siteversion "github.com/ankasoftco/provider-avi/internal/controller/avi/siteversion"
	snmptrapprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/snmptrapprofile"
	sslkeyandcertificate "github.com/ankasoftco/provider-avi/internal/controller/avi/sslkeyandcertificate"
	sslprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/sslprofile"
	ssopolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/ssopolicy"
	statediffoperation "github.com/ankasoftco/provider-avi/internal/controller/avi/statediffoperation"
	statediffsnapshot "github.com/ankasoftco/provider-avi/internal/controller/avi/statediffsnapshot"
	stringgroup "github.com/ankasoftco/provider-avi/internal/controller/avi/stringgroup"
	systemconfiguration "github.com/ankasoftco/provider-avi/internal/controller/avi/systemconfiguration"
	systemlimits "github.com/ankasoftco/provider-avi/internal/controller/avi/systemlimits"
	tenant "github.com/ankasoftco/provider-avi/internal/controller/avi/tenant"
	testsedatastorelevel1 "github.com/ankasoftco/provider-avi/internal/controller/avi/testsedatastorelevel1"
	testsedatastorelevel2 "github.com/ankasoftco/provider-avi/internal/controller/avi/testsedatastorelevel2"
	testsedatastorelevel3 "github.com/ankasoftco/provider-avi/internal/controller/avi/testsedatastorelevel3"
	trafficcloneprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/trafficcloneprofile"
	upgradestatusinfo "github.com/ankasoftco/provider-avi/internal/controller/avi/upgradestatusinfo"
	upgradestatussummary "github.com/ankasoftco/provider-avi/internal/controller/avi/upgradestatussummary"
	user "github.com/ankasoftco/provider-avi/internal/controller/avi/user"
	useraccount "github.com/ankasoftco/provider-avi/internal/controller/avi/useraccount"
	useraccountprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/useraccountprofile"
	vcenterserver "github.com/ankasoftco/provider-avi/internal/controller/avi/vcenterserver"
	virtualservice "github.com/ankasoftco/provider-avi/internal/controller/avi/virtualservice"
	vrfcontext "github.com/ankasoftco/provider-avi/internal/controller/avi/vrfcontext"
	vsdatascriptset "github.com/ankasoftco/provider-avi/internal/controller/avi/vsdatascriptset"
	vsgs "github.com/ankasoftco/provider-avi/internal/controller/avi/vsgs"
	vsvip "github.com/ankasoftco/provider-avi/internal/controller/avi/vsvip"
	wafapplicationsignatureprovider "github.com/ankasoftco/provider-avi/internal/controller/avi/wafapplicationsignatureprovider"
	wafcrs "github.com/ankasoftco/provider-avi/internal/controller/avi/wafcrs"
	wafpolicy "github.com/ankasoftco/provider-avi/internal/controller/avi/wafpolicy"
	wafpolicypsmgroup "github.com/ankasoftco/provider-avi/internal/controller/avi/wafpolicypsmgroup"
	wafprofile "github.com/ankasoftco/provider-avi/internal/controller/avi/wafprofile"
	webapput "github.com/ankasoftco/provider-avi/internal/controller/avi/webapput"
	webhook "github.com/ankasoftco/provider-avi/internal/controller/avi/webhook"
	providerconfig "github.com/ankasoftco/provider-avi/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		actiongroupconfig.Setup,
		albservicesconfig.Setup,
		albservicesfileupload.Setup,
		albservicesjob.Setup,
		alertconfig.Setup,
		alertemailconfig.Setup,
		alertscriptconfig.Setup,
		alertsyslogconfig.Setup,
		analyticsprofile.Setup,
		applicationpersistenceprofile.Setup,
		applicationprofile.Setup,
		authmappingprofile.Setup,
		authprofile.Setup,
		autoscalelaunchconfig.Setup,
		availabilityzone.Setup,
		backup.Setup,
		backupconfiguration.Setup,
		botconfigconsolidator.Setup,
		botdetectionpolicy.Setup,
		botipreputationtypemapping.Setup,
		botmal4policyset.Setup,
		botmapping.Setup,
		certificatemanagementprofile.Setup,
		cloud.Setup,
		cloudconnectoruser.Setup,
		cloudproperties.Setup,
		cluster.Setup,
		clusterclouddetails.Setup,
		controllerportalregistration.Setup,
		controllerproperties.Setup,
		controllersite.Setup,
		dnspolicy.Setup,
		dynamicdnsrecord.Setup,
		errorpagebody.Setup,
		errorpageprofile.Setup,
		federationcheckpoint.Setup,
		fileobject.Setup,
		fileservice.Setup,
		geodb.Setup,
		gslb.Setup,
		gslbgeodbprofile.Setup,
		gslbservice.Setup,
		hardwaresecuritymodulegroup.Setup,
		healthmonitor.Setup,
		httppolicyset.Setup,
		icapprofile.Setup,
		image.Setup,
		inventoryfaultconfig.Setup,
		ipaddrgroup.Setup,
		ipamdnsproviderprofile.Setup,
		ipreputationdb.Setup,
		jwtserverprofile.Setup,
		labelgroup.Setup,
		licenseledgerdetails.Setup,
		licensestatus.Setup,
		memorybalancerrequest.Setup,
		microservicegroup.Setup,
		natpolicy.Setup,
		network.Setup,
		networkprofile.Setup,
		networksecuritypolicy.Setup,
		networkservice.Setup,
		nsxtsegmentruntime.Setup,
		pingaccessagent.Setup,
		pkiprofile.Setup,
		pool.Setup,
		poolgroup.Setup,
		poolgroupdeploymentpolicy.Setup,
		prioritylabels.Setup,
		protocolparser.Setup,
		rmcloudopsproto.Setup,
		role.Setup,
		scheduler.Setup,
		securitymanagerdata.Setup,
		securitypolicy.Setup,
		seproperties.Setup,
		server.Setup,
		serverautoscalepolicy.Setup,
		serviceengine.Setup,
		serviceenginegroup.Setup,
		siteversion.Setup,
		snmptrapprofile.Setup,
		sslkeyandcertificate.Setup,
		sslprofile.Setup,
		ssopolicy.Setup,
		statediffoperation.Setup,
		statediffsnapshot.Setup,
		stringgroup.Setup,
		systemconfiguration.Setup,
		systemlimits.Setup,
		tenant.Setup,
		testsedatastorelevel1.Setup,
		testsedatastorelevel2.Setup,
		testsedatastorelevel3.Setup,
		trafficcloneprofile.Setup,
		upgradestatusinfo.Setup,
		upgradestatussummary.Setup,
		user.Setup,
		useraccount.Setup,
		useraccountprofile.Setup,
		vcenterserver.Setup,
		virtualservice.Setup,
		vrfcontext.Setup,
		vsdatascriptset.Setup,
		vsgs.Setup,
		vsvip.Setup,
		wafapplicationsignatureprovider.Setup,
		wafcrs.Setup,
		wafpolicy.Setup,
		wafpolicypsmgroup.Setup,
		wafprofile.Setup,
		webapput.Setup,
		webhook.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
