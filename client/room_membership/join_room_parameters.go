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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewJoinRoomParams creates a new JoinRoomParams object
// with the default values initialized.
func NewJoinRoomParams() *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJoinRoomParamsWithTimeout creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJoinRoomParamsWithTimeout(timeout time.Duration) *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		timeout: timeout,
	}
}

// NewJoinRoomParamsWithContext creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewJoinRoomParamsWithContext(ctx context.Context) *JoinRoomParams {
	var ()
	return &JoinRoomParams{

		Context: ctx,
	}
}

// NewJoinRoomParamsWithHTTPClient creates a new JoinRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJoinRoomParamsWithHTTPClient(client *http.Client) *JoinRoomParams {
	var ()
	return &JoinRoomParams{
		HTTPClient: client,
	}
}

/*JoinRoomParams contains all the parameters to send to the API endpoint
for the join room operation typically these are written to a http.Request
*/
type JoinRoomParams struct {

	/*RoomIDOrAlias
	  The room identifier or alias to join.

	*/
	RoomIDOrAlias string
	/*ServerName
	  The servers to attempt to join the room through. One of the servers
	must be participating in the room.

	*/
	ServerName []string
	/*ThirdPartySigned*/
	ThirdPartySigned JoinRoomBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the join room params
func (o *JoinRoomParams) WithTimeout(timeout time.Duration) *JoinRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the join room params
func (o *JoinRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the join room params
func (o *JoinRoomParams) WithContext(ctx context.Context) *JoinRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the join room params
func (o *JoinRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the join room params
func (o *JoinRoomParams) WithHTTPClient(client *http.Client) *JoinRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the join room params
func (o *JoinRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomIDOrAlias adds the roomIDOrAlias to the join room params
func (o *JoinRoomParams) WithRoomIDOrAlias(roomIDOrAlias string) *JoinRoomParams {
	o.SetRoomIDOrAlias(roomIDOrAlias)
	return o
}

// SetRoomIDOrAlias adds the roomIdOrAlias to the join room params
func (o *JoinRoomParams) SetRoomIDOrAlias(roomIDOrAlias string) {
	o.RoomIDOrAlias = roomIDOrAlias
}

// WithServerName adds the serverName to the join room params
func (o *JoinRoomParams) WithServerName(serverName []string) *JoinRoomParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the join room params
func (o *JoinRoomParams) SetServerName(serverName []string) {
	o.ServerName = serverName
}

// WithThirdPartySigned adds the thirdPartySigned to the join room params
func (o *JoinRoomParams) WithThirdPartySigned(thirdPartySigned JoinRoomBody) *JoinRoomParams {
	o.SetThirdPartySigned(thirdPartySigned)
	return o
}

// SetThirdPartySigned adds the thirdPartySigned to the join room params
func (o *JoinRoomParams) SetThirdPartySigned(thirdPartySigned JoinRoomBody) {
	o.ThirdPartySigned = thirdPartySigned
}

// WriteToRequest writes these params to a swagger request
func (o *JoinRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roomIdOrAlias
	if err := r.SetPathParam("roomIdOrAlias", o.RoomIDOrAlias); err != nil {
		return err
	}

	valuesServerName := o.ServerName

	joinedServerName := swag.JoinByFormat(valuesServerName, "")
	// query array param server_name
	if err := r.SetQueryParam("server_name", joinedServerName...); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ThirdPartySigned); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
