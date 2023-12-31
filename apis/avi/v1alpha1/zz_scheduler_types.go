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

type SchedulerConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SchedulerConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SchedulerConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SchedulerInitParameters struct {
	BackupConfigRef *string `json:"backupConfigRef,omitempty" tf:"backup_config_ref,omitempty"`

	ConfigpbAttributes []SchedulerConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`

	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	FrequencyUnit *string `json:"frequencyUnit,omitempty" tf:"frequency_unit,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RunMode *string `json:"runMode,omitempty" tf:"run_mode,omitempty"`

	RunScriptRef *string `json:"runScriptRef,omitempty" tf:"run_script_ref,omitempty"`

	SchedulerAction *string `json:"schedulerAction,omitempty" tf:"scheduler_action,omitempty"`

	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SchedulerObservation struct {
	BackupConfigRef *string `json:"backupConfigRef,omitempty" tf:"backup_config_ref,omitempty"`

	ConfigpbAttributes []SchedulerConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`

	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	FrequencyUnit *string `json:"frequencyUnit,omitempty" tf:"frequency_unit,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RunMode *string `json:"runMode,omitempty" tf:"run_mode,omitempty"`

	RunScriptRef *string `json:"runScriptRef,omitempty" tf:"run_script_ref,omitempty"`

	SchedulerAction *string `json:"schedulerAction,omitempty" tf:"scheduler_action,omitempty"`

	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SchedulerParameters struct {

	// +kubebuilder:validation:Optional
	BackupConfigRef *string `json:"backupConfigRef,omitempty" tf:"backup_config_ref,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []SchedulerConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`

	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Optional
	FrequencyUnit *string `json:"frequencyUnit,omitempty" tf:"frequency_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RunMode *string `json:"runMode,omitempty" tf:"run_mode,omitempty"`

	// +kubebuilder:validation:Optional
	RunScriptRef *string `json:"runScriptRef,omitempty" tf:"run_script_ref,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulerAction *string `json:"schedulerAction,omitempty" tf:"scheduler_action,omitempty"`

	// +kubebuilder:validation:Optional
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// SchedulerSpec defines the desired state of Scheduler
type SchedulerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchedulerParameters `json:"forProvider"`
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
	InitProvider SchedulerInitParameters `json:"initProvider,omitempty"`
}

// SchedulerStatus defines the observed state of Scheduler.
type SchedulerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchedulerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Scheduler is the Schema for the Schedulers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type Scheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SchedulerSpec   `json:"spec"`
	Status SchedulerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchedulerList contains a list of Schedulers
type SchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scheduler `json:"items"`
}

// Repository type metadata.
var (
	Scheduler_Kind             = "Scheduler"
	Scheduler_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Scheduler_Kind}.String()
	Scheduler_KindAPIVersion   = Scheduler_Kind + "." + CRDGroupVersion.String()
	Scheduler_GroupVersionKind = CRDGroupVersion.WithKind(Scheduler_Kind)
)

func init() {
	SchemeBuilder.Register(&Scheduler{}, &SchedulerList{})
}
