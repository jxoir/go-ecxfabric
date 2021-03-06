// Code generated by go-swagger; DO NOT EDIT.

package metros

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new metros API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metros API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMetrosUsingGET returns list of all metros

The Get Metros API allows ECX participants to either retrieve a list of all metros where they have ports or to retrieve a list of all metros where ECX is enabled, depending on the API parameters.
*/
func (a *Client) GetMetrosUsingGET(params *GetMetrosUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetMetrosUsingGETOK, *GetMetrosUsingGETNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetrosUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMetrosUsingGET",
		Method:             "GET",
		PathPattern:        "/ecx/v3/l2/common/metros",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetrosUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetMetrosUsingGETOK:
		return value, nil, nil
	case *GetMetrosUsingGETNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
