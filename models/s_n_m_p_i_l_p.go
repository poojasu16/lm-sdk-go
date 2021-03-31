// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SNMPILP s n m p i l p
//
// swagger:model SNMPILP
type SNMPILP struct {

	// o ID
	// Required: true
	OID *string `json:"OID"`

	// method
	Method string `json:"method,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this s n m p i l p
func (m *SNMPILP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SNMPILP) validateOID(formats strfmt.Registry) error {

	if err := validate.Required("OID", "body", m.OID); err != nil {
		return err
	}

	return nil
}

func (m *SNMPILP) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s n m p i l p based on context it is used
func (m *SNMPILP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SNMPILP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPILP) UnmarshalBinary(b []byte) error {
	var res SNMPILP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
