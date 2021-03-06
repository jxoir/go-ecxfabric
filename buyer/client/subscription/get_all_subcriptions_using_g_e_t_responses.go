// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// GetAllSubcriptionsUsingGETReader is a Reader for the GetAllSubcriptionsUsingGET structure.
type GetAllSubcriptionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSubcriptionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllSubcriptionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllSubcriptionsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAllSubcriptionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllSubcriptionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetAllSubcriptionsUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetAllSubcriptionsUsingGETConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllSubcriptionsUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAllSubcriptionsUsingGETServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllSubcriptionsUsingGETOK creates a GetAllSubcriptionsUsingGETOK with default headers values
func NewGetAllSubcriptionsUsingGETOK() *GetAllSubcriptionsUsingGETOK {
	return &GetAllSubcriptionsUsingGETOK{}
}

/*GetAllSubcriptionsUsingGETOK handles this case with default header values.

Success
*/
type GetAllSubcriptionsUsingGETOK struct {
	Payload *models.SubscriptionGetResponse
}

func (o *GetAllSubcriptionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETBadRequest creates a GetAllSubcriptionsUsingGETBadRequest with default headers values
func NewGetAllSubcriptionsUsingGETBadRequest() *GetAllSubcriptionsUsingGETBadRequest {
	return &GetAllSubcriptionsUsingGETBadRequest{}
}

/*GetAllSubcriptionsUsingGETBadRequest handles this case with default header values.

Bad request
*/
type GetAllSubcriptionsUsingGETBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETUnauthorized creates a GetAllSubcriptionsUsingGETUnauthorized with default headers values
func NewGetAllSubcriptionsUsingGETUnauthorized() *GetAllSubcriptionsUsingGETUnauthorized {
	return &GetAllSubcriptionsUsingGETUnauthorized{}
}

/*GetAllSubcriptionsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllSubcriptionsUsingGETUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETForbidden creates a GetAllSubcriptionsUsingGETForbidden with default headers values
func NewGetAllSubcriptionsUsingGETForbidden() *GetAllSubcriptionsUsingGETForbidden {
	return &GetAllSubcriptionsUsingGETForbidden{}
}

/*GetAllSubcriptionsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllSubcriptionsUsingGETForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETMethodNotAllowed creates a GetAllSubcriptionsUsingGETMethodNotAllowed with default headers values
func NewGetAllSubcriptionsUsingGETMethodNotAllowed() *GetAllSubcriptionsUsingGETMethodNotAllowed {
	return &GetAllSubcriptionsUsingGETMethodNotAllowed{}
}

/*GetAllSubcriptionsUsingGETMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAllSubcriptionsUsingGETMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETConflict creates a GetAllSubcriptionsUsingGETConflict with default headers values
func NewGetAllSubcriptionsUsingGETConflict() *GetAllSubcriptionsUsingGETConflict {
	return &GetAllSubcriptionsUsingGETConflict{}
}

/*GetAllSubcriptionsUsingGETConflict handles this case with default header values.

Conflict
*/
type GetAllSubcriptionsUsingGETConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETConflict  %+v", 409, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETInternalServerError creates a GetAllSubcriptionsUsingGETInternalServerError with default headers values
func NewGetAllSubcriptionsUsingGETInternalServerError() *GetAllSubcriptionsUsingGETInternalServerError {
	return &GetAllSubcriptionsUsingGETInternalServerError{}
}

/*GetAllSubcriptionsUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAllSubcriptionsUsingGETInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubcriptionsUsingGETServiceUnavailable creates a GetAllSubcriptionsUsingGETServiceUnavailable with default headers values
func NewGetAllSubcriptionsUsingGETServiceUnavailable() *GetAllSubcriptionsUsingGETServiceUnavailable {
	return &GetAllSubcriptionsUsingGETServiceUnavailable{}
}

/*GetAllSubcriptionsUsingGETServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetAllSubcriptionsUsingGETServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *GetAllSubcriptionsUsingGETServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/subscription][%d] getAllSubcriptionsUsingGETServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAllSubcriptionsUsingGETServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
