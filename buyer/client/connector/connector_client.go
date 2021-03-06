// Code generated by go-swagger; DO NOT EDIT.

package connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new connector API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connector API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateConnectorUsingPOST creates connector

This API is used to create Connector.
*/
func (a *Client) CreateConnectorUsingPOST(params *CreateConnectorUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConnectorUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConnectorUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createConnectorUsingPOST",
		Method:             "POST",
		PathPattern:        "/ecx/v3/l3/connector",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConnectorUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateConnectorUsingPOSTCreated), nil

}

/*
DeleteConnectorUsingDELETE deletes connector

This API is used to delete Connector for given uuid
*/
func (a *Client) DeleteConnectorUsingDELETE(params *DeleteConnectorUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConnectorUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConnectorUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConnectorUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/ecx/v3/l3/connector/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConnectorUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteConnectorUsingDELETENoContent), nil

}

/*
GetAllConnectorsUsingGET returns list of connectors

This API is used to get all Connectors with respective to query params
*/
func (a *Client) GetAllConnectorsUsingGET(params *GetAllConnectorsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllConnectorsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllConnectorsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllConnectorsUsingGET",
		Method:             "GET",
		PathPattern:        "/ecx/v3/l3/connector",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllConnectorsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllConnectorsUsingGETOK), nil

}

/*
IsConnectorExistUsingGET validates connector name already exists or not

This API is used to check connector name exists or not
*/
func (a *Client) IsConnectorExistUsingGET(params *IsConnectorExistUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*IsConnectorExistUsingGETOK, *IsConnectorExistUsingGETNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsConnectorExistUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isConnectorExistUsingGET",
		Method:             "GET",
		PathPattern:        "/ecx/v3/l3/connector/exist/{metroCode}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IsConnectorExistUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *IsConnectorExistUsingGETOK:
		return value, nil, nil
	case *IsConnectorExistUsingGETNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateConnectorUsingPATCH updates connector

This API is used to update connector for given uuid
*/
func (a *Client) UpdateConnectorUsingPATCH(params *UpdateConnectorUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConnectorUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConnectorUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConnectorUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/ecx/v3/l3/connector/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateConnectorUsingPATCHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateConnectorUsingPATCHNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
