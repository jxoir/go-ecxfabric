// Code generated by go-swagger; DO NOT EDIT.

package connections

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

	models "github.com/jxoir/go-ecxfabric/seller/models"
)

// NewPerformUserActionUsingPATCHParams creates a new PerformUserActionUsingPATCHParams object
// with the default values initialized.
func NewPerformUserActionUsingPATCHParams() *PerformUserActionUsingPATCHParams {
	var ()
	return &PerformUserActionUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPerformUserActionUsingPATCHParamsWithTimeout creates a new PerformUserActionUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPerformUserActionUsingPATCHParamsWithTimeout(timeout time.Duration) *PerformUserActionUsingPATCHParams {
	var ()
	return &PerformUserActionUsingPATCHParams{

		timeout: timeout,
	}
}

// NewPerformUserActionUsingPATCHParamsWithContext creates a new PerformUserActionUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewPerformUserActionUsingPATCHParamsWithContext(ctx context.Context) *PerformUserActionUsingPATCHParams {
	var ()
	return &PerformUserActionUsingPATCHParams{

		Context: ctx,
	}
}

// NewPerformUserActionUsingPATCHParamsWithHTTPClient creates a new PerformUserActionUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPerformUserActionUsingPATCHParamsWithHTTPClient(client *http.Client) *PerformUserActionUsingPATCHParams {
	var ()
	return &PerformUserActionUsingPATCHParams{
		HTTPClient: client,
	}
}

/*PerformUserActionUsingPATCHParams contains all the parameters to send to the API endpoint
for the perform user action using p a t c h operation typically these are written to a http.Request
*/
type PerformUserActionUsingPATCHParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*Action
	  action

	*/
	Action string
	/*ConnID
	  connId

	*/
	ConnID string
	/*Request
	  request

	*/
	Request *models.PatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithTimeout(timeout time.Duration) *PerformUserActionUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithContext(ctx context.Context) *PerformUserActionUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithHTTPClient(client *http.Client) *PerformUserActionUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithAuthorization(authorization string) *PerformUserActionUsingPATCHParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAction adds the action to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithAction(action string) *PerformUserActionUsingPATCHParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetAction(action string) {
	o.Action = action
}

// WithConnID adds the connID to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithConnID(connID string) *PerformUserActionUsingPATCHParams {
	o.SetConnID(connID)
	return o
}

// SetConnID adds the connId to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetConnID(connID string) {
	o.ConnID = connID
}

// WithRequest adds the request to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) WithRequest(request *models.PatchRequest) *PerformUserActionUsingPATCHParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the perform user action using p a t c h params
func (o *PerformUserActionUsingPATCHParams) SetRequest(request *models.PatchRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PerformUserActionUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {
		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
	}

	// path param connId
	if err := r.SetPathParam("connId", o.ConnID); err != nil {
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
