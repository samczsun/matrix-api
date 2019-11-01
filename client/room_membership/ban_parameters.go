// Code generated by go-swagger; DO NOT EDIT.

package room_membership

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

// NewBanParams creates a new BanParams object
// with the default values initialized.
func NewBanParams() *BanParams {
	var ()
	return &BanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBanParamsWithTimeout creates a new BanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBanParamsWithTimeout(timeout time.Duration) *BanParams {
	var ()
	return &BanParams{

		timeout: timeout,
	}
}

// NewBanParamsWithContext creates a new BanParams object
// with the default values initialized, and the ability to set a context for a request
func NewBanParamsWithContext(ctx context.Context) *BanParams {
	var ()
	return &BanParams{

		Context: ctx,
	}
}

// NewBanParamsWithHTTPClient creates a new BanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBanParamsWithHTTPClient(client *http.Client) *BanParams {
	var ()
	return &BanParams{
		HTTPClient: client,
	}
}

/*BanParams contains all the parameters to send to the API endpoint
for the ban operation typically these are written to a http.Request
*/
type BanParams struct {

	/*Body*/
	Body BanBody
	/*RoomID
	  The room identifier (not alias) from which the user should be banned.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ban params
func (o *BanParams) WithTimeout(timeout time.Duration) *BanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ban params
func (o *BanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ban params
func (o *BanParams) WithContext(ctx context.Context) *BanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ban params
func (o *BanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ban params
func (o *BanParams) WithHTTPClient(client *http.Client) *BanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ban params
func (o *BanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ban params
func (o *BanParams) WithBody(body BanBody) *BanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ban params
func (o *BanParams) SetBody(body BanBody) {
	o.Body = body
}

// WithRoomID adds the roomID to the ban params
func (o *BanParams) WithRoomID(roomID string) *BanParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the ban params
func (o *BanParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *BanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param roomId
	if err := r.SetPathParam("roomId", o.RoomID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
