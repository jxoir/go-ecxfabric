// Code generated by go-swagger; DO NOT EDIT.

package seller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/buyer/models"
)

// GetProfilesByMetroUsingGETReader is a Reader for the GetProfilesByMetroUsingGET structure.
type GetProfilesByMetroUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfilesByMetroUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfilesByMetroUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetProfilesByMetroUsingGETNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProfilesByMetroUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProfilesByMetroUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProfilesByMetroUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProfilesByMetroUsingGETOK creates a GetProfilesByMetroUsingGETOK with default headers values
func NewGetProfilesByMetroUsingGETOK() *GetProfilesByMetroUsingGETOK {
	return &GetProfilesByMetroUsingGETOK{}
}

/*GetProfilesByMetroUsingGETOK handles this case with default header values.

Success
*/
type GetProfilesByMetroUsingGETOK struct {
	Payload *models.GetServProfServicesResp
}

func (o *GetProfilesByMetroUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services][%d] getProfilesByMetroUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetProfilesByMetroUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServProfServicesResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesByMetroUsingGETNoContent creates a GetProfilesByMetroUsingGETNoContent with default headers values
func NewGetProfilesByMetroUsingGETNoContent() *GetProfilesByMetroUsingGETNoContent {
	return &GetProfilesByMetroUsingGETNoContent{}
}

/*GetProfilesByMetroUsingGETNoContent handles this case with default header values.

No Content
*/
type GetProfilesByMetroUsingGETNoContent struct {
}

func (o *GetProfilesByMetroUsingGETNoContent) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services][%d] getProfilesByMetroUsingGETNoContent ", 204)
}

func (o *GetProfilesByMetroUsingGETNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProfilesByMetroUsingGETBadRequest creates a GetProfilesByMetroUsingGETBadRequest with default headers values
func NewGetProfilesByMetroUsingGETBadRequest() *GetProfilesByMetroUsingGETBadRequest {
	return &GetProfilesByMetroUsingGETBadRequest{}
}

/*GetProfilesByMetroUsingGETBadRequest handles this case with default header values.

Bad Request
*/
type GetProfilesByMetroUsingGETBadRequest struct {
	Payload models.ErrorResponseArray
}

func (o *GetProfilesByMetroUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services][%d] getProfilesByMetroUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetProfilesByMetroUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesByMetroUsingGETNotFound creates a GetProfilesByMetroUsingGETNotFound with default headers values
func NewGetProfilesByMetroUsingGETNotFound() *GetProfilesByMetroUsingGETNotFound {
	return &GetProfilesByMetroUsingGETNotFound{}
}

/*GetProfilesByMetroUsingGETNotFound handles this case with default header values.

Record Not Found
*/
type GetProfilesByMetroUsingGETNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetProfilesByMetroUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services][%d] getProfilesByMetroUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetProfilesByMetroUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesByMetroUsingGETInternalServerError creates a GetProfilesByMetroUsingGETInternalServerError with default headers values
func NewGetProfilesByMetroUsingGETInternalServerError() *GetProfilesByMetroUsingGETInternalServerError {
	return &GetProfilesByMetroUsingGETInternalServerError{}
}

/*GetProfilesByMetroUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetProfilesByMetroUsingGETInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetProfilesByMetroUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ecx/v3/l2/serviceprofiles/services][%d] getProfilesByMetroUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProfilesByMetroUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
