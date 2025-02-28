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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common/resource"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (s *AWSSQSSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AWSSQSSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *AWSSQSSource) GetConditionSet() apis.ConditionSet {
	return eventSourceConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *AWSSQSSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements Reconcilable.
func (s *AWSSQSSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *AWSSQSSource) GetStatusManager() *StatusManager {
	return &StatusManager{
		ConditionSet:      s.GetConditionSet(),
		EventSourceStatus: &s.Status,
	}
}

// Supported event types
const (
	AWSSQSGenericEventType = "message"
)

// GetEventTypes implements Reconcilable.
func (s *AWSSQSSource) GetEventTypes() []string {
	return []string{
		AWSEventType(s.Spec.ARN.Service, AWSSQSGenericEventType),
	}
}

// AsEventSource implements Reconcilable.
func (s *AWSSQSSource) AsEventSource() string {
	return s.Spec.ARN.String()
}

// WantsOwnServiceAccount implements serviceAccountProvider.
func (s *AWSSQSSource) WantsOwnServiceAccount() bool {
	return s.Spec.Auth.EksIAMRole != nil
}

// ServiceAccountOptions implements serviceAccountProvider.
func (s *AWSSQSSource) ServiceAccountOptions() []resource.ServiceAccountOption {
	var saOpts []resource.ServiceAccountOption

	if iamRole := s.Spec.Auth.EksIAMRole; iamRole != nil {
		setIAMRoleAnnotation := func(sa *corev1.ServiceAccount) {
			metav1.SetMetaDataAnnotation(&sa.ObjectMeta, annotationEksIAMRole, iamRole.String())
		}

		saOpts = append(saOpts, setIAMRoleAnnotation)
	}

	return saOpts
}
