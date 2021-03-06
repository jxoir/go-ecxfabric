// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PutServiceProfileModelv3 put service profile modelv3
// swagger:model PutServiceProfileModelv3
type PutServiceProfileModelv3 struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this put service profile modelv3
func (m *PutServiceProfileModelv3) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutServiceProfileModelv3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutServiceProfileModelv3) UnmarshalBinary(b []byte) error {
	var res PutServiceProfileModelv3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
