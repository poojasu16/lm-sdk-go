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

// NewUpdateDevicePropertyByNameParams creates a new UpdateDevicePropertyByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDevicePropertyByNameParams() *UpdateDevicePropertyByNameParams {
	return &UpdateDevicePropertyByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDevicePropertyByNameParamsWithTimeout creates a new UpdateDevicePropertyByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateDevicePropertyByNameParamsWithTimeout(timeout time.Duration) *UpdateDevicePropertyByNameParams {
	return &UpdateDevicePropertyByNameParams{
		timeout: timeout,
	}
}

// NewUpdateDevicePropertyByNameParamsWithContext creates a new UpdateDevicePropertyByNameParams object
// with the ability to set a context for a request.
func NewUpdateDevicePropertyByNameParamsWithContext(ctx context.Context) *UpdateDevicePropertyByNameParams {
	return &UpdateDevicePropertyByNameParams{
		Context: ctx,
	}
}

// NewUpdateDevicePropertyByNameParamsWithHTTPClient creates a new UpdateDevicePropertyByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDevicePropertyByNameParamsWithHTTPClient(client *http.Client) *UpdateDevicePropertyByNameParams {
	return &UpdateDevicePropertyByNameParams{
		HTTPClient: client,
	}
}

/* UpdateDevicePropertyByNameParams contains all the parameters to send to the API endpoint
   for the update device property by name operation.

   Typically these are written to a http.Request.
*/
type UpdateDevicePropertyByNameParams struct {

	// Body.
	Body *models.EntityProperty

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDevicePropertyByNameParams) WithDefaults() *UpdateDevicePropertyByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDevicePropertyByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithTimeout(timeout time.Duration) *UpdateDevicePropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithContext(ctx context.Context) *UpdateDevicePropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithHTTPClient(client *http.Client) *UpdateDevicePropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithBody(body *models.EntityProperty) *UpdateDevicePropertyByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithDeviceID(deviceID int32) *UpdateDevicePropertyByNameParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithName adds the name to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) WithName(name string) *UpdateDevicePropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update device property by name params
func (o *UpdateDevicePropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDevicePropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
