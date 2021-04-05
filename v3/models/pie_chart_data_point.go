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

// PieChartDataPoint pie chart data point
//
// swagger:model PieChartDataPoint
type PieChartDataPoint struct {

	// aggregate
	Aggregate bool `json:"aggregate,omitempty"`

	// aggregate function
	AggregateFunction string `json:"aggregateFunction,omitempty"`

	// data point Id
	DataPointID int32 `json:"dataPointId,omitempty"`

	// data point name
	DataPointName string `json:"dataPointName,omitempty"`

	// data source full name
	DataSourceFullName string `json:"dataSourceFullName,omitempty"`

	// data source Id
	DataSourceID int32 `json:"dataSourceId,omitempty"`

	// device display name
	// Required: true
	DeviceDisplayName *string `json:"deviceDisplayName"`

	// device group full path
	// Required: true
	DeviceGroupFullPath *string `json:"deviceGroupFullPath"`

	// instance name
	// Required: true
	InstanceName *string `json:"instanceName"`

	// name
	// Required: true
	Name *string `json:"name"`

	// top10
	Top10 bool `json:"top10,omitempty"`
}

// Validate validates this pie chart data point
func (m *PieChartDataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceGroupFullPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
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

func (m *PieChartDataPoint) validateDeviceDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("deviceDisplayName", "body", m.DeviceDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *PieChartDataPoint) validateDeviceGroupFullPath(formats strfmt.Registry) error {

	if err := validate.Required("deviceGroupFullPath", "body", m.DeviceGroupFullPath); err != nil {
		return err
	}

	return nil
}

func (m *PieChartDataPoint) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *PieChartDataPoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pie chart data point based on context it is used
func (m *PieChartDataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PieChartDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieChartDataPoint) UnmarshalBinary(b []byte) error {
	var res PieChartDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}