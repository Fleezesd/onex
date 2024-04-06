// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	rest "k8s.io/client-go/rest"

	v1 "github.com/superproj/onex/pkg/apis/coordination/v1"
	"github.com/superproj/onex/pkg/generated/clientset/versioned/scheme"
)

type CoordinationV1Interface interface {
	RESTClient() rest.Interface
	LeasesGetter
}

// CoordinationV1Client is used to interact with features provided by the coordination.k8s.io group.
type CoordinationV1Client struct {
	restClient rest.Interface
}

func (c *CoordinationV1Client) Leases(namespace string) LeaseInterface {
	return newLeases(c, namespace)
}

// NewForConfig creates a new CoordinationV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CoordinationV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CoordinationV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CoordinationV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CoordinationV1Client{client}, nil
}

// NewForConfigOrDie creates a new CoordinationV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoordinationV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoordinationV1Client for the given RESTClient.
func New(c rest.Interface) *CoordinationV1Client {
	return &CoordinationV1Client{c}
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
func (c *CoordinationV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}