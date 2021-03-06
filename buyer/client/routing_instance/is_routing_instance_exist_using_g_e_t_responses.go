// Code generated by go-swagger; DO NOT EDIT.

package routing_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// IsRoutingInstanceExistUsingGETReader is a Reader for the IsRoutingInstanceExistUsingGET structure.
type IsRoutingInstanceExistUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsRoutingInstanceExistUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIsRoutingInstanceExistUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewIsRoutingInstanceExistUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewIsRoutingInstanceExistUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewIsRoutingInstanceExistUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewIsRoutingInstanceExistUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewIsRoutingInstanceExistUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewIsRoutingInstanceExistUsingGETConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewIsRoutingInstanceExistUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewIsRoutingInstanceExistUsingGETServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsRoutingInstanceExistUsingGETOK creates a IsRoutingInstanceExistUsingGETOK with default headers values
func NewIsRoutingInstanceExistUsingGETOK() *IsRoutingInstanceExistUsingGETOK {
	return &IsRoutingInstanceExistUsingGETOK{}
}

/*IsRoutingInstanceExistUsingGETOK handles this case with default header values.

success
*/
type IsRoutingInstanceExistUsingGETOK struct {
	Payload *models.RoutingInstanceExistenceResponse
}

func (o *IsRoutingInstanceExistUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingInstanceExistenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETNoContent creates a IsRoutingInstanceExistUsingGETNoContent with default headers values
func NewIsRoutingInstanceExistUsingGETNoContent() *IsRoutingInstanceExistUsingGETNoContent {
	return &IsRoutingInstanceExistUsingGETNoContent{}
}

/*IsRoutingInstanceExistUsingGETNoContent handles this case with default header values.

No Content
*/
type IsRoutingInstanceExistUsingGETNoContent struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETNoContent  %+v", 204, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETBadRequest creates a IsRoutingInstanceExistUsingGETBadRequest with default headers values
func NewIsRoutingInstanceExistUsingGETBadRequest() *IsRoutingInstanceExistUsingGETBadRequest {
	return &IsRoutingInstanceExistUsingGETBadRequest{}
}

/*IsRoutingInstanceExistUsingGETBadRequest handles this case with default header values.

Bad request
*/
type IsRoutingInstanceExistUsingGETBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETUnauthorized creates a IsRoutingInstanceExistUsingGETUnauthorized with default headers values
func NewIsRoutingInstanceExistUsingGETUnauthorized() *IsRoutingInstanceExistUsingGETUnauthorized {
	return &IsRoutingInstanceExistUsingGETUnauthorized{}
}

/*IsRoutingInstanceExistUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsRoutingInstanceExistUsingGETUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETForbidden creates a IsRoutingInstanceExistUsingGETForbidden with default headers values
func NewIsRoutingInstanceExistUsingGETForbidden() *IsRoutingInstanceExistUsingGETForbidden {
	return &IsRoutingInstanceExistUsingGETForbidden{}
}

/*IsRoutingInstanceExistUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsRoutingInstanceExistUsingGETForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETMethodNotAllowed creates a IsRoutingInstanceExistUsingGETMethodNotAllowed with default headers values
func NewIsRoutingInstanceExistUsingGETMethodNotAllowed() *IsRoutingInstanceExistUsingGETMethodNotAllowed {
	return &IsRoutingInstanceExistUsingGETMethodNotAllowed{}
}

/*IsRoutingInstanceExistUsingGETMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type IsRoutingInstanceExistUsingGETMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETConflict creates a IsRoutingInstanceExistUsingGETConflict with default headers values
func NewIsRoutingInstanceExistUsingGETConflict() *IsRoutingInstanceExistUsingGETConflict {
	return &IsRoutingInstanceExistUsingGETConflict{}
}

/*IsRoutingInstanceExistUsingGETConflict handles this case with default header values.

Conflict
*/
type IsRoutingInstanceExistUsingGETConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETConflict  %+v", 409, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETInternalServerError creates a IsRoutingInstanceExistUsingGETInternalServerError with default headers values
func NewIsRoutingInstanceExistUsingGETInternalServerError() *IsRoutingInstanceExistUsingGETInternalServerError {
	return &IsRoutingInstanceExistUsingGETInternalServerError{}
}

/*IsRoutingInstanceExistUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type IsRoutingInstanceExistUsingGETInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsRoutingInstanceExistUsingGETServiceUnavailable creates a IsRoutingInstanceExistUsingGETServiceUnavailable with default headers values
func NewIsRoutingInstanceExistUsingGETServiceUnavailable() *IsRoutingInstanceExistUsingGETServiceUnavailable {
	return &IsRoutingInstanceExistUsingGETServiceUnavailable{}
}

/*IsRoutingInstanceExistUsingGETServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type IsRoutingInstanceExistUsingGETServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *IsRoutingInstanceExistUsingGETServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/routinginstance/exist/{metroCode}/{name}][%d] isRoutingInstanceExistUsingGETServiceUnavailable  %+v", 503, o.Payload)
}

func (o *IsRoutingInstanceExistUsingGETServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
