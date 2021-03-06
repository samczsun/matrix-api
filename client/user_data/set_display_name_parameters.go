// Code generated by go-swagger; DO NOT EDIT.

package user_data

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

// NewSetDisplayNameParams creates a new SetDisplayNameParams object
// with the default values initialized.
func NewSetDisplayNameParams() *SetDisplayNameParams {
	var ()
	return &SetDisplayNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDisplayNameParamsWithTimeout creates a new SetDisplayNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDisplayNameParamsWithTimeout(timeout time.Duration) *SetDisplayNameParams {
	var ()
	return &SetDisplayNameParams{

		timeout: timeout,
	}
}

// NewSetDisplayNameParamsWithContext creates a new SetDisplayNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDisplayNameParamsWithContext(ctx context.Context) *SetDisplayNameParams {
	var ()
	return &SetDisplayNameParams{

		Context: ctx,
	}
}

// NewSetDisplayNameParamsWithHTTPClient creates a new SetDisplayNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDisplayNameParamsWithHTTPClient(client *http.Client) *SetDisplayNameParams {
	var ()
	return &SetDisplayNameParams{
		HTTPClient: client,
	}
}

/*SetDisplayNameParams contains all the parameters to send to the API endpoint
for the set display name operation typically these are written to a http.Request
*/
type SetDisplayNameParams struct {

	/*DisplayName
	  The display name info.

	*/
	DisplayName SetDisplayNameBody
	/*UserID
	  The user whose display name to set.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set display name params
func (o *SetDisplayNameParams) WithTimeout(timeout time.Duration) *SetDisplayNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set display name params
func (o *SetDisplayNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set display name params
func (o *SetDisplayNameParams) WithContext(ctx context.Context) *SetDisplayNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set display name params
func (o *SetDisplayNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set display name params
func (o *SetDisplayNameParams) WithHTTPClient(client *http.Client) *SetDisplayNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set display name params
func (o *SetDisplayNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the set display name params
func (o *SetDisplayNameParams) WithDisplayName(displayName SetDisplayNameBody) *SetDisplayNameParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the set display name params
func (o *SetDisplayNameParams) SetDisplayName(displayName SetDisplayNameBody) {
	o.DisplayName = displayName
}

// WithUserID adds the userID to the set display name params
func (o *SetDisplayNameParams) WithUserID(userID string) *SetDisplayNameParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set display name params
func (o *SetDisplayNameParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetDisplayNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DisplayName); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
