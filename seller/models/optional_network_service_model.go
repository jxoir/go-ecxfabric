// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OptionalNetworkServiceModel optional network service model
// swagger:model OptionalNetworkServiceModel
type OptionalNetworkServiceModel struct {

	// equinix routes network prefix size
	EquinixRoutesNetworkPrefixSize int32 `json:"equinixRoutesNetworkPrefixSize,omitempty"`

	// filter incoming routes
	FilterIncomingRoutes bool `json:"filterIncomingRoutes,omitempty"`
}

// Validate validates this optional network service model
func (m *OptionalNetworkServiceModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OptionalNetworkServiceModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptionalNetworkServiceModel) UnmarshalBinary(b []byte) error {
	var res OptionalNetworkServiceModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
