/*
Copyright 2018 Infoblox, Inc.

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
	v1alpha1 "github.com/infobloxopen/atlas/pkg/apis/atlasauthz/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppRoles implements AppRoleInterface
type FakeAppRoles struct {
	Fake *FakeAtlasauthzV1alpha1
	ns   string
}

var approlesResource = schema.GroupVersionResource{Group: "atlasauthz.infoblox.com", Version: "v1alpha1", Resource: "approles"}

var approlesKind = schema.GroupVersionKind{Group: "atlasauthz.infoblox.com", Version: "v1alpha1", Kind: "AppRole"}

// Get takes name of the appRole, and returns the corresponding appRole object, and an error if there is any.
func (c *FakeAppRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.AppRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(approlesResource, c.ns, name), &v1alpha1.AppRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppRole), err
}

// List takes label and field selectors, and returns the list of AppRoles that match those selectors.
func (c *FakeAppRoles) List(opts v1.ListOptions) (result *v1alpha1.AppRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(approlesResource, approlesKind, c.ns, opts), &v1alpha1.AppRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppRoleList{}
	for _, item := range obj.(*v1alpha1.AppRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appRoles.
func (c *FakeAppRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(approlesResource, c.ns, opts))

}

// Create takes the representation of a appRole and creates it.  Returns the server's representation of the appRole, and an error, if there is any.
func (c *FakeAppRoles) Create(appRole *v1alpha1.AppRole) (result *v1alpha1.AppRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(approlesResource, c.ns, appRole), &v1alpha1.AppRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppRole), err
}

// Update takes the representation of a appRole and updates it. Returns the server's representation of the appRole, and an error, if there is any.
func (c *FakeAppRoles) Update(appRole *v1alpha1.AppRole) (result *v1alpha1.AppRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(approlesResource, c.ns, appRole), &v1alpha1.AppRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppRole), err
}

// Delete takes name of the appRole and deletes it. Returns an error if one occurs.
func (c *FakeAppRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(approlesResource, c.ns, name), &v1alpha1.AppRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(approlesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppRoleList{})
	return err
}

// Patch applies the patch and returns the patched appRole.
func (c *FakeAppRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(approlesResource, c.ns, name, data, subresources...), &v1alpha1.AppRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppRole), err
}