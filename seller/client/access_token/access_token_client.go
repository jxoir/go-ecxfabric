// Code generated by go-swagger; DO NOT EDIT.

package access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new access token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccessToken returns oauth2 access token

The ECX Fabric APIs use a custom per request authentication system. All calls to the API must use the API Oauth key that can be obtained by calling the authentication API. In the case of an incorrect key, the appropriate error message will be returned.
*/
func (a *Client) GetAccessToken(params *GetAccessTokenParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccessTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessToken",
		Method:             "POST",
		PathPattern:        "/oauth2/v1/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccessTokenOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
