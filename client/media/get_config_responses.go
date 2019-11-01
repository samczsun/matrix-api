// Code generated by go-swagger; DO NOT EDIT.

package media

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

// GetConfigReader is a Reader for the GetConfig structure.
type GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/*GetConfigOK handles this case with default header values.

The public content repository configuration for the matrix server.
*/
type GetConfigOK struct {
	Payload *GetConfigOKBody
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/config][%d] getConfigOK  %+v", 200, o.Payload)
}

func (o *GetConfigOK) GetPayload() *GetConfigOKBody {
	return o.Payload
}

func (o *GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConfigOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigTooManyRequests creates a GetConfigTooManyRequests with default headers values
func NewGetConfigTooManyRequests() *GetConfigTooManyRequests {
	return &GetConfigTooManyRequests{}
}

/*GetConfigTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type GetConfigTooManyRequests struct {
	Payload *GetConfigTooManyRequestsBody
}

func (o *GetConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/config][%d] getConfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigTooManyRequests) GetPayload() *GetConfigTooManyRequestsBody {
	return o.Payload
}

func (o *GetConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConfigTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetConfigOKBody get config o k body
swagger:model GetConfigOKBody
*/
type GetConfigOKBody struct {

	// The maximum size an upload can be in bytes.
	// Clients SHOULD use this as a guide when uploading content.
	// If not listed or null, the size limit should be treated as unknown.
	MUploadSize int64 `json:"m.upload.size,omitempty"`
}

// Validate validates this get config o k body
func (o *GetConfigOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetConfigTooManyRequestsBody A Matrix-level Error
swagger:model GetConfigTooManyRequestsBody
*/
type GetConfigTooManyRequestsBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get config too many requests body
func (o *GetConfigTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getConfigTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetConfigTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}