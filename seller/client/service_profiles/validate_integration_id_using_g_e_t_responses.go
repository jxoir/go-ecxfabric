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

// ValidateIntegrationIDUsingGETReader is a Reader for the ValidateIntegrationIDUsingGET structure.
type ValidateIntegrationIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateIntegrationIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateIntegrationIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewValidateIntegrationIDUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewValidateIntegrationIDUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewValidateIntegrationIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewValidateIntegrationIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateIntegrationIDUsingGETOK creates a ValidateIntegrationIDUsingGETOK with default headers values
func NewValidateIntegrationIDUsingGETOK() *ValidateIntegrationIDUsingGETOK {
	return &ValidateIntegrationIDUsingGETOK{}
}

/*ValidateIntegrationIDUsingGETOK handles this case with default header values.

Success
*/
type ValidateIntegrationIDUsingGETOK struct {
	Payload *models.ValidateIntegrationIDResponse
}

func (o *ValidateIntegrationIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/validateIntegrationId/{integrationId}][%d] validateIntegrationIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *ValidateIntegrationIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidateIntegrationIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateIntegrationIDUsingGETNoContent creates a ValidateIntegrationIDUsingGETNoContent with default headers values
func NewValidateIntegrationIDUsingGETNoContent() *ValidateIntegrationIDUsingGETNoContent {
	return &ValidateIntegrationIDUsingGETNoContent{}
}

/*ValidateIntegrationIDUsingGETNoContent handles this case with default header values.

No Content
*/
type ValidateIntegrationIDUsingGETNoContent struct {
}

func (o *ValidateIntegrationIDUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/validateIntegrationId/{integrationId}][%d] validateIntegrationIdUsingGETNoContent ", 204)
}

func (o *ValidateIntegrationIDUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateIntegrationIDUsingGETBadRequest creates a ValidateIntegrationIDUsingGETBadRequest with default headers values
func NewValidateIntegrationIDUsingGETBadRequest() *ValidateIntegrationIDUsingGETBadRequest {
	return &ValidateIntegrationIDUsingGETBadRequest{}
}

/*ValidateIntegrationIDUsingGETBadRequest handles this case with default header values.

Bad Request
*/
type ValidateIntegrationIDUsingGETBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ValidateIntegrationIDUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/validateIntegrationId/{integrationId}][%d] validateIntegrationIdUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateIntegrationIDUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateIntegrationIDUsingGETNotFound creates a ValidateIntegrationIDUsingGETNotFound with default headers values
func NewValidateIntegrationIDUsingGETNotFound() *ValidateIntegrationIDUsingGETNotFound {
	return &ValidateIntegrationIDUsingGETNotFound{}
}

/*ValidateIntegrationIDUsingGETNotFound handles this case with default header values.

Record Not Found
*/
type ValidateIntegrationIDUsingGETNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ValidateIntegrationIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/validateIntegrationId/{integrationId}][%d] validateIntegrationIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *ValidateIntegrationIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateIntegrationIDUsingGETInternalServerError creates a ValidateIntegrationIDUsingGETInternalServerError with default headers values
func NewValidateIntegrationIDUsingGETInternalServerError() *ValidateIntegrationIDUsingGETInternalServerError {
	return &ValidateIntegrationIDUsingGETInternalServerError{}
}

/*ValidateIntegrationIDUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type ValidateIntegrationIDUsingGETInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ValidateIntegrationIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/validateIntegrationId/{integrationId}][%d] validateIntegrationIdUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateIntegrationIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
