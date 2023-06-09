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

type ClusterTypeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

type ClusterTypeParameters struct {

	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

// ClusterTypeSpec defines the desired state of ClusterType
type ClusterTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterTypeParameters `json:"forProvider"`
}

// ClusterTypeStatus defines the observed state of ClusterType.
type ClusterTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterType is the Schema for the ClusterTypes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type ClusterType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterTypeSpec   `json:"spec"`
	Status            ClusterTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterTypeList contains a list of ClusterTypes
type ClusterTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterType `json:"items"`
}

// Repository type metadata.
var (
	ClusterType_Kind             = "ClusterType"
	ClusterType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterType_Kind}.String()
	ClusterType_KindAPIVersion   = ClusterType_Kind + "." + CRDGroupVersion.String()
	ClusterType_GroupVersionKind = CRDGroupVersion.WithKind(ClusterType_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterType{}, &ClusterTypeList{})
}
