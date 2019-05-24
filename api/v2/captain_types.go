/*
Copyright 2019 The Kubernetes authors.

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

package v2

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/play-kubebuilder/api/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CaptainSpec defines the desired state of Captain
type CaptainSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +optional
	NextRun int `json:"nextRun,omitempty"`
}

// CaptainStatus defines the observed state of Captain
type CaptainStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastRun int `json:"lastRun,omitempty"`
}

// +kubebuilder:object:root=true

// Captain is the Schema for the captains API
type Captain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaptainSpec   `json:"spec,omitempty"`
	Status CaptainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaptainList contains a list of Captain
type CaptainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Captain `json:"items"`
}

func (ej *Captain) ConvertTo(dst conversion.Hub) error {
	switch t := dst.(type) {
	case *v1.Captain:
		jobv1 := dst.(*v1.Captain)
		jobv1.ObjectMeta = ej.ObjectMeta
		jobv1.Spec.NextStop = ej.Spec.NextRun
		return nil
	default:
		return fmt.Errorf("unsupported type %v", t)
	}
}

func (ej *Captain) ConvertFrom(src conversion.Hub) error {
	switch t := src.(type) {
	case *v1.Captain:
		jobv1 := src.(*v1.Captain)
		ej.ObjectMeta = jobv1.ObjectMeta
		ej.Spec.NextRun = jobv1.Spec.NextStop
		return nil
	default:
		return fmt.Errorf("unsupported type %v", t)
	}
}

func init() {
	SchemeBuilder.Register(&Captain{}, &CaptainList{})
}
