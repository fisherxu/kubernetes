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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	apps "k8s.io/kubernetes/pkg/apis/apps"
	scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

// ControllerRevisionsGetter has a method to return a ControllerRevisionInterface.
// A group's client should implement this interface.
type ControllerRevisionsGetter interface {
	ControllerRevisions(namespace string) ControllerRevisionInterface
}

// ControllerRevisionInterface has methods to work with ControllerRevision resources.
type ControllerRevisionInterface interface {
	Create(*apps.ControllerRevision) (*apps.ControllerRevision, error)
	Update(*apps.ControllerRevision) (*apps.ControllerRevision, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*apps.ControllerRevision, error)
	List(opts v1.ListOptions) (*apps.ControllerRevisionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ControllerRevision, err error)
	ControllerRevisionExpansion
}

// controllerRevisions implements ControllerRevisionInterface
type controllerRevisions struct {
	client rest.Interface
	ns     string
}

// newControllerRevisions returns a ControllerRevisions
func newControllerRevisions(c *AppsClient, namespace string) *controllerRevisions {
	return &controllerRevisions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the controllerRevision, and returns the corresponding controllerRevision object, and an error if there is any.
func (c *controllerRevisions) Get(name string, options v1.GetOptions) (result *apps.ControllerRevision, err error) {
	result = &apps.ControllerRevision{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ControllerRevisions that match those selectors.
func (c *controllerRevisions) List(opts v1.ListOptions) (result *apps.ControllerRevisionList, err error) {
	result = &apps.ControllerRevisionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested controllerRevisions.
func (c *controllerRevisions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a controllerRevision and creates it.  Returns the server's representation of the controllerRevision, and an error, if there is any.
func (c *controllerRevisions) Create(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
	result = &apps.ControllerRevision{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("controllerrevisions").
		Body(controllerRevision).
		Do().
		Into(result)
	return
}

// Update takes the representation of a controllerRevision and updates it. Returns the server's representation of the controllerRevision, and an error, if there is any.
func (c *controllerRevisions) Update(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
	result = &apps.ControllerRevision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(controllerRevision.Name).
		Body(controllerRevision).
		Do().
		Into(result)
	return
}

// Delete takes name of the controllerRevision and deletes it. Returns an error if one occurs.
func (c *controllerRevisions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *controllerRevisions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched controllerRevision.
func (c *controllerRevisions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ControllerRevision, err error) {
	result = &apps.ControllerRevision{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("controllerrevisions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
