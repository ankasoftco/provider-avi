/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BotLimitsInitParameters struct {
	AllowRules *string `json:"allowRules,omitempty" tf:"allow_rules,omitempty"`

	Hdrs *string `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	MappingRules *string `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`
}

type BotLimitsObservation struct {
	AllowRules *string `json:"allowRules,omitempty" tf:"allow_rules,omitempty"`

	Hdrs *string `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	MappingRules *string `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`
}

type BotLimitsParameters struct {

	// +kubebuilder:validation:Optional
	AllowRules *string `json:"allowRules,omitempty" tf:"allow_rules,omitempty"`

	// +kubebuilder:validation:Optional
	Hdrs *string `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	// +kubebuilder:validation:Optional
	MappingRules *string `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`
}

type ControllerCloudLimitsInitParameters struct {
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerCloudLimitsObservation struct {
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerCloudLimitsParameters struct {

	// +kubebuilder:validation:Optional
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerLimitsInitParameters struct {
	BotLimits []BotLimitsInitParameters `json:"botLimits,omitempty" tf:"bot_limits,omitempty"`

	CertificatesPerVirtualservice *string `json:"certificatesPerVirtualservice,omitempty" tf:"certificates_per_virtualservice,omitempty"`

	ControllerCloudLimits []ControllerCloudLimitsInitParameters `json:"controllerCloudLimits,omitempty" tf:"controller_cloud_limits,omitempty"`

	ControllerSizingLimits []ControllerSizingLimitsInitParameters `json:"controllerSizingLimits,omitempty" tf:"controller_sizing_limits,omitempty"`

	DefaultRoutesPerVrfcontext *string `json:"defaultRoutesPerVrfcontext,omitempty" tf:"default_routes_per_vrfcontext,omitempty"`

	GatewayMonPerVrf *string `json:"gatewayMonPerVrf,omitempty" tf:"gateway_mon_per_vrf,omitempty"`

	IpaddressLimits []IpaddressLimitsInitParameters `json:"ipaddressLimits,omitempty" tf:"ipaddress_limits,omitempty"`

	IpsPerIpgroup *string `json:"ipsPerIpgroup,omitempty" tf:"ips_per_ipgroup,omitempty"`

	L7Limits []L7LimitsInitParameters `json:"l7Limits,omitempty" tf:"l7_limits,omitempty"`

	PoolgroupsPerVirtualservice *string `json:"poolgroupsPerVirtualservice,omitempty" tf:"poolgroups_per_virtualservice,omitempty"`

	PoolsPerPoolgroup *string `json:"poolsPerPoolgroup,omitempty" tf:"pools_per_poolgroup,omitempty"`

	PoolsPerVirtualservice *string `json:"poolsPerVirtualservice,omitempty" tf:"pools_per_virtualservice,omitempty"`

	RoutesPerVrfcontext *string `json:"routesPerVrfcontext,omitempty" tf:"routes_per_vrfcontext,omitempty"`

	RulesPerNATPolicy *string `json:"rulesPerNatPolicy,omitempty" tf:"rules_per_nat_policy,omitempty"`

	RulesPerNetworksecuritypolicy *string `json:"rulesPerNetworksecuritypolicy,omitempty" tf:"rules_per_networksecuritypolicy,omitempty"`

	ServersPerPool *string `json:"serversPerPool,omitempty" tf:"servers_per_pool,omitempty"`

	SniChildrenPerParent *string `json:"sniChildrenPerParent,omitempty" tf:"sni_children_per_parent,omitempty"`

	StringsPerStringgroup *string `json:"stringsPerStringgroup,omitempty" tf:"strings_per_stringgroup,omitempty"`

	VsBGPScaleout *string `json:"vsBgpScaleout,omitempty" tf:"vs_bgp_scaleout,omitempty"`

	VsL2Scaleout *string `json:"vsL2Scaleout,omitempty" tf:"vs_l2_scaleout,omitempty"`

	WafLimits []WafLimitsInitParameters `json:"wafLimits,omitempty" tf:"waf_limits,omitempty"`
}

type ControllerLimitsObservation struct {
	BotLimits []BotLimitsObservation `json:"botLimits,omitempty" tf:"bot_limits,omitempty"`

	CertificatesPerVirtualservice *string `json:"certificatesPerVirtualservice,omitempty" tf:"certificates_per_virtualservice,omitempty"`

	ControllerCloudLimits []ControllerCloudLimitsObservation `json:"controllerCloudLimits,omitempty" tf:"controller_cloud_limits,omitempty"`

	ControllerSizingLimits []ControllerSizingLimitsObservation `json:"controllerSizingLimits,omitempty" tf:"controller_sizing_limits,omitempty"`

	DefaultRoutesPerVrfcontext *string `json:"defaultRoutesPerVrfcontext,omitempty" tf:"default_routes_per_vrfcontext,omitempty"`

	GatewayMonPerVrf *string `json:"gatewayMonPerVrf,omitempty" tf:"gateway_mon_per_vrf,omitempty"`

	IpaddressLimits []IpaddressLimitsObservation `json:"ipaddressLimits,omitempty" tf:"ipaddress_limits,omitempty"`

	IpsPerIpgroup *string `json:"ipsPerIpgroup,omitempty" tf:"ips_per_ipgroup,omitempty"`

	L7Limits []L7LimitsObservation `json:"l7Limits,omitempty" tf:"l7_limits,omitempty"`

	PoolgroupsPerVirtualservice *string `json:"poolgroupsPerVirtualservice,omitempty" tf:"poolgroups_per_virtualservice,omitempty"`

	PoolsPerPoolgroup *string `json:"poolsPerPoolgroup,omitempty" tf:"pools_per_poolgroup,omitempty"`

	PoolsPerVirtualservice *string `json:"poolsPerVirtualservice,omitempty" tf:"pools_per_virtualservice,omitempty"`

	RoutesPerVrfcontext *string `json:"routesPerVrfcontext,omitempty" tf:"routes_per_vrfcontext,omitempty"`

	RulesPerNATPolicy *string `json:"rulesPerNatPolicy,omitempty" tf:"rules_per_nat_policy,omitempty"`

	RulesPerNetworksecuritypolicy *string `json:"rulesPerNetworksecuritypolicy,omitempty" tf:"rules_per_networksecuritypolicy,omitempty"`

	ServersPerPool *string `json:"serversPerPool,omitempty" tf:"servers_per_pool,omitempty"`

	SniChildrenPerParent *string `json:"sniChildrenPerParent,omitempty" tf:"sni_children_per_parent,omitempty"`

	StringsPerStringgroup *string `json:"stringsPerStringgroup,omitempty" tf:"strings_per_stringgroup,omitempty"`

	VsBGPScaleout *string `json:"vsBgpScaleout,omitempty" tf:"vs_bgp_scaleout,omitempty"`

	VsL2Scaleout *string `json:"vsL2Scaleout,omitempty" tf:"vs_l2_scaleout,omitempty"`

	WafLimits []WafLimitsObservation `json:"wafLimits,omitempty" tf:"waf_limits,omitempty"`
}

type ControllerLimitsParameters struct {

	// +kubebuilder:validation:Optional
	BotLimits []BotLimitsParameters `json:"botLimits,omitempty" tf:"bot_limits,omitempty"`

	// +kubebuilder:validation:Optional
	CertificatesPerVirtualservice *string `json:"certificatesPerVirtualservice,omitempty" tf:"certificates_per_virtualservice,omitempty"`

	// +kubebuilder:validation:Optional
	ControllerCloudLimits []ControllerCloudLimitsParameters `json:"controllerCloudLimits,omitempty" tf:"controller_cloud_limits,omitempty"`

	// +kubebuilder:validation:Optional
	ControllerSizingLimits []ControllerSizingLimitsParameters `json:"controllerSizingLimits,omitempty" tf:"controller_sizing_limits,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRoutesPerVrfcontext *string `json:"defaultRoutesPerVrfcontext,omitempty" tf:"default_routes_per_vrfcontext,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayMonPerVrf *string `json:"gatewayMonPerVrf,omitempty" tf:"gateway_mon_per_vrf,omitempty"`

	// +kubebuilder:validation:Optional
	IpaddressLimits []IpaddressLimitsParameters `json:"ipaddressLimits,omitempty" tf:"ipaddress_limits,omitempty"`

	// +kubebuilder:validation:Optional
	IpsPerIpgroup *string `json:"ipsPerIpgroup,omitempty" tf:"ips_per_ipgroup,omitempty"`

	// +kubebuilder:validation:Optional
	L7Limits []L7LimitsParameters `json:"l7Limits,omitempty" tf:"l7_limits,omitempty"`

	// +kubebuilder:validation:Optional
	PoolgroupsPerVirtualservice *string `json:"poolgroupsPerVirtualservice,omitempty" tf:"poolgroups_per_virtualservice,omitempty"`

	// +kubebuilder:validation:Optional
	PoolsPerPoolgroup *string `json:"poolsPerPoolgroup,omitempty" tf:"pools_per_poolgroup,omitempty"`

	// +kubebuilder:validation:Optional
	PoolsPerVirtualservice *string `json:"poolsPerVirtualservice,omitempty" tf:"pools_per_virtualservice,omitempty"`

	// +kubebuilder:validation:Optional
	RoutesPerVrfcontext *string `json:"routesPerVrfcontext,omitempty" tf:"routes_per_vrfcontext,omitempty"`

	// +kubebuilder:validation:Optional
	RulesPerNATPolicy *string `json:"rulesPerNatPolicy,omitempty" tf:"rules_per_nat_policy,omitempty"`

	// +kubebuilder:validation:Optional
	RulesPerNetworksecuritypolicy *string `json:"rulesPerNetworksecuritypolicy,omitempty" tf:"rules_per_networksecuritypolicy,omitempty"`

	// +kubebuilder:validation:Optional
	ServersPerPool *string `json:"serversPerPool,omitempty" tf:"servers_per_pool,omitempty"`

	// +kubebuilder:validation:Optional
	SniChildrenPerParent *string `json:"sniChildrenPerParent,omitempty" tf:"sni_children_per_parent,omitempty"`

	// +kubebuilder:validation:Optional
	StringsPerStringgroup *string `json:"stringsPerStringgroup,omitempty" tf:"strings_per_stringgroup,omitempty"`

	// +kubebuilder:validation:Optional
	VsBGPScaleout *string `json:"vsBgpScaleout,omitempty" tf:"vs_bgp_scaleout,omitempty"`

	// +kubebuilder:validation:Optional
	VsL2Scaleout *string `json:"vsL2Scaleout,omitempty" tf:"vs_l2_scaleout,omitempty"`

	// +kubebuilder:validation:Optional
	WafLimits []WafLimitsParameters `json:"wafLimits,omitempty" tf:"waf_limits,omitempty"`
}

