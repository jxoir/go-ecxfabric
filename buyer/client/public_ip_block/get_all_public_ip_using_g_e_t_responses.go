// Code generated by go-swagger; DO NOT EDIT.

package public_ip_block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// GetAllPublicIPUsingGETReader is a Reader for the GetAllPublicIPUsingGET structure.
type GetAllPublicIPUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllPublicIPUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllPublicIPUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllPublicIPUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAllPublicIPUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllPublicIPUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetAllPublicIPUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetAllPublicIPUsingGETConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllPublicIPUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAllPublicIPUsingGETServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllPublicIPUsingGETOK creates a GetAllPublicIPUsingGETOK with default headers values
func NewGetAllPublicIPUsingGETOK() *GetAllPublicIPUsingGETOK {
	return &GetAllPublicIPUsingGETOK{}
}

/*GetAllPublicIPUsingGETOK handles this case with default header values.

Success
*/
type GetAllPublicIPUsingGETOK struct {
	Payload *models.IPBlockGetResponse
}

func (o *GetAllPublicIPUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllPublicIPUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPBlockGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETBadRequest creates a GetAllPublicIPUsingGETBadRequest with default headers values
func NewGetAllPublicIPUsingGETBadRequest() *GetAllPublicIPUsingGETBadRequest {
	return &GetAllPublicIPUsingGETBadRequest{}
}

/*GetAllPublicIPUsingGETBadRequest handles this case with default header values.

Bad request
*/
type GetAllPublicIPUsingGETBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllPublicIPUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETUnauthorized creates a GetAllPublicIPUsingGETUnauthorized with default headers values
func NewGetAllPublicIPUsingGETUnauthorized() *GetAllPublicIPUsingGETUnauthorized {
	return &GetAllPublicIPUsingGETUnauthorized{}
}

/*GetAllPublicIPUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllPublicIPUsingGETUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllPublicIPUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETForbidden creates a GetAllPublicIPUsingGETForbidden with default headers values
func NewGetAllPublicIPUsingGETForbidden() *GetAllPublicIPUsingGETForbidden {
	return &GetAllPublicIPUsingGETForbidden{}
}

/*GetAllPublicIPUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllPublicIPUsingGETForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *GetAllPublicIPUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETMethodNotAllowed creates a GetAllPublicIPUsingGETMethodNotAllowed with default headers values
func NewGetAllPublicIPUsingGETMethodNotAllowed() *GetAllPublicIPUsingGETMethodNotAllowed {
	return &GetAllPublicIPUsingGETMethodNotAllowed{}
}

/*GetAllPublicIPUsingGETMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAllPublicIPUsingGETMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAllPublicIPUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETConflict creates a GetAllPublicIPUsingGETConflict with default headers values
func NewGetAllPublicIPUsingGETConflict() *GetAllPublicIPUsingGETConflict {
	return &GetAllPublicIPUsingGETConflict{}
}

/*GetAllPublicIPUsingGETConflict handles this case with default header values.

Conflict
*/
type GetAllPublicIPUsingGETConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETConflict  %+v", 409, o.Payload)
}

func (o *GetAllPublicIPUsingGETConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETInternalServerError creates a GetAllPublicIPUsingGETInternalServerError with default headers values
func NewGetAllPublicIPUsingGETInternalServerError() *GetAllPublicIPUsingGETInternalServerError {
	return &GetAllPublicIPUsingGETInternalServerError{}
}

/*GetAllPublicIPUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAllPublicIPUsingGETInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllPublicIPUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPublicIPUsingGETServiceUnavailable creates a GetAllPublicIPUsingGETServiceUnavailable with default headers values
func NewGetAllPublicIPUsingGETServiceUnavailable() *GetAllPublicIPUsingGETServiceUnavailable {
	return &GetAllPublicIPUsingGETServiceUnavailable{}
}

/*GetAllPublicIPUsingGETServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetAllPublicIPUsingGETServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllPublicIPUsingGETServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/publicIpAddress][%d] getAllPublicIpUsingGETServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAllPublicIPUsingGETServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
