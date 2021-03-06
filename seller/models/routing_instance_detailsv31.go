// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RoutingInstanceDetailsv31 routing instance detailsv31
// swagger:model routingInstanceDetailsv31
type RoutingInstanceDetailsv31 struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this routing instance detailsv31
func (m *RoutingInstanceDetailsv31) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoutingInstanceDetailsv31) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingInstanceDetailsv31) UnmarshalBinary(b []byte) error {
	var res RoutingInstanceDetailsv31
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
