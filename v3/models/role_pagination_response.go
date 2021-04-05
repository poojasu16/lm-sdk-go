// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RolePaginationResponse role pagination response
//
// swagger:model RolePaginationResponse
type RolePaginationResponse struct {

	// items
	Items []*Role `json:"items,omitempty"`

	// search Id
	// Read Only: true
	SearchID string `json:"searchId,omitempty"`

	// total
	// Read Only: true
	Total int32 `json:"total,omitempty"`
}

// Validate validates this role pagination response
func (m *RolePaginationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RolePaginationResponse) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this role pagination response based on the context it is used
func (m *RolePaginationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RolePaginationResponse) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RolePaginationResponse) contextValidateSearchID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "searchId", "body", string(m.SearchID)); err != nil {
		return err
	}

	return nil
}

func (m *RolePaginationResponse) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", int32(m.Total)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RolePaginationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RolePaginationResponse) UnmarshalBinary(b []byte) error {
	var res RolePaginationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