type ControllerSizesInitParameters struct {
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	MinCpus *string `json:"minCpus,omitempty" tf:"min_cpus,omitempty"`

	MinMemory *string `json:"minMemory,omitempty" tf:"min_memory,omitempty"`
}

type ControllerSizesObservation struct {
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	MinCpus *string `json:"minCpus,omitempty" tf:"min_cpus,omitempty"`

	MinMemory *string `json:"minMemory,omitempty" tf:"min_memory,omitempty"`
}

type ControllerSizesParameters struct {

	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// +kubebuilder:validation:Optional
	MinCpus *string `json:"minCpus,omitempty" tf:"min_cpus,omitempty"`

	// +kubebuilder:validation:Optional
	MinMemory *string `json:"minMemory,omitempty" tf:"min_memory,omitempty"`
}

type ControllerSizingCloudLimitsInitParameters struct {
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerSizingCloudLimitsObservation struct {
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerSizingCloudLimitsParameters struct {

	// +kubebuilder:validation:Optional
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ControllerSizingLimitsInitParameters struct {
	ControllerSizingCloudLimits []ControllerSizingCloudLimitsInitParameters `json:"controllerSizingCloudLimits,omitempty" tf:"controller_sizing_cloud_limits,omitempty"`

	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	NumEastWestVirtualservices *string `json:"numEastWestVirtualservices,omitempty" tf:"num_east_west_virtualservices,omitempty"`

	NumServers *string `json:"numServers,omitempty" tf:"num_servers,omitempty"`

	NumServiceengines *string `json:"numServiceengines,omitempty" tf:"num_serviceengines,omitempty"`

	NumTenants *string `json:"numTenants,omitempty" tf:"num_tenants,omitempty"`

	NumVirtualservices *string `json:"numVirtualservices,omitempty" tf:"num_virtualservices,omitempty"`

	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	NumVrfs *string `json:"numVrfs,omitempty" tf:"num_vrfs,omitempty"`
}

type ControllerSizingLimitsObservation struct {
	ControllerSizingCloudLimits []ControllerSizingCloudLimitsObservation `json:"controllerSizingCloudLimits,omitempty" tf:"controller_sizing_cloud_limits,omitempty"`

	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	NumEastWestVirtualservices *string `json:"numEastWestVirtualservices,omitempty" tf:"num_east_west_virtualservices,omitempty"`

	NumServers *string `json:"numServers,omitempty" tf:"num_servers,omitempty"`

	NumServiceengines *string `json:"numServiceengines,omitempty" tf:"num_serviceengines,omitempty"`

	NumTenants *string `json:"numTenants,omitempty" tf:"num_tenants,omitempty"`

	NumVirtualservices *string `json:"numVirtualservices,omitempty" tf:"num_virtualservices,omitempty"`

	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	NumVrfs *string `json:"numVrfs,omitempty" tf:"num_vrfs,omitempty"`
}

type ControllerSizingLimitsParameters struct {

	// +kubebuilder:validation:Optional
	ControllerSizingCloudLimits []ControllerSizingCloudLimitsParameters `json:"controllerSizingCloudLimits,omitempty" tf:"controller_sizing_cloud_limits,omitempty"`

	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// +kubebuilder:validation:Optional
	NumClouds *string `json:"numClouds,omitempty" tf:"num_clouds,omitempty"`

	// +kubebuilder:validation:Optional
	NumEastWestVirtualservices *string `json:"numEastWestVirtualservices,omitempty" tf:"num_east_west_virtualservices,omitempty"`

	// +kubebuilder:validation:Optional
	NumServers *string `json:"numServers,omitempty" tf:"num_servers,omitempty"`

	// +kubebuilder:validation:Optional
	NumServiceengines *string `json:"numServiceengines,omitempty" tf:"num_serviceengines,omitempty"`

	// +kubebuilder:validation:Optional
	NumTenants *string `json:"numTenants,omitempty" tf:"num_tenants,omitempty"`

	// +kubebuilder:validation:Optional
	NumVirtualservices *string `json:"numVirtualservices,omitempty" tf:"num_virtualservices,omitempty"`

	// +kubebuilder:validation:Optional
	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	// +kubebuilder:validation:Optional
	NumVrfs *string `json:"numVrfs,omitempty" tf:"num_vrfs,omitempty"`
}

type IpaddressLimitsInitParameters struct {
	IPAddressGroupPerMatchCriteria *string `json:"ipAddressGroupPerMatchCriteria,omitempty" tf:"ip_address_group_per_match_criteria,omitempty"`

	IPAddressPrefixPerMatchCriteria *string `json:"ipAddressPrefixPerMatchCriteria,omitempty" tf:"ip_address_prefix_per_match_criteria,omitempty"`

	IPAddressRangePerMatchCriteria *string `json:"ipAddressRangePerMatchCriteria,omitempty" tf:"ip_address_range_per_match_criteria,omitempty"`

	IPAddressesPerMatchCriteria *string `json:"ipAddressesPerMatchCriteria,omitempty" tf:"ip_addresses_per_match_criteria,omitempty"`
}

type IpaddressLimitsObservation struct {
	IPAddressGroupPerMatchCriteria *string `json:"ipAddressGroupPerMatchCriteria,omitempty" tf:"ip_address_group_per_match_criteria,omitempty"`

	IPAddressPrefixPerMatchCriteria *string `json:"ipAddressPrefixPerMatchCriteria,omitempty" tf:"ip_address_prefix_per_match_criteria,omitempty"`

	IPAddressRangePerMatchCriteria *string `json:"ipAddressRangePerMatchCriteria,omitempty" tf:"ip_address_range_per_match_criteria,omitempty"`

	IPAddressesPerMatchCriteria *string `json:"ipAddressesPerMatchCriteria,omitempty" tf:"ip_addresses_per_match_criteria,omitempty"`
}

type IpaddressLimitsParameters struct {

	// +kubebuilder:validation:Optional
	IPAddressGroupPerMatchCriteria *string `json:"ipAddressGroupPerMatchCriteria,omitempty" tf:"ip_address_group_per_match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressPrefixPerMatchCriteria *string `json:"ipAddressPrefixPerMatchCriteria,omitempty" tf:"ip_address_prefix_per_match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressRangePerMatchCriteria *string `json:"ipAddressRangePerMatchCriteria,omitempty" tf:"ip_address_range_per_match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressesPerMatchCriteria *string `json:"ipAddressesPerMatchCriteria,omitempty" tf:"ip_addresses_per_match_criteria,omitempty"`
}

type L7LimitsInitParameters struct {
	HTTPPoliciesPerVs *string `json:"httpPoliciesPerVs,omitempty" tf:"http_policies_per_vs,omitempty"`

	NumCompressionFilters *string `json:"numCompressionFilters,omitempty" tf:"num_compression_filters,omitempty"`

	NumCustomStr *string `json:"numCustomStr,omitempty" tf:"num_custom_str,omitempty"`

	NumMatchesPerRule *string `json:"numMatchesPerRule,omitempty" tf:"num_matches_per_rule,omitempty"`

	NumRulesPerEvhHost *string `json:"numRulesPerEvhHost,omitempty" tf:"num_rules_per_evh_host,omitempty"`

	NumRulesPerHTTPPolicy *string `json:"numRulesPerHttpPolicy,omitempty" tf:"num_rules_per_http_policy,omitempty"`

	NumStrgroupsPerMatch *string `json:"numStrgroupsPerMatch,omitempty" tf:"num_strgroups_per_match,omitempty"`

	StrCacheMime *string `json:"strCacheMime,omitempty" tf:"str_cache_mime,omitempty"`

	StrGroupsCacheMime *string `json:"strGroupsCacheMime,omitempty" tf:"str_groups_cache_mime,omitempty"`

	StrGroupsNoCacheMime *string `json:"strGroupsNoCacheMime,omitempty" tf:"str_groups_no_cache_mime,omitempty"`

	StrGroupsNoCacheURI *string `json:"strGroupsNoCacheUri,omitempty" tf:"str_groups_no_cache_uri,omitempty"`

	StrNoCacheMime *string `json:"strNoCacheMime,omitempty" tf:"str_no_cache_mime,omitempty"`

	StrNoCacheURI *string `json:"strNoCacheUri,omitempty" tf:"str_no_cache_uri,omitempty"`
}

type L7LimitsObservation struct {
	HTTPPoliciesPerVs *string `json:"httpPoliciesPerVs,omitempty" tf:"http_policies_per_vs,omitempty"`

	NumCompressionFilters *string `json:"numCompressionFilters,omitempty" tf:"num_compression_filters,omitempty"`

	NumCustomStr *string `json:"numCustomStr,omitempty" tf:"num_custom_str,omitempty"`

	NumMatchesPerRule *string `json:"numMatchesPerRule,omitempty" tf:"num_matches_per_rule,omitempty"`

	NumRulesPerEvhHost *string `json:"numRulesPerEvhHost,omitempty" tf:"num_rules_per_evh_host,omitempty"`

	NumRulesPerHTTPPolicy *string `json:"numRulesPerHttpPolicy,omitempty" tf:"num_rules_per_http_policy,omitempty"`

	NumStrgroupsPerMatch *string `json:"numStrgroupsPerMatch,omitempty" tf:"num_strgroups_per_match,omitempty"`

	StrCacheMime *string `json:"strCacheMime,omitempty" tf:"str_cache_mime,omitempty"`

	StrGroupsCacheMime *string `json:"strGroupsCacheMime,omitempty" tf:"str_groups_cache_mime,omitempty"`

	StrGroupsNoCacheMime *string `json:"strGroupsNoCacheMime,omitempty" tf:"str_groups_no_cache_mime,omitempty"`

	StrGroupsNoCacheURI *string `json:"strGroupsNoCacheUri,omitempty" tf:"str_groups_no_cache_uri,omitempty"`

	StrNoCacheMime *string `json:"strNoCacheMime,omitempty" tf:"str_no_cache_mime,omitempty"`

	StrNoCacheURI *string `json:"strNoCacheUri,omitempty" tf:"str_no_cache_uri,omitempty"`
}

type L7LimitsParameters struct {

	// +kubebuilder:validation:Optional
	HTTPPoliciesPerVs *string `json:"httpPoliciesPerVs,omitempty" tf:"http_policies_per_vs,omitempty"`

	// +kubebuilder:validation:Optional
	NumCompressionFilters *string `json:"numCompressionFilters,omitempty" tf:"num_compression_filters,omitempty"`

	// +kubebuilder:validation:Optional
	NumCustomStr *string `json:"numCustomStr,omitempty" tf:"num_custom_str,omitempty"`

	// +kubebuilder:validation:Optional
	NumMatchesPerRule *string `json:"numMatchesPerRule,omitempty" tf:"num_matches_per_rule,omitempty"`

	// +kubebuilder:validation:Optional
	NumRulesPerEvhHost *string `json:"numRulesPerEvhHost,omitempty" tf:"num_rules_per_evh_host,omitempty"`

	// +kubebuilder:validation:Optional
	NumRulesPerHTTPPolicy *string `json:"numRulesPerHttpPolicy,omitempty" tf:"num_rules_per_http_policy,omitempty"`

	// +kubebuilder:validation:Optional
	NumStrgroupsPerMatch *string `json:"numStrgroupsPerMatch,omitempty" tf:"num_strgroups_per_match,omitempty"`

	// +kubebuilder:validation:Optional
	StrCacheMime *string `json:"strCacheMime,omitempty" tf:"str_cache_mime,omitempty"`

	// +kubebuilder:validation:Optional
	StrGroupsCacheMime *string `json:"strGroupsCacheMime,omitempty" tf:"str_groups_cache_mime,omitempty"`

	// +kubebuilder:validation:Optional
	StrGroupsNoCacheMime *string `json:"strGroupsNoCacheMime,omitempty" tf:"str_groups_no_cache_mime,omitempty"`

	// +kubebuilder:validation:Optional
	StrGroupsNoCacheURI *string `json:"strGroupsNoCacheUri,omitempty" tf:"str_groups_no_cache_uri,omitempty"`

	// +kubebuilder:validation:Optional
	StrNoCacheMime *string `json:"strNoCacheMime,omitempty" tf:"str_no_cache_mime,omitempty"`

	// +kubebuilder:validation:Optional
	StrNoCacheURI *string `json:"strNoCacheUri,omitempty" tf:"str_no_cache_uri,omitempty"`
}

type ServiceengineCloudLimitsInitParameters struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	VrfsPerServiceengine *string `json:"vrfsPerServiceengine,omitempty" tf:"vrfs_per_serviceengine,omitempty"`
}

type ServiceengineCloudLimitsObservation struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	VrfsPerServiceengine *string `json:"vrfsPerServiceengine,omitempty" tf:"vrfs_per_serviceengine,omitempty"`
}

type ServiceengineCloudLimitsParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VrfsPerServiceengine *string `json:"vrfsPerServiceengine,omitempty" tf:"vrfs_per_serviceengine,omitempty"`
}

