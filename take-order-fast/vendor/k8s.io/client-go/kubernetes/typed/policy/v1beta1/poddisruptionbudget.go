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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// PodDisruptionBudgetsGetter has a method to return a PodDisruptionBudgetInterface.
// A group's client should implement this interface.
type PodDisruptionBudgetsGetter interface {
	PodDisruptionBudgets(namespace string) PodDisruptionBudgetInterface
}

// PodDisruptionBudgetInterface has methods to work with PodDisruptionBudget resources.
type PodDisruptionBudgetInterface interface {
	Create(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.CreateOptions) (*v1beta1.PodDisruptionBudget, error)
	Update(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.UpdateOptions) (*v1beta1.PodDisruptionBudget, error)
	UpdateStatus(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.UpdateOptions) (*v1beta1.PodDisruptionBudget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.PodDisruptionBudget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PodDisruptionBudgetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodDisruptionBudget, err error)
	PodDisruptionBudgetExpansion
}

// podDisruptionBudgets implements PodDisruptionBudgetInterface
type podDisruptionBudgets struct {
	client rest.Interface
	ns     string
}

// newPodDisruptionBudgets returns a PodDisruptionBudgets
func newPodDisruptionBudgets(c *PolicyV1beta1Client, namespace string) *podDisruptionBudgets {
	return &podDisruptionBudgets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podDisruptionBudget, and returns the corresponding podDisruptionBudget object, and an error if there is any.
func (c *podDisruptionBudgets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PodDisruptionBudget, err error) {
	result = &v1beta1.PodDisruptionBudget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodDisruptionBudgets that match those selectors.
func (c *podDisruptionBudgets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PodDisruptionBudgetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PodDisruptionBudgetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podDisruptionBudgets.
func (c *podDisruptionBudgets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podDisruptionBudget and creates it.  Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (c *podDisruptionBudgets) Create(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.CreateOptions) (result *v1beta1.PodDisruptionBudget, err error) {
	result = &v1beta1.PodDisruptionBudget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podDisruptionBudget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podDisruptionBudget and updates it. Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (c *podDisruptionBudgets) Update(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.UpdateOptions) (result *v1beta1.PodDisruptionBudget, err error) {
	result = &v1beta1.PodDisruptionBudget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		Name(podDisruptionBudget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podDisruptionBudget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *podDisruptionBudgets) UpdateStatus(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts v1.UpdateOptions) (result *v1beta1.PodDisruptionBudget, err error) {
	result = &v1beta1.PodDisruptionBudget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		Name(podDisruptionBudget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podDisruptionBudget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podDisruptionBudget and deletes it. Returns an error if one occurs.
func (c *podDisruptionBudgets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podDisruptionBudgets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podDisruptionBudget.
func (c *podDisruptionBudgets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodDisruptionBudget, err error) {
	result = &v1beta1.PodDisruptionBudget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("poddisruptionbudgets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
