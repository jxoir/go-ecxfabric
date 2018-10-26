// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OptionalNetworkServiceBundleOrdering optional network service bundle ordering
// swagger:model optionalNetworkServiceBundleOrdering
type OptionalNetworkServiceBundleOrdering struct {

	// dedicated connector Uuid
	DedicatedConnectorUUID string `json:"dedicatedConnectorUuid,omitempty"`

	// is eqx enforced nat
	IsEqxEnforcedNat bool `json:"isEqxEnforcedNat,omitempty"`

	// perform nat
	PerformNat bool `json:"performNat,omitempty"`

	// policer
	Policer *Policer `json:"policer,omitempty"`
}

// Validate validates this optional network service bundle ordering
func (m *OptionalNetworkServiceBundleOrdering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OptionalNetworkServiceBundleOrdering) validatePolicer(formats strfmt.Registry) error {

	if swag.IsZero(m.Policer) { // not required
		return nil
	}

	if m.Policer != nil {
		if err := m.Policer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OptionalNetworkServiceBundleOrdering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptionalNetworkServiceBundleOrdering) UnmarshalBinary(b []byte) error {
	var res OptionalNetworkServiceBundleOrdering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}