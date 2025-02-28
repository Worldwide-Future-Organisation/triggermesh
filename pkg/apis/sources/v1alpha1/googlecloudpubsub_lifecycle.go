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
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (s *GoogleCloudPubSubSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleCloudPubSubSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *GoogleCloudPubSubSource) GetConditionSet() apis.ConditionSet {
	return googleCloudPubSubSourceConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *GoogleCloudPubSubSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements Reconcilable.
func (s *GoogleCloudPubSubSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *GoogleCloudPubSubSource) GetStatusManager() *StatusManager {
	return &StatusManager{
		ConditionSet:      s.GetConditionSet(),
		EventSourceStatus: &s.Status.EventSourceStatus,
	}
}

// AsEventSource implements Reconcilable.
func (s *GoogleCloudPubSubSource) AsEventSource() string {
	return s.Spec.Topic.String()
}

// Supported event types
const (
	GoogleCloudPubSubGenericEventType = "com.google.cloud.pubsub.message"
)

// GetEventTypes returns the event types generated by the source.
func (s *GoogleCloudPubSubSource) GetEventTypes() []string {
	return []string{
		GoogleCloudPubSubGenericEventType,
	}
}

// Status conditions
const (
	// GoogleCloudPubSubConditionSubscribed has status True when the source has subscribed to a topic.
	GoogleCloudPubSubConditionSubscribed apis.ConditionType = "Subscribed"
)

// googleCloudPubSubSourceConditionSet is a set of conditions for
// GoogleCloudPubSubSource objects.
var googleCloudPubSubSourceConditionSet = NewEventSourceConditionSet(
	GoogleCloudPubSubConditionSubscribed,
)

// MarkSubscribed sets the Subscribed condition to True.
func (s *GoogleCloudPubSubSourceStatus) MarkSubscribed() {
	googleCloudPubSubSourceConditionSet.Manage(s).MarkTrue(GoogleCloudPubSubConditionSubscribed)
}

// MarkNotSubscribed sets the Subscribed condition to False with the given
// reason and message.
func (s *GoogleCloudPubSubSourceStatus) MarkNotSubscribed(reason, msg string) {
	googleCloudPubSubSourceConditionSet.Manage(s).MarkFalse(GoogleCloudPubSubConditionSubscribed, reason, msg)
}
