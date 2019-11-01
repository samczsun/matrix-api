// Code generated by go-swagger; DO NOT EDIT.

package device_management

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

// NewDeleteDevicesParams creates a new DeleteDevicesParams object
// with the default values initialized.
func NewDeleteDevicesParams() *DeleteDevicesParams {
	var ()
	return &DeleteDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesParamsWithTimeout creates a new DeleteDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDevicesParamsWithTimeout(timeout time.Duration) *DeleteDevicesParams {
	var ()
	return &DeleteDevicesParams{

		timeout: timeout,
	}
}

// NewDeleteDevicesParamsWithContext creates a new DeleteDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDevicesParamsWithContext(ctx context.Context) *DeleteDevicesParams {
	var ()
	return &DeleteDevicesParams{

		Context: ctx,
	}
}

// NewDeleteDevicesParamsWithHTTPClient creates a new DeleteDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDevicesParamsWithHTTPClient(client *http.Client) *DeleteDevicesParams {
	var ()
	return &DeleteDevicesParams{
		HTTPClient: client,
	}
}

/*DeleteDevicesParams contains all the parameters to send to the API endpoint
for the delete devices operation typically these are written to a http.Request
*/
type DeleteDevicesParams struct {

	/*Body*/
	Body DeleteDevicesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete devices params
func (o *DeleteDevicesParams) WithTimeout(timeout time.Duration) *DeleteDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices params
func (o *DeleteDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices params
func (o *DeleteDevicesParams) WithContext(ctx context.Context) *DeleteDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices params
func (o *DeleteDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices params
func (o *DeleteDevicesParams) WithHTTPClient(client *http.Client) *DeleteDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices params
func (o *DeleteDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete devices params
func (o *DeleteDevicesParams) WithBody(body DeleteDevicesBody) *DeleteDevicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete devices params
func (o *DeleteDevicesParams) SetBody(body DeleteDevicesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
