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
	EventTypeElasticsearchStore    = "io.triggermesh.elasticsearch.doc.index"
	EventTypeElasticsearchResponse = "io.triggermesh.elasticsearch.doc.index.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*ElasticsearchTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("ElasticsearchTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*ElasticsearchTarget) GetConditionSet() apis.ConditionSet {
	return targetConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *ElasticsearchTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *ElasticsearchTarget) GetStatusManager() *StatusManager {
	return &StatusManager{
		ConditionSet: t.GetConditionSet(),
		TargetStatus: &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*ElasticsearchTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeElasticsearchStore,
	}
}

// GetEventTypes implements EventSource.
func (*ElasticsearchTarget) GetEventTypes() []string {
	return []string{
		EventTypeElasticsearchResponse,
	}
}

// AsEventSource implements EventSource.
func (t *ElasticsearchTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "io.triggermesh." + kind + "." + t.Namespace + "." + t.Name
}
