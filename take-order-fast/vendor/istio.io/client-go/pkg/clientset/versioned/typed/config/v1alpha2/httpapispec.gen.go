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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPAPISpecsGetter has a method to return a HTTPAPISpecInterface.
// A group's client should implement this interface.
type HTTPAPISpecsGetter interface {
	HTTPAPISpecs(namespace string) HTTPAPISpecInterface
}

// HTTPAPISpecInterface has methods to work with HTTPAPISpec resources.
type HTTPAPISpecInterface interface {
	Create(ctx context.Context, hTTPAPISpec *v1alpha2.HTTPAPISpec, opts v1.CreateOptions) (*v1alpha2.HTTPAPISpec, error)
	Update(ctx context.Context, hTTPAPISpec *v1alpha2.HTTPAPISpec, opts v1.UpdateOptions) (*v1alpha2.HTTPAPISpec, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.HTTPAPISpec, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.HTTPAPISpecList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HTTPAPISpec, err error)
	HTTPAPISpecExpansion
}

// hTTPAPISpecs implements HTTPAPISpecInterface
type hTTPAPISpecs struct {
	client rest.Interface
	ns     string
}

// newHTTPAPISpecs returns a HTTPAPISpecs
func newHTTPAPISpecs(c *ConfigV1alpha2Client, namespace string) *hTTPAPISpecs {
	return &hTTPAPISpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPAPISpec, and returns the corresponding hTTPAPISpec object, and an error if there is any.
func (c *hTTPAPISpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPAPISpecs that match those selectors.
func (c *hTTPAPISpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.HTTPAPISpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecs.
func (c *hTTPAPISpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hTTPAPISpec and creates it.  Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *hTTPAPISpecs) Create(ctx context.Context, hTTPAPISpec *v1alpha2.HTTPAPISpec, opts v1.CreateOptions) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPAPISpec).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hTTPAPISpec and updates it. Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *hTTPAPISpecs) Update(ctx context.Context, hTTPAPISpec *v1alpha2.HTTPAPISpec, opts v1.UpdateOptions) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(hTTPAPISpec.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPAPISpec).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hTTPAPISpec and deletes it. Returns an error if one occurs.
func (c *hTTPAPISpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPAPISpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hTTPAPISpec.
func (c *hTTPAPISpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
