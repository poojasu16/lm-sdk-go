// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSDTHistoryByDeviceGroupIDParams creates a new GetSDTHistoryByDeviceGroupIDParams object
// with the default values initialized.
func NewGetSDTHistoryByDeviceGroupIDParams() *GetSDTHistoryByDeviceGroupIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetSDTHistoryByDeviceGroupIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSDTHistoryByDeviceGroupIDParamsWithTimeout creates a new GetSDTHistoryByDeviceGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSDTHistoryByDeviceGroupIDParamsWithTimeout(timeout time.Duration) *GetSDTHistoryByDeviceGroupIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetSDTHistoryByDeviceGroupIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetSDTHistoryByDeviceGroupIDParamsWithContext creates a new GetSDTHistoryByDeviceGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSDTHistoryByDeviceGroupIDParamsWithContext(ctx context.Context) *GetSDTHistoryByDeviceGroupIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetSDTHistoryByDeviceGroupIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetSDTHistoryByDeviceGroupIDParamsWithHTTPClient creates a new GetSDTHistoryByDeviceGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSDTHistoryByDeviceGroupIDParamsWithHTTPClient(client *http.Client) *GetSDTHistoryByDeviceGroupIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetSDTHistoryByDeviceGroupIDParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetSDTHistoryByDeviceGroupIDParams contains all the parameters to send to the API endpoint
for the get SDT history by device group Id operation typically these are written to a http.Request
*/
type GetSDTHistoryByDeviceGroupIDParams struct {

	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*ID*/
	ID int32
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithTimeout(timeout time.Duration) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithContext(ctx context.Context) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithHTTPClient(client *http.Client) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithFields(fields *string) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithFilter(filter *string) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithID(id int32) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithOffset(offset *int32) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) WithSize(size *int32) *GetSDTHistoryByDeviceGroupIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get SDT history by device group Id params
func (o *GetSDTHistoryByDeviceGroupIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetSDTHistoryByDeviceGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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
