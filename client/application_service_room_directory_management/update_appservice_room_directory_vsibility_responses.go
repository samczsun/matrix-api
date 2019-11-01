// Code generated by go-swagger; DO NOT EDIT.

package application_service_room_directory_management

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

// UpdateAppserviceRoomDirectoryVsibilityReader is a Reader for the UpdateAppserviceRoomDirectoryVsibility structure.
type UpdateAppserviceRoomDirectoryVsibilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppserviceRoomDirectoryVsibilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAppserviceRoomDirectoryVsibilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAppserviceRoomDirectoryVsibilityOK creates a UpdateAppserviceRoomDirectoryVsibilityOK with default headers values
func NewUpdateAppserviceRoomDirectoryVsibilityOK() *UpdateAppserviceRoomDirectoryVsibilityOK {
	return &UpdateAppserviceRoomDirectoryVsibilityOK{}
}

/*UpdateAppserviceRoomDirectoryVsibilityOK handles this case with default header values.

The room's directory visibility has been updated.
*/
type UpdateAppserviceRoomDirectoryVsibilityOK struct {
	Payload interface{}
}

func (o *UpdateAppserviceRoomDirectoryVsibilityOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/directory/list/appservice/{networkId}/{roomId}][%d] updateAppserviceRoomDirectoryVsibilityOK  %+v", 200, o.Payload)
}

func (o *UpdateAppserviceRoomDirectoryVsibilityOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAppserviceRoomDirectoryVsibilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateAppserviceRoomDirectoryVsibilityBody update appservice room directory vsibility body
swagger:model UpdateAppserviceRoomDirectoryVsibilityBody
*/
type UpdateAppserviceRoomDirectoryVsibilityBody struct {

	// Whether the room should be visible (public) in the directory
	// or not (private).
	// Required: true
	// Enum: [public private]
	Visibility *string `json:"visibility"`
}

// Validate validates this update appservice room directory vsibility body
func (o *UpdateAppserviceRoomDirectoryVsibilityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateAppserviceRoomDirectoryVsibilityBodyTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateAppserviceRoomDirectoryVsibilityBodyTypeVisibilityPropEnum = append(updateAppserviceRoomDirectoryVsibilityBodyTypeVisibilityPropEnum, v)
	}
}

const (

	// UpdateAppserviceRoomDirectoryVsibilityBodyVisibilityPublic captures enum value "public"
	UpdateAppserviceRoomDirectoryVsibilityBodyVisibilityPublic string = "public"

	// UpdateAppserviceRoomDirectoryVsibilityBodyVisibilityPrivate captures enum value "private"
	UpdateAppserviceRoomDirectoryVsibilityBodyVisibilityPrivate string = "private"
)

// prop value enum
func (o *UpdateAppserviceRoomDirectoryVsibilityBody) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateAppserviceRoomDirectoryVsibilityBodyTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAppserviceRoomDirectoryVsibilityBody) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"visibility", "body", o.Visibility); err != nil {
		return err
	}

	// value enum
	if err := o.validateVisibilityEnum("body"+"."+"visibility", "body", *o.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAppserviceRoomDirectoryVsibilityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAppserviceRoomDirectoryVsibilityBody) UnmarshalBinary(b []byte) error {
	var res UpdateAppserviceRoomDirectoryVsibilityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
