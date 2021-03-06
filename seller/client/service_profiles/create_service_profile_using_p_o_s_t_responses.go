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

// CreateServiceProfileUsingPOSTReader is a Reader for the CreateServiceProfileUsingPOST structure.
type CreateServiceProfileUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceProfileUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateServiceProfileUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateServiceProfileUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateServiceProfileUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateServiceProfileUsingPOSTOK creates a CreateServiceProfileUsingPOSTOK with default headers values
func NewCreateServiceProfileUsingPOSTOK() *CreateServiceProfileUsingPOSTOK {
	return &CreateServiceProfileUsingPOSTOK{}
}

/*CreateServiceProfileUsingPOSTOK handles this case with default header values.

success
*/
type CreateServiceProfileUsingPOSTOK struct {
	Payload *models.PostServiceprofilesResponse
}

func (o *CreateServiceProfileUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l2/serviceprofiles][%d] createServiceProfileUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateServiceProfileUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostServiceprofilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceProfileUsingPOSTBadRequest creates a CreateServiceProfileUsingPOSTBadRequest with default headers values
func NewCreateServiceProfileUsingPOSTBadRequest() *CreateServiceProfileUsingPOSTBadRequest {
	return &CreateServiceProfileUsingPOSTBadRequest{}
}

/*CreateServiceProfileUsingPOSTBadRequest handles this case with default header values.

Bad request
*/
type CreateServiceProfileUsingPOSTBadRequest struct {
	Payload models.ErrorResponseArray
}

func (o *CreateServiceProfileUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l2/serviceprofiles][%d] createServiceProfileUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceProfileUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceProfileUsingPOSTInternalServerError creates a CreateServiceProfileUsingPOSTInternalServerError with default headers values
func NewCreateServiceProfileUsingPOSTInternalServerError() *CreateServiceProfileUsingPOSTInternalServerError {
	return &CreateServiceProfileUsingPOSTInternalServerError{}
}

/*CreateServiceProfileUsingPOSTInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateServiceProfileUsingPOSTInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateServiceProfileUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ecx/v3/l2/serviceprofiles][%d] createServiceProfileUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateServiceProfileUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
