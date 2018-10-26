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

// NewGetProfilesByMetroUsingGETParams creates a new GetProfilesByMetroUsingGETParams object
// with the default values initialized.
func NewGetProfilesByMetroUsingGETParams() *GetProfilesByMetroUsingGETParams {
	var (
		pageNumberDefault = int32(0)
		pageSizeDefault   = int32(20)
	)
	return &GetProfilesByMetroUsingGETParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfilesByMetroUsingGETParamsWithTimeout creates a new GetProfilesByMetroUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfilesByMetroUsingGETParamsWithTimeout(timeout time.Duration) *GetProfilesByMetroUsingGETParams {
	var (
		pageNumberDefault = int32(0)
		pageSizeDefault   = int32(20)
	)
	return &GetProfilesByMetroUsingGETParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetProfilesByMetroUsingGETParamsWithContext creates a new GetProfilesByMetroUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfilesByMetroUsingGETParamsWithContext(ctx context.Context) *GetProfilesByMetroUsingGETParams {
	var (
		pageNumberDefault = int32(0)
		pageSizeDefault   = int32(20)
	)
	return &GetProfilesByMetroUsingGETParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetProfilesByMetroUsingGETParamsWithHTTPClient creates a new GetProfilesByMetroUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfilesByMetroUsingGETParamsWithHTTPClient(client *http.Client) *GetProfilesByMetroUsingGETParams {
	var (
		pageNumberDefault = int32(0)
		pageSizeDefault   = int32(20)
	)
	return &GetProfilesByMetroUsingGETParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetProfilesByMetroUsingGETParams contains all the parameters to send to the API endpoint
for the get profiles by metro using g e t operation typically these are written to a http.Request
*/
type GetProfilesByMetroUsingGETParams struct {

	/*Authorization
	  Specify the OAuth Bearer token with prefix 'Bearer '.

	*/
	Authorization string
	/*MetroCode
	  metroCode

	*/
	MetroCode []string
	/*PageNumber
	  page number

	*/
	PageNumber *int32
	/*PageSize
	  total number of records

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithTimeout(timeout time.Duration) *GetProfilesByMetroUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithContext(ctx context.Context) *GetProfilesByMetroUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithHTTPClient(client *http.Client) *GetProfilesByMetroUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithAuthorization(authorization string) *GetProfilesByMetroUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithMetroCode adds the metroCode to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithMetroCode(metroCode []string) *GetProfilesByMetroUsingGETParams {
	o.SetMetroCode(metroCode)
	return o
}

// SetMetroCode adds the metroCode to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetMetroCode(metroCode []string) {
	o.MetroCode = metroCode
}

// WithPageNumber adds the pageNumber to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithPageNumber(pageNumber *int32) *GetProfilesByMetroUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) WithPageSize(pageSize *int32) *GetProfilesByMetroUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get profiles by metro using g e t params
func (o *GetProfilesByMetroUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilesByMetroUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	valuesMetroCode := o.MetroCode

	joinedMetroCode := swag.JoinByFormat(valuesMetroCode, "multi")
	// query array param metroCode
	if err := r.SetQueryParam("metroCode", joinedMetroCode...); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}