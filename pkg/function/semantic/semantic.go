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

package semantic

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"

	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

// Semantic can do semantic deep equality checks for Kubernetes API objects.
//
// For a given comparison function
//
//   comp(a, b interface{})
//
// 'a' should always be the desired state, and 'b' the current state for
// DeepDerivative comparisons to work as expected.
var Semantic = conversion.EqualitiesOrDie(
	knServiceEqual,
)

// eq is an instance of Equalities for internal deep derivative comparisons
// of API objects. Adapted from "k8s.io/apimachinery/pkg/api/equality".Semantic.
var eq = conversion.EqualitiesOrDie(
	func(a, b resource.Quantity) bool {
		if a.IsZero() {
			return true
		}
		return a.Cmp(b) == 0
	},
	func(a, b metav1.Time) bool { // e.g. metadata.creationTimestamp
		if a.IsZero() {
			return true
		}
		return a.UTC() == b.UTC()
	},
	func(a, b int64) bool { // e.g. metadata.generation
		if a == 0 {
			return true
		}
		return a == b
	},
	// Needed because DeepDerivative compares int values directly, which
	// doesn't yield the expected result with defaulted int32 probe fields.
	func(a, b *corev1.Probe) bool {
		if a == nil {
			return true
		}
		if b == nil {
			return false
		}

		if a.InitialDelaySeconds != 0 && a.InitialDelaySeconds != b.InitialDelaySeconds {
			return false
		}
		if a.TimeoutSeconds != 0 && a.TimeoutSeconds != b.TimeoutSeconds {
			return false
		}
		if a.PeriodSeconds != 0 && a.PeriodSeconds != b.PeriodSeconds {
			return false
		}
		if a.SuccessThreshold != 0 && a.SuccessThreshold != b.SuccessThreshold {
			return false
		}
		if a.FailureThreshold != 0 && a.FailureThreshold != b.FailureThreshold {
			return false
		}

		return (conversion.Equalities{}).DeepDerivative(a.Handler, b.Handler)
	},
)

// knServiceEqual returns whether two Knative Services are semantically equivalent.
func knServiceEqual(a, b *servingv1.Service) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if !eq.DeepDerivative(&a.ObjectMeta, &b.ObjectMeta) {
		return false
	}

	if !eq.DeepDerivative(&a.Spec.Template, &b.Spec.Template) {
		return false
	}

	return true
}
