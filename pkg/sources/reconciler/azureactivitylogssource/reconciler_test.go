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

package azureactivitylogssource

import (
	"context"
	"testing"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"

	"github.com/triggermesh/triggermesh/pkg/apis/sources"
	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	fakeinjectionclient "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/azureactivitylogssource"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common"
	. "github.com/triggermesh/triggermesh/pkg/sources/reconciler/testing"
)

func TestReconcileSource(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:   "registry/image:tag",
		configs: &source.EmptyVarsGenerator{},
	}

	ctor := reconcilerCtor(adapterCfg)
	src := newEventSource()
	ab := adapterBuilder(adapterCfg)

	TestReconcileAdapter(t, ctor, src, ab)
}

// reconcilerCtor returns a Ctor for a source Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			base:       NewTestDeploymentReconciler(ctx, ls),
			adapterCfg: cfg,
			srcLister:  ls.GetAzureActivityLogsSourceLister().AzureActivityLogsSources,
		}

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetAzureActivityLogsSourceLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newEventSource returns a test source object with a minimal set of pre-filled attributes.
func newEventSource() *v1alpha1.AzureActivityLogsSource {
	src := &v1alpha1.AzureActivityLogsSource{
		Spec: v1alpha1.AzureActivityLogsSourceSpec{
			SubscriptionID: tSubscriptionID,
			Destination: v1alpha1.AzureActivityLogsSourceDestination{
				EventHubs: v1alpha1.AzureActivityLogsSourceDestinationEventHubs{
					NamespaceID: tEventHubNamespaceID,
				},
			},
			Categories: []string{"Administrative"},
			Auth: v1alpha1.AzureAuth{
				ServicePrincipal: &v1alpha1.AzureServicePrincipal{
					TenantID: v1alpha1.ValueFromField{
						Value: "00000000-0000-0000-0000-000000000000",
					},
					ClientID: v1alpha1.ValueFromField{
						Value: "00000000-0000-0000-0000-000000000000",
					},
					ClientSecret: v1alpha1.ValueFromField{
						Value: "some_secret",
					},
				},
			},
		},
	}

	// assume finalizer is already set to prevent the generated reconciler
	// from generating an extra Patch action
	src.Finalizers = []string{sources.AzureActivityLogsSourceResource.String()}

	Populate(src)

	return src
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterDeploymentBuilder {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
