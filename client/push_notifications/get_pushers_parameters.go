// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPushersParams creates a new GetPushersParams object
// with the default values initialized.
func NewGetPushersParams() *GetPushersParams {

	return &GetPushersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPushersParamsWithTimeout creates a new GetPushersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPushersParamsWithTimeout(timeout time.Duration) *GetPushersParams {

	return &GetPushersParams{

		timeout: timeout,
	}
}

// NewGetPushersParamsWithContext creates a new GetPushersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPushersParamsWithContext(ctx context.Context) *GetPushersParams {

	return &GetPushersParams{

		Context: ctx,
	}
}

// NewGetPushersParamsWithHTTPClient creates a new GetPushersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPushersParamsWithHTTPClient(client *http.Client) *GetPushersParams {

	return &GetPushersParams{
		HTTPClient: client,
	}
}

/*GetPushersParams contains all the parameters to send to the API endpoint
for the get pushers operation typically these are written to a http.Request
*/
type GetPushersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pushers params
func (o *GetPushersParams) WithTimeout(timeout time.Duration) *GetPushersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pushers params
func (o *GetPushersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pushers params
func (o *GetPushersParams) WithContext(ctx context.Context) *GetPushersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pushers params
func (o *GetPushersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pushers params
func (o *GetPushersParams) WithHTTPClient(client *http.Client) *GetPushersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pushers params
func (o *GetPushersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPushersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
