// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// RoomInitialSyncReader is a Reader for the RoomInitialSync structure.
type RoomInitialSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoomInitialSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoomInitialSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRoomInitialSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRoomInitialSyncOK creates a RoomInitialSyncOK with default headers values
func NewRoomInitialSyncOK() *RoomInitialSyncOK {
	return &RoomInitialSyncOK{}
}

/*RoomInitialSyncOK handles this case with default header values.

The current state of the room
*/
type RoomInitialSyncOK struct {
	Payload *RoomInitialSyncOKBody
}

func (o *RoomInitialSyncOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/initialSync][%d] roomInitialSyncOK  %+v", 200, o.Payload)
}

func (o *RoomInitialSyncOK) GetPayload() *RoomInitialSyncOKBody {
	return o.Payload
}

func (o *RoomInitialSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RoomInitialSyncOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoomInitialSyncForbidden creates a RoomInitialSyncForbidden with default headers values
func NewRoomInitialSyncForbidden() *RoomInitialSyncForbidden {
	return &RoomInitialSyncForbidden{}
}

/*RoomInitialSyncForbidden handles this case with default header values.

You aren't a member of the room and weren't previously a member of the room.

*/
type RoomInitialSyncForbidden struct {
}

func (o *RoomInitialSyncForbidden) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/initialSync][%d] roomInitialSyncForbidden ", 403)
}

func (o *RoomInitialSyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*RoomInitialSyncOKBody RoomInfo
swagger:model RoomInitialSyncOKBody
*/
type RoomInitialSyncOKBody struct {

	// The private data that this user has attached to this room.
	AccountData []*AccountDataItems0 `json:"account_data"`

	// The user's membership state in this room.
	// Enum: [invite join leave ban]
	Membership string `json:"membership,omitempty"`

	// messages
	Messages *RoomInitialSyncOKBodyMessages `json:"messages,omitempty"`

	// The ID of this room.
	// Required: true
	RoomID *string `json:"room_id"`

	// If the user is a member of the room this will be the
	// current state of the room as a list of events. If the
	// user has left the room this will be the state of the
	// room when they left it.
	State []*StateItems0 `json:"state"`

	// Whether this room is visible to the ``/publicRooms`` API
	// or not."
	// Enum: [private public]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this room initial sync o k body
func (o *RoomInitialSyncOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMembership(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoomID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RoomInitialSyncOKBody) validateAccountData(formats strfmt.Registry) error {

	if swag.IsZero(o.AccountData) { // not required
		return nil
	}

	for i := 0; i < len(o.AccountData); i++ {
		if swag.IsZero(o.AccountData[i]) { // not required
			continue
		}

		if o.AccountData[i] != nil {
			if err := o.AccountData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roomInitialSyncOK" + "." + "account_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var roomInitialSyncOKBodyTypeMembershipPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invite","join","leave","ban"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roomInitialSyncOKBodyTypeMembershipPropEnum = append(roomInitialSyncOKBodyTypeMembershipPropEnum, v)
	}
}

const (

	// RoomInitialSyncOKBodyMembershipInvite captures enum value "invite"
	RoomInitialSyncOKBodyMembershipInvite string = "invite"

	// RoomInitialSyncOKBodyMembershipJoin captures enum value "join"
	RoomInitialSyncOKBodyMembershipJoin string = "join"

	// RoomInitialSyncOKBodyMembershipLeave captures enum value "leave"
	RoomInitialSyncOKBodyMembershipLeave string = "leave"

	// RoomInitialSyncOKBodyMembershipBan captures enum value "ban"
	RoomInitialSyncOKBodyMembershipBan string = "ban"
)

// prop value enum
func (o *RoomInitialSyncOKBody) validateMembershipEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, roomInitialSyncOKBodyTypeMembershipPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RoomInitialSyncOKBody) validateMembership(formats strfmt.Registry) error {

	if swag.IsZero(o.Membership) { // not required
		return nil
	}

	// value enum
	if err := o.validateMembershipEnum("roomInitialSyncOK"+"."+"membership", "body", o.Membership); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBody) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(o.Messages) { // not required
		return nil
	}

	if o.Messages != nil {
		if err := o.Messages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roomInitialSyncOK" + "." + "messages")
			}
			return err
		}
	}

	return nil
}

