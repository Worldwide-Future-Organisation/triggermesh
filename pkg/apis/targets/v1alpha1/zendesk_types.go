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

// ZendeskTarget is the Schema for an Zendesk Target.
type ZendeskTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZendeskTargetSpec `json:"spec"`
	Status TargetStatus      `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ Reconcilable              = (*ZendeskTarget)(nil)
	_ targets.IntegrationTarget = (*ZendeskTarget)(nil)
	_ targets.EventSource       = (*ZendeskTarget)(nil)
)

// ZendeskTargetSpec holds the desired state of the ZendeskTarget.
type ZendeskTargetSpec struct {
	// Token contains the Zendesk account Token.
	Token SecretValueFromSource `json:"token"`

	// Subdomain the Zendesk subdomain.
	Subdomain string `json:"subdomain"`

	// Email the registered Zendesk email account.
	Email string `json:"email"`

	// Subject a static subject assignemnt for every ticket.
	// +optional
	Subject string `json:"subject,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ZendeskTargetList is a list of ZendeskTarget resources
type ZendeskTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ZendeskTarget `json:"items"`
}
