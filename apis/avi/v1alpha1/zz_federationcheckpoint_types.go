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

type FederationCheckpointConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FederationCheckpointConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FederationCheckpointConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FederationCheckpointInitParameters struct {
	ConfigpbAttributes []FederationCheckpointConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	IsFederated *string `json:"isFederated,omitempty" tf:"is_federated,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type FederationCheckpointObservation struct {
	ConfigpbAttributes []FederationCheckpointConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsFederated *string `json:"isFederated,omitempty" tf:"is_federated,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type FederationCheckpointParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []FederationCheckpointConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IsFederated *string `json:"isFederated,omitempty" tf:"is_federated,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// FederationCheckpointSpec defines the desired state of FederationCheckpoint
type FederationCheckpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FederationCheckpointParameters `json:"forProvider"`
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
	InitProvider FederationCheckpointInitParameters `json:"initProvider,omitempty"`
}

// FederationCheckpointStatus defines the observed state of FederationCheckpoint.
type FederationCheckpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FederationCheckpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FederationCheckpoint is the Schema for the FederationCheckpoints API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type FederationCheckpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   FederationCheckpointSpec   `json:"spec"`
	Status FederationCheckpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FederationCheckpointList contains a list of FederationCheckpoints
type FederationCheckpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederationCheckpoint `json:"items"`
}

// Repository type metadata.
var (
	FederationCheckpoint_Kind             = "FederationCheckpoint"
	FederationCheckpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FederationCheckpoint_Kind}.String()
	FederationCheckpoint_KindAPIVersion   = FederationCheckpoint_Kind + "." + CRDGroupVersion.String()
	FederationCheckpoint_GroupVersionKind = CRDGroupVersion.WithKind(FederationCheckpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&FederationCheckpoint{}, &FederationCheckpointList{})
}