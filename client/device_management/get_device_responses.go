// Code generated by go-swagger; DO NOT EDIT.

package device_management

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

// GetDeviceReader is a Reader for the GetDevice structure.
type GetDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceOK creates a GetDeviceOK with default headers values
func NewGetDeviceOK() *GetDeviceOK {
	return &GetDeviceOK{}
}

/*GetDeviceOK handles this case with default header values.

Device information
*/
type GetDeviceOK struct {
	Payload *GetDeviceOKBody
}

func (o *GetDeviceOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/devices/{deviceId}][%d] getDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceOK) GetPayload() *GetDeviceOKBody {
	return o.Payload
}

func (o *GetDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceNotFound creates a GetDeviceNotFound with default headers values
func NewGetDeviceNotFound() *GetDeviceNotFound {
	return &GetDeviceNotFound{}
}

/*GetDeviceNotFound handles this case with default header values.

The current user has no device with the given ID.
*/
type GetDeviceNotFound struct {
}

func (o *GetDeviceNotFound) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/devices/{deviceId}][%d] getDeviceNotFound ", 404)
}

func (o *GetDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetDeviceOKBody get device o k body
swagger:model GetDeviceOKBody
*/
type GetDeviceOKBody struct {

	// Identifier of this device.
	// Required: true
	DeviceID *string `json:"device_id"`

	// Display name set by the user for this device. Absent if no name has been
	// set.
	DisplayName string `json:"display_name,omitempty"`

	// The IP address where this device was last seen. (May be a few minutes out
	// of date, for efficiency reasons).
	LastSeenIP string `json:"last_seen_ip,omitempty"`

	// The timestamp (in milliseconds since the unix epoch) when this devices
	// was last seen. (May be a few minutes out of date, for efficiency
	// reasons).
	LastSeenTs int64 `json:"last_seen_ts,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetDeviceOKBody) UnmarshalJSON(raw []byte) error {
	// GetDeviceOKBodyAO0
	var dataGetDeviceOKBodyAO0 struct {
		DeviceID *string `json:"device_id"`

		DisplayName string `json:"display_name,omitempty"`

		LastSeenIP string `json:"last_seen_ip,omitempty"`

		LastSeenTs int64 `json:"last_seen_ts,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetDeviceOKBodyAO0); err != nil {
		return err
	}

	o.DeviceID = dataGetDeviceOKBodyAO0.DeviceID

	o.DisplayName = dataGetDeviceOKBodyAO0.DisplayName

	o.LastSeenIP = dataGetDeviceOKBodyAO0.LastSeenIP

	o.LastSeenTs = dataGetDeviceOKBodyAO0.LastSeenTs

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetDeviceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataGetDeviceOKBodyAO0 struct {
		DeviceID *string `json:"device_id"`

		DisplayName string `json:"display_name,omitempty"`

		LastSeenIP string `json:"last_seen_ip,omitempty"`

		LastSeenTs int64 `json:"last_seen_ts,omitempty"`
	}

	dataGetDeviceOKBodyAO0.DeviceID = o.DeviceID

	dataGetDeviceOKBodyAO0.DisplayName = o.DisplayName

	dataGetDeviceOKBodyAO0.LastSeenIP = o.LastSeenIP

	dataGetDeviceOKBodyAO0.LastSeenTs = o.LastSeenTs

	jsonDataGetDeviceOKBodyAO0, errGetDeviceOKBodyAO0 := swag.WriteJSON(dataGetDeviceOKBodyAO0)
	if errGetDeviceOKBodyAO0 != nil {
		return nil, errGetDeviceOKBodyAO0
	}
	_parts = append(_parts, jsonDataGetDeviceOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get device o k body
func (o *GetDeviceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceOKBody) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("getDeviceOK"+"."+"device_id", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}