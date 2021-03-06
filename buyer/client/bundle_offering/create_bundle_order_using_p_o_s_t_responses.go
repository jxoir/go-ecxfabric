// Code generated by go-swagger; DO NOT EDIT.

package bundle_offering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// CreateBundleOrderUsingPOSTReader is a Reader for the CreateBundleOrderUsingPOST structure.
type CreateBundleOrderUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBundleOrderUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateBundleOrderUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateBundleOrderUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateBundleOrderUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateBundleOrderUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateBundleOrderUsingPOSTMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateBundleOrderUsingPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateBundleOrderUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateBundleOrderUsingPOSTServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateBundleOrderUsingPOSTOK creates a CreateBundleOrderUsingPOSTOK with default headers values
func NewCreateBundleOrderUsingPOSTOK() *CreateBundleOrderUsingPOSTOK {
	return &CreateBundleOrderUsingPOSTOK{}
}

/*CreateBundleOrderUsingPOSTOK handles this case with default header values.

success
*/
type CreateBundleOrderUsingPOSTOK struct {
	Payload *models.BundleResponse
}

func (o *CreateBundleOrderUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BundleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTBadRequest creates a CreateBundleOrderUsingPOSTBadRequest with default headers values
func NewCreateBundleOrderUsingPOSTBadRequest() *CreateBundleOrderUsingPOSTBadRequest {
	return &CreateBundleOrderUsingPOSTBadRequest{}
}

/*CreateBundleOrderUsingPOSTBadRequest handles this case with default header values.

Bad request
*/
type CreateBundleOrderUsingPOSTBadRequest struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTUnauthorized creates a CreateBundleOrderUsingPOSTUnauthorized with default headers values
func NewCreateBundleOrderUsingPOSTUnauthorized() *CreateBundleOrderUsingPOSTUnauthorized {
	return &CreateBundleOrderUsingPOSTUnauthorized{}
}

/*CreateBundleOrderUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateBundleOrderUsingPOSTUnauthorized struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTForbidden creates a CreateBundleOrderUsingPOSTForbidden with default headers values
func NewCreateBundleOrderUsingPOSTForbidden() *CreateBundleOrderUsingPOSTForbidden {
	return &CreateBundleOrderUsingPOSTForbidden{}
}

/*CreateBundleOrderUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateBundleOrderUsingPOSTForbidden struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTForbidden  %+v", 403, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTMethodNotAllowed creates a CreateBundleOrderUsingPOSTMethodNotAllowed with default headers values
func NewCreateBundleOrderUsingPOSTMethodNotAllowed() *CreateBundleOrderUsingPOSTMethodNotAllowed {
	return &CreateBundleOrderUsingPOSTMethodNotAllowed{}
}

/*CreateBundleOrderUsingPOSTMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type CreateBundleOrderUsingPOSTMethodNotAllowed struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTConflict creates a CreateBundleOrderUsingPOSTConflict with default headers values
func NewCreateBundleOrderUsingPOSTConflict() *CreateBundleOrderUsingPOSTConflict {
	return &CreateBundleOrderUsingPOSTConflict{}
}

/*CreateBundleOrderUsingPOSTConflict handles this case with default header values.

Conflict
*/
type CreateBundleOrderUsingPOSTConflict struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTConflict  %+v", 409, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTInternalServerError creates a CreateBundleOrderUsingPOSTInternalServerError with default headers values
func NewCreateBundleOrderUsingPOSTInternalServerError() *CreateBundleOrderUsingPOSTInternalServerError {
	return &CreateBundleOrderUsingPOSTInternalServerError{}
}

/*CreateBundleOrderUsingPOSTInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateBundleOrderUsingPOSTInternalServerError struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleOrderUsingPOSTServiceUnavailable creates a CreateBundleOrderUsingPOSTServiceUnavailable with default headers values
func NewCreateBundleOrderUsingPOSTServiceUnavailable() *CreateBundleOrderUsingPOSTServiceUnavailable {
	return &CreateBundleOrderUsingPOSTServiceUnavailable{}
}

/*CreateBundleOrderUsingPOSTServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type CreateBundleOrderUsingPOSTServiceUnavailable struct {
	Payload models.GeneralErrorMessage
}

func (o *CreateBundleOrderUsingPOSTServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l3/bundle/{bundleCode}][%d] createBundleOrderUsingPOSTServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateBundleOrderUsingPOSTServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
