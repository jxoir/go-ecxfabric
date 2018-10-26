// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConnectorCreateRequest connector create request
// swagger:model ConnectorCreateRequest
type ConnectorCreateRequest struct {

	// optional network service
	OptionalNetworkService *ConnectorCreateRequestOptionalNetworkService `json:"OptionalNetworkService,omitempty"`

	// ctag
	Ctag int64 `json:"ctag,omitempty"`

	// is tagged
	IsTagged bool `json:"isTagged,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notification emails
	NotificationEmails []string `json:"notificationEmails"`

	// port Uuid
	PortUUID int64 `json:"portUuid,omitempty"`

	// ri Uuid
	RiUUID string `json:"riUuid,omitempty"`

	// stag
	Stag int64 `json:"stag,omitempty"`
}

// Validate validates this connector create request
func (m *ConnectorCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptionalNetworkService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectorCreateRequest) validateOptionalNetworkService(formats strfmt.Registry) error {

	if swag.IsZero(m.OptionalNetworkService) { // not required
		return nil
	}

	if m.OptionalNetworkService != nil {
		if err := m.OptionalNetworkService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OptionalNetworkService")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorCreateRequest) UnmarshalBinary(b []byte) error {
	var res ConnectorCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConnectorCreateRequestOptionalNetworkService connector create request optional network service
// swagger:model ConnectorCreateRequestOptionalNetworkService
type ConnectorCreateRequestOptionalNetworkService struct {

	// bi forward detection
	BiForwardDetection bool `json:"biForwardDetection,omitempty"`

	// policer
	Policer *Policer `json:"policer,omitempty"`
}

// Validate validates this connector create request optional network service
func (m *ConnectorCreateRequestOptionalNetworkService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectorCreateRequestOptionalNetworkService) validatePolicer(formats strfmt.Registry) error {

	if swag.IsZero(m.Policer) { // not required
		return nil
	}

	if m.Policer != nil {
		if err := m.Policer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OptionalNetworkService" + "." + "policer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorCreateRequestOptionalNetworkService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorCreateRequestOptionalNetworkService) UnmarshalBinary(b []byte) error {
	var res ConnectorCreateRequestOptionalNetworkService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
