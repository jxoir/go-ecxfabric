// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionBundleOrdering subscription bundle ordering
// swagger:model SubscriptionBundleOrdering
type SubscriptionBundleOrdering struct {

	// authorization key
	AuthorizationKey string `json:"authorizationKey,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notification emails
	NotificationEmails []string `json:"notificationEmails"`

	// optional network service
	OptionalNetworkService *OptionalNetworkServiceBundleOrdering `json:"optionalNetworkService,omitempty"`

	// service profile metro code
	ServiceProfileMetroCode string `json:"serviceProfileMetroCode,omitempty"`

	// service profile Uuid
	ServiceProfileUUID string `json:"serviceProfileUuid,omitempty"`

	// subscriber metro code
	SubscriberMetroCode string `json:"subscriberMetroCode,omitempty"`

	// subscriber ri Uuid
	SubscriberRiUUID string `json:"subscriberRiUuid,omitempty"`
}

// Validate validates this subscription bundle ordering
func (m *SubscriptionBundleOrdering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptionalNetworkService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionBundleOrdering) validateOptionalNetworkService(formats strfmt.Registry) error {

	if swag.IsZero(m.OptionalNetworkService) { // not required
		return nil
	}

	if m.OptionalNetworkService != nil {
		if err := m.OptionalNetworkService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optionalNetworkService")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionBundleOrdering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionBundleOrdering) UnmarshalBinary(b []byte) error {
	var res SubscriptionBundleOrdering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}