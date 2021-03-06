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

// GetRoomIDByAliasReader is a Reader for the GetRoomIDByAlias structure.
type GetRoomIDByAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoomIDByAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoomIDByAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoomIDByAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRoomIDByAliasOK creates a GetRoomIDByAliasOK with default headers values
func NewGetRoomIDByAliasOK() *GetRoomIDByAliasOK {
	return &GetRoomIDByAliasOK{}
}

/*GetRoomIDByAliasOK handles this case with default header values.

The room ID and other information for this alias.
*/
type GetRoomIDByAliasOK struct {
	Payload *GetRoomIDByAliasOKBody
}

func (o *GetRoomIDByAliasOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/directory/room/{roomAlias}][%d] getRoomIdByAliasOK  %+v", 200, o.Payload)
}

func (o *GetRoomIDByAliasOK) GetPayload() *GetRoomIDByAliasOKBody {
	return o.Payload
}

func (o *GetRoomIDByAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRoomIDByAliasOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoomIDByAliasNotFound creates a GetRoomIDByAliasNotFound with default headers values
func NewGetRoomIDByAliasNotFound() *GetRoomIDByAliasNotFound {
	return &GetRoomIDByAliasNotFound{}
}

/*GetRoomIDByAliasNotFound handles this case with default header values.

There is no mapped room ID for this room alias.
*/
type GetRoomIDByAliasNotFound struct {
	Payload *GetRoomIDByAliasNotFoundBody
}

func (o *GetRoomIDByAliasNotFound) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/directory/room/{roomAlias}][%d] getRoomIdByAliasNotFound  %+v", 404, o.Payload)
}

func (o *GetRoomIDByAliasNotFound) GetPayload() *GetRoomIDByAliasNotFoundBody {
	return o.Payload
}

func (o *GetRoomIDByAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRoomIDByAliasNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRoomIDByAliasNotFoundBody A Matrix-level Error
swagger:model GetRoomIDByAliasNotFoundBody
*/
type GetRoomIDByAliasNotFoundBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get room ID by alias not found body
func (o *GetRoomIDByAliasNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoomIDByAliasNotFoundBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getRoomIdByAliasNotFound"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomIDByAliasNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomIDByAliasNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetRoomIDByAliasNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRoomIDByAliasOKBody get room ID by alias o k body
swagger:model GetRoomIDByAliasOKBody
*/
type GetRoomIDByAliasOKBody struct {

	// The room ID for this room alias.
	RoomID string `json:"room_id,omitempty"`

	// A list of servers that are aware of this room alias.
	Servers []string `json:"servers"`
}

// Validate validates this get room ID by alias o k body
func (o *GetRoomIDByAliasOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomIDByAliasOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomIDByAliasOKBody) UnmarshalBinary(b []byte) error {
	var res GetRoomIDByAliasOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
