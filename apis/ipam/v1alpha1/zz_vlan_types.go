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

type VlanObservation struct {

	// Defaults to `""`.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Defaults to `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	Vid *float64 `json:"vid,omitempty" tf:"vid,omitempty"`
}

type VlanParameters struct {

	// Defaults to `""`.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=Group
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

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

	// Defaults to `active`.
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

	// +kubebuilder:validation:Optional
	Vid *float64 `json:"vid,omitempty" tf:"vid,omitempty"`
}

// VlanSpec defines the desired state of Vlan
type VlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VlanParameters `json:"forProvider"`
}

// VlanStatus defines the observed state of Vlan.
type VlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vlan is the Schema for the Vlans API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Vlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vid)",message="vid is a required parameter"
	Spec   VlanSpec   `json:"spec"`
	Status VlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VlanList contains a list of Vlans
type VlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vlan `json:"items"`
}

// Repository type metadata.
var (
	Vlan_Kind             = "Vlan"
	Vlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vlan_Kind}.String()
	Vlan_KindAPIVersion   = Vlan_Kind + "." + CRDGroupVersion.String()
	Vlan_GroupVersionKind = CRDGroupVersion.WithKind(Vlan_Kind)
)

func init() {
	SchemeBuilder.Register(&Vlan{}, &VlanList{})
}
