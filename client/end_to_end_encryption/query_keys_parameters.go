// Code generated by go-swagger; DO NOT EDIT.

package end_to_end_encryption

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

// NewQueryKeysParams creates a new QueryKeysParams object
// with the default values initialized.
func NewQueryKeysParams() *QueryKeysParams {
	var ()
	return &QueryKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryKeysParamsWithTimeout creates a new QueryKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryKeysParamsWithTimeout(timeout time.Duration) *QueryKeysParams {
	var ()
	return &QueryKeysParams{

		timeout: timeout,
	}
}

// NewQueryKeysParamsWithContext creates a new QueryKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryKeysParamsWithContext(ctx context.Context) *QueryKeysParams {
	var ()
	return &QueryKeysParams{

		Context: ctx,
	}
}

// NewQueryKeysParamsWithHTTPClient creates a new QueryKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryKeysParamsWithHTTPClient(client *http.Client) *QueryKeysParams {
	var ()
	return &QueryKeysParams{
		HTTPClient: client,
	}
}

/*QueryKeysParams contains all the parameters to send to the API endpoint
for the query keys operation typically these are written to a http.Request
*/
type QueryKeysParams struct {

	/*Query
	  Query defining the keys to be downloaded

	*/
	Query QueryKeysBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query keys params
func (o *QueryKeysParams) WithTimeout(timeout time.Duration) *QueryKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query keys params
func (o *QueryKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query keys params
func (o *QueryKeysParams) WithContext(ctx context.Context) *QueryKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query keys params
func (o *QueryKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query keys params
func (o *QueryKeysParams) WithHTTPClient(client *http.Client) *QueryKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query keys params
func (o *QueryKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the query keys params
func (o *QueryKeysParams) WithQuery(query QueryKeysBody) *QueryKeysParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the query keys params
func (o *QueryKeysParams) SetQuery(query QueryKeysBody) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *QueryKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Query); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
