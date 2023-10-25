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

type DsrProfileInitParameters struct {
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type DsrProfileObservation struct {
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type DsrProfileParameters struct {

	// +kubebuilder:validation:Optional
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	// +kubebuilder:validation:Optional
	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type NetworkProfileConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkProfileConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkProfileConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkProfileInitParameters struct {
	ConfigpbAttributes []NetworkProfileConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ConnectionMirror *string `json:"connectionMirror,omitempty" tf:"connection_mirror,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Markers []NetworkProfileMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Profile []ProfileInitParameters `json:"profile,omitempty" tf:"profile,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkProfileMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkProfileMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkProfileMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NetworkProfileObservation struct {
	ConfigpbAttributes []NetworkProfileConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	ConnectionMirror *string `json:"connectionMirror,omitempty" tf:"connection_mirror,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Markers []NetworkProfileMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Profile []ProfileObservation `json:"profile,omitempty" tf:"profile,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type NetworkProfileParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []NetworkProfileConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionMirror *string `json:"connectionMirror,omitempty" tf:"connection_mirror,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []NetworkProfileMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Profile []ProfileParameters `json:"profile,omitempty" tf:"profile,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ProfileInitParameters struct {
	SctpFastPathProfile []SctpFastPathProfileInitParameters `json:"sctpFastPathProfile,omitempty" tf:"sctp_fast_path_profile,omitempty"`

	SctpProxyProfile []SctpProxyProfileInitParameters `json:"sctpProxyProfile,omitempty" tf:"sctp_proxy_profile,omitempty"`

	TCPFastPathProfile []TCPFastPathProfileInitParameters `json:"tcpFastPathProfile,omitempty" tf:"tcp_fast_path_profile,omitempty"`

	TCPProxyProfile []TCPProxyProfileInitParameters `json:"tcpProxyProfile,omitempty" tf:"tcp_proxy_profile,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UDPFastPathProfile []UDPFastPathProfileInitParameters `json:"udpFastPathProfile,omitempty" tf:"udp_fast_path_profile,omitempty"`

	UDPProxyProfile []UDPProxyProfileInitParameters `json:"udpProxyProfile,omitempty" tf:"udp_proxy_profile,omitempty"`
}

type ProfileObservation struct {
	SctpFastPathProfile []SctpFastPathProfileObservation `json:"sctpFastPathProfile,omitempty" tf:"sctp_fast_path_profile,omitempty"`

	SctpProxyProfile []SctpProxyProfileObservation `json:"sctpProxyProfile,omitempty" tf:"sctp_proxy_profile,omitempty"`

	TCPFastPathProfile []TCPFastPathProfileObservation `json:"tcpFastPathProfile,omitempty" tf:"tcp_fast_path_profile,omitempty"`

	TCPProxyProfile []TCPProxyProfileObservation `json:"tcpProxyProfile,omitempty" tf:"tcp_proxy_profile,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UDPFastPathProfile []UDPFastPathProfileObservation `json:"udpFastPathProfile,omitempty" tf:"udp_fast_path_profile,omitempty"`

	UDPProxyProfile []UDPProxyProfileObservation `json:"udpProxyProfile,omitempty" tf:"udp_proxy_profile,omitempty"`
}

type ProfileParameters struct {

	// +kubebuilder:validation:Optional
	SctpFastPathProfile []SctpFastPathProfileParameters `json:"sctpFastPathProfile,omitempty" tf:"sctp_fast_path_profile,omitempty"`

	// +kubebuilder:validation:Optional
	SctpProxyProfile []SctpProxyProfileParameters `json:"sctpProxyProfile,omitempty" tf:"sctp_proxy_profile,omitempty"`

	// +kubebuilder:validation:Optional
	TCPFastPathProfile []TCPFastPathProfileParameters `json:"tcpFastPathProfile,omitempty" tf:"tcp_fast_path_profile,omitempty"`

	// +kubebuilder:validation:Optional
	TCPProxyProfile []TCPProxyProfileParameters `json:"tcpProxyProfile,omitempty" tf:"tcp_proxy_profile,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UDPFastPathProfile []UDPFastPathProfileParameters `json:"udpFastPathProfile,omitempty" tf:"udp_fast_path_profile,omitempty"`

	// +kubebuilder:validation:Optional
	UDPProxyProfile []UDPProxyProfileParameters `json:"udpProxyProfile,omitempty" tf:"udp_proxy_profile,omitempty"`
}

type SctpFastPathProfileInitParameters struct {
	EnableInitChunkProtection *string `json:"enableInitChunkProtection,omitempty" tf:"enable_init_chunk_protection,omitempty"`

	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
}

type SctpFastPathProfileObservation struct {
	EnableInitChunkProtection *string `json:"enableInitChunkProtection,omitempty" tf:"enable_init_chunk_protection,omitempty"`

	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
}

type SctpFastPathProfileParameters struct {

	// +kubebuilder:validation:Optional
	EnableInitChunkProtection *string `json:"enableInitChunkProtection,omitempty" tf:"enable_init_chunk_protection,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
}

type SctpProxyProfileInitParameters struct {
	CookieExpirationTimeout *string `json:"cookieExpirationTimeout,omitempty" tf:"cookie_expiration_timeout,omitempty"`

	HeartbeatInterval *string `json:"heartbeatInterval,omitempty" tf:"heartbeat_interval,omitempty"`

	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	MaxRetransmissionsAssociation *string `json:"maxRetransmissionsAssociation,omitempty" tf:"max_retransmissions_association,omitempty"`

	MaxRetransmissionsInitChunks *string `json:"maxRetransmissionsInitChunks,omitempty" tf:"max_retransmissions_init_chunks,omitempty"`

	NumberOfStreams *string `json:"numberOfStreams,omitempty" tf:"number_of_streams,omitempty"`

	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	ResetTimeout *string `json:"resetTimeout,omitempty" tf:"reset_timeout,omitempty"`
}

type SctpProxyProfileObservation struct {
	CookieExpirationTimeout *string `json:"cookieExpirationTimeout,omitempty" tf:"cookie_expiration_timeout,omitempty"`

	HeartbeatInterval *string `json:"heartbeatInterval,omitempty" tf:"heartbeat_interval,omitempty"`

	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	MaxRetransmissionsAssociation *string `json:"maxRetransmissionsAssociation,omitempty" tf:"max_retransmissions_association,omitempty"`

	MaxRetransmissionsInitChunks *string `json:"maxRetransmissionsInitChunks,omitempty" tf:"max_retransmissions_init_chunks,omitempty"`

	NumberOfStreams *string `json:"numberOfStreams,omitempty" tf:"number_of_streams,omitempty"`

	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	ResetTimeout *string `json:"resetTimeout,omitempty" tf:"reset_timeout,omitempty"`
}

type SctpProxyProfileParameters struct {

	// +kubebuilder:validation:Optional
	CookieExpirationTimeout *string `json:"cookieExpirationTimeout,omitempty" tf:"cookie_expiration_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	HeartbeatInterval *string `json:"heartbeatInterval,omitempty" tf:"heartbeat_interval,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeout *string `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	MaxRetransmissionsAssociation *string `json:"maxRetransmissionsAssociation,omitempty" tf:"max_retransmissions_association,omitempty"`

	// +kubebuilder:validation:Optional
	MaxRetransmissionsInitChunks *string `json:"maxRetransmissionsInitChunks,omitempty" tf:"max_retransmissions_init_chunks,omitempty"`

	// +kubebuilder:validation:Optional
	NumberOfStreams *string `json:"numberOfStreams,omitempty" tf:"number_of_streams,omitempty"`

	// +kubebuilder:validation:Optional
	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	// +kubebuilder:validation:Optional
	ResetTimeout *string `json:"resetTimeout,omitempty" tf:"reset_timeout,omitempty"`
}

type TCPFastPathProfileInitParameters struct {
	DsrProfile []DsrProfileInitParameters `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	EnableSynProtection *string `json:"enableSynProtection,omitempty" tf:"enable_syn_protection,omitempty"`

	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

type TCPFastPathProfileObservation struct {
	DsrProfile []DsrProfileObservation `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	EnableSynProtection *string `json:"enableSynProtection,omitempty" tf:"enable_syn_protection,omitempty"`

	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

type TCPFastPathProfileParameters struct {

	// +kubebuilder:validation:Optional
	DsrProfile []DsrProfileParameters `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSynProtection *string `json:"enableSynProtection,omitempty" tf:"enable_syn_protection,omitempty"`

	// +kubebuilder:validation:Optional
	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

type TCPProxyProfileInitParameters struct {
	AggressiveCongestionAvoidance *string `json:"aggressiveCongestionAvoidance,omitempty" tf:"aggressive_congestion_avoidance,omitempty"`

	AutoWindowGrowth *string `json:"autoWindowGrowth,omitempty" tf:"auto_window_growth,omitempty"`

	Automatic *string `json:"automatic,omitempty" tf:"automatic,omitempty"`

	CcAlgo *string `json:"ccAlgo,omitempty" tf:"cc_algo,omitempty"`

	CongestionRecoveryScalingFactor *string `json:"congestionRecoveryScalingFactor,omitempty" tf:"congestion_recovery_scaling_factor,omitempty"`

	IPDscp *string `json:"ipDscp,omitempty" tf:"ip_dscp,omitempty"`

	IdleConnectionTimeout *string `json:"idleConnectionTimeout,omitempty" tf:"idle_connection_timeout,omitempty"`

	IdleConnectionType *string `json:"idleConnectionType,omitempty" tf:"idle_connection_type,omitempty"`

	IgnoreTimeWait *string `json:"ignoreTimeWait,omitempty" tf:"ignore_time_wait,omitempty"`

	KeepaliveInHalfcloseState *string `json:"keepaliveInHalfcloseState,omitempty" tf:"keepalive_in_halfclose_state,omitempty"`

	MaxRetransmissions *string `json:"maxRetransmissions,omitempty" tf:"max_retransmissions,omitempty"`

	MaxSegmentSize *string `json:"maxSegmentSize,omitempty" tf:"max_segment_size,omitempty"`

	MaxSynRetransmissions *string `json:"maxSynRetransmissions,omitempty" tf:"max_syn_retransmissions,omitempty"`

	MinRexmtTimeout *string `json:"minRexmtTimeout,omitempty" tf:"min_rexmt_timeout,omitempty"`

	NaglesAlgorithm *string `json:"naglesAlgorithm,omitempty" tf:"nagles_algorithm,omitempty"`

	ReassemblyQueueSize *string `json:"reassemblyQueueSize,omitempty" tf:"reassembly_queue_size,omitempty"`

	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	ReorderThreshold *string `json:"reorderThreshold,omitempty" tf:"reorder_threshold,omitempty"`

	SlowStartScalingFactor *string `json:"slowStartScalingFactor,omitempty" tf:"slow_start_scaling_factor,omitempty"`

	TimeWaitDelay *string `json:"timeWaitDelay,omitempty" tf:"time_wait_delay,omitempty"`

	UseInterfaceMtu *string `json:"useInterfaceMtu,omitempty" tf:"use_interface_mtu,omitempty"`
}

type TCPProxyProfileObservation struct {
	AggressiveCongestionAvoidance *string `json:"aggressiveCongestionAvoidance,omitempty" tf:"aggressive_congestion_avoidance,omitempty"`

	AutoWindowGrowth *string `json:"autoWindowGrowth,omitempty" tf:"auto_window_growth,omitempty"`

	Automatic *string `json:"automatic,omitempty" tf:"automatic,omitempty"`

	CcAlgo *string `json:"ccAlgo,omitempty" tf:"cc_algo,omitempty"`

	CongestionRecoveryScalingFactor *string `json:"congestionRecoveryScalingFactor,omitempty" tf:"congestion_recovery_scaling_factor,omitempty"`

	IPDscp *string `json:"ipDscp,omitempty" tf:"ip_dscp,omitempty"`

	IdleConnectionTimeout *string `json:"idleConnectionTimeout,omitempty" tf:"idle_connection_timeout,omitempty"`

	IdleConnectionType *string `json:"idleConnectionType,omitempty" tf:"idle_connection_type,omitempty"`

	IgnoreTimeWait *string `json:"ignoreTimeWait,omitempty" tf:"ignore_time_wait,omitempty"`

	KeepaliveInHalfcloseState *string `json:"keepaliveInHalfcloseState,omitempty" tf:"keepalive_in_halfclose_state,omitempty"`

	MaxRetransmissions *string `json:"maxRetransmissions,omitempty" tf:"max_retransmissions,omitempty"`

	MaxSegmentSize *string `json:"maxSegmentSize,omitempty" tf:"max_segment_size,omitempty"`

	MaxSynRetransmissions *string `json:"maxSynRetransmissions,omitempty" tf:"max_syn_retransmissions,omitempty"`

	MinRexmtTimeout *string `json:"minRexmtTimeout,omitempty" tf:"min_rexmt_timeout,omitempty"`

	NaglesAlgorithm *string `json:"naglesAlgorithm,omitempty" tf:"nagles_algorithm,omitempty"`

	ReassemblyQueueSize *string `json:"reassemblyQueueSize,omitempty" tf:"reassembly_queue_size,omitempty"`

	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	ReorderThreshold *string `json:"reorderThreshold,omitempty" tf:"reorder_threshold,omitempty"`

	SlowStartScalingFactor *string `json:"slowStartScalingFactor,omitempty" tf:"slow_start_scaling_factor,omitempty"`

	TimeWaitDelay *string `json:"timeWaitDelay,omitempty" tf:"time_wait_delay,omitempty"`

	UseInterfaceMtu *string `json:"useInterfaceMtu,omitempty" tf:"use_interface_mtu,omitempty"`
}

type TCPProxyProfileParameters struct {

	// +kubebuilder:validation:Optional
	AggressiveCongestionAvoidance *string `json:"aggressiveCongestionAvoidance,omitempty" tf:"aggressive_congestion_avoidance,omitempty"`

	// +kubebuilder:validation:Optional
	AutoWindowGrowth *string `json:"autoWindowGrowth,omitempty" tf:"auto_window_growth,omitempty"`

	// +kubebuilder:validation:Optional
	Automatic *string `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// +kubebuilder:validation:Optional
	CcAlgo *string `json:"ccAlgo,omitempty" tf:"cc_algo,omitempty"`

	// +kubebuilder:validation:Optional
	CongestionRecoveryScalingFactor *string `json:"congestionRecoveryScalingFactor,omitempty" tf:"congestion_recovery_scaling_factor,omitempty"`

	// +kubebuilder:validation:Optional
	IPDscp *string `json:"ipDscp,omitempty" tf:"ip_dscp,omitempty"`

	// +kubebuilder:validation:Optional
	IdleConnectionTimeout *string `json:"idleConnectionTimeout,omitempty" tf:"idle_connection_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	IdleConnectionType *string `json:"idleConnectionType,omitempty" tf:"idle_connection_type,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreTimeWait *string `json:"ignoreTimeWait,omitempty" tf:"ignore_time_wait,omitempty"`

	// +kubebuilder:validation:Optional
	KeepaliveInHalfcloseState *string `json:"keepaliveInHalfcloseState,omitempty" tf:"keepalive_in_halfclose_state,omitempty"`

	// +kubebuilder:validation:Optional
	MaxRetransmissions *string `json:"maxRetransmissions,omitempty" tf:"max_retransmissions,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSegmentSize *string `json:"maxSegmentSize,omitempty" tf:"max_segment_size,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSynRetransmissions *string `json:"maxSynRetransmissions,omitempty" tf:"max_syn_retransmissions,omitempty"`

	// +kubebuilder:validation:Optional
	MinRexmtTimeout *string `json:"minRexmtTimeout,omitempty" tf:"min_rexmt_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	NaglesAlgorithm *string `json:"naglesAlgorithm,omitempty" tf:"nagles_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	ReassemblyQueueSize *string `json:"reassemblyQueueSize,omitempty" tf:"reassembly_queue_size,omitempty"`

	// +kubebuilder:validation:Optional
	ReceiveWindow *string `json:"receiveWindow,omitempty" tf:"receive_window,omitempty"`

	// +kubebuilder:validation:Optional
	ReorderThreshold *string `json:"reorderThreshold,omitempty" tf:"reorder_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	SlowStartScalingFactor *string `json:"slowStartScalingFactor,omitempty" tf:"slow_start_scaling_factor,omitempty"`

	// +kubebuilder:validation:Optional
	TimeWaitDelay *string `json:"timeWaitDelay,omitempty" tf:"time_wait_delay,omitempty"`

	// +kubebuilder:validation:Optional
	UseInterfaceMtu *string `json:"useInterfaceMtu,omitempty" tf:"use_interface_mtu,omitempty"`
}

type UDPFastPathProfileDsrProfileInitParameters struct {
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type UDPFastPathProfileDsrProfileObservation struct {
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type UDPFastPathProfileDsrProfileParameters struct {

	// +kubebuilder:validation:Optional
	DsrEncapType *string `json:"dsrEncapType,omitempty" tf:"dsr_encap_type,omitempty"`

	// +kubebuilder:validation:Optional
	DsrType *string `json:"dsrType,omitempty" tf:"dsr_type,omitempty"`
}

type UDPFastPathProfileInitParameters struct {
	DsrProfile []UDPFastPathProfileDsrProfileInitParameters `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	PerPktLoadbalance *string `json:"perPktLoadbalance,omitempty" tf:"per_pkt_loadbalance,omitempty"`

	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`

	Snat *string `json:"snat,omitempty" tf:"snat,omitempty"`
}

type UDPFastPathProfileObservation struct {
	DsrProfile []UDPFastPathProfileDsrProfileObservation `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	PerPktLoadbalance *string `json:"perPktLoadbalance,omitempty" tf:"per_pkt_loadbalance,omitempty"`

	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`

	Snat *string `json:"snat,omitempty" tf:"snat,omitempty"`
}

type UDPFastPathProfileParameters struct {

	// +kubebuilder:validation:Optional
	DsrProfile []UDPFastPathProfileDsrProfileParameters `json:"dsrProfile,omitempty" tf:"dsr_profile,omitempty"`

	// +kubebuilder:validation:Optional
	PerPktLoadbalance *string `json:"perPktLoadbalance,omitempty" tf:"per_pkt_loadbalance,omitempty"`

	// +kubebuilder:validation:Optional
	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Snat *string `json:"snat,omitempty" tf:"snat,omitempty"`
}

type UDPProxyProfileInitParameters struct {
	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

type UDPProxyProfileObservation struct {
	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

type UDPProxyProfileParameters struct {

	// +kubebuilder:validation:Optional
	SessionIdleTimeout *string `json:"sessionIdleTimeout,omitempty" tf:"session_idle_timeout,omitempty"`
}

// NetworkProfileSpec defines the desired state of NetworkProfile
type NetworkProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkProfileParameters `json:"forProvider"`
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
	InitProvider NetworkProfileInitParameters `json:"initProvider,omitempty"`
}

// NetworkProfileStatus defines the observed state of NetworkProfile.
type NetworkProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkProfile is the Schema for the NetworkProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type NetworkProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profile) || has(self.initProvider.profile)",message="profile is a required parameter"
	Spec   NetworkProfileSpec   `json:"spec"`
	Status NetworkProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkProfileList contains a list of NetworkProfiles
type NetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkProfile `json:"items"`
}

// Repository type metadata.
var (
	NetworkProfile_Kind             = "NetworkProfile"
	NetworkProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkProfile_Kind}.String()
	NetworkProfile_KindAPIVersion   = NetworkProfile_Kind + "." + CRDGroupVersion.String()
	NetworkProfile_GroupVersionKind = CRDGroupVersion.WithKind(NetworkProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkProfile{}, &NetworkProfileList{})
}