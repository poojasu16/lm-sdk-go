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

// NewExecuteDebugCommandParams creates a new ExecuteDebugCommandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteDebugCommandParams() *ExecuteDebugCommandParams {
	return &ExecuteDebugCommandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteDebugCommandParamsWithTimeout creates a new ExecuteDebugCommandParams object
// with the ability to set a timeout on a request.
func NewExecuteDebugCommandParamsWithTimeout(timeout time.Duration) *ExecuteDebugCommandParams {
	return &ExecuteDebugCommandParams{
		timeout: timeout,
	}
}

// NewExecuteDebugCommandParamsWithContext creates a new ExecuteDebugCommandParams object
// with the ability to set a context for a request.
func NewExecuteDebugCommandParamsWithContext(ctx context.Context) *ExecuteDebugCommandParams {
	return &ExecuteDebugCommandParams{
		Context: ctx,
	}
}

// NewExecuteDebugCommandParamsWithHTTPClient creates a new ExecuteDebugCommandParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteDebugCommandParamsWithHTTPClient(client *http.Client) *ExecuteDebugCommandParams {
	return &ExecuteDebugCommandParams{
		HTTPClient: client,
	}
}

/* ExecuteDebugCommandParams contains all the parameters to send to the API endpoint
   for the execute debug command operation.

   Typically these are written to a http.Request.
*/
type ExecuteDebugCommandParams struct {

	// Body.
	Body *models.Debug

	// CollectorID.
	//
	// Format: int32
	// Default: -1
	CollectorID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute debug command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDebugCommandParams) WithDefaults() *ExecuteDebugCommandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute debug command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDebugCommandParams) SetDefaults() {
	var (
		collectorIDDefault = int32(-1)
	)

	val := ExecuteDebugCommandParams{
		CollectorID: &collectorIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the execute debug command params
func (o *ExecuteDebugCommandParams) WithTimeout(timeout time.Duration) *ExecuteDebugCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute debug command params
func (o *ExecuteDebugCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute debug command params
func (o *ExecuteDebugCommandParams) WithContext(ctx context.Context) *ExecuteDebugCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute debug command params
func (o *ExecuteDebugCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute debug command params
func (o *ExecuteDebugCommandParams) WithHTTPClient(client *http.Client) *ExecuteDebugCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute debug command params
func (o *ExecuteDebugCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute debug command params
func (o *ExecuteDebugCommandParams) WithBody(body *models.Debug) *ExecuteDebugCommandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute debug command params
func (o *ExecuteDebugCommandParams) SetBody(body *models.Debug) {
	o.Body = body
}

// WithCollectorID adds the collectorID to the execute debug command params
func (o *ExecuteDebugCommandParams) WithCollectorID(collectorID *int32) *ExecuteDebugCommandParams {
	o.SetCollectorID(collectorID)
	return o
}

// SetCollectorID adds the collectorId to the execute debug command params
func (o *ExecuteDebugCommandParams) SetCollectorID(collectorID *int32) {
	o.CollectorID = collectorID
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteDebugCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CollectorID != nil {

		// query param collectorId
		var qrCollectorID int32

		if o.CollectorID != nil {
			qrCollectorID = *o.CollectorID
		}
		qCollectorID := swag.FormatInt32(qrCollectorID)
		if qCollectorID != "" {

			if err := r.SetQueryParam("collectorId", qCollectorID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
