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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddDeviceDatasourceInstanceGroupParams creates a new AddDeviceDatasourceInstanceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDeviceDatasourceInstanceGroupParams() *AddDeviceDatasourceInstanceGroupParams {
	return &AddDeviceDatasourceInstanceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceDatasourceInstanceGroupParamsWithTimeout creates a new AddDeviceDatasourceInstanceGroupParams object
// with the ability to set a timeout on a request.
func NewAddDeviceDatasourceInstanceGroupParamsWithTimeout(timeout time.Duration) *AddDeviceDatasourceInstanceGroupParams {
	return &AddDeviceDatasourceInstanceGroupParams{
		timeout: timeout,
	}
}

// NewAddDeviceDatasourceInstanceGroupParamsWithContext creates a new AddDeviceDatasourceInstanceGroupParams object
// with the ability to set a context for a request.
func NewAddDeviceDatasourceInstanceGroupParamsWithContext(ctx context.Context) *AddDeviceDatasourceInstanceGroupParams {
	return &AddDeviceDatasourceInstanceGroupParams{
		Context: ctx,
	}
}

// NewAddDeviceDatasourceInstanceGroupParamsWithHTTPClient creates a new AddDeviceDatasourceInstanceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDeviceDatasourceInstanceGroupParamsWithHTTPClient(client *http.Client) *AddDeviceDatasourceInstanceGroupParams {
	return &AddDeviceDatasourceInstanceGroupParams{
		HTTPClient: client,
	}
}

/* AddDeviceDatasourceInstanceGroupParams contains all the parameters to send to the API endpoint
   for the add device datasource instance group operation.

   Typically these are written to a http.Request.
*/
type AddDeviceDatasourceInstanceGroupParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add device datasource instance group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceDatasourceInstanceGroupParams) WithDefaults() *AddDeviceDatasourceInstanceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add device datasource instance group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceDatasourceInstanceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithTimeout(timeout time.Duration) *AddDeviceDatasourceInstanceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithContext(ctx context.Context) *AddDeviceDatasourceInstanceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithHTTPClient(client *http.Client) *AddDeviceDatasourceInstanceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithBody(body *models.DeviceDataSourceInstanceGroup) *AddDeviceDatasourceInstanceGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetBody(body *models.DeviceDataSourceInstanceGroup) {
	o.Body = body
}

// WithDeviceDsID adds the deviceDsID to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithDeviceDsID(deviceDsID int32) *AddDeviceDatasourceInstanceGroupParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) WithDeviceID(deviceID int32) *AddDeviceDatasourceInstanceGroupParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the add device datasource instance group params
func (o *AddDeviceDatasourceInstanceGroupParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceDatasourceInstanceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
