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

// NewAddDevicePropertyParams creates a new AddDevicePropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDevicePropertyParams() *AddDevicePropertyParams {
	return &AddDevicePropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDevicePropertyParamsWithTimeout creates a new AddDevicePropertyParams object
// with the ability to set a timeout on a request.
func NewAddDevicePropertyParamsWithTimeout(timeout time.Duration) *AddDevicePropertyParams {
	return &AddDevicePropertyParams{
		timeout: timeout,
	}
}

// NewAddDevicePropertyParamsWithContext creates a new AddDevicePropertyParams object
// with the ability to set a context for a request.
func NewAddDevicePropertyParamsWithContext(ctx context.Context) *AddDevicePropertyParams {
	return &AddDevicePropertyParams{
		Context: ctx,
	}
}

// NewAddDevicePropertyParamsWithHTTPClient creates a new AddDevicePropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDevicePropertyParamsWithHTTPClient(client *http.Client) *AddDevicePropertyParams {
	return &AddDevicePropertyParams{
		HTTPClient: client,
	}
}

/* AddDevicePropertyParams contains all the parameters to send to the API endpoint
   for the add device property operation.

   Typically these are written to a http.Request.
*/
type AddDevicePropertyParams struct {

	// Body.
	Body *models.EntityProperty

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add device property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDevicePropertyParams) WithDefaults() *AddDevicePropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add device property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDevicePropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add device property params
func (o *AddDevicePropertyParams) WithTimeout(timeout time.Duration) *AddDevicePropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device property params
func (o *AddDevicePropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device property params
func (o *AddDevicePropertyParams) WithContext(ctx context.Context) *AddDevicePropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device property params
func (o *AddDevicePropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device property params
func (o *AddDevicePropertyParams) WithHTTPClient(client *http.Client) *AddDevicePropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device property params
func (o *AddDevicePropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device property params
func (o *AddDevicePropertyParams) WithBody(body *models.EntityProperty) *AddDevicePropertyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device property params
func (o *AddDevicePropertyParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the add device property params
func (o *AddDevicePropertyParams) WithDeviceID(deviceID int32) *AddDevicePropertyParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the add device property params
func (o *AddDevicePropertyParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDevicePropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
