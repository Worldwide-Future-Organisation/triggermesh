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

	"github.com/triggermesh/triggermesh/pkg/apis/targets"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonTarget defines the schema for the Tekton target.
type TektonTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   TektonTargetSpec `json:"spec"`
	Status TargetStatus     `json:"status,omitempty"`
}

// TektonTargetSpec holds the desired state of event target.
type TektonTargetSpec struct {
	// ReapPolicy dictates the reaping policy to be applied for the target
	// +optional
	ReapPolicy *TektonTargetReapPolicy `json:"reapPolicy,omitempty"`
}

// TektonTargetReapPolicy defines desired Repeating Policy.
type TektonTargetReapPolicy struct {
	// ReapSuccessAge How long to wait before reaping runs that were successful
	ReapSuccessAge *string `json:"success,omitempty"`
	// ReapFailAge How long to wait before reaping runs that failed
	ReapFailAge *string `json:"fail,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ Reconcilable              = (*TektonTarget)(nil)
	_ targets.IntegrationTarget = (*TektonTarget)(nil)
	_ targets.EventSource       = (*TektonTarget)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonTargetList is a list of event targets.
type TektonTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TektonTarget `json:"items"`
}
