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

type MicroServiceGroupConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MicroServiceGroupConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MicroServiceGroupConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MicroServiceGroupInitParameters struct {
	ConfigpbAttributes []MicroServiceGroupConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServiceRefs []*string `json:"serviceRefs,omitempty" tf:"service_refs,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type MicroServiceGroupObservation struct {
	ConfigpbAttributes []MicroServiceGroupConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServiceRefs []*string `json:"serviceRefs,omitempty" tf:"service_refs,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type MicroServiceGroupParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []MicroServiceGroupConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRefs []*string `json:"serviceRefs,omitempty" tf:"service_refs,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// MicroServiceGroupSpec defines the desired state of MicroServiceGroup
type MicroServiceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MicroServiceGroupParameters `json:"forProvider"`
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
	InitProvider MicroServiceGroupInitParameters `json:"initProvider,omitempty"`
}

// MicroServiceGroupStatus defines the observed state of MicroServiceGroup.
type MicroServiceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MicroServiceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MicroServiceGroup is the Schema for the MicroServiceGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type MicroServiceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   MicroServiceGroupSpec   `json:"spec"`
	Status MicroServiceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MicroServiceGroupList contains a list of MicroServiceGroups
type MicroServiceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicroServiceGroup `json:"items"`
}

// Repository type metadata.
var (
	MicroServiceGroup_Kind             = "MicroServiceGroup"
	MicroServiceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MicroServiceGroup_Kind}.String()
	MicroServiceGroup_KindAPIVersion   = MicroServiceGroup_Kind + "." + CRDGroupVersion.String()
	MicroServiceGroup_GroupVersionKind = CRDGroupVersion.WithKind(MicroServiceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MicroServiceGroup{}, &MicroServiceGroupList{})
}
