package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DDNS is the Schema for the ddns API
type DDNS struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of DDNS
	// +required
	Spec DDNSSpec `json:"spec"`

	// status defines the observed state of DDNS
	// +optional
	Status DDNSStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// DDNSList contains a list of DDNS
type DDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []DDNS `json:"items"`
}

// DDNSSpec defines the desired state of DDNS
type DDNSSpec struct {
}

// DDNSStatus defines the observed state of DDNS.
type DDNSStatus struct {
	// conditions represent the current state of the DDNS resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
