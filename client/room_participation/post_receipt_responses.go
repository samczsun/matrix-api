// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// PostReceiptReader is a Reader for the PostReceipt structure.
type PostReceiptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReceiptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostReceiptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPostReceiptTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostReceiptOK creates a PostReceiptOK with default headers values
func NewPostReceiptOK() *PostReceiptOK {
	return &PostReceiptOK{}
}

/*PostReceiptOK handles this case with default header values.

The receipt was sent.
*/
type PostReceiptOK struct {
	Payload interface{}
}

func (o *PostReceiptOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/receipt/{receiptType}/{eventId}][%d] postReceiptOK  %+v", 200, o.Payload)
}

func (o *PostReceiptOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostReceiptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReceiptTooManyRequests creates a PostReceiptTooManyRequests with default headers values
func NewPostReceiptTooManyRequests() *PostReceiptTooManyRequests {
	return &PostReceiptTooManyRequests{}
}

/*PostReceiptTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type PostReceiptTooManyRequests struct {
	Payload *PostReceiptTooManyRequestsBody
}

func (o *PostReceiptTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/receipt/{receiptType}/{eventId}][%d] postReceiptTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostReceiptTooManyRequests) GetPayload() *PostReceiptTooManyRequestsBody {
	return o.Payload
}

func (o *PostReceiptTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostReceiptTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostReceiptTooManyRequestsBody The rate limit was reached for this request
swagger:model PostReceiptTooManyRequestsBody
*/
type PostReceiptTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this post receipt too many requests body
func (o *PostReceiptTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostReceiptTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("postReceiptTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostReceiptTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostReceiptTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res PostReceiptTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
