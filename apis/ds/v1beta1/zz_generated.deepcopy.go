//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarder) DeepCopyInto(out *ConditionalForwarder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarder.
func (in *ConditionalForwarder) DeepCopy() *ConditionalForwarder {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConditionalForwarder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderInitParameters) DeepCopyInto(out *ConditionalForwarderInitParameters) {
	*out = *in
	if in.DNSIps != nil {
		in, out := &in.DNSIps, &out.DNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderInitParameters.
func (in *ConditionalForwarderInitParameters) DeepCopy() *ConditionalForwarderInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderList) DeepCopyInto(out *ConditionalForwarderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConditionalForwarder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderList.
func (in *ConditionalForwarderList) DeepCopy() *ConditionalForwarderList {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConditionalForwarderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderObservation) DeepCopyInto(out *ConditionalForwarderObservation) {
	*out = *in
	if in.DNSIps != nil {
		in, out := &in.DNSIps, &out.DNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RemoteDomainName != nil {
		in, out := &in.RemoteDomainName, &out.RemoteDomainName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderObservation.
func (in *ConditionalForwarderObservation) DeepCopy() *ConditionalForwarderObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderParameters) DeepCopyInto(out *ConditionalForwarderParameters) {
	*out = *in
	if in.DNSIps != nil {
		in, out := &in.DNSIps, &out.DNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.DirectoryIDRef != nil {
		in, out := &in.DirectoryIDRef, &out.DirectoryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DirectoryIDSelector != nil {
		in, out := &in.DirectoryIDSelector, &out.DirectoryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RemoteDomainName != nil {
		in, out := &in.RemoteDomainName, &out.RemoteDomainName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderParameters.
func (in *ConditionalForwarderParameters) DeepCopy() *ConditionalForwarderParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderSpec) DeepCopyInto(out *ConditionalForwarderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderSpec.
func (in *ConditionalForwarderSpec) DeepCopy() *ConditionalForwarderSpec {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalForwarderStatus) DeepCopyInto(out *ConditionalForwarderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalForwarderStatus.
func (in *ConditionalForwarderStatus) DeepCopy() *ConditionalForwarderStatus {
	if in == nil {
		return nil
	}
	out := new(ConditionalForwarderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectSettingsInitParameters) DeepCopyInto(out *ConnectSettingsInitParameters) {
	*out = *in
	if in.CustomerDNSIps != nil {
		in, out := &in.CustomerDNSIps, &out.CustomerDNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomerUsername != nil {
		in, out := &in.CustomerUsername, &out.CustomerUsername
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectSettingsInitParameters.
func (in *ConnectSettingsInitParameters) DeepCopy() *ConnectSettingsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectSettingsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectSettingsObservation) DeepCopyInto(out *ConnectSettingsObservation) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConnectIps != nil {
		in, out := &in.ConnectIps, &out.ConnectIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomerDNSIps != nil {
		in, out := &in.CustomerDNSIps, &out.CustomerDNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomerUsername != nil {
		in, out := &in.CustomerUsername, &out.CustomerUsername
		*out = new(string)
		**out = **in
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectSettingsObservation.
func (in *ConnectSettingsObservation) DeepCopy() *ConnectSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(ConnectSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectSettingsParameters) DeepCopyInto(out *ConnectSettingsParameters) {
	*out = *in
	if in.CustomerDNSIps != nil {
		in, out := &in.CustomerDNSIps, &out.CustomerDNSIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomerUsername != nil {
		in, out := &in.CustomerUsername, &out.CustomerUsername
		*out = new(string)
		**out = **in
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIdsRefs != nil {
		in, out := &in.SubnetIdsRefs, &out.SubnetIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIdsSelector != nil {
		in, out := &in.SubnetIdsSelector, &out.SubnetIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectSettingsParameters.
func (in *ConnectSettingsParameters) DeepCopy() *ConnectSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Directory) DeepCopyInto(out *Directory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Directory.
func (in *Directory) DeepCopy() *Directory {
	if in == nil {
		return nil
	}
	out := new(Directory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Directory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryInitParameters) DeepCopyInto(out *DirectoryInitParameters) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ConnectSettings != nil {
		in, out := &in.ConnectSettings, &out.ConnectSettings
		*out = make([]ConnectSettingsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DesiredNumberOfDomainControllers != nil {
		in, out := &in.DesiredNumberOfDomainControllers, &out.DesiredNumberOfDomainControllers
		*out = new(float64)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.EnableSso != nil {
		in, out := &in.EnableSso, &out.EnableSso
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VPCSettings != nil {
		in, out := &in.VPCSettings, &out.VPCSettings
		*out = make([]VPCSettingsInitParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryInitParameters.
func (in *DirectoryInitParameters) DeepCopy() *DirectoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryList) DeepCopyInto(out *DirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Directory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryList.
func (in *DirectoryList) DeepCopy() *DirectoryList {
	if in == nil {
		return nil
	}
	out := new(DirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryObservation) DeepCopyInto(out *DirectoryObservation) {
	*out = *in
	if in.AccessURL != nil {
		in, out := &in.AccessURL, &out.AccessURL
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ConnectSettings != nil {
		in, out := &in.ConnectSettings, &out.ConnectSettings
		*out = make([]ConnectSettingsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSIPAddresses != nil {
		in, out := &in.DNSIPAddresses, &out.DNSIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DesiredNumberOfDomainControllers != nil {
		in, out := &in.DesiredNumberOfDomainControllers, &out.DesiredNumberOfDomainControllers
		*out = new(float64)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.EnableSso != nil {
		in, out := &in.EnableSso, &out.EnableSso
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VPCSettings != nil {
		in, out := &in.VPCSettings, &out.VPCSettings
		*out = make([]VPCSettingsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryObservation.
func (in *DirectoryObservation) DeepCopy() *DirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(DirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryParameters) DeepCopyInto(out *DirectoryParameters) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ConnectSettings != nil {
		in, out := &in.ConnectSettings, &out.ConnectSettings
		*out = make([]ConnectSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DesiredNumberOfDomainControllers != nil {
		in, out := &in.DesiredNumberOfDomainControllers, &out.DesiredNumberOfDomainControllers
		*out = new(float64)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.EnableSso != nil {
		in, out := &in.EnableSso, &out.EnableSso
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VPCSettings != nil {
		in, out := &in.VPCSettings, &out.VPCSettings
		*out = make([]VPCSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryParameters.
func (in *DirectoryParameters) DeepCopy() *DirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectorySpec) DeepCopyInto(out *DirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectorySpec.
func (in *DirectorySpec) DeepCopy() *DirectorySpec {
	if in == nil {
		return nil
	}
	out := new(DirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryStatus) DeepCopyInto(out *DirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryStatus.
func (in *DirectoryStatus) DeepCopy() *DirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(DirectoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectory) DeepCopyInto(out *SharedDirectory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectory.
func (in *SharedDirectory) DeepCopy() *SharedDirectory {
	if in == nil {
		return nil
	}
	out := new(SharedDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedDirectory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectoryInitParameters) DeepCopyInto(out *SharedDirectoryInitParameters) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectoryInitParameters.
func (in *SharedDirectoryInitParameters) DeepCopy() *SharedDirectoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(SharedDirectoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectoryList) DeepCopyInto(out *SharedDirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedDirectory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectoryList.
func (in *SharedDirectoryList) DeepCopy() *SharedDirectoryList {
	if in == nil {
		return nil
	}
	out := new(SharedDirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedDirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectoryObservation) DeepCopyInto(out *SharedDirectoryObservation) {
	*out = *in
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.SharedDirectoryID != nil {
		in, out := &in.SharedDirectoryID, &out.SharedDirectoryID
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectoryObservation.
func (in *SharedDirectoryObservation) DeepCopy() *SharedDirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(SharedDirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectoryParameters) DeepCopyInto(out *SharedDirectoryParameters) {
	*out = *in
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.DirectoryIDRef != nil {
		in, out := &in.DirectoryIDRef, &out.DirectoryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DirectoryIDSelector != nil {
		in, out := &in.DirectoryIDSelector, &out.DirectoryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.NotesSecretRef != nil {
		in, out := &in.NotesSecretRef, &out.NotesSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectoryParameters.
func (in *SharedDirectoryParameters) DeepCopy() *SharedDirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(SharedDirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectorySpec) DeepCopyInto(out *SharedDirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectorySpec.
func (in *SharedDirectorySpec) DeepCopy() *SharedDirectorySpec {
	if in == nil {
		return nil
	}
	out := new(SharedDirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDirectoryStatus) DeepCopyInto(out *SharedDirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDirectoryStatus.
func (in *SharedDirectoryStatus) DeepCopy() *SharedDirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(SharedDirectoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetInitParameters) DeepCopyInto(out *TargetInitParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetInitParameters.
func (in *TargetInitParameters) DeepCopy() *TargetInitParameters {
	if in == nil {
		return nil
	}
	out := new(TargetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetObservation) DeepCopyInto(out *TargetObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetObservation.
func (in *TargetObservation) DeepCopy() *TargetObservation {
	if in == nil {
		return nil
	}
	out := new(TargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParameters) DeepCopyInto(out *TargetParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParameters.
func (in *TargetParameters) DeepCopy() *TargetParameters {
	if in == nil {
		return nil
	}
	out := new(TargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSettingsInitParameters) DeepCopyInto(out *VPCSettingsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSettingsInitParameters.
func (in *VPCSettingsInitParameters) DeepCopy() *VPCSettingsInitParameters {
	if in == nil {
		return nil
	}
	out := new(VPCSettingsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSettingsObservation) DeepCopyInto(out *VPCSettingsObservation) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSettingsObservation.
func (in *VPCSettingsObservation) DeepCopy() *VPCSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(VPCSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSettingsParameters) DeepCopyInto(out *VPCSettingsParameters) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIdsRefs != nil {
		in, out := &in.SubnetIdsRefs, &out.SubnetIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIdsSelector != nil {
		in, out := &in.SubnetIdsSelector, &out.SubnetIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSettingsParameters.
func (in *VPCSettingsParameters) DeepCopy() *VPCSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(VPCSettingsParameters)
	in.DeepCopyInto(out)
	return out
}
