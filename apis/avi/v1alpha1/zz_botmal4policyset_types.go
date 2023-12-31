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

type BotMaL4PolicySetConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type BotMaL4PolicySetConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type BotMaL4PolicySetConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type BotMaL4PolicySetInitParameters struct {
	ConfigpbAttributes []BotMaL4PolicySetConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	IsInternalPolicy *string `json:"isInternalPolicy,omitempty" tf:"is_internal_policy,omitempty"`

	L4ConnectionPolicy []L4ConnectionPolicyInitParameters `json:"l4ConnectionPolicy,omitempty" tf:"l4_connection_policy,omitempty"`

	Markers []BotMaL4PolicySetMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type BotMaL4PolicySetMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BotMaL4PolicySetMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BotMaL4PolicySetMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BotMaL4PolicySetObservation struct {
	ConfigpbAttributes []BotMaL4PolicySetConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsInternalPolicy *string `json:"isInternalPolicy,omitempty" tf:"is_internal_policy,omitempty"`

	L4ConnectionPolicy []L4ConnectionPolicyObservation `json:"l4ConnectionPolicy,omitempty" tf:"l4_connection_policy,omitempty"`

	Markers []BotMaL4PolicySetMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type BotMaL4PolicySetParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []BotMaL4PolicySetConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IsInternalPolicy *string `json:"isInternalPolicy,omitempty" tf:"is_internal_policy,omitempty"`

	// +kubebuilder:validation:Optional
	L4ConnectionPolicy []L4ConnectionPolicyParameters `json:"l4ConnectionPolicy,omitempty" tf:"l4_connection_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []BotMaL4PolicySetMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type L4ConnectionPolicyInitParameters struct {
	Rules []L4ConnectionPolicyRulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type L4ConnectionPolicyObservation struct {
	Rules []L4ConnectionPolicyRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type L4ConnectionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Rules []L4ConnectionPolicyRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type L4ConnectionPolicyRulesActionInitParameters struct {
	SelectPool []SelectPoolInitParameters `json:"selectPool,omitempty" tf:"select_pool,omitempty"`
}

type L4ConnectionPolicyRulesActionObservation struct {
	SelectPool []SelectPoolObservation `json:"selectPool,omitempty" tf:"select_pool,omitempty"`
}

type L4ConnectionPolicyRulesActionParameters struct {

	// +kubebuilder:validation:Optional
	SelectPool []SelectPoolParameters `json:"selectPool,omitempty" tf:"select_pool,omitempty"`
}

type L4ConnectionPolicyRulesInitParameters struct {
	Action []L4ConnectionPolicyRulesActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Match []L4ConnectionPolicyRulesMatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPAddrsInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPAddrsObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPAddrsParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPInitParameters struct {
	Addrs []L4ConnectionPolicyRulesMatchClientIPAddrsInitParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []L4ConnectionPolicyRulesMatchClientIPPrefixesInitParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []L4ConnectionPolicyRulesMatchClientIPRangesInitParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPObservation struct {
	Addrs []L4ConnectionPolicyRulesMatchClientIPAddrsObservation `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []L4ConnectionPolicyRulesMatchClientIPPrefixesObservation `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []L4ConnectionPolicyRulesMatchClientIPRangesObservation `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPParameters struct {

	// +kubebuilder:validation:Optional
	Addrs []L4ConnectionPolicyRulesMatchClientIPAddrsParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Prefixes []L4ConnectionPolicyRulesMatchClientIPPrefixesParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	Ranges []L4ConnectionPolicyRulesMatchClientIPRangesParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesInitParameters struct {
	IPAddr []L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrInitParameters `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesObservation struct {
	IPAddr []L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrObservation `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPPrefixesParameters struct {

	// +kubebuilder:validation:Optional
	IPAddr []L4ConnectionPolicyRulesMatchClientIPPrefixesIPAddrParameters `json:"ipAddr" tf:"ip_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Mask *string `json:"mask" tf:"mask,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesBeginInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesBeginObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesBeginParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesEndInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesEndObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesEndParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesInitParameters struct {
	Begin []L4ConnectionPolicyRulesMatchClientIPRangesBeginInitParameters `json:"begin,omitempty" tf:"begin,omitempty"`

	End []L4ConnectionPolicyRulesMatchClientIPRangesEndInitParameters `json:"end,omitempty" tf:"end,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesObservation struct {
	Begin []L4ConnectionPolicyRulesMatchClientIPRangesBeginObservation `json:"begin,omitempty" tf:"begin,omitempty"`

	End []L4ConnectionPolicyRulesMatchClientIPRangesEndObservation `json:"end,omitempty" tf:"end,omitempty"`
}

type L4ConnectionPolicyRulesMatchClientIPRangesParameters struct {

	// +kubebuilder:validation:Optional
	Begin []L4ConnectionPolicyRulesMatchClientIPRangesBeginParameters `json:"begin" tf:"begin,omitempty"`

	// +kubebuilder:validation:Optional
	End []L4ConnectionPolicyRulesMatchClientIPRangesEndParameters `json:"end" tf:"end,omitempty"`
}

type L4ConnectionPolicyRulesMatchInitParameters struct {
	ClientIP []L4ConnectionPolicyRulesMatchClientIPInitParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	Port []PortInitParameters `json:"port,omitempty" tf:"port,omitempty"`

	Protocol []L4ConnectionPolicyRulesMatchProtocolInitParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesMatchObservation struct {
	ClientIP []L4ConnectionPolicyRulesMatchClientIPObservation `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	Port []PortObservation `json:"port,omitempty" tf:"port,omitempty"`

	Protocol []L4ConnectionPolicyRulesMatchProtocolObservation `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesMatchParameters struct {

	// +kubebuilder:validation:Optional
	ClientIP []L4ConnectionPolicyRulesMatchClientIPParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Port []PortParameters `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol []L4ConnectionPolicyRulesMatchProtocolParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesMatchProtocolInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesMatchProtocolObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesMatchProtocolParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type L4ConnectionPolicyRulesObservation struct {
	Action []L4ConnectionPolicyRulesActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Match []L4ConnectionPolicyRulesMatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type L4ConnectionPolicyRulesParameters struct {

	// +kubebuilder:validation:Optional
	Action []L4ConnectionPolicyRulesActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// +kubebuilder:validation:Optional
	Match []L4ConnectionPolicyRulesMatchParameters `json:"match,omitempty" tf:"match,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type PortInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	PortRanges []PortRangesInitParameters `json:"portRanges,omitempty" tf:"port_ranges,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type PortObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	PortRanges []PortRangesObservation `json:"portRanges,omitempty" tf:"port_ranges,omitempty"`

	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type PortParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	PortRanges []PortRangesParameters `json:"portRanges,omitempty" tf:"port_ranges,omitempty"`

	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type PortRangesInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type PortRangesObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type PortRangesParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type SelectPoolInitParameters struct {
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	PoolGroupRef *string `json:"poolGroupRef,omitempty" tf:"pool_group_ref,omitempty"`

	PoolRef *string `json:"poolRef,omitempty" tf:"pool_ref,omitempty"`
}

type SelectPoolObservation struct {
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	PoolGroupRef *string `json:"poolGroupRef,omitempty" tf:"pool_group_ref,omitempty"`

	PoolRef *string `json:"poolRef,omitempty" tf:"pool_ref,omitempty"`
}

type SelectPoolParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	PoolGroupRef *string `json:"poolGroupRef,omitempty" tf:"pool_group_ref,omitempty"`

	// +kubebuilder:validation:Optional
	PoolRef *string `json:"poolRef,omitempty" tf:"pool_ref,omitempty"`
}

// BotMaL4PolicySetSpec defines the desired state of BotMaL4PolicySet
type BotMaL4PolicySetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotMaL4PolicySetParameters `json:"forProvider"`
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
	InitProvider BotMaL4PolicySetInitParameters `json:"initProvider,omitempty"`
}

// BotMaL4PolicySetStatus defines the observed state of BotMaL4PolicySet.
type BotMaL4PolicySetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotMaL4PolicySetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotMaL4PolicySet is the Schema for the BotMaL4PolicySets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type BotMaL4PolicySet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   BotMaL4PolicySetSpec   `json:"spec"`
	Status BotMaL4PolicySetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotMaL4PolicySetList contains a list of BotMaL4PolicySets
type BotMaL4PolicySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotMaL4PolicySet `json:"items"`
}

// Repository type metadata.
var (
	BotMaL4PolicySet_Kind             = "BotMaL4PolicySet"
	BotMaL4PolicySet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotMaL4PolicySet_Kind}.String()
	BotMaL4PolicySet_KindAPIVersion   = BotMaL4PolicySet_Kind + "." + CRDGroupVersion.String()
	BotMaL4PolicySet_GroupVersionKind = CRDGroupVersion.WithKind(BotMaL4PolicySet_Kind)
)

func init() {
	SchemeBuilder.Register(&BotMaL4PolicySet{}, &BotMaL4PolicySetList{})
}
