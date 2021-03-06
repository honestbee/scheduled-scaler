/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.restdev.com/operators/pkg/apis/scaling/v1alpha1"
)

// FakeScheduledScalers implements ScheduledScalerInterface
type FakeScheduledScalers struct {
	Fake *FakeScalingV1alpha1
	ns   string
}

var scheduledscalersResource = schema.GroupVersionResource{Group: "scaling.k8s.restdev.com", Version: "v1alpha1", Resource: "scheduledscalers"}

var scheduledscalersKind = schema.GroupVersionKind{Group: "scaling.k8s.restdev.com", Version: "v1alpha1", Kind: "ScheduledScaler"}

// Get takes name of the scheduledScaler, and returns the corresponding scheduledScaler object, and an error if there is any.
func (c *FakeScheduledScalers) Get(name string, options v1.GetOptions) (result *v1alpha1.ScheduledScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scheduledscalersResource, c.ns, name), &v1alpha1.ScheduledScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledScaler), err
}

// List takes label and field selectors, and returns the list of ScheduledScalers that match those selectors.
func (c *FakeScheduledScalers) List(opts v1.ListOptions) (result *v1alpha1.ScheduledScalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scheduledscalersResource, scheduledscalersKind, c.ns, opts), &v1alpha1.ScheduledScalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScheduledScalerList{}
	for _, item := range obj.(*v1alpha1.ScheduledScalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduledScalers.
func (c *FakeScheduledScalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scheduledscalersResource, c.ns, opts))

}

// Create takes the representation of a scheduledScaler and creates it.  Returns the server's representation of the scheduledScaler, and an error, if there is any.
func (c *FakeScheduledScalers) Create(scheduledScaler *v1alpha1.ScheduledScaler) (result *v1alpha1.ScheduledScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scheduledscalersResource, c.ns, scheduledScaler), &v1alpha1.ScheduledScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledScaler), err
}

// Update takes the representation of a scheduledScaler and updates it. Returns the server's representation of the scheduledScaler, and an error, if there is any.
func (c *FakeScheduledScalers) Update(scheduledScaler *v1alpha1.ScheduledScaler) (result *v1alpha1.ScheduledScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scheduledscalersResource, c.ns, scheduledScaler), &v1alpha1.ScheduledScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledScaler), err
}

// Delete takes name of the scheduledScaler and deletes it. Returns an error if one occurs.
func (c *FakeScheduledScalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scheduledscalersResource, c.ns, name), &v1alpha1.ScheduledScaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduledScalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scheduledscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScheduledScalerList{})
	return err
}

// Patch applies the patch and returns the patched scheduledScaler.
func (c *FakeScheduledScalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScheduledScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scheduledscalersResource, c.ns, name, data, subresources...), &v1alpha1.ScheduledScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledScaler), err
}
