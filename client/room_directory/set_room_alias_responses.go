// Code generated by go-swagger; DO NOT EDIT.

package room_directory

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

// SetRoomAliasReader is a Reader for the SetRoomAlias structure.
type SetRoomAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoomAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoomAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewSetRoomAliasConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRoomAliasOK creates a SetRoomAliasOK with default headers values
func NewSetRoomAliasOK() *SetRoomAliasOK {
	return &SetRoomAliasOK{}
}

/*SetRoomAliasOK handles this case with default header values.

The mapping was created.
*/
type SetRoomAliasOK struct {
	Payload interface{}
}

func (o *SetRoomAliasOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/directory/room/{roomAlias}][%d] setRoomAliasOK  %+v", 200, o.Payload)
}

func (o *SetRoomAliasOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SetRoomAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoomAliasConflict creates a SetRoomAliasConflict with default headers values
func NewSetRoomAliasConflict() *SetRoomAliasConflict {
	return &SetRoomAliasConflict{}
}

/*SetRoomAliasConflict handles this case with default header values.

A room alias with that name already exists.
*/
type SetRoomAliasConflict struct {
	Payload *SetRoomAliasConflictBody
}

func (o *SetRoomAliasConflict) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/directory/room/{roomAlias}][%d] setRoomAliasConflict  %+v", 409, o.Payload)
}

func (o *SetRoomAliasConflict) GetPayload() *SetRoomAliasConflictBody {
	return o.Payload
}

func (o *SetRoomAliasConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetRoomAliasConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetRoomAliasBody set room alias body
swagger:model SetRoomAliasBody
*/
type SetRoomAliasBody struct {

	// The room ID to set.
	// Required: true
	RoomID *string `json:"room_id"`
}

// Validate validates this set room alias body
func (o *SetRoomAliasBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRoomID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetRoomAliasBody) validateRoomID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"room_id", "body", o.RoomID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetRoomAliasBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetRoomAliasBody) UnmarshalBinary(b []byte) error {
	var res SetRoomAliasBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetRoomAliasConflictBody A Matrix-level Error
swagger:model SetRoomAliasConflictBody
*/
type SetRoomAliasConflictBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this set room alias conflict body
func (o *SetRoomAliasConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetRoomAliasConflictBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("setRoomAliasConflict"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetRoomAliasConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetRoomAliasConflictBody) UnmarshalBinary(b []byte) error {
	var res SetRoomAliasConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}