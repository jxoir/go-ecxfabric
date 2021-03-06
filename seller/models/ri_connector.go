// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RiConnector ri connector
// swagger:model RiConnector
type RiConnector struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created date
	CreatedDate string `json:"createdDate,omitempty"`

	// ctag
	Ctag int32 `json:"ctag,omitempty"`

	// device
	Device string `json:"device,omitempty"`

	// last updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// last updated date
	LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// optional network services
	OptionalNetworkServices *RIConnectorOptionalNetwork `json:"optionalNetworkServices,omitempty"`

	// port name
	PortName string `json:"portName,omitempty"`

	// stag
	Stag int32 `json:"stag,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this ri connector
func (m *RiConnector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptionalNetworkServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RiConnector) validateOptionalNetworkServices(formats strfmt.Registry) error {

	if swag.IsZero(m.OptionalNetworkServices) { // not required
		return nil
	}

	if m.OptionalNetworkServices != nil {
		if err := m.OptionalNetworkServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optionalNetworkServices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RiConnector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RiConnector) UnmarshalBinary(b []byte) error {
	var res RiConnector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
