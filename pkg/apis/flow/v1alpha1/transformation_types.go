/*
Copyright 2022 TriggerMesh Inc.

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

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Transformation is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
type Transformation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransformationSpec `json:"spec,omitempty"`
	Status TargetStatus       `json:"status,omitempty"`
}

var (
	_ apis.Validatable = (*Transformation)(nil)
	_ apis.Defaultable = (*Transformation)(nil)

	_ Reconcilable = (*Transformation)(nil)
	_ EventSender  = (*Transformation)(nil)
)

// TransformationSpec holds the desired state of the Transformation (from the client).
type TransformationSpec struct {
	// Context contains Transformations that must be applied on CE Context
	Context []Transform `json:"context,omitempty"`
	// Data contains Transformations that must be applied on CE Data
	Data []Transform `json:"data,omitempty"`

	// Support sending to an event sink instead of replying.
	duckv1.SourceSpec `json:",inline"`
}

// Transform describes transformation schemes for different CE types.
type Transform struct {
	Operation string `json:"operation"`
	Paths     []Path `json:"paths"`
}

// Path is a key-value pair that represents JSON object path
type Path struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TransformationList is a list of Transformation resources
type TransformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Transformation `json:"items"`
}
