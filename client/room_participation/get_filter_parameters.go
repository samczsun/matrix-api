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

// NewGetFilterParams creates a new GetFilterParams object
// with the default values initialized.
func NewGetFilterParams() *GetFilterParams {
	var ()
	return &GetFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilterParamsWithTimeout creates a new GetFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFilterParamsWithTimeout(timeout time.Duration) *GetFilterParams {
	var ()
	return &GetFilterParams{

		timeout: timeout,
	}
}

// NewGetFilterParamsWithContext creates a new GetFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFilterParamsWithContext(ctx context.Context) *GetFilterParams {
	var ()
	return &GetFilterParams{

		Context: ctx,
	}
}

// NewGetFilterParamsWithHTTPClient creates a new GetFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFilterParamsWithHTTPClient(client *http.Client) *GetFilterParams {
	var ()
	return &GetFilterParams{
		HTTPClient: client,
	}
}

/*GetFilterParams contains all the parameters to send to the API endpoint
for the get filter operation typically these are written to a http.Request
*/
type GetFilterParams struct {

	/*FilterID
	  The filter ID to download.

	*/
	FilterID string
	/*UserID
	  The user ID to download a filter for.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get filter params
func (o *GetFilterParams) WithTimeout(timeout time.Duration) *GetFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get filter params
func (o *GetFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get filter params
func (o *GetFilterParams) WithContext(ctx context.Context) *GetFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get filter params
func (o *GetFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get filter params
func (o *GetFilterParams) WithHTTPClient(client *http.Client) *GetFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get filter params
func (o *GetFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterID adds the filterID to the get filter params
func (o *GetFilterParams) WithFilterID(filterID string) *GetFilterParams {
	o.SetFilterID(filterID)
	return o
}

// SetFilterID adds the filterId to the get filter params
func (o *GetFilterParams) SetFilterID(filterID string) {
	o.FilterID = filterID
}

// WithUserID adds the userID to the get filter params
func (o *GetFilterParams) WithUserID(userID string) *GetFilterParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get filter params
func (o *GetFilterParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filterId
	if err := r.SetPathParam("filterId", o.FilterID); err != nil {
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
