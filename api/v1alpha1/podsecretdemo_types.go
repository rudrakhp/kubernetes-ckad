/*
Copyright 2022.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodSecretDemoSpec defines the desired state of PodSecretDemo
type PodSecretDemoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PodSecretDemo. Edit podsecretdemo_types.go to remove/update
	Secret string `json:"secret,omitempty"`
	Data   string `json:"data,omitempty"`
	Opt    string `json:"opt,omitempty"`
}

// PodSecretDemoStatus defines the observed state of PodSecretDemo
type PodSecretDemoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodSecretDemo is the Schema for the podsecretdemoes API
type PodSecretDemo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodSecretDemoSpec   `json:"spec,omitempty"`
	Status PodSecretDemoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodSecretDemoList contains a list of PodSecretDemo
type PodSecretDemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodSecretDemo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodSecretDemo{}, &PodSecretDemoList{})
}
