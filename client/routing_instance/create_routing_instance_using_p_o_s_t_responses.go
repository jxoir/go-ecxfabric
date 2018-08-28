// Code generated by go-swagger; DO NOT EDIT.

package routing_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/models"
)

// CreateRoutingInstanceUsingPOSTReader is a Reader for the CreateRoutingInstanceUsingPOST structure.
type CreateRoutingInstanceUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoutingInstanceUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateRoutingInstanceUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewCreateRoutingInstanceUsingPOSTNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateRoutingInstanceUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateRoutingInstanceUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateRoutingInstanceUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateRoutingInstanceUsingPOSTMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateRoutingInstanceUsingPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateRoutingInstanceUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateRoutingInstanceUsingPOSTServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRoutingInstanceUsingPOSTCreated creates a CreateRoutingInstanceUsingPOSTCreated with default headers values
func NewCreateRoutingInstanceUsingPOSTCreated() *CreateRoutingInstanceUsingPOSTCreated {
	return &CreateRoutingInstanceUsingPOSTCreated{}
}

/*CreateRoutingInstanceUsingPOSTCreated handles this case with default header values.

created
*/
type CreateRoutingInstanceUsingPOSTCreated struct {
	Payload *models.RoutingInstanceCreateResponse
}

func (o *CreateRoutingInstanceUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingInstanceCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTNoContent creates a CreateRoutingInstanceUsingPOSTNoContent with default headers values
func NewCreateRoutingInstanceUsingPOSTNoContent() *CreateRoutingInstanceUsingPOSTNoContent {
	return &CreateRoutingInstanceUsingPOSTNoContent{}
}

/*CreateRoutingInstanceUsingPOSTNoContent handles this case with default header values.

No Content
*/
type CreateRoutingInstanceUsingPOSTNoContent struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTNoContent) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTNoContent  %+v", 204, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTBadRequest creates a CreateRoutingInstanceUsingPOSTBadRequest with default headers values
func NewCreateRoutingInstanceUsingPOSTBadRequest() *CreateRoutingInstanceUsingPOSTBadRequest {
	return &CreateRoutingInstanceUsingPOSTBadRequest{}
}

/*CreateRoutingInstanceUsingPOSTBadRequest handles this case with default header values.

Bad request
*/
type CreateRoutingInstanceUsingPOSTBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTUnauthorized creates a CreateRoutingInstanceUsingPOSTUnauthorized with default headers values
func NewCreateRoutingInstanceUsingPOSTUnauthorized() *CreateRoutingInstanceUsingPOSTUnauthorized {
	return &CreateRoutingInstanceUsingPOSTUnauthorized{}
}

/*CreateRoutingInstanceUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRoutingInstanceUsingPOSTUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTForbidden creates a CreateRoutingInstanceUsingPOSTForbidden with default headers values
func NewCreateRoutingInstanceUsingPOSTForbidden() *CreateRoutingInstanceUsingPOSTForbidden {
	return &CreateRoutingInstanceUsingPOSTForbidden{}
}

/*CreateRoutingInstanceUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateRoutingInstanceUsingPOSTForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTForbidden  %+v", 403, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTMethodNotAllowed creates a CreateRoutingInstanceUsingPOSTMethodNotAllowed with default headers values
func NewCreateRoutingInstanceUsingPOSTMethodNotAllowed() *CreateRoutingInstanceUsingPOSTMethodNotAllowed {
	return &CreateRoutingInstanceUsingPOSTMethodNotAllowed{}
}

/*CreateRoutingInstanceUsingPOSTMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type CreateRoutingInstanceUsingPOSTMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTConflict creates a CreateRoutingInstanceUsingPOSTConflict with default headers values
func NewCreateRoutingInstanceUsingPOSTConflict() *CreateRoutingInstanceUsingPOSTConflict {
	return &CreateRoutingInstanceUsingPOSTConflict{}
}

/*CreateRoutingInstanceUsingPOSTConflict handles this case with default header values.

Conflict
*/
type CreateRoutingInstanceUsingPOSTConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTConflict  %+v", 409, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTInternalServerError creates a CreateRoutingInstanceUsingPOSTInternalServerError with default headers values
func NewCreateRoutingInstanceUsingPOSTInternalServerError() *CreateRoutingInstanceUsingPOSTInternalServerError {
	return &CreateRoutingInstanceUsingPOSTInternalServerError{}
}

/*CreateRoutingInstanceUsingPOSTInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateRoutingInstanceUsingPOSTInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoutingInstanceUsingPOSTServiceUnavailable creates a CreateRoutingInstanceUsingPOSTServiceUnavailable with default headers values
func NewCreateRoutingInstanceUsingPOSTServiceUnavailable() *CreateRoutingInstanceUsingPOSTServiceUnavailable {
	return &CreateRoutingInstanceUsingPOSTServiceUnavailable{}
}

/*CreateRoutingInstanceUsingPOSTServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type CreateRoutingInstanceUsingPOSTServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateRoutingInstanceUsingPOSTServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/routinginstance][%d] createRoutingInstanceUsingPOSTServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateRoutingInstanceUsingPOSTServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
