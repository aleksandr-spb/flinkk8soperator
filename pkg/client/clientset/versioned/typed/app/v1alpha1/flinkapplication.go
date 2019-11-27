// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/aleksandr-spb/flinkk8soperator/pkg/apis/app/v1alpha1"
	scheme "github.com/aleksandr-spb/flinkk8soperator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FlinkApplicationsGetter has a method to return a FlinkApplicationInterface.
// A group's client should implement this interface.
type FlinkApplicationsGetter interface {
	FlinkApplications(namespace string) FlinkApplicationInterface
}

// FlinkApplicationInterface has methods to work with FlinkApplication resources.
type FlinkApplicationInterface interface {
	Create(*v1alpha1.FlinkApplication) (*v1alpha1.FlinkApplication, error)
	Update(*v1alpha1.FlinkApplication) (*v1alpha1.FlinkApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FlinkApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.FlinkApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FlinkApplication, err error)
	FlinkApplicationExpansion
}

// flinkApplications implements FlinkApplicationInterface
type flinkApplications struct {
	client rest.Interface
	ns     string
}

// newFlinkApplications returns a FlinkApplications
func newFlinkApplications(c *FlinkV1alpha1Client, namespace string) *flinkApplications {
	return &flinkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flinkApplication, and returns the corresponding flinkApplication object, and an error if there is any.
func (c *flinkApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.FlinkApplication, err error) {
	result = &v1alpha1.FlinkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlinkApplications that match those selectors.
func (c *flinkApplications) List(opts v1.ListOptions) (result *v1alpha1.FlinkApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FlinkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flinkApplications.
func (c *flinkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a flinkApplication and creates it.  Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *flinkApplications) Create(flinkApplication *v1alpha1.FlinkApplication) (result *v1alpha1.FlinkApplication, err error) {
	result = &v1alpha1.FlinkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flinkapplications").
		Body(flinkApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a flinkApplication and updates it. Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *flinkApplications) Update(flinkApplication *v1alpha1.FlinkApplication) (result *v1alpha1.FlinkApplication, err error) {
	result = &v1alpha1.FlinkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(flinkApplication.Name).
		Body(flinkApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the flinkApplication and deletes it. Returns an error if one occurs.
func (c *flinkApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flinkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched flinkApplication.
func (c *flinkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FlinkApplication, err error) {
	result = &v1alpha1.FlinkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flinkapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
