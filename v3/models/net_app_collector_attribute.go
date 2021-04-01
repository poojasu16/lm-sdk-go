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

// NetAppCollectorAttribute net app collector attribute
//
// swagger:model NetAppCollectorAttribute
type NetAppCollectorAttribute struct {

	// ip
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// net app API
	NetAppAPI string `json:"netAppAPI,omitempty"`

	// net app aggregate
	NetAppAggregate string `json:"netAppAggregate,omitempty"`

	// net app index
	NetAppIndex string `json:"netAppIndex,omitempty"`

	// net app instance
	NetAppInstance string `json:"netAppInstance,omitempty"`

	// net app object
	NetAppObject string `json:"netAppObject,omitempty"`

	// net app type
	NetAppType string `json:"netAppType,omitempty"`

	// net app XML
	NetAppXML string `json:"netAppXML,omitempty"`

	// net app XML bulk
	NetAppXMLBulk string `json:"netAppXMLBulk,omitempty"`

	// net app XML bulk locator
	NetAppXMLBulkLocator string `json:"netAppXMLBulkLocator,omitempty"`

	// net app XML index
	NetAppXMLIndex string `json:"netAppXMLIndex,omitempty"`

	// net app XML instance
	NetAppXMLInstance string `json:"netAppXMLInstance,omitempty"`

	// net app XML locator
	NetAppXMLLocator string `json:"netAppXMLLocator,omitempty"`

	// params
	Params interface{} `json:"params,omitempty"`
}

// Name gets the name of this subtype
func (m *NetAppCollectorAttribute) Name() string {
	return "netapp"
}

// SetName sets the name of this subtype
func (m *NetAppCollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NetAppCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// net app API
		NetAppAPI string `json:"netAppAPI,omitempty"`

		// net app aggregate
		NetAppAggregate string `json:"netAppAggregate,omitempty"`

		// net app index
		NetAppIndex string `json:"netAppIndex,omitempty"`

		// net app instance
		NetAppInstance string `json:"netAppInstance,omitempty"`

		// net app object
		NetAppObject string `json:"netAppObject,omitempty"`

		// net app type
		NetAppType string `json:"netAppType,omitempty"`

		// net app XML
		NetAppXML string `json:"netAppXML,omitempty"`

		// net app XML bulk
		NetAppXMLBulk string `json:"netAppXMLBulk,omitempty"`

		// net app XML bulk locator
		NetAppXMLBulkLocator string `json:"netAppXMLBulkLocator,omitempty"`

		// net app XML index
		NetAppXMLIndex string `json:"netAppXMLIndex,omitempty"`

		// net app XML instance
		NetAppXMLInstance string `json:"netAppXMLInstance,omitempty"`

		// net app XML locator
		NetAppXMLLocator string `json:"netAppXMLLocator,omitempty"`

		// params
		Params interface{} `json:"params,omitempty"`
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

	var result NetAppCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.IP = data.IP
	result.NetAppAPI = data.NetAppAPI
	result.NetAppAggregate = data.NetAppAggregate
	result.NetAppIndex = data.NetAppIndex
	result.NetAppInstance = data.NetAppInstance
	result.NetAppObject = data.NetAppObject
	result.NetAppType = data.NetAppType
	result.NetAppXML = data.NetAppXML
	result.NetAppXMLBulk = data.NetAppXMLBulk
	result.NetAppXMLBulkLocator = data.NetAppXMLBulkLocator
	result.NetAppXMLIndex = data.NetAppXMLIndex
	result.NetAppXMLInstance = data.NetAppXMLInstance
	result.NetAppXMLLocator = data.NetAppXMLLocator
	result.Params = data.Params

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NetAppCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// net app API
		NetAppAPI string `json:"netAppAPI,omitempty"`

		// net app aggregate
		NetAppAggregate string `json:"netAppAggregate,omitempty"`

		// net app index
		NetAppIndex string `json:"netAppIndex,omitempty"`

		// net app instance
		NetAppInstance string `json:"netAppInstance,omitempty"`

		// net app object
		NetAppObject string `json:"netAppObject,omitempty"`

		// net app type
		NetAppType string `json:"netAppType,omitempty"`

		// net app XML
		NetAppXML string `json:"netAppXML,omitempty"`

		// net app XML bulk
		NetAppXMLBulk string `json:"netAppXMLBulk,omitempty"`

		// net app XML bulk locator
		NetAppXMLBulkLocator string `json:"netAppXMLBulkLocator,omitempty"`

		// net app XML index
		NetAppXMLIndex string `json:"netAppXMLIndex,omitempty"`

		// net app XML instance
		NetAppXMLInstance string `json:"netAppXMLInstance,omitempty"`

		// net app XML locator
		NetAppXMLLocator string `json:"netAppXMLLocator,omitempty"`

		// params
		Params interface{} `json:"params,omitempty"`
	}{

		IP: m.IP,

		NetAppAPI: m.NetAppAPI,

		NetAppAggregate: m.NetAppAggregate,

		NetAppIndex: m.NetAppIndex,

		NetAppInstance: m.NetAppInstance,

		NetAppObject: m.NetAppObject,

		NetAppType: m.NetAppType,

		NetAppXML: m.NetAppXML,

		NetAppXMLBulk: m.NetAppXMLBulk,

		NetAppXMLBulkLocator: m.NetAppXMLBulkLocator,

		NetAppXMLIndex: m.NetAppXMLIndex,

		NetAppXMLInstance: m.NetAppXMLInstance,

		NetAppXMLLocator: m.NetAppXMLLocator,

		Params: m.Params,
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

// Validate validates this net app collector attribute
func (m *NetAppCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this net app collector attribute based on the context it is used
func (m *NetAppCollectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetAppCollectorAttribute) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetAppCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetAppCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res NetAppCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
