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

// AwsEC2ReservedInstanceDiscoveryMethod aws e c2 reserved instance discovery method
//
// swagger:model AwsEC2ReservedInstanceDiscoveryMethod
type AwsEC2ReservedInstanceDiscoveryMethod struct {

	// instance name
	// Required: true
	InstanceName *string `json:"instanceName"`
}

// Name gets the name of this subtype
func (m *AwsEC2ReservedInstanceDiscoveryMethod) Name() string {
	return "ad_awsec2reservedinstance"
}

// SetName sets the name of this subtype
func (m *AwsEC2ReservedInstanceDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsEC2ReservedInstanceDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// instance name
		// Required: true
		InstanceName *string `json:"instanceName"`
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

	var result AwsEC2ReservedInstanceDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.InstanceName = data.InstanceName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsEC2ReservedInstanceDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// instance name
		// Required: true
		InstanceName *string `json:"instanceName"`
	}{

		InstanceName: m.InstanceName,
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

// Validate validates this aws e c2 reserved instance discovery method
func (m *AwsEC2ReservedInstanceDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsEC2ReservedInstanceDiscoveryMethod) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws e c2 reserved instance discovery method based on the context it is used
func (m *AwsEC2ReservedInstanceDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsEC2ReservedInstanceDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEC2ReservedInstanceDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AwsEC2ReservedInstanceDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
