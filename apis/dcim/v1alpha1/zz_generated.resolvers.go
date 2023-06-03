/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1"
	common "github.com/fire-ant/provider-netbox/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DeviceInterface.
func (mg *DeviceInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeviceType.
func (mg *DeviceType) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ManufacturerID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ManufacturerIDRef,
		Selector:     mg.Spec.ForProvider.ManufacturerIDSelector,
		To: reference.To{
			List:    &ManufacturerList{},
			Managed: &Manufacturer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManufacturerID")
	}
	mg.Spec.ForProvider.ManufacturerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManufacturerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Location.
func (mg *Location) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &SiteList{},
			Managed: &Site{},
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
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}
