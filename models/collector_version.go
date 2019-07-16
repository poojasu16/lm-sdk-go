// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CollectorVersion collector version
// swagger:model CollectorVersion
type CollectorVersion struct {

	// True if Linux collector available
	// Read Only: true
	Has32bitLinux *bool `json:"has32bitLinux,omitempty"`

	// True if Windows collector available
	// Read Only: true
	Has32bitWindows *bool `json:"has32bitWindows,omitempty"`

	// The collector major version
	// Read Only: true
	MajorVersion int32 `json:"majorVersion,omitempty"`

	// True if collector is a required release
	// Read Only: true
	Mandatory *bool `json:"mandatory,omitempty"`

	// The collector minor version
	// Read Only: true
	MinorVersion int32 `json:"minorVersion,omitempty"`

	// Release Epoch for official releases
	// Read Only: true
	ReleaseEpoch int64 `json:"releaseEpoch,omitempty"`

	// False for early release. True for general release
	// Read Only: true
	Stable *bool `json:"stable,omitempty"`
}

// Validate validates this collector version
func (m *CollectorVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectorVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorVersion) UnmarshalBinary(b []byte) error {
	var res CollectorVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}