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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchRecipientGroupByIDParams creates a new PatchRecipientGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRecipientGroupByIDParams() *PatchRecipientGroupByIDParams {
	return &PatchRecipientGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRecipientGroupByIDParamsWithTimeout creates a new PatchRecipientGroupByIDParams object
// with the ability to set a timeout on a request.
func NewPatchRecipientGroupByIDParamsWithTimeout(timeout time.Duration) *PatchRecipientGroupByIDParams {
	return &PatchRecipientGroupByIDParams{
		timeout: timeout,
	}
}

// NewPatchRecipientGroupByIDParamsWithContext creates a new PatchRecipientGroupByIDParams object
// with the ability to set a context for a request.
func NewPatchRecipientGroupByIDParamsWithContext(ctx context.Context) *PatchRecipientGroupByIDParams {
	return &PatchRecipientGroupByIDParams{
		Context: ctx,
	}
}

// NewPatchRecipientGroupByIDParamsWithHTTPClient creates a new PatchRecipientGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRecipientGroupByIDParamsWithHTTPClient(client *http.Client) *PatchRecipientGroupByIDParams {
	return &PatchRecipientGroupByIDParams{
		HTTPClient: client,
	}
}

/* PatchRecipientGroupByIDParams contains all the parameters to send to the API endpoint
   for the patch recipient group by Id operation.

   Typically these are written to a http.Request.
*/
type PatchRecipientGroupByIDParams struct {

	// Body.
	Body *models.RecipientGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch recipient group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRecipientGroupByIDParams) WithDefaults() *PatchRecipientGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch recipient group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRecipientGroupByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) WithTimeout(timeout time.Duration) *PatchRecipientGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) WithContext(ctx context.Context) *PatchRecipientGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) WithHTTPClient(client *http.Client) *PatchRecipientGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) WithBody(body *models.RecipientGroup) *PatchRecipientGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) SetBody(body *models.RecipientGroup) {
	o.Body = body
}

// WithID adds the id to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) WithID(id int32) *PatchRecipientGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch recipient group by Id params
func (o *PatchRecipientGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRecipientGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
