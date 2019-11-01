// Code generated by go-swagger; DO NOT EDIT.

package user_data

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

// NewDeactivateAccountParams creates a new DeactivateAccountParams object
// with the default values initialized.
func NewDeactivateAccountParams() *DeactivateAccountParams {
	var ()
	return &DeactivateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeactivateAccountParamsWithTimeout creates a new DeactivateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeactivateAccountParamsWithTimeout(timeout time.Duration) *DeactivateAccountParams {
	var ()
	return &DeactivateAccountParams{

		timeout: timeout,
	}
}

// NewDeactivateAccountParamsWithContext creates a new DeactivateAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeactivateAccountParamsWithContext(ctx context.Context) *DeactivateAccountParams {
	var ()
	return &DeactivateAccountParams{

		Context: ctx,
	}
}

// NewDeactivateAccountParamsWithHTTPClient creates a new DeactivateAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeactivateAccountParamsWithHTTPClient(client *http.Client) *DeactivateAccountParams {
	var ()
	return &DeactivateAccountParams{
		HTTPClient: client,
	}
}

/*DeactivateAccountParams contains all the parameters to send to the API endpoint
for the deactivate account operation typically these are written to a http.Request
*/
type DeactivateAccountParams struct {

	/*Body*/
	Body DeactivateAccountBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deactivate account params
func (o *DeactivateAccountParams) WithTimeout(timeout time.Duration) *DeactivateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deactivate account params
func (o *DeactivateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deactivate account params
func (o *DeactivateAccountParams) WithContext(ctx context.Context) *DeactivateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deactivate account params
func (o *DeactivateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deactivate account params
func (o *DeactivateAccountParams) WithHTTPClient(client *http.Client) *DeactivateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deactivate account params
func (o *DeactivateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deactivate account params
func (o *DeactivateAccountParams) WithBody(body DeactivateAccountBody) *DeactivateAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deactivate account params
func (o *DeactivateAccountParams) SetBody(body DeactivateAccountBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeactivateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
