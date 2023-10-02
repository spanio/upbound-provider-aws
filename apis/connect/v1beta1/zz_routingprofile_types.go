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

type MediaConcurrenciesInitParameters struct {

	// Specifies the channels that agents can handle in the Contact Control Panel (CCP). Valid values are VOICE, CHAT, TASK.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Specifies the number of contacts an agent can have on a channel simultaneously. Valid Range for VOICE: Minimum value of 1. Maximum value of 1. Valid Range for CHAT: Minimum value of 1. Maximum value of 10. Valid Range for TASK: Minimum value of 1. Maximum value of 10.
	Concurrency *float64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`
}

type MediaConcurrenciesObservation struct {

	// Specifies the channels that agents can handle in the Contact Control Panel (CCP). Valid values are VOICE, CHAT, TASK.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Specifies the number of contacts an agent can have on a channel simultaneously. Valid Range for VOICE: Minimum value of 1. Maximum value of 1. Valid Range for CHAT: Minimum value of 1. Maximum value of 10. Valid Range for TASK: Minimum value of 1. Maximum value of 10.
	Concurrency *float64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`
}

type MediaConcurrenciesParameters struct {

	// Specifies the channels that agents can handle in the Contact Control Panel (CCP). Valid values are VOICE, CHAT, TASK.
	// +kubebuilder:validation:Optional
	Channel *string `json:"channel" tf:"channel,omitempty"`

	// Specifies the number of contacts an agent can have on a channel simultaneously. Valid Range for VOICE: Minimum value of 1. Maximum value of 1. Valid Range for CHAT: Minimum value of 1. Maximum value of 10. Valid Range for TASK: Minimum value of 1. Maximum value of 10.
	// +kubebuilder:validation:Optional
	Concurrency *float64 `json:"concurrency" tf:"concurrency,omitempty"`
}

type QueueConfigsAssociatedInitParameters struct {
}

type QueueConfigsAssociatedObservation struct {

	// Specifies the channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are VOICE, CHAT, TASK.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Specifies the delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// Specifies the order in which contacts are to be handled for the queue.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// ARN for the queue.
	QueueArn *string `json:"queueArn,omitempty" tf:"queue_arn,omitempty"`

	// Specifies the identifier for the queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`

	// Name for the queue.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`
}

type QueueConfigsAssociatedParameters struct {
}

type QueueConfigsInitParameters struct {

	// Specifies the channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are VOICE, CHAT, TASK.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Specifies the delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// Specifies the order in which contacts are to be handled for the queue.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the identifier for the queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`
}

type QueueConfigsObservation struct {

	// Specifies the channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are VOICE, CHAT, TASK.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Specifies the delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	Delay *float64 `json:"delay,omitempty" tf:"delay,omitempty"`

	// Specifies the order in which contacts are to be handled for the queue.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// ARN for the queue.
	QueueArn *string `json:"queueArn,omitempty" tf:"queue_arn,omitempty"`

	// Specifies the identifier for the queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`

	// Name for the queue.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`
}

type QueueConfigsParameters struct {

	// Specifies the channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are VOICE, CHAT, TASK.
	// +kubebuilder:validation:Optional
	Channel *string `json:"channel" tf:"channel,omitempty"`

	// Specifies the delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	// +kubebuilder:validation:Optional
	Delay *float64 `json:"delay" tf:"delay,omitempty"`

	// Specifies the order in which contacts are to be handled for the queue.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Specifies the identifier for the queue.
	// +kubebuilder:validation:Optional
	QueueID *string `json:"queueId" tf:"queue_id,omitempty"`
}

