// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// SetPresenceReader is a Reader for the SetPresence structure.
type SetPresenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPresenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetPresenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewSetPresenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetPresenceOK creates a SetPresenceOK with default headers values
func NewSetPresenceOK() *SetPresenceOK {
	return &SetPresenceOK{}
}

/*SetPresenceOK handles this case with default header values.

The new presence state was set.
*/
type SetPresenceOK struct {
	Payload interface{}
}

func (o *SetPresenceOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/presence/{userId}/status][%d] setPresenceOK  %+v", 200, o.Payload)
}

func (o *SetPresenceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SetPresenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPresenceTooManyRequests creates a SetPresenceTooManyRequests with default headers values
func NewSetPresenceTooManyRequests() *SetPresenceTooManyRequests {
	return &SetPresenceTooManyRequests{}
}

/*SetPresenceTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type SetPresenceTooManyRequests struct {
	Payload *SetPresenceTooManyRequestsBody
}

func (o *SetPresenceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/presence/{userId}/status][%d] setPresenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *SetPresenceTooManyRequests) GetPayload() *SetPresenceTooManyRequestsBody {
	return o.Payload
}

func (o *SetPresenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetPresenceTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetPresenceBody set presence body
swagger:model SetPresenceBody
*/
type SetPresenceBody struct {

	// The new presence state.
	// Required: true
	// Enum: [online offline unavailable]
	Presence *string `json:"presence"`

	// The status message to attach to this state.
	StatusMsg string `json:"status_msg,omitempty"`
}

// Validate validates this set presence body
func (o *SetPresenceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePresence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var setPresenceBodyTypePresencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setPresenceBodyTypePresencePropEnum = append(setPresenceBodyTypePresencePropEnum, v)
	}
}

const (

	// SetPresenceBodyPresenceOnline captures enum value "online"
	SetPresenceBodyPresenceOnline string = "online"

	// SetPresenceBodyPresenceOffline captures enum value "offline"
	SetPresenceBodyPresenceOffline string = "offline"

	// SetPresenceBodyPresenceUnavailable captures enum value "unavailable"
	SetPresenceBodyPresenceUnavailable string = "unavailable"
)

// prop value enum
func (o *SetPresenceBody) validatePresenceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, setPresenceBodyTypePresencePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *SetPresenceBody) validatePresence(formats strfmt.Registry) error {

	if err := validate.Required("presenceState"+"."+"presence", "body", o.Presence); err != nil {
		return err
	}

	// value enum
	if err := o.validatePresenceEnum("presenceState"+"."+"presence", "body", *o.Presence); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetPresenceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetPresenceBody) UnmarshalBinary(b []byte) error {
	var res SetPresenceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetPresenceTooManyRequestsBody The rate limit was reached for this request
swagger:model SetPresenceTooManyRequestsBody
*/
type SetPresenceTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this set presence too many requests body
func (o *SetPresenceTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetPresenceTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("setPresenceTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetPresenceTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetPresenceTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res SetPresenceTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
