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

package testing

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"

	fakek8sinjectionclient "knative.dev/pkg/client/injection/kube/client/fake"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	logtesting "knative.dev/pkg/logging/testing"
	"knative.dev/pkg/reconciler"
	rt "knative.dev/pkg/reconciler/testing"
	fakeservinginjectionclient "knative.dev/serving/pkg/client/injection/client/fake"

	fakeinjectionclient "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client/fake"
	"github.com/triggermesh/triggermesh/pkg/targets/reconciler/common"
)

// Ctor constructs a controller.Reconciler.
type Ctor func(*testing.T, context.Context, *rt.TableRow, *Listers) controller.Reconciler

// MakeFactory creates a testing factory for our controller.Reconciler, and
// initializes a Reconciler using the given Ctor as part of the process.
func MakeFactory(ctor Ctor) rt.Factory {
	return func(t *testing.T, tr *rt.TableRow) (controller.Reconciler, rt.ActionRecorderList, rt.EventList) {
		scheme := NewScheme()

		ls := NewListers(scheme, tr.Objects)

		// enable values injected by the test case (TableRow) to be consumed in ctor
		ctx := tr.Ctx
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = logging.WithLogger(ctx, logtesting.TestLogger(t))

		// the controller.Reconciler uses an internal client to handle
		// target objects
		ctx, client := fakeinjectionclient.With(ctx, ls.GetTargetsObjects()...)

		// all clients used inside reconciler implementations should be
		// injected as well
		ctx, k8sClient := fakek8sinjectionclient.With(ctx, ls.GetKubeObjects()...)
		ctx, servingClient := fakeservinginjectionclient.With(ctx, ls.GetServingObjects()...)

		const eventRecorderBufferSize = 10
		eventRecorder := record.NewFakeRecorder(eventRecorderBufferSize)
		ctx = controller.WithEventRecorder(ctx, eventRecorder)

		// set up Reconciler from fakes
		r := ctor(t, ctx, tr, &ls)

		// promote the reconciler if it is leader aware
		if la, ok := r.(reconciler.LeaderAware); ok {
			err := la.Promote(reconciler.UniversalBucket(), func(reconciler.Bucket, types.NamespacedName) {})
			if err != nil {
				t.Fatalf("Failed to promote reconciler to leader: %s", err)
			}
		}

		// inject reactors from table row
		for _, reactor := range tr.WithReactors {
			client.PrependReactor("*", "*", reactor)
			k8sClient.PrependReactor("*", "*", reactor)
			servingClient.PrependReactor("*", "*", reactor)
		}

		actionRecorderList := rt.ActionRecorderList{
			client, // record status updates
			k8sClient,
			servingClient,
		}

		eventList := rt.EventList{
			Recorder: eventRecorder,
		}

		return r, actionRecorderList, eventList
	}
}

// NewTestServiceReconciler returns a GenericServiceReconciler initialized with
// test clients.
func NewTestServiceReconciler(ctx context.Context, ls *Listers) common.GenericServiceReconciler {
	return common.GenericServiceReconciler{
		Lister:                ls.GetServiceLister().Services,
		Client:                fakeservinginjectionclient.Get(ctx).ServingV1().Services,
		GenericRBACReconciler: newTestRBACReconciler(ctx, ls),
	}
}

// newTestRBACReconciler returns a GenericRBACReconciler initialized with test clients.
func newTestRBACReconciler(ctx context.Context, ls *Listers) *common.GenericRBACReconciler {
	return &common.GenericRBACReconciler{
		SALister: ls.GetServiceAccountLister().ServiceAccounts,
		RBLister: ls.GetRoleBindingLister().RoleBindings,
		SAClient: fakek8sinjectionclient.Get(ctx).CoreV1().ServiceAccounts,
		RBClient: fakek8sinjectionclient.Get(ctx).RbacV1().RoleBindings,
	}
}
