// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GETConnectionByUUIDResponse g e t connection by u Uid response
// swagger:model GETConnectionByUUidResponse
type GETConnectionByUUIDResponse struct {

	// aside encapsulation
	AsideEncapsulation string `json:"asideEncapsulation,omitempty"`

	// authorization key
	AuthorizationKey string `json:"authorizationKey,omitempty"`

	// billing tier
	BillingTier string `json:"billingTier,omitempty"`

	// buyer organization name
	BuyerOrganizationName string `json:"buyerOrganizationName,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created by email
	CreatedByEmail string `json:"createdByEmail,omitempty"`

	// created by full name
	CreatedByFullName string `json:"createdByFullName,omitempty"`

	// created date
	CreatedDate string `json:"createdDate,omitempty"`

	// last updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// last updated by email
	LastUpdatedByEmail string `json:"lastUpdatedByEmail,omitempty"`

	// last updated by full name
	LastUpdatedByFullName string `json:"lastUpdatedByFullName,omitempty"`

	// last updated date
	LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`

	// metadata
	Metadata *MetadataV3 `json:"metadata,omitempty"`

	// metro code
	MetroCode string `json:"metroCode,omitempty"`

	// metro description
	MetroDescription string `json:"metroDescription,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// named tag
	NamedTag string `json:"namedTag,omitempty"`

	// notifications
	Notifications []string `json:"notifications"`

	// port name
	PortName string `json:"portName,omitempty"`

	// port UUID
	PortUUID string `json:"portUUID,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// purchase order number
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// redundancy group
	RedundancyGroup string `json:"redundancyGroup,omitempty"`

	// redundancy type
	RedundancyType string `json:"redundancyType,omitempty"`

	// redundant UUID
	RedundantUUID string `json:"redundantUUID,omitempty"`

	// remote
	Remote bool `json:"remote,omitempty"`

	// self
	Self bool `json:"self,omitempty"`

	// seller metro code
	SellerMetroCode string `json:"sellerMetroCode,omitempty"`

	// seller metro description
	SellerMetroDescription string `json:"sellerMetroDescription,omitempty"`

	// seller organization name
	SellerOrganizationName string `json:"sellerOrganizationName,omitempty"`

	// seller service name
	SellerServiceName string `json:"sellerServiceName,omitempty"`

	// seller service UUID
	SellerServiceUUID string `json:"sellerServiceUUID,omitempty"`

	// speed
	Speed int32 `json:"speed,omitempty"`

	// speed unit
	SpeedUnit string `json:"speedUnit,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// vlan s tag
	VlanSTag int32 `json:"vlanSTag,omitempty"`

	// z side port name
	ZSidePortName string `json:"zSidePortName,omitempty"`

	// z side port UUID
	ZSidePortUUID string `json:"zSidePortUUID,omitempty"`

	// z side vlan c tag
	ZSideVlanCTag int32 `json:"zSideVlanCTag,omitempty"`

	// z side vlan s tag
	ZSideVlanSTag int32 `json:"zSideVlanSTag,omitempty"`
}

// Validate validates this g e t connection by u Uid response
func (m *GETConnectionByUUIDResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GETConnectionByUUIDResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GETConnectionByUUIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETConnectionByUUIDResponse) UnmarshalBinary(b []byte) error {
	var res GETConnectionByUUIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}