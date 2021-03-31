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

// NetflowWidget netflow widget
//
// swagger:model NetflowWidget
type NetflowWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// data type
	DataType string `json:"dataType,omitempty"`

	// device display name
	// Required: true
	DeviceDisplayName *string `json:"deviceDisplayName"`

	// device Id
	// Read Only: true
	DeviceID int32 `json:"deviceId,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// netflow filter
	NetflowFilter *NetflowFilters `json:"netflowFilter,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *NetflowWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *NetflowWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *NetflowWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *NetflowWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *NetflowWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *NetflowWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *NetflowWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *NetflowWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *NetflowWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *NetflowWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *NetflowWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *NetflowWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *NetflowWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *NetflowWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *NetflowWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *NetflowWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *NetflowWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *NetflowWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *NetflowWidget) Type() string {
	return "netflow"
}

// SetType sets the type of this subtype
func (m *NetflowWidget) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *NetflowWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *NetflowWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NetflowWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data type
		DataType string `json:"dataType,omitempty"`

		// device display name
		// Required: true
		DeviceDisplayName *string `json:"deviceDisplayName"`

		// device Id
		// Read Only: true
		DeviceID int32 `json:"deviceId,omitempty"`

		// filter
		Filter string `json:"filter,omitempty"`

		// netflow filter
		NetflowFilter *NetflowFilters `json:"netflowFilter,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result NetflowWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.userPermissionField = base.UserPermission

	result.DataType = data.DataType
	result.DeviceDisplayName = data.DeviceDisplayName
	result.DeviceID = data.DeviceID
	result.Filter = data.Filter
	result.NetflowFilter = data.NetflowFilter

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NetflowWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data type
		DataType string `json:"dataType,omitempty"`

		// device display name
		// Required: true
		DeviceDisplayName *string `json:"deviceDisplayName"`

		// device Id
		// Read Only: true
		DeviceID int32 `json:"deviceId,omitempty"`

		// filter
		Filter string `json:"filter,omitempty"`

		// netflow filter
		NetflowFilter *NetflowFilters `json:"netflowFilter,omitempty"`
	}{

		DataType: m.DataType,

		DeviceDisplayName: m.DeviceDisplayName,

		DeviceID: m.DeviceID,

		Filter: m.Filter,

		NetflowFilter: m.NetflowFilter,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this netflow widget
func (m *NetflowWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetflowFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetflowWidget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) validateDeviceDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("deviceDisplayName", "body", m.DeviceDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) validateNetflowFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.NetflowFilter) { // not required
		return nil
	}

	if m.NetflowFilter != nil {
		if err := m.NetflowFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netflowFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this netflow widget based on the context it is used
func (m *NetflowWidget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetflowFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetflowWidget) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedBy", "body", string(m.LastUpdatedBy())); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) contextValidateLastUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedOn", "body", int64(m.LastUpdatedOn())); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) contextValidateDeviceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceId", "body", int32(m.DeviceID)); err != nil {
		return err
	}

	return nil
}

func (m *NetflowWidget) contextValidateNetflowFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.NetflowFilter != nil {
		if err := m.NetflowFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netflowFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetflowWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetflowWidget) UnmarshalBinary(b []byte) error {
	var res NetflowWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
