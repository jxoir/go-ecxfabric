// Code generated by go-swagger; DO NOT EDIT.

package connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteConnectorUsingDELETEParams creates a new DeleteConnectorUsingDELETEParams object
// with the default values initialized.
func NewDeleteConnectorUsingDELETEParams() *DeleteConnectorUsingDELETEParams {
	var ()
	return &DeleteConnectorUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConnectorUsingDELETEParamsWithTimeout creates a new DeleteConnectorUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConnectorUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteConnectorUsingDELETEParams {
	var ()
	return &DeleteConnectorUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteConnectorUsingDELETEParamsWithContext creates a new DeleteConnectorUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConnectorUsingDELETEParamsWithContext(ctx context.Context) *DeleteConnectorUsingDELETEParams {
	var ()
	return &DeleteConnectorUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteConnectorUsingDELETEParamsWithHTTPClient creates a new DeleteConnectorUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConnectorUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteConnectorUsingDELETEParams {
	var ()
	return &DeleteConnectorUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteConnectorUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete connector using d e l e t e operation typically these are written to a http.Request
*/
type DeleteConnectorUsingDELETEParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*UUID
	  uuid

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteConnectorUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) WithContext(ctx context.Context) *DeleteConnectorUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteConnectorUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) WithAuthorization(authorization string) *DeleteConnectorUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUUID adds the uuid to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) WithUUID(uuid string) *DeleteConnectorUsingDELETEParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete connector using d e l e t e params
func (o *DeleteConnectorUsingDELETEParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConnectorUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
