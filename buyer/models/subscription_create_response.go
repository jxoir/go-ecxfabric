// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SubscriptionCreateResponse subscription create response
// swagger:model SubscriptionCreateResponse
type SubscriptionCreateResponse struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this subscription create response
func (m *SubscriptionCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionCreateResponse) UnmarshalBinary(b []byte) error {
	var res SubscriptionCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
