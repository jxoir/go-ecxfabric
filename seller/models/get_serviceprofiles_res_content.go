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

// GetServiceprofilesResContent get serviceprofiles res content
// swagger:model GetServiceprofilesResContent
type GetServiceprofilesResContent struct {

	// alert percentage
	AlertPercentage float64 `json:"alertPercentage,omitempty"`

	// allow custom speed
	AllowCustomSpeed bool `json:"allowCustomSpeed,omitempty"`

	// api available
	APIAvailable bool `json:"apiAvailable,omitempty"`

	// auth key label
	AuthKeyLabel string `json:"authKeyLabel,omitempty"`

	// authorization key
	AuthorizationKey string `json:"authorizationKey,omitempty"`

	// connection accessibility
	ConnectionAccessibility string `json:"connectionAccessibility,omitempty"`

	// connection name label
	ConnectionNameLabel string `json:"connectionNameLabel,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created by email
	CreatedByEmail string `json:"createdByEmail,omitempty"`

	// created by full name
	CreatedByFullName string `json:"createdByFullName,omitempty"`

	// created date
	CreatedDate string `json:"createdDate,omitempty"`

	// ctag label
	CtagLabel string `json:"ctagLabel,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// equinix managed port and vlan
	EquinixManagedPortAndVlan bool `json:"equinixManagedPortAndVlan,omitempty"`

	// features
	Features *GetservProfResContentFeatures `json:"features,omitempty"`

	// global organization
	GlobalOrganization string `json:"globalOrganization,omitempty"`

	// integration Id
	IntegrationID string `json:"integrationId,omitempty"`

	// last updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// last updated by full name
	LastUpdatedByFullName string `json:"lastUpdatedByFullName,omitempty"`

	// last updated date
	LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`

	// metadata
	Metadata *GetSErvProfRespContentMetadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// named tags
	NamedTags GetServiceprofilesResContentNamedTags `json:"namedTags"`

	// on bandwidth threshold notification
	OnBandwidthThresholdNotification []string `json:"onBandwidthThresholdNotification"`

	// on profile approval reject notification
	OnProfileApprovalRejectNotification []string `json:"onProfileApprovalRejectNotification"`

	// on vc approval rejection notification
	OnVcApprovalRejectionNotification []string `json:"onVcApprovalRejectionNotification"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// ports
	Ports []*PortDetail `json:"ports"`

	// private
	Private bool `json:"private,omitempty"`

	// required redundancy
	RequiredRedundancy bool `json:"requiredRedundancy,omitempty"`

	// speed bands
	SpeedBands []*SpeedBand `json:"speedBands"`

	// speed from API
	SpeedFromAPI bool `json:"speedFromAPI,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// tag type
	TagType string `json:"tagType,omitempty"`

	// updated by email
	UpdatedByEmail string `json:"updatedByEmail,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// vlan same as primary
	VlanSameAsPrimary bool `json:"vlanSameAsPrimary,omitempty"`
}

// Validate validates this get serviceprofiles res content
func (m *GetServiceprofilesResContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamedTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
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

func (m *GetServiceprofilesResContent) validateFeatures(formats strfmt.Registry) error {

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

func (m *GetServiceprofilesResContent) validateMetadata(formats strfmt.Registry) error {

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

func (m *GetServiceprofilesResContent) validateNamedTags(formats strfmt.Registry) error {

	if swag.IsZero(m.NamedTags) { // not required
		return nil
	}

	if err := m.NamedTags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("namedTags")
		}
		return err
	}

	return nil
}

func (m *GetServiceprofilesResContent) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetServiceprofilesResContent) validateSpeedBands(formats strfmt.Registry) error {

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
func (m *GetServiceprofilesResContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetServiceprofilesResContent) UnmarshalBinary(b []byte) error {
	var res GetServiceprofilesResContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
