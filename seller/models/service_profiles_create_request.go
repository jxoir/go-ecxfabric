// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceProfilesCreateRequest service profiles create request
// swagger:model ServiceProfilesCreateRequest
type ServiceProfilesCreateRequest struct {

	// additional info
	AdditionalInfo *AdditionalInfoModel `json:"additionalInfo,omitempty"`

	// allow custom speed
	AllowCustomSpeed bool `json:"allowCustomSpeed,omitempty"`

	// auth key label
	AuthKeyLabel string `json:"authKeyLabel,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// features
	Features *EnabledFeaturesModel `json:"features,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// on profile approval reject notification
	OnProfileApprovalRejectNotification []string `json:"onProfileApprovalRejectNotification"`

	// optional network services
	OptionalNetworkServices *OptionalNetworkServicesv3 `json:"optionalNetworkServices,omitempty"`

	// routing instance details
	RoutingInstanceDetails *RoutingInstanceDetailsv3 `json:"routingInstanceDetails,omitempty"`

	// service type
	ServiceType string `json:"serviceType,omitempty"`

	// speed bands
	SpeedBands []*SpeedBand `json:"speedBands"`

	// status for display
	StatusForDisplay string `json:"statusForDisplay,omitempty"`
}

// Validate validates this service profiles create request
func (m *ServiceProfilesCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionalNetworkServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingInstanceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeedBands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceProfilesCreateRequest) validateAdditionalInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalInfo) { // not required
		return nil
	}

	if m.AdditionalInfo != nil {
		if err := m.AdditionalInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additionalInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceProfilesCreateRequest) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceProfilesCreateRequest) validateOptionalNetworkServices(formats strfmt.Registry) error {

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

func (m *ServiceProfilesCreateRequest) validateRoutingInstanceDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingInstanceDetails) { // not required
		return nil
	}

	if m.RoutingInstanceDetails != nil {
		if err := m.RoutingInstanceDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routingInstanceDetails")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceProfilesCreateRequest) validateSpeedBands(formats strfmt.Registry) error {

	if swag.IsZero(m.SpeedBands) { // not required
		return nil
	}

	for i := 0; i < len(m.SpeedBands); i++ {
		if swag.IsZero(m.SpeedBands[i]) { // not required
			continue
		}

		if m.SpeedBands[i] != nil {
			if err := m.SpeedBands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("speedBands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceProfilesCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceProfilesCreateRequest) UnmarshalBinary(b []byte) error {
	var res ServiceProfilesCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