type ServiceengineLimitsInitParameters struct {
	AllVirtualservicesPerServiceengine *string `json:"allVirtualservicesPerServiceengine,omitempty" tf:"all_virtualservices_per_serviceengine,omitempty"`

	EwVirtualservicesPerServiceengine *string `json:"ewVirtualservicesPerServiceengine,omitempty" tf:"ew_virtualservices_per_serviceengine,omitempty"`

	NsVirtualservicesPerServiceengine *string `json:"nsVirtualservicesPerServiceengine,omitempty" tf:"ns_virtualservices_per_serviceengine,omitempty"`

	NumLogicalIntfPerSe *string `json:"numLogicalIntfPerSe,omitempty" tf:"num_logical_intf_per_se,omitempty"`

	NumPhyIntfPerSe *string `json:"numPhyIntfPerSe,omitempty" tf:"num_phy_intf_per_se,omitempty"`

	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	NumVlanIntfPerPhyIntf *string `json:"numVlanIntfPerPhyIntf,omitempty" tf:"num_vlan_intf_per_phy_intf,omitempty"`

	NumVlanIntfPerSe *string `json:"numVlanIntfPerSe,omitempty" tf:"num_vlan_intf_per_se,omitempty"`

	ServiceengineCloudLimits []ServiceengineCloudLimitsInitParameters `json:"serviceengineCloudLimits,omitempty" tf:"serviceengine_cloud_limits,omitempty"`
}

