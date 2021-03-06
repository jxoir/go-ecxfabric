// Code generated by go-swagger; DO NOT EDIT.

package seller_services

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

// NewGetProfileByIDUsingGETParams creates a new GetProfileByIDUsingGETParams object
// with the default values initialized.
func NewGetProfileByIDUsingGETParams() *GetProfileByIDUsingGETParams {
	var ()
	return &GetProfileByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfileByIDUsingGETParamsWithTimeout creates a new GetProfileByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfileByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetProfileByIDUsingGETParams {
	var ()
	return &GetProfileByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetProfileByIDUsingGETParamsWithContext creates a new GetProfileByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfileByIDUsingGETParamsWithContext(ctx context.Context) *GetProfileByIDUsingGETParams {
	var ()
	return &GetProfileByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetProfileByIDUsingGETParamsWithHTTPClient creates a new GetProfileByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfileByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetProfileByIDUsingGETParams {
	var ()
	return &GetProfileByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetProfileByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get profile by Id using g e t operation typically these are written to a http.Request
*/
type GetProfileByIDUsingGETParams struct {

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

// WithTimeout adds the timeout to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetProfileByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) WithContext(ctx context.Context) *GetProfileByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetProfileByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) WithAuthorization(authorization string) *GetProfileByIDUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUUID adds the uuid to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) WithUUID(uuid string) *GetProfileByIDUsingGETParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get profile by Id using g e t params
func (o *GetProfileByIDUsingGETParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfileByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
