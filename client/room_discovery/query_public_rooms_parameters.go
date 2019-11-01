// Code generated by go-swagger; DO NOT EDIT.

package room_discovery

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

// NewQueryPublicRoomsParams creates a new QueryPublicRoomsParams object
// with the default values initialized.
func NewQueryPublicRoomsParams() *QueryPublicRoomsParams {
	var ()
	return &QueryPublicRoomsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryPublicRoomsParamsWithTimeout creates a new QueryPublicRoomsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryPublicRoomsParamsWithTimeout(timeout time.Duration) *QueryPublicRoomsParams {
	var ()
	return &QueryPublicRoomsParams{

		timeout: timeout,
	}
}

// NewQueryPublicRoomsParamsWithContext creates a new QueryPublicRoomsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryPublicRoomsParamsWithContext(ctx context.Context) *QueryPublicRoomsParams {
	var ()
	return &QueryPublicRoomsParams{

		Context: ctx,
	}
}

// NewQueryPublicRoomsParamsWithHTTPClient creates a new QueryPublicRoomsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryPublicRoomsParamsWithHTTPClient(client *http.Client) *QueryPublicRoomsParams {
	var ()
	return &QueryPublicRoomsParams{
		HTTPClient: client,
	}
}

/*QueryPublicRoomsParams contains all the parameters to send to the API endpoint
for the query public rooms operation typically these are written to a http.Request
*/
type QueryPublicRoomsParams struct {

	/*Body
	  Options for which rooms to return.

	*/
	Body QueryPublicRoomsBody
	/*Server
	  The server to fetch the public room lists from. Defaults to the
	local server.

	*/
	Server *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query public rooms params
func (o *QueryPublicRoomsParams) WithTimeout(timeout time.Duration) *QueryPublicRoomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query public rooms params
func (o *QueryPublicRoomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query public rooms params
func (o *QueryPublicRoomsParams) WithContext(ctx context.Context) *QueryPublicRoomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query public rooms params
func (o *QueryPublicRoomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query public rooms params
func (o *QueryPublicRoomsParams) WithHTTPClient(client *http.Client) *QueryPublicRoomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query public rooms params
func (o *QueryPublicRoomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query public rooms params
func (o *QueryPublicRoomsParams) WithBody(body QueryPublicRoomsBody) *QueryPublicRoomsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query public rooms params
func (o *QueryPublicRoomsParams) SetBody(body QueryPublicRoomsBody) {
	o.Body = body
}

// WithServer adds the server to the query public rooms params
func (o *QueryPublicRoomsParams) WithServer(server *string) *QueryPublicRoomsParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the query public rooms params
func (o *QueryPublicRoomsParams) SetServer(server *string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *QueryPublicRoomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Server != nil {

		// query param server
		var qrServer string
		if o.Server != nil {
			qrServer = *o.Server
		}
		qServer := qrServer
		if qServer != "" {
			if err := r.SetQueryParam("server", qServer); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}