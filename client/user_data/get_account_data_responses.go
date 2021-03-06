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

// GetAccountDataReader is a Reader for the GetAccountData structure.
type GetAccountDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountDataOK creates a GetAccountDataOK with default headers values
func NewGetAccountDataOK() *GetAccountDataOK {
	return &GetAccountDataOK{}
}

/*GetAccountDataOK handles this case with default header values.

The account data content for the given type.
*/
type GetAccountDataOK struct {
	Payload interface{}
}

func (o *GetAccountDataOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/user/{userId}/account_data/{type}][%d] getAccountDataOK  %+v", 200, o.Payload)
}

func (o *GetAccountDataOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAccountDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
