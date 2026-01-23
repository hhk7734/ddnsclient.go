// Package v1alpha1 contains API Schema definitions for the networking v1alpha1 API group.
// +kubebuilder:object:generate=true
// +groupName=networking.loliot.net
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

func init() {
	SchemeBuilder.Register(&DDNS{}, &DDNSList{})
}

var (
	// GroupVersion is group version used to register these objects.
	GroupVersion = schema.GroupVersion{Group: "networking.loliot.net", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// DDNSGVK is the GroupVersionKind for DDNS.
	DDNSGVK = GroupVersion.WithKind("DDNS")

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
