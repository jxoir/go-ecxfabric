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

// NewDeleteProfileByProfileNameOrUUIDUsingDELETEParams creates a new DeleteProfileByProfileNameOrUUIDUsingDELETEParams object
// with the default values initialized.
func NewDeleteProfileByProfileNameOrUUIDUsingDELETEParams() *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	var ()
	return &DeleteProfileByProfileNameOrUUIDUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithTimeout creates a new DeleteProfileByProfileNameOrUUIDUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	var ()
	return &DeleteProfileByProfileNameOrUUIDUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithContext creates a new DeleteProfileByProfileNameOrUUIDUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	var ()
	return &DeleteProfileByProfileNameOrUUIDUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithHTTPClient creates a new DeleteProfileByProfileNameOrUUIDUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProfileByProfileNameOrUUIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	var ()
	return &DeleteProfileByProfileNameOrUUIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteProfileByProfileNameOrUUIDUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete profile by profile name or UUID using d e l e t e operation typically these are written to a http.Request
*/
type DeleteProfileByProfileNameOrUUIDUsingDELETEParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*UUID
	  uuids

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUUID adds the uuid to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WithUUID(uuid string) *DeleteProfileByProfileNameOrUUIDUsingDELETEParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete profile by profile name or UUID using d e l e t e params
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProfileByProfileNameOrUUIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