type ServiceengineLimitsObservation struct {
	AllVirtualservicesPerServiceengine *string `json:"allVirtualservicesPerServiceengine,omitempty" tf:"all_virtualservices_per_serviceengine,omitempty"`

	EwVirtualservicesPerServiceengine *string `json:"ewVirtualservicesPerServiceengine,omitempty" tf:"ew_virtualservices_per_serviceengine,omitempty"`

	NsVirtualservicesPerServiceengine *string `json:"nsVirtualservicesPerServiceengine,omitempty" tf:"ns_virtualservices_per_serviceengine,omitempty"`

	NumLogicalIntfPerSe *string `json:"numLogicalIntfPerSe,omitempty" tf:"num_logical_intf_per_se,omitempty"`

	NumPhyIntfPerSe *string `json:"numPhyIntfPerSe,omitempty" tf:"num_phy_intf_per_se,omitempty"`

	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	NumVlanIntfPerPhyIntf *string `json:"numVlanIntfPerPhyIntf,omitempty" tf:"num_vlan_intf_per_phy_intf,omitempty"`

	NumVlanIntfPerSe *string `json:"numVlanIntfPerSe,omitempty" tf:"num_vlan_intf_per_se,omitempty"`

	ServiceengineCloudLimits []ServiceengineCloudLimitsObservation `json:"serviceengineCloudLimits,omitempty" tf:"serviceengine_cloud_limits,omitempty"`
}

