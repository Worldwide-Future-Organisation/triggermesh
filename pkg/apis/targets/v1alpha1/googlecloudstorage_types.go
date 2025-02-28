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

// GoogleCloudStorageTarget is the Schema for an Google Storage Target.
type GoogleCloudStorageTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleCloudStorageTargetSpec `json:"spec"`
	Status TargetStatus                 `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ Reconcilable              = (*GoogleCloudStorageTarget)(nil)
	_ targets.IntegrationTarget = (*GoogleCloudStorageTarget)(nil)
	_ targets.EventSource       = (*GoogleCloudStorageTarget)(nil)
)

// GoogleCloudStorageTargetSpec holds the desired state of the GoogleCloudStorageTarget.
type GoogleCloudStorageTargetSpec struct {
	// Credentials represents how Google Storage credentials should be provided in the secret
	Credentials SecretValueFromSource `json:"credentialsJson"`

	// BucketName specifies the Google Storage Bucket
	BucketName string `json:"bucketName"`

	// Whether to omit CloudEvent context attributes in objects created in Google Cloud Storage.
	// When this property is false (default), the entire CloudEvent payload is included.
	// When this property is true, only the CloudEvent data is included.
	DiscardCEContext bool `json:"discardCloudEventContext"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudStorageTargetList is a list of GoogleCloudStorageTarget resources
type GoogleCloudStorageTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GoogleCloudStorageTarget `json:"items"`
}
