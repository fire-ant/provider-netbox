//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddress) DeepCopyInto(out *IPAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddress.
func (in *IPAddress) DeepCopy() *IPAddress {
	if in == nil {
		return nil
	}
	out := new(IPAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressList) DeepCopyInto(out *IPAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressList.
func (in *IPAddressList) DeepCopy() *IPAddressList {
	if in == nil {
		return nil
	}
	out := new(IPAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressObservation) DeepCopyInto(out *IPAddressObservation) {
	*out = *in
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPRangeID != nil {
		in, out := &in.IPRangeID, &out.IPRangeID
		*out = new(float64)
		**out = **in
	}
	if in.InterfaceID != nil {
		in, out := &in.InterfaceID, &out.InterfaceID
		*out = new(float64)
		**out = **in
	}
	if in.PrefixID != nil {
		in, out := &in.PrefixID, &out.PrefixID
		*out = new(float64)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
	if in.VrfID != nil {
		in, out := &in.VrfID, &out.VrfID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressObservation.
func (in *IPAddressObservation) DeepCopy() *IPAddressObservation {
	if in == nil {
		return nil
	}
	out := new(IPAddressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressParameters) DeepCopyInto(out *IPAddressParameters) {
	*out = *in
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IPRangeID != nil {
		in, out := &in.IPRangeID, &out.IPRangeID
		*out = new(float64)
		**out = **in
	}
	if in.IPRangeIDRef != nil {
		in, out := &in.IPRangeIDRef, &out.IPRangeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IPRangeIDSelector != nil {
		in, out := &in.IPRangeIDSelector, &out.IPRangeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.InterfaceID != nil {
		in, out := &in.InterfaceID, &out.InterfaceID
		*out = new(float64)
		**out = **in
	}
	if in.InterfaceIDRef != nil {
		in, out := &in.InterfaceIDRef, &out.InterfaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.InterfaceIDSelector != nil {
		in, out := &in.InterfaceIDSelector, &out.InterfaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PrefixID != nil {
		in, out := &in.PrefixID, &out.PrefixID
		*out = new(float64)
		**out = **in
	}
	if in.PrefixIDRef != nil {
		in, out := &in.PrefixIDRef, &out.PrefixIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrefixIDSelector != nil {
		in, out := &in.PrefixIDSelector, &out.PrefixIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
	if in.TenantIDRef != nil {
		in, out := &in.TenantIDRef, &out.TenantIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TenantIDSelector != nil {
		in, out := &in.TenantIDSelector, &out.TenantIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VrfID != nil {
		in, out := &in.VrfID, &out.VrfID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressParameters.
func (in *IPAddressParameters) DeepCopy() *IPAddressParameters {
	if in == nil {
		return nil
	}
	out := new(IPAddressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressSpec) DeepCopyInto(out *IPAddressSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressSpec.
func (in *IPAddressSpec) DeepCopy() *IPAddressSpec {
	if in == nil {
		return nil
	}
	out := new(IPAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressStatus) DeepCopyInto(out *IPAddressStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressStatus.
func (in *IPAddressStatus) DeepCopy() *IPAddressStatus {
	if in == nil {
		return nil
	}
	out := new(IPAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prefix) DeepCopyInto(out *Prefix) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prefix.
func (in *Prefix) DeepCopy() *Prefix {
	if in == nil {
		return nil
	}
	out := new(Prefix)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Prefix) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixList) DeepCopyInto(out *PrefixList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Prefix, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixList.
func (in *PrefixList) DeepCopy() *PrefixList {
	if in == nil {
		return nil
	}
	out := new(PrefixList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrefixList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixObservation) DeepCopyInto(out *PrefixObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsPool != nil {
		in, out := &in.IsPool, &out.IsPool
		*out = new(bool)
		**out = **in
	}
	if in.MarkUtilized != nil {
		in, out := &in.MarkUtilized, &out.MarkUtilized
		*out = new(bool)
		**out = **in
	}
	if in.ParentPrefixID != nil {
		in, out := &in.ParentPrefixID, &out.ParentPrefixID
		*out = new(float64)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(float64)
		**out = **in
	}
	if in.SiteID != nil {
		in, out := &in.SiteID, &out.SiteID
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(float64)
		**out = **in
	}
	if in.VrfID != nil {
		in, out := &in.VrfID, &out.VrfID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixObservation.
func (in *PrefixObservation) DeepCopy() *PrefixObservation {
	if in == nil {
		return nil
	}
	out := new(PrefixObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixParameters) DeepCopyInto(out *PrefixParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IsPool != nil {
		in, out := &in.IsPool, &out.IsPool
		*out = new(bool)
		**out = **in
	}
	if in.MarkUtilized != nil {
		in, out := &in.MarkUtilized, &out.MarkUtilized
		*out = new(bool)
		**out = **in
	}
	if in.ParentPrefixID != nil {
		in, out := &in.ParentPrefixID, &out.ParentPrefixID
		*out = new(float64)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(float64)
		**out = **in
	}
	if in.RoleIDRef != nil {
		in, out := &in.RoleIDRef, &out.RoleIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleIDSelector != nil {
		in, out := &in.RoleIDSelector, &out.RoleIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SiteID != nil {
		in, out := &in.SiteID, &out.SiteID
		*out = new(float64)
		**out = **in
	}
	if in.SiteIDRef != nil {
		in, out := &in.SiteIDRef, &out.SiteIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SiteIDSelector != nil {
		in, out := &in.SiteIDSelector, &out.SiteIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
	if in.TenantIDRef != nil {
		in, out := &in.TenantIDRef, &out.TenantIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TenantIDSelector != nil {
		in, out := &in.TenantIDSelector, &out.TenantIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(float64)
		**out = **in
	}
	if in.VlanIDRef != nil {
		in, out := &in.VlanIDRef, &out.VlanIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VlanIDSelector != nil {
		in, out := &in.VlanIDSelector, &out.VlanIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VrfID != nil {
		in, out := &in.VrfID, &out.VrfID
		*out = new(float64)
		**out = **in
	}
	if in.VrfIDRef != nil {
		in, out := &in.VrfIDRef, &out.VrfIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VrfIDSelector != nil {
		in, out := &in.VrfIDSelector, &out.VrfIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixParameters.
func (in *PrefixParameters) DeepCopy() *PrefixParameters {
	if in == nil {
		return nil
	}
	out := new(PrefixParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixSpec) DeepCopyInto(out *PrefixSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixSpec.
func (in *PrefixSpec) DeepCopy() *PrefixSpec {
	if in == nil {
		return nil
	}
	out := new(PrefixSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixStatus) DeepCopyInto(out *PrefixStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixStatus.
func (in *PrefixStatus) DeepCopy() *PrefixStatus {
	if in == nil {
		return nil
	}
	out := new(PrefixStatus)
	in.DeepCopyInto(out)
	return out
}