type ServiceengineLimitsParameters struct {

	// +kubebuilder:validation:Optional
	AllVirtualservicesPerServiceengine *string `json:"allVirtualservicesPerServiceengine,omitempty" tf:"all_virtualservices_per_serviceengine,omitempty"`

	// +kubebuilder:validation:Optional
	EwVirtualservicesPerServiceengine *string `json:"ewVirtualservicesPerServiceengine,omitempty" tf:"ew_virtualservices_per_serviceengine,omitempty"`

	// +kubebuilder:validation:Optional
	NsVirtualservicesPerServiceengine *string `json:"nsVirtualservicesPerServiceengine,omitempty" tf:"ns_virtualservices_per_serviceengine,omitempty"`

	// +kubebuilder:validation:Optional
	NumLogicalIntfPerSe *string `json:"numLogicalIntfPerSe,omitempty" tf:"num_logical_intf_per_se,omitempty"`

	// +kubebuilder:validation:Optional
	NumPhyIntfPerSe *string `json:"numPhyIntfPerSe,omitempty" tf:"num_phy_intf_per_se,omitempty"`

	// +kubebuilder:validation:Optional
	NumVirtualservicesRtMetrics *string `json:"numVirtualservicesRtMetrics,omitempty" tf:"num_virtualservices_rt_metrics,omitempty"`

	// +kubebuilder:validation:Optional
	NumVlanIntfPerPhyIntf *string `json:"numVlanIntfPerPhyIntf,omitempty" tf:"num_vlan_intf_per_phy_intf,omitempty"`

	// +kubebuilder:validation:Optional
	NumVlanIntfPerSe *string `json:"numVlanIntfPerSe,omitempty" tf:"num_vlan_intf_per_se,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceengineCloudLimits []ServiceengineCloudLimitsParameters `json:"serviceengineCloudLimits,omitempty" tf:"serviceengine_cloud_limits,omitempty"`
}

type SystemLimitsConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SystemLimitsConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SystemLimitsConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SystemLimitsInitParameters struct {
	ConfigpbAttributes []SystemLimitsConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ControllerLimits []ControllerLimitsInitParameters `json:"controllerLimits,omitempty" tf:"controller_limits,omitempty"`

	ControllerSizes []ControllerSizesInitParameters `json:"controllerSizes,omitempty" tf:"controller_sizes,omitempty"`

	ServiceengineLimits []ServiceengineLimitsInitParameters `json:"serviceengineLimits,omitempty" tf:"serviceengine_limits,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SystemLimitsObservation struct {
	ConfigpbAttributes []SystemLimitsConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ControllerLimits []ControllerLimitsObservation `json:"controllerLimits,omitempty" tf:"controller_limits,omitempty"`

	ControllerSizes []ControllerSizesObservation `json:"controllerSizes,omitempty" tf:"controller_sizes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ServiceengineLimits []ServiceengineLimitsObservation `json:"serviceengineLimits,omitempty" tf:"serviceengine_limits,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SystemLimitsParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []SystemLimitsConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	ControllerLimits []ControllerLimitsParameters `json:"controllerLimits,omitempty" tf:"controller_limits,omitempty"`

	// +kubebuilder:validation:Optional
	ControllerSizes []ControllerSizesParameters `json:"controllerSizes,omitempty" tf:"controller_sizes,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceengineLimits []ServiceengineLimitsParameters `json:"serviceengineLimits,omitempty" tf:"serviceengine_limits,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type WafLimitsInitParameters struct {
	NumAllowedContentTypes *string `json:"numAllowedContentTypes,omitempty" tf:"num_allowed_content_types,omitempty"`

	NumAllowedRequestContentTypeCharsets *string `json:"numAllowedRequestContentTypeCharsets,omitempty" tf:"num_allowed_request_content_type_charsets,omitempty"`

	NumAllowlistPolicyRules *string `json:"numAllowlistPolicyRules,omitempty" tf:"num_allowlist_policy_rules,omitempty"`

	NumApplications *string `json:"numApplications,omitempty" tf:"num_applications,omitempty"`

	NumDataFiles *string `json:"numDataFiles,omitempty" tf:"num_data_files,omitempty"`

	NumPrePostCrsGroups *string `json:"numPrePostCrsGroups,omitempty" tf:"num_pre_post_crs_groups,omitempty"`

	NumPsmGroups *string `json:"numPsmGroups,omitempty" tf:"num_psm_groups,omitempty"`

	NumPsmMatchElements *string `json:"numPsmMatchElements,omitempty" tf:"num_psm_match_elements,omitempty"`

	NumPsmMatchRulesPerLoc *string `json:"numPsmMatchRulesPerLoc,omitempty" tf:"num_psm_match_rules_per_loc,omitempty"`

	NumPsmTotalLocations *string `json:"numPsmTotalLocations,omitempty" tf:"num_psm_total_locations,omitempty"`

	NumRestrictedExtensions *string `json:"numRestrictedExtensions,omitempty" tf:"num_restricted_extensions,omitempty"`

	NumRestrictedHeaders *string `json:"numRestrictedHeaders,omitempty" tf:"num_restricted_headers,omitempty"`

	NumRuleTags *string `json:"numRuleTags,omitempty" tf:"num_rule_tags,omitempty"`

	NumRulesPerRulegroup *string `json:"numRulesPerRulegroup,omitempty" tf:"num_rules_per_rulegroup,omitempty"`

	NumStaticExtensions *string `json:"numStaticExtensions,omitempty" tf:"num_static_extensions,omitempty"`
}

