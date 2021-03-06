// Code generated by go-swagger; DO NOT EDIT.

package connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// CreateConnectorUsingPOSTReader is a Reader for the CreateConnectorUsingPOST structure.
type CreateConnectorUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConnectorUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateConnectorUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateConnectorUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateConnectorUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateConnectorUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateConnectorUsingPOSTMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateConnectorUsingPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateConnectorUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateConnectorUsingPOSTServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateConnectorUsingPOSTCreated creates a CreateConnectorUsingPOSTCreated with default headers values
func NewCreateConnectorUsingPOSTCreated() *CreateConnectorUsingPOSTCreated {
	return &CreateConnectorUsingPOSTCreated{}
}

/*CreateConnectorUsingPOSTCreated handles this case with default header values.

created
*/
type CreateConnectorUsingPOSTCreated struct {
	Payload *models.ConnectorCreateResponse
}

func (o *CreateConnectorUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateConnectorUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectorCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTBadRequest creates a CreateConnectorUsingPOSTBadRequest with default headers values
func NewCreateConnectorUsingPOSTBadRequest() *CreateConnectorUsingPOSTBadRequest {
	return &CreateConnectorUsingPOSTBadRequest{}
}

/*CreateConnectorUsingPOSTBadRequest handles this case with default header values.

Bad request
*/
type CreateConnectorUsingPOSTBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConnectorUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTUnauthorized creates a CreateConnectorUsingPOSTUnauthorized with default headers values
func NewCreateConnectorUsingPOSTUnauthorized() *CreateConnectorUsingPOSTUnauthorized {
	return &CreateConnectorUsingPOSTUnauthorized{}
}

/*CreateConnectorUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateConnectorUsingPOSTUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateConnectorUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTForbidden creates a CreateConnectorUsingPOSTForbidden with default headers values
func NewCreateConnectorUsingPOSTForbidden() *CreateConnectorUsingPOSTForbidden {
	return &CreateConnectorUsingPOSTForbidden{}
}

/*CreateConnectorUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateConnectorUsingPOSTForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTForbidden  %+v", 403, o.Payload)
}

func (o *CreateConnectorUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTMethodNotAllowed creates a CreateConnectorUsingPOSTMethodNotAllowed with default headers values
func NewCreateConnectorUsingPOSTMethodNotAllowed() *CreateConnectorUsingPOSTMethodNotAllowed {
	return &CreateConnectorUsingPOSTMethodNotAllowed{}
}

/*CreateConnectorUsingPOSTMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type CreateConnectorUsingPOSTMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateConnectorUsingPOSTMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTConflict creates a CreateConnectorUsingPOSTConflict with default headers values
func NewCreateConnectorUsingPOSTConflict() *CreateConnectorUsingPOSTConflict {
	return &CreateConnectorUsingPOSTConflict{}
}

/*CreateConnectorUsingPOSTConflict handles this case with default header values.

Conflict
*/
type CreateConnectorUsingPOSTConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTConflict  %+v", 409, o.Payload)
}

func (o *CreateConnectorUsingPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTInternalServerError creates a CreateConnectorUsingPOSTInternalServerError with default headers values
func NewCreateConnectorUsingPOSTInternalServerError() *CreateConnectorUsingPOSTInternalServerError {
	return &CreateConnectorUsingPOSTInternalServerError{}
}

/*CreateConnectorUsingPOSTInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateConnectorUsingPOSTInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateConnectorUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectorUsingPOSTServiceUnavailable creates a CreateConnectorUsingPOSTServiceUnavailable with default headers values
func NewCreateConnectorUsingPOSTServiceUnavailable() *CreateConnectorUsingPOSTServiceUnavailable {
	return &CreateConnectorUsingPOSTServiceUnavailable{}
}

/*CreateConnectorUsingPOSTServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type CreateConnectorUsingPOSTServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateConnectorUsingPOSTServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/connector][%d] createConnectorUsingPOSTServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateConnectorUsingPOSTServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
