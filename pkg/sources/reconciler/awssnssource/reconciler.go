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

package awssnssource

import (
	"context"
	"fmt"

	"knative.dev/pkg/reconciler"

	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/awssnssource"
	listersv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/sources/v1alpha1"
	snsclient "github.com/triggermesh/triggermesh/pkg/sources/client/sns"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common"
)

// Reconciler implements controller.Reconciler for the event source type.
type Reconciler struct {
	base       common.GenericServiceReconciler
	adapterCfg *adapterConfig

	srcLister func(namespace string) listersv1alpha1.AWSSNSSourceNamespaceLister

	// SNS client interface to interact with the SNS API
	snsCg snsclient.ClientGetter
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// Check that our Reconciler implements Finalizer
var _ reconcilerv1alpha1.Finalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, src *v1alpha1.AWSSNSSource) reconciler.Event {
	// inject source into context for usage in reconciliation logic
	ctx = v1alpha1.WithReconcilable(ctx, src)

	if err := r.base.ReconcileAdapter(ctx, r); err != nil {
		return fmt.Errorf("failed to reconcile source: %w", err)
	}

	return r.ensureSubscribed(ctx)
}

// FinalizeKind is called when the resource is deleted.
func (r *Reconciler) FinalizeKind(ctx context.Context, src *v1alpha1.AWSSNSSource) reconciler.Event {
	// inject source into context for usage in finalization logic
	ctx = v1alpha1.WithReconcilable(ctx, src)

	// The finalizer blocks the deletion of the source object until
	// ensureUnsubscribed succeeds to ensure that we don't leave any
	// dangling subscription behind us.
	return r.ensureUnsubscribed(ctx)
}
