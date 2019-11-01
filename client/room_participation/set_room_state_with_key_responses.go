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

// SetRoomStateWithKeyReader is a Reader for the SetRoomStateWithKey structure.
type SetRoomStateWithKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoomStateWithKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoomStateWithKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSetRoomStateWithKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRoomStateWithKeyOK creates a SetRoomStateWithKeyOK with default headers values
func NewSetRoomStateWithKeyOK() *SetRoomStateWithKeyOK {
	return &SetRoomStateWithKeyOK{}
}

/*SetRoomStateWithKeyOK handles this case with default header values.

An ID for the sent event.
*/
type SetRoomStateWithKeyOK struct {
	Payload *SetRoomStateWithKeyOKBody
}

func (o *SetRoomStateWithKeyOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey}][%d] setRoomStateWithKeyOK  %+v", 200, o.Payload)
}

func (o *SetRoomStateWithKeyOK) GetPayload() *SetRoomStateWithKeyOKBody {
	return o.Payload
}

func (o *SetRoomStateWithKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetRoomStateWithKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoomStateWithKeyForbidden creates a SetRoomStateWithKeyForbidden with default headers values
func NewSetRoomStateWithKeyForbidden() *SetRoomStateWithKeyForbidden {
	return &SetRoomStateWithKeyForbidden{}
}

/*SetRoomStateWithKeyForbidden handles this case with default header values.

The sender doesn't have permission to send the event into the room.
*/
type SetRoomStateWithKeyForbidden struct {
	Payload *SetRoomStateWithKeyForbiddenBody
}

func (o *SetRoomStateWithKeyForbidden) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey}][%d] setRoomStateWithKeyForbidden  %+v", 403, o.Payload)
}

func (o *SetRoomStateWithKeyForbidden) GetPayload() *SetRoomStateWithKeyForbiddenBody {
	return o.Payload
}

func (o *SetRoomStateWithKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetRoomStateWithKeyForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetRoomStateWithKeyForbiddenBody A Matrix-level Error
swagger:model SetRoomStateWithKeyForbiddenBody
*/
type SetRoomStateWithKeyForbiddenBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this set room state with key forbidden body
func (o *SetRoomStateWithKeyForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetRoomStateWithKeyForbiddenBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("setRoomStateWithKeyForbidden"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetRoomStateWithKeyForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetRoomStateWithKeyForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SetRoomStateWithKeyForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetRoomStateWithKeyOKBody set room state with key o k body
swagger:model SetRoomStateWithKeyOKBody
*/
type SetRoomStateWithKeyOKBody struct {

	// A unique identifier for the event.
	EventID string `json:"event_id,omitempty"`
}

// Validate validates this set room state with key o k body
func (o *SetRoomStateWithKeyOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetRoomStateWithKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetRoomStateWithKeyOKBody) UnmarshalBinary(b []byte) error {
	var res SetRoomStateWithKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}