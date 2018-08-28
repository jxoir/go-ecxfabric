// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreatePublicIPRequest create public Ip request
// swagger:model CreatePublicIpRequest
type CreatePublicIPRequest struct {

	// metro code
	MetroCode string `json:"metroCode,omitempty"`
}

// Validate validates this create public Ip request
func (m *CreatePublicIPRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePublicIPRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePublicIPRequest) UnmarshalBinary(b []byte) error {
	var res CreatePublicIPRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
