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

type TargetObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type TargetParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Tenant
	// +crossplane:generate:reference:extractor=github.com/fire-ant/provider-netbox/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Tenant in netbox to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Tenant in netbox to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`
}

// TargetSpec defines the desired state of Target
type TargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetParameters `json:"forProvider"`
}

// TargetStatus defines the observed state of Target.
type TargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Target is the Schema for the Targets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetSpec   `json:"spec"`
	Status            TargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetList contains a list of Targets
type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target `json:"items"`
}

// Repository type metadata.
var (
	Target_Kind             = "Target"
	Target_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Target_Kind}.String()
	Target_KindAPIVersion   = Target_Kind + "." + CRDGroupVersion.String()
	Target_GroupVersionKind = CRDGroupVersion.WithKind(Target_Kind)
)

func init() {
	SchemeBuilder.Register(&Target{}, &TargetList{})
}
