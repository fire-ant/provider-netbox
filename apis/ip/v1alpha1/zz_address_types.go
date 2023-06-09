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

type AddressObservation struct {
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	InterfaceID *float64 `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Defaults to `virtualization.vminterface`.
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type AddressParameters struct {

	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.VirtInterface
	// +crossplane:generate:reference:extractor=github.com/fire-ant/provider-netbox/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InterfaceID *float64 `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Reference to a VirtInterface in netbox to populate interfaceId.
	// +kubebuilder:validation:Optional
	InterfaceIDRef *v1.Reference `json:"interfaceIdRef,omitempty" tf:"-"`

	// Selector for a VirtInterface in netbox to populate interfaceId.
	// +kubebuilder:validation:Optional
	InterfaceIDSelector *v1.Selector `json:"interfaceIdSelector,omitempty" tf:"-"`

	// Defaults to `virtualization.vminterface`.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vrf
	// +crossplane:generate:reference:extractor=github.com/fire-ant/provider-netbox/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// Reference to a Vrf in netbox to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// Selector for a Vrf in netbox to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`
}

// AddressSpec defines the desired state of Address
type AddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressParameters `json:"forProvider"`
}

// AddressStatus defines the observed state of Address.
type AddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Address is the Schema for the Addresss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ipAddress)",message="ipAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.status)",message="status is a required parameter"
	Spec   AddressSpec   `json:"spec"`
	Status AddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressList contains a list of Addresss
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Address `json:"items"`
}

// Repository type metadata.
var (
	Address_Kind             = "Address"
	Address_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Address_Kind}.String()
	Address_KindAPIVersion   = Address_Kind + "." + CRDGroupVersion.String()
	Address_GroupVersionKind = CRDGroupVersion.WithKind(Address_Kind)
)

func init() {
	SchemeBuilder.Register(&Address{}, &AddressList{})
}
