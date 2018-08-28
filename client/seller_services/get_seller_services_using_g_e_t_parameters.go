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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSellerServicesUsingGETParams creates a new GetSellerServicesUsingGETParams object
// with the default values initialized.
func NewGetSellerServicesUsingGETParams() *GetSellerServicesUsingGETParams {
	var (
		pageDefault  = int32(0)
		totalDefault = int32(900)
	)
	return &GetSellerServicesUsingGETParams{
		Page:  &pageDefault,
		Total: &totalDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSellerServicesUsingGETParamsWithTimeout creates a new GetSellerServicesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSellerServicesUsingGETParamsWithTimeout(timeout time.Duration) *GetSellerServicesUsingGETParams {
	var (
		pageDefault  = int32(0)
		totalDefault = int32(900)
	)
	return &GetSellerServicesUsingGETParams{
		Page:  &pageDefault,
		Total: &totalDefault,

		timeout: timeout,
	}
}

// NewGetSellerServicesUsingGETParamsWithContext creates a new GetSellerServicesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSellerServicesUsingGETParamsWithContext(ctx context.Context) *GetSellerServicesUsingGETParams {
	var (
		pageDefault  = int32(0)
		totalDefault = int32(900)
	)
	return &GetSellerServicesUsingGETParams{
		Page:  &pageDefault,
		Total: &totalDefault,

		Context: ctx,
	}
}

// NewGetSellerServicesUsingGETParamsWithHTTPClient creates a new GetSellerServicesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSellerServicesUsingGETParamsWithHTTPClient(client *http.Client) *GetSellerServicesUsingGETParams {
	var (
		pageDefault  = int32(0)
		totalDefault = int32(900)
	)
	return &GetSellerServicesUsingGETParams{
		Page:       &pageDefault,
		Total:      &totalDefault,
		HTTPClient: client,
	}
}

/*GetSellerServicesUsingGETParams contains all the parameters to send to the API endpoint
for the get seller services using g e t operation typically these are written to a http.Request
*/
type GetSellerServicesUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*Metros
	  metros

	*/
	Metros []string
	/*Page
	  page

	*/
	Page *int32
	/*Total
	  total

	*/
	Total *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithTimeout(timeout time.Duration) *GetSellerServicesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithContext(ctx context.Context) *GetSellerServicesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithHTTPClient(client *http.Client) *GetSellerServicesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithAuthorization(authorization string) *GetSellerServicesUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithMetros adds the metros to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithMetros(metros []string) *GetSellerServicesUsingGETParams {
	o.SetMetros(metros)
	return o
}

// SetMetros adds the metros to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetMetros(metros []string) {
	o.Metros = metros
}

// WithPage adds the page to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithPage(page *int32) *GetSellerServicesUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithTotal adds the total to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) WithTotal(total *int32) *GetSellerServicesUsingGETParams {
	o.SetTotal(total)
	return o
}

// SetTotal adds the total to the get seller services using g e t params
func (o *GetSellerServicesUsingGETParams) SetTotal(total *int32) {
	o.Total = total
}

// WriteToRequest writes these params to a swagger request
func (o *GetSellerServicesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	valuesMetros := o.Metros

	joinedMetros := swag.JoinByFormat(valuesMetros, "multi")
	// query array param metros
	if err := r.SetQueryParam("metros", joinedMetros...); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Total != nil {

		// query param total
		var qrTotal int32
		if o.Total != nil {
			qrTotal = *o.Total
		}
		qTotal := swag.FormatInt32(qrTotal)
		if qTotal != "" {
			if err := r.SetQueryParam("total", qTotal); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}