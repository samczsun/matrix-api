// Code generated by go-swagger; DO NOT EDIT.

package open_id

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

// NewRequestOpenIDTokenParams creates a new RequestOpenIDTokenParams object
// with the default values initialized.
func NewRequestOpenIDTokenParams() *RequestOpenIDTokenParams {
	var ()
	return &RequestOpenIDTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestOpenIDTokenParamsWithTimeout creates a new RequestOpenIDTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestOpenIDTokenParamsWithTimeout(timeout time.Duration) *RequestOpenIDTokenParams {
	var ()
	return &RequestOpenIDTokenParams{

		timeout: timeout,
	}
}

// NewRequestOpenIDTokenParamsWithContext creates a new RequestOpenIDTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestOpenIDTokenParamsWithContext(ctx context.Context) *RequestOpenIDTokenParams {
	var ()
	return &RequestOpenIDTokenParams{

		Context: ctx,
	}
}

// NewRequestOpenIDTokenParamsWithHTTPClient creates a new RequestOpenIDTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestOpenIDTokenParamsWithHTTPClient(client *http.Client) *RequestOpenIDTokenParams {
	var ()
	return &RequestOpenIDTokenParams{
		HTTPClient: client,
	}
}

/*RequestOpenIDTokenParams contains all the parameters to send to the API endpoint
for the request open Id token operation typically these are written to a http.Request
*/
type RequestOpenIDTokenParams struct {

	/*Body
	  An empty object. Reserved for future expansion.

	*/
	Body interface{}
	/*UserID
	  The user to request and OpenID token for. Should be the user who
	is authenticated for the request.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request open Id token params
func (o *RequestOpenIDTokenParams) WithTimeout(timeout time.Duration) *RequestOpenIDTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request open Id token params
func (o *RequestOpenIDTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request open Id token params
func (o *RequestOpenIDTokenParams) WithContext(ctx context.Context) *RequestOpenIDTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request open Id token params
func (o *RequestOpenIDTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request open Id token params
func (o *RequestOpenIDTokenParams) WithHTTPClient(client *http.Client) *RequestOpenIDTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request open Id token params
func (o *RequestOpenIDTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the request open Id token params
func (o *RequestOpenIDTokenParams) WithBody(body interface{}) *RequestOpenIDTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the request open Id token params
func (o *RequestOpenIDTokenParams) SetBody(body interface{}) {
	o.Body = body
}

// WithUserID adds the userID to the request open Id token params
func (o *RequestOpenIDTokenParams) WithUserID(userID string) *RequestOpenIDTokenParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the request open Id token params
func (o *RequestOpenIDTokenParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RequestOpenIDTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