type WafLimitsObservation struct {
	NumAllowedContentTypes *string `json:"numAllowedContentTypes,omitempty" tf:"num_allowed_content_types,omitempty"`

	NumAllowedRequestContentTypeCharsets *string `json:"numAllowedRequestContentTypeCharsets,omitempty" tf:"num_allowed_request_content_type_charsets,omitempty"`

	NumAllowlistPolicyRules *string `json:"numAllowlistPolicyRules,omitempty" tf:"num_allowlist_policy_rules,omitempty"`

	NumApplications *string `json:"numApplications,omitempty" tf:"num_applications,omitempty"`

	NumDataFiles *string `json:"numDataFiles,omitempty" tf:"num_data_files,omitempty"`

	NumPrePostCrsGroups *string `json:"numPrePostCrsGroups,omitempty" tf:"num_pre_post_crs_groups,omitempty"`

	NumPsmGroups *string `json:"numPsmGroups,omitempty" tf:"num_psm_groups,omitempty"`

	NumPsmMatchElements *string `json:"numPsmMatchElements,omitempty" tf:"num_psm_match_elements,omitempty"`

	NumPsmMatchRulesPerLoc *string `json:"numPsmMatchRulesPerLoc,omitempty" tf:"num_psm_match_rules_per_loc,omitempty"`

	NumPsmTotalLocations *string `json:"numPsmTotalLocations,omitempty" tf:"num_psm_total_locations,omitempty"`

	NumRestrictedExtensions *string `json:"numRestrictedExtensions,omitempty" tf:"num_restricted_extensions,omitempty"`

	NumRestrictedHeaders *string `json:"numRestrictedHeaders,omitempty" tf:"num_restricted_headers,omitempty"`

	NumRuleTags *string `json:"numRuleTags,omitempty" tf:"num_rule_tags,omitempty"`

	NumRulesPerRulegroup *string `json:"numRulesPerRulegroup,omitempty" tf:"num_rules_per_rulegroup,omitempty"`

	NumStaticExtensions *string `json:"numStaticExtensions,omitempty" tf:"num_static_extensions,omitempty"`
}

