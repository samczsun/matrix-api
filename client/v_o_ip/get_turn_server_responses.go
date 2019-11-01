// Code generated by go-swagger; DO NOT EDIT.

package v_o_ip

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

// GetTurnServerReader is a Reader for the GetTurnServer structure.
type GetTurnServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTurnServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTurnServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetTurnServerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTurnServerOK creates a GetTurnServerOK with default headers values
func NewGetTurnServerOK() *GetTurnServerOK {
	return &GetTurnServerOK{}
}

/*GetTurnServerOK handles this case with default header values.

The TURN server credentials.
*/
type GetTurnServerOK struct {
	Payload *GetTurnServerOKBody
}

func (o *GetTurnServerOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/voip/turnServer][%d] getTurnServerOK  %+v", 200, o.Payload)
}

func (o *GetTurnServerOK) GetPayload() *GetTurnServerOKBody {
	return o.Payload
}

func (o *GetTurnServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTurnServerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTurnServerTooManyRequests creates a GetTurnServerTooManyRequests with default headers values
func NewGetTurnServerTooManyRequests() *GetTurnServerTooManyRequests {
	return &GetTurnServerTooManyRequests{}
}

/*GetTurnServerTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type GetTurnServerTooManyRequests struct {
	Payload *GetTurnServerTooManyRequestsBody
}

func (o *GetTurnServerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/voip/turnServer][%d] getTurnServerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTurnServerTooManyRequests) GetPayload() *GetTurnServerTooManyRequestsBody {
	return o.Payload
}

func (o *GetTurnServerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTurnServerTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTurnServerOKBody get turn server o k body
swagger:model GetTurnServerOKBody
*/
type GetTurnServerOKBody struct {

	// The password to use.
	// Required: true
	Password *string `json:"password"`

	// The time-to-live in seconds
	// Required: true
	TTL *int64 `json:"ttl"`

	// A list of TURN URIs
	// Required: true
	Uris []string `json:"uris"`

	// The username to use.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this get turn server o k body
func (o *GetTurnServerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUris(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTurnServerOKBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("getTurnServerOK"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *GetTurnServerOKBody) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("getTurnServerOK"+"."+"ttl", "body", o.TTL); err != nil {
		return err
	}

	return nil
}

func (o *GetTurnServerOKBody) validateUris(formats strfmt.Registry) error {

	if err := validate.Required("getTurnServerOK"+"."+"uris", "body", o.Uris); err != nil {
		return err
	}

	return nil
}

func (o *GetTurnServerOKBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("getTurnServerOK"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTurnServerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTurnServerOKBody) UnmarshalBinary(b []byte) error {
	var res GetTurnServerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTurnServerTooManyRequestsBody The rate limit was reached for this request
swagger:model GetTurnServerTooManyRequestsBody
*/
type GetTurnServerTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this get turn server too many requests body
func (o *GetTurnServerTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTurnServerTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getTurnServerTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTurnServerTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTurnServerTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetTurnServerTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}