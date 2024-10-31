// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DNSInitParameters struct {

	// (String) The TTL to send when returning DNS results.
	// The TTL to send when returning DNS results.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type DNSObservation struct {

	// (String) The TTL to send when returning DNS results.
	// The TTL to send when returning DNS results.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type DNSParameters struct {

	// (String) The TTL to send when returning DNS results.
	// The TTL to send when returning DNS results.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type FailoverInitParameters struct {

	// (List of String) Remote datacenters to return results from.
	// Remote datacenters to return results from.
	Datacenters []*string `json:"datacenters,omitempty" tf:"datacenters,omitempty"`

	// (Number) Return results from this many datacenters, sorted in ascending order of estimated RTT.
	// Return results from this many datacenters, sorted in ascending order of estimated RTT.
	NearestN *float64 `json:"nearestN,omitempty" tf:"nearest_n,omitempty"`

	// (Block List) Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with nearest_n or datacenters. (see below for nested schema)
	// Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with `nearest_n` or `datacenters`.
	Targets []TargetsInitParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type FailoverObservation struct {

	// (List of String) Remote datacenters to return results from.
	// Remote datacenters to return results from.
	Datacenters []*string `json:"datacenters,omitempty" tf:"datacenters,omitempty"`

	// (Number) Return results from this many datacenters, sorted in ascending order of estimated RTT.
	// Return results from this many datacenters, sorted in ascending order of estimated RTT.
	NearestN *float64 `json:"nearestN,omitempty" tf:"nearest_n,omitempty"`

	// (Block List) Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with nearest_n or datacenters. (see below for nested schema)
	// Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with `nearest_n` or `datacenters`.
	Targets []TargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`
}

type FailoverParameters struct {

	// (List of String) Remote datacenters to return results from.
	// Remote datacenters to return results from.
	// +kubebuilder:validation:Optional
	Datacenters []*string `json:"datacenters,omitempty" tf:"datacenters,omitempty"`

	// (Number) Return results from this many datacenters, sorted in ascending order of estimated RTT.
	// Return results from this many datacenters, sorted in ascending order of estimated RTT.
	// +kubebuilder:validation:Optional
	NearestN *float64 `json:"nearestN,omitempty" tf:"nearest_n,omitempty"`

	// (Block List) Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with nearest_n or datacenters. (see below for nested schema)
	// Specifies a sequential list of remote datacenters and cluster peers to failover to if there are no healthy service instances in the local datacenter. This option cannot be used with `nearest_n` or `datacenters`.
	// +kubebuilder:validation:Optional
	Targets []TargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type QueryInitParameters struct {

	// (Boolean) When true the prepared query will return connect proxy services for a queried service.  Conditions such as tags in the prepared query will be matched against the proxy service. Defaults to false.
	// When `true` the prepared query will return connect proxy services for a queried service.  Conditions such as `tags` in the prepared query will be matched against the proxy service. Defaults to false.
	Connect *bool `json:"connect,omitempty" tf:"connect,omitempty"`

	// (Block List, Max: 1) Settings for controlling the DNS response details. (see below for nested schema)
	// Settings for controlling the DNS response details.
	DNS []DNSInitParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (Block List, Max: 1) Options for controlling behavior when no healthy nodes are available in the local DC. (see below for nested schema)
	// Options for controlling behavior when no healthy nodes are available in the local DC.
	Failover []FailoverInitParameters `json:"failover,omitempty" tf:"failover,omitempty"`

	// defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	// Specifies a list of check IDs that should be ignored when filtering unhealthy instances. This is mostly useful in an emergency or as a temporary measure when a health check is found to be unreliable. Being able to ignore it in centrally-defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	IgnoreCheckIds []*string `json:"ignoreCheckIds,omitempty" tf:"ignore_check_ids,omitempty"`

	// all.
	// The name of the prepared query. Used to identify the prepared query during requests. Can be specified as an empty string to configure the query as a catch-all.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic _agent value can be used to always sort nearest the node servicing the request.
	// Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic `_agent` value can be used to always sort nearest the node servicing the request.
	Near *string `json:"near,omitempty" tf:"near,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// +mapType=granular
	NodeMeta map[string]*string `json:"nodeMeta,omitempty" tf:"node_meta,omitempty"`

	// (Boolean) When true, the prepared query will only return nodes with passing health checks in the result.
	// When `true`, the prepared query will only return nodes with passing health checks in the result.
	OnlyPassing *bool `json:"onlyPassing,omitempty" tf:"only_passing,omitempty"`

	// (String) The name of the service to query
	// The name of the service to query
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// +mapType=granular
	ServiceMeta map[string]*string `json:"serviceMeta,omitempty" tf:"service_meta,omitempty"`

	// (String) The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	// The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	Session *string `json:"session,omitempty" tf:"session,omitempty"`

	// (String) The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	// The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	StoredToken *string `json:"storedToken,omitempty" tf:"stored_token,omitempty"`

	// (Set of String) The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Query templating options. This is used to make a single prepared query respond to many different requests (see below for nested schema)
	// Query templating options. This is used to make a single prepared query respond to many different requests
	Template []TemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (String, Sensitive, Deprecated) The ACL token to use when saving the prepared query. This overrides the token that the agent provides by default.
	// The ACL token to use when saving the prepared query. This overrides the token that the agent provides by default.
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type QueryObservation struct {

	// (Boolean) When true the prepared query will return connect proxy services for a queried service.  Conditions such as tags in the prepared query will be matched against the proxy service. Defaults to false.
	// When `true` the prepared query will return connect proxy services for a queried service.  Conditions such as `tags` in the prepared query will be matched against the proxy service. Defaults to false.
	Connect *bool `json:"connect,omitempty" tf:"connect,omitempty"`

	// (Block List, Max: 1) Settings for controlling the DNS response details. (see below for nested schema)
	// Settings for controlling the DNS response details.
	DNS []DNSObservation `json:"dns,omitempty" tf:"dns,omitempty"`

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (Block List, Max: 1) Options for controlling behavior when no healthy nodes are available in the local DC. (see below for nested schema)
	// Options for controlling behavior when no healthy nodes are available in the local DC.
	Failover []FailoverObservation `json:"failover,omitempty" tf:"failover,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	// Specifies a list of check IDs that should be ignored when filtering unhealthy instances. This is mostly useful in an emergency or as a temporary measure when a health check is found to be unreliable. Being able to ignore it in centrally-defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	IgnoreCheckIds []*string `json:"ignoreCheckIds,omitempty" tf:"ignore_check_ids,omitempty"`

	// all.
	// The name of the prepared query. Used to identify the prepared query during requests. Can be specified as an empty string to configure the query as a catch-all.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic _agent value can be used to always sort nearest the node servicing the request.
	// Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic `_agent` value can be used to always sort nearest the node servicing the request.
	Near *string `json:"near,omitempty" tf:"near,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// +mapType=granular
	NodeMeta map[string]*string `json:"nodeMeta,omitempty" tf:"node_meta,omitempty"`

	// (Boolean) When true, the prepared query will only return nodes with passing health checks in the result.
	// When `true`, the prepared query will only return nodes with passing health checks in the result.
	OnlyPassing *bool `json:"onlyPassing,omitempty" tf:"only_passing,omitempty"`

	// (String) The name of the service to query
	// The name of the service to query
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// +mapType=granular
	ServiceMeta map[string]*string `json:"serviceMeta,omitempty" tf:"service_meta,omitempty"`

	// (String) The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	// The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	Session *string `json:"session,omitempty" tf:"session,omitempty"`

	// (String) The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	// The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	StoredToken *string `json:"storedToken,omitempty" tf:"stored_token,omitempty"`

	// (Set of String) The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Query templating options. This is used to make a single prepared query respond to many different requests (see below for nested schema)
	// Query templating options. This is used to make a single prepared query respond to many different requests
	Template []TemplateObservation `json:"template,omitempty" tf:"template,omitempty"`
}

type QueryParameters struct {

	// (Boolean) When true the prepared query will return connect proxy services for a queried service.  Conditions such as tags in the prepared query will be matched against the proxy service. Defaults to false.
	// When `true` the prepared query will return connect proxy services for a queried service.  Conditions such as `tags` in the prepared query will be matched against the proxy service. Defaults to false.
	// +kubebuilder:validation:Optional
	Connect *bool `json:"connect,omitempty" tf:"connect,omitempty"`

	// (Block List, Max: 1) Settings for controlling the DNS response details. (see below for nested schema)
	// Settings for controlling the DNS response details.
	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (Block List, Max: 1) Options for controlling behavior when no healthy nodes are available in the local DC. (see below for nested schema)
	// Options for controlling behavior when no healthy nodes are available in the local DC.
	// +kubebuilder:validation:Optional
	Failover []FailoverParameters `json:"failover,omitempty" tf:"failover,omitempty"`

	// defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	// Specifies a list of check IDs that should be ignored when filtering unhealthy instances. This is mostly useful in an emergency or as a temporary measure when a health check is found to be unreliable. Being able to ignore it in centrally-defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	// +kubebuilder:validation:Optional
	IgnoreCheckIds []*string `json:"ignoreCheckIds,omitempty" tf:"ignore_check_ids,omitempty"`

	// all.
	// The name of the prepared query. Used to identify the prepared query during requests. Can be specified as an empty string to configure the query as a catch-all.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic _agent value can be used to always sort nearest the node servicing the request.
	// Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic `_agent` value can be used to always sort nearest the node servicing the request.
	// +kubebuilder:validation:Optional
	Near *string `json:"near,omitempty" tf:"near,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	NodeMeta map[string]*string `json:"nodeMeta,omitempty" tf:"node_meta,omitempty"`

	// (Boolean) When true, the prepared query will only return nodes with passing health checks in the result.
	// When `true`, the prepared query will only return nodes with passing health checks in the result.
	// +kubebuilder:validation:Optional
	OnlyPassing *bool `json:"onlyPassing,omitempty" tf:"only_passing,omitempty"`

	// (String) The name of the service to query
	// The name of the service to query
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ServiceMeta map[string]*string `json:"serviceMeta,omitempty" tf:"service_meta,omitempty"`

	// (String) The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	// The name of the Consul session to tie this query's lifetime to.  This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	// +kubebuilder:validation:Optional
	Session *string `json:"session,omitempty" tf:"session,omitempty"`

	// (String) The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	// The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.
	// +kubebuilder:validation:Optional
	StoredToken *string `json:"storedToken,omitempty" tf:"stored_token,omitempty"`

	// (Set of String) The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// The list of required and/or disallowed tags.  If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Query templating options. This is used to make a single prepared query respond to many different requests (see below for nested schema)
	// Query templating options. This is used to make a single prepared query respond to many different requests
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (String, Sensitive, Deprecated) The ACL token to use when saving the prepared query. This overrides the token that the agent provides by default.
	// The ACL token to use when saving the prepared query. This overrides the token that the agent provides by default.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type TargetsInitParameters struct {

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// Specifies a WAN federated datacenter to forward the query to.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (String) Specifies a cluster peer to use for failover.
	// Specifies a cluster peer to use for failover.
	Peer *string `json:"peer,omitempty" tf:"peer,omitempty"`
}

type TargetsObservation struct {

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// Specifies a WAN federated datacenter to forward the query to.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (String) Specifies a cluster peer to use for failover.
	// Specifies a cluster peer to use for failover.
	Peer *string `json:"peer,omitempty" tf:"peer,omitempty"`
}

type TargetsParameters struct {

	// (String) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	// Specifies a WAN federated datacenter to forward the query to.
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// (String) Specifies a cluster peer to use for failover.
	// Specifies a cluster peer to use for failover.
	// +kubebuilder:validation:Optional
	Peer *string `json:"peer,omitempty" tf:"peer,omitempty"`
}

type TemplateInitParameters struct {

	// (String) The regular expression to match with. When using name_prefix_match, this regex is applied against the query name.
	// The regular expression to match with. When using `name_prefix_match`, this regex is applied against the query name.
	Regexp *string `json:"regexp,omitempty" tf:"regexp,omitempty"`

	// (Boolean) If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	// If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	RemoveEmptyTags *bool `json:"removeEmptyTags,omitempty" tf:"remove_empty_tags,omitempty"`

	// (String) The type of template matching to perform. Currently only name_prefix_match is supported.
	// The type of template matching to perform. Currently only `name_prefix_match` is supported.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TemplateObservation struct {

	// (String) The regular expression to match with. When using name_prefix_match, this regex is applied against the query name.
	// The regular expression to match with. When using `name_prefix_match`, this regex is applied against the query name.
	Regexp *string `json:"regexp,omitempty" tf:"regexp,omitempty"`

	// (Boolean) If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	// If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	RemoveEmptyTags *bool `json:"removeEmptyTags,omitempty" tf:"remove_empty_tags,omitempty"`

	// (String) The type of template matching to perform. Currently only name_prefix_match is supported.
	// The type of template matching to perform. Currently only `name_prefix_match` is supported.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TemplateParameters struct {

	// (String) The regular expression to match with. When using name_prefix_match, this regex is applied against the query name.
	// The regular expression to match with. When using `name_prefix_match`, this regex is applied against the query name.
	// +kubebuilder:validation:Optional
	Regexp *string `json:"regexp" tf:"regexp,omitempty"`

	// (Boolean) If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	// If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	// +kubebuilder:validation:Optional
	RemoveEmptyTags *bool `json:"removeEmptyTags,omitempty" tf:"remove_empty_tags,omitempty"`

	// (String) The type of template matching to perform. Currently only name_prefix_match is supported.
	// The type of template matching to perform. Currently only `name_prefix_match` is supported.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// QuerySpec defines the desired state of Query
type QuerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueryParameters `json:"forProvider"`
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
	InitProvider QueryInitParameters `json:"initProvider,omitempty"`
}

// QueryStatus defines the observed state of Query.
type QueryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Query is the Schema for the Querys API.  Managing prepared queries is done using Consul's REST API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,consul}
type Query struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.service) || (has(self.initProvider) && has(self.initProvider.service))",message="spec.forProvider.service is a required parameter"
	Spec   QuerySpec   `json:"spec"`
	Status QueryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueryList contains a list of Querys
type QueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Query `json:"items"`
}

// Repository type metadata.
var (
	Query_Kind             = "Query"
	Query_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Query_Kind}.String()
	Query_KindAPIVersion   = Query_Kind + "." + CRDGroupVersion.String()
	Query_GroupVersionKind = CRDGroupVersion.WithKind(Query_Kind)
)

func init() {
	SchemeBuilder.Register(&Query{}, &QueryList{})
}
