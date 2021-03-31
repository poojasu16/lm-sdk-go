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

// NewDeleteCollectorGroupByIDParams creates a new DeleteCollectorGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCollectorGroupByIDParams() *DeleteCollectorGroupByIDParams {
	return &DeleteCollectorGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCollectorGroupByIDParamsWithTimeout creates a new DeleteCollectorGroupByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteCollectorGroupByIDParamsWithTimeout(timeout time.Duration) *DeleteCollectorGroupByIDParams {
	return &DeleteCollectorGroupByIDParams{
		timeout: timeout,
	}
}

// NewDeleteCollectorGroupByIDParamsWithContext creates a new DeleteCollectorGroupByIDParams object
// with the ability to set a context for a request.
func NewDeleteCollectorGroupByIDParamsWithContext(ctx context.Context) *DeleteCollectorGroupByIDParams {
	return &DeleteCollectorGroupByIDParams{
		Context: ctx,
	}
}

// NewDeleteCollectorGroupByIDParamsWithHTTPClient creates a new DeleteCollectorGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCollectorGroupByIDParamsWithHTTPClient(client *http.Client) *DeleteCollectorGroupByIDParams {
	return &DeleteCollectorGroupByIDParams{
		HTTPClient: client,
	}
}

/* DeleteCollectorGroupByIDParams contains all the parameters to send to the API endpoint
   for the delete collector group by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteCollectorGroupByIDParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete collector group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectorGroupByIDParams) WithDefaults() *DeleteCollectorGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete collector group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectorGroupByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) WithTimeout(timeout time.Duration) *DeleteCollectorGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) WithContext(ctx context.Context) *DeleteCollectorGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) WithHTTPClient(client *http.Client) *DeleteCollectorGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) WithID(id int32) *DeleteCollectorGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete collector group by Id params
func (o *DeleteCollectorGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCollectorGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
