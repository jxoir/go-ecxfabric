// Code generated by go-swagger; DO NOT EDIT.

package seller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/models"
)

// GetProfileByIDUsingGETReader is a Reader for the GetProfileByIDUsingGET structure.
type GetProfileByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfileByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetProfileByIDUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProfileByIDUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProfileByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProfileByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProfileByIDUsingGETOK creates a GetProfileByIDUsingGETOK with default headers values
func NewGetProfileByIDUsingGETOK() *GetProfileByIDUsingGETOK {
	return &GetProfileByIDUsingGETOK{}
}

/*GetProfileByIDUsingGETOK handles this case with default header values.

Success
*/
type GetProfileByIDUsingGETOK struct {
	Payload *models.GetServProfServicesResp
}

func (o *GetProfileByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services/{uuid}][%d] getProfileByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetProfileByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServProfServicesResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDUsingGETNoContent creates a GetProfileByIDUsingGETNoContent with default headers values
func NewGetProfileByIDUsingGETNoContent() *GetProfileByIDUsingGETNoContent {
	return &GetProfileByIDUsingGETNoContent{}
}

/*GetProfileByIDUsingGETNoContent handles this case with default header values.

No Content
*/
type GetProfileByIDUsingGETNoContent struct {
}

func (o *GetProfileByIDUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services/{uuid}][%d] getProfileByIdUsingGETNoContent ", 204)
}

func (o *GetProfileByIDUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProfileByIDUsingGETBadRequest creates a GetProfileByIDUsingGETBadRequest with default headers values
func NewGetProfileByIDUsingGETBadRequest() *GetProfileByIDUsingGETBadRequest {
	return &GetProfileByIDUsingGETBadRequest{}
}

/*GetProfileByIDUsingGETBadRequest handles this case with default header values.

Bad Request
*/
type GetProfileByIDUsingGETBadRequest struct {
	Payload models.ErrorResponseArray
}

func (o *GetProfileByIDUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services/{uuid}][%d] getProfileByIdUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetProfileByIDUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDUsingGETNotFound creates a GetProfileByIDUsingGETNotFound with default headers values
func NewGetProfileByIDUsingGETNotFound() *GetProfileByIDUsingGETNotFound {
	return &GetProfileByIDUsingGETNotFound{}
}

/*GetProfileByIDUsingGETNotFound handles this case with default header values.

Record Not Found
*/
type GetProfileByIDUsingGETNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetProfileByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services/{uuid}][%d] getProfileByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetProfileByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByIDUsingGETInternalServerError creates a GetProfileByIDUsingGETInternalServerError with default headers values
func NewGetProfileByIDUsingGETInternalServerError() *GetProfileByIDUsingGETInternalServerError {
	return &GetProfileByIDUsingGETInternalServerError{}
}

/*GetProfileByIDUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetProfileByIDUsingGETInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetProfileByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services/{uuid}][%d] getProfileByIdUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProfileByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}