/*
Copyright 2021 TriggerMesh Inc.

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

// Code generated by injection-gen. DO NOT EDIT.

package googlecloudrepositoriessource

import (
	context "context"

	apissourcesv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	internalclientset "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset"
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/sources/v1alpha1"
	client "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client"
	factory "github.com/triggermesh/triggermesh/pkg/client/generated/injection/informers/factory"
	sourcesv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Sources().V1alpha1().GoogleCloudRepositoriesSources()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

func withDynamicInformer(ctx context.Context) context.Context {
	inf := &wrapper{client: client.Get(ctx)}
	return context.WithValue(ctx, Key{}, inf)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.GoogleCloudRepositoriesSourceInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/sources/v1alpha1.GoogleCloudRepositoriesSourceInformer from context.")
	}
	return untyped.(v1alpha1.GoogleCloudRepositoriesSourceInformer)
}

type wrapper struct {
	client internalclientset.Interface

	namespace string
}

var _ v1alpha1.GoogleCloudRepositoriesSourceInformer = (*wrapper)(nil)
var _ sourcesv1alpha1.GoogleCloudRepositoriesSourceLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apissourcesv1alpha1.GoogleCloudRepositoriesSource{}, 0, nil)
}

func (w *wrapper) Lister() sourcesv1alpha1.GoogleCloudRepositoriesSourceLister {
	return w
}

func (w *wrapper) GoogleCloudRepositoriesSources(namespace string) sourcesv1alpha1.GoogleCloudRepositoriesSourceNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apissourcesv1alpha1.GoogleCloudRepositoriesSource, err error) {
	lo, err := w.client.SourcesV1alpha1().GoogleCloudRepositoriesSources(w.namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: selector.String(),
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apissourcesv1alpha1.GoogleCloudRepositoriesSource, error) {
	return w.client.SourcesV1alpha1().GoogleCloudRepositoriesSources(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}
