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
	"context"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMutatingWebhookConfigurations implements MutatingWebhookConfigurationInterface
type FakeMutatingWebhookConfigurations struct {
	Fake *FakeAdmissionregistrationV1
}

var mutatingwebhookconfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1", Resource: "mutatingwebhookconfigurations"}

var mutatingwebhookconfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "v1", Kind: "MutatingWebhookConfiguration"}

// Get takes name of the mutatingWebhookConfiguration, and returns the corresponding mutatingWebhookConfiguration object, and an error if there is any.
func (c *FakeMutatingWebhookConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *admissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(mutatingwebhookconfigurationsResource, name), &admissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.MutatingWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of MutatingWebhookConfigurations that match those selectors.
func (c *FakeMutatingWebhookConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *admissionregistrationv1.MutatingWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, opts), &admissionregistrationv1.MutatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1.MutatingWebhookConfigurationList{ListMeta: obj.(*admissionregistrationv1.MutatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1.MutatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mutatingWebhookConfigurations.
func (c *FakeMutatingWebhookConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(mutatingwebhookconfigurationsResource, opts))
}

// Create takes the representation of a mutatingWebhookConfiguration and creates it.  Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Create(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1.MutatingWebhookConfiguration, opts v1.CreateOptions) (result *admissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.MutatingWebhookConfiguration), err
}

// Update takes the representation of a mutatingWebhookConfiguration and updates it. Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Update(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1.MutatingWebhookConfiguration, opts v1.UpdateOptions) (result *admissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.MutatingWebhookConfiguration), err
}

// Delete takes name of the mutatingWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeMutatingWebhookConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(mutatingwebhookconfigurationsResource, name), &admissionregistrationv1.MutatingWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMutatingWebhookConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(mutatingwebhookconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &admissionregistrationv1.MutatingWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched mutatingWebhookConfiguration.
func (c *FakeMutatingWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *admissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, name, pt, data, subresources...), &admissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.MutatingWebhookConfiguration), err
}
