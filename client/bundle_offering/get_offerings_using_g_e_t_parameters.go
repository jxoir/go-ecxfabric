// Code generated by go-swagger; DO NOT EDIT.

package bundle_offering

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

// NewGetOfferingsUsingGETParams creates a new GetOfferingsUsingGETParams object
// with the default values initialized.
func NewGetOfferingsUsingGETParams() *GetOfferingsUsingGETParams {
	var ()
	return &GetOfferingsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOfferingsUsingGETParamsWithTimeout creates a new GetOfferingsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOfferingsUsingGETParamsWithTimeout(timeout time.Duration) *GetOfferingsUsingGETParams {
	var ()
	return &GetOfferingsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetOfferingsUsingGETParamsWithContext creates a new GetOfferingsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOfferingsUsingGETParamsWithContext(ctx context.Context) *GetOfferingsUsingGETParams {
	var ()
	return &GetOfferingsUsingGETParams{

		Context: ctx,
	}
}

// NewGetOfferingsUsingGETParamsWithHTTPClient creates a new GetOfferingsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOfferingsUsingGETParamsWithHTTPClient(client *http.Client) *GetOfferingsUsingGETParams {
	var ()
	return &GetOfferingsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetOfferingsUsingGETParams contains all the parameters to send to the API endpoint
for the get offerings using g e t operation typically these are written to a http.Request
*/
type GetOfferingsUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) WithTimeout(timeout time.Duration) *GetOfferingsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) WithContext(ctx context.Context) *GetOfferingsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) WithHTTPClient(client *http.Client) *GetOfferingsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) WithAuthorization(authorization string) *GetOfferingsUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get offerings using g e t params
func (o *GetOfferingsUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetOfferingsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
