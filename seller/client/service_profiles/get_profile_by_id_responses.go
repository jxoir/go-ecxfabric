// Code generated by go-swagger; DO NOT EDIT.

package service_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/seller/models"
)

// GetProfileByIDReader is a Reader for the GetProfileByID structure.
type GetProfileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfileByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetProfileByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProfileByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetProfileByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetProfileByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetProfileByIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetProfileByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProfileByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetProfileByIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProfileByIDOK creates a GetProfileByIDOK with default headers values
func NewGetProfileByIDOK() *GetProfileByIDOK {
	return &GetProfileByIDOK{}
}

/*GetProfileByIDOK handles this case with default header values.

Success
*/
type GetProfileByIDOK struct {
	Payload *models.GETServiceProfileModelv3
}

func (o *GetProfileByIDOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdOK  %+v", 200, o.Payload)
}

func (o *GetProfileByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GETServiceProfileModelv3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDNoContent creates a GetProfileByIDNoContent with default headers values
func NewGetProfileByIDNoContent() *GetProfileByIDNoContent {
	return &GetProfileByIDNoContent{}
}

/*GetProfileByIDNoContent handles this case with default header values.

No Content
*/
type GetProfileByIDNoContent struct {
}

func (o *GetProfileByIDNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdNoContent ", 204)
}

func (o *GetProfileByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProfileByIDBadRequest creates a GetProfileByIDBadRequest with default headers values
func NewGetProfileByIDBadRequest() *GetProfileByIDBadRequest {
	return &GetProfileByIDBadRequest{}
}

/*GetProfileByIDBadRequest handles this case with default header values.

Bad request
*/
type GetProfileByIDBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetProfileByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDUnauthorized creates a GetProfileByIDUnauthorized with default headers values
func NewGetProfileByIDUnauthorized() *GetProfileByIDUnauthorized {
	return &GetProfileByIDUnauthorized{}
}

/*GetProfileByIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetProfileByIDUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProfileByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDForbidden creates a GetProfileByIDForbidden with default headers values
func NewGetProfileByIDForbidden() *GetProfileByIDForbidden {
	return &GetProfileByIDForbidden{}
}

/*GetProfileByIDForbidden handles this case with default header values.

Forbidden
*/
type GetProfileByIDForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetProfileByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDMethodNotAllowed creates a GetProfileByIDMethodNotAllowed with default headers values
func NewGetProfileByIDMethodNotAllowed() *GetProfileByIDMethodNotAllowed {
	return &GetProfileByIDMethodNotAllowed{}
}

/*GetProfileByIDMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetProfileByIDMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetProfileByIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDConflict creates a GetProfileByIDConflict with default headers values
func NewGetProfileByIDConflict() *GetProfileByIDConflict {
	return &GetProfileByIDConflict{}
}

/*GetProfileByIDConflict handles this case with default header values.

Conflict
*/
type GetProfileByIDConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDConflict) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdConflict  %+v", 409, o.Payload)
}

func (o *GetProfileByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDInternalServerError creates a GetProfileByIDInternalServerError with default headers values
func NewGetProfileByIDInternalServerError() *GetProfileByIDInternalServerError {
	return &GetProfileByIDInternalServerError{}
}

/*GetProfileByIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetProfileByIDInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProfileByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDServiceUnavailable creates a GetProfileByIDServiceUnavailable with default headers values
func NewGetProfileByIDServiceUnavailable() *GetProfileByIDServiceUnavailable {
	return &GetProfileByIDServiceUnavailable{}
}

/*GetProfileByIDServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetProfileByIDServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *GetProfileByIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l3/serviceprofiles/{uuid}][%d] getProfileByIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProfileByIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
