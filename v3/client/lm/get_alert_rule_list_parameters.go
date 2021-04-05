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

// NewGetAlertRuleListParams creates a new GetAlertRuleListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertRuleListParams() *GetAlertRuleListParams {
	return &GetAlertRuleListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertRuleListParamsWithTimeout creates a new GetAlertRuleListParams object
// with the ability to set a timeout on a request.
func NewGetAlertRuleListParamsWithTimeout(timeout time.Duration) *GetAlertRuleListParams {
	return &GetAlertRuleListParams{
		timeout: timeout,
	}
}

// NewGetAlertRuleListParamsWithContext creates a new GetAlertRuleListParams object
// with the ability to set a context for a request.
func NewGetAlertRuleListParamsWithContext(ctx context.Context) *GetAlertRuleListParams {
	return &GetAlertRuleListParams{
		Context: ctx,
	}
}

// NewGetAlertRuleListParamsWithHTTPClient creates a new GetAlertRuleListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertRuleListParamsWithHTTPClient(client *http.Client) *GetAlertRuleListParams {
	return &GetAlertRuleListParams{
		HTTPClient: client,
	}
}

/* GetAlertRuleListParams contains all the parameters to send to the API endpoint
   for the get alert rule list operation.

   Typically these are written to a http.Request.
*/
type GetAlertRuleListParams struct {

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

// WithDefaults hydrates default values in the get alert rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRuleListParams) WithDefaults() *GetAlertRuleListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRuleListParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAlertRuleListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alert rule list params
func (o *GetAlertRuleListParams) WithTimeout(timeout time.Duration) *GetAlertRuleListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert rule list params
func (o *GetAlertRuleListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert rule list params
func (o *GetAlertRuleListParams) WithContext(ctx context.Context) *GetAlertRuleListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert rule list params
func (o *GetAlertRuleListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert rule list params
func (o *GetAlertRuleListParams) WithHTTPClient(client *http.Client) *GetAlertRuleListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert rule list params
func (o *GetAlertRuleListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get alert rule list params
func (o *GetAlertRuleListParams) WithFields(fields *string) *GetAlertRuleListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get alert rule list params
func (o *GetAlertRuleListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get alert rule list params
func (o *GetAlertRuleListParams) WithFilter(filter *string) *GetAlertRuleListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get alert rule list params
func (o *GetAlertRuleListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get alert rule list params
func (o *GetAlertRuleListParams) WithOffset(offset *int32) *GetAlertRuleListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get alert rule list params
func (o *GetAlertRuleListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get alert rule list params
func (o *GetAlertRuleListParams) WithSize(size *int32) *GetAlertRuleListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get alert rule list params
func (o *GetAlertRuleListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertRuleListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
