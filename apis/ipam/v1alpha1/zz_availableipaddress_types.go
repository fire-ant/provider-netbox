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

type AvailableIPAddressObservation struct {
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	IPRangeID *string `json:"ipRangeId,omitempty" tf:"ip_range_id,omitempty"`

	InterfaceID *float64 `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	PrefixID *string `json:"prefixId,omitempty" tf:"prefix_id,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Defaults to `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type AvailableIPAddressParameters struct {

	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=IPRange
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IPRangeID *string `json:"ipRangeId,omitempty" tf:"ip_range_id,omitempty"`

	// Reference to a IPRange to populate ipRangeId.
	// +kubebuilder:validation:Optional
	IPRangeIDRef *v1.Reference `json:"ipRangeIdRef,omitempty" tf:"-"`

	// Selector for a IPRange to populate ipRangeId.
	// +kubebuilder:validation:Optional
	IPRangeIDSelector *v1.Selector `json:"ipRangeIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InterfaceID *float64 `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Reference to a DeviceInterface in dcim to populate interfaceId.
	// +kubebuilder:validation:Optional
	InterfaceIDRef *v1.Reference `json:"interfaceIdRef,omitempty" tf:"-"`

	// Selector for a DeviceInterface in dcim to populate interfaceId.
	// +kubebuilder:validation:Optional
	InterfaceIDSelector *v1.Selector `json:"interfaceIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Prefix
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrefixID *string `json:"prefixId,omitempty" tf:"prefix_id,omitempty"`

	// Reference to a Prefix to populate prefixId.
	// +kubebuilder:validation:Optional
	PrefixIDRef *v1.Reference `json:"prefixIdRef,omitempty" tf:"-"`

	// Selector for a Prefix to populate prefixId.
	// +kubebuilder:validation:Optional
	PrefixIDSelector *v1.Selector `json:"prefixIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

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

// AvailableIPAddressSpec defines the desired state of AvailableIPAddress
type AvailableIPAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AvailableIPAddressParameters `json:"forProvider"`
}

// AvailableIPAddressStatus defines the observed state of AvailableIPAddress.
type AvailableIPAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AvailableIPAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AvailableIPAddress is the Schema for the AvailableIPAddresss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type AvailableIPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AvailableIPAddressSpec   `json:"spec"`
	Status            AvailableIPAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailableIPAddressList contains a list of AvailableIPAddresss
type AvailableIPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailableIPAddress `json:"items"`
}

// Repository type metadata.
var (
	AvailableIPAddress_Kind             = "AvailableIPAddress"
	AvailableIPAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AvailableIPAddress_Kind}.String()
	AvailableIPAddress_KindAPIVersion   = AvailableIPAddress_Kind + "." + CRDGroupVersion.String()
	AvailableIPAddress_GroupVersionKind = CRDGroupVersion.WithKind(AvailableIPAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&AvailableIPAddress{}, &AvailableIPAddressList{})
}
