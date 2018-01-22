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

package fake

import (
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePersistentVolumeClaims implements PersistentVolumeClaimInterface
type FakePersistentVolumeClaims struct {
	Fake *FakeCoreV1
	ns   string
}

var persistentvolumeclaimsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "persistentvolumeclaims"}

var persistentvolumeclaimsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PersistentVolumeClaim"}

// Get takes name of the persistentVolumeClaim, and returns the corresponding persistentVolumeClaim object, and an error if there is any.
func (c *FakePersistentVolumeClaims) Get(name string, options v1.GetOptions) (result *core_v1.PersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(persistentvolumeclaimsResource, c.ns, name), &core_v1.PersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.PersistentVolumeClaim), err
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors.
func (c *FakePersistentVolumeClaims) List(opts v1.ListOptions) (result *core_v1.PersistentVolumeClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(persistentvolumeclaimsResource, persistentvolumeclaimsKind, c.ns, opts), &core_v1.PersistentVolumeClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core_v1.PersistentVolumeClaimList{}
	for _, item := range obj.(*core_v1.PersistentVolumeClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persistentVolumeClaims.
func (c *FakePersistentVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(persistentvolumeclaimsResource, c.ns, opts))

}

// Create takes the representation of a persistentVolumeClaim and creates it.  Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *FakePersistentVolumeClaims) Create(persistentVolumeClaim *core_v1.PersistentVolumeClaim) (result *core_v1.PersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), &core_v1.PersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.PersistentVolumeClaim), err
}

// Update takes the representation of a persistentVolumeClaim and updates it. Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *FakePersistentVolumeClaims) Update(persistentVolumeClaim *core_v1.PersistentVolumeClaim) (result *core_v1.PersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), &core_v1.PersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.PersistentVolumeClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersistentVolumeClaims) UpdateStatus(persistentVolumeClaim *core_v1.PersistentVolumeClaim) (*core_v1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(persistentvolumeclaimsResource, "status", c.ns, persistentVolumeClaim), &core_v1.PersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.PersistentVolumeClaim), err
}

// Delete takes name of the persistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *FakePersistentVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(persistentvolumeclaimsResource, c.ns, name), &core_v1.PersistentVolumeClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePersistentVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(persistentvolumeclaimsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core_v1.PersistentVolumeClaimList{})
	return err
}

// Patch applies the patch and returns the patched persistentVolumeClaim.
func (c *FakePersistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core_v1.PersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(persistentvolumeclaimsResource, c.ns, name, data, subresources...), &core_v1.PersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.PersistentVolumeClaim), err
}
