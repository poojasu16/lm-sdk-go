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
)

// NewGetAdminByIDParams creates a new GetAdminByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdminByIDParams() *GetAdminByIDParams {
	return &GetAdminByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdminByIDParamsWithTimeout creates a new GetAdminByIDParams object
// with the ability to set a timeout on a request.
func NewGetAdminByIDParamsWithTimeout(timeout time.Duration) *GetAdminByIDParams {
	return &GetAdminByIDParams{
		timeout: timeout,
	}
}

// NewGetAdminByIDParamsWithContext creates a new GetAdminByIDParams object
// with the ability to set a context for a request.
func NewGetAdminByIDParamsWithContext(ctx context.Context) *GetAdminByIDParams {
	return &GetAdminByIDParams{
		Context: ctx,
	}
}

// NewGetAdminByIDParamsWithHTTPClient creates a new GetAdminByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdminByIDParamsWithHTTPClient(client *http.Client) *GetAdminByIDParams {
	return &GetAdminByIDParams{
		HTTPClient: client,
	}
}

/* GetAdminByIDParams contains all the parameters to send to the API endpoint
   for the get admin by Id operation.

   Typically these are written to a http.Request.
*/
type GetAdminByIDParams struct {

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get admin by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdminByIDParams) WithDefaults() *GetAdminByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get admin by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdminByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get admin by Id params
func (o *GetAdminByIDParams) WithTimeout(timeout time.Duration) *GetAdminByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get admin by Id params
func (o *GetAdminByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get admin by Id params
func (o *GetAdminByIDParams) WithContext(ctx context.Context) *GetAdminByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get admin by Id params
func (o *GetAdminByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get admin by Id params
func (o *GetAdminByIDParams) WithHTTPClient(client *http.Client) *GetAdminByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get admin by Id params
func (o *GetAdminByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get admin by Id params
func (o *GetAdminByIDParams) WithFields(fields *string) *GetAdminByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get admin by Id params
func (o *GetAdminByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get admin by Id params
func (o *GetAdminByIDParams) WithID(id int32) *GetAdminByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get admin by Id params
func (o *GetAdminByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdminByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
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
