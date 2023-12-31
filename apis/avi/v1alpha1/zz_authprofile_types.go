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

type AuthProfileConfigpbAttributesInitParameters struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AuthProfileConfigpbAttributesObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AuthProfileConfigpbAttributesParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AuthProfileInitParameters struct {
	ConfigpbAttributes []AuthProfileConfigpbAttributesInitParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	HTTP []HTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	JwtProfileRef *string `json:"jwtProfileRef,omitempty" tf:"jwt_profile_ref,omitempty"`

	Ldap []LdapInitParameters `json:"ldap,omitempty" tf:"ldap,omitempty"`

	Markers []AuthProfileMarkersInitParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OauthProfile []OauthProfileInitParameters `json:"oauthProfile,omitempty" tf:"oauth_profile,omitempty"`

	PaAgentRef *string `json:"paAgentRef,omitempty" tf:"pa_agent_ref,omitempty"`

	SAML []SAMLInitParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	TacacsPlus []TacacsPlusInitParameters `json:"tacacsPlus,omitempty" tf:"tacacs_plus,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type AuthProfileMarkersInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AuthProfileMarkersObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AuthProfileMarkersParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AuthProfileObservation struct {
	ConfigpbAttributes []AuthProfileConfigpbAttributesObservation `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	HTTP []HTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	JwtProfileRef *string `json:"jwtProfileRef,omitempty" tf:"jwt_profile_ref,omitempty"`

	Ldap []LdapObservation `json:"ldap,omitempty" tf:"ldap,omitempty"`

	Markers []AuthProfileMarkersObservation `json:"markers,omitempty" tf:"markers,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OauthProfile []OauthProfileObservation `json:"oauthProfile,omitempty" tf:"oauth_profile,omitempty"`

	PaAgentRef *string `json:"paAgentRef,omitempty" tf:"pa_agent_ref,omitempty"`

	SAML []SAMLObservation `json:"saml,omitempty" tf:"saml,omitempty"`

	TacacsPlus []TacacsPlusObservation `json:"tacacsPlus,omitempty" tf:"tacacs_plus,omitempty"`

	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type AuthProfileParameters struct {

	// +kubebuilder:validation:Optional
	ConfigpbAttributes []AuthProfileConfigpbAttributesParameters `json:"configpbAttributes,omitempty" tf:"configpb_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// +kubebuilder:validation:Optional
	JwtProfileRef *string `json:"jwtProfileRef,omitempty" tf:"jwt_profile_ref,omitempty"`

	// +kubebuilder:validation:Optional
	Ldap []LdapParameters `json:"ldap,omitempty" tf:"ldap,omitempty"`

	// +kubebuilder:validation:Optional
	Markers []AuthProfileMarkersParameters `json:"markers,omitempty" tf:"markers,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OauthProfile []OauthProfileParameters `json:"oauthProfile,omitempty" tf:"oauth_profile,omitempty"`

	// +kubebuilder:validation:Optional
	PaAgentRef *string `json:"paAgentRef,omitempty" tf:"pa_agent_ref,omitempty"`

	// +kubebuilder:validation:Optional
	SAML []SAMLParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	// +kubebuilder:validation:Optional
	TacacsPlus []TacacsPlusParameters `json:"tacacsPlus,omitempty" tf:"tacacs_plus,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *string `json:"tenantRef,omitempty" tf:"tenant_ref,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type AuthorizationAttrsInitParameters struct {
	Mandatory *string `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AuthorizationAttrsObservation struct {
	Mandatory *string `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AuthorizationAttrsParameters struct {

	// +kubebuilder:validation:Optional
	Mandatory *string `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HTTPInitParameters struct {
	CacheExpirationTime *string `json:"cacheExpirationTime,omitempty" tf:"cache_expiration_time,omitempty"`

	RequestHeader *string `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	RequireUserGroups []*string `json:"requireUserGroups,omitempty" tf:"require_user_groups,omitempty"`
}

type HTTPObservation struct {
	CacheExpirationTime *string `json:"cacheExpirationTime,omitempty" tf:"cache_expiration_time,omitempty"`

	RequestHeader *string `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	RequireUserGroups []*string `json:"requireUserGroups,omitempty" tf:"require_user_groups,omitempty"`
}

type HTTPParameters struct {

	// +kubebuilder:validation:Optional
	CacheExpirationTime *string `json:"cacheExpirationTime,omitempty" tf:"cache_expiration_time,omitempty"`

	// +kubebuilder:validation:Optional
	RequestHeader *string `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// +kubebuilder:validation:Optional
	RequireUserGroups []*string `json:"requireUserGroups,omitempty" tf:"require_user_groups,omitempty"`
}

type IdpInitParameters struct {
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type IdpObservation struct {
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type IdpParameters struct {

	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type LdapInitParameters struct {
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	BindAsAdministrator *string `json:"bindAsAdministrator,omitempty" tf:"bind_as_administrator,omitempty"`

	EmailAttribute *string `json:"emailAttribute,omitempty" tf:"email_attribute,omitempty"`

	FullNameAttribute *string `json:"fullNameAttribute,omitempty" tf:"full_name_attribute,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	SecurityMode *string `json:"securityMode,omitempty" tf:"security_mode,omitempty"`

	Server []*string `json:"server,omitempty" tf:"server,omitempty"`

	Settings []SettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	UserBind []UserBindInitParameters `json:"userBind,omitempty" tf:"user_bind,omitempty"`
}

type LdapObservation struct {
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	BindAsAdministrator *string `json:"bindAsAdministrator,omitempty" tf:"bind_as_administrator,omitempty"`

	EmailAttribute *string `json:"emailAttribute,omitempty" tf:"email_attribute,omitempty"`

	FullNameAttribute *string `json:"fullNameAttribute,omitempty" tf:"full_name_attribute,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	SecurityMode *string `json:"securityMode,omitempty" tf:"security_mode,omitempty"`

	Server []*string `json:"server,omitempty" tf:"server,omitempty"`

	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	UserBind []UserBindObservation `json:"userBind,omitempty" tf:"user_bind,omitempty"`
}

type LdapParameters struct {

	// +kubebuilder:validation:Optional
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// +kubebuilder:validation:Optional
	BindAsAdministrator *string `json:"bindAsAdministrator,omitempty" tf:"bind_as_administrator,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAttribute *string `json:"emailAttribute,omitempty" tf:"email_attribute,omitempty"`

	// +kubebuilder:validation:Optional
	FullNameAttribute *string `json:"fullNameAttribute,omitempty" tf:"full_name_attribute,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityMode *string `json:"securityMode,omitempty" tf:"security_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Server []*string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// +kubebuilder:validation:Optional
	UserBind []UserBindParameters `json:"userBind,omitempty" tf:"user_bind,omitempty"`
}

type OauthProfileInitParameters struct {
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	EndSessionEndpoint *string `json:"endSessionEndpoint,omitempty" tf:"end_session_endpoint,omitempty"`

	IntrospectionEndpoint *string `json:"introspectionEndpoint,omitempty" tf:"introspection_endpoint,omitempty"`

	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	JwksTimeout *string `json:"jwksTimeout,omitempty" tf:"jwks_timeout,omitempty"`

	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	OauthRespBufferSz *string `json:"oauthRespBufferSz,omitempty" tf:"oauth_resp_buffer_sz,omitempty"`

	PoolRef *string `json:"poolRef,omitempty" tf:"pool_ref,omitempty"`

	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" tf:"userinfo_endpoint,omitempty"`
}

type OauthProfileObservation struct {
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	EndSessionEndpoint *string `json:"endSessionEndpoint,omitempty" tf:"end_session_endpoint,omitempty"`

	IntrospectionEndpoint *string `json:"introspectionEndpoint,omitempty" tf:"introspection_endpoint,omitempty"`

	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	JwksTimeout *string `json:"jwksTimeout,omitempty" tf:"jwks_timeout,omitempty"`

	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	OauthRespBufferSz *string `json:"oauthRespBufferSz,omitempty" tf:"oauth_resp_buffer_sz,omitempty"`

	PoolRef *string `json:"poolRef,omitempty" tf:"pool_ref,omitempty"`

	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" tf:"userinfo_endpoint,omitempty"`
}

type OauthProfileParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	EndSessionEndpoint *string `json:"endSessionEndpoint,omitempty" tf:"end_session_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	IntrospectionEndpoint *string `json:"introspectionEndpoint,omitempty" tf:"introspection_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// +kubebuilder:validation:Optional
	JwksTimeout *string `json:"jwksTimeout,omitempty" tf:"jwks_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	// +kubebuilder:validation:Optional
	OauthRespBufferSz *string `json:"oauthRespBufferSz,omitempty" tf:"oauth_resp_buffer_sz,omitempty"`

	// +kubebuilder:validation:Optional
	PoolRef *string `json:"poolRef" tf:"pool_ref,omitempty"`

	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" tf:"userinfo_endpoint,omitempty"`
}

type SAMLInitParameters struct {
	Idp []IdpInitParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	Sp []SpInitParameters `json:"sp,omitempty" tf:"sp,omitempty"`
}

type SAMLObservation struct {
	Idp []IdpObservation `json:"idp,omitempty" tf:"idp,omitempty"`

	Sp []SpObservation `json:"sp,omitempty" tf:"sp,omitempty"`
}

type SAMLParameters struct {

	// +kubebuilder:validation:Optional
	Idp []IdpParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	// +kubebuilder:validation:Optional
	Sp []SpParameters `json:"sp" tf:"sp,omitempty"`
}

type SettingsInitParameters struct {
	AdminBindDn *string `json:"adminBindDn,omitempty" tf:"admin_bind_dn,omitempty"`

	GroupFilter *string `json:"groupFilter,omitempty" tf:"group_filter,omitempty"`

	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	GroupMemberIsFullDn *string `json:"groupMemberIsFullDn,omitempty" tf:"group_member_is_full_dn,omitempty"`

	GroupSearchDn *string `json:"groupSearchDn,omitempty" tf:"group_search_dn,omitempty"`

	GroupSearchScope *string `json:"groupSearchScope,omitempty" tf:"group_search_scope,omitempty"`

	IgnoreReferrals *string `json:"ignoreReferrals,omitempty" tf:"ignore_referrals,omitempty"`

	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	UserSearchDn *string `json:"userSearchDn,omitempty" tf:"user_search_dn,omitempty"`

	UserSearchScope *string `json:"userSearchScope,omitempty" tf:"user_search_scope,omitempty"`
}

type SettingsObservation struct {
	AdminBindDn *string `json:"adminBindDn,omitempty" tf:"admin_bind_dn,omitempty"`

	GroupFilter *string `json:"groupFilter,omitempty" tf:"group_filter,omitempty"`

	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	GroupMemberIsFullDn *string `json:"groupMemberIsFullDn,omitempty" tf:"group_member_is_full_dn,omitempty"`

	GroupSearchDn *string `json:"groupSearchDn,omitempty" tf:"group_search_dn,omitempty"`

	GroupSearchScope *string `json:"groupSearchScope,omitempty" tf:"group_search_scope,omitempty"`

	IgnoreReferrals *string `json:"ignoreReferrals,omitempty" tf:"ignore_referrals,omitempty"`

	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	UserSearchDn *string `json:"userSearchDn,omitempty" tf:"user_search_dn,omitempty"`

	UserSearchScope *string `json:"userSearchScope,omitempty" tf:"user_search_scope,omitempty"`
}

type SettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdminBindDn *string `json:"adminBindDn,omitempty" tf:"admin_bind_dn,omitempty"`

	// +kubebuilder:validation:Optional
	GroupFilter *string `json:"groupFilter,omitempty" tf:"group_filter,omitempty"`

	// +kubebuilder:validation:Optional
	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	// +kubebuilder:validation:Optional
	GroupMemberIsFullDn *string `json:"groupMemberIsFullDn,omitempty" tf:"group_member_is_full_dn,omitempty"`

	// +kubebuilder:validation:Optional
	GroupSearchDn *string `json:"groupSearchDn,omitempty" tf:"group_search_dn,omitempty"`

	// +kubebuilder:validation:Optional
	GroupSearchScope *string `json:"groupSearchScope,omitempty" tf:"group_search_scope,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreReferrals *string `json:"ignoreReferrals,omitempty" tf:"ignore_referrals,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	// +kubebuilder:validation:Optional
	UserSearchDn *string `json:"userSearchDn,omitempty" tf:"user_search_dn,omitempty"`

	// +kubebuilder:validation:Optional
	UserSearchScope *string `json:"userSearchScope,omitempty" tf:"user_search_scope,omitempty"`
}

type SpInitParameters struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	OrgDisplayName *string `json:"orgDisplayName,omitempty" tf:"org_display_name,omitempty"`

	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	OrgURL *string `json:"orgUrl,omitempty" tf:"org_url,omitempty"`

	SAMLEntityType *string `json:"samlEntityType,omitempty" tf:"saml_entity_type,omitempty"`

	SpNodes []SpNodesInitParameters `json:"spNodes,omitempty" tf:"sp_nodes,omitempty"`

	TechContactEmail *string `json:"techContactEmail,omitempty" tf:"tech_contact_email,omitempty"`

	TechContactName *string `json:"techContactName,omitempty" tf:"tech_contact_name,omitempty"`
}

type SpNodesInitParameters struct {
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	SigningSSLKeyAndCertificateRef *string `json:"signingSslKeyAndCertificateRef,omitempty" tf:"signing_ssl_key_and_certificate_ref,omitempty"`

	SingleSignonURL *string `json:"singleSignonUrl,omitempty" tf:"single_signon_url,omitempty"`
}

type SpNodesObservation struct {
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	SigningSSLKeyAndCertificateRef *string `json:"signingSslKeyAndCertificateRef,omitempty" tf:"signing_ssl_key_and_certificate_ref,omitempty"`

	SingleSignonURL *string `json:"singleSignonUrl,omitempty" tf:"single_signon_url,omitempty"`
}

type SpNodesParameters struct {

	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SigningSSLKeyAndCertificateRef *string `json:"signingSslKeyAndCertificateRef,omitempty" tf:"signing_ssl_key_and_certificate_ref,omitempty"`

	// +kubebuilder:validation:Optional
	SingleSignonURL *string `json:"singleSignonUrl,omitempty" tf:"single_signon_url,omitempty"`
}

type SpObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	OrgDisplayName *string `json:"orgDisplayName,omitempty" tf:"org_display_name,omitempty"`

	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	OrgURL *string `json:"orgUrl,omitempty" tf:"org_url,omitempty"`

	SAMLEntityType *string `json:"samlEntityType,omitempty" tf:"saml_entity_type,omitempty"`

	SpNodes []SpNodesObservation `json:"spNodes,omitempty" tf:"sp_nodes,omitempty"`

	TechContactEmail *string `json:"techContactEmail,omitempty" tf:"tech_contact_email,omitempty"`

	TechContactName *string `json:"techContactName,omitempty" tf:"tech_contact_name,omitempty"`
}

type SpParameters struct {

	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	OrgDisplayName *string `json:"orgDisplayName,omitempty" tf:"org_display_name,omitempty"`

	// +kubebuilder:validation:Optional
	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	// +kubebuilder:validation:Optional
	OrgURL *string `json:"orgUrl,omitempty" tf:"org_url,omitempty"`

	// +kubebuilder:validation:Optional
	SAMLEntityType *string `json:"samlEntityType,omitempty" tf:"saml_entity_type,omitempty"`

	// +kubebuilder:validation:Optional
	SpNodes []SpNodesParameters `json:"spNodes,omitempty" tf:"sp_nodes,omitempty"`

	// +kubebuilder:validation:Optional
	TechContactEmail *string `json:"techContactEmail,omitempty" tf:"tech_contact_email,omitempty"`

	// +kubebuilder:validation:Optional
	TechContactName *string `json:"techContactName,omitempty" tf:"tech_contact_name,omitempty"`
}

type TacacsPlusInitParameters struct {
	AuthorizationAttrs []AuthorizationAttrsInitParameters `json:"authorizationAttrs,omitempty" tf:"authorization_attrs,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Server []*string `json:"server,omitempty" tf:"server,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type TacacsPlusObservation struct {
	AuthorizationAttrs []AuthorizationAttrsObservation `json:"authorizationAttrs,omitempty" tf:"authorization_attrs,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Server []*string `json:"server,omitempty" tf:"server,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type TacacsPlusParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizationAttrs []AuthorizationAttrsParameters `json:"authorizationAttrs,omitempty" tf:"authorization_attrs,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Server []*string `json:"server" tf:"server,omitempty"`

	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type UserBindInitParameters struct {
	DnTemplate *string `json:"dnTemplate,omitempty" tf:"dn_template,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`
}

type UserBindObservation struct {
	DnTemplate *string `json:"dnTemplate,omitempty" tf:"dn_template,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`
}

type UserBindParameters struct {

	// +kubebuilder:validation:Optional
	DnTemplate *string `json:"dnTemplate,omitempty" tf:"dn_template,omitempty"`

	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// +kubebuilder:validation:Optional
	UserAttributes []*string `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`
}

// AuthProfileSpec defines the desired state of AuthProfile
type AuthProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthProfileParameters `json:"forProvider"`
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
	InitProvider AuthProfileInitParameters `json:"initProvider,omitempty"`
}

// AuthProfileStatus defines the observed state of AuthProfile.
type AuthProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthProfile is the Schema for the AuthProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,avi}
type AuthProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   AuthProfileSpec   `json:"spec"`
	Status AuthProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthProfileList contains a list of AuthProfiles
type AuthProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthProfile `json:"items"`
}

// Repository type metadata.
var (
	AuthProfile_Kind             = "AuthProfile"
	AuthProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthProfile_Kind}.String()
	AuthProfile_KindAPIVersion   = AuthProfile_Kind + "." + CRDGroupVersion.String()
	AuthProfile_GroupVersionKind = CRDGroupVersion.WithKind(AuthProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthProfile{}, &AuthProfileList{})
}
