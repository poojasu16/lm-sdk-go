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

// ILP i l p
//
// swagger:model ILP
type ILP struct {

	// lm name
	// Required: true
	LMName *string `json:"lmName"`

	// wmi name
	// Required: true
	WmiName *string `json:"wmiName"`
}

// Validate validates this i l p
func (m *ILP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLMName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWmiName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ILP) validateLMName(formats strfmt.Registry) error {

	if err := validate.Required("lmName", "body", m.LMName); err != nil {
		return err
	}

	return nil
}

func (m *ILP) validateWmiName(formats strfmt.Registry) error {

	if err := validate.Required("wmiName", "body", m.WmiName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this i l p based on context it is used
func (m *ILP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ILP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ILP) UnmarshalBinary(b []byte) error {
	var res ILP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
