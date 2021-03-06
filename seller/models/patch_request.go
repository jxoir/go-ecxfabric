// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PatchRequest patch request
// swagger:model PatchRequest
type PatchRequest struct {

	// primary port UUID
	PrimaryPortUUID string `json:"primaryPortUUID,omitempty"`

	// primary vlan c tag
	PrimaryVlanCTag string `json:"primaryVlanCTag,omitempty"`

	// primary vlan s tag
	PrimaryVlanSTag string `json:"primaryVlanSTag,omitempty"`

	// reject comment
	RejectComment string `json:"rejectComment,omitempty"`

	// secondary port UUID
	SecondaryPortUUID string `json:"secondaryPortUUID,omitempty"`

	// secondary vlan c tag
	SecondaryVlanCTag string `json:"secondaryVlanCTag,omitempty"`

	// secondary vlan s tag
	SecondaryVlanSTag string `json:"secondaryVlanSTag,omitempty"`
}

// Validate validates this patch request
func (m *PatchRequest) Validate(formats strfmt.Registry) error {
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
