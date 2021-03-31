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
)

// ScriptCollectorAttribute script collector attribute
//
// swagger:model ScriptCollectorAttribute
type ScriptCollectorAttribute struct {

	// groovy script
	GroovyScript string `json:"groovyScript,omitempty"`

	// linux cmdline
	LinuxCmdline string `json:"linuxCmdline,omitempty"`

	// linux script
	LinuxScript string `json:"linuxScript,omitempty"`

	// script type
	ScriptType string `json:"scriptType,omitempty"`

	// windows cmdline
	WindowsCmdline string `json:"windowsCmdline,omitempty"`

	// windows script
	WindowsScript string `json:"windowsScript,omitempty"`
}

// Name gets the name of this subtype
func (m *ScriptCollectorAttribute) Name() string {
	return "script"
}

// SetName sets the name of this subtype
func (m *ScriptCollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ScriptCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// groovy script
		GroovyScript string `json:"groovyScript,omitempty"`

		// linux cmdline
		LinuxCmdline string `json:"linuxCmdline,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// script type
		ScriptType string `json:"scriptType,omitempty"`

		// windows cmdline
		WindowsCmdline string `json:"windowsCmdline,omitempty"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`
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

	var result ScriptCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.GroovyScript = data.GroovyScript
	result.LinuxCmdline = data.LinuxCmdline
	result.LinuxScript = data.LinuxScript
	result.ScriptType = data.ScriptType
	result.WindowsCmdline = data.WindowsCmdline
	result.WindowsScript = data.WindowsScript

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ScriptCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// groovy script
		GroovyScript string `json:"groovyScript,omitempty"`

		// linux cmdline
		LinuxCmdline string `json:"linuxCmdline,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// script type
		ScriptType string `json:"scriptType,omitempty"`

		// windows cmdline
		WindowsCmdline string `json:"windowsCmdline,omitempty"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`
	}{

		GroovyScript: m.GroovyScript,

		LinuxCmdline: m.LinuxCmdline,

		LinuxScript: m.LinuxScript,

		ScriptType: m.ScriptType,

		WindowsCmdline: m.WindowsCmdline,

		WindowsScript: m.WindowsScript,
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

// Validate validates this script collector attribute
func (m *ScriptCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this script collector attribute based on the context it is used
func (m *ScriptCollectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ScriptCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res ScriptCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
