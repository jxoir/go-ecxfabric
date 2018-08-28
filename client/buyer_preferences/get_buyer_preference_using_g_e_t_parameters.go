// Code generated by go-swagger; DO NOT EDIT.

package buyer_preferences

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

// NewGetBuyerPreferenceUsingGETParams creates a new GetBuyerPreferenceUsingGETParams object
// with the default values initialized.
func NewGetBuyerPreferenceUsingGETParams() *GetBuyerPreferenceUsingGETParams {
	var ()
	return &GetBuyerPreferenceUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuyerPreferenceUsingGETParamsWithTimeout creates a new GetBuyerPreferenceUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuyerPreferenceUsingGETParamsWithTimeout(timeout time.Duration) *GetBuyerPreferenceUsingGETParams {
	var ()
	return &GetBuyerPreferenceUsingGETParams{

		timeout: timeout,
	}
}

// NewGetBuyerPreferenceUsingGETParamsWithContext creates a new GetBuyerPreferenceUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuyerPreferenceUsingGETParamsWithContext(ctx context.Context) *GetBuyerPreferenceUsingGETParams {
	var ()
	return &GetBuyerPreferenceUsingGETParams{

		Context: ctx,
	}
}

// NewGetBuyerPreferenceUsingGETParamsWithHTTPClient creates a new GetBuyerPreferenceUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuyerPreferenceUsingGETParamsWithHTTPClient(client *http.Client) *GetBuyerPreferenceUsingGETParams {
	var ()
	return &GetBuyerPreferenceUsingGETParams{
		HTTPClient: client,
	}
}

/*GetBuyerPreferenceUsingGETParams contains all the parameters to send to the API endpoint
for the get buyer preference using g e t operation typically these are written to a http.Request
*/
type GetBuyerPreferenceUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) WithTimeout(timeout time.Duration) *GetBuyerPreferenceUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) WithContext(ctx context.Context) *GetBuyerPreferenceUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) WithHTTPClient(client *http.Client) *GetBuyerPreferenceUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) WithAuthorization(authorization string) *GetBuyerPreferenceUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get buyer preference using g e t params
func (o *GetBuyerPreferenceUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuyerPreferenceUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