type WafLimitsParameters struct {

	// +kubebuilder:validation:Optional
	NumAllowedContentTypes *string `json:"numAllowedContentTypes,omitempty" tf:"num_allowed_content_types,omitempty"`

	// +kubebuilder:validation:Optional
	NumAllowedRequestContentTypeCharsets *string `json:"numAllowedRequestContentTypeCharsets,omitempty" tf:"num_allowed_request_content_type_charsets,omitempty"`

	// +kubebuilder:validation:Optional
	NumAllowlistPolicyRules *string `json:"numAllowlistPolicyRules,omitempty" tf:"num_allowlist_policy_rules,omitempty"`

	// +kubebuilder:validation:Optional
	NumApplications *string `json:"numApplications,omitempty" tf:"num_applications,omitempty"`

	// +kubebuilder:validation:Optional
	NumDataFiles *string `json:"numDataFiles,omitempty" tf:"num_data_files,omitempty"`

	// +kubebuilder:validation:Optional
	NumPrePostCrsGroups *string `json:"numPrePostCrsGroups,omitempty" tf:"num_pre_post_crs_groups,omitempty"`

	// +kubebuilder:validation:Optional
	NumPsmGroups *string `json:"numPsmGroups,omitempty" tf:"num_psm_groups,omitempty"`

	// +kubebuilder:validation:Optional
	NumPsmMatchElements *string `json:"numPsmMatchElements,omitempty" tf:"num_psm_match_elements,omitempty"`

	// +kubebuilder:validation:Optional
	NumPsmMatchRulesPerLoc *string `json:"numPsmMatchRulesPerLoc,omitempty" tf:"num_psm_match_rules_per_loc,omitempty"`

	// +kubebuilder:validation:Optional
	NumPsmTotalLocations *string `json:"numPsmTotalLocations,omitempty" tf:"num_psm_total_locations,omitempty"`

	// +kubebuilder:validation:Optional
	NumRestrictedExtensions *string `json:"numRestrictedExtensions,omitempty" tf:"num_restricted_extensions,omitempty"`

	// +kubebuilder:validation:Optional
	NumRestrictedHeaders *string `json:"numRestrictedHeaders,omitempty" tf:"num_restricted_headers,omitempty"`

	// +kubebuilder:validation:Optional
	NumRuleTags *string `json:"numRuleTags,omitempty" tf:"num_rule_tags,omitempty"`

	// +kubebuilder:validation:Optional
	NumRulesPerRulegroup *string `json:"numRulesPerRulegroup,omitempty" tf:"num_rules_per_rulegroup,omitempty"`

	// +kubebuilder:validation:Optional
	NumStaticExtensions *string `json:"numStaticExtensions,omitempty" tf:"num_static_extensions,omitempty"`
}

// SystemLimitsSpec defines the desired state of SystemLimits
type SystemLimitsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SystemLimitsParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SystemLimitsInitParameters `json:"initProvider,omitempty"`
}

// SystemLimitsStatus defines the observed state of SystemLimits.
type SystemLimitsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SystemLimitsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SystemLimits is the Schema for the SystemLimitss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type SystemLimits struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SystemLimitsSpec   `json:"spec"`
	Status            SystemLimitsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SystemLimitsList contains a list of SystemLimitss
type SystemLimitsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SystemLimits `json:"items"`
}

// Repository type metadata.
var (
	SystemLimits_Kind             = "SystemLimits"
	SystemLimits_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SystemLimits_Kind}.String()
	SystemLimits_KindAPIVersion   = SystemLimits_Kind + "." + CRDGroupVersion.String()
	SystemLimits_GroupVersionKind = CRDGroupVersion.WithKind(SystemLimits_Kind)
)

func init() {
	SchemeBuilder.Register(&SystemLimits{}, &SystemLimitsList{})
}