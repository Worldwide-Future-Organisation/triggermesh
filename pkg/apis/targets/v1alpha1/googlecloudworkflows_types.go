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

// GoogleCloudWorkflowsTarget is the Schema for an Google Cloud Workflows Target.
type GoogleCloudWorkflowsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleCloudWorkflowsTargetSpec `json:"spec"`
	Status TargetStatus                   `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ Reconcilable              = (*GoogleCloudWorkflowsTarget)(nil)
	_ targets.IntegrationTarget = (*GoogleCloudWorkflowsTarget)(nil)
	_ targets.EventSource       = (*GoogleCloudWorkflowsTarget)(nil)
)

// GoogleCloudWorkflowsTargetSpec holds the desired state of the GoogleCloudWorkflowsTarget.
type GoogleCloudWorkflowsTargetSpec struct {
	// GoogleCloudWorkflowsApiKey represents how GoogleCloudWorkflows credentials should be provided in the secret
	Credentials SecretValueFromSource `json:"credentialsJson"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudWorkflowsTargetList is a list of GoogleCloudWorkflowsTarget resources
type GoogleCloudWorkflowsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GoogleCloudWorkflowsTarget `json:"items"`
}
