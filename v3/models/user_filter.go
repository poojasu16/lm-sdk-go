// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserFilter user filter
//
// swagger:model UserFilter
type UserFilter struct {

	// If the user is API only user. Acceptable values are: all, yes, no
	APIOnlyUser string `json:"apiOnlyUser,omitempty"`

	// The emails of the user
	Email string `json:"email,omitempty"`

	// Whether does the user enabled 2FA. Acceptable values are: all, yes, no
	Enable2fa string `json:"enable2fa,omitempty"`

	// The first name of the user
	FirstName string `json:"firstName,omitempty"`

	// The last name of the user
	LastName string `json:"lastName,omitempty"`

	// Which roles is the user belongs to
	RoleAssignment string `json:"roleAssignment,omitempty"`

	// The user status. Acceptable values are: all, active, suspended
	Status string `json:"status,omitempty"`

	// The username of the user
	Username string `json:"username,omitempty"`
}

// Validate validates this user filter
func (m *UserFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user filter based on context it is used
func (m *UserFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserFilter) UnmarshalBinary(b []byte) error {
	var res UserFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
