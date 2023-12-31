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

type EscrowInfosInitParameters struct {
	LastUpdated *string `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`

	ServiceCores *string `json:"serviceCores,omitempty" tf:"service_cores,omitempty"`

	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type EscrowInfosObservation struct {
	LastUpdated *string `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`

	ServiceCores *string `json:"serviceCores,omitempty" tf:"service_cores,omitempty"`

	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type EscrowInfosParameters struct {

	// +kubebuilder:validation:Optional
	LastUpdated *string `json:"lastUpdated" tf:"last_updated,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceCores *string `json:"serviceCores" tf:"service_cores,omitempty"`

	// +kubebuilder:validation:Optional
	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier" tf:"tier,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type LicenseLedgerDetailsInitParameters struct {
	EscrowInfos []EscrowInfosInitParameters `json:"escrowInfos,omitempty" tf:"escrow_infos,omitempty"`

	SeInfos []SeInfosInitParameters `json:"seInfos,omitempty" tf:"se_infos,omitempty"`

	TierUsages []TierUsagesInitParameters `json:"tierUsages,omitempty" tf:"tier_usages,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type LicenseLedgerDetailsObservation struct {
	EscrowInfos []EscrowInfosObservation `json:"escrowInfos,omitempty" tf:"escrow_infos,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SeInfos []SeInfosObservation `json:"seInfos,omitempty" tf:"se_infos,omitempty"`

	TierUsages []TierUsagesObservation `json:"tierUsages,omitempty" tf:"tier_usages,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type LicenseLedgerDetailsParameters struct {

	// +kubebuilder:validation:Optional
	EscrowInfos []EscrowInfosParameters `json:"escrowInfos,omitempty" tf:"escrow_infos,omitempty"`

	// +kubebuilder:validation:Optional
	SeInfos []SeInfosParameters `json:"seInfos,omitempty" tf:"se_infos,omitempty"`

	// +kubebuilder:validation:Optional
	TierUsages []TierUsagesParameters `json:"tierUsages,omitempty" tf:"tier_usages,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SeInfosInitParameters struct {
	LastUpdated *string `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`

	ServiceCores *string `json:"serviceCores,omitempty" tf:"service_cores,omitempty"`

	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SeInfosObservation struct {
	LastUpdated *string `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`

	ServiceCores *string `json:"serviceCores,omitempty" tf:"service_cores,omitempty"`

	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SeInfosParameters struct {

	// +kubebuilder:validation:Optional
	LastUpdated *string `json:"lastUpdated" tf:"last_updated,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceCores *string `json:"serviceCores" tf:"service_cores,omitempty"`

	// +kubebuilder:validation:Optional
	TenantUUID *string `json:"tenantUuid,omitempty" tf:"tenant_uuid,omitempty"`

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier" tf:"tier,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type TierUsagesInitParameters struct {
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	Usage []UsageInitParameters `json:"usage,omitempty" tf:"usage,omitempty"`
}

type TierUsagesObservation struct {
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	Usage []UsageObservation `json:"usage,omitempty" tf:"usage,omitempty"`
}

type TierUsagesParameters struct {

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// +kubebuilder:validation:Optional
	Usage []UsageParameters `json:"usage,omitempty" tf:"usage,omitempty"`
}

type UsageInitParameters struct {
	Available *string `json:"available,omitempty" tf:"available,omitempty"`

	Consumed *string `json:"consumed,omitempty" tf:"consumed,omitempty"`

	Escrow *string `json:"escrow,omitempty" tf:"escrow,omitempty"`

	Remaining *string `json:"remaining,omitempty" tf:"remaining,omitempty"`
}

type UsageObservation struct {
	Available *string `json:"available,omitempty" tf:"available,omitempty"`

	Consumed *string `json:"consumed,omitempty" tf:"consumed,omitempty"`

	Escrow *string `json:"escrow,omitempty" tf:"escrow,omitempty"`

	Remaining *string `json:"remaining,omitempty" tf:"remaining,omitempty"`
}

type UsageParameters struct {

	// +kubebuilder:validation:Optional
	Available *string `json:"available,omitempty" tf:"available,omitempty"`

	// +kubebuilder:validation:Optional
	Consumed *string `json:"consumed,omitempty" tf:"consumed,omitempty"`

	// +kubebuilder:validation:Optional
	Escrow *string `json:"escrow,omitempty" tf:"escrow,omitempty"`

	// +kubebuilder:validation:Optional
	Remaining *string `json:"remaining,omitempty" tf:"remaining,omitempty"`
}

// LicenseLedgerDetailsSpec defines the desired state of LicenseLedgerDetails
type LicenseLedgerDetailsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseLedgerDetailsParameters `json:"forProvider"`
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
	InitProvider LicenseLedgerDetailsInitParameters `json:"initProvider,omitempty"`
}

// LicenseLedgerDetailsStatus defines the observed state of LicenseLedgerDetails.
type LicenseLedgerDetailsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseLedgerDetailsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseLedgerDetails is the Schema for the LicenseLedgerDetailss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type LicenseLedgerDetails struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseLedgerDetailsSpec   `json:"spec"`
	Status            LicenseLedgerDetailsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseLedgerDetailsList contains a list of LicenseLedgerDetailss
type LicenseLedgerDetailsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseLedgerDetails `json:"items"`
}

// Repository type metadata.
var (
	LicenseLedgerDetails_Kind             = "LicenseLedgerDetails"
	LicenseLedgerDetails_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseLedgerDetails_Kind}.String()
	LicenseLedgerDetails_KindAPIVersion   = LicenseLedgerDetails_Kind + "." + CRDGroupVersion.String()
	LicenseLedgerDetails_GroupVersionKind = CRDGroupVersion.WithKind(LicenseLedgerDetails_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseLedgerDetails{}, &LicenseLedgerDetailsList{})
}
