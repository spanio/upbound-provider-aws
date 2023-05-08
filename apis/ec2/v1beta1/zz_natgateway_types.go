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

type NATGatewayObservation_2 struct {

	// The Allocation ID of the Elastic IP address for the gateway. Required for connectivity_type of public.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The association ID of the Elastic IP address that's associated with the NAT gateway. Only available when connectivity_type is public.
	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	// Connectivity type for the gateway. Valid values are private and public. Defaults to public.
	ConnectivityType *string `json:"connectivityType,omitempty" tf:"connectivity_type,omitempty"`

	// The ID of the NAT Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the network interface associated with the NAT gateway.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The Elastic IP address associated with the NAT gateway.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// The Subnet ID of the subnet in which to place the gateway.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NATGatewayParameters_2 struct {

	// The Allocation ID of the Elastic IP address for the gateway. Required for connectivity_type of public.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EIP
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// Reference to a EIP in ec2 to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDRef *v1.Reference `json:"allocationIdRef,omitempty" tf:"-"`

	// Selector for a EIP in ec2 to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDSelector *v1.Selector `json:"allocationIdSelector,omitempty" tf:"-"`

	// Connectivity type for the gateway. Valid values are private and public. Defaults to public.
	// +kubebuilder:validation:Optional
	ConnectivityType *string `json:"connectivityType,omitempty" tf:"connectivity_type,omitempty"`

	// The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Subnet ID of the subnet in which to place the gateway.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NATGatewaySpec defines the desired state of NATGateway
type NATGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATGatewayParameters_2 `json:"forProvider"`
}

// NATGatewayStatus defines the observed state of NATGateway.
type NATGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATGatewayObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATGateway is the Schema for the NATGateways API. Provides a resource to create a VPC NAT Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NATGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATGatewaySpec   `json:"spec"`
	Status            NATGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayList contains a list of NATGateways
type NATGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATGateway `json:"items"`
}

// Repository type metadata.
var (
	NATGateway_Kind             = "NATGateway"
	NATGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATGateway_Kind}.String()
	NATGateway_KindAPIVersion   = NATGateway_Kind + "." + CRDGroupVersion.String()
	NATGateway_GroupVersionKind = CRDGroupVersion.WithKind(NATGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&NATGateway{}, &NATGatewayList{})
}
