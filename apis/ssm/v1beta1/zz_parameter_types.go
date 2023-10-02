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

type ParameterInitParameters_2 struct {

	// Regular expression used to validate the parameter value.
	AllowedPattern *string `json:"allowedPattern,omitempty" tf:"allowed_pattern,omitempty"`

	// ARN of the parameter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Data type of the parameter. Valid values: text, aws:ssm:integration and aws:ec2:image for AMI format, see the Native parameter support for Amazon Machine Image IDs.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Description of the parameter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Value of the parameter. This argument is not valid with a type of SecureString.
	InsecureValue *string `json:"insecureValue,omitempty" tf:"insecure_value,omitempty"`

	// KMS key ID or ARN for encrypting a SecureString.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Overwrite an existing parameter.
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are Standard, Advanced, and Intelligent-Tiering. Downgrading an Advanced tier parameter to Standard will recreate the resource. For more information on parameter tiers, see the AWS SSM Parameter tier comparison and guide.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// Type of the parameter. Valid types are String, StringList and SecureString.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParameterObservation_2 struct {

	// Regular expression used to validate the parameter value.
	AllowedPattern *string `json:"allowedPattern,omitempty" tf:"allowed_pattern,omitempty"`

	// ARN of the parameter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Data type of the parameter. Valid values: text, aws:ssm:integration and aws:ec2:image for AMI format, see the Native parameter support for Amazon Machine Image IDs.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Description of the parameter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Value of the parameter. This argument is not valid with a type of SecureString.
	InsecureValue *string `json:"insecureValue,omitempty" tf:"insecure_value,omitempty"`

	// KMS key ID or ARN for encrypting a SecureString.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Overwrite an existing parameter.
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are Standard, Advanced, and Intelligent-Tiering. Downgrading an Advanced tier parameter to Standard will recreate the resource. For more information on parameter tiers, see the AWS SSM Parameter tier comparison and guide.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// Type of the parameter. Valid types are String, StringList and SecureString.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Version of the parameter.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ParameterParameters_2 struct {

	// Regular expression used to validate the parameter value.
	// +kubebuilder:validation:Optional
	AllowedPattern *string `json:"allowedPattern,omitempty" tf:"allowed_pattern,omitempty"`

	// ARN of the parameter.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Data type of the parameter. Valid values: text, aws:ssm:integration and aws:ec2:image for AMI format, see the Native parameter support for Amazon Machine Image IDs.
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Description of the parameter.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Value of the parameter. This argument is not valid with a type of SecureString.
	// +kubebuilder:validation:Optional
	InsecureValue *string `json:"insecureValue,omitempty" tf:"insecure_value,omitempty"`

	// KMS key ID or ARN for encrypting a SecureString.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Overwrite an existing parameter.
	// +kubebuilder:validation:Optional
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are Standard, Advanced, and Intelligent-Tiering. Downgrading an Advanced tier parameter to Standard will recreate the resource. For more information on parameter tiers, see the AWS SSM Parameter tier comparison and guide.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// Type of the parameter. Valid types are String, StringList and SecureString.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value of the parameter.15 and later, this may require additional configuration handling for certain scenarios.15 Upgrade Guide.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

// ParameterSpec defines the desired state of Parameter
type ParameterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterParameters_2 `json:"forProvider"`
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
	InitProvider ParameterInitParameters_2 `json:"initProvider,omitempty"`
}

// ParameterStatus defines the observed state of Parameter.
type ParameterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Parameter is the Schema for the Parameters API. Provides a SSM Parameter resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Parameter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ParameterSpec   `json:"spec"`
	Status ParameterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterList contains a list of Parameters
type ParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Parameter `json:"items"`
}

// Repository type metadata.
var (
	Parameter_Kind             = "Parameter"
	Parameter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Parameter_Kind}.String()
	Parameter_KindAPIVersion   = Parameter_Kind + "." + CRDGroupVersion.String()
	Parameter_GroupVersionKind = CRDGroupVersion.WithKind(Parameter_Kind)
)

func init() {
	SchemeBuilder.Register(&Parameter{}, &ParameterList{})
}
