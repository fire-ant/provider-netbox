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

type VrfObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type VrfParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

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
}

// VrfSpec defines the desired state of Vrf
type VrfSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VrfParameters `json:"forProvider"`
}

// VrfStatus defines the observed state of Vrf.
type VrfStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VrfObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vrf is the Schema for the Vrfs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Vrf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VrfSpec   `json:"spec"`
	Status            VrfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VrfList contains a list of Vrfs
type VrfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vrf `json:"items"`
}

// Repository type metadata.
var (
	Vrf_Kind             = "Vrf"
	Vrf_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vrf_Kind}.String()
	Vrf_KindAPIVersion   = Vrf_Kind + "." + CRDGroupVersion.String()
	Vrf_GroupVersionKind = CRDGroupVersion.WithKind(Vrf_Kind)
)

func init() {
	SchemeBuilder.Register(&Vrf{}, &VrfList{})
}
