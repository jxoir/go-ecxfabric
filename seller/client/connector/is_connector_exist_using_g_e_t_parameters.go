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

// NewIsConnectorExistUsingGETParams creates a new IsConnectorExistUsingGETParams object
// with the default values initialized.
func NewIsConnectorExistUsingGETParams() *IsConnectorExistUsingGETParams {
	var ()
	return &IsConnectorExistUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsConnectorExistUsingGETParamsWithTimeout creates a new IsConnectorExistUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsConnectorExistUsingGETParamsWithTimeout(timeout time.Duration) *IsConnectorExistUsingGETParams {
	var ()
	return &IsConnectorExistUsingGETParams{

		timeout: timeout,
	}
}

// NewIsConnectorExistUsingGETParamsWithContext creates a new IsConnectorExistUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsConnectorExistUsingGETParamsWithContext(ctx context.Context) *IsConnectorExistUsingGETParams {
	var ()
	return &IsConnectorExistUsingGETParams{

		Context: ctx,
	}
}

// NewIsConnectorExistUsingGETParamsWithHTTPClient creates a new IsConnectorExistUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsConnectorExistUsingGETParamsWithHTTPClient(client *http.Client) *IsConnectorExistUsingGETParams {
	var ()
	return &IsConnectorExistUsingGETParams{
		HTTPClient: client,
	}
}

/*IsConnectorExistUsingGETParams contains all the parameters to send to the API endpoint
for the is connector exist using g e t operation typically these are written to a http.Request
*/
type IsConnectorExistUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*MetroCode
	  Metro Code

	*/
	MetroCode string
	/*Name
	  Connector Name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithTimeout(timeout time.Duration) *IsConnectorExistUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithContext(ctx context.Context) *IsConnectorExistUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithHTTPClient(client *http.Client) *IsConnectorExistUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithAuthorization(authorization string) *IsConnectorExistUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithMetroCode adds the metroCode to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithMetroCode(metroCode string) *IsConnectorExistUsingGETParams {
	o.SetMetroCode(metroCode)
	return o
}

// SetMetroCode adds the metroCode to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetMetroCode(metroCode string) {
	o.MetroCode = metroCode
}

// WithName adds the name to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) WithName(name string) *IsConnectorExistUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the is connector exist using g e t params
func (o *IsConnectorExistUsingGETParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *IsConnectorExistUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param metroCode
	if err := r.SetPathParam("metroCode", o.MetroCode); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
