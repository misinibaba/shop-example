/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFlowSchemas implements FlowSchemaInterface
type FakeFlowSchemas struct {
	Fake *FakeFlowcontrolV1alpha1
}

var flowschemasResource = schema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1alpha1", Resource: "flowschemas"}

var flowschemasKind = schema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1alpha1", Kind: "FlowSchema"}

// Get takes name of the flowSchema, and returns the corresponding flowSchema object, and an error if there is any.
func (c *FakeFlowSchemas) Get(name string, options v1.GetOptions) (result *v1alpha1.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(flowschemasResource, name), &v1alpha1.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlowSchema), err
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *FakeFlowSchemas) List(opts v1.ListOptions) (result *v1alpha1.FlowSchemaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(flowschemasResource, flowschemasKind, opts), &v1alpha1.FlowSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FlowSchemaList{ListMeta: obj.(*v1alpha1.FlowSchemaList).ListMeta}
	for _, item := range obj.(*v1alpha1.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flowSchemas.
func (c *FakeFlowSchemas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(flowschemasResource, opts))
}

// Create takes the representation of a flowSchema and creates it.  Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *FakeFlowSchemas) Create(flowSchema *v1alpha1.FlowSchema) (result *v1alpha1.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(flowschemasResource, flowSchema), &v1alpha1.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlowSchema), err
}

// Update takes the representation of a flowSchema and updates it. Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *FakeFlowSchemas) Update(flowSchema *v1alpha1.FlowSchema) (result *v1alpha1.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(flowschemasResource, flowSchema), &v1alpha1.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlowSchema), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFlowSchemas) UpdateStatus(flowSchema *v1alpha1.FlowSchema) (*v1alpha1.FlowSchema, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(flowschemasResource, "status", flowSchema), &v1alpha1.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlowSchema), err
}

// Delete takes name of the flowSchema and deletes it. Returns an error if one occurs.
func (c *FakeFlowSchemas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(flowschemasResource, name), &v1alpha1.FlowSchema{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFlowSchemas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(flowschemasResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FlowSchemaList{})
	return err
}

// Patch applies the patch and returns the patched flowSchema.
func (c *FakeFlowSchemas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(flowschemasResource, name, pt, data, subresources...), &v1alpha1.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlowSchema), err
}
