/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CAProviderClassSpec defines the desired state of CAProviderClass
type CAProviderClassSpec struct {
	// Selects a key of secrets
	// +optional
	SecretRefs []SecretKeySelector `json:"secretRefs,omitempty"`
	// Selects issuers or cluster issuers
	// +optional
	IssuerRefs []TypedObjectReference `json:"issuerRefs,omitempty"`
	// Selects certificates
	// +optional
	CertificateRefs []ObjectReference `json:"certificateRefs,omitempty"`
}

// CAProviderClassStatus defines the observed state of CAProviderClass
type CAProviderClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CAProviderClass is the Schema for the caproviderclasses API
type CAProviderClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CAProviderClassSpec   `json:"spec,omitempty"`
	Status CAProviderClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CAProviderClassList contains a list of CAProviderClass
type CAProviderClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CAProviderClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CAProviderClass{}, &CAProviderClassList{})
}

// SecretKeySelector selects a key of a Secret.
// +structType=atomic
type SecretKeySelector struct {
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key"`
	// Specify whether the Secret or its key must be defined
	// +optional
	Optional *bool `json:"optional,omitempty"`
}

// TypedObjectReference contains enough information to let you locate the typed referenced object.
// +structType=atomic
type TypedObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
}

// ObjectReference contains enough information to let you locate the referenced objects of a known type.
// +structType=atomic
type ObjectReference struct {
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
}
