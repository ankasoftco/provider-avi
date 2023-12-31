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

type ALBServicesFileUploadInitParameters struct {
	CaseID *string `json:"caseId,omitempty" tf:"case_id,omitempty"`

	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	S3Directory *string `json:"s3Directory,omitempty" tf:"s3_directory,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ALBServicesFileUploadObservation struct {
	CaseID *string `json:"caseId,omitempty" tf:"case_id,omitempty"`

	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	S3Directory *string `json:"s3Directory,omitempty" tf:"s3_directory,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ALBServicesFileUploadParameters struct {

	// +kubebuilder:validation:Optional
	CaseID *string `json:"caseId,omitempty" tf:"case_id,omitempty"`

	// +kubebuilder:validation:Optional
	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	S3Directory *string `json:"s3Directory,omitempty" tf:"s3_directory,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// ALBServicesFileUploadSpec defines the desired state of ALBServicesFileUpload
type ALBServicesFileUploadSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ALBServicesFileUploadParameters `json:"forProvider"`
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
	InitProvider ALBServicesFileUploadInitParameters `json:"initProvider,omitempty"`
}

// ALBServicesFileUploadStatus defines the observed state of ALBServicesFileUpload.
type ALBServicesFileUploadStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ALBServicesFileUploadObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ALBServicesFileUpload is the Schema for the ALBServicesFileUploads API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type ALBServicesFileUpload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filePath) || has(self.initProvider.filePath)",message="filePath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ALBServicesFileUploadSpec   `json:"spec"`
	Status ALBServicesFileUploadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ALBServicesFileUploadList contains a list of ALBServicesFileUploads
type ALBServicesFileUploadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ALBServicesFileUpload `json:"items"`
}

// Repository type metadata.
var (
	ALBServicesFileUpload_Kind             = "ALBServicesFileUpload"
	ALBServicesFileUpload_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ALBServicesFileUpload_Kind}.String()
	ALBServicesFileUpload_KindAPIVersion   = ALBServicesFileUpload_Kind + "." + CRDGroupVersion.String()
	ALBServicesFileUpload_GroupVersionKind = CRDGroupVersion.WithKind(ALBServicesFileUpload_Kind)
)

func init() {
	SchemeBuilder.Register(&ALBServicesFileUpload{}, &ALBServicesFileUploadList{})
}
