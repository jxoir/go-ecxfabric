// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostConnectionRequest post connection request
// swagger:model PostConnectionRequest
type PostConnectionRequest struct {

	// authorization key
	AuthorizationKey string `json:"authorizationKey,omitempty"`

	// named tag
	NamedTag string `json:"namedTag,omitempty"`

	// notifications
	Notifications []string `json:"notifications"`

	// primary name
	PrimaryName string `json:"primaryName,omitempty"`

	// primary port UUID
	PrimaryPortUUID string `json:"primaryPortUUID,omitempty"`

	// primary vlan c tag
	PrimaryVlanCTag string `json:"primaryVlanCTag,omitempty"`

	// primary vlan s tag
	PrimaryVlanSTag int64 `json:"primaryVlanSTag,omitempty"`

	// primary z side port UUID
	PrimaryZSidePortUUID string `json:"primaryZSidePortUUID,omitempty"`

	// primary z side vlan c tag
	PrimaryZSideVlanCTag int64 `json:"primaryZSideVlanCTag,omitempty"`

	// primary z side vlan s tag
	PrimaryZSideVlanSTag int64 `json:"primaryZSideVlanSTag,omitempty"`

	// profile UUID
	ProfileUUID string `json:"profileUUID,omitempty"`

	// purchase order number
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// secondary name
	SecondaryName string `json:"secondaryName,omitempty"`

	// secondary port UUID
	SecondaryPortUUID string `json:"secondaryPortUUID,omitempty"`

	// secondary vlan c tag
	SecondaryVlanCTag string `json:"secondaryVlanCTag,omitempty"`

	// secondary vlan s tag
	SecondaryVlanSTag int64 `json:"secondaryVlanSTag,omitempty"`

	// secondary z side port UUID
	SecondaryZSidePortUUID string `json:"secondaryZSidePortUUID,omitempty"`

	// secondary z side vlan c tag
	SecondaryZSideVlanCTag int64 `json:"secondaryZSideVlanCTag,omitempty"`

	// secondary z side vlan s tag
	SecondaryZSideVlanSTag int64 `json:"secondaryZSideVlanSTag,omitempty"`

	// seller metro code
	SellerMetroCode string `json:"sellerMetroCode,omitempty"`

	// seller region
	SellerRegion string `json:"sellerRegion,omitempty"`

	// speed
	Speed int64 `json:"speed,omitempty"`

	// speed unit
	SpeedUnit string `json:"speedUnit,omitempty"`
}

// Validate validates this post connection request
func (m *PostConnectionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostConnectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConnectionRequest) UnmarshalBinary(b []byte) error {
	var res PostConnectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
