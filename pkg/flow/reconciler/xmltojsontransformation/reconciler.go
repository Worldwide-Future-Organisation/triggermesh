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

package xmltojsontransformation

import (
	"context"

	"knative.dev/pkg/reconciler"

	"github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/flow/v1alpha1/xmltojsontransformation"
	listersv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/flow/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/flow/reconciler/common"
)

// Reconciler implements controller.Reconciler for the event target type.
type Reconciler struct {
	base       common.GenericServiceReconciler
	adapterCfg *adapterConfig

	trgLister func(namespace string) listersv1alpha1.XMLToJSONTransformationNamespaceLister
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, trg *v1alpha1.XMLToJSONTransformation) reconciler.Event {
	// inject target into context for usage in reconciliation logic
	ctx = v1alpha1.WithReconcilable(ctx, trg)

	return r.base.ReconcileAdapter(ctx, r)
}
