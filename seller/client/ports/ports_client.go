// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPortInfoUsingGET2 fetches list of virtual ports

The Get Ports API allow ECX participants to retrieve a list of all their Ports assigned to their customer id (known as uuID). This is useful and necessary when creating connections as all connections within the ECX platform are associated with the customer�s port and require the port id when calling the Post Connections API. Detailed information about all ports are returned including the encapsulation (Dot1Q or QinQ), the metro and region where the ports are located, port size/bandwidth, and other pertinent information.
*/
func (a *Client) GetPortInfoUsingGET2(params *GetPortInfoUsingGET2Params) (*GetPortInfoUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortInfoUsingGET2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPortInfoUsingGET_2",
		Method:             "GET",
		PathPattern:        "/ecx/v3/port/userport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortInfoUsingGET2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPortInfoUsingGET2OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
