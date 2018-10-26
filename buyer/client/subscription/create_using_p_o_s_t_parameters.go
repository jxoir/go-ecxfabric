// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewCreateUsingPOSTParams creates a new CreateUsingPOSTParams object
// with the default values initialized.
func NewCreateUsingPOSTParams() *CreateUsingPOSTParams {
	var ()
	return &CreateUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOSTParamsWithTimeout creates a new CreateUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateUsingPOSTParams {
	var ()
	return &CreateUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateUsingPOSTParamsWithContext creates a new CreateUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUsingPOSTParamsWithContext(ctx context.Context) *CreateUsingPOSTParams {
	var ()
	return &CreateUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateUsingPOSTParamsWithHTTPClient creates a new CreateUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateUsingPOSTParams {
	var ()
	return &CreateUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateUsingPOSTParams contains all the parameters to send to the API endpoint
for the create using p o s t operation typically these are written to a http.Request
*/
type CreateUsingPOSTParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*Request
	  request

	*/
	Request *models.SubscriptionBundleOrdering

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create using p o s t params
func (o *CreateUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t params
func (o *CreateUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t params
func (o *CreateUsingPOSTParams) WithContext(ctx context.Context) *CreateUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t params
func (o *CreateUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t params
func (o *CreateUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t params
func (o *CreateUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create using p o s t params
func (o *CreateUsingPOSTParams) WithAuthorization(authorization string) *CreateUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create using p o s t params
func (o *CreateUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the create using p o s t params
func (o *CreateUsingPOSTParams) WithRequest(request *models.SubscriptionBundleOrdering) *CreateUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create using p o s t params
func (o *CreateUsingPOSTParams) SetRequest(request *models.SubscriptionBundleOrdering) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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