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

// InviteBy3PIDReader is a Reader for the InviteBy3PID structure.
type InviteBy3PIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InviteBy3PIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInviteBy3PIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInviteBy3PIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInviteBy3PIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInviteBy3PIDOK creates a InviteBy3PIDOK with default headers values
func NewInviteBy3PIDOK() *InviteBy3PIDOK {
	return &InviteBy3PIDOK{}
}

/*InviteBy3PIDOK handles this case with default header values.

The user has been invited to join the room.
*/
type InviteBy3PIDOK struct {
	Payload interface{}
}

func (o *InviteBy3PIDOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/invite][%d] inviteBy3PIdOK  %+v", 200, o.Payload)
}

func (o *InviteBy3PIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *InviteBy3PIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInviteBy3PIDForbidden creates a InviteBy3PIDForbidden with default headers values
func NewInviteBy3PIDForbidden() *InviteBy3PIDForbidden {
	return &InviteBy3PIDForbidden{}
}

/*InviteBy3PIDForbidden handles this case with default header values.

You do not have permission to invite the user to the room. A meaningful ``errcode`` and description error text will be returned. Example reasons for rejections are:

- The invitee has been banned from the room.
- The invitee is already a member of the room.
- The inviter is not currently in the room.
- The inviter's power level is insufficient to invite users to the room.
*/
type InviteBy3PIDForbidden struct {
	Payload *InviteBy3PIDForbiddenBody
}

func (o *InviteBy3PIDForbidden) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/invite][%d] inviteBy3PIdForbidden  %+v", 403, o.Payload)
}

func (o *InviteBy3PIDForbidden) GetPayload() *InviteBy3PIDForbiddenBody {
	return o.Payload
}

func (o *InviteBy3PIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InviteBy3PIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInviteBy3PIDTooManyRequests creates a InviteBy3PIDTooManyRequests with default headers values
func NewInviteBy3PIDTooManyRequests() *InviteBy3PIDTooManyRequests {
	return &InviteBy3PIDTooManyRequests{}
}

/*InviteBy3PIDTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type InviteBy3PIDTooManyRequests struct {
	Payload *InviteBy3PIDTooManyRequestsBody
}

func (o *InviteBy3PIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/invite][%d] inviteBy3PIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *InviteBy3PIDTooManyRequests) GetPayload() *InviteBy3PIDTooManyRequestsBody {
	return o.Payload
}

func (o *InviteBy3PIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InviteBy3PIDTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*InviteBy3PIDBody invite by3 p ID body
swagger:model InviteBy3PIDBody
*/
type InviteBy3PIDBody struct {

	// The invitee's third party identifier.
	// Required: true
	Address *string `json:"address"`

	// An access token previously registered with the identity server. Servers
	// can treat this as optional to distinguish between r0.5-compatible clients
	// and this specification version.
	// Required: true
	IDAccessToken *string `json:"id_access_token"`

	// The hostname+port of the identity server which should be used for third party identifier lookups.
	// Required: true
	IDServer *string `json:"id_server"`

	// The kind of address being passed in the address field, for example ``email``.
	// Required: true
	Medium *string `json:"medium"`
}

// Validate validates this invite by3 p ID body
func (o *InviteBy3PIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIDAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIDServer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InviteBy3PIDBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	return nil
}

func (o *InviteBy3PIDBody) validateIDAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"id_access_token", "body", o.IDAccessToken); err != nil {
		return err
	}

	return nil
}

func (o *InviteBy3PIDBody) validateIDServer(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"id_server", "body", o.IDServer); err != nil {
		return err
	}

	return nil
}

func (o *InviteBy3PIDBody) validateMedium(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"medium", "body", o.Medium); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InviteBy3PIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InviteBy3PIDBody) UnmarshalBinary(b []byte) error {
	var res InviteBy3PIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InviteBy3PIDForbiddenBody A Matrix-level Error
swagger:model InviteBy3PIDForbiddenBody
*/
type InviteBy3PIDForbiddenBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this invite by3 p ID forbidden body
func (o *InviteBy3PIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InviteBy3PIDForbiddenBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("inviteBy3PIdForbidden"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InviteBy3PIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InviteBy3PIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res InviteBy3PIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InviteBy3PIDTooManyRequestsBody The rate limit was reached for this request
swagger:model InviteBy3PIDTooManyRequestsBody
*/
type InviteBy3PIDTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this invite by3 p ID too many requests body
func (o *InviteBy3PIDTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InviteBy3PIDTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("inviteBy3PIdTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InviteBy3PIDTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InviteBy3PIDTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res InviteBy3PIDTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}