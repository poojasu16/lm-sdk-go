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

// NewUpdateWidgetByIDParams creates a new UpdateWidgetByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWidgetByIDParams() *UpdateWidgetByIDParams {
	return &UpdateWidgetByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWidgetByIDParamsWithTimeout creates a new UpdateWidgetByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateWidgetByIDParamsWithTimeout(timeout time.Duration) *UpdateWidgetByIDParams {
	return &UpdateWidgetByIDParams{
		timeout: timeout,
	}
}

// NewUpdateWidgetByIDParamsWithContext creates a new UpdateWidgetByIDParams object
// with the ability to set a context for a request.
func NewUpdateWidgetByIDParamsWithContext(ctx context.Context) *UpdateWidgetByIDParams {
	return &UpdateWidgetByIDParams{
		Context: ctx,
	}
}

// NewUpdateWidgetByIDParamsWithHTTPClient creates a new UpdateWidgetByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWidgetByIDParamsWithHTTPClient(client *http.Client) *UpdateWidgetByIDParams {
	return &UpdateWidgetByIDParams{
		HTTPClient: client,
	}
}

/* UpdateWidgetByIDParams contains all the parameters to send to the API endpoint
   for the update widget by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateWidgetByIDParams struct {

	// Body.
	Body models.Widget

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update widget by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWidgetByIDParams) WithDefaults() *UpdateWidgetByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update widget by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWidgetByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update widget by Id params
func (o *UpdateWidgetByIDParams) WithTimeout(timeout time.Duration) *UpdateWidgetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update widget by Id params
func (o *UpdateWidgetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update widget by Id params
func (o *UpdateWidgetByIDParams) WithContext(ctx context.Context) *UpdateWidgetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update widget by Id params
func (o *UpdateWidgetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update widget by Id params
func (o *UpdateWidgetByIDParams) WithHTTPClient(client *http.Client) *UpdateWidgetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update widget by Id params
func (o *UpdateWidgetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update widget by Id params
func (o *UpdateWidgetByIDParams) WithBody(body models.Widget) *UpdateWidgetByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update widget by Id params
func (o *UpdateWidgetByIDParams) SetBody(body models.Widget) {
	o.Body = body
}

// WithID adds the id to the update widget by Id params
func (o *UpdateWidgetByIDParams) WithID(id int32) *UpdateWidgetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update widget by Id params
func (o *UpdateWidgetByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWidgetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
