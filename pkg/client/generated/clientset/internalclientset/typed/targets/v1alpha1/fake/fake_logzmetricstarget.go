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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLogzMetricsTargets implements LogzMetricsTargetInterface
type FakeLogzMetricsTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var logzmetricstargetsResource = schema.GroupVersionResource{Group: "targets.triggermesh.io", Version: "v1alpha1", Resource: "logzmetricstargets"}

var logzmetricstargetsKind = schema.GroupVersionKind{Group: "targets.triggermesh.io", Version: "v1alpha1", Kind: "LogzMetricsTarget"}

// Get takes name of the logzMetricsTarget, and returns the corresponding logzMetricsTarget object, and an error if there is any.
func (c *FakeLogzMetricsTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogzMetricsTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logzmetricstargetsResource, c.ns, name), &v1alpha1.LogzMetricsTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogzMetricsTarget), err
}

// List takes label and field selectors, and returns the list of LogzMetricsTargets that match those selectors.
func (c *FakeLogzMetricsTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogzMetricsTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logzmetricstargetsResource, logzmetricstargetsKind, c.ns, opts), &v1alpha1.LogzMetricsTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogzMetricsTargetList{ListMeta: obj.(*v1alpha1.LogzMetricsTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogzMetricsTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logzMetricsTargets.
func (c *FakeLogzMetricsTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logzmetricstargetsResource, c.ns, opts))

}

// Create takes the representation of a logzMetricsTarget and creates it.  Returns the server's representation of the logzMetricsTarget, and an error, if there is any.
func (c *FakeLogzMetricsTargets) Create(ctx context.Context, logzMetricsTarget *v1alpha1.LogzMetricsTarget, opts v1.CreateOptions) (result *v1alpha1.LogzMetricsTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logzmetricstargetsResource, c.ns, logzMetricsTarget), &v1alpha1.LogzMetricsTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogzMetricsTarget), err
}

// Update takes the representation of a logzMetricsTarget and updates it. Returns the server's representation of the logzMetricsTarget, and an error, if there is any.
func (c *FakeLogzMetricsTargets) Update(ctx context.Context, logzMetricsTarget *v1alpha1.LogzMetricsTarget, opts v1.UpdateOptions) (result *v1alpha1.LogzMetricsTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logzmetricstargetsResource, c.ns, logzMetricsTarget), &v1alpha1.LogzMetricsTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogzMetricsTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogzMetricsTargets) UpdateStatus(ctx context.Context, logzMetricsTarget *v1alpha1.LogzMetricsTarget, opts v1.UpdateOptions) (*v1alpha1.LogzMetricsTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logzmetricstargetsResource, "status", c.ns, logzMetricsTarget), &v1alpha1.LogzMetricsTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogzMetricsTarget), err
}

// Delete takes name of the logzMetricsTarget and deletes it. Returns an error if one occurs.
func (c *FakeLogzMetricsTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logzmetricstargetsResource, c.ns, name), &v1alpha1.LogzMetricsTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogzMetricsTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logzmetricstargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogzMetricsTargetList{})
	return err
}

// Patch applies the patch and returns the patched logzMetricsTarget.
func (c *FakeLogzMetricsTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogzMetricsTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logzmetricstargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogzMetricsTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogzMetricsTarget), err
}
