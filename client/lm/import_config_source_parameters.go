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
)

// NewImportConfigSourceParams creates a new ImportConfigSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportConfigSourceParams() *ImportConfigSourceParams {
	return &ImportConfigSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportConfigSourceParamsWithTimeout creates a new ImportConfigSourceParams object
// with the ability to set a timeout on a request.
func NewImportConfigSourceParamsWithTimeout(timeout time.Duration) *ImportConfigSourceParams {
	return &ImportConfigSourceParams{
		timeout: timeout,
	}
}

// NewImportConfigSourceParamsWithContext creates a new ImportConfigSourceParams object
// with the ability to set a context for a request.
func NewImportConfigSourceParamsWithContext(ctx context.Context) *ImportConfigSourceParams {
	return &ImportConfigSourceParams{
		Context: ctx,
	}
}

// NewImportConfigSourceParamsWithHTTPClient creates a new ImportConfigSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportConfigSourceParamsWithHTTPClient(client *http.Client) *ImportConfigSourceParams {
	return &ImportConfigSourceParams{
		HTTPClient: client,
	}
}

/* ImportConfigSourceParams contains all the parameters to send to the API endpoint
   for the import config source operation.

   Typically these are written to a http.Request.
*/
type ImportConfigSourceParams struct {

	// File.
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import config source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportConfigSourceParams) WithDefaults() *ImportConfigSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import config source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportConfigSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import config source params
func (o *ImportConfigSourceParams) WithTimeout(timeout time.Duration) *ImportConfigSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import config source params
func (o *ImportConfigSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import config source params
func (o *ImportConfigSourceParams) WithContext(ctx context.Context) *ImportConfigSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import config source params
func (o *ImportConfigSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import config source params
func (o *ImportConfigSourceParams) WithHTTPClient(client *http.Client) *ImportConfigSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import config source params
func (o *ImportConfigSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the import config source params
func (o *ImportConfigSourceParams) WithFile(file runtime.NamedReadCloser) *ImportConfigSourceParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the import config source params
func (o *ImportConfigSourceParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *ImportConfigSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
