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

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// NewSaveBuyerPreferenceUsingPOSTParams creates a new SaveBuyerPreferenceUsingPOSTParams object
// with the default values initialized.
func NewSaveBuyerPreferenceUsingPOSTParams() *SaveBuyerPreferenceUsingPOSTParams {
	var ()
	return &SaveBuyerPreferenceUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveBuyerPreferenceUsingPOSTParamsWithTimeout creates a new SaveBuyerPreferenceUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveBuyerPreferenceUsingPOSTParamsWithTimeout(timeout time.Duration) *SaveBuyerPreferenceUsingPOSTParams {
	var ()
	return &SaveBuyerPreferenceUsingPOSTParams{

		timeout: timeout,
	}
}

// NewSaveBuyerPreferenceUsingPOSTParamsWithContext creates a new SaveBuyerPreferenceUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveBuyerPreferenceUsingPOSTParamsWithContext(ctx context.Context) *SaveBuyerPreferenceUsingPOSTParams {
	var ()
	return &SaveBuyerPreferenceUsingPOSTParams{

		Context: ctx,
	}
}

// NewSaveBuyerPreferenceUsingPOSTParamsWithHTTPClient creates a new SaveBuyerPreferenceUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveBuyerPreferenceUsingPOSTParamsWithHTTPClient(client *http.Client) *SaveBuyerPreferenceUsingPOSTParams {
	var ()
	return &SaveBuyerPreferenceUsingPOSTParams{
		HTTPClient: client,
	}
}

/*SaveBuyerPreferenceUsingPOSTParams contains all the parameters to send to the API endpoint
for the save buyer preference using p o s t operation typically these are written to a http.Request
*/
type SaveBuyerPreferenceUsingPOSTParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*Request
	  request

	*/
	Request *models.BuyerPreferenceModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) WithTimeout(timeout time.Duration) *SaveBuyerPreferenceUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) WithContext(ctx context.Context) *SaveBuyerPreferenceUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) WithHTTPClient(client *http.Client) *SaveBuyerPreferenceUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) WithAuthorization(authorization string) *SaveBuyerPreferenceUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) WithRequest(request *models.BuyerPreferenceModel) *SaveBuyerPreferenceUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the save buyer preference using p o s t params
func (o *SaveBuyerPreferenceUsingPOSTParams) SetRequest(request *models.BuyerPreferenceModel) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SaveBuyerPreferenceUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
