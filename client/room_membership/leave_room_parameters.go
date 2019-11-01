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

// NewLeaveRoomParams creates a new LeaveRoomParams object
// with the default values initialized.
func NewLeaveRoomParams() *LeaveRoomParams {
	var ()
	return &LeaveRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLeaveRoomParamsWithTimeout creates a new LeaveRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLeaveRoomParamsWithTimeout(timeout time.Duration) *LeaveRoomParams {
	var ()
	return &LeaveRoomParams{

		timeout: timeout,
	}
}

// NewLeaveRoomParamsWithContext creates a new LeaveRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewLeaveRoomParamsWithContext(ctx context.Context) *LeaveRoomParams {
	var ()
	return &LeaveRoomParams{

		Context: ctx,
	}
}

// NewLeaveRoomParamsWithHTTPClient creates a new LeaveRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLeaveRoomParamsWithHTTPClient(client *http.Client) *LeaveRoomParams {
	var ()
	return &LeaveRoomParams{
		HTTPClient: client,
	}
}

/*LeaveRoomParams contains all the parameters to send to the API endpoint
for the leave room operation typically these are written to a http.Request
*/
type LeaveRoomParams struct {

	/*RoomID
	  The room identifier to leave.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the leave room params
func (o *LeaveRoomParams) WithTimeout(timeout time.Duration) *LeaveRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the leave room params
func (o *LeaveRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the leave room params
func (o *LeaveRoomParams) WithContext(ctx context.Context) *LeaveRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the leave room params
func (o *LeaveRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the leave room params
func (o *LeaveRoomParams) WithHTTPClient(client *http.Client) *LeaveRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the leave room params
func (o *LeaveRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomID adds the roomID to the leave room params
func (o *LeaveRoomParams) WithRoomID(roomID string) *LeaveRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the leave room params
func (o *LeaveRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *LeaveRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roomId
	if err := r.SetPathParam("roomId", o.RoomID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
