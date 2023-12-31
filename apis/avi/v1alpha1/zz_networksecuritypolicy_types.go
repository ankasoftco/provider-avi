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

type ClientPortInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	Ranges []ClientPortRangesInitParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type ClientPortObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	Ranges []ClientPortRangesObservation `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type ClientPortParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	Ranges []ClientPortRangesParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type ClientPortRangesInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type ClientPortRangesObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type ClientPortRangesParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type MicroserviceInitParameters struct {
	GroupRef *string `json:"groupRef,omitempty" tf:"group_ref,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`
}

type MicroserviceObservation struct {
	GroupRef *string `json:"groupRef,omitempty" tf:"group_ref,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`
}

type MicroserviceParameters struct {

	// +kubebuilder:validation:Optional
	GroupRef *string `json:"groupRef" tf:"group_ref,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`
}

type NetworkSecurityPolicyConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkSecurityPolicyConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkSecurityPolicyConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkSecurityPolicyInitParameters struct {
	CloudConfigCksum *string `json:"cloudConfigCksum,omitempty" tf:"cloud_config_cksum,omitempty"`

	ConfigpbAttributes []NetworkSecurityPolicyConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	GeoDBRef *string `json:"geoDbRef,omitempty" tf:"geo_db_ref,omitempty"`

	IPReputationDBRef *string `json:"ipReputationDbRef,omitempty" tf:"ip_reputation_db_ref,omitempty"`

	Internal *string `json:"internal,omitempty" tf:"internal,omitempty"`

	Markers []NetworkSecurityPolicyMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Rules []NetworkSecurityPolicyRulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkSecurityPolicyMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkSecurityPolicyMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkSecurityPolicyMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkSecurityPolicyObservation struct {
	CloudConfigCksum *string `json:"cloudConfigCksum,omitempty" tf:"cloud_config_cksum,omitempty"`

	ConfigpbAttributes []NetworkSecurityPolicyConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	GeoDBRef *string `json:"geoDbRef,omitempty" tf:"geo_db_ref,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPReputationDBRef *string `json:"ipReputationDbRef,omitempty" tf:"ip_reputation_db_ref,omitempty"`

	Internal *string `json:"internal,omitempty" tf:"internal,omitempty"`

	Markers []NetworkSecurityPolicyMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Rules []NetworkSecurityPolicyRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkSecurityPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CloudConfigCksum *string `json:"cloudConfigCksum,omitempty" tf:"cloud_config_cksum,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []NetworkSecurityPolicyConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GeoDBRef *string `json:"geoDbRef,omitempty" tf:"geo_db_ref,omitempty"`

	// +kubebuilder:validation:Optional
	IPReputationDBRef *string `json:"ipReputationDbRef,omitempty" tf:"ip_reputation_db_ref,omitempty"`

	// +kubebuilder:validation:Optional
	Internal *string `json:"internal,omitempty" tf:"internal,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []NetworkSecurityPolicyMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Rules []NetworkSecurityPolicyRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkSecurityPolicyRulesInitParameters struct {
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	Age *string `json:"age,omitempty" tf:"age,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Log *string `json:"log,omitempty" tf:"log,omitempty"`

	Match []NetworkSecurityPolicyRulesMatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RlParam []RlParamInitParameters `json:"rlParam,omitempty" tf:"rl_param,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPAddrsInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPAddrsObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPAddrsParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPInitParameters struct {
	Addrs []NetworkSecurityPolicyRulesMatchClientIPAddrsInitParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []NetworkSecurityPolicyRulesMatchClientIPPrefixesInitParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []NetworkSecurityPolicyRulesMatchClientIPRangesInitParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPObservation struct {
	Addrs []NetworkSecurityPolicyRulesMatchClientIPAddrsObservation `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []NetworkSecurityPolicyRulesMatchClientIPPrefixesObservation `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []NetworkSecurityPolicyRulesMatchClientIPRangesObservation `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPParameters struct {

	// +kubebuilder:validation:Optional
	Addrs []NetworkSecurityPolicyRulesMatchClientIPAddrsParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Prefixes []NetworkSecurityPolicyRulesMatchClientIPPrefixesParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	Ranges []NetworkSecurityPolicyRulesMatchClientIPRangesParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesInitParameters struct {
	IPAddr []NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrInitParameters `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesObservation struct {
	IPAddr []NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrObservation `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPPrefixesParameters struct {

	// +kubebuilder:validation:Optional
	IPAddr []NetworkSecurityPolicyRulesMatchClientIPPrefixesIPAddrParameters `json:"ipAddr" tf:"ip_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Mask *string `json:"mask" tf:"mask,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesBeginInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesBeginObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesBeginParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesEndInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesEndObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesEndParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesInitParameters struct {
	Begin []NetworkSecurityPolicyRulesMatchClientIPRangesBeginInitParameters `json:"begin,omitempty" tf:"begin,omitempty"`

	End []NetworkSecurityPolicyRulesMatchClientIPRangesEndInitParameters `json:"end,omitempty" tf:"end,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesObservation struct {
	Begin []NetworkSecurityPolicyRulesMatchClientIPRangesBeginObservation `json:"begin,omitempty" tf:"begin,omitempty"`

	End []NetworkSecurityPolicyRulesMatchClientIPRangesEndObservation `json:"end,omitempty" tf:"end,omitempty"`
}

type NetworkSecurityPolicyRulesMatchClientIPRangesParameters struct {

	// +kubebuilder:validation:Optional
	Begin []NetworkSecurityPolicyRulesMatchClientIPRangesBeginParameters `json:"begin" tf:"begin,omitempty"`

	// +kubebuilder:validation:Optional
	End []NetworkSecurityPolicyRulesMatchClientIPRangesEndParameters `json:"end" tf:"end,omitempty"`
}

type NetworkSecurityPolicyRulesMatchGeoMatchesInitParameters struct {
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	MatchOperation *string `json:"matchOperation,omitempty" tf:"match_operation,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkSecurityPolicyRulesMatchGeoMatchesObservation struct {
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	MatchOperation *string `json:"matchOperation,omitempty" tf:"match_operation,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkSecurityPolicyRulesMatchGeoMatchesParameters struct {

	// +kubebuilder:validation:Optional
	Attribute *string `json:"attribute" tf:"attribute,omitempty"`

	// +kubebuilder:validation:Optional
	MatchOperation *string `json:"matchOperation" tf:"match_operation,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type NetworkSecurityPolicyRulesMatchIPReputationTypeInitParameters struct {
	MatchOperation *string `json:"matchOperation,omitempty" tf:"match_operation,omitempty"`

	ReputationTypes []*string `json:"reputationTypes,omitempty" tf:"reputation_types,omitempty"`
}

type NetworkSecurityPolicyRulesMatchIPReputationTypeObservation struct {
	MatchOperation *string `json:"matchOperation,omitempty" tf:"match_operation,omitempty"`

	ReputationTypes []*string `json:"reputationTypes,omitempty" tf:"reputation_types,omitempty"`
}

type NetworkSecurityPolicyRulesMatchIPReputationTypeParameters struct {

	// +kubebuilder:validation:Optional
	MatchOperation *string `json:"matchOperation" tf:"match_operation,omitempty"`

	// +kubebuilder:validation:Optional
	ReputationTypes []*string `json:"reputationTypes" tf:"reputation_types,omitempty"`
}

type NetworkSecurityPolicyRulesMatchInitParameters struct {
	ClientIP []NetworkSecurityPolicyRulesMatchClientIPInitParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	ClientPort []ClientPortInitParameters `json:"clientPort,omitempty" tf:"client_port,omitempty"`

	GeoMatches []NetworkSecurityPolicyRulesMatchGeoMatchesInitParameters `json:"geoMatches,omitempty" tf:"geo_matches,omitempty"`

	IPReputationType []NetworkSecurityPolicyRulesMatchIPReputationTypeInitParameters `json:"ipReputationType,omitempty" tf:"ip_reputation_type,omitempty"`

	Microservice []MicroserviceInitParameters `json:"microservice,omitempty" tf:"microservice,omitempty"`

	VsPort []NetworkSecurityPolicyRulesMatchVsPortInitParameters `json:"vsPort,omitempty" tf:"vs_port,omitempty"`
}

type NetworkSecurityPolicyRulesMatchObservation struct {
	ClientIP []NetworkSecurityPolicyRulesMatchClientIPObservation `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	ClientPort []ClientPortObservation `json:"clientPort,omitempty" tf:"client_port,omitempty"`

	GeoMatches []NetworkSecurityPolicyRulesMatchGeoMatchesObservation `json:"geoMatches,omitempty" tf:"geo_matches,omitempty"`

	IPReputationType []NetworkSecurityPolicyRulesMatchIPReputationTypeObservation `json:"ipReputationType,omitempty" tf:"ip_reputation_type,omitempty"`

	Microservice []MicroserviceObservation `json:"microservice,omitempty" tf:"microservice,omitempty"`

	VsPort []NetworkSecurityPolicyRulesMatchVsPortObservation `json:"vsPort,omitempty" tf:"vs_port,omitempty"`
}

type NetworkSecurityPolicyRulesMatchParameters struct {

	// +kubebuilder:validation:Optional
	ClientIP []NetworkSecurityPolicyRulesMatchClientIPParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	ClientPort []ClientPortParameters `json:"clientPort,omitempty" tf:"client_port,omitempty"`

	// +kubebuilder:validation:Optional
	GeoMatches []NetworkSecurityPolicyRulesMatchGeoMatchesParameters `json:"geoMatches,omitempty" tf:"geo_matches,omitempty"`

	// +kubebuilder:validation:Optional
	IPReputationType []NetworkSecurityPolicyRulesMatchIPReputationTypeParameters `json:"ipReputationType,omitempty" tf:"ip_reputation_type,omitempty"`

	// +kubebuilder:validation:Optional
	Microservice []MicroserviceParameters `json:"microservice,omitempty" tf:"microservice,omitempty"`

	// +kubebuilder:validation:Optional
	VsPort []NetworkSecurityPolicyRulesMatchVsPortParameters `json:"vsPort,omitempty" tf:"vs_port,omitempty"`
}

type NetworkSecurityPolicyRulesMatchVsPortInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type NetworkSecurityPolicyRulesMatchVsPortObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type NetworkSecurityPolicyRulesMatchVsPortParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports" tf:"ports,omitempty"`
}

type NetworkSecurityPolicyRulesObservation struct {
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	Age *string `json:"age,omitempty" tf:"age,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Log *string `json:"log,omitempty" tf:"log,omitempty"`

	Match []NetworkSecurityPolicyRulesMatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RlParam []RlParamObservation `json:"rlParam,omitempty" tf:"rl_param,omitempty"`
}

type NetworkSecurityPolicyRulesParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Age *string `json:"age,omitempty" tf:"age,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// +kubebuilder:validation:Optional
	Enable *string `json:"enable" tf:"enable,omitempty"`

	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// +kubebuilder:validation:Optional
	Log *string `json:"log,omitempty" tf:"log,omitempty"`

	// +kubebuilder:validation:Optional
	Match []NetworkSecurityPolicyRulesMatchParameters `json:"match" tf:"match,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RlParam []RlParamParameters `json:"rlParam,omitempty" tf:"rl_param,omitempty"`
}

type RlParamInitParameters struct {
	BurstSize *string `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	MaxRate *string `json:"maxRate,omitempty" tf:"max_rate,omitempty"`
}

type RlParamObservation struct {
	BurstSize *string `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	MaxRate *string `json:"maxRate,omitempty" tf:"max_rate,omitempty"`
}

type RlParamParameters struct {

	// +kubebuilder:validation:Optional
	BurstSize *string `json:"burstSize" tf:"burst_size,omitempty"`

	// +kubebuilder:validation:Optional
	MaxRate *string `json:"maxRate" tf:"max_rate,omitempty"`
}

// NetworkSecurityPolicySpec defines the desired state of NetworkSecurityPolicy
type NetworkSecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkSecurityPolicyParameters `json:"forProvider"`
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
	InitProvider NetworkSecurityPolicyInitParameters `json:"initProvider,omitempty"`
}

// NetworkSecurityPolicyStatus defines the observed state of NetworkSecurityPolicy.
type NetworkSecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkSecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityPolicy is the Schema for the NetworkSecurityPolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type NetworkSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityPolicySpec   `json:"spec"`
	Status            NetworkSecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityPolicyList contains a list of NetworkSecurityPolicys
type NetworkSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	NetworkSecurityPolicy_Kind             = "NetworkSecurityPolicy"
	NetworkSecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkSecurityPolicy_Kind}.String()
	NetworkSecurityPolicy_KindAPIVersion   = NetworkSecurityPolicy_Kind + "." + CRDGroupVersion.String()
	NetworkSecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(NetworkSecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkSecurityPolicy{}, &NetworkSecurityPolicyList{})
}
