//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manufacturer) DeepCopyInto(out *Manufacturer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manufacturer.
func (in *Manufacturer) DeepCopy() *Manufacturer {
	if in == nil {
		return nil
	}
	out := new(Manufacturer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Manufacturer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManufacturerList) DeepCopyInto(out *ManufacturerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Manufacturer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufacturerList.
func (in *ManufacturerList) DeepCopy() *ManufacturerList {
	if in == nil {
		return nil
	}
	out := new(ManufacturerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManufacturerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManufacturerObservation) DeepCopyInto(out *ManufacturerObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufacturerObservation.
func (in *ManufacturerObservation) DeepCopy() *ManufacturerObservation {
	if in == nil {
		return nil
	}
	out := new(ManufacturerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManufacturerParameters) DeepCopyInto(out *ManufacturerParameters) {
	*out = *in
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufacturerParameters.
func (in *ManufacturerParameters) DeepCopy() *ManufacturerParameters {
	if in == nil {
		return nil
	}
	out := new(ManufacturerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManufacturerSpec) DeepCopyInto(out *ManufacturerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufacturerSpec.
func (in *ManufacturerSpec) DeepCopy() *ManufacturerSpec {
	if in == nil {
		return nil
	}
	out := new(ManufacturerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManufacturerStatus) DeepCopyInto(out *ManufacturerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManufacturerStatus.
func (in *ManufacturerStatus) DeepCopy() *ManufacturerStatus {
	if in == nil {
		return nil
	}
	out := new(ManufacturerStatus)
	in.DeepCopyInto(out)
	return out
}