// Code generated by go-swagger; DO NOT EDIT.

package connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/seller/models"
)

// IsConnectorExistUsingGETReader is a Reader for the IsConnectorExistUsingGET structure.
type IsConnectorExistUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsConnectorExistUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIsConnectorExistUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewIsConnectorExistUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewIsConnectorExistUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewIsConnectorExistUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewIsConnectorExistUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewIsConnectorExistUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewIsConnectorExistUsingGETConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewIsConnectorExistUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewIsConnectorExistUsingGETServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsConnectorExistUsingGETOK creates a IsConnectorExistUsingGETOK with default headers values
func NewIsConnectorExistUsingGETOK() *IsConnectorExistUsingGETOK {
	return &IsConnectorExistUsingGETOK{}
}

/*IsConnectorExistUsingGETOK handles this case with default header values.

success
*/
type IsConnectorExistUsingGETOK struct {
	Payload *models.ConnectorExistenceResponse
}

func (o *IsConnectorExistUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsConnectorExistUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectorExistenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETNoContent creates a IsConnectorExistUsingGETNoContent with default headers values
func NewIsConnectorExistUsingGETNoContent() *IsConnectorExistUsingGETNoContent {
	return &IsConnectorExistUsingGETNoContent{}
}

/*IsConnectorExistUsingGETNoContent handles this case with default header values.

No Content
*/
type IsConnectorExistUsingGETNoContent struct {
}

func (o *IsConnectorExistUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETNoContent ", 204)
}

func (o *IsConnectorExistUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsConnectorExistUsingGETBadRequest creates a IsConnectorExistUsingGETBadRequest with default headers values
func NewIsConnectorExistUsingGETBadRequest() *IsConnectorExistUsingGETBadRequest {
	return &IsConnectorExistUsingGETBadRequest{}
}

/*IsConnectorExistUsingGETBadRequest handles this case with default header values.

Bad request
*/
type IsConnectorExistUsingGETBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *IsConnectorExistUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETUnauthorized creates a IsConnectorExistUsingGETUnauthorized with default headers values
func NewIsConnectorExistUsingGETUnauthorized() *IsConnectorExistUsingGETUnauthorized {
	return &IsConnectorExistUsingGETUnauthorized{}
}

/*IsConnectorExistUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsConnectorExistUsingGETUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *IsConnectorExistUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETForbidden creates a IsConnectorExistUsingGETForbidden with default headers values
func NewIsConnectorExistUsingGETForbidden() *IsConnectorExistUsingGETForbidden {
	return &IsConnectorExistUsingGETForbidden{}
}

/*IsConnectorExistUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsConnectorExistUsingGETForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *IsConnectorExistUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETMethodNotAllowed creates a IsConnectorExistUsingGETMethodNotAllowed with default headers values
func NewIsConnectorExistUsingGETMethodNotAllowed() *IsConnectorExistUsingGETMethodNotAllowed {
	return &IsConnectorExistUsingGETMethodNotAllowed{}
}

/*IsConnectorExistUsingGETMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type IsConnectorExistUsingGETMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *IsConnectorExistUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETConflict creates a IsConnectorExistUsingGETConflict with default headers values
func NewIsConnectorExistUsingGETConflict() *IsConnectorExistUsingGETConflict {
	return &IsConnectorExistUsingGETConflict{}
}

/*IsConnectorExistUsingGETConflict handles this case with default header values.

Conflict
*/
type IsConnectorExistUsingGETConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETConflict  %+v", 409, o.Payload)
}

func (o *IsConnectorExistUsingGETConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETInternalServerError creates a IsConnectorExistUsingGETInternalServerError with default headers values
func NewIsConnectorExistUsingGETInternalServerError() *IsConnectorExistUsingGETInternalServerError {
	return &IsConnectorExistUsingGETInternalServerError{}
}

/*IsConnectorExistUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type IsConnectorExistUsingGETInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *IsConnectorExistUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsConnectorExistUsingGETServiceUnavailable creates a IsConnectorExistUsingGETServiceUnavailable with default headers values
func NewIsConnectorExistUsingGETServiceUnavailable() *IsConnectorExistUsingGETServiceUnavailable {
	return &IsConnectorExistUsingGETServiceUnavailable{}
}

/*IsConnectorExistUsingGETServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type IsConnectorExistUsingGETServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *IsConnectorExistUsingGETServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/connector/exist/{metroCode}/{name}][%d] isConnectorExistUsingGETServiceUnavailable  %+v", 503, o.Payload)
}

func (o *IsConnectorExistUsingGETServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
