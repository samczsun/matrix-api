// Code generated by go-swagger; DO NOT EDIT.

package server_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWhoIsReader is a Reader for the GetWhoIs structure.
type GetWhoIsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWhoIsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWhoIsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWhoIsOK creates a GetWhoIsOK with default headers values
func NewGetWhoIsOK() *GetWhoIsOK {
	return &GetWhoIsOK{}
}

/*GetWhoIsOK handles this case with default header values.

The lookup was successful.
*/
type GetWhoIsOK struct {
	Payload *GetWhoIsOKBody
}

func (o *GetWhoIsOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/admin/whois/{userId}][%d] getWhoIsOK  %+v", 200, o.Payload)
}

func (o *GetWhoIsOK) GetPayload() *GetWhoIsOKBody {
	return o.Payload
}

func (o *GetWhoIsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWhoIsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DevicesAnon DeviceInfo
swagger:model DevicesAnon
*/
type DevicesAnon struct {

	// A user's sessions (i.e. what they did with an access token from one login).
	Sessions []*DevicesAnonSessionsItems0 `json:"sessions"`
}

// Validate validates this devices anon
func (o *DevicesAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DevicesAnon) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(o.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(o.Sessions); i++ {
		if swag.IsZero(o.Sessions[i]) { // not required
			continue
		}

		if o.Sessions[i] != nil {
			if err := o.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DevicesAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DevicesAnon) UnmarshalBinary(b []byte) error {
	var res DevicesAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DevicesAnonSessionsItems0 SessionInfo
swagger:model DevicesAnonSessionsItems0
*/
type DevicesAnonSessionsItems0 struct {

	// Information particular connections in the session.
	Connections []*DevicesAnonSessionsItems0ConnectionsItems0 `json:"connections"`
}

// Validate validates this devices anon sessions items0
func (o *DevicesAnonSessionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DevicesAnonSessionsItems0) validateConnections(formats strfmt.Registry) error {

	if swag.IsZero(o.Connections) { // not required
		return nil
	}

	for i := 0; i < len(o.Connections); i++ {
		if swag.IsZero(o.Connections[i]) { // not required
			continue
		}

		if o.Connections[i] != nil {
			if err := o.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DevicesAnonSessionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DevicesAnonSessionsItems0) UnmarshalBinary(b []byte) error {
	var res DevicesAnonSessionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DevicesAnonSessionsItems0ConnectionsItems0 ConnectionInfo
swagger:model DevicesAnonSessionsItems0ConnectionsItems0
*/
type DevicesAnonSessionsItems0ConnectionsItems0 struct {

	// Most recently seen IP address of the session.
	IP string `json:"ip,omitempty"`

	// Unix timestamp that the session was last active.
	LastSeen int64 `json:"last_seen,omitempty"`

	// User agent string last seen in the session.
	UserAgent string `json:"user_agent,omitempty"`
}

// Validate validates this devices anon sessions items0 connections items0
func (o *DevicesAnonSessionsItems0ConnectionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DevicesAnonSessionsItems0ConnectionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DevicesAnonSessionsItems0ConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res DevicesAnonSessionsItems0ConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWhoIsOKBody get who is o k body
swagger:model GetWhoIsOKBody
*/
type GetWhoIsOKBody struct {

	// Each key is an identitfier for one of the user's devices.
	Devices map[string]DevicesAnon `json:"devices,omitempty"`

	// The Matrix user ID of the user.
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this get who is o k body
func (o *GetWhoIsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWhoIsOKBody) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(o.Devices) { // not required
		return nil
	}

	for k := range o.Devices {

		if swag.IsZero(o.Devices[k]) { // not required
			continue
		}
		if val, ok := o.Devices[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWhoIsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWhoIsOKBody) UnmarshalBinary(b []byte) error {
	var res GetWhoIsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}