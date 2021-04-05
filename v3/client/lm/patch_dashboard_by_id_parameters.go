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

// NewPatchDashboardByIDParams creates a new PatchDashboardByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDashboardByIDParams() *PatchDashboardByIDParams {
	return &PatchDashboardByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDashboardByIDParamsWithTimeout creates a new PatchDashboardByIDParams object
// with the ability to set a timeout on a request.
func NewPatchDashboardByIDParamsWithTimeout(timeout time.Duration) *PatchDashboardByIDParams {
	return &PatchDashboardByIDParams{
		timeout: timeout,
	}
}

// NewPatchDashboardByIDParamsWithContext creates a new PatchDashboardByIDParams object
// with the ability to set a context for a request.
func NewPatchDashboardByIDParamsWithContext(ctx context.Context) *PatchDashboardByIDParams {
	return &PatchDashboardByIDParams{
		Context: ctx,
	}
}

// NewPatchDashboardByIDParamsWithHTTPClient creates a new PatchDashboardByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDashboardByIDParamsWithHTTPClient(client *http.Client) *PatchDashboardByIDParams {
	return &PatchDashboardByIDParams{
		HTTPClient: client,
	}
}

/* PatchDashboardByIDParams contains all the parameters to send to the API endpoint
   for the patch dashboard by Id operation.

   Typically these are written to a http.Request.
*/
type PatchDashboardByIDParams struct {

	// Body.
	Body *models.Dashboard

	// ID.
	//
	// Format: int32
	ID int32

	// OverwriteGroupFields.
	OverwriteGroupFields *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDashboardByIDParams) WithDefaults() *PatchDashboardByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDashboardByIDParams) SetDefaults() {
	var (
		overwriteGroupFieldsDefault = bool(false)
	)

	val := PatchDashboardByIDParams{
		OverwriteGroupFields: &overwriteGroupFieldsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithTimeout(timeout time.Duration) *PatchDashboardByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithContext(ctx context.Context) *PatchDashboardByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithHTTPClient(client *http.Client) *PatchDashboardByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithBody(body *models.Dashboard) *PatchDashboardByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetBody(body *models.Dashboard) {
	o.Body = body
}

// WithID adds the id to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithID(id int32) *PatchDashboardByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOverwriteGroupFields adds the overwriteGroupFields to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) WithOverwriteGroupFields(overwriteGroupFields *bool) *PatchDashboardByIDParams {
	o.SetOverwriteGroupFields(overwriteGroupFields)
	return o
}

// SetOverwriteGroupFields adds the overwriteGroupFields to the patch dashboard by Id params
func (o *PatchDashboardByIDParams) SetOverwriteGroupFields(overwriteGroupFields *bool) {
	o.OverwriteGroupFields = overwriteGroupFields
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDashboardByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.OverwriteGroupFields != nil {

		// query param overwriteGroupFields
		var qrOverwriteGroupFields bool

		if o.OverwriteGroupFields != nil {
			qrOverwriteGroupFields = *o.OverwriteGroupFields
		}
		qOverwriteGroupFields := swag.FormatBool(qrOverwriteGroupFields)
		if qOverwriteGroupFields != "" {

			if err := r.SetQueryParam("overwriteGroupFields", qOverwriteGroupFields); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}