// Code generated by go-swagger; DO NOT EDIT.

package session_management

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

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoginTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

The user has been authenticated.
*/
type LoginOK struct {
	Payload *LoginOKBody
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/login][%d] loginOK  %+v", 200, o.Payload)
}

func (o *LoginOK) GetPayload() *LoginOKBody {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginBadRequest creates a LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {
	return &LoginBadRequest{}
}

/*LoginBadRequest handles this case with default header values.

Part of the request was invalid. For example, the login type may not be recognised.
*/
type LoginBadRequest struct {
	Payload *LoginBadRequestBody
}

func (o *LoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/login][%d] loginBadRequest  %+v", 400, o.Payload)
}

func (o *LoginBadRequest) GetPayload() *LoginBadRequestBody {
	return o.Payload
}

func (o *LoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginForbidden creates a LoginForbidden with default headers values
func NewLoginForbidden() *LoginForbidden {
	return &LoginForbidden{}
}

/*LoginForbidden handles this case with default header values.

The login attempt failed. This can include one of the following error codes:
  * ``M_FORBIDDEN``: The provided authentication data was incorrect.
  * ``M_USER_DEACTIVATED``: The user has been deactivated.
*/
type LoginForbidden struct {
	Payload *LoginForbiddenBody
}

func (o *LoginForbidden) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/login][%d] loginForbidden  %+v", 403, o.Payload)
}

func (o *LoginForbidden) GetPayload() *LoginForbiddenBody {
	return o.Payload
}

func (o *LoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginTooManyRequests creates a LoginTooManyRequests with default headers values
func NewLoginTooManyRequests() *LoginTooManyRequests {
	return &LoginTooManyRequests{}
}

/*LoginTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type LoginTooManyRequests struct {
	Payload *LoginTooManyRequestsBody
}

func (o *LoginTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/login][%d] loginTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoginTooManyRequests) GetPayload() *LoginTooManyRequestsBody {
	return o.Payload
}

func (o *LoginTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LoginBadRequestBody A Matrix-level Error
swagger:model LoginBadRequestBody
*/
type LoginBadRequestBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this login bad request body
func (o *LoginBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginBadRequestBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("loginBadRequest"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LoginBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginBody login body
swagger:model LoginBody
*/
type LoginBody struct {

	// Third party identifier for the user.  Deprecated in favour of ``identifier``.
	Address string `json:"address,omitempty"`

	// ID of the client device. If this does not correspond to a
	// known client device, a new device will be created. The server
	// will auto-generate a device_id if this is not specified.
	DeviceID string `json:"device_id,omitempty"`

	// identifier
	Identifier *LoginParamsBodyIdentifier `json:"identifier,omitempty"`

	// A display name to assign to the newly-created device. Ignored
	// if ``device_id`` corresponds to a known device.
	InitialDeviceDisplayName string `json:"initial_device_display_name,omitempty"`

	// When logging in using a third party identifier, the medium of the identifier. Must be 'email'.  Deprecated in favour of ``identifier``.
	Medium string `json:"medium,omitempty"`

	// Required when ``type`` is ``m.login.password``. The user's
	// password.
	Password string `json:"password,omitempty"`

	// Required when ``type`` is ``m.login.token``. Part of `Token-based`_ login.
	Token string `json:"token,omitempty"`

	// The login type being used.
	// Required: true
	// Enum: [m.login.password m.login.token]
	Type *string `json:"type"`

	// The fully qualified user ID or just local part of the user ID, to log in.  Deprecated in favour of ``identifier``.
	User string `json:"user,omitempty"`
}

// Validate validates this login body
func (o *LoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginBody) validateIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(o.Identifier) { // not required
		return nil
	}

	if o.Identifier != nil {
		if err := o.Identifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "identifier")
			}
			return err
		}
	}

	return nil
}

var loginBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m.login.password","m.login.token"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loginBodyTypeTypePropEnum = append(loginBodyTypeTypePropEnum, v)
	}
}

const (

	// LoginBodyTypeMLoginPassword captures enum value "m.login.password"
	LoginBodyTypeMLoginPassword string = "m.login.password"

	// LoginBodyTypeMLoginToken captures enum value "m.login.token"
	LoginBodyTypeMLoginToken string = "m.login.token"
)

// prop value enum
func (o *LoginBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loginBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *LoginBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginBody) UnmarshalBinary(b []byte) error {
	var res LoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginForbiddenBody A Matrix-level Error
swagger:model LoginForbiddenBody
*/
type LoginForbiddenBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this login forbidden body
func (o *LoginForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginForbiddenBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("loginForbidden"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginForbiddenBody) UnmarshalBinary(b []byte) error {
	var res LoginForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginOKBody login o k body
swagger:model LoginOKBody
*/
type LoginOKBody struct {

	// An access token for the account.
	// This access token can then be used to authorize other requests.
	AccessToken string `json:"access_token,omitempty"`

	// ID of the logged-in device. Will be the same as the
	// corresponding parameter in the request, if one was specified.
	DeviceID string `json:"device_id,omitempty"`

	// The server_name of the homeserver on which the account has
	// been registered.
	//
	// **Deprecated**. Clients should extract the server_name from
	// ``user_id`` (by splitting at the first colon) if they require
	// it. Note also that ``homeserver`` is not spelt this way.
	HomeServer string `json:"home_server,omitempty"`

	// The fully-qualified Matrix ID that has been registered.
	UserID string `json:"user_id,omitempty"`

	// well known
	WellKnown *LoginOKBodyWellKnown `json:"well_known,omitempty"`
}

// Validate validates this login o k body
func (o *LoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWellKnown(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBody) validateWellKnown(formats strfmt.Registry) error {

	if swag.IsZero(o.WellKnown) { // not required
		return nil
	}

	if o.WellKnown != nil {
		if err := o.WellKnown.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "well_known")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBody) UnmarshalBinary(b []byte) error {
	var res LoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginOKBodyWellKnown Discovery Information
//
// Optional client configuration provided by the server. If present,
// clients SHOULD use the provided object to reconfigure themselves,
// optionally validating the URLs within. This object takes the same
// form as the one returned from .well-known autodiscovery.
swagger:model LoginOKBodyWellKnown
*/
type LoginOKBodyWellKnown struct {

	// m homeserver
	// Required: true
	MHomeserver *LoginOKBodyWellKnownMHomeserver `json:"m.homeserver"`

	// m identity server
	MIdentityServer *LoginOKBodyWellKnownMIdentityServer `json:"m.identity_server,omitempty"`

	// Application-dependent keys using Java package naming convention.
	LoginOKBodyWellKnown map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *LoginOKBodyWellKnown) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// m homeserver
		// Required: true
		MHomeserver *LoginOKBodyWellKnownMHomeserver `json:"m.homeserver"`

		// m identity server
		MIdentityServer *LoginOKBodyWellKnownMIdentityServer `json:"m.identity_server,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LoginOKBodyWellKnown

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
		o.LoginOKBodyWellKnown = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o LoginOKBodyWellKnown) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// m homeserver
		// Required: true
		MHomeserver *LoginOKBodyWellKnownMHomeserver `json:"m.homeserver"`

		// m identity server
		MIdentityServer *LoginOKBodyWellKnownMIdentityServer `json:"m.identity_server,omitempty"`
	}

	stage1.MHomeserver = o.MHomeserver

	stage1.MIdentityServer = o.MIdentityServer

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.LoginOKBodyWellKnown) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.LoginOKBodyWellKnown)
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

// Validate validates this login o k body well known
func (o *LoginOKBodyWellKnown) Validate(formats strfmt.Registry) error {
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

func (o *LoginOKBodyWellKnown) validateMHomeserver(formats strfmt.Registry) error {

	if err := validate.Required("loginOK"+"."+"well_known"+"."+"m.homeserver", "body", o.MHomeserver); err != nil {
		return err
	}

	if o.MHomeserver != nil {
		if err := o.MHomeserver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "well_known" + "." + "m.homeserver")
			}
			return err
		}
	}

	return nil
}

func (o *LoginOKBodyWellKnown) validateMIdentityServer(formats strfmt.Registry) error {

	if swag.IsZero(o.MIdentityServer) { // not required
		return nil
	}

	if o.MIdentityServer != nil {
		if err := o.MIdentityServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "well_known" + "." + "m.identity_server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBodyWellKnown) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBodyWellKnown) UnmarshalBinary(b []byte) error {
	var res LoginOKBodyWellKnown
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginOKBodyWellKnownMHomeserver Homeserver Information
//
// Used by clients to discover homeserver information.
swagger:model LoginOKBodyWellKnownMHomeserver
*/
type LoginOKBodyWellKnownMHomeserver struct {

	// The base URL for the homeserver for client-server connections.
	// Required: true
	BaseURL *string `json:"base_url"`
}

// Validate validates this login o k body well known m homeserver
func (o *LoginOKBodyWellKnownMHomeserver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBodyWellKnownMHomeserver) validateBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("loginOK"+"."+"well_known"+"."+"m.homeserver"+"."+"base_url", "body", o.BaseURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBodyWellKnownMHomeserver) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBodyWellKnownMHomeserver) UnmarshalBinary(b []byte) error {
	var res LoginOKBodyWellKnownMHomeserver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginOKBodyWellKnownMIdentityServer Identity Server Information
//
// Used by clients to discover identity server information.
swagger:model LoginOKBodyWellKnownMIdentityServer
*/
type LoginOKBodyWellKnownMIdentityServer struct {

	// The base URL for the identity server for client-server connections.
	// Required: true
	BaseURL *string `json:"base_url"`
}

// Validate validates this login o k body well known m identity server
func (o *LoginOKBodyWellKnownMIdentityServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBodyWellKnownMIdentityServer) validateBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("loginOK"+"."+"well_known"+"."+"m.identity_server"+"."+"base_url", "body", o.BaseURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBodyWellKnownMIdentityServer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBodyWellKnownMIdentityServer) UnmarshalBinary(b []byte) error {
	var res LoginOKBodyWellKnownMIdentityServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginParamsBodyIdentifier User identifier
//
// Identification information for the user.
swagger:model LoginParamsBodyIdentifier
*/
type LoginParamsBodyIdentifier struct {

	// The type of identification.  See `Identifier types`_ for supported values and additional property descriptions.
	// Required: true
	Type *string `json:"type"`

	// login params body identifier additional properties
	LoginParamsBodyIdentifierAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *LoginParamsBodyIdentifier) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The type of identification.  See `Identifier types`_ for supported values and additional property descriptions.
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LoginParamsBodyIdentifier

	rcv.Type = stage1.Type

	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

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
		o.LoginParamsBodyIdentifierAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o LoginParamsBodyIdentifier) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The type of identification.  See `Identifier types`_ for supported values and additional property descriptions.
		// Required: true
		Type *string `json:"type"`
	}

	stage1.Type = o.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.LoginParamsBodyIdentifierAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.LoginParamsBodyIdentifierAdditionalProperties)
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

// Validate validates this login params body identifier
func (o *LoginParamsBodyIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginParamsBodyIdentifier) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"identifier"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginParamsBodyIdentifier) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginParamsBodyIdentifier) UnmarshalBinary(b []byte) error {
	var res LoginParamsBodyIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LoginTooManyRequestsBody The rate limit was reached for this request
swagger:model LoginTooManyRequestsBody
*/
type LoginTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this login too many requests body
func (o *LoginTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("loginTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res LoginTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
