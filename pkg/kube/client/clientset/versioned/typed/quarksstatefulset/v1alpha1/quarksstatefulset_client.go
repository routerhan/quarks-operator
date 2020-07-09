/*

Don't alter this file, it was generated.

*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "code.cloudfoundry.org/quarks-statefulset/pkg/kube/apis/quarksstatefulset/v1alpha1"
	"code.cloudfoundry.org/quarks-statefulset/pkg/kube/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type QuarksstatefulsetV1alpha1Interface interface {
	RESTClient() rest.Interface
	QuarksStatefulSetsGetter
}

// QuarksstatefulsetV1alpha1Client is used to interact with features provided by the quarksstatefulset group.
type QuarksstatefulsetV1alpha1Client struct {
	restClient rest.Interface
}

func (c *QuarksstatefulsetV1alpha1Client) QuarksStatefulSets(namespace string) QuarksStatefulSetInterface {
	return newQuarksStatefulSets(c, namespace)
}

// NewForConfig creates a new QuarksstatefulsetV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*QuarksstatefulsetV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &QuarksstatefulsetV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new QuarksstatefulsetV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *QuarksstatefulsetV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new QuarksstatefulsetV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *QuarksstatefulsetV1alpha1Client {
	return &QuarksstatefulsetV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *QuarksstatefulsetV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
