// Code generated by go-swagger; DO NOT EDIT.

package room_membership

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

// ForgetRoomReader is a Reader for the ForgetRoom structure.
type ForgetRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForgetRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForgetRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewForgetRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewForgetRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewForgetRoomOK creates a ForgetRoomOK with default headers values
func NewForgetRoomOK() *ForgetRoomOK {
	return &ForgetRoomOK{}
}

/*ForgetRoomOK handles this case with default header values.

The room has been forgotten.
*/
type ForgetRoomOK struct {
	Payload interface{}
}

func (o *ForgetRoomOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/forget][%d] forgetRoomOK  %+v", 200, o.Payload)
}

func (o *ForgetRoomOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ForgetRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgetRoomBadRequest creates a ForgetRoomBadRequest with default headers values
func NewForgetRoomBadRequest() *ForgetRoomBadRequest {
	return &ForgetRoomBadRequest{}
}

/*ForgetRoomBadRequest handles this case with default header values.

The user has not left the room
*/
type ForgetRoomBadRequest struct {
	Payload *ForgetRoomBadRequestBody
}

func (o *ForgetRoomBadRequest) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/forget][%d] forgetRoomBadRequest  %+v", 400, o.Payload)
}

func (o *ForgetRoomBadRequest) GetPayload() *ForgetRoomBadRequestBody {
	return o.Payload
}

func (o *ForgetRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ForgetRoomBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgetRoomTooManyRequests creates a ForgetRoomTooManyRequests with default headers values
func NewForgetRoomTooManyRequests() *ForgetRoomTooManyRequests {
	return &ForgetRoomTooManyRequests{}
}

/*ForgetRoomTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type ForgetRoomTooManyRequests struct {
	Payload *ForgetRoomTooManyRequestsBody
}

func (o *ForgetRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/forget][%d] forgetRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *ForgetRoomTooManyRequests) GetPayload() *ForgetRoomTooManyRequestsBody {
	return o.Payload
}

func (o *ForgetRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ForgetRoomTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ForgetRoomBadRequestBody A Matrix-level Error
swagger:model ForgetRoomBadRequestBody
*/
type ForgetRoomBadRequestBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this forget room bad request body
func (o *ForgetRoomBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ForgetRoomBadRequestBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("forgetRoomBadRequest"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ForgetRoomBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ForgetRoomBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ForgetRoomBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ForgetRoomTooManyRequestsBody The rate limit was reached for this request
swagger:model ForgetRoomTooManyRequestsBody
*/
type ForgetRoomTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this forget room too many requests body
func (o *ForgetRoomTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ForgetRoomTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("forgetRoomTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ForgetRoomTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ForgetRoomTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res ForgetRoomTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}