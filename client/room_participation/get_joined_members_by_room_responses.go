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

	strfmt "github.com/go-openapi/strfmt"
)

// GetJoinedMembersByRoomReader is a Reader for the GetJoinedMembersByRoom structure.
type GetJoinedMembersByRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJoinedMembersByRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJoinedMembersByRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetJoinedMembersByRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJoinedMembersByRoomOK creates a GetJoinedMembersByRoomOK with default headers values
func NewGetJoinedMembersByRoomOK() *GetJoinedMembersByRoomOK {
	return &GetJoinedMembersByRoomOK{}
}

/*GetJoinedMembersByRoomOK handles this case with default header values.

A map of MXID to room member objects.
*/
type GetJoinedMembersByRoomOK struct {
	Payload *GetJoinedMembersByRoomOKBody
}

func (o *GetJoinedMembersByRoomOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/joined_members][%d] getJoinedMembersByRoomOK  %+v", 200, o.Payload)
}

func (o *GetJoinedMembersByRoomOK) GetPayload() *GetJoinedMembersByRoomOKBody {
	return o.Payload
}

func (o *GetJoinedMembersByRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetJoinedMembersByRoomOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJoinedMembersByRoomForbidden creates a GetJoinedMembersByRoomForbidden with default headers values
func NewGetJoinedMembersByRoomForbidden() *GetJoinedMembersByRoomForbidden {
	return &GetJoinedMembersByRoomForbidden{}
}

/*GetJoinedMembersByRoomForbidden handles this case with default header values.

You aren't a member of the room.

*/
type GetJoinedMembersByRoomForbidden struct {
}

func (o *GetJoinedMembersByRoomForbidden) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/joined_members][%d] getJoinedMembersByRoomForbidden ", 403)
}

func (o *GetJoinedMembersByRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetJoinedMembersByRoomOKBody get joined members by room o k body
swagger:model GetJoinedMembersByRoomOKBody
*/
type GetJoinedMembersByRoomOKBody struct {

	// A map from user ID to a RoomMember object.
	Joined map[string]JoinedAnon `json:"joined,omitempty"`
}

// Validate validates this get joined members by room o k body
func (o *GetJoinedMembersByRoomOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJoined(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetJoinedMembersByRoomOKBody) validateJoined(formats strfmt.Registry) error {

	if swag.IsZero(o.Joined) { // not required
		return nil
	}

	for k := range o.Joined {

		if swag.IsZero(o.Joined[k]) { // not required
			continue
		}
		if val, ok := o.Joined[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetJoinedMembersByRoomOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetJoinedMembersByRoomOKBody) UnmarshalBinary(b []byte) error {
	var res GetJoinedMembersByRoomOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*JoinedAnon RoomMember
swagger:model JoinedAnon
*/
type JoinedAnon struct {

	// The mxc avatar url of the user this object is representing.
	AvatarURL string `json:"avatar_url,omitempty"`

	// The display name of the user this object is representing.
	DisplayName string `json:"display_name,omitempty"`
}

// Validate validates this joined anon
func (o *JoinedAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *JoinedAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JoinedAnon) UnmarshalBinary(b []byte) error {
	var res JoinedAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
