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

package internalclientset

import (
	"fmt"

	extensionsv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/extensions/v1alpha1"
	flowv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1"
	routingv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/routing/v1alpha1"
	sourcesv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1"
	targetsv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ExtensionsV1alpha1() extensionsv1alpha1.ExtensionsV1alpha1Interface
	FlowV1alpha1() flowv1alpha1.FlowV1alpha1Interface
	RoutingV1alpha1() routingv1alpha1.RoutingV1alpha1Interface
	SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface
	TargetsV1alpha1() targetsv1alpha1.TargetsV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	extensionsV1alpha1 *extensionsv1alpha1.ExtensionsV1alpha1Client
	flowV1alpha1       *flowv1alpha1.FlowV1alpha1Client
	routingV1alpha1    *routingv1alpha1.RoutingV1alpha1Client
	sourcesV1alpha1    *sourcesv1alpha1.SourcesV1alpha1Client
	targetsV1alpha1    *targetsv1alpha1.TargetsV1alpha1Client
}

// ExtensionsV1alpha1 retrieves the ExtensionsV1alpha1Client
func (c *Clientset) ExtensionsV1alpha1() extensionsv1alpha1.ExtensionsV1alpha1Interface {
	return c.extensionsV1alpha1
}

// FlowV1alpha1 retrieves the FlowV1alpha1Client
func (c *Clientset) FlowV1alpha1() flowv1alpha1.FlowV1alpha1Interface {
	return c.flowV1alpha1
}

// RoutingV1alpha1 retrieves the RoutingV1alpha1Client
func (c *Clientset) RoutingV1alpha1() routingv1alpha1.RoutingV1alpha1Interface {
	return c.routingV1alpha1
}

// SourcesV1alpha1 retrieves the SourcesV1alpha1Client
func (c *Clientset) SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface {
	return c.sourcesV1alpha1
}

// TargetsV1alpha1 retrieves the TargetsV1alpha1Client
func (c *Clientset) TargetsV1alpha1() targetsv1alpha1.TargetsV1alpha1Interface {
	return c.targetsV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.extensionsV1alpha1, err = extensionsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.flowV1alpha1, err = flowv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.routingV1alpha1, err = routingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1alpha1, err = sourcesv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.targetsV1alpha1, err = targetsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.extensionsV1alpha1 = extensionsv1alpha1.NewForConfigOrDie(c)
	cs.flowV1alpha1 = flowv1alpha1.NewForConfigOrDie(c)
	cs.routingV1alpha1 = routingv1alpha1.NewForConfigOrDie(c)
	cs.sourcesV1alpha1 = sourcesv1alpha1.NewForConfigOrDie(c)
	cs.targetsV1alpha1 = targetsv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.extensionsV1alpha1 = extensionsv1alpha1.New(c)
	cs.flowV1alpha1 = flowv1alpha1.New(c)
	cs.routingV1alpha1 = routingv1alpha1.New(c)
	cs.sourcesV1alpha1 = sourcesv1alpha1.New(c)
	cs.targetsV1alpha1 = targetsv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
