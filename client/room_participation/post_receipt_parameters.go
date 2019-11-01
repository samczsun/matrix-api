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

// NewPostReceiptParams creates a new PostReceiptParams object
// with the default values initialized.
func NewPostReceiptParams() *PostReceiptParams {
	var ()
	return &PostReceiptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostReceiptParamsWithTimeout creates a new PostReceiptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostReceiptParamsWithTimeout(timeout time.Duration) *PostReceiptParams {
	var ()
	return &PostReceiptParams{

		timeout: timeout,
	}
}

// NewPostReceiptParamsWithContext creates a new PostReceiptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostReceiptParamsWithContext(ctx context.Context) *PostReceiptParams {
	var ()
	return &PostReceiptParams{

		Context: ctx,
	}
}

// NewPostReceiptParamsWithHTTPClient creates a new PostReceiptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostReceiptParamsWithHTTPClient(client *http.Client) *PostReceiptParams {
	var ()
	return &PostReceiptParams{
		HTTPClient: client,
	}
}

/*PostReceiptParams contains all the parameters to send to the API endpoint
for the post receipt operation typically these are written to a http.Request
*/
type PostReceiptParams struct {

	/*EventID
	  The event ID to acknowledge up to.

	*/
	EventID string
	/*Receipt
	  Extra receipt information to attach to ``content`` if any. The
	server will automatically set the ``ts`` field.

	*/
	Receipt interface{}
	/*ReceiptType
	  The type of receipt to send.

	*/
	ReceiptType string
	/*RoomID
	  The room in which to send the event.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post receipt params
func (o *PostReceiptParams) WithTimeout(timeout time.Duration) *PostReceiptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post receipt params
func (o *PostReceiptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post receipt params
func (o *PostReceiptParams) WithContext(ctx context.Context) *PostReceiptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post receipt params
func (o *PostReceiptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post receipt params
func (o *PostReceiptParams) WithHTTPClient(client *http.Client) *PostReceiptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post receipt params
func (o *PostReceiptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the post receipt params
func (o *PostReceiptParams) WithEventID(eventID string) *PostReceiptParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the post receipt params
func (o *PostReceiptParams) SetEventID(eventID string) {
	o.EventID = eventID
}

// WithReceipt adds the receipt to the post receipt params
func (o *PostReceiptParams) WithReceipt(receipt interface{}) *PostReceiptParams {
	o.SetReceipt(receipt)
	return o
}

// SetReceipt adds the receipt to the post receipt params
func (o *PostReceiptParams) SetReceipt(receipt interface{}) {
	o.Receipt = receipt
}

// WithReceiptType adds the receiptType to the post receipt params
func (o *PostReceiptParams) WithReceiptType(receiptType string) *PostReceiptParams {
	o.SetReceiptType(receiptType)
	return o
}

// SetReceiptType adds the receiptType to the post receipt params
func (o *PostReceiptParams) SetReceiptType(receiptType string) {
	o.ReceiptType = receiptType
}

// WithRoomID adds the roomID to the post receipt params
func (o *PostReceiptParams) WithRoomID(roomID string) *PostReceiptParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the post receipt params
func (o *PostReceiptParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *PostReceiptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param eventId
	if err := r.SetPathParam("eventId", o.EventID); err != nil {
		return err
	}

	if o.Receipt != nil {
		if err := r.SetBodyParam(o.Receipt); err != nil {
			return err
		}
	}

	// path param receiptType
	if err := r.SetPathParam("receiptType", o.ReceiptType); err != nil {
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