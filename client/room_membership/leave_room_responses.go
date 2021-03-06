// Code generated by go-swagger; DO NOT EDIT.

package room_membership

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

// LeaveRoomReader is a Reader for the LeaveRoom structure.
type LeaveRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeaveRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLeaveRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewLeaveRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeaveRoomOK creates a LeaveRoomOK with default headers values
func NewLeaveRoomOK() *LeaveRoomOK {
	return &LeaveRoomOK{}
}

/*LeaveRoomOK handles this case with default header values.

The room has been left.
*/
type LeaveRoomOK struct {
	Payload interface{}
}

func (o *LeaveRoomOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/leave][%d] leaveRoomOK  %+v", 200, o.Payload)
}

func (o *LeaveRoomOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LeaveRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaveRoomTooManyRequests creates a LeaveRoomTooManyRequests with default headers values
func NewLeaveRoomTooManyRequests() *LeaveRoomTooManyRequests {
	return &LeaveRoomTooManyRequests{}
}

/*LeaveRoomTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type LeaveRoomTooManyRequests struct {
	Payload *LeaveRoomTooManyRequestsBody
}

func (o *LeaveRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/leave][%d] leaveRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *LeaveRoomTooManyRequests) GetPayload() *LeaveRoomTooManyRequestsBody {
	return o.Payload
}

func (o *LeaveRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LeaveRoomTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LeaveRoomTooManyRequestsBody The rate limit was reached for this request
swagger:model LeaveRoomTooManyRequestsBody
*/
type LeaveRoomTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this leave room too many requests body
func (o *LeaveRoomTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LeaveRoomTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("leaveRoomTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LeaveRoomTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LeaveRoomTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res LeaveRoomTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
