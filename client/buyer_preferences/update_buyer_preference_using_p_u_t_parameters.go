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

	models "github.com/jxoir/go-ecxfabric/models"
)

// NewUpdateBuyerPreferenceUsingPUTParams creates a new UpdateBuyerPreferenceUsingPUTParams object
// with the default values initialized.
func NewUpdateBuyerPreferenceUsingPUTParams() *UpdateBuyerPreferenceUsingPUTParams {
	var ()
	return &UpdateBuyerPreferenceUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBuyerPreferenceUsingPUTParamsWithTimeout creates a new UpdateBuyerPreferenceUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBuyerPreferenceUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateBuyerPreferenceUsingPUTParams {
	var ()
	return &UpdateBuyerPreferenceUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateBuyerPreferenceUsingPUTParamsWithContext creates a new UpdateBuyerPreferenceUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBuyerPreferenceUsingPUTParamsWithContext(ctx context.Context) *UpdateBuyerPreferenceUsingPUTParams {
	var ()
	return &UpdateBuyerPreferenceUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateBuyerPreferenceUsingPUTParamsWithHTTPClient creates a new UpdateBuyerPreferenceUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBuyerPreferenceUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateBuyerPreferenceUsingPUTParams {
	var ()
	return &UpdateBuyerPreferenceUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateBuyerPreferenceUsingPUTParams contains all the parameters to send to the API endpoint
for the update buyer preference using p u t operation typically these are written to a http.Request
*/
type UpdateBuyerPreferenceUsingPUTParams struct {

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

// WithTimeout adds the timeout to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateBuyerPreferenceUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) WithContext(ctx context.Context) *UpdateBuyerPreferenceUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateBuyerPreferenceUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) WithAuthorization(authorization string) *UpdateBuyerPreferenceUsingPUTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) WithRequest(request *models.BuyerPreferenceModel) *UpdateBuyerPreferenceUsingPUTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update buyer preference using p u t params
func (o *UpdateBuyerPreferenceUsingPUTParams) SetRequest(request *models.BuyerPreferenceModel) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBuyerPreferenceUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
