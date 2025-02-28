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

package awscloudwatchsource

import (
	"context"
	"testing"
	"time"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"

	"github.com/triggermesh/triggermesh/pkg/apis"
	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	fakeinjectionclient "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/awscloudwatchsource"
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

// reconcilerCtor returns a Ctor for a AWSCloudWatchSource Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			base:       NewTestDeploymentReconciler(ctx, ls),
			adapterCfg: cfg,
			srcLister:  ls.GetAWSCloudWatchSourceLister().AWSCloudWatchSources,
		}

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetAWSCloudWatchSourceLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newEventSource returns a populated source object.
func newEventSource() *v1alpha1.AWSCloudWatchSource {
	pollingInterval := apis.Duration(5 * time.Minute)

	src := &v1alpha1.AWSCloudWatchSource{
		Spec: v1alpha1.AWSCloudWatchSourceSpec{
			Region: "us-west-2",
			MetricQueries: []v1alpha1.AWSCloudWatchMetricQuery{{
				Name:       "testquery",
				Expression: nil,
				Metric: &v1alpha1.AWSCloudWatchMetricStat{
					Metric: v1alpha1.AWSCloudWatchMetric{
						Dimensions: []v1alpha1.AWSCloudWatchMetricDimension{{
							Name:  "FunctionName",
							Value: "makemoney",
						}},
						MetricName: "Duration",
						Namespace:  "AWS/Lambda",
					},
					Period: 60,
					Stat:   "sum",
					Unit:   "seconds",
				},
			}},
			PollingInterval: &pollingInterval,
		},
		Status: v1alpha1.EventSourceStatus{},
	}

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
