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

type AppLearningInfoInitParameters struct {
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	URIInfo []URIInfoInitParameters `json:"uriInfo,omitempty" tf:"uri_info,omitempty"`

	VsUUID *string `json:"vsUuid,omitempty" tf:"vs_uuid,omitempty"`
}

type AppLearningInfoObservation struct {
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	URIInfo []URIInfoObservation `json:"uriInfo,omitempty" tf:"uri_info,omitempty"`

	VsUUID *string `json:"vsUuid,omitempty" tf:"vs_uuid,omitempty"`
}

type AppLearningInfoParameters struct {

	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	URIInfo []URIInfoParameters `json:"uriInfo,omitempty" tf:"uri_info,omitempty"`

	// +kubebuilder:validation:Optional
	VsUUID *string `json:"vsUuid,omitempty" tf:"vs_uuid,omitempty"`
}

type ParamInfoInitParameters struct {
	ParamHits *string `json:"paramHits,omitempty" tf:"param_hits,omitempty"`

	ParamKey *string `json:"paramKey,omitempty" tf:"param_key,omitempty"`

	ParamSizeClasses []ParamSizeClassesInitParameters `json:"paramSizeClasses,omitempty" tf:"param_size_classes,omitempty"`

	ParamTypeClasses []ParamTypeClassesInitParameters `json:"paramTypeClasses,omitempty" tf:"param_type_classes,omitempty"`
}

type ParamInfoObservation struct {
	ParamHits *string `json:"paramHits,omitempty" tf:"param_hits,omitempty"`

	ParamKey *string `json:"paramKey,omitempty" tf:"param_key,omitempty"`

	ParamSizeClasses []ParamSizeClassesObservation `json:"paramSizeClasses,omitempty" tf:"param_size_classes,omitempty"`

	ParamTypeClasses []ParamTypeClassesObservation `json:"paramTypeClasses,omitempty" tf:"param_type_classes,omitempty"`
}

type ParamInfoParameters struct {

	// +kubebuilder:validation:Optional
	ParamHits *string `json:"paramHits,omitempty" tf:"param_hits,omitempty"`

	// +kubebuilder:validation:Optional
	ParamKey *string `json:"paramKey,omitempty" tf:"param_key,omitempty"`

	// +kubebuilder:validation:Optional
	ParamSizeClasses []ParamSizeClassesParameters `json:"paramSizeClasses,omitempty" tf:"param_size_classes,omitempty"`

	// +kubebuilder:validation:Optional
	ParamTypeClasses []ParamTypeClassesParameters `json:"paramTypeClasses,omitempty" tf:"param_type_classes,omitempty"`
}

type ParamSizeClassesInitParameters struct {
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	Len *string `json:"len,omitempty" tf:"len,omitempty"`
}

type ParamSizeClassesObservation struct {
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	Len *string `json:"len,omitempty" tf:"len,omitempty"`
}

type ParamSizeClassesParameters struct {

	// +kubebuilder:validation:Optional
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	// +kubebuilder:validation:Optional
	Len *string `json:"len,omitempty" tf:"len,omitempty"`
}

type ParamTypeClassesInitParameters struct {
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParamTypeClassesObservation struct {
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParamTypeClassesParameters struct {

	// +kubebuilder:validation:Optional
	Hits *string `json:"hits,omitempty" tf:"hits,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityManagerDataInitParameters struct {
	AppLearningInfo []AppLearningInfoInitParameters `json:"appLearningInfo,omitempty" tf:"app_learning_info,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SecurityManagerDataObservation struct {
	AppLearningInfo []AppLearningInfoObservation `json:"appLearningInfo,omitempty" tf:"app_learning_info,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SecurityManagerDataParameters struct {

	// +kubebuilder:validation:Optional
	AppLearningInfo []AppLearningInfoParameters `json:"appLearningInfo,omitempty" tf:"app_learning_info,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type URIInfoInitParameters struct {
	ParamInfo []ParamInfoInitParameters `json:"paramInfo,omitempty" tf:"param_info,omitempty"`

	URIHits *string `json:"uriHits,omitempty" tf:"uri_hits,omitempty"`

	URIKey *string `json:"uriKey,omitempty" tf:"uri_key,omitempty"`
}

type URIInfoObservation struct {
	ParamInfo []ParamInfoObservation `json:"paramInfo,omitempty" tf:"param_info,omitempty"`

	URIHits *string `json:"uriHits,omitempty" tf:"uri_hits,omitempty"`

	URIKey *string `json:"uriKey,omitempty" tf:"uri_key,omitempty"`
}

type URIInfoParameters struct {

	// +kubebuilder:validation:Optional
	ParamInfo []ParamInfoParameters `json:"paramInfo,omitempty" tf:"param_info,omitempty"`

	// +kubebuilder:validation:Optional
	URIHits *string `json:"uriHits,omitempty" tf:"uri_hits,omitempty"`

	// +kubebuilder:validation:Optional
	URIKey *string `json:"uriKey,omitempty" tf:"uri_key,omitempty"`
}

// SecurityManagerDataSpec defines the desired state of SecurityManagerData
type SecurityManagerDataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityManagerDataParameters `json:"forProvider"`
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
	InitProvider SecurityManagerDataInitParameters `json:"initProvider,omitempty"`
}

// SecurityManagerDataStatus defines the observed state of SecurityManagerData.
type SecurityManagerDataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityManagerDataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityManagerData is the Schema for the SecurityManagerDatas API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type SecurityManagerData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SecurityManagerDataSpec   `json:"spec"`
	Status SecurityManagerDataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityManagerDataList contains a list of SecurityManagerDatas
type SecurityManagerDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityManagerData `json:"items"`
}

// Repository type metadata.
var (
	SecurityManagerData_Kind             = "SecurityManagerData"
	SecurityManagerData_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityManagerData_Kind}.String()
	SecurityManagerData_KindAPIVersion   = SecurityManagerData_Kind + "." + CRDGroupVersion.String()
	SecurityManagerData_GroupVersionKind = CRDGroupVersion.WithKind(SecurityManagerData_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityManagerData{}, &SecurityManagerDataList{})
}
