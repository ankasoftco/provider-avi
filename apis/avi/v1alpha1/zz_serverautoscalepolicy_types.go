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

type ScheduledScalingsInitParameters struct {
	AutoscalingDuration *string `json:"autoscalingDuration,omitempty" tf:"autoscaling_duration,omitempty"`

	CronExpression *string `json:"cronExpression,omitempty" tf:"cron_expression,omitempty"`

	DesiredCapacity *string `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	ScheduleMaxStep *string `json:"scheduleMaxStep,omitempty" tf:"schedule_max_step,omitempty"`

	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type ScheduledScalingsObservation struct {
	AutoscalingDuration *string `json:"autoscalingDuration,omitempty" tf:"autoscaling_duration,omitempty"`

	CronExpression *string `json:"cronExpression,omitempty" tf:"cron_expression,omitempty"`

	DesiredCapacity *string `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	ScheduleMaxStep *string `json:"scheduleMaxStep,omitempty" tf:"schedule_max_step,omitempty"`

	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type ScheduledScalingsParameters struct {

	// +kubebuilder:validation:Optional
	AutoscalingDuration *string `json:"autoscalingDuration,omitempty" tf:"autoscaling_duration,omitempty"`

	// +kubebuilder:validation:Optional
	CronExpression *string `json:"cronExpression,omitempty" tf:"cron_expression,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCapacity *string `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Enable *string `json:"enable,omitempty" tf:"enable,omitempty"`

	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduleMaxStep *string `json:"scheduleMaxStep,omitempty" tf:"schedule_max_step,omitempty"`

	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type ServerAutoScalePolicyConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ServerAutoScalePolicyConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ServerAutoScalePolicyConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ServerAutoScalePolicyInitParameters struct {
	ConfigpbAttributes []ServerAutoScalePolicyConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	DelayForServerGarbageCollection *string `json:"delayForServerGarbageCollection,omitempty" tf:"delay_for_server_garbage_collection,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	IntelligentAutoscale *string `json:"intelligentAutoscale,omitempty" tf:"intelligent_autoscale,omitempty"`

	IntelligentScaleinMargin *string `json:"intelligentScaleinMargin,omitempty" tf:"intelligent_scalein_margin,omitempty"`

	IntelligentScaleoutMargin *string `json:"intelligentScaleoutMargin,omitempty" tf:"intelligent_scaleout_margin,omitempty"`

	Markers []ServerAutoScalePolicyMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	MaxScaleinAdjustmentStep *string `json:"maxScaleinAdjustmentStep,omitempty" tf:"max_scalein_adjustment_step,omitempty"`

	MaxScaleoutAdjustmentStep *string `json:"maxScaleoutAdjustmentStep,omitempty" tf:"max_scaleout_adjustment_step,omitempty"`

	MaxSize *string `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	MinSize *string `json:"minSize,omitempty" tf:"min_size,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ScaleinAlertconfigRefs []*string `json:"scaleinAlertconfigRefs,omitempty" tf:"scalein_alertconfig_refs,omitempty"`

	ScaleinCooldown *string `json:"scaleinCooldown,omitempty" tf:"scalein_cooldown,omitempty"`

	ScaleoutAlertconfigRefs []*string `json:"scaleoutAlertconfigRefs,omitempty" tf:"scaleout_alertconfig_refs,omitempty"`

	ScaleoutCooldown *string `json:"scaleoutCooldown,omitempty" tf:"scaleout_cooldown,omitempty"`

	ScheduledScalings []ScheduledScalingsInitParameters `json:"scheduledScalings,omitempty" tf:"scheduled_scalings,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	UsePredictedLoad *string `json:"usePredictedLoad,omitempty" tf:"use_predicted_load,omitempty"`
}

type ServerAutoScalePolicyMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ServerAutoScalePolicyMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ServerAutoScalePolicyMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ServerAutoScalePolicyObservation struct {
	ConfigpbAttributes []ServerAutoScalePolicyConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	DelayForServerGarbageCollection *string `json:"delayForServerGarbageCollection,omitempty" tf:"delay_for_server_garbage_collection,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IntelligentAutoscale *string `json:"intelligentAutoscale,omitempty" tf:"intelligent_autoscale,omitempty"`

	IntelligentScaleinMargin *string `json:"intelligentScaleinMargin,omitempty" tf:"intelligent_scalein_margin,omitempty"`

	IntelligentScaleoutMargin *string `json:"intelligentScaleoutMargin,omitempty" tf:"intelligent_scaleout_margin,omitempty"`

	Markers []ServerAutoScalePolicyMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	MaxScaleinAdjustmentStep *string `json:"maxScaleinAdjustmentStep,omitempty" tf:"max_scalein_adjustment_step,omitempty"`

	MaxScaleoutAdjustmentStep *string `json:"maxScaleoutAdjustmentStep,omitempty" tf:"max_scaleout_adjustment_step,omitempty"`

	MaxSize *string `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	MinSize *string `json:"minSize,omitempty" tf:"min_size,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ScaleinAlertconfigRefs []*string `json:"scaleinAlertconfigRefs,omitempty" tf:"scalein_alertconfig_refs,omitempty"`

	ScaleinCooldown *string `json:"scaleinCooldown,omitempty" tf:"scalein_cooldown,omitempty"`

	ScaleoutAlertconfigRefs []*string `json:"scaleoutAlertconfigRefs,omitempty" tf:"scaleout_alertconfig_refs,omitempty"`

	ScaleoutCooldown *string `json:"scaleoutCooldown,omitempty" tf:"scaleout_cooldown,omitempty"`

	ScheduledScalings []ScheduledScalingsObservation `json:"scheduledScalings,omitempty" tf:"scheduled_scalings,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	UsePredictedLoad *string `json:"usePredictedLoad,omitempty" tf:"use_predicted_load,omitempty"`
}

type ServerAutoScalePolicyParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []ServerAutoScalePolicyConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	DelayForServerGarbageCollection *string `json:"delayForServerGarbageCollection,omitempty" tf:"delay_for_server_garbage_collection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntelligentAutoscale *string `json:"intelligentAutoscale,omitempty" tf:"intelligent_autoscale,omitempty"`

	// +kubebuilder:validation:Optional
	IntelligentScaleinMargin *string `json:"intelligentScaleinMargin,omitempty" tf:"intelligent_scalein_margin,omitempty"`

	// +kubebuilder:validation:Optional
	IntelligentScaleoutMargin *string `json:"intelligentScaleoutMargin,omitempty" tf:"intelligent_scaleout_margin,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []ServerAutoScalePolicyMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	MaxScaleinAdjustmentStep *string `json:"maxScaleinAdjustmentStep,omitempty" tf:"max_scalein_adjustment_step,omitempty"`

	// +kubebuilder:validation:Optional
	MaxScaleoutAdjustmentStep *string `json:"maxScaleoutAdjustmentStep,omitempty" tf:"max_scaleout_adjustment_step,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSize *string `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *string `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleinAlertconfigRefs []*string `json:"scaleinAlertconfigRefs,omitempty" tf:"scalein_alertconfig_refs,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleinCooldown *string `json:"scaleinCooldown,omitempty" tf:"scalein_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleoutAlertconfigRefs []*string `json:"scaleoutAlertconfigRefs,omitempty" tf:"scaleout_alertconfig_refs,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleoutCooldown *string `json:"scaleoutCooldown,omitempty" tf:"scaleout_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledScalings []ScheduledScalingsParameters `json:"scheduledScalings,omitempty" tf:"scheduled_scalings,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// +kubebuilder:validation:Optional
	UsePredictedLoad *string `json:"usePredictedLoad,omitempty" tf:"use_predicted_load,omitempty"`
}

// ServerAutoScalePolicySpec defines the desired state of ServerAutoScalePolicy
type ServerAutoScalePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerAutoScalePolicyParameters `json:"forProvider"`
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
	InitProvider ServerAutoScalePolicyInitParameters `json:"initProvider,omitempty"`
}

// ServerAutoScalePolicyStatus defines the observed state of ServerAutoScalePolicy.
type ServerAutoScalePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerAutoScalePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServerAutoScalePolicy is the Schema for the ServerAutoScalePolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type ServerAutoScalePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ServerAutoScalePolicySpec   `json:"spec"`
	Status ServerAutoScalePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerAutoScalePolicyList contains a list of ServerAutoScalePolicys
type ServerAutoScalePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerAutoScalePolicy `json:"items"`
}

// Repository type metadata.
var (
	ServerAutoScalePolicy_Kind             = "ServerAutoScalePolicy"
	ServerAutoScalePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerAutoScalePolicy_Kind}.String()
	ServerAutoScalePolicy_KindAPIVersion   = ServerAutoScalePolicy_Kind + "." + CRDGroupVersion.String()
	ServerAutoScalePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServerAutoScalePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerAutoScalePolicy{}, &ServerAutoScalePolicyList{})
}
