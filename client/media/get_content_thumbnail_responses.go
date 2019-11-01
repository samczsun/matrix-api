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

// GetContentThumbnailReader is a Reader for the GetContentThumbnail structure.
type GetContentThumbnailReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetContentThumbnailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentThumbnailOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentThumbnailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentThumbnailRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentThumbnailTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetContentThumbnailBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContentThumbnailOK creates a GetContentThumbnailOK with default headers values
func NewGetContentThumbnailOK(writer io.Writer) *GetContentThumbnailOK {
	return &GetContentThumbnailOK{
		Payload: writer,
	}
}

/*GetContentThumbnailOK handles this case with default header values.

A thumbnail of the requested content.
*/
type GetContentThumbnailOK struct {
	/*The content type of the thumbnail.
	 */
	ContentType string

	Payload io.Writer
}

func (o *GetContentThumbnailOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailOK  %+v", 200, o.Payload)
}

func (o *GetContentThumbnailOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetContentThumbnailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentThumbnailBadRequest creates a GetContentThumbnailBadRequest with default headers values
func NewGetContentThumbnailBadRequest() *GetContentThumbnailBadRequest {
	return &GetContentThumbnailBadRequest{}
}

/*GetContentThumbnailBadRequest handles this case with default header values.

The request does not make sense to the server, or the server cannot thumbnail
the content. For example, the client requested non-integer dimensions or asked
for negatively-sized images.
*/
type GetContentThumbnailBadRequest struct {
	Payload *GetContentThumbnailBadRequestBody
}

func (o *GetContentThumbnailBadRequest) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentThumbnailBadRequest) GetPayload() *GetContentThumbnailBadRequestBody {
	return o.Payload
}

func (o *GetContentThumbnailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContentThumbnailBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentThumbnailRequestEntityTooLarge creates a GetContentThumbnailRequestEntityTooLarge with default headers values
func NewGetContentThumbnailRequestEntityTooLarge() *GetContentThumbnailRequestEntityTooLarge {
	return &GetContentThumbnailRequestEntityTooLarge{}
}

/*GetContentThumbnailRequestEntityTooLarge handles this case with default header values.

The local content is too large for the server to thumbnail.
*/
type GetContentThumbnailRequestEntityTooLarge struct {
	Payload *GetContentThumbnailRequestEntityTooLargeBody
}

func (o *GetContentThumbnailRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentThumbnailRequestEntityTooLarge) GetPayload() *GetContentThumbnailRequestEntityTooLargeBody {
	return o.Payload
}

func (o *GetContentThumbnailRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContentThumbnailRequestEntityTooLargeBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentThumbnailTooManyRequests creates a GetContentThumbnailTooManyRequests with default headers values
func NewGetContentThumbnailTooManyRequests() *GetContentThumbnailTooManyRequests {
	return &GetContentThumbnailTooManyRequests{}
}

/*GetContentThumbnailTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type GetContentThumbnailTooManyRequests struct {
	Payload *GetContentThumbnailTooManyRequestsBody
}

func (o *GetContentThumbnailTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentThumbnailTooManyRequests) GetPayload() *GetContentThumbnailTooManyRequestsBody {
	return o.Payload
}

func (o *GetContentThumbnailTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContentThumbnailTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentThumbnailBadGateway creates a GetContentThumbnailBadGateway with default headers values
func NewGetContentThumbnailBadGateway() *GetContentThumbnailBadGateway {
	return &GetContentThumbnailBadGateway{}
}

/*GetContentThumbnailBadGateway handles this case with default header values.

The remote content is too large for the server to thumbnail.
*/
type GetContentThumbnailBadGateway struct {
	Payload *GetContentThumbnailBadGatewayBody
}

func (o *GetContentThumbnailBadGateway) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailBadGateway  %+v", 502, o.Payload)
}

func (o *GetContentThumbnailBadGateway) GetPayload() *GetContentThumbnailBadGatewayBody {
	return o.Payload
}

func (o *GetContentThumbnailBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContentThumbnailBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetContentThumbnailBadGatewayBody A Matrix-level Error
swagger:model GetContentThumbnailBadGatewayBody
*/
type GetContentThumbnailBadGatewayBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get content thumbnail bad gateway body
func (o *GetContentThumbnailBadGatewayBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContentThumbnailBadGatewayBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getContentThumbnailBadGateway"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContentThumbnailBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContentThumbnailBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res GetContentThumbnailBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContentThumbnailBadRequestBody A Matrix-level Error
swagger:model GetContentThumbnailBadRequestBody
*/
type GetContentThumbnailBadRequestBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get content thumbnail bad request body
func (o *GetContentThumbnailBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContentThumbnailBadRequestBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getContentThumbnailBadRequest"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContentThumbnailBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContentThumbnailBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetContentThumbnailBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContentThumbnailRequestEntityTooLargeBody A Matrix-level Error
swagger:model GetContentThumbnailRequestEntityTooLargeBody
*/
type GetContentThumbnailRequestEntityTooLargeBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get content thumbnail request entity too large body
func (o *GetContentThumbnailRequestEntityTooLargeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContentThumbnailRequestEntityTooLargeBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getContentThumbnailRequestEntityTooLarge"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContentThumbnailRequestEntityTooLargeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContentThumbnailRequestEntityTooLargeBody) UnmarshalBinary(b []byte) error {
	var res GetContentThumbnailRequestEntityTooLargeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContentThumbnailTooManyRequestsBody The rate limit was reached for this request
swagger:model GetContentThumbnailTooManyRequestsBody
*/
type GetContentThumbnailTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this get content thumbnail too many requests body
func (o *GetContentThumbnailTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContentThumbnailTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("getContentThumbnailTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContentThumbnailTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContentThumbnailTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetContentThumbnailTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
