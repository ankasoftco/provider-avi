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

type ALBServicesJobConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ALBServicesJobConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ALBServicesJobConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ALBServicesJobInitParameters struct {
	Command *string `json:"command,omitempty" tf:"command,omitempty"`

	ConfigpbAttributes []ALBServicesJobConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	EndTime []EndTimeInitParameters `json:"endTime,omitempty" tf:"end_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Params []ParamsInitParameters `json:"params,omitempty" tf:"params,omitempty"`

	PulseJobID *string `json:"pulseJobId,omitempty" tf:"pulse_job_id,omitempty"`

	PulseSyncStatus *string `json:"pulseSyncStatus,omitempty" tf:"pulse_sync_status,omitempty"`

	Result *string `json:"result,omitempty" tf:"result,omitempty"`

	StartTime []StartTimeInitParameters `json:"startTime,omitempty" tf:"start_time,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ALBServicesJobObservation struct {
	Command *string `json:"command,omitempty" tf:"command,omitempty"`

	ConfigpbAttributes []ALBServicesJobConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	EndTime []EndTimeObservation `json:"endTime,omitempty" tf:"end_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Params []ParamsObservation `json:"params,omitempty" tf:"params,omitempty"`

	PulseJobID *string `json:"pulseJobId,omitempty" tf:"pulse_job_id,omitempty"`

	PulseSyncStatus *string `json:"pulseSyncStatus,omitempty" tf:"pulse_sync_status,omitempty"`

	Result *string `json:"result,omitempty" tf:"result,omitempty"`

	StartTime []StartTimeObservation `json:"startTime,omitempty" tf:"start_time,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ALBServicesJobParameters struct {

	// +kubebuilder:validation:Optional
	Command *string `json:"command,omitempty" tf:"command,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []ALBServicesJobConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	EndTime []EndTimeParameters `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Params []ParamsParameters `json:"params,omitempty" tf:"params,omitempty"`

	// +kubebuilder:validation:Optional
	PulseJobID *string `json:"pulseJobId,omitempty" tf:"pulse_job_id,omitempty"`

	// +kubebuilder:validation:Optional
	PulseSyncStatus *string `json:"pulseSyncStatus,omitempty" tf:"pulse_sync_status,omitempty"`

	// +kubebuilder:validation:Optional
	Result *string `json:"result,omitempty" tf:"result,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime []StartTimeParameters `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type EndTimeInitParameters struct {
	Secs *string `json:"secs,omitempty" tf:"secs,omitempty"`

	Usecs *string `json:"usecs,omitempty" tf:"usecs,omitempty"`
}

type EndTimeObservation struct {
	Secs *string `json:"secs,omitempty" tf:"secs,omitempty"`

	Usecs *string `json:"usecs,omitempty" tf:"usecs,omitempty"`
}

type EndTimeParameters struct {

	// +kubebuilder:validation:Optional
	Secs *string `json:"secs" tf:"secs,omitempty"`

	// +kubebuilder:validation:Optional
	Usecs *string `json:"usecs" tf:"usecs,omitempty"`
}

type ParamsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParamsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParamsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StartTimeInitParameters struct {
	Secs *string `json:"secs,omitempty" tf:"secs,omitempty"`

	Usecs *string `json:"usecs,omitempty" tf:"usecs,omitempty"`
}

type StartTimeObservation struct {
	Secs *string `json:"secs,omitempty" tf:"secs,omitempty"`

	Usecs *string `json:"usecs,omitempty" tf:"usecs,omitempty"`
}

type StartTimeParameters struct {

	// +kubebuilder:validation:Optional
	Secs *string `json:"secs" tf:"secs,omitempty"`

	// +kubebuilder:validation:Optional
	Usecs *string `json:"usecs" tf:"usecs,omitempty"`
}

// ALBServicesJobSpec defines the desired state of ALBServicesJob
type ALBServicesJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ALBServicesJobParameters `json:"forProvider"`
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
	InitProvider ALBServicesJobInitParameters `json:"initProvider,omitempty"`
}

// ALBServicesJobStatus defines the observed state of ALBServicesJob.
type ALBServicesJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ALBServicesJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ALBServicesJob is the Schema for the ALBServicesJobs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type ALBServicesJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.command) || has(self.initProvider.command)",message="command is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ALBServicesJobSpec   `json:"spec"`
	Status ALBServicesJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ALBServicesJobList contains a list of ALBServicesJobs
type ALBServicesJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ALBServicesJob `json:"items"`
}

// Repository type metadata.
var (
	ALBServicesJob_Kind             = "ALBServicesJob"
	ALBServicesJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ALBServicesJob_Kind}.String()
	ALBServicesJob_KindAPIVersion   = ALBServicesJob_Kind + "." + CRDGroupVersion.String()
	ALBServicesJob_GroupVersionKind = CRDGroupVersion.WithKind(ALBServicesJob_Kind)
)

func init() {
	SchemeBuilder.Register(&ALBServicesJob{}, &ALBServicesJobList{})
}