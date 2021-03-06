// Code generated by go-swagger; DO NOT EDIT.

package application_service_room_directory_management

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

// NewUpdateAppserviceRoomDirectoryVsibilityParams creates a new UpdateAppserviceRoomDirectoryVsibilityParams object
// with the default values initialized.
func NewUpdateAppserviceRoomDirectoryVsibilityParams() *UpdateAppserviceRoomDirectoryVsibilityParams {
	var ()
	return &UpdateAppserviceRoomDirectoryVsibilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppserviceRoomDirectoryVsibilityParamsWithTimeout creates a new UpdateAppserviceRoomDirectoryVsibilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppserviceRoomDirectoryVsibilityParamsWithTimeout(timeout time.Duration) *UpdateAppserviceRoomDirectoryVsibilityParams {
	var ()
	return &UpdateAppserviceRoomDirectoryVsibilityParams{

		timeout: timeout,
	}
}

// NewUpdateAppserviceRoomDirectoryVsibilityParamsWithContext creates a new UpdateAppserviceRoomDirectoryVsibilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppserviceRoomDirectoryVsibilityParamsWithContext(ctx context.Context) *UpdateAppserviceRoomDirectoryVsibilityParams {
	var ()
	return &UpdateAppserviceRoomDirectoryVsibilityParams{

		Context: ctx,
	}
}

// NewUpdateAppserviceRoomDirectoryVsibilityParamsWithHTTPClient creates a new UpdateAppserviceRoomDirectoryVsibilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppserviceRoomDirectoryVsibilityParamsWithHTTPClient(client *http.Client) *UpdateAppserviceRoomDirectoryVsibilityParams {
	var ()
	return &UpdateAppserviceRoomDirectoryVsibilityParams{
		HTTPClient: client,
	}
}

/*UpdateAppserviceRoomDirectoryVsibilityParams contains all the parameters to send to the API endpoint
for the update appservice room directory vsibility operation typically these are written to a http.Request
*/
type UpdateAppserviceRoomDirectoryVsibilityParams struct {

	/*Body*/
	Body UpdateAppserviceRoomDirectoryVsibilityBody
	/*NetworkID
	  The protocol (network) ID to update the room list for. This would
	have been provided by the application service as being listed as
	a supported protocol.

	*/
	NetworkID string
	/*RoomID
	  The room ID to add to the directory.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithTimeout(timeout time.Duration) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithContext(ctx context.Context) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithHTTPClient(client *http.Client) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithBody(body UpdateAppserviceRoomDirectoryVsibilityBody) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetBody(body UpdateAppserviceRoomDirectoryVsibilityBody) {
	o.Body = body
}

// WithNetworkID adds the networkID to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithNetworkID(networkID string) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRoomID adds the roomID to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WithRoomID(roomID string) *UpdateAppserviceRoomDirectoryVsibilityParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the update appservice room directory vsibility params
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppserviceRoomDirectoryVsibilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
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
