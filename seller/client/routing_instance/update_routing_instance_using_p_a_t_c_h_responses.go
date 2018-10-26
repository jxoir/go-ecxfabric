// Code generated by go-swagger; DO NOT EDIT.

package routing_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/seller/models"
)

// UpdateRoutingInstanceUsingPATCHReader is a Reader for the UpdateRoutingInstanceUsingPATCH structure.
type UpdateRoutingInstanceUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoutingInstanceUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateRoutingInstanceUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRoutingInstanceUsingPATCHBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRoutingInstanceUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateRoutingInstanceUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewUpdateRoutingInstanceUsingPATCHMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateRoutingInstanceUsingPATCHConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateRoutingInstanceUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewUpdateRoutingInstanceUsingPATCHServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRoutingInstanceUsingPATCHNoContent creates a UpdateRoutingInstanceUsingPATCHNoContent with default headers values
func NewUpdateRoutingInstanceUsingPATCHNoContent() *UpdateRoutingInstanceUsingPATCHNoContent {
	return &UpdateRoutingInstanceUsingPATCHNoContent{}
}

/*UpdateRoutingInstanceUsingPATCHNoContent handles this case with default header values.

OK
*/
type UpdateRoutingInstanceUsingPATCHNoContent struct {
}

func (o *UpdateRoutingInstanceUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHNoContent ", 204)
}

func (o *UpdateRoutingInstanceUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHBadRequest creates a UpdateRoutingInstanceUsingPATCHBadRequest with default headers values
func NewUpdateRoutingInstanceUsingPATCHBadRequest() *UpdateRoutingInstanceUsingPATCHBadRequest {
	return &UpdateRoutingInstanceUsingPATCHBadRequest{}
}

/*UpdateRoutingInstanceUsingPATCHBadRequest handles this case with default header values.

Bad request
*/
type UpdateRoutingInstanceUsingPATCHBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHUnauthorized creates a UpdateRoutingInstanceUsingPATCHUnauthorized with default headers values
func NewUpdateRoutingInstanceUsingPATCHUnauthorized() *UpdateRoutingInstanceUsingPATCHUnauthorized {
	return &UpdateRoutingInstanceUsingPATCHUnauthorized{}
}

/*UpdateRoutingInstanceUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateRoutingInstanceUsingPATCHUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHForbidden creates a UpdateRoutingInstanceUsingPATCHForbidden with default headers values
func NewUpdateRoutingInstanceUsingPATCHForbidden() *UpdateRoutingInstanceUsingPATCHForbidden {
	return &UpdateRoutingInstanceUsingPATCHForbidden{}
}

/*UpdateRoutingInstanceUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type UpdateRoutingInstanceUsingPATCHForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHMethodNotAllowed creates a UpdateRoutingInstanceUsingPATCHMethodNotAllowed with default headers values
func NewUpdateRoutingInstanceUsingPATCHMethodNotAllowed() *UpdateRoutingInstanceUsingPATCHMethodNotAllowed {
	return &UpdateRoutingInstanceUsingPATCHMethodNotAllowed{}
}

/*UpdateRoutingInstanceUsingPATCHMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type UpdateRoutingInstanceUsingPATCHMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHConflict creates a UpdateRoutingInstanceUsingPATCHConflict with default headers values
func NewUpdateRoutingInstanceUsingPATCHConflict() *UpdateRoutingInstanceUsingPATCHConflict {
	return &UpdateRoutingInstanceUsingPATCHConflict{}
}

/*UpdateRoutingInstanceUsingPATCHConflict handles this case with default header values.

Conflict
*/
type UpdateRoutingInstanceUsingPATCHConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHConflict) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHConflict  %+v", 409, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHInternalServerError creates a UpdateRoutingInstanceUsingPATCHInternalServerError with default headers values
func NewUpdateRoutingInstanceUsingPATCHInternalServerError() *UpdateRoutingInstanceUsingPATCHInternalServerError {
	return &UpdateRoutingInstanceUsingPATCHInternalServerError{}
}

/*UpdateRoutingInstanceUsingPATCHInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateRoutingInstanceUsingPATCHInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoutingInstanceUsingPATCHServiceUnavailable creates a UpdateRoutingInstanceUsingPATCHServiceUnavailable with default headers values
func NewUpdateRoutingInstanceUsingPATCHServiceUnavailable() *UpdateRoutingInstanceUsingPATCHServiceUnavailable {
	return &UpdateRoutingInstanceUsingPATCHServiceUnavailable{}
}

/*UpdateRoutingInstanceUsingPATCHServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type UpdateRoutingInstanceUsingPATCHServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *UpdateRoutingInstanceUsingPATCHServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /ecx/v3/l3/routinginstance/{uuid}][%d] updateRoutingInstanceUsingPATCHServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateRoutingInstanceUsingPATCHServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
