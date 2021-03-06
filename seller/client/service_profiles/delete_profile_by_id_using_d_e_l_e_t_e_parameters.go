// Code generated by go-swagger; DO NOT EDIT.

package service_profiles

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

// NewDeleteProfileByIDUsingDELETEParams creates a new DeleteProfileByIDUsingDELETEParams object
// with the default values initialized.
func NewDeleteProfileByIDUsingDELETEParams() *DeleteProfileByIDUsingDELETEParams {
	var ()
	return &DeleteProfileByIDUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProfileByIDUsingDELETEParamsWithTimeout creates a new DeleteProfileByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProfileByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteProfileByIDUsingDELETEParams {
	var ()
	return &DeleteProfileByIDUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteProfileByIDUsingDELETEParamsWithContext creates a new DeleteProfileByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProfileByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteProfileByIDUsingDELETEParams {
	var ()
	return &DeleteProfileByIDUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteProfileByIDUsingDELETEParamsWithHTTPClient creates a new DeleteProfileByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProfileByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteProfileByIDUsingDELETEParams {
	var ()
	return &DeleteProfileByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteProfileByIDUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete profile by Id using d e l e t e operation typically these are written to a http.Request
*/
type DeleteProfileByIDUsingDELETEParams struct {

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

// WithTimeout adds the timeout to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteProfileByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteProfileByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteProfileByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteProfileByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUUID adds the uuid to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) WithUUID(uuid string) *DeleteProfileByIDUsingDELETEParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete profile by Id using d e l e t e params
func (o *DeleteProfileByIDUsingDELETEParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProfileByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
