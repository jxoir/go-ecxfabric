// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ConnectorCreateResponse connector create response
// swagger:model ConnectorCreateResponse
type ConnectorCreateResponse struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this connector create response
func (m *ConnectorCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorCreateResponse) UnmarshalBinary(b []byte) error {
	var res ConnectorCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}