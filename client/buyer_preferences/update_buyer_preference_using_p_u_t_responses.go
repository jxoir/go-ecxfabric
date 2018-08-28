// Code generated by go-swagger; DO NOT EDIT.

package buyer_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jxoir/go-ecxfabric/models"
)

// UpdateBuyerPreferenceUsingPUTReader is a Reader for the UpdateBuyerPreferenceUsingPUT structure.
type UpdateBuyerPreferenceUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuyerPreferenceUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBuyerPreferenceUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateBuyerPreferenceUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateBuyerPreferenceUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBuyerPreferenceUsingPUTOK creates a UpdateBuyerPreferenceUsingPUTOK with default headers values
func NewUpdateBuyerPreferenceUsingPUTOK() *UpdateBuyerPreferenceUsingPUTOK {
	return &UpdateBuyerPreferenceUsingPUTOK{}
}

/*UpdateBuyerPreferenceUsingPUTOK handles this case with default header values.

Success
*/
type UpdateBuyerPreferenceUsingPUTOK struct {
	Payload *models.BuyerPreferenceModel
}

func (o *UpdateBuyerPreferenceUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /ecx/v3/l2/buyerPreference][%d] updateBuyerPreferenceUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateBuyerPreferenceUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuyerPreferenceModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuyerPreferenceUsingPUTNotFound creates a UpdateBuyerPreferenceUsingPUTNotFound with default headers values
func NewUpdateBuyerPreferenceUsingPUTNotFound() *UpdateBuyerPreferenceUsingPUTNotFound {
	return &UpdateBuyerPreferenceUsingPUTNotFound{}
}

/*UpdateBuyerPreferenceUsingPUTNotFound handles this case with default header values.

Record Not Found
*/
type UpdateBuyerPreferenceUsingPUTNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateBuyerPreferenceUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /ecx/v3/l2/buyerPreference][%d] updateBuyerPreferenceUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBuyerPreferenceUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuyerPreferenceUsingPUTInternalServerError creates a UpdateBuyerPreferenceUsingPUTInternalServerError with default headers values
func NewUpdateBuyerPreferenceUsingPUTInternalServerError() *UpdateBuyerPreferenceUsingPUTInternalServerError {
	return &UpdateBuyerPreferenceUsingPUTInternalServerError{}
}

/*UpdateBuyerPreferenceUsingPUTInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateBuyerPreferenceUsingPUTInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *UpdateBuyerPreferenceUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ecx/v3/l2/buyerPreference][%d] updateBuyerPreferenceUsingPUTInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateBuyerPreferenceUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
