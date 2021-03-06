// Code generated by go-swagger; DO NOT EDIT.

package routing_instance

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

// NewIsRoutingInstanceExistUsingGETParams creates a new IsRoutingInstanceExistUsingGETParams object
// with the default values initialized.
func NewIsRoutingInstanceExistUsingGETParams() *IsRoutingInstanceExistUsingGETParams {
	var ()
	return &IsRoutingInstanceExistUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsRoutingInstanceExistUsingGETParamsWithTimeout creates a new IsRoutingInstanceExistUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsRoutingInstanceExistUsingGETParamsWithTimeout(timeout time.Duration) *IsRoutingInstanceExistUsingGETParams {
	var ()
	return &IsRoutingInstanceExistUsingGETParams{

		timeout: timeout,
	}
}

// NewIsRoutingInstanceExistUsingGETParamsWithContext creates a new IsRoutingInstanceExistUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsRoutingInstanceExistUsingGETParamsWithContext(ctx context.Context) *IsRoutingInstanceExistUsingGETParams {
	var ()
	return &IsRoutingInstanceExistUsingGETParams{

		Context: ctx,
	}
}

// NewIsRoutingInstanceExistUsingGETParamsWithHTTPClient creates a new IsRoutingInstanceExistUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsRoutingInstanceExistUsingGETParamsWithHTTPClient(client *http.Client) *IsRoutingInstanceExistUsingGETParams {
	var ()
	return &IsRoutingInstanceExistUsingGETParams{
		HTTPClient: client,
	}
}

/*IsRoutingInstanceExistUsingGETParams contains all the parameters to send to the API endpoint
for the is routing instance exist using g e t operation typically these are written to a http.Request
*/
type IsRoutingInstanceExistUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*MetroCode
	  Metro Code

	*/
	MetroCode string
	/*Name
	  Routing Instance Name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithTimeout(timeout time.Duration) *IsRoutingInstanceExistUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithContext(ctx context.Context) *IsRoutingInstanceExistUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithHTTPClient(client *http.Client) *IsRoutingInstanceExistUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithAuthorization(authorization string) *IsRoutingInstanceExistUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithMetroCode adds the metroCode to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithMetroCode(metroCode string) *IsRoutingInstanceExistUsingGETParams {
	o.SetMetroCode(metroCode)
	return o
}

// SetMetroCode adds the metroCode to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetMetroCode(metroCode string) {
	o.MetroCode = metroCode
}

// WithName adds the name to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) WithName(name string) *IsRoutingInstanceExistUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the is routing instance exist using g e t params
func (o *IsRoutingInstanceExistUsingGETParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *IsRoutingInstanceExistUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
