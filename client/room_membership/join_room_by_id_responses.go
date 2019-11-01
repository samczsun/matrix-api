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

// JoinRoomByIDReader is a Reader for the JoinRoomByID structure.
type JoinRoomByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JoinRoomByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJoinRoomByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewJoinRoomByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewJoinRoomByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJoinRoomByIDOK creates a JoinRoomByIDOK with default headers values
func NewJoinRoomByIDOK() *JoinRoomByIDOK {
	return &JoinRoomByIDOK{}
}

/*JoinRoomByIDOK handles this case with default header values.

The room has been joined.

The joined room ID must be returned in the ``room_id`` field.
*/
type JoinRoomByIDOK struct {
	Payload *JoinRoomByIDOKBody
}

func (o *JoinRoomByIDOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/join][%d] joinRoomByIdOK  %+v", 200, o.Payload)
}

func (o *JoinRoomByIDOK) GetPayload() *JoinRoomByIDOKBody {
	return o.Payload
}

func (o *JoinRoomByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(JoinRoomByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinRoomByIDForbidden creates a JoinRoomByIDForbidden with default headers values
func NewJoinRoomByIDForbidden() *JoinRoomByIDForbidden {
	return &JoinRoomByIDForbidden{}
}

/*JoinRoomByIDForbidden handles this case with default header values.

You do not have permission to join the room. A meaningful ``errcode`` and description error text will be returned. Example reasons for rejection are:

- The room is invite-only and the user was not invited.
- The user has been banned from the room.
*/
type JoinRoomByIDForbidden struct {
	Payload *JoinRoomByIDForbiddenBody
}

func (o *JoinRoomByIDForbidden) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/join][%d] joinRoomByIdForbidden  %+v", 403, o.Payload)
}

func (o *JoinRoomByIDForbidden) GetPayload() *JoinRoomByIDForbiddenBody {
	return o.Payload
}

func (o *JoinRoomByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(JoinRoomByIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinRoomByIDTooManyRequests creates a JoinRoomByIDTooManyRequests with default headers values
func NewJoinRoomByIDTooManyRequests() *JoinRoomByIDTooManyRequests {
	return &JoinRoomByIDTooManyRequests{}
}

/*JoinRoomByIDTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type JoinRoomByIDTooManyRequests struct {
	Payload *JoinRoomByIDTooManyRequestsBody
}

func (o *JoinRoomByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/join][%d] joinRoomByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *JoinRoomByIDTooManyRequests) GetPayload() *JoinRoomByIDTooManyRequestsBody {
	return o.Payload
}

func (o *JoinRoomByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(JoinRoomByIDTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JoinRoomByIDBody join room by ID body
swagger:model JoinRoomByIDBody
*/
type JoinRoomByIDBody struct {

	// third party signed
	ThirdPartySigned *JoinRoomByIDParamsBodyThirdPartySigned `json:"third_party_signed,omitempty"`
}

// Validate validates this join room by ID body
func (o *JoinRoomByIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateThirdPartySigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JoinRoomByIDBody) validateThirdPartySigned(formats strfmt.Registry) error {

	if swag.IsZero(o.ThirdPartySigned) { // not required
		return nil
	}

	if o.ThirdPartySigned != nil {
		if err := o.ThirdPartySigned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("third_party_signed" + "." + "third_party_signed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JoinRoomByIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinRoomByIDBody) UnmarshalBinary(b []byte) error {
	var res JoinRoomByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JoinRoomByIDForbiddenBody A Matrix-level Error
swagger:model JoinRoomByIDForbiddenBody
*/
type JoinRoomByIDForbiddenBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this join room by ID forbidden body
func (o *JoinRoomByIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JoinRoomByIDForbiddenBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("joinRoomByIdForbidden"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JoinRoomByIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinRoomByIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res JoinRoomByIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JoinRoomByIDOKBody join room by ID o k body
swagger:model JoinRoomByIDOKBody
*/
type JoinRoomByIDOKBody struct {

	// The joined room ID.
	// Required: true
	RoomID *string `json:"room_id"`
}

// Validate validates this join room by ID o k body
func (o *JoinRoomByIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRoomID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JoinRoomByIDOKBody) validateRoomID(formats strfmt.Registry) error {

	if err := validate.Required("joinRoomByIdOK"+"."+"room_id", "body", o.RoomID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JoinRoomByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinRoomByIDOKBody) UnmarshalBinary(b []byte) error {
	var res JoinRoomByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JoinRoomByIDParamsBodyThirdPartySigned Third Party Signed
//
// A signature of an ``m.third_party_invite`` token to prove that this user
// owns a third party identity which has been invited to the room.
swagger:model JoinRoomByIDParamsBodyThirdPartySigned
*/
type JoinRoomByIDParamsBodyThirdPartySigned struct {

	// The Matrix ID of the invitee.
	// Required: true
	Mxid *string `json:"mxid"`

	// The Matrix ID of the user who issued the invite.
	// Required: true
	Sender *string `json:"sender"`

	// Signatures
	//
	// A signatures object containing a signature of the entire signed object.
	// Required: true
	Signatures interface{} `json:"signatures"`

	// The state key of the m.third_party_invite event.
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this join room by ID params body third party signed
func (o *JoinRoomByIDParamsBodyThirdPartySigned) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMxid(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JoinRoomByIDParamsBodyThirdPartySigned) validateMxid(formats strfmt.Registry) error {

	if err := validate.Required("third_party_signed"+"."+"third_party_signed"+"."+"mxid", "body", o.Mxid); err != nil {
		return err
	}

	return nil
}

func (o *JoinRoomByIDParamsBodyThirdPartySigned) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("third_party_signed"+"."+"third_party_signed"+"."+"sender", "body", o.Sender); err != nil {
		return err
	}

	return nil
}

func (o *JoinRoomByIDParamsBodyThirdPartySigned) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("third_party_signed"+"."+"third_party_signed"+"."+"signatures", "body", o.Signatures); err != nil {
		return err
	}

	return nil
}

func (o *JoinRoomByIDParamsBodyThirdPartySigned) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("third_party_signed"+"."+"third_party_signed"+"."+"token", "body", o.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JoinRoomByIDParamsBodyThirdPartySigned) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinRoomByIDParamsBodyThirdPartySigned) UnmarshalBinary(b []byte) error {
	var res JoinRoomByIDParamsBodyThirdPartySigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JoinRoomByIDTooManyRequestsBody The rate limit was reached for this request
swagger:model JoinRoomByIDTooManyRequestsBody
*/
type JoinRoomByIDTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this join room by ID too many requests body
func (o *JoinRoomByIDTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JoinRoomByIDTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("joinRoomByIdTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *JoinRoomByIDTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinRoomByIDTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res JoinRoomByIDTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}