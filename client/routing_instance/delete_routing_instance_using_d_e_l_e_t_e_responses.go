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

// DeleteRoutingInstanceUsingDELETEReader is a Reader for the DeleteRoutingInstanceUsingDELETE structure.
type DeleteRoutingInstanceUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingInstanceUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRoutingInstanceUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteRoutingInstanceUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteRoutingInstanceUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteRoutingInstanceUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteRoutingInstanceUsingDELETEMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteRoutingInstanceUsingDELETEConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteRoutingInstanceUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteRoutingInstanceUsingDELETEServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoutingInstanceUsingDELETENoContent creates a DeleteRoutingInstanceUsingDELETENoContent with default headers values
func NewDeleteRoutingInstanceUsingDELETENoContent() *DeleteRoutingInstanceUsingDELETENoContent {
	return &DeleteRoutingInstanceUsingDELETENoContent{}
}

/*DeleteRoutingInstanceUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteRoutingInstanceUsingDELETENoContent struct {
}

func (o *DeleteRoutingInstanceUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETENoContent ", 204)
}

func (o *DeleteRoutingInstanceUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEBadRequest creates a DeleteRoutingInstanceUsingDELETEBadRequest with default headers values
func NewDeleteRoutingInstanceUsingDELETEBadRequest() *DeleteRoutingInstanceUsingDELETEBadRequest {
	return &DeleteRoutingInstanceUsingDELETEBadRequest{}
}

/*DeleteRoutingInstanceUsingDELETEBadRequest handles this case with default header values.

Bad request
*/
type DeleteRoutingInstanceUsingDELETEBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEUnauthorized creates a DeleteRoutingInstanceUsingDELETEUnauthorized with default headers values
func NewDeleteRoutingInstanceUsingDELETEUnauthorized() *DeleteRoutingInstanceUsingDELETEUnauthorized {
	return &DeleteRoutingInstanceUsingDELETEUnauthorized{}
}

/*DeleteRoutingInstanceUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRoutingInstanceUsingDELETEUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEForbidden creates a DeleteRoutingInstanceUsingDELETEForbidden with default headers values
func NewDeleteRoutingInstanceUsingDELETEForbidden() *DeleteRoutingInstanceUsingDELETEForbidden {
	return &DeleteRoutingInstanceUsingDELETEForbidden{}
}

/*DeleteRoutingInstanceUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteRoutingInstanceUsingDELETEForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEMethodNotAllowed creates a DeleteRoutingInstanceUsingDELETEMethodNotAllowed with default headers values
func NewDeleteRoutingInstanceUsingDELETEMethodNotAllowed() *DeleteRoutingInstanceUsingDELETEMethodNotAllowed {
	return &DeleteRoutingInstanceUsingDELETEMethodNotAllowed{}
}

/*DeleteRoutingInstanceUsingDELETEMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type DeleteRoutingInstanceUsingDELETEMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEConflict creates a DeleteRoutingInstanceUsingDELETEConflict with default headers values
func NewDeleteRoutingInstanceUsingDELETEConflict() *DeleteRoutingInstanceUsingDELETEConflict {
	return &DeleteRoutingInstanceUsingDELETEConflict{}
}

/*DeleteRoutingInstanceUsingDELETEConflict handles this case with default header values.

Conflict
*/
type DeleteRoutingInstanceUsingDELETEConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEConflict) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEConflict  %+v", 409, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEInternalServerError creates a DeleteRoutingInstanceUsingDELETEInternalServerError with default headers values
func NewDeleteRoutingInstanceUsingDELETEInternalServerError() *DeleteRoutingInstanceUsingDELETEInternalServerError {
	return &DeleteRoutingInstanceUsingDELETEInternalServerError{}
}

/*DeleteRoutingInstanceUsingDELETEInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteRoutingInstanceUsingDELETEInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingInstanceUsingDELETEServiceUnavailable creates a DeleteRoutingInstanceUsingDELETEServiceUnavailable with default headers values
func NewDeleteRoutingInstanceUsingDELETEServiceUnavailable() *DeleteRoutingInstanceUsingDELETEServiceUnavailable {
	return &DeleteRoutingInstanceUsingDELETEServiceUnavailable{}
}

/*DeleteRoutingInstanceUsingDELETEServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type DeleteRoutingInstanceUsingDELETEServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *DeleteRoutingInstanceUsingDELETEServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /ecx/v3/l3/routinginstance/{uuid}][%d] deleteRoutingInstanceUsingDELETEServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingInstanceUsingDELETEServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
