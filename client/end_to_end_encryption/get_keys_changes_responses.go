// Code generated by go-swagger; DO NOT EDIT.

package end_to_end_encryption

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetKeysChangesReader is a Reader for the GetKeysChanges structure.
type GetKeysChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeysChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeysChangesOK creates a GetKeysChangesOK with default headers values
func NewGetKeysChangesOK() *GetKeysChangesOK {
	return &GetKeysChangesOK{}
}

/*GetKeysChangesOK handles this case with default header values.

The list of users who updated their devices.
*/
type GetKeysChangesOK struct {
	Payload *GetKeysChangesOKBody
}

func (o *GetKeysChangesOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/keys/changes][%d] getKeysChangesOK  %+v", 200, o.Payload)
}

func (o *GetKeysChangesOK) GetPayload() *GetKeysChangesOKBody {
	return o.Payload
}

func (o *GetKeysChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetKeysChangesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetKeysChangesOKBody get keys changes o k body
swagger:model GetKeysChangesOKBody
*/
type GetKeysChangesOKBody struct {

	// The Matrix User IDs of all users who updated their device
	// identity keys.
	Changed []string `json:"changed"`

	// The Matrix User IDs of all users who may have left all
	// the end-to-end encrypted rooms they previously shared
	// with the user.
	Left []string `json:"left"`
}

// Validate validates this get keys changes o k body
func (o *GetKeysChangesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetKeysChangesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKeysChangesOKBody) UnmarshalBinary(b []byte) error {
	var res GetKeysChangesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}