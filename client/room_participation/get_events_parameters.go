// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEventsParams creates a new GetEventsParams object
// with the default values initialized.
func NewGetEventsParams() *GetEventsParams {
	var ()
	return &GetEventsParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewGetEventsParamsWithTimeout creates a new GetEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsParamsWithTimeout(timeout time.Duration) *GetEventsParams {
	var ()
	return &GetEventsParams{

		requestTimeout: timeout,
	}
}

// NewGetEventsParamsWithContext creates a new GetEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsParamsWithContext(ctx context.Context) *GetEventsParams {
	var ()
	return &GetEventsParams{

		Context: ctx,
	}
}

// NewGetEventsParamsWithHTTPClient creates a new GetEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsParamsWithHTTPClient(client *http.Client) *GetEventsParams {
	var ()
	return &GetEventsParams{
		HTTPClient: client,
	}
}

/*GetEventsParams contains all the parameters to send to the API endpoint
for the get events operation typically these are written to a http.Request
*/
type GetEventsParams struct {

	/*From
	  The token to stream from. This token is either from a previous
	request to this API or from the initial sync API.

	*/
	From *string
	/*Timeout
	  The maximum time in milliseconds to wait for an event.

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the get events params
func (o *GetEventsParams) WithRequestTimeout(timeout time.Duration) *GetEventsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the get events params
func (o *GetEventsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the get events params
func (o *GetEventsParams) WithContext(ctx context.Context) *GetEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events params
func (o *GetEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) WithHTTPClient(client *http.Client) *GetEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get events params
func (o *GetEventsParams) WithFrom(from *string) *GetEventsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get events params
func (o *GetEventsParams) SetFrom(from *string) {
	o.From = from
}

// WithTimeout adds the timeout to the get events params
func (o *GetEventsParams) WithTimeout(timeout *int64) *GetEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events params
func (o *GetEventsParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
