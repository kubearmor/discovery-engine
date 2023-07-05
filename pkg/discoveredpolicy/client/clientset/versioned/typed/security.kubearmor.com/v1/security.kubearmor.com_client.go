// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/accuknox/auto-policy-discovery/pkg/discoveredpolicy/api/security.kubearmor.com/v1"
	"github.com/accuknox/auto-policy-discovery/pkg/discoveredpolicy/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SecurityV1Interface interface {
	RESTClient() rest.Interface
	DiscoveredPoliciesGetter
}

// SecurityV1Client is used to interact with features provided by the security.kubearmor.com group.
type SecurityV1Client struct {
	restClient rest.Interface
}

func (c *SecurityV1Client) DiscoveredPolicies(namespace string) DiscoveredPolicyInterface {
	return newDiscoveredPolicies(c, namespace)
}

// NewForConfig creates a new SecurityV1Client for the given config.
func NewForConfig(c *rest.Config) (*SecurityV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SecurityV1Client{client}, nil
}

// NewForConfigOrDie creates a new SecurityV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SecurityV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SecurityV1Client for the given RESTClient.
func New(c rest.Interface) *SecurityV1Client {
	return &SecurityV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
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
func (c *SecurityV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
