// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IPBlockGetResponse IP block get response
// swagger:model IPBlockGetResponse
type IPBlockGetResponse struct {

	// content
	Content []*IPBlockGetResponseModel `json:"content"`

	// is first page
	IsFirstPage bool `json:"isFirstPage,omitempty"`

	// is last page
	IsLastPage bool `json:"isLastPage,omitempty"`

	// page number
	PageNumber int64 `json:"pageNumber,omitempty"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this IP block get response
func (m *IPBlockGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPBlockGetResponse) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPBlockGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPBlockGetResponse) UnmarshalBinary(b []byte) error {
	var res IPBlockGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
