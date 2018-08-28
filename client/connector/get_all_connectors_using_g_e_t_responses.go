// Code generated by go-swagger; DO NOT EDIT.

package connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/models"
)

// GetAllConnectorsUsingGETReader is a Reader for the GetAllConnectorsUsingGET structure.
type GetAllConnectorsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllConnectorsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllConnectorsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllConnectorsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAllConnectorsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllConnectorsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetAllConnectorsUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetAllConnectorsUsingGETConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllConnectorsUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAllConnectorsUsingGETServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllConnectorsUsingGETOK creates a GetAllConnectorsUsingGETOK with default headers values
func NewGetAllConnectorsUsingGETOK() *GetAllConnectorsUsingGETOK {
	return &GetAllConnectorsUsingGETOK{}
}

/*GetAllConnectorsUsingGETOK handles this case with default header values.

success
*/
type GetAllConnectorsUsingGETOK struct {
	Payload *models.GetConnectorsResponse
}

func (o *GetAllConnectorsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllConnectorsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetConnectorsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETBadRequest creates a GetAllConnectorsUsingGETBadRequest with default headers values
func NewGetAllConnectorsUsingGETBadRequest() *GetAllConnectorsUsingGETBadRequest {
	return &GetAllConnectorsUsingGETBadRequest{}
}

/*GetAllConnectorsUsingGETBadRequest handles this case with default header values.

Bad request
*/
type GetAllConnectorsUsingGETBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllConnectorsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETUnauthorized creates a GetAllConnectorsUsingGETUnauthorized with default headers values
func NewGetAllConnectorsUsingGETUnauthorized() *GetAllConnectorsUsingGETUnauthorized {
	return &GetAllConnectorsUsingGETUnauthorized{}
}

/*GetAllConnectorsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllConnectorsUsingGETUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllConnectorsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETForbidden creates a GetAllConnectorsUsingGETForbidden with default headers values
func NewGetAllConnectorsUsingGETForbidden() *GetAllConnectorsUsingGETForbidden {
	return &GetAllConnectorsUsingGETForbidden{}
}

/*GetAllConnectorsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllConnectorsUsingGETForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *GetAllConnectorsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETMethodNotAllowed creates a GetAllConnectorsUsingGETMethodNotAllowed with default headers values
func NewGetAllConnectorsUsingGETMethodNotAllowed() *GetAllConnectorsUsingGETMethodNotAllowed {
	return &GetAllConnectorsUsingGETMethodNotAllowed{}
}

/*GetAllConnectorsUsingGETMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAllConnectorsUsingGETMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAllConnectorsUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETConflict creates a GetAllConnectorsUsingGETConflict with default headers values
func NewGetAllConnectorsUsingGETConflict() *GetAllConnectorsUsingGETConflict {
	return &GetAllConnectorsUsingGETConflict{}
}

/*GetAllConnectorsUsingGETConflict handles this case with default header values.

Conflict
*/
type GetAllConnectorsUsingGETConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETConflict  %+v", 409, o.Payload)
}

func (o *GetAllConnectorsUsingGETConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETInternalServerError creates a GetAllConnectorsUsingGETInternalServerError with default headers values
func NewGetAllConnectorsUsingGETInternalServerError() *GetAllConnectorsUsingGETInternalServerError {
	return &GetAllConnectorsUsingGETInternalServerError{}
}

/*GetAllConnectorsUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAllConnectorsUsingGETInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllConnectorsUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllConnectorsUsingGETServiceUnavailable creates a GetAllConnectorsUsingGETServiceUnavailable with default headers values
func NewGetAllConnectorsUsingGETServiceUnavailable() *GetAllConnectorsUsingGETServiceUnavailable {
	return &GetAllConnectorsUsingGETServiceUnavailable{}
}

/*GetAllConnectorsUsingGETServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetAllConnectorsUsingGETServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllConnectorsUsingGETServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector][%d] getAllConnectorsUsingGETServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAllConnectorsUsingGETServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}