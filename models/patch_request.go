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

// PatchRequest patch request
// swagger:model PatchRequest
type PatchRequest struct {

	// action details list
	ActionDetailsList []*PatchActionDetailsList `json:"ActionDetailsList"`
}

// Validate validates this patch request
func (m *PatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionDetailsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchRequest) validateActionDetailsList(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionDetailsList) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionDetailsList); i++ {
		if swag.IsZero(m.ActionDetailsList[i]) { // not required
			continue
		}

		if m.ActionDetailsList[i] != nil {
			if err := m.ActionDetailsList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ActionDetailsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchRequest) UnmarshalBinary(b []byte) error {
	var res PatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