type RoutingProfileInitParameters struct {

	// Specifies the description of the Routing Profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more media_concurrencies blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The media_concurrencies block is documented below.
	MediaConcurrencies []MediaConcurrenciesInitParameters `json:"mediaConcurrencies,omitempty" tf:"media_concurrencies,omitempty"`

	// Specifies the name of the Routing Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more queue_configs blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The queue_configs block is documented below.
	QueueConfigs []QueueConfigsInitParameters `json:"queueConfigs,omitempty" tf:"queue_configs,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RoutingProfileObservation struct {

	// The Amazon Resource Name (ARN) of the Routing Profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies the default outbound queue for the Routing Profile.
	DefaultOutboundQueueID *string `json:"defaultOutboundQueueId,omitempty" tf:"default_outbound_queue_id,omitempty"`

	// Specifies the description of the Routing Profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Routing Profile separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// One or more media_concurrencies blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The media_concurrencies block is documented below.
	MediaConcurrencies []MediaConcurrenciesObservation `json:"mediaConcurrencies,omitempty" tf:"media_concurrencies,omitempty"`

	// Specifies the name of the Routing Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more queue_configs blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The queue_configs block is documented below.
	QueueConfigs []QueueConfigsObservation `json:"queueConfigs,omitempty" tf:"queue_configs,omitempty"`

	QueueConfigsAssociated []QueueConfigsAssociatedObservation `json:"queueConfigsAssociated,omitempty" tf:"queue_configs_associated,omitempty"`

	// The identifier for the Routing Profile.
	RoutingProfileID *string `json:"routingProfileId,omitempty" tf:"routing_profile_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RoutingProfileParameters struct {

	// Specifies the default outbound queue for the Routing Profile.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("queue_id",true)
	// +kubebuilder:validation:Optional
	DefaultOutboundQueueID *string `json:"defaultOutboundQueueId,omitempty" tf:"default_outbound_queue_id,omitempty"`

	// Reference to a Queue in connect to populate defaultOutboundQueueId.
	// +kubebuilder:validation:Optional
	DefaultOutboundQueueIDRef *v1.Reference `json:"defaultOutboundQueueIdRef,omitempty" tf:"-"`

	// Selector for a Queue in connect to populate defaultOutboundQueueId.
	// +kubebuilder:validation:Optional
	DefaultOutboundQueueIDSelector *v1.Selector `json:"defaultOutboundQueueIdSelector,omitempty" tf:"-"`

	// Specifies the description of the Routing Profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// One or more media_concurrencies blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The media_concurrencies block is documented below.
	// +kubebuilder:validation:Optional
	MediaConcurrencies []MediaConcurrenciesParameters `json:"mediaConcurrencies,omitempty" tf:"media_concurrencies,omitempty"`

	// Specifies the name of the Routing Profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more queue_configs blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The queue_configs block is documented below.
	// +kubebuilder:validation:Optional
	QueueConfigs []QueueConfigsParameters `json:"queueConfigs,omitempty" tf:"queue_configs,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RoutingProfileSpec defines the desired state of RoutingProfile
type RoutingProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingProfileParameters `json:"forProvider"`
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
	InitProvider RoutingProfileInitParameters `json:"initProvider,omitempty"`
}

// RoutingProfileStatus defines the observed state of RoutingProfile.
type RoutingProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingProfile is the Schema for the RoutingProfiles API. Provides details about a specific Amazon Connect Routing Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RoutingProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mediaConcurrencies) || (has(self.initProvider) && has(self.initProvider.mediaConcurrencies))",message="spec.forProvider.mediaConcurrencies is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RoutingProfileSpec   `json:"spec"`
	Status RoutingProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingProfileList contains a list of RoutingProfiles
type RoutingProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingProfile `json:"items"`
}

// Repository type metadata.
var (
	RoutingProfile_Kind             = "RoutingProfile"
	RoutingProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingProfile_Kind}.String()
	RoutingProfile_KindAPIVersion   = RoutingProfile_Kind + "." + CRDGroupVersion.String()
	RoutingProfile_GroupVersionKind = CRDGroupVersion.WithKind(RoutingProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingProfile{}, &RoutingProfileList{})
}
