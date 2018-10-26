// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetSErvProfRespContentMetadata get s erv prof resp content metadata
// swagger:model GetSErvProfRespContentMetadata
type GetSErvProfRespContentMetadata struct {

	// is release vlan
	IsReleaseVlan bool `json:"is_release_vlan,omitempty"`

	// limit auth key conn
	LimitAuthKeyConn bool `json:"limit_auth_key_conn,omitempty"`
}

// Validate validates this get s erv prof resp content metadata
func (m *GetSErvProfRespContentMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSErvProfRespContentMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSErvProfRespContentMetadata) UnmarshalBinary(b []byte) error {
	var res GetSErvProfRespContentMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}