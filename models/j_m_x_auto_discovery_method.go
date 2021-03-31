// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JMXAutoDiscoveryMethod j m x auto discovery method
//
// swagger:model JMXAutoDiscoveryMethod
type JMXAutoDiscoveryMethod struct {

	// path
	// Required: true
	Path *string `json:"path"`

	// ports
	// Required: true
	Ports *string `json:"ports"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Name gets the name of this subtype
func (m *JMXAutoDiscoveryMethod) Name() string {
	return "ad_jmx"
}

// SetName sets the name of this subtype
func (m *JMXAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *JMXAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// path
		// Required: true
		Path *string `json:"path"`

		// ports
		// Required: true
		Ports *string `json:"ports"`

		// url
		// Required: true
		URL *string `json:"url"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result JMXAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Path = data.Path
	result.Ports = data.Ports
	result.URL = data.URL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m JMXAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// path
		// Required: true
		Path *string `json:"path"`

		// ports
		// Required: true
		Ports *string `json:"ports"`

		// url
		// Required: true
		URL *string `json:"url"`
	}{

		Path: m.Path,

		Ports: m.Ports,

		URL: m.URL,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this j m x auto discovery method
func (m *JMXAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JMXAutoDiscoveryMethod) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *JMXAutoDiscoveryMethod) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

func (m *JMXAutoDiscoveryMethod) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this j m x auto discovery method based on the context it is used
func (m *JMXAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JMXAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JMXAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res JMXAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
