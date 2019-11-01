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

// DeleteRoomTagReader is a Reader for the DeleteRoomTag structure.
type DeleteRoomTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoomTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoomTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoomTagOK creates a DeleteRoomTagOK with default headers values
func NewDeleteRoomTagOK() *DeleteRoomTagOK {
	return &DeleteRoomTagOK{}
}

/*DeleteRoomTagOK handles this case with default header values.

The tag was successfully removed.
*/
type DeleteRoomTagOK struct {
	Payload interface{}
}

func (o *DeleteRoomTagOK) Error() string {
	return fmt.Sprintf("[DELETE /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags/{tag}][%d] deleteRoomTagOK  %+v", 200, o.Payload)
}

func (o *DeleteRoomTagOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRoomTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
