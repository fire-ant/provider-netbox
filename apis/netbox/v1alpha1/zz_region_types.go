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

type RegionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ParentRegionID *float64 `json:"parentRegionId,omitempty" tf:"parent_region_id,omitempty"`

	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

type RegionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ParentRegionID *float64 `json:"parentRegionId,omitempty" tf:"parent_region_id,omitempty"`

	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

// RegionSpec defines the desired state of Region
type RegionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionParameters `json:"forProvider"`
}

// RegionStatus defines the observed state of Region.
type RegionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Region is the Schema for the Regions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Region struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionSpec   `json:"spec"`
	Status            RegionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionList contains a list of Regions
type RegionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Region `json:"items"`
}

// Repository type metadata.
var (
	Region_Kind             = "Region"
	Region_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Region_Kind}.String()
	Region_KindAPIVersion   = Region_Kind + "." + CRDGroupVersion.String()
	Region_GroupVersionKind = CRDGroupVersion.WithKind(Region_Kind)
)

func init() {
	SchemeBuilder.Register(&Region{}, &RegionList{})
}
