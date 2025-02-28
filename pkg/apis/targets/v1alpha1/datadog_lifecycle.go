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
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// Managed event types
const (
	EventTypeDatadogMetric = "io.triggermesh.datadog.metric.submit"
	EventTypeDatadogEvent  = "io.triggermesh.datadog.event.post"
	EventTypeDatadogLog    = "io.triggermesh.datadog.log.send"

	EventTypeDatadogResponse = "io.triggermesh.datadog.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*DatadogTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("DatadogTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*DatadogTarget) GetConditionSet() apis.ConditionSet {
	return targetConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *DatadogTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *DatadogTarget) GetStatusManager() *StatusManager {
	return &StatusManager{
		ConditionSet: t.GetConditionSet(),
		TargetStatus: &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*DatadogTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeDatadogMetric,
		EventTypeDatadogEvent,
		EventTypeDatadogLog,
	}
}

// GetEventTypes implements EventSource.
func (*DatadogTarget) GetEventTypes() []string {
	return []string{
		EventTypeDatadogResponse,
	}
}

// AsEventSource implements EventSource.
func (t *DatadogTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "io.triggermesh." + kind + "." + t.Namespace + "." + t.Name
}
