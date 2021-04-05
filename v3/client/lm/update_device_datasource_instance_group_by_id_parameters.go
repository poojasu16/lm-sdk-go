// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewUpdateDeviceDatasourceInstanceGroupByIDParams creates a new UpdateDeviceDatasourceInstanceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceDatasourceInstanceGroupByIDParams() *UpdateDeviceDatasourceInstanceGroupByIDParams {
	return &UpdateDeviceDatasourceInstanceGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithTimeout creates a new UpdateDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	return &UpdateDeviceDatasourceInstanceGroupByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithContext creates a new UpdateDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a context for a request.
func NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	return &UpdateDeviceDatasourceInstanceGroupByIDParams{
		Context: ctx,
	}
}

// NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient creates a new UpdateDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	return &UpdateDeviceDatasourceInstanceGroupByIDParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceDatasourceInstanceGroupByIDParams contains all the parameters to send to the API endpoint
   for the update device datasource instance group by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceDatasourceInstanceGroupByIDParams struct {

	// Body.
	Body *models.DeviceDataSourceInstanceGroup

	/* DeviceDsID.

	   The device-datasource ID you'd like to add an instance group for

	   Format: int32
	*/
	DeviceDsID int32

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithDefaults() *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithBody(body *models.DeviceDataSourceInstanceGroup) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetBody(body *models.DeviceDataSourceInstanceGroup) {
	o.Body = body
}

// WithDeviceDsID adds the deviceDsID to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithDeviceDsID(deviceDsID int32) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithDeviceID(deviceID int32) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithID adds the id to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WithID(id int32) *UpdateDeviceDatasourceInstanceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device datasource instance group by Id params
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceDatasourceInstanceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}