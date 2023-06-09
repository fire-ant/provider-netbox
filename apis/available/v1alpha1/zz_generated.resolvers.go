/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/fire-ant/provider-netbox/apis/ip/v1alpha1"
	v1alpha12 "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1"
	v1alpha11 "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1"
	common "github.com/fire-ant/provider-netbox/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IPAddress.
func (mg *IPAddress) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.IPRangeID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IPRangeIDRef,
		Selector:     mg.Spec.ForProvider.IPRangeIDSelector,
		To: reference.To{
			List:    &v1alpha1.IPRangeList{},
			Managed: &v1alpha1.IPRange{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPRangeID")
	}
	mg.Spec.ForProvider.IPRangeID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPRangeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.InterfaceID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InterfaceIDRef,
		Selector:     mg.Spec.ForProvider.InterfaceIDSelector,
		To: reference.To{
			List:    &v1alpha11.VrfList{},
			Managed: &v1alpha11.Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InterfaceID")
	}
	mg.Spec.ForProvider.InterfaceID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InterfaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.PrefixID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrefixIDRef,
		Selector:     mg.Spec.ForProvider.PrefixIDSelector,
		To: reference.To{
			List:    &v1alpha11.PrefixList{},
			Managed: &v1alpha11.Prefix{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrefixID")
	}
	mg.Spec.ForProvider.PrefixID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrefixIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha11.TenantList{},
			Managed: &v1alpha11.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Prefix.
func (mg *Prefix) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &v1alpha12.RoleList{},
			Managed: &v1alpha12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &v1alpha11.SiteList{},
			Managed: &v1alpha11.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha11.TenantList{},
			Managed: &v1alpha11.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VlanID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VlanIDRef,
		Selector:     mg.Spec.ForProvider.VlanIDSelector,
		To: reference.To{
			List:    &v1alpha11.VlanList{},
			Managed: &v1alpha11.Vlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VlanID")
	}
	mg.Spec.ForProvider.VlanID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VlanIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &v1alpha11.VrfList{},
			Managed: &v1alpha11.Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}
