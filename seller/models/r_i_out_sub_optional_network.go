// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RIOutSubOptionalNetwork r i out sub optional network
// swagger:model RIOutSubOptionalNetwork
type RIOutSubOptionalNetwork struct {

	// dedicated connector name
	DedicatedConnectorName string `json:"dedicatedConnectorName,omitempty"`

	// dedicated connector Uuid
	DedicatedConnectorUUID string `json:"dedicatedConnectorUuid,omitempty"`

	// is custom speed
	IsCustomSpeed bool `json:"isCustomSpeed,omitempty"`

	// is eqx enforced nat
	IsEqxEnforcedNat bool `json:"isEqxEnforcedNat,omitempty"`

	// is perform nat
	IsPerformNat bool `json:"isPerformNat,omitempty"`

	// nat eq address
	NatEqAddress string `json:"natEqAddress,omitempty"`

	// policer speed unit
	PolicerSpeedUnit string `json:"policerSpeedUnit,omitempty"`

	// policer speed value
	PolicerSpeedValue int64 `json:"policerSpeedValue,omitempty"`
}

// Validate validates this r i out sub optional network
func (m *RIOutSubOptionalNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RIOutSubOptionalNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RIOutSubOptionalNetwork) UnmarshalBinary(b []byte) error {
	var res RIOutSubOptionalNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}