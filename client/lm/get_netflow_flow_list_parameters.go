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

// NewGetNetflowFlowListParams creates a new GetNetflowFlowListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetflowFlowListParams() *GetNetflowFlowListParams {
	return &GetNetflowFlowListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetflowFlowListParamsWithTimeout creates a new GetNetflowFlowListParams object
// with the ability to set a timeout on a request.
func NewGetNetflowFlowListParamsWithTimeout(timeout time.Duration) *GetNetflowFlowListParams {
	return &GetNetflowFlowListParams{
		timeout: timeout,
	}
}

// NewGetNetflowFlowListParamsWithContext creates a new GetNetflowFlowListParams object
// with the ability to set a context for a request.
func NewGetNetflowFlowListParamsWithContext(ctx context.Context) *GetNetflowFlowListParams {
	return &GetNetflowFlowListParams{
		Context: ctx,
	}
}

// NewGetNetflowFlowListParamsWithHTTPClient creates a new GetNetflowFlowListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetflowFlowListParamsWithHTTPClient(client *http.Client) *GetNetflowFlowListParams {
	return &GetNetflowFlowListParams{
		HTTPClient: client,
	}
}

/* GetNetflowFlowListParams contains all the parameters to send to the API endpoint
   for the get netflow flow list operation.

   Typically these are written to a http.Request.
*/
type GetNetflowFlowListParams struct {

	// End.
	//
	// Format: int64
	End *int64

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// NetflowFilter.
	NetflowFilter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get netflow flow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetflowFlowListParams) WithDefaults() *GetNetflowFlowListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get netflow flow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetflowFlowListParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetNetflowFlowListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithTimeout(timeout time.Duration) *GetNetflowFlowListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithContext(ctx context.Context) *GetNetflowFlowListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithHTTPClient(client *http.Client) *GetNetflowFlowListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithEnd(end *int64) *GetNetflowFlowListParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithFields(fields *string) *GetNetflowFlowListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithFilter(filter *string) *GetNetflowFlowListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithID(id int32) *GetNetflowFlowListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithNetflowFilter(netflowFilter *string) *GetNetflowFlowListParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOffset adds the offset to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithOffset(offset *int32) *GetNetflowFlowListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithSize(size *int32) *GetNetflowFlowListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetSize(size *int32) {
	o.Size = size
}

// WithStart adds the start to the get netflow flow list params
func (o *GetNetflowFlowListParams) WithStart(start *int64) *GetNetflowFlowListParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get netflow flow list params
func (o *GetNetflowFlowListParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetflowFlowListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
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

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
