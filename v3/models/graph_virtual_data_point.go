// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GraphVirtualDataPoint graph virtual data point
//
// swagger:model GraphVirtualDataPoint
type GraphVirtualDataPoint struct {

	// name
	Name string `json:"name,omitempty"`

	// rpn
	Rpn string `json:"rpn,omitempty"`
}

// Validate validates this graph virtual data point
func (m *GraphVirtualDataPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this graph virtual data point based on context it is used
func (m *GraphVirtualDataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphVirtualDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphVirtualDataPoint) UnmarshalBinary(b []byte) error {
	var res GraphVirtualDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
