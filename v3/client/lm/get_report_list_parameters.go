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

// NewGetReportListParams creates a new GetReportListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportListParams() *GetReportListParams {
	return &GetReportListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportListParamsWithTimeout creates a new GetReportListParams object
// with the ability to set a timeout on a request.
func NewGetReportListParamsWithTimeout(timeout time.Duration) *GetReportListParams {
	return &GetReportListParams{
		timeout: timeout,
	}
}

// NewGetReportListParamsWithContext creates a new GetReportListParams object
// with the ability to set a context for a request.
func NewGetReportListParamsWithContext(ctx context.Context) *GetReportListParams {
	return &GetReportListParams{
		Context: ctx,
	}
}

// NewGetReportListParamsWithHTTPClient creates a new GetReportListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportListParamsWithHTTPClient(client *http.Client) *GetReportListParams {
	return &GetReportListParams{
		HTTPClient: client,
	}
}

/* GetReportListParams contains all the parameters to send to the API endpoint
   for the get report list operation.

   Typically these are written to a http.Request.
*/
type GetReportListParams struct {

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get report list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportListParams) WithDefaults() *GetReportListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get report list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportListParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetReportListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get report list params
func (o *GetReportListParams) WithTimeout(timeout time.Duration) *GetReportListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get report list params
func (o *GetReportListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get report list params
func (o *GetReportListParams) WithContext(ctx context.Context) *GetReportListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get report list params
func (o *GetReportListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get report list params
func (o *GetReportListParams) WithHTTPClient(client *http.Client) *GetReportListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get report list params
func (o *GetReportListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get report list params
func (o *GetReportListParams) WithFields(fields *string) *GetReportListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get report list params
func (o *GetReportListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get report list params
func (o *GetReportListParams) WithFilter(filter *string) *GetReportListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get report list params
func (o *GetReportListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get report list params
func (o *GetReportListParams) WithOffset(offset *int32) *GetReportListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get report list params
func (o *GetReportListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get report list params
func (o *GetReportListParams) WithSize(size *int32) *GetReportListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get report list params
func (o *GetReportListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
