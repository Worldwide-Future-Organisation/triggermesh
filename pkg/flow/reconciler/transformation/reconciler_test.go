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

package transformation

import (
	"context"
	"testing"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"

	"github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	fakeinjectionclient "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/flow/v1alpha1/transformation"
	"github.com/triggermesh/triggermesh/pkg/flow/reconciler/common"
	. "github.com/triggermesh/triggermesh/pkg/flow/reconciler/testing"
)

func TestReconcile(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:     "registry/image:tag",
		obsConfig: &source.EmptyVarsGenerator{},
	}

	ctor := reconcilerCtor(adapterCfg)
	trg := newTarget()
	ab := adapterBuilder(adapterCfg)

	TestReconcileAdapter(t, ctor, trg, ab)
}

// reconcilerCtor returns a Ctor for a Transformation Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			base:       NewTestServiceReconciler(ctx, ls),
			adapterCfg: cfg,
			trgLister:  ls.GetTransformationLister().Transformations,
		}

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetTransformationLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newTarget returns a populated target object.
func newTarget() *v1alpha1.Transformation {
	trg := &v1alpha1.Transformation{
		Spec: v1alpha1.TransformationSpec{
			Context: []v1alpha1.Transform{
				{
					Operation: "store",
					Paths: []v1alpha1.Path{
						{
							Key:   "$time",
							Value: "time",
						},
					},
				},
			},
			Data: []v1alpha1.Transform{
				{
					Operation: "add",
					Paths: []v1alpha1.Path{
						{
							Key:   "event.time",
							Value: "$time",
						},
					},
				},
			},
		},
	}

	Populate(trg)

	return trg
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterServiceBuilder {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
