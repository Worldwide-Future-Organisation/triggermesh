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
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type JQTransformation struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JQTransformationSpec `json:"spec"`
	Status TargetStatus         `json:"status,omitempty"`
}

var (
	_ Reconcilable = (*JQTransformation)(nil)
	_ EventSender  = (*JQTransformation)(nil)
)

// JQTransformationSpec holds the desired state of the JQTransformation (from the client).
type JQTransformationSpec struct {
	// The query that gets passed to the JQ library
	Query string `json:"query"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Support sending to an event sink instead of replying.
	duckv1.SourceSpec `json:",inline"`
}

// JQTransformationStatus communicates the observed state of the JQTransformation (from the controller).
type JQTransformationStatus struct {
	duckv1.SourceStatus  `json:",inline"`
	duckv1.AddressStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JQTransformationList is a list of JQTransformation resources
type JQTransformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []JQTransformation `json:"items"`
}
