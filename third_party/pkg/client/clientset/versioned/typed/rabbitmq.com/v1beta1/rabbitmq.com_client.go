/*
Copyright 2020 The Knative Authors

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
	rest "k8s.io/client-go/rest"
	v1beta1 "knative.dev/eventing-rabbitmq/third_party/pkg/apis/rabbitmq.com/v1beta1"
	"knative.dev/eventing-rabbitmq/third_party/pkg/client/clientset/versioned/scheme"
)

type RabbitmqV1beta1Interface interface {
	RESTClient() rest.Interface
	BindingsGetter
	ExchangesGetter
	PermissionsGetter
	PoliciesGetter
	QueuesGetter
	SchemaReplicationsGetter
	UsersGetter
	VhostsGetter
}

// RabbitmqV1beta1Client is used to interact with features provided by the rabbitmq.com group.
type RabbitmqV1beta1Client struct {
	restClient rest.Interface
}

func (c *RabbitmqV1beta1Client) Bindings(namespace string) BindingInterface {
	return newBindings(c, namespace)
}

func (c *RabbitmqV1beta1Client) Exchanges(namespace string) ExchangeInterface {
	return newExchanges(c, namespace)
}

func (c *RabbitmqV1beta1Client) Permissions(namespace string) PermissionInterface {
	return newPermissions(c, namespace)
}

func (c *RabbitmqV1beta1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *RabbitmqV1beta1Client) Queues(namespace string) QueueInterface {
	return newQueues(c, namespace)
}

func (c *RabbitmqV1beta1Client) SchemaReplications(namespace string) SchemaReplicationInterface {
	return newSchemaReplications(c, namespace)
}

func (c *RabbitmqV1beta1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *RabbitmqV1beta1Client) Vhosts(namespace string) VhostInterface {
	return newVhosts(c, namespace)
}

// NewForConfig creates a new RabbitmqV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*RabbitmqV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &RabbitmqV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new RabbitmqV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *RabbitmqV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new RabbitmqV1beta1Client for the given RESTClient.
func New(c rest.Interface) *RabbitmqV1beta1Client {
	return &RabbitmqV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
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
func (c *RabbitmqV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}