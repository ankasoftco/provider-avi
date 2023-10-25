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

type AutoScaleLaunchConfigConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutoScaleLaunchConfigConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutoScaleLaunchConfigConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutoScaleLaunchConfigInitParameters struct {
	ConfigpbAttributes []AutoScaleLaunchConfigConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	Markers []AutoScaleLaunchConfigMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	Mesos []MesosInitParameters `json:"mesos,omitempty" tf:"mesos,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Openstack []OpenstackInitParameters `json:"openstack,omitempty" tf:"openstack,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	UseExternalAsg *string `json:"useExternalAsg,omitempty" tf:"use_external_asg,omitempty"`
}

type AutoScaleLaunchConfigMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AutoScaleLaunchConfigMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AutoScaleLaunchConfigMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AutoScaleLaunchConfigObservation struct {
	ConfigpbAttributes []AutoScaleLaunchConfigConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	Markers []AutoScaleLaunchConfigMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	Mesos []MesosObservation `json:"mesos,omitempty" tf:"mesos,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Openstack []OpenstackObservation `json:"openstack,omitempty" tf:"openstack,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	UseExternalAsg *string `json:"useExternalAsg,omitempty" tf:"use_external_asg,omitempty"`
}

type AutoScaleLaunchConfigParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []AutoScaleLaunchConfigConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []AutoScaleLaunchConfigMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	Mesos []MesosParameters `json:"mesos,omitempty" tf:"mesos,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Openstack []OpenstackParameters `json:"openstack,omitempty" tf:"openstack,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// +kubebuilder:validation:Optional
	UseExternalAsg *string `json:"useExternalAsg,omitempty" tf:"use_external_asg,omitempty"`
}

type MesosInitParameters struct {
	Force *string `json:"force,omitempty" tf:"force,omitempty"`
}

type MesosObservation struct {
	Force *string `json:"force,omitempty" tf:"force,omitempty"`
}

type MesosParameters struct {

	// +kubebuilder:validation:Optional
	Force *string `json:"force,omitempty" tf:"force,omitempty"`
}

type OpenstackInitParameters struct {
	HeatScaleDownURL *string `json:"heatScaleDownUrl,omitempty" tf:"heat_scale_down_url,omitempty"`

	HeatScaleUpURL *string `json:"heatScaleUpUrl,omitempty" tf:"heat_scale_up_url,omitempty"`
}

type OpenstackObservation struct {
	HeatScaleDownURL *string `json:"heatScaleDownUrl,omitempty" tf:"heat_scale_down_url,omitempty"`

	HeatScaleUpURL *string `json:"heatScaleUpUrl,omitempty" tf:"heat_scale_up_url,omitempty"`
}

type OpenstackParameters struct {

	// +kubebuilder:validation:Optional
	HeatScaleDownURL *string `json:"heatScaleDownUrl,omitempty" tf:"heat_scale_down_url,omitempty"`

	// +kubebuilder:validation:Optional
	HeatScaleUpURL *string `json:"heatScaleUpUrl,omitempty" tf:"heat_scale_up_url,omitempty"`
}

// AutoScaleLaunchConfigSpec defines the desired state of AutoScaleLaunchConfig
type AutoScaleLaunchConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoScaleLaunchConfigParameters `json:"forProvider"`
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
	InitProvider AutoScaleLaunchConfigInitParameters `json:"initProvider,omitempty"`
}

// AutoScaleLaunchConfigStatus defines the observed state of AutoScaleLaunchConfig.
type AutoScaleLaunchConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoScaleLaunchConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScaleLaunchConfig is the Schema for the AutoScaleLaunchConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type AutoScaleLaunchConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   AutoScaleLaunchConfigSpec   `json:"spec"`
	Status AutoScaleLaunchConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScaleLaunchConfigList contains a list of AutoScaleLaunchConfigs
type AutoScaleLaunchConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoScaleLaunchConfig `json:"items"`
}

// Repository type metadata.
var (
	AutoScaleLaunchConfig_Kind             = "AutoScaleLaunchConfig"
	AutoScaleLaunchConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoScaleLaunchConfig_Kind}.String()
	AutoScaleLaunchConfig_KindAPIVersion   = AutoScaleLaunchConfig_Kind + "." + CRDGroupVersion.String()
	AutoScaleLaunchConfig_GroupVersionKind = CRDGroupVersion.WithKind(AutoScaleLaunchConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoScaleLaunchConfig{}, &AutoScaleLaunchConfigList{})
}