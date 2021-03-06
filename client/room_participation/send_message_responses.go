// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// SendMessageReader is a Reader for the SendMessage structure.
type SendMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendMessageOK creates a SendMessageOK with default headers values
func NewSendMessageOK() *SendMessageOK {
	return &SendMessageOK{}
}

/*SendMessageOK handles this case with default header values.

An ID for the sent event.
*/
type SendMessageOK struct {
	Payload *SendMessageOKBody
}

func (o *SendMessageOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/rooms/{roomId}/send/{eventType}/{txnId}][%d] sendMessageOK  %+v", 200, o.Payload)
}

func (o *SendMessageOK) GetPayload() *SendMessageOKBody {
	return o.Payload
}

func (o *SendMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendMessageOKBody send message o k body
swagger:model SendMessageOKBody
*/
type SendMessageOKBody struct {

	// A unique identifier for the event.
	EventID string `json:"event_id,omitempty"`
}

// Validate validates this send message o k body
func (o *SendMessageOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageOKBody) UnmarshalBinary(b []byte) error {
	var res SendMessageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
