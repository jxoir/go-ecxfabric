// Code generated by go-swagger; DO NOT EDIT.

package metros

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// GetMetrosUsingGETReader is a Reader for the GetMetrosUsingGET structure.
type GetMetrosUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetrosUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetrosUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetMetrosUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMetrosUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetMetrosUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMetrosUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMetrosUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMetrosUsingGETOK creates a GetMetrosUsingGETOK with default headers values
func NewGetMetrosUsingGETOK() *GetMetrosUsingGETOK {
	return &GetMetrosUsingGETOK{}
}

/*GetMetrosUsingGETOK handles this case with default header values.

success
*/
type GetMetrosUsingGETOK struct {
	Payload models.GETCommonMetroResp
}

func (o *GetMetrosUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetMetrosUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetrosUsingGETNoContent creates a GetMetrosUsingGETNoContent with default headers values
func NewGetMetrosUsingGETNoContent() *GetMetrosUsingGETNoContent {
	return &GetMetrosUsingGETNoContent{}
}

/*GetMetrosUsingGETNoContent handles this case with default header values.

No Content
*/
type GetMetrosUsingGETNoContent struct {
}

func (o *GetMetrosUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETNoContent ", 204)
}

func (o *GetMetrosUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMetrosUsingGETBadRequest creates a GetMetrosUsingGETBadRequest with default headers values
func NewGetMetrosUsingGETBadRequest() *GetMetrosUsingGETBadRequest {
	return &GetMetrosUsingGETBadRequest{}
}

/*GetMetrosUsingGETBadRequest handles this case with default header values.

Bad request
*/
type GetMetrosUsingGETBadRequest struct {
	Payload models.ErrorResponseArray
}

func (o *GetMetrosUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetMetrosUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetrosUsingGETUnauthorized creates a GetMetrosUsingGETUnauthorized with default headers values
func NewGetMetrosUsingGETUnauthorized() *GetMetrosUsingGETUnauthorized {
	return &GetMetrosUsingGETUnauthorized{}
}

/*GetMetrosUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetMetrosUsingGETUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetMetrosUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMetrosUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetrosUsingGETNotFound creates a GetMetrosUsingGETNotFound with default headers values
func NewGetMetrosUsingGETNotFound() *GetMetrosUsingGETNotFound {
	return &GetMetrosUsingGETNotFound{}
}

/*GetMetrosUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetMetrosUsingGETNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetMetrosUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetMetrosUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetrosUsingGETInternalServerError creates a GetMetrosUsingGETInternalServerError with default headers values
func NewGetMetrosUsingGETInternalServerError() *GetMetrosUsingGETInternalServerError {
	return &GetMetrosUsingGETInternalServerError{}
}

/*GetMetrosUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetMetrosUsingGETInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetMetrosUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/common/metros][%d] getMetrosUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMetrosUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}