func (o *RoomInitialSyncOKBody) validateRoomID(formats strfmt.Registry) error {

	if err := validate.Required("roomInitialSyncOK"+"."+"room_id", "body", o.RoomID); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBody) validateState(formats strfmt.Registry) error {

	if swag.IsZero(o.State) { // not required
		return nil
	}

	for i := 0; i < len(o.State); i++ {
		if swag.IsZero(o.State[i]) { // not required
			continue
		}

		if o.State[i] != nil {
			if err := o.State[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roomInitialSyncOK" + "." + "state" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var roomInitialSyncOKBodyTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roomInitialSyncOKBodyTypeVisibilityPropEnum = append(roomInitialSyncOKBodyTypeVisibilityPropEnum, v)
	}
}

const (

	// RoomInitialSyncOKBodyVisibilityPrivate captures enum value "private"
	RoomInitialSyncOKBodyVisibilityPrivate string = "private"

	// RoomInitialSyncOKBodyVisibilityPublic captures enum value "public"
	RoomInitialSyncOKBodyVisibilityPublic string = "public"
)

// prop value enum
func (o *RoomInitialSyncOKBody) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, roomInitialSyncOKBodyTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RoomInitialSyncOKBody) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(o.Visibility) { // not required
		return nil
	}

	// value enum
	if err := o.validateVisibilityEnum("roomInitialSyncOK"+"."+"visibility", "body", o.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBody) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RoomInitialSyncOKBodyMessages PaginationChunk
//
// The pagination chunk for this room.
swagger:model RoomInitialSyncOKBodyMessages
*/
type RoomInitialSyncOKBodyMessages struct {

	// If the user is a member of the room this will be a
	// list of the most recent messages for this room. If
	// the user has left the room this will be the
	// messages that preceeded them leaving. This array
	// will consist of at most ``limit`` elements.
	// Required: true
	Chunk []*RoomInitialSyncOKBodyMessagesChunkItems0 `json:"chunk"`

	// A token which correlates to the last value in ``chunk``.
	// Used for pagination.
	// Required: true
	End *string `json:"end"`

	// A token which correlates to the first value in ``chunk``.
	// Used for pagination.
	// Required: true
	Start *string `json:"start"`
}

// Validate validates this room initial sync o k body messages
func (o *RoomInitialSyncOKBodyMessages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChunk(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RoomInitialSyncOKBodyMessages) validateChunk(formats strfmt.Registry) error {

	if err := validate.Required("roomInitialSyncOK"+"."+"messages"+"."+"chunk", "body", o.Chunk); err != nil {
		return err
	}

	for i := 0; i < len(o.Chunk); i++ {
		if swag.IsZero(o.Chunk[i]) { // not required
			continue
		}

		if o.Chunk[i] != nil {
			if err := o.Chunk[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roomInitialSyncOK" + "." + "messages" + "." + "chunk" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessages) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("roomInitialSyncOK"+"."+"messages"+"."+"end", "body", o.End); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessages) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("roomInitialSyncOK"+"."+"messages"+"."+"start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessages) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessages) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBodyMessages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RoomInitialSyncOKBodyMessagesChunkItems0 RoomEvent
swagger:model RoomInitialSyncOKBodyMessagesChunkItems0
*/
type RoomInitialSyncOKBodyMessagesChunkItems0 struct {
	RoomInitialSyncOKBodyMessagesChunkItems0AllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RoomInitialSyncOKBodyMessagesChunkItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RoomInitialSyncOKBodyMessagesChunkItems0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RoomInitialSyncOKBodyMessagesChunkItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this room initial sync o k body messages chunk items0
func (o *RoomInitialSyncOKBodyMessagesChunkItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RoomInitialSyncOKBodyMessagesChunkItems0AllOf0
	if err := o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBodyMessagesChunkItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RoomInitialSyncOKBodyMessagesChunkItems0AllOf0 Room Event
//
// Room Events have the following fields.
swagger:model RoomInitialSyncOKBodyMessagesChunkItems0AllOf0
*/
type RoomInitialSyncOKBodyMessagesChunkItems0AllOf0 struct {
	RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0

	// The ID of the room associated with this event. Will not be present on events
	// that arrive through ``/sync``, despite being required everywhere else.
	// Required: true
	RoomID *string `json:"room_id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 = aO0

	// now for regular properties
	var propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0 struct {
		RoomID *string `json:"room_id"`
	}
	if err := swag.ReadJSON(raw, &propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0); err != nil {
		return err
	}
	o.RoomID = propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0.RoomID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0 struct {
		RoomID *string `json:"room_id"`
	}
	propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0.RoomID = o.RoomID

	jsonDataPropsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0, errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0 := swag.WriteJSON(propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0)
	if errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0 != nil {
		return nil, errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0
	}
	_parts = append(_parts, jsonDataPropsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this room initial sync o k body messages chunk items0 all of0
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0
	if err := o.RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoomID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) validateRoomID(formats strfmt.Registry) error {

	if err := validate.Required("room_id", "body", o.RoomID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBodyMessagesChunkItems0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 Room Event
//
// In addition to the Event fields, Room Events have the following additional fields.
swagger:model RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0
*/
type RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 struct {

	// The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	// Required: true
	Content interface{} `json:"content"`

	// The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	// Required: true
	Type *string `json:"type"`

	// The globally unique event identifier.
	// Required: true
	EventID *string `json:"event_id"`

	// Timestamp in milliseconds on originating homeserver when this event was sent.
	// Required: true
	OriginServerTs *int64 `json:"origin_server_ts"`

	// Contains the fully-qualified ID of the user who sent this event.
	// Required: true
	Sender *string `json:"sender"`

	// unsigned
	Unsigned *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned `json:"unsigned,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Content interface{} `json:"content"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.Content = dataAO0.Content

	o.Type = dataAO0.Type

	// now for regular properties
	var propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 struct {
		EventID *string `json:"event_id"`

		OriginServerTs *int64 `json:"origin_server_ts"`

		Sender *string `json:"sender"`

		Unsigned *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned `json:"unsigned,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0); err != nil {
		return err
	}
	o.EventID = propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.EventID

	o.OriginServerTs = propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.OriginServerTs

	o.Sender = propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.Sender

	o.Unsigned = propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.Unsigned

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Content interface{} `json:"content"`

		Type *string `json:"type"`
	}

	dataAO0.Content = o.Content

	dataAO0.Type = o.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	// now for regular properties
	var propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 struct {
		EventID *string `json:"event_id"`

		OriginServerTs *int64 `json:"origin_server_ts"`

		Sender *string `json:"sender"`

		Unsigned *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned `json:"unsigned,omitempty"`
	}
	propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.EventID = o.EventID

	propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.OriginServerTs = o.OriginServerTs

	propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.Sender = o.Sender

	propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0.Unsigned = o.Unsigned

	jsonDataPropsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0, errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 := swag.WriteJSON(propsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0)
	if errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0 != nil {
		return nil, errRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0
	}
	_parts = append(_parts, jsonDataPropsRoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this room initial sync o k body messages chunk items0 all of0 all of0
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginServerTs(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnsigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("event_id", "body", o.EventID); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateOriginServerTs(formats strfmt.Registry) error {

	if err := validate.Required("origin_server_ts", "body", o.OriginServerTs); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", o.Sender); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateUnsigned(formats strfmt.Registry) error {

	if swag.IsZero(o.Unsigned) { // not required
		return nil
	}

	if o.Unsigned != nil {
		if err := o.Unsigned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unsigned")
			}
			return err
		}
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", o.Content); err != nil {
		return err
	}

	return nil
}

func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned UnsignedData
//
// Contains optional extra information about the event.
swagger:model RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned
*/
type RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned struct {

	// The time in milliseconds that has elapsed since the event was sent. This field is generated by the local homeserver, and may be incorrect if the local time on at least one of the two servers is out of sync, which can cause the age to either be negative or greater than it actually is.
	Age int64 `json:"age,omitempty"`

	// Event
	//
	// Optional. The event that redacted this event, if any.
	RedactedBecause interface{} `json:"redacted_because,omitempty"`

	// The client-supplied transaction ID, if the client being given the event is the same one which sent it.
	TransactionID string `json:"transaction_id,omitempty"`
}

// Validate validates this room initial sync o k body messages chunk items0 all of0 all of0 unsigned
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned) UnmarshalBinary(b []byte) error {
	var res RoomInitialSyncOKBodyMessagesChunkItems0AllOf0AllOf0Unsigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
