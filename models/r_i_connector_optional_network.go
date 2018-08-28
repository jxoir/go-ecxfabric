// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RIConnectorOptionalNetwork r i connector optional network
// swagger:model RIConnectorOptionalNetwork
type RIConnectorOptionalNetwork struct {

	// bgp Ip subnet size
	BgpIPSubnetSize float64 `json:"bgpIpSubnetSize,omitempty"`

	// bi forward detection
	BiForwardDetection bool `json:"biForwardDetection,omitempty"`

	// customer bgp peering Ip
	CustomerBgpPeeringIP string `json:"customerBgpPeeringIp,omitempty"`

	// dedicated subscription name
	DedicatedSubscriptionName string `json:"dedicatedSubscriptionName,omitempty"`

	// dedicated subscription Uuid
	DedicatedSubscriptionUUID string `json:"dedicatedSubscriptionUuid,omitempty"`

	// equinix bgp peering Ip
	EquinixBgpPeeringIP float64 `json:"equinixBgpPeeringIp,omitempty"`

	// is user provided peering Ip
	IsUserProvidedPeeringIP bool `json:"isUserProvidedPeeringIp,omitempty"`

	// policer speed unit
	PolicerSpeedUnit string `json:"policerSpeedUnit,omitempty"`

	// policer speed value
	PolicerSpeedValue float64 `json:"policerSpeedValue,omitempty"`
}

// Validate validates this r i connector optional network
func (m *RIConnectorOptionalNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RIConnectorOptionalNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RIConnectorOptionalNetwork) UnmarshalBinary(b []byte) error {
	var res RIConnectorOptionalNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
