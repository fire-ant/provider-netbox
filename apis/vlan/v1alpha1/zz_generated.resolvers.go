/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1"
	common "github.com/fire-ant/provider-netbox/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ScopeID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ScopeIDRef,
		Selector:     mg.Spec.ForProvider.ScopeIDSelector,
		To: reference.To{
			List:    &v1alpha1.SiteList{},
			Managed: &v1alpha1.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScopeID")
	}
	mg.Spec.ForProvider.ScopeID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeIDRef = rsp.ResolvedReference

	return nil
}
