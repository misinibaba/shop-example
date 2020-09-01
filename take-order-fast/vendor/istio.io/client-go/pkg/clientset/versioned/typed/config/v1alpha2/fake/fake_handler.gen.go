// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHandlers implements HandlerInterface
type FakeHandlers struct {
	Fake *FakeConfigV1alpha2
	ns   string
}

var handlersResource = schema.GroupVersionResource{Group: "config.istio.io", Version: "v1alpha2", Resource: "handlers"}

var handlersKind = schema.GroupVersionKind{Group: "config.istio.io", Version: "v1alpha2", Kind: "Handler"}

// Get takes name of the handler, and returns the corresponding handler object, and an error if there is any.
func (c *FakeHandlers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Handler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(handlersResource, c.ns, name), &v1alpha2.Handler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Handler), err
}

// List takes label and field selectors, and returns the list of Handlers that match those selectors.
func (c *FakeHandlers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HandlerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(handlersResource, handlersKind, c.ns, opts), &v1alpha2.HandlerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HandlerList{ListMeta: obj.(*v1alpha2.HandlerList).ListMeta}
	for _, item := range obj.(*v1alpha2.HandlerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested handlers.
func (c *FakeHandlers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(handlersResource, c.ns, opts))

}

// Create takes the representation of a handler and creates it.  Returns the server's representation of the handler, and an error, if there is any.
func (c *FakeHandlers) Create(ctx context.Context, handler *v1alpha2.Handler, opts v1.CreateOptions) (result *v1alpha2.Handler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(handlersResource, c.ns, handler), &v1alpha2.Handler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Handler), err
}

// Update takes the representation of a handler and updates it. Returns the server's representation of the handler, and an error, if there is any.
func (c *FakeHandlers) Update(ctx context.Context, handler *v1alpha2.Handler, opts v1.UpdateOptions) (result *v1alpha2.Handler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(handlersResource, c.ns, handler), &v1alpha2.Handler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Handler), err
}

// Delete takes name of the handler and deletes it. Returns an error if one occurs.
func (c *FakeHandlers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(handlersResource, c.ns, name), &v1alpha2.Handler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHandlers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(handlersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.HandlerList{})
	return err
}

// Patch applies the patch and returns the patched handler.
func (c *FakeHandlers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Handler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(handlersResource, c.ns, name, pt, data, subresources...), &v1alpha2.Handler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Handler), err
}
