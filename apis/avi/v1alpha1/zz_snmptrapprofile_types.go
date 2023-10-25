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

type SnmpTrapProfileConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SnmpTrapProfileConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SnmpTrapProfileConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SnmpTrapProfileInitParameters struct {
	ConfigpbAttributes []SnmpTrapProfileConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	TrapServers []TrapServersInitParameters `json:"trapServers,omitempty" tf:"trap_servers,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SnmpTrapProfileObservation struct {
	ConfigpbAttributes []SnmpTrapProfileConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	TrapServers []TrapServersObservation `json:"trapServers,omitempty" tf:"trap_servers,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SnmpTrapProfileParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []SnmpTrapProfileConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	TrapServers []TrapServersParameters `json:"trapServers,omitempty" tf:"trap_servers,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type TrapServersIPAddrInitParameters struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrapServersIPAddrObservation struct {
	Addr *string `json:"addr,omitempty" tf:"addr,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrapServersIPAddrParameters struct {

	// +kubebuilder:validation:Optional
	Addr *string `json:"addr" tf:"addr,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type TrapServersInitParameters struct {
	IPAddr []TrapServersIPAddrInitParameters `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	User []UserInitParameters `json:"user,omitempty" tf:"user,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TrapServersObservation struct {
	IPAddr []TrapServersIPAddrObservation `json:"ipAddr,omitempty" tf:"ip_addr,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TrapServersParameters struct {

	// +kubebuilder:validation:Optional
	CommunitySecretRef *v1.SecretKeySelector `json:"communitySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IPAddr []TrapServersIPAddrParameters `json:"ipAddr" tf:"ip_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type UserInitParameters struct {
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	PrivType *string `json:"privType,omitempty" tf:"priv_type,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	PrivType *string `json:"privType,omitempty" tf:"priv_type,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// +kubebuilder:validation:Optional
	AuthPassphraseSecretRef *v1.SecretKeySelector `json:"authPassphraseSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// +kubebuilder:validation:Optional
	PrivPassphraseSecretRef *v1.SecretKeySelector `json:"privPassphraseSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrivType *string `json:"privType,omitempty" tf:"priv_type,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// SnmpTrapProfileSpec defines the desired state of SnmpTrapProfile
type SnmpTrapProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnmpTrapProfileParameters `json:"forProvider"`
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
	InitProvider SnmpTrapProfileInitParameters `json:"initProvider,omitempty"`
}

// SnmpTrapProfileStatus defines the observed state of SnmpTrapProfile.
type SnmpTrapProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnmpTrapProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnmpTrapProfile is the Schema for the SnmpTrapProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type SnmpTrapProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SnmpTrapProfileSpec   `json:"spec"`
	Status SnmpTrapProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnmpTrapProfileList contains a list of SnmpTrapProfiles
type SnmpTrapProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnmpTrapProfile `json:"items"`
}

// Repository type metadata.
var (
	SnmpTrapProfile_Kind             = "SnmpTrapProfile"
	SnmpTrapProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnmpTrapProfile_Kind}.String()
	SnmpTrapProfile_KindAPIVersion   = SnmpTrapProfile_Kind + "." + CRDGroupVersion.String()
	SnmpTrapProfile_GroupVersionKind = CRDGroupVersion.WithKind(SnmpTrapProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&SnmpTrapProfile{}, &SnmpTrapProfileList{})
}