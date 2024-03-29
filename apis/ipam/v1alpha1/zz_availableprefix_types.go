/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AvailablePrefixObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsPool *bool `json:"isPool,omitempty" tf:"is_pool,omitempty"`

	MarkUtilized *bool `json:"markUtilized,omitempty" tf:"mark_utilized,omitempty"`

	ParentPrefixID *float64 `json:"parentPrefixId,omitempty" tf:"parent_prefix_id,omitempty"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type AvailablePrefixParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IsPool *bool `json:"isPool,omitempty" tf:"is_pool,omitempty"`

	// +kubebuilder:validation:Optional
	MarkUtilized *bool `json:"markUtilized,omitempty" tf:"mark_utilized,omitempty"`

	// +crossplane:generate:reference:type=Prefix
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ParentPrefixID *float64 `json:"parentPrefixId,omitempty" tf:"parent_prefix_id,omitempty"`

	// Reference to a Prefix to populate parentPrefixId.
	// +kubebuilder:validation:Optional
	ParentPrefixIDRef *v1.Reference `json:"parentPrefixIdRef,omitempty" tf:"-"`

	// Selector for a Prefix to populate parentPrefixId.
	// +kubebuilder:validation:Optional
	ParentPrefixIDSelector *v1.Selector `json:"parentPrefixIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// +crossplane:generate:reference:type=IpamRole
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a IpamRole to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a IpamRole to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Site
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Reference to a Site in dcim to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDRef *v1.Reference `json:"siteIdRef,omitempty" tf:"-"`

	// Selector for a Site in dcim to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDSelector *v1.Selector `json:"siteIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Tenant in tenant to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Tenant in tenant to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Vlan
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// Reference to a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Vrf
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// Reference to a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// Selector for a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`
}

// AvailablePrefixSpec defines the desired state of AvailablePrefix
type AvailablePrefixSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AvailablePrefixParameters `json:"forProvider"`
}

// AvailablePrefixStatus defines the observed state of AvailablePrefix.
type AvailablePrefixStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AvailablePrefixObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AvailablePrefix is the Schema for the AvailablePrefixs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type AvailablePrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.prefixLength)",message="prefixLength is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.status)",message="status is a required parameter"
	Spec   AvailablePrefixSpec   `json:"spec"`
	Status AvailablePrefixStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailablePrefixList contains a list of AvailablePrefixs
type AvailablePrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailablePrefix `json:"items"`
}

// Repository type metadata.
var (
	AvailablePrefix_Kind             = "AvailablePrefix"
	AvailablePrefix_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AvailablePrefix_Kind}.String()
	AvailablePrefix_KindAPIVersion   = AvailablePrefix_Kind + "." + CRDGroupVersion.String()
	AvailablePrefix_GroupVersionKind = CRDGroupVersion.WithKind(AvailablePrefix_Kind)
)

func init() {
	SchemeBuilder.Register(&AvailablePrefix{}, &AvailablePrefixList{})
}
