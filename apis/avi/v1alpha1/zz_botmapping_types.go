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

type BotMappingInitParameters struct {
	MappingRules []BotMappingMappingRulesInitParameters `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type BotMappingMappingRulesInitParameters struct {
	Classification []ClassificationInitParameters `json:"classification,omitempty" tf:"classification,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Match []MappingRulesMatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BotMappingMappingRulesObservation struct {
	Classification []ClassificationObservation `json:"classification,omitempty" tf:"classification,omitempty"`

	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	Match []MappingRulesMatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BotMappingMappingRulesParameters struct {

	// +kubebuilder:validation:Optional
	Classification []ClassificationParameters `json:"classification" tf:"classification,omitempty"`

	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// +kubebuilder:validation:Optional
	Match []MappingRulesMatchParameters `json:"match" tf:"match,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type BotMappingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MappingRules []BotMappingMappingRulesObservation `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type BotMappingParameters struct {

	// +kubebuilder:validation:Optional
	MappingRules []BotMappingMappingRulesParameters `json:"mappingRules,omitempty" tf:"mapping_rules,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ClassMatcherInitParameters struct {
	ClientClasses []*string `json:"clientClasses,omitempty" tf:"client_classes,omitempty"`

	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

type ClassMatcherObservation struct {
	ClientClasses []*string `json:"clientClasses,omitempty" tf:"client_classes,omitempty"`

	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

type ClassMatcherParameters struct {

	// +kubebuilder:validation:Optional
	ClientClasses []*string `json:"clientClasses" tf:"client_classes,omitempty"`

	// +kubebuilder:validation:Optional
	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

type ClassificationInitParameters struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UserDefinedType *string `json:"userDefinedType,omitempty" tf:"user_defined_type,omitempty"`
}

type ClassificationObservation struct {
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UserDefinedType *string `json:"userDefinedType,omitempty" tf:"user_defined_type,omitempty"`
}

type ClassificationParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UserDefinedType *string `json:"userDefinedType,omitempty" tf:"user_defined_type,omitempty"`
}

type ClientIPAddrsInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPAddrsObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPAddrsParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClientIPPrefixesIPAddrInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPPrefixesIPAddrObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPPrefixesIPAddrParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClientIPPrefixesInitParameters struct {
	IPAddr []ClientIPPrefixesIPAddrInitParameters `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type ClientIPPrefixesObservation struct {
	IPAddr []ClientIPPrefixesIPAddrObservation `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`
}

type ClientIPPrefixesParameters struct {

	// +kubebuilder:validation:Optional
	IPAddr []ClientIPPrefixesIPAddrParameters `json:"ipAddr" tf:"ip_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Mask *string `json:"mask" tf:"mask,omitempty"`
}

type ClientIPRangesBeginInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPRangesBeginObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPRangesBeginParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClientIPRangesEndInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPRangesEndObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientIPRangesEndParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type IdentifierMatcherInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type IdentifierMatcherObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type IdentifierMatcherParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	// +kubebuilder:validation:Optional
	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type MappingRulesMatchInitParameters struct {
	ClassMatcher []ClassMatcherInitParameters `json:"classMatcher,omitempty" tf:"class_matcher,omitempty"`

	ClientIP []MatchClientIPInitParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	ComponentMatcher *string `json:"componentMatcher,omitempty" tf:"component_matcher,omitempty"`

	Hdrs []MatchHdrsInitParameters `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	HostHdr []MatchHostHdrInitParameters `json:"hostHdr,omitempty" tf:"host_hdr,omitempty"`

	IdentifierMatcher []IdentifierMatcherInitParameters `json:"identifierMatcher,omitempty" tf:"identifier_matcher,omitempty"`

	Method []MatchMethodInitParameters `json:"method,omitempty" tf:"method,omitempty"`

	Path []MatchPathInitParameters `json:"path,omitempty" tf:"path,omitempty"`

	TypeMatcher []TypeMatcherInitParameters `json:"typeMatcher,omitempty" tf:"type_matcher,omitempty"`
}

type MappingRulesMatchObservation struct {
	ClassMatcher []ClassMatcherObservation `json:"classMatcher,omitempty" tf:"class_matcher,omitempty"`

	ClientIP []MatchClientIPObservation `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	ComponentMatcher *string `json:"componentMatcher,omitempty" tf:"component_matcher,omitempty"`

	Hdrs []MatchHdrsObservation `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	HostHdr []MatchHostHdrObservation `json:"hostHdr,omitempty" tf:"host_hdr,omitempty"`

	IdentifierMatcher []IdentifierMatcherObservation `json:"identifierMatcher,omitempty" tf:"identifier_matcher,omitempty"`

	Method []MatchMethodObservation `json:"method,omitempty" tf:"method,omitempty"`

	Path []MatchPathObservation `json:"path,omitempty" tf:"path,omitempty"`

	TypeMatcher []TypeMatcherObservation `json:"typeMatcher,omitempty" tf:"type_matcher,omitempty"`
}

type MappingRulesMatchParameters struct {

	// +kubebuilder:validation:Optional
	ClassMatcher []ClassMatcherParameters `json:"classMatcher,omitempty" tf:"class_matcher,omitempty"`

	// +kubebuilder:validation:Optional
	ClientIP []MatchClientIPParameters `json:"clientIp,omitempty" tf:"client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	ComponentMatcher *string `json:"componentMatcher,omitempty" tf:"component_matcher,omitempty"`

	// +kubebuilder:validation:Optional
	Hdrs []MatchHdrsParameters `json:"hdrs,omitempty" tf:"hdrs,omitempty"`

	// +kubebuilder:validation:Optional
	HostHdr []MatchHostHdrParameters `json:"hostHdr,omitempty" tf:"host_hdr,omitempty"`

	// +kubebuilder:validation:Optional
	IdentifierMatcher []IdentifierMatcherParameters `json:"identifierMatcher,omitempty" tf:"identifier_matcher,omitempty"`

	// +kubebuilder:validation:Optional
	Method []MatchMethodParameters `json:"method,omitempty" tf:"method,omitempty"`

	// +kubebuilder:validation:Optional
	Path []MatchPathParameters `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	TypeMatcher []TypeMatcherParameters `json:"typeMatcher,omitempty" tf:"type_matcher,omitempty"`
}

type MatchClientIPInitParameters struct {
	Addrs []ClientIPAddrsInitParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []ClientIPPrefixesInitParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []MatchClientIPRangesInitParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type MatchClientIPObservation struct {
	Addrs []ClientIPAddrsObservation `json:"addrs,omitempty" tf:"addrs,omitempty"`

	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Prefixes []ClientIPPrefixesObservation `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	Ranges []MatchClientIPRangesObservation `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type MatchClientIPParameters struct {

	// +kubebuilder:validation:Optional
	Addrs []ClientIPAddrsParameters `json:"addrs,omitempty" tf:"addrs,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRefs []*string `json:"groupRefs,omitempty" tf:"group_refs,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Prefixes []ClientIPPrefixesParameters `json:"prefixes,omitempty" tf:"prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	Ranges []MatchClientIPRangesParameters `json:"ranges,omitempty" tf:"ranges,omitempty"`
}

type MatchClientIPRangesInitParameters struct {
	Begin []ClientIPRangesBeginInitParameters `json:"begin,omitempty" tf:"begin,omitempty"`

	End []ClientIPRangesEndInitParameters `json:"end,omitempty" tf:"end,omitempty"`
}

type MatchClientIPRangesObservation struct {
	Begin []ClientIPRangesBeginObservation `json:"begin,omitempty" tf:"begin,omitempty"`

	End []ClientIPRangesEndObservation `json:"end,omitempty" tf:"end,omitempty"`
}

type MatchClientIPRangesParameters struct {

	// +kubebuilder:validation:Optional
	Begin []ClientIPRangesBeginParameters `json:"begin" tf:"begin,omitempty"`

	// +kubebuilder:validation:Optional
	End []ClientIPRangesEndParameters `json:"end" tf:"end,omitempty"`
}

type MatchHdrsInitParameters struct {
	Hdr *string `json:"hdr,omitempty" tf:"hdr,omitempty"`

	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchHdrsObservation struct {
	Hdr *string `json:"hdr,omitempty" tf:"hdr,omitempty"`

	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchHdrsParameters struct {

	// +kubebuilder:validation:Optional
	Hdr *string `json:"hdr" tf:"hdr,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchHostHdrInitParameters struct {
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchHostHdrObservation struct {
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchHostHdrParameters struct {

	// +kubebuilder:validation:Optional
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchMethodInitParameters struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Methods []*string `json:"methods,omitempty" tf:"methods,omitempty"`
}

type MatchMethodObservation struct {
	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	Methods []*string `json:"methods,omitempty" tf:"methods,omitempty"`
}

type MatchMethodParameters struct {

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Methods []*string `json:"methods" tf:"methods,omitempty"`
}

type MatchPathInitParameters struct {
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	MatchDecodedString *string `json:"matchDecodedString,omitempty" tf:"match_decoded_string,omitempty"`

	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type MatchPathObservation struct {
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	MatchCriteria *string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	MatchDecodedString *string `json:"matchDecodedString,omitempty" tf:"match_decoded_string,omitempty"`

	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type MatchPathParameters struct {

	// +kubebuilder:validation:Optional
	MatchCase *string `json:"matchCase,omitempty" tf:"match_case,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCriteria *string `json:"matchCriteria" tf:"match_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	MatchDecodedString *string `json:"matchDecodedString,omitempty" tf:"match_decoded_string,omitempty"`

	// +kubebuilder:validation:Optional
	MatchStr []*string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	// +kubebuilder:validation:Optional
	StringGroupRefs []*string `json:"stringGroupRefs,omitempty" tf:"string_group_refs,omitempty"`
}

type TypeMatcherInitParameters struct {
	ClientTypes []*string `json:"clientTypes,omitempty" tf:"client_types,omitempty"`

	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

type TypeMatcherObservation struct {
	ClientTypes []*string `json:"clientTypes,omitempty" tf:"client_types,omitempty"`

	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

type TypeMatcherParameters struct {

	// +kubebuilder:validation:Optional
	ClientTypes []*string `json:"clientTypes" tf:"client_types,omitempty"`

	// +kubebuilder:validation:Optional
	Op *string `json:"op,omitempty" tf:"op,omitempty"`
}

// BotMappingSpec defines the desired state of BotMapping
type BotMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotMappingParameters `json:"forProvider"`
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
	InitProvider BotMappingInitParameters `json:"initProvider,omitempty"`
}

// BotMappingStatus defines the observed state of BotMapping.
type BotMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotMapping is the Schema for the BotMappings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type BotMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   BotMappingSpec   `json:"spec"`
	Status BotMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotMappingList contains a list of BotMappings
type BotMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotMapping `json:"items"`
}

// Repository type metadata.
var (
	BotMapping_Kind             = "BotMapping"
	BotMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotMapping_Kind}.String()
	BotMapping_KindAPIVersion   = BotMapping_Kind + "." + CRDGroupVersion.String()
	BotMapping_GroupVersionKind = CRDGroupVersion.WithKind(BotMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&BotMapping{}, &BotMappingList{})
}