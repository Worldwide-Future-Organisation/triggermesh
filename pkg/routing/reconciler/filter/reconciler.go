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

package filter

import (
	"context"

	"knative.dev/pkg/reconciler"

	"github.com/triggermesh/triggermesh/pkg/apis/routing/v1alpha1"
	filterreconciler "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/routing/v1alpha1/filter"
	listersv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/routing/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/routing/reconciler/common"
)

// Reconciler implements addressableservicereconciler.Interface for
// AddressableService resources.
type Reconciler struct {
	base       common.GenericServiceReconciler
	rtrLister  func(namespace string) listersv1alpha1.FilterNamespaceLister
	adapterCfg *adapterConfig
}

// Check that our Reconciler implements Interface
var _ filterreconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.Filter) reconciler.Event {
	// inject source into context for usage in reconciliation logic
	ctx = v1alpha1.WithReconcilable(ctx, o)

	return r.base.ReconcileAdapter(ctx, r)
}
