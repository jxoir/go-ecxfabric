// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetOptionalNetworkService get optional network service
// swagger:model GetOptionalNetworkService
type GetOptionalNetworkService struct {

	// bi forward detection
	BiForwardDetection bool `json:"biForwardDetection,omitempty"`

	// customer provided peering Ip
	CustomerProvidedPeeringIP bool `json:"customerProvidedPeeringIp,omitempty"`

	// dedicated subscription Uuid
	DedicatedSubscriptionUUID string `json:"dedicatedSubscriptionUuid,omitempty"`

	// policer
	Policer *Policer `json:"policer,omitempty"`
}

// Validate validates this get optional network service
func (m *GetOptionalNetworkService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOptionalNetworkService) validatePolicer(formats strfmt.Registry) error {

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
func (m *GetOptionalNetworkService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOptionalNetworkService) UnmarshalBinary(b []byte) error {
	var res GetOptionalNetworkService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}