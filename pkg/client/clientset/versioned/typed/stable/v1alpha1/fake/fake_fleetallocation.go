// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "agones.dev/agones/pkg/apis/stable/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFleetAllocations implements FleetAllocationInterface
type FakeFleetAllocations struct {
	Fake *FakeStableV1alpha1
	ns   string
}

var fleetallocationsResource = schema.GroupVersionResource{Group: "stable.agones.dev", Version: "v1alpha1", Resource: "fleetallocations"}

var fleetallocationsKind = schema.GroupVersionKind{Group: "stable.agones.dev", Version: "v1alpha1", Kind: "FleetAllocation"}

// Get takes name of the fleetAllocation, and returns the corresponding fleetAllocation object, and an error if there is any.
func (c *FakeFleetAllocations) Get(name string, options v1.GetOptions) (result *v1alpha1.FleetAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fleetallocationsResource, c.ns, name), &v1alpha1.FleetAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FleetAllocation), err
}

// List takes label and field selectors, and returns the list of FleetAllocations that match those selectors.
func (c *FakeFleetAllocations) List(opts v1.ListOptions) (result *v1alpha1.FleetAllocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fleetallocationsResource, fleetallocationsKind, c.ns, opts), &v1alpha1.FleetAllocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FleetAllocationList{ListMeta: obj.(*v1alpha1.FleetAllocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.FleetAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fleetAllocations.
func (c *FakeFleetAllocations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fleetallocationsResource, c.ns, opts))

}

// Create takes the representation of a fleetAllocation and creates it.  Returns the server's representation of the fleetAllocation, and an error, if there is any.
func (c *FakeFleetAllocations) Create(fleetAllocation *v1alpha1.FleetAllocation) (result *v1alpha1.FleetAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fleetallocationsResource, c.ns, fleetAllocation), &v1alpha1.FleetAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FleetAllocation), err
}

// Update takes the representation of a fleetAllocation and updates it. Returns the server's representation of the fleetAllocation, and an error, if there is any.
func (c *FakeFleetAllocations) Update(fleetAllocation *v1alpha1.FleetAllocation) (result *v1alpha1.FleetAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fleetallocationsResource, c.ns, fleetAllocation), &v1alpha1.FleetAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FleetAllocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFleetAllocations) UpdateStatus(fleetAllocation *v1alpha1.FleetAllocation) (*v1alpha1.FleetAllocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fleetallocationsResource, "status", c.ns, fleetAllocation), &v1alpha1.FleetAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FleetAllocation), err
}

// Delete takes name of the fleetAllocation and deletes it. Returns an error if one occurs.
func (c *FakeFleetAllocations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fleetallocationsResource, c.ns, name), &v1alpha1.FleetAllocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFleetAllocations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fleetallocationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FleetAllocationList{})
	return err
}

// Patch applies the patch and returns the patched fleetAllocation.
func (c *FakeFleetAllocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FleetAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fleetallocationsResource, c.ns, name, data, subresources...), &v1alpha1.FleetAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FleetAllocation), err
}
