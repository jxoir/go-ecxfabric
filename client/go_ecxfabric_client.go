// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jxoir/go-ecxfabric/client/access_token"
	"github.com/jxoir/go-ecxfabric/client/bundle_offering"
	"github.com/jxoir/go-ecxfabric/client/buyer_preferences"
	"github.com/jxoir/go-ecxfabric/client/connections"
	"github.com/jxoir/go-ecxfabric/client/connector"
	"github.com/jxoir/go-ecxfabric/client/metros"
	"github.com/jxoir/go-ecxfabric/client/ports"
	"github.com/jxoir/go-ecxfabric/client/public_ip_block"
	"github.com/jxoir/go-ecxfabric/client/routing_instance"
	"github.com/jxoir/go-ecxfabric/client/seller_services"
	"github.com/jxoir/go-ecxfabric/client/subscription"
)

// Default go ecxfabric HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.equinix.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new go ecxfabric HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GoEcxfabric {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new go ecxfabric HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GoEcxfabric {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new go ecxfabric client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GoEcxfabric {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GoEcxfabric)
	cli.Transport = transport

	cli.AccessToken = access_token.New(transport, formats)

	cli.BundleOffering = bundle_offering.New(transport, formats)

	cli.BuyerPreferences = buyer_preferences.New(transport, formats)

	cli.Connections = connections.New(transport, formats)

	cli.Connector = connector.New(transport, formats)

	cli.Metros = metros.New(transport, formats)

	cli.Ports = ports.New(transport, formats)

	cli.PublicIPBlock = public_ip_block.New(transport, formats)

	cli.RoutingInstance = routing_instance.New(transport, formats)

	cli.SellerServices = seller_services.New(transport, formats)

	cli.Subscription = subscription.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GoEcxfabric is a client for go ecxfabric
type GoEcxfabric struct {
	AccessToken *access_token.Client

	BundleOffering *bundle_offering.Client

	BuyerPreferences *buyer_preferences.Client

	Connections *connections.Client

	Connector *connector.Client

	Metros *metros.Client

	Ports *ports.Client

	PublicIPBlock *public_ip_block.Client

	RoutingInstance *routing_instance.Client

	SellerServices *seller_services.Client

	Subscription *subscription.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GoEcxfabric) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AccessToken.SetTransport(transport)

	c.BundleOffering.SetTransport(transport)

	c.BuyerPreferences.SetTransport(transport)

	c.Connections.SetTransport(transport)

	c.Connector.SetTransport(transport)

	c.Metros.SetTransport(transport)

	c.Ports.SetTransport(transport)

	c.PublicIPBlock.SetTransport(transport)

	c.RoutingInstance.SetTransport(transport)

	c.SellerServices.SetTransport(transport)

	c.Subscription.SetTransport(transport)

}
