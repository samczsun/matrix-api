// Code generated by go-swagger; DO NOT EDIT.

package media

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

// NewGetContentParams creates a new GetContentParams object
// with the default values initialized.
func NewGetContentParams() *GetContentParams {
	var (
		allowRemoteDefault = bool(true)
	)
	return &GetContentParams{
		AllowRemote: &allowRemoteDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentParamsWithTimeout creates a new GetContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentParamsWithTimeout(timeout time.Duration) *GetContentParams {
	var (
		allowRemoteDefault = bool(true)
	)
	return &GetContentParams{
		AllowRemote: &allowRemoteDefault,

		timeout: timeout,
	}
}

// NewGetContentParamsWithContext creates a new GetContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentParamsWithContext(ctx context.Context) *GetContentParams {
	var (
		allowRemoteDefault = bool(true)
	)
	return &GetContentParams{
		AllowRemote: &allowRemoteDefault,

		Context: ctx,
	}
}

// NewGetContentParamsWithHTTPClient creates a new GetContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentParamsWithHTTPClient(client *http.Client) *GetContentParams {
	var (
		allowRemoteDefault = bool(true)
	)
	return &GetContentParams{
		AllowRemote: &allowRemoteDefault,
		HTTPClient:  client,
	}
}

/*GetContentParams contains all the parameters to send to the API endpoint
for the get content operation typically these are written to a http.Request
*/
type GetContentParams struct {

	/*AllowRemote
	  Indicates to the server that it should not attempt to fetch the media if it is deemed
	remote. This is to prevent routing loops where the server contacts itself. Defaults to
	true if not provided.


	*/
	AllowRemote *bool
	/*MediaID
	  The media ID from the ``mxc://`` URI (the path component)


	*/
	MediaID string
	/*ServerName
	  The server name from the ``mxc://`` URI (the authoritory component)


	*/
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get content params
func (o *GetContentParams) WithTimeout(timeout time.Duration) *GetContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get content params
func (o *GetContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get content params
func (o *GetContentParams) WithContext(ctx context.Context) *GetContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get content params
func (o *GetContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get content params
func (o *GetContentParams) WithHTTPClient(client *http.Client) *GetContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get content params
func (o *GetContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowRemote adds the allowRemote to the get content params
func (o *GetContentParams) WithAllowRemote(allowRemote *bool) *GetContentParams {
	o.SetAllowRemote(allowRemote)
	return o
}

// SetAllowRemote adds the allowRemote to the get content params
func (o *GetContentParams) SetAllowRemote(allowRemote *bool) {
	o.AllowRemote = allowRemote
}

// WithMediaID adds the mediaID to the get content params
func (o *GetContentParams) WithMediaID(mediaID string) *GetContentParams {
	o.SetMediaID(mediaID)
	return o
}

// SetMediaID adds the mediaId to the get content params
func (o *GetContentParams) SetMediaID(mediaID string) {
	o.MediaID = mediaID
}

// WithServerName adds the serverName to the get content params
func (o *GetContentParams) WithServerName(serverName string) *GetContentParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the get content params
func (o *GetContentParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowRemote != nil {

		// query param allow_remote
		var qrAllowRemote bool
		if o.AllowRemote != nil {
			qrAllowRemote = *o.AllowRemote
		}
		qAllowRemote := swag.FormatBool(qrAllowRemote)
		if qAllowRemote != "" {
			if err := r.SetQueryParam("allow_remote", qAllowRemote); err != nil {
				return err
			}
		}

	}

	// path param mediaId
	if err := r.SetPathParam("mediaId", o.MediaID); err != nil {
		return err
	}

	// path param serverName
	if err := r.SetPathParam("serverName", o.ServerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
