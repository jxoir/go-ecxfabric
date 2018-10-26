// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OptionalNetworkServicesv3 optional network servicesv3
// swagger:model optionalNetworkServicesv3
type OptionalNetworkServicesv3 struct {

	// filter incoming routes
	FilterIncomingRoutes bool `json:"filterIncomingRoutes,omitempty"`
}

// Validate validates this optional network servicesv3
func (m *OptionalNetworkServicesv3) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OptionalNetworkServicesv3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptionalNetworkServicesv3) UnmarshalBinary(b []byte) error {
	var res OptionalNetworkServicesv3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}