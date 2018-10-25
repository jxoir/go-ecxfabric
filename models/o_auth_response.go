// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OAuthResponse o auth response
// swagger:model OAuthResponse
type OAuthResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// refresh token timeout
	RefreshTokenTimeout string `json:"refresh_token_timeout,omitempty"`

	// token timeout
	// TODO: Fix go-swagger, add string to int64 timeout
	TokenTimeout int64 `json:"token_timeout,string,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this o auth response
func (m *OAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthResponse) UnmarshalBinary(b []byte) error {
	var res OAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
