// Code generated by go-swagger; DO NOT EDIT.

package user_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAccountDataPerRoomReader is a Reader for the GetAccountDataPerRoom structure.
type GetAccountDataPerRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountDataPerRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountDataPerRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountDataPerRoomOK creates a GetAccountDataPerRoomOK with default headers values
func NewGetAccountDataPerRoomOK() *GetAccountDataPerRoomOK {
	return &GetAccountDataPerRoomOK{}
}

/*GetAccountDataPerRoomOK handles this case with default header values.

The account data content for the given type.
*/
type GetAccountDataPerRoomOK struct {
	Payload interface{}
}

func (o *GetAccountDataPerRoomOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/user/{userId}/rooms/{roomId}/account_data/{type}][%d] getAccountDataPerRoomOK  %+v", 200, o.Payload)
}

func (o *GetAccountDataPerRoomOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAccountDataPerRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
