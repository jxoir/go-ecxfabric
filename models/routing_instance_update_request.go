// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RoutingInstanceUpdateRequest routing instance update request
// swagger:model RoutingInstanceUpdateRequest
type RoutingInstanceUpdateRequest struct {

	// name
	Name string `json:"name,omitempty"`

	// notification emails
	NotificationEmails []string `json:"notificationEmails"`
}

// Validate validates this routing instance update request
func (m *RoutingInstanceUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoutingInstanceUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingInstanceUpdateRequest) UnmarshalBinary(b []byte) error {
	var res RoutingInstanceUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
