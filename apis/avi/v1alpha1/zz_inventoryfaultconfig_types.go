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

type ControllerFaultsInitParameters struct {
	BackupSchedulerFaults *string `json:"backupSchedulerFaults,omitempty" tf:"backup_scheduler_faults,omitempty"`

	ClusterFaults *string `json:"clusterFaults,omitempty" tf:"cluster_faults,omitempty"`

	DeprecatedAPIVersionFaults *string `json:"deprecatedApiVersionFaults,omitempty" tf:"deprecated_api_version_faults,omitempty"`

	LicenseFaults *string `json:"licenseFaults,omitempty" tf:"license_faults,omitempty"`

	MigrationFaults *string `json:"migrationFaults,omitempty" tf:"migration_faults,omitempty"`

	SslprofileFaults *string `json:"sslprofileFaults,omitempty" tf:"sslprofile_faults,omitempty"`
}

type ControllerFaultsObservation struct {
	BackupSchedulerFaults *string `json:"backupSchedulerFaults,omitempty" tf:"backup_scheduler_faults,omitempty"`

	ClusterFaults *string `json:"clusterFaults,omitempty" tf:"cluster_faults,omitempty"`

	DeprecatedAPIVersionFaults *string `json:"deprecatedApiVersionFaults,omitempty" tf:"deprecated_api_version_faults,omitempty"`

	LicenseFaults *string `json:"licenseFaults,omitempty" tf:"license_faults,omitempty"`

	MigrationFaults *string `json:"migrationFaults,omitempty" tf:"migration_faults,omitempty"`

	SslprofileFaults *string `json:"sslprofileFaults,omitempty" tf:"sslprofile_faults,omitempty"`
}

type ControllerFaultsParameters struct {

	// +kubebuilder:validation:Optional
	BackupSchedulerFaults *string `json:"backupSchedulerFaults,omitempty" tf:"backup_scheduler_faults,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterFaults *string `json:"clusterFaults,omitempty" tf:"cluster_faults,omitempty"`

	// +kubebuilder:validation:Optional
	DeprecatedAPIVersionFaults *string `json:"deprecatedApiVersionFaults,omitempty" tf:"deprecated_api_version_faults,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseFaults *string `json:"licenseFaults,omitempty" tf:"license_faults,omitempty"`

	// +kubebuilder:validation:Optional
	MigrationFaults *string `json:"migrationFaults,omitempty" tf:"migration_faults,omitempty"`

	// +kubebuilder:validation:Optional
	SslprofileFaults *string `json:"sslprofileFaults,omitempty" tf:"sslprofile_faults,omitempty"`
}

type InventoryFaultConfigConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InventoryFaultConfigConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InventoryFaultConfigConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InventoryFaultConfigInitParameters struct {
	ConfigpbAttributes []InventoryFaultConfigConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ControllerFaults []ControllerFaultsInitParameters `json:"controllerFaults,omitempty" tf:"controller_faults,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServiceengineFaults []ServiceengineFaultsInitParameters `json:"serviceengineFaults,omitempty" tf:"serviceengine_faults,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	VirtualserviceFaults []VirtualserviceFaultsInitParameters `json:"virtualserviceFaults,omitempty" tf:"virtualservice_faults,omitempty"`
}

type InventoryFaultConfigObservation struct {
	ConfigpbAttributes []InventoryFaultConfigConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ControllerFaults []ControllerFaultsObservation `json:"controllerFaults,omitempty" tf:"controller_faults,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServiceengineFaults []ServiceengineFaultsObservation `json:"serviceengineFaults,omitempty" tf:"serviceengine_faults,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	VirtualserviceFaults []VirtualserviceFaultsObservation `json:"virtualserviceFaults,omitempty" tf:"virtualservice_faults,omitempty"`
}

type InventoryFaultConfigParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []InventoryFaultConfigConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	ControllerFaults []ControllerFaultsParameters `json:"controllerFaults,omitempty" tf:"controller_faults,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceengineFaults []ServiceengineFaultsParameters `json:"serviceengineFaults,omitempty" tf:"serviceengine_faults,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualserviceFaults []VirtualserviceFaultsParameters `json:"virtualserviceFaults,omitempty" tf:"virtualservice_faults,omitempty"`
}

