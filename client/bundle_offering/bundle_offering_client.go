// Code generated by go-swagger; DO NOT EDIT.

package bundle_offering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new bundle offering API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bundle offering API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateBundleOrderUsingPOST creates bundle order

This API is used to create bundle order.
*/
func (a *Client) CreateBundleOrderUsingPOST(params *CreateBundleOrderUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBundleOrderUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBundleOrderUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBundleOrderUsingPOST",
		Method:             "POST",
		PathPattern:        "/ecx/v3/l3/bundle/{bundleCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBundleOrderUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateBundleOrderUsingPOSTOK), nil

}

/*
GetOfferingsUsingGET returns list of bundle offerings

This API returns list of bundle offerings
*/
func (a *Client) GetOfferingsUsingGET(params *GetOfferingsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetOfferingsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOfferingsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOfferingsUsingGET",
		Method:             "GET",
		PathPattern:        "/ecx/v3/l3/bundle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOfferingsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOfferingsUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}