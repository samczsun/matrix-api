// Code generated by go-swagger; DO NOT EDIT.

package server_administration

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

// GetWellknownReader is a Reader for the GetWellknown structure.
type GetWellknownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWellknownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWellknownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWellknownNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWellknownOK creates a GetWellknownOK with default headers values
func NewGetWellknownOK() *GetWellknownOK {
	return &GetWellknownOK{}
}

/*GetWellknownOK handles this case with default header values.

Server discovery information.
*/
type GetWellknownOK struct {
	Payload *GetWellknownOKBodyP0
}

func (o *GetWellknownOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/matrix/client][%d] getWellknownOK  %+v", 200, o.Payload)
}

func (o *GetWellknownOK) GetPayload() *GetWellknownOKBodyP0 {
	return o.Payload
}

func (o *GetWellknownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWellknownOKBodyP0)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWellknownNotFound creates a GetWellknownNotFound with default headers values
func NewGetWellknownNotFound() *GetWellknownNotFound {
	return &GetWellknownNotFound{}
}

/*GetWellknownNotFound handles this case with default header values.

No server discovery information available.
*/
type GetWellknownNotFound struct {
}

func (o *GetWellknownNotFound) Error() string {
	return fmt.Sprintf("[GET /.well-known/matrix/client][%d] getWellknownNotFound ", 404)
}

func (o *GetWellknownNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetWellknownOKBodyP0 Discovery Information
//
// Used by clients to determine the homeserver, identity server, and other
// optional components they should be interacting with.
swagger:model GetWellknownOKBodyP0
*/
type GetWellknownOKBodyP0 struct {

	// m homeserver
	// Required: true
	MHomeserver *GetWellknownOKBodyP0MHomeserver `json:"m.homeserver"`

	// m identity server
	MIdentityServer *GetWellknownOKBodyP0MIdentityServer `json:"m.identity_server,omitempty"`

	// Application-dependent keys using Java package naming convention.
	GetWellknownOKBodyP0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *GetWellknownOKBodyP0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// m homeserver
		// Required: true
		MHomeserver *GetWellknownOKBodyP0MHomeserver `json:"m.homeserver"`

		// m identity server
		MIdentityServer *GetWellknownOKBodyP0MIdentityServer `json:"m.identity_server,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GetWellknownOKBodyP0

	rcv.MHomeserver = stage1.MHomeserver

	rcv.MIdentityServer = stage1.MIdentityServer

	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "m.homeserver")

	delete(stage2, "m.identity_server")

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
		o.GetWellknownOKBodyP0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o GetWellknownOKBodyP0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// m homeserver
		// Required: true
		MHomeserver *GetWellknownOKBodyP0MHomeserver `json:"m.homeserver"`

		// m identity server
		MIdentityServer *GetWellknownOKBodyP0MIdentityServer `json:"m.identity_server,omitempty"`
	}

	stage1.MHomeserver = o.MHomeserver

	stage1.MIdentityServer = o.MIdentityServer

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.GetWellknownOKBodyP0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.GetWellknownOKBodyP0)
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

// Validate validates this get wellknown o k body p0
func (o *GetWellknownOKBodyP0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMHomeserver(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMIdentityServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWellknownOKBodyP0) validateMHomeserver(formats strfmt.Registry) error {

	if err := validate.Required("m.homeserver", "body", o.MHomeserver); err != nil {
		return err
	}

	if o.MHomeserver != nil {
		if err := o.MHomeserver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("m.homeserver")
			}
			return err
		}
	}

	return nil
}

func (o *GetWellknownOKBodyP0) validateMIdentityServer(formats strfmt.Registry) error {

	if swag.IsZero(o.MIdentityServer) { // not required
		return nil
	}

	if o.MIdentityServer != nil {
		if err := o.MIdentityServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("m.identity_server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWellknownOKBodyP0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellknownOKBodyP0) UnmarshalBinary(b []byte) error {
	var res GetWellknownOKBodyP0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWellknownOKBodyP0MHomeserver Homeserver Information
//
// Used by clients to discover homeserver information.
swagger:model GetWellknownOKBodyP0MHomeserver
*/
type GetWellknownOKBodyP0MHomeserver struct {

	// The base URL for the homeserver for client-server connections.
	// Required: true
	BaseURL *string `json:"base_url"`
}

// Validate validates this get wellknown o k body p0 m homeserver
func (o *GetWellknownOKBodyP0MHomeserver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWellknownOKBodyP0MHomeserver) validateBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("m.homeserver"+"."+"base_url", "body", o.BaseURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWellknownOKBodyP0MHomeserver) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellknownOKBodyP0MHomeserver) UnmarshalBinary(b []byte) error {
	var res GetWellknownOKBodyP0MHomeserver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWellknownOKBodyP0MIdentityServer Identity Server Information
//
// Used by clients to discover identity server information.
swagger:model GetWellknownOKBodyP0MIdentityServer
*/
type GetWellknownOKBodyP0MIdentityServer struct {

	// The base URL for the identity server for client-server connections.
	// Required: true
	BaseURL *string `json:"base_url"`
}

// Validate validates this get wellknown o k body p0 m identity server
func (o *GetWellknownOKBodyP0MIdentityServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWellknownOKBodyP0MIdentityServer) validateBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("m.identity_server"+"."+"base_url", "body", o.BaseURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWellknownOKBodyP0MIdentityServer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellknownOKBodyP0MIdentityServer) UnmarshalBinary(b []byte) error {
	var res GetWellknownOKBodyP0MIdentityServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}