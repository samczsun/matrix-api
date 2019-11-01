// Code generated by go-swagger; DO NOT EDIT.

package room_ugprades

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

// NewUpgradeRoomParams creates a new UpgradeRoomParams object
// with the default values initialized.
func NewUpgradeRoomParams() *UpgradeRoomParams {
	var ()
	return &UpgradeRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeRoomParamsWithTimeout creates a new UpgradeRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeRoomParamsWithTimeout(timeout time.Duration) *UpgradeRoomParams {
	var ()
	return &UpgradeRoomParams{

		timeout: timeout,
	}
}

// NewUpgradeRoomParamsWithContext creates a new UpgradeRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeRoomParamsWithContext(ctx context.Context) *UpgradeRoomParams {
	var ()
	return &UpgradeRoomParams{

		Context: ctx,
	}
}

// NewUpgradeRoomParamsWithHTTPClient creates a new UpgradeRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeRoomParamsWithHTTPClient(client *http.Client) *UpgradeRoomParams {
	var ()
	return &UpgradeRoomParams{
		HTTPClient: client,
	}
}

/*UpgradeRoomParams contains all the parameters to send to the API endpoint
for the upgrade room operation typically these are written to a http.Request
*/
type UpgradeRoomParams struct {

	/*Body
	  The request body

	*/
	Body UpgradeRoomBody
	/*RoomID
	  The ID of the room to upgrade.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade room params
func (o *UpgradeRoomParams) WithTimeout(timeout time.Duration) *UpgradeRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade room params
func (o *UpgradeRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade room params
func (o *UpgradeRoomParams) WithContext(ctx context.Context) *UpgradeRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade room params
func (o *UpgradeRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade room params
func (o *UpgradeRoomParams) WithHTTPClient(client *http.Client) *UpgradeRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade room params
func (o *UpgradeRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade room params
func (o *UpgradeRoomParams) WithBody(body UpgradeRoomBody) *UpgradeRoomParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade room params
func (o *UpgradeRoomParams) SetBody(body UpgradeRoomBody) {
	o.Body = body
}

// WithRoomID adds the roomID to the upgrade room params
func (o *UpgradeRoomParams) WithRoomID(roomID string) *UpgradeRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the upgrade room params
func (o *UpgradeRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
