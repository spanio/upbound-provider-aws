// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CognitoConfigInitParameters struct {

	// The client ID for your Amazon Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// ID for your Amazon Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_pool_id",false)
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolRef *v1.Reference `json:"userPoolRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolSelector *v1.Selector `json:"userPoolSelector,omitempty" tf:"-"`
}

type CognitoConfigObservation struct {

	// The client ID for your Amazon Cognito user pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// ID for your Amazon Cognito user pool.
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`
}

type CognitoConfigParameters struct {

	// The client ID for your Amazon Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// ID for your Amazon Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_pool_id",false)
	// +kubebuilder:validation:Optional
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolRef *v1.Reference `json:"userPoolRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolSelector *v1.Selector `json:"userPoolSelector,omitempty" tf:"-"`
}

type OidcConfigInitParameters struct {

	// The OIDC IdP authorization endpoint used to configure your private workforce.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The client ID for your Amazon Cognito user pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The OIDC IdP issuer used to configure your private workforce.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	// The OIDC IdP logout endpoint used to configure your private workforce.
	LogoutEndpoint *string `json:"logoutEndpoint,omitempty" tf:"logout_endpoint,omitempty"`

	// The OIDC IdP token endpoint used to configure your private workforce.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	// The OIDC IdP user information endpoint used to configure your private workforce.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tf:"user_info_endpoint,omitempty"`
}

type OidcConfigObservation struct {

	// The OIDC IdP authorization endpoint used to configure your private workforce.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The client ID for your Amazon Cognito user pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The OIDC IdP issuer used to configure your private workforce.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	// The OIDC IdP logout endpoint used to configure your private workforce.
	LogoutEndpoint *string `json:"logoutEndpoint,omitempty" tf:"logout_endpoint,omitempty"`

	// The OIDC IdP token endpoint used to configure your private workforce.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	// The OIDC IdP user information endpoint used to configure your private workforce.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tf:"user_info_endpoint,omitempty"`
}

type OidcConfigParameters struct {

	// The OIDC IdP authorization endpoint used to configure your private workforce.
	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// The client ID for your Amazon Cognito user pool.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The OIDC IdP client secret used to configure your private workforce.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The OIDC IdP issuer used to configure your private workforce.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	// +kubebuilder:validation:Optional
	JwksURI *string `json:"jwksUri" tf:"jwks_uri,omitempty"`

	// The OIDC IdP logout endpoint used to configure your private workforce.
	// +kubebuilder:validation:Optional
	LogoutEndpoint *string `json:"logoutEndpoint" tf:"logout_endpoint,omitempty"`

	// The OIDC IdP token endpoint used to configure your private workforce.
	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// The OIDC IdP user information endpoint used to configure your private workforce.
	// +kubebuilder:validation:Optional
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint,omitempty"`
}

type SourceIPConfigInitParameters struct {

	// A list of up to 10 CIDR values.
	// +listType=set
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`
}

type SourceIPConfigObservation struct {

	// A list of up to 10 CIDR values.
	// +listType=set
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`
}

type SourceIPConfigParameters struct {

	// A list of up to 10 CIDR values.
	// +kubebuilder:validation:Optional
	// +listType=set
	Cidrs []*string `json:"cidrs" tf:"cidrs,omitempty"`
}

type WorkforceInitParameters struct {

	// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with oidc_config. see Cognito Config details below.
	CognitoConfig []CognitoConfigInitParameters `json:"cognitoConfig,omitempty" tf:"cognito_config,omitempty"`

	// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with cognito_config. see OIDC Config details below.
	OidcConfig []OidcConfigInitParameters `json:"oidcConfig,omitempty" tf:"oidc_config,omitempty"`

	// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
	SourceIPConfig []SourceIPConfigInitParameters `json:"sourceIpConfig,omitempty" tf:"source_ip_config,omitempty"`

	// configure a workforce using VPC. see Workforce VPC Config details below.
	WorkforceVPCConfig []WorkforceVPCConfigInitParameters `json:"workforceVpcConfig,omitempty" tf:"workforce_vpc_config,omitempty"`
}

type WorkforceObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Workforce.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with oidc_config. see Cognito Config details below.
	CognitoConfig []CognitoConfigObservation `json:"cognitoConfig,omitempty" tf:"cognito_config,omitempty"`

	// The name of the Workforce.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with cognito_config. see OIDC Config details below.
	OidcConfig []OidcConfigObservation `json:"oidcConfig,omitempty" tf:"oidc_config,omitempty"`

	// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
	SourceIPConfig []SourceIPConfigObservation `json:"sourceIpConfig,omitempty" tf:"source_ip_config,omitempty"`

	// The subdomain for your OIDC Identity Provider.
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// configure a workforce using VPC. see Workforce VPC Config details below.
	WorkforceVPCConfig []WorkforceVPCConfigObservation `json:"workforceVpcConfig,omitempty" tf:"workforce_vpc_config,omitempty"`
}

type WorkforceParameters struct {

	// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with oidc_config. see Cognito Config details below.
	// +kubebuilder:validation:Optional
	CognitoConfig []CognitoConfigParameters `json:"cognitoConfig,omitempty" tf:"cognito_config,omitempty"`

	// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with cognito_config. see OIDC Config details below.
	// +kubebuilder:validation:Optional
	OidcConfig []OidcConfigParameters `json:"oidcConfig,omitempty" tf:"oidc_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
	// +kubebuilder:validation:Optional
	SourceIPConfig []SourceIPConfigParameters `json:"sourceIpConfig,omitempty" tf:"source_ip_config,omitempty"`

	// configure a workforce using VPC. see Workforce VPC Config details below.
	// +kubebuilder:validation:Optional
	WorkforceVPCConfig []WorkforceVPCConfigParameters `json:"workforceVpcConfig,omitempty" tf:"workforce_vpc_config,omitempty"`
}

type WorkforceVPCConfigInitParameters struct {

	// The VPC security group IDs. The security groups must be for the same VPC as specified in the subnet.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnets in the VPC that you want to connect.
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// The ID of the VPC that the workforce uses for communication.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type WorkforceVPCConfigObservation struct {

	// The VPC security group IDs. The security groups must be for the same VPC as specified in the subnet.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnets in the VPC that you want to connect.
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// The IDs for the VPC service endpoints of your VPC workforce.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC that the workforce uses for communication.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type WorkforceVPCConfigParameters struct {

	// The VPC security group IDs. The security groups must be for the same VPC as specified in the subnet.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnets in the VPC that you want to connect.
	// +kubebuilder:validation:Optional
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// The ID of the VPC that the workforce uses for communication.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// WorkforceSpec defines the desired state of Workforce
type WorkforceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkforceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WorkforceInitParameters `json:"initProvider,omitempty"`
}

// WorkforceStatus defines the observed state of Workforce.
type WorkforceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkforceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workforce is the Schema for the Workforces API. Provides a SageMaker Workforce resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workforce struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkforceSpec   `json:"spec"`
	Status            WorkforceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkforceList contains a list of Workforces
type WorkforceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workforce `json:"items"`
}

// Repository type metadata.
var (
	Workforce_Kind             = "Workforce"
	Workforce_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workforce_Kind}.String()
	Workforce_KindAPIVersion   = Workforce_Kind + "." + CRDGroupVersion.String()
	Workforce_GroupVersionKind = CRDGroupVersion.WithKind(Workforce_Kind)
)

func init() {
	SchemeBuilder.Register(&Workforce{}, &WorkforceList{})
}
