// Code generated by go-swagger; DO NOT EDIT.

package device_management

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

// DeleteDevicesReader is a Reader for the DeleteDevices structure.
type DeleteDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDevicesOK creates a DeleteDevicesOK with default headers values
func NewDeleteDevicesOK() *DeleteDevicesOK {
	return &DeleteDevicesOK{}
}

/*DeleteDevicesOK handles this case with default header values.

The devices were successfully removed, or had been removed
previously.
*/
type DeleteDevicesOK struct {
	Payload interface{}
}

func (o *DeleteDevicesOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/delete_devices][%d] deleteDevicesOK  %+v", 200, o.Payload)
}

func (o *DeleteDevicesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUnauthorized creates a DeleteDevicesUnauthorized with default headers values
func NewDeleteDevicesUnauthorized() *DeleteDevicesUnauthorized {
	return &DeleteDevicesUnauthorized{}
}

/*DeleteDevicesUnauthorized handles this case with default header values.

The homeserver requires additional authentication information.
*/
type DeleteDevicesUnauthorized struct {
	Payload *DeleteDevicesUnauthorizedBody
}

func (o *DeleteDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/delete_devices][%d] deleteDevicesUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDevicesUnauthorized) GetPayload() *DeleteDevicesUnauthorizedBody {
	return o.Payload
}

func (o *DeleteDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDevicesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteDevicesBody delete devices body
swagger:model DeleteDevicesBody
*/
type DeleteDevicesBody struct {

	// auth
	Auth *DeleteDevicesParamsBodyAuth `json:"auth,omitempty"`

	// The list of device IDs to delete.
	// Required: true
	Devices []string `json:"devices"`
}

// Validate validates this delete devices body
func (o *DeleteDevicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDevicesBody) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(o.Auth) { // not required
		return nil
	}

	if o.Auth != nil {
		if err := o.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "auth")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteDevicesBody) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"devices", "body", o.Devices); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDevicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDevicesBody) UnmarshalBinary(b []byte) error {
	var res DeleteDevicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteDevicesParamsBodyAuth Authentication Data
//
// Additional authentication information for the
// user-interactive authentication API.
swagger:model DeleteDevicesParamsBodyAuth
*/
type DeleteDevicesParamsBodyAuth struct {

	// The value of the session key given by the homeserver.
	Session string `json:"session,omitempty"`

	// The login type that the client is attempting to complete.
	// Required: true
	Type *string `json:"type"`

	// Keys dependent on the login type
	DeleteDevicesParamsBodyAuth map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *DeleteDevicesParamsBodyAuth) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The value of the session key given by the homeserver.
		Session string `json:"session,omitempty"`

		// The login type that the client is attempting to complete.
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DeleteDevicesParamsBodyAuth

	rcv.Session = stage1.Session

	rcv.Type = stage1.Type

	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "session")

	delete(stage2, "type")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.DeleteDevicesParamsBodyAuth = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o DeleteDevicesParamsBodyAuth) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The value of the session key given by the homeserver.
		Session string `json:"session,omitempty"`

		// The login type that the client is attempting to complete.
		// Required: true
		Type *string `json:"type"`
	}

	stage1.Session = o.Session

	stage1.Type = o.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.DeleteDevicesParamsBodyAuth) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.DeleteDevicesParamsBodyAuth)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this delete devices params body auth
func (o *DeleteDevicesParamsBodyAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDevicesParamsBodyAuth) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"auth"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDevicesParamsBodyAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDevicesParamsBodyAuth) UnmarshalBinary(b []byte) error {
	var res DeleteDevicesParamsBodyAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteDevicesUnauthorizedBody Authentication response
//
// Used by servers to indicate that additional authentication information is required,
swagger:model DeleteDevicesUnauthorizedBody
*/
type DeleteDevicesUnauthorizedBody struct {

	// A list of the stages the client has completed successfully
	Completed []string `json:"completed"`

	// Flow information
	//
	// A list of the login flows supported by the server for this API.
	// Required: true
	Flows []*FlowsItems0 `json:"flows"`

	// Contains any information that the client will need to know in order to
	// use a given type of authentication. For each login type presented,
	// that type may be present as a key in this dictionary. For example, the
	// public part of an OAuth client ID could be given here.
	Params map[string]interface{} `json:"params,omitempty"`

	// This is a session identifier that the client must pass back to the home
	// server, if one is provided, in subsequent attempts to authenticate in the
	// same API call.
	Session string `json:"session,omitempty"`
}

// Validate validates this delete devices unauthorized body
func (o *DeleteDevicesUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFlows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDevicesUnauthorizedBody) validateFlows(formats strfmt.Registry) error {

	if err := validate.Required("deleteDevicesUnauthorized"+"."+"flows", "body", o.Flows); err != nil {
		return err
	}

	for i := 0; i < len(o.Flows); i++ {
		if swag.IsZero(o.Flows[i]) { // not required
			continue
		}

		if o.Flows[i] != nil {
			if err := o.Flows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteDevicesUnauthorized" + "." + "flows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDevicesUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDevicesUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteDevicesUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
