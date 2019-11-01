// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

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

// PostPusherReader is a Reader for the PostPusher structure.
type PostPusherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPusherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPusherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPusherBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPusherTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPusherOK creates a PostPusherOK with default headers values
func NewPostPusherOK() *PostPusherOK {
	return &PostPusherOK{}
}

/*PostPusherOK handles this case with default header values.

The pusher was set.
*/
type PostPusherOK struct {
	Payload interface{}
}

func (o *PostPusherOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/pushers/set][%d] postPusherOK  %+v", 200, o.Payload)
}

func (o *PostPusherOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPusherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPusherBadRequest creates a PostPusherBadRequest with default headers values
func NewPostPusherBadRequest() *PostPusherBadRequest {
	return &PostPusherBadRequest{}
}

/*PostPusherBadRequest handles this case with default header values.

One or more of the pusher values were invalid.
*/
type PostPusherBadRequest struct {
	Payload *PostPusherBadRequestBody
}

func (o *PostPusherBadRequest) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/pushers/set][%d] postPusherBadRequest  %+v", 400, o.Payload)
}

func (o *PostPusherBadRequest) GetPayload() *PostPusherBadRequestBody {
	return o.Payload
}

func (o *PostPusherBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPusherBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPusherTooManyRequests creates a PostPusherTooManyRequests with default headers values
func NewPostPusherTooManyRequests() *PostPusherTooManyRequests {
	return &PostPusherTooManyRequests{}
}

/*PostPusherTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type PostPusherTooManyRequests struct {
	Payload *PostPusherTooManyRequestsBody
}

func (o *PostPusherTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/pushers/set][%d] postPusherTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPusherTooManyRequests) GetPayload() *PostPusherTooManyRequestsBody {
	return o.Payload
}

func (o *PostPusherTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPusherTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostPusherBadRequestBody A Matrix-level Error
swagger:model PostPusherBadRequestBody
*/
type PostPusherBadRequestBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this post pusher bad request body
func (o *PostPusherBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPusherBadRequestBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("postPusherBadRequest"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPusherBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPusherBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostPusherBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPusherBody post pusher body
swagger:model PostPusherBody
*/
type PostPusherBody struct {

	// A string that will allow the user to identify what application
	// owns this pusher.
	// Required: true
	AppDisplayName *string `json:"app_display_name"`

	// This is a reverse-DNS style identifier for the application.
	// It is recommended that this end with the platform, such that
	// different platform versions get different app identifiers.
	// Max length, 64 chars.
	//
	// If the ``kind`` is ``"email"``, this is ``"m.email"``.
	// Required: true
	AppID *string `json:"app_id"`

	// If true, the homeserver should add another pusher with the
	// given pushkey and App ID in addition to any others with
	// different user IDs. Otherwise, the homeserver must remove any
	// other pushers with the same App ID and pushkey for different
	// users. The default is ``false``.
	Append bool `json:"append,omitempty"`

	// data
	// Required: true
	Data *PostPusherParamsBodyData `json:"data"`

	// A string that will allow the user to identify what device owns
	// this pusher.
	// Required: true
	DeviceDisplayName *string `json:"device_display_name"`

	// The kind of pusher to configure. ``"http"`` makes a pusher that
	// sends HTTP pokes. ``"email"`` makes a pusher that emails the
	// user with unread notifications. ``null`` deletes the pusher.
	// Required: true
	Kind *string `json:"kind"`

	// The preferred language for receiving notifications (e.g. 'en'
	// or 'en-US').
	// Required: true
	Lang *string `json:"lang"`

	// This string determines which set of device specific rules this
	// pusher executes.
	ProfileTag string `json:"profile_tag,omitempty"`

	// This is a unique identifier for this pusher. The value you
	// should use for this is the routing or destination address
	// information for the notification, for example, the APNS token
	// for APNS or the Registration ID for GCM. If your notification
	// client has no such concept, use any unique identifier.
	// Max length, 512 bytes.
	//
	// If the ``kind`` is ``"email"``, this is the email address to
	// send notifications to.
	// Required: true
	Pushkey *string `json:"pushkey"`
}

// Validate validates this post pusher body
func (o *PostPusherBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAppDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeviceDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLang(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePushkey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPusherBody) validateAppDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"app_display_name", "body", o.AppDisplayName); err != nil {
		return err
	}

	return nil
}

func (o *PostPusherBody) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"app_id", "body", o.AppID); err != nil {
		return err
	}

	return nil
}

func (o *PostPusherBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pusher" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *PostPusherBody) validateDeviceDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"device_display_name", "body", o.DeviceDisplayName); err != nil {
		return err
	}

	return nil
}

func (o *PostPusherBody) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

func (o *PostPusherBody) validateLang(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"lang", "body", o.Lang); err != nil {
		return err
	}

	return nil
}

func (o *PostPusherBody) validatePushkey(formats strfmt.Registry) error {

	if err := validate.Required("pusher"+"."+"pushkey", "body", o.Pushkey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPusherBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPusherBody) UnmarshalBinary(b []byte) error {
	var res PostPusherBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPusherParamsBodyData PusherData
//
// A dictionary of information for the pusher implementation
// itself. If ``kind`` is ``http``, this should contain ``url``
// which is the URL to use to send notifications to.
swagger:model PostPusherParamsBodyData
*/
type PostPusherParamsBodyData struct {

	// The format to send notifications in to Push Gateways if the
	// ``kind`` is ``http``. The details about what fields the
	// homeserver should send to the push gateway are defined in the
	// `Push Gateway Specification`_. Currently the only format
	// available is 'event_id_only'.
	Format string `json:"format,omitempty"`

	// Required if ``kind`` is ``http``. The URL to use to send
	// notifications to. MUST be an HTTPS URL with a path of
	// ``/_matrix/push/v1/notify``.
	URL string `json:"url,omitempty"`
}

// Validate validates this post pusher params body data
func (o *PostPusherParamsBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPusherParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPusherParamsBodyData) UnmarshalBinary(b []byte) error {
	var res PostPusherParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPusherTooManyRequestsBody The rate limit was reached for this request
swagger:model PostPusherTooManyRequestsBody
*/
type PostPusherTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this post pusher too many requests body
func (o *PostPusherTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPusherTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("postPusherTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPusherTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPusherTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res PostPusherTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}