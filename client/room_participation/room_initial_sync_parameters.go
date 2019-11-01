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

	strfmt "github.com/go-openapi/strfmt"
)

// NewRoomInitialSyncParams creates a new RoomInitialSyncParams object
// with the default values initialized.
func NewRoomInitialSyncParams() *RoomInitialSyncParams {
	var ()
	return &RoomInitialSyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoomInitialSyncParamsWithTimeout creates a new RoomInitialSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoomInitialSyncParamsWithTimeout(timeout time.Duration) *RoomInitialSyncParams {
	var ()
	return &RoomInitialSyncParams{

		timeout: timeout,
	}
}

// NewRoomInitialSyncParamsWithContext creates a new RoomInitialSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoomInitialSyncParamsWithContext(ctx context.Context) *RoomInitialSyncParams {
	var ()
	return &RoomInitialSyncParams{

		Context: ctx,
	}
}

// NewRoomInitialSyncParamsWithHTTPClient creates a new RoomInitialSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRoomInitialSyncParamsWithHTTPClient(client *http.Client) *RoomInitialSyncParams {
	var ()
	return &RoomInitialSyncParams{
		HTTPClient: client,
	}
}

/*RoomInitialSyncParams contains all the parameters to send to the API endpoint
for the room initial sync operation typically these are written to a http.Request
*/
type RoomInitialSyncParams struct {

	/*RoomID
	  The room to get the data.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the room initial sync params
func (o *RoomInitialSyncParams) WithTimeout(timeout time.Duration) *RoomInitialSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the room initial sync params
func (o *RoomInitialSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the room initial sync params
func (o *RoomInitialSyncParams) WithContext(ctx context.Context) *RoomInitialSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the room initial sync params
func (o *RoomInitialSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the room initial sync params
func (o *RoomInitialSyncParams) WithHTTPClient(client *http.Client) *RoomInitialSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the room initial sync params
func (o *RoomInitialSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomID adds the roomID to the room initial sync params
func (o *RoomInitialSyncParams) WithRoomID(roomID string) *RoomInitialSyncParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the room initial sync params
func (o *RoomInitialSyncParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *RoomInitialSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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