type ServiceengineFaultsInitParameters struct {
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`
}

type ServiceengineFaultsObservation struct {
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`
}

type ServiceengineFaultsParameters struct {

	// +kubebuilder:validation:Optional
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`
}

type VirtualserviceFaultsInitParameters struct {
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`

	PoolServerFaults *string `json:"poolServerFaults,omitempty" tf:"pool_server_faults,omitempty"`

	SSLCertExpiryFaults *string `json:"sslCertExpiryFaults,omitempty" tf:"ssl_cert_expiry_faults,omitempty"`

	SSLCertStatusFaults *string `json:"sslCertStatusFaults,omitempty" tf:"ssl_cert_status_faults,omitempty"`

	ScaleoutFaults *string `json:"scaleoutFaults,omitempty" tf:"scaleout_faults,omitempty"`

	SharedVipFaults *string `json:"sharedVipFaults,omitempty" tf:"shared_vip_faults,omitempty"`
}

type VirtualserviceFaultsObservation struct {
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`

	PoolServerFaults *string `json:"poolServerFaults,omitempty" tf:"pool_server_faults,omitempty"`

	SSLCertExpiryFaults *string `json:"sslCertExpiryFaults,omitempty" tf:"ssl_cert_expiry_faults,omitempty"`

	SSLCertStatusFaults *string `json:"sslCertStatusFaults,omitempty" tf:"ssl_cert_status_faults,omitempty"`

	ScaleoutFaults *string `json:"scaleoutFaults,omitempty" tf:"scaleout_faults,omitempty"`

	SharedVipFaults *string `json:"sharedVipFaults,omitempty" tf:"shared_vip_faults,omitempty"`
}

type VirtualserviceFaultsParameters struct {

	// +kubebuilder:validation:Optional
	DebugFaults *string `json:"debugFaults,omitempty" tf:"debug_faults,omitempty"`

	// +kubebuilder:validation:Optional
	PoolServerFaults *string `json:"poolServerFaults,omitempty" tf:"pool_server_faults,omitempty"`

	// +kubebuilder:validation:Optional
	SSLCertExpiryFaults *string `json:"sslCertExpiryFaults,omitempty" tf:"ssl_cert_expiry_faults,omitempty"`

	// +kubebuilder:validation:Optional
	SSLCertStatusFaults *string `json:"sslCertStatusFaults,omitempty" tf:"ssl_cert_status_faults,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleoutFaults *string `json:"scaleoutFaults,omitempty" tf:"scaleout_faults,omitempty"`

	// +kubebuilder:validation:Optional
	SharedVipFaults *string `json:"sharedVipFaults,omitempty" tf:"shared_vip_faults,omitempty"`
}

// InventoryFaultConfigSpec defines the desired state of InventoryFaultConfig
type InventoryFaultConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InventoryFaultConfigParameters `json:"forProvider"`
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
	InitProvider InventoryFaultConfigInitParameters `json:"initProvider,omitempty"`
}

// InventoryFaultConfigStatus defines the observed state of InventoryFaultConfig.
type InventoryFaultConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InventoryFaultConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InventoryFaultConfig is the Schema for the InventoryFaultConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type InventoryFaultConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InventoryFaultConfigSpec   `json:"spec"`
	Status            InventoryFaultConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InventoryFaultConfigList contains a list of InventoryFaultConfigs
type InventoryFaultConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InventoryFaultConfig `json:"items"`
}

// Repository type metadata.
var (
	InventoryFaultConfig_Kind             = "InventoryFaultConfig"
	InventoryFaultConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InventoryFaultConfig_Kind}.String()
	InventoryFaultConfig_KindAPIVersion   = InventoryFaultConfig_Kind + "." + CRDGroupVersion.String()
	InventoryFaultConfig_GroupVersionKind = CRDGroupVersion.WithKind(InventoryFaultConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&InventoryFaultConfig{}, &InventoryFaultConfigList{})
}
