// Code generated by go-swagger; DO NOT EDIT.

package user_data

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

// SetDisplayNameReader is a Reader for the SetDisplayName structure.
type SetDisplayNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDisplayNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDisplayNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewSetDisplayNameTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDisplayNameOK creates a SetDisplayNameOK with default headers values
func NewSetDisplayNameOK() *SetDisplayNameOK {
	return &SetDisplayNameOK{}
}

/*SetDisplayNameOK handles this case with default header values.

The display name was set.
*/
type SetDisplayNameOK struct {
	Payload interface{}
}

func (o *SetDisplayNameOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/profile/{userId}/displayname][%d] setDisplayNameOK  %+v", 200, o.Payload)
}

func (o *SetDisplayNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SetDisplayNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDisplayNameTooManyRequests creates a SetDisplayNameTooManyRequests with default headers values
func NewSetDisplayNameTooManyRequests() *SetDisplayNameTooManyRequests {
	return &SetDisplayNameTooManyRequests{}
}

/*SetDisplayNameTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type SetDisplayNameTooManyRequests struct {
	Payload *SetDisplayNameTooManyRequestsBody
}

func (o *SetDisplayNameTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/profile/{userId}/displayname][%d] setDisplayNameTooManyRequests  %+v", 429, o.Payload)
}

func (o *SetDisplayNameTooManyRequests) GetPayload() *SetDisplayNameTooManyRequestsBody {
	return o.Payload
}

func (o *SetDisplayNameTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetDisplayNameTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetDisplayNameBody set display name body
swagger:model SetDisplayNameBody
*/
type SetDisplayNameBody struct {

	// The new display name for this user.
	Displayname string `json:"displayname,omitempty"`
}

// Validate validates this set display name body
func (o *SetDisplayNameBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetDisplayNameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetDisplayNameBody) UnmarshalBinary(b []byte) error {
	var res SetDisplayNameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetDisplayNameTooManyRequestsBody The rate limit was reached for this request
swagger:model SetDisplayNameTooManyRequestsBody
*/
type SetDisplayNameTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this set display name too many requests body
func (o *SetDisplayNameTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetDisplayNameTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("setDisplayNameTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetDisplayNameTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetDisplayNameTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res SetDisplayNameTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
