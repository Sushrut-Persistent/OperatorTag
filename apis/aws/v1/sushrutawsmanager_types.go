/*
Copyright 2023.

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

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SushrutAWSManagerSpec defines the desired state of SushrutAWSManager
type SushrutAWSManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SushrutAWSManager. Edit sushrutawsmanager_types.go to remove/update
	Image           string           `json:"image,omitempty"`
	ImagePullPolicy v1.PullPolicy    `json:"imagePullPolicy,omitempty"`
	RestartPolicy   v1.RestartPolicy `json:"restartPolicy,omitempty"`
}

// SushrutAWSManagerStatus defines the observed state of SushrutAWSManager
type SushrutAWSManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SushrutAWSManager is the Schema for the sushrutawsmanagers API
type SushrutAWSManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SushrutAWSManagerSpec   `json:"spec,omitempty"`
	Status SushrutAWSManagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SushrutAWSManagerList contains a list of SushrutAWSManager
type SushrutAWSManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SushrutAWSManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SushrutAWSManager{}, &SushrutAWSManagerList{})
}
