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
	"time"

	"github.com/kelseyhightower/envconfig"

	"knative.dev/eventing/pkg/reconciler/source"
	k8sclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	informerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/informers/sources/v1alpha1/awssnssource"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/awssnssource"
	"github.com/triggermesh/triggermesh/pkg/sources/client/sns"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common"
)

// the resync period ensures we regularly re-check the state of SNS subscriptions.
const informerResyncPeriod = time.Minute * 5

// NewController creates a Reconciler for the event source and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {

	typ := (*v1alpha1.AWSSNSSource)(nil)
	app := common.ComponentName(typ)

	// Calling envconfig.Process() with a prefix appends that prefix
	// (uppercased) to the Go field name, e.g. MYSOURCE_IMAGE.
	adapterCfg := &adapterConfig{
		configs: source.WatchConfigurations(ctx, app, cmw, source.WithLogging, source.WithMetrics),
	}
	envconfig.MustProcess(app, adapterCfg)

	informer := informerv1alpha1.Get(ctx)

	r := &Reconciler{
		adapterCfg: adapterCfg,
		srcLister:  informer.Lister().AWSSNSSources,
		snsCg:      sns.NewClientGetter(k8sclient.Get(ctx).CoreV1().Secrets),
	}
	impl := reconcilerv1alpha1.NewImpl(ctx, r)

	logger := logging.FromContext(ctx)

	r.base = common.NewMTGenericServiceReconciler(
		ctx,
		typ,
		impl.Tracker,
		common.EnqueueObjectsInNamespaceOf(informer.Informer(), impl.FilteredGlobalResync, logger),
	)

	informer.Informer().AddEventHandlerWithResyncPeriod(controller.HandleAll(impl.Enqueue), informerResyncPeriod)

	return impl
}
