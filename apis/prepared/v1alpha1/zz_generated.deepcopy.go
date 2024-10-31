//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSInitParameters) DeepCopyInto(out *DNSInitParameters) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSInitParameters.
func (in *DNSInitParameters) DeepCopy() *DNSInitParameters {
	if in == nil {
		return nil
	}
	out := new(DNSInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSObservation) DeepCopyInto(out *DNSObservation) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSObservation.
func (in *DNSObservation) DeepCopy() *DNSObservation {
	if in == nil {
		return nil
	}
	out := new(DNSObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSParameters) DeepCopyInto(out *DNSParameters) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSParameters.
func (in *DNSParameters) DeepCopy() *DNSParameters {
	if in == nil {
		return nil
	}
	out := new(DNSParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverInitParameters) DeepCopyInto(out *FailoverInitParameters) {
	*out = *in
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NearestN != nil {
		in, out := &in.NearestN, &out.NearestN
		*out = new(float64)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverInitParameters.
func (in *FailoverInitParameters) DeepCopy() *FailoverInitParameters {
	if in == nil {
		return nil
	}
	out := new(FailoverInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverObservation) DeepCopyInto(out *FailoverObservation) {
	*out = *in
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NearestN != nil {
		in, out := &in.NearestN, &out.NearestN
		*out = new(float64)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverObservation.
func (in *FailoverObservation) DeepCopy() *FailoverObservation {
	if in == nil {
		return nil
	}
	out := new(FailoverObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverParameters) DeepCopyInto(out *FailoverParameters) {
	*out = *in
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NearestN != nil {
		in, out := &in.NearestN, &out.NearestN
		*out = new(float64)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverParameters.
func (in *FailoverParameters) DeepCopy() *FailoverParameters {
	if in == nil {
		return nil
	}
	out := new(FailoverParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Query) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryInitParameters) DeepCopyInto(out *QueryInitParameters) {
	*out = *in
	if in.Connect != nil {
		in, out := &in.Connect, &out.Connect
		*out = new(bool)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]DNSInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = make([]FailoverInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreCheckIds != nil {
		in, out := &in.IgnoreCheckIds, &out.IgnoreCheckIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Near != nil {
		in, out := &in.Near, &out.Near
		*out = new(string)
		**out = **in
	}
	if in.NodeMeta != nil {
		in, out := &in.NodeMeta, &out.NodeMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OnlyPassing != nil {
		in, out := &in.OnlyPassing, &out.OnlyPassing
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.ServiceMeta != nil {
		in, out := &in.ServiceMeta, &out.ServiceMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(string)
		**out = **in
	}
	if in.StoredToken != nil {
		in, out := &in.StoredToken, &out.StoredToken
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]TemplateInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryInitParameters.
func (in *QueryInitParameters) DeepCopy() *QueryInitParameters {
	if in == nil {
		return nil
	}
	out := new(QueryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryList) DeepCopyInto(out *QueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Query, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryList.
func (in *QueryList) DeepCopy() *QueryList {
	if in == nil {
		return nil
	}
	out := new(QueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryObservation) DeepCopyInto(out *QueryObservation) {
	*out = *in
	if in.Connect != nil {
		in, out := &in.Connect, &out.Connect
		*out = new(bool)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]DNSObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = make([]FailoverObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCheckIds != nil {
		in, out := &in.IgnoreCheckIds, &out.IgnoreCheckIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Near != nil {
		in, out := &in.Near, &out.Near
		*out = new(string)
		**out = **in
	}
	if in.NodeMeta != nil {
		in, out := &in.NodeMeta, &out.NodeMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OnlyPassing != nil {
		in, out := &in.OnlyPassing, &out.OnlyPassing
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.ServiceMeta != nil {
		in, out := &in.ServiceMeta, &out.ServiceMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(string)
		**out = **in
	}
	if in.StoredToken != nil {
		in, out := &in.StoredToken, &out.StoredToken
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]TemplateObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryObservation.
func (in *QueryObservation) DeepCopy() *QueryObservation {
	if in == nil {
		return nil
	}
	out := new(QueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryParameters) DeepCopyInto(out *QueryParameters) {
	*out = *in
	if in.Connect != nil {
		in, out := &in.Connect, &out.Connect
		*out = new(bool)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]DNSParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = make([]FailoverParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreCheckIds != nil {
		in, out := &in.IgnoreCheckIds, &out.IgnoreCheckIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Near != nil {
		in, out := &in.Near, &out.Near
		*out = new(string)
		**out = **in
	}
	if in.NodeMeta != nil {
		in, out := &in.NodeMeta, &out.NodeMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OnlyPassing != nil {
		in, out := &in.OnlyPassing, &out.OnlyPassing
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.ServiceMeta != nil {
		in, out := &in.ServiceMeta, &out.ServiceMeta
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(string)
		**out = **in
	}
	if in.StoredToken != nil {
		in, out := &in.StoredToken, &out.StoredToken
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]TemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameters.
func (in *QueryParameters) DeepCopy() *QueryParameters {
	if in == nil {
		return nil
	}
	out := new(QueryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuerySpec) DeepCopyInto(out *QuerySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuerySpec.
func (in *QuerySpec) DeepCopy() *QuerySpec {
	if in == nil {
		return nil
	}
	out := new(QuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryStatus) DeepCopyInto(out *QueryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryStatus.
func (in *QueryStatus) DeepCopy() *QueryStatus {
	if in == nil {
		return nil
	}
	out := new(QueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsInitParameters) DeepCopyInto(out *TargetsInitParameters) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Peer != nil {
		in, out := &in.Peer, &out.Peer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsInitParameters.
func (in *TargetsInitParameters) DeepCopy() *TargetsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsObservation) DeepCopyInto(out *TargetsObservation) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Peer != nil {
		in, out := &in.Peer, &out.Peer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsObservation.
func (in *TargetsObservation) DeepCopy() *TargetsObservation {
	if in == nil {
		return nil
	}
	out := new(TargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsParameters) DeepCopyInto(out *TargetsParameters) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Peer != nil {
		in, out := &in.Peer, &out.Peer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsParameters.
func (in *TargetsParameters) DeepCopy() *TargetsParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateInitParameters) DeepCopyInto(out *TemplateInitParameters) {
	*out = *in
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(string)
		**out = **in
	}
	if in.RemoveEmptyTags != nil {
		in, out := &in.RemoveEmptyTags, &out.RemoveEmptyTags
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateInitParameters.
func (in *TemplateInitParameters) DeepCopy() *TemplateInitParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateObservation) DeepCopyInto(out *TemplateObservation) {
	*out = *in
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(string)
		**out = **in
	}
	if in.RemoveEmptyTags != nil {
		in, out := &in.RemoveEmptyTags, &out.RemoveEmptyTags
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateObservation.
func (in *TemplateObservation) DeepCopy() *TemplateObservation {
	if in == nil {
		return nil
	}
	out := new(TemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParameters) DeepCopyInto(out *TemplateParameters) {
	*out = *in
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(string)
		**out = **in
	}
	if in.RemoveEmptyTags != nil {
		in, out := &in.RemoveEmptyTags, &out.RemoveEmptyTags
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParameters.
func (in *TemplateParameters) DeepCopy() *TemplateParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateParameters)
	in.DeepCopyInto(out)
	return out
}
