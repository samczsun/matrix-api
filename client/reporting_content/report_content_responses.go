// Code generated by go-swagger; DO NOT EDIT.

package reporting_content

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

// ReportContentReader is a Reader for the ReportContent structure.
type ReportContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportContentOK creates a ReportContentOK with default headers values
func NewReportContentOK() *ReportContentOK {
	return &ReportContentOK{}
}

/*ReportContentOK handles this case with default header values.

The event has been reported successfully.
*/
type ReportContentOK struct {
	Payload interface{}
}

func (o *ReportContentOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/rooms/{roomId}/report/{eventId}][%d] reportContentOK  %+v", 200, o.Payload)
}

func (o *ReportContentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReportContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ReportContentBody report content body
swagger:model ReportContentBody
*/
type ReportContentBody struct {

	// The reason the content is being reported. May be blank.
	// Required: true
	Reason *string `json:"reason"`

	// The score to rate this content as where -100 is most offensive
	// and 0 is inoffensive.
	// Required: true
	Score *int64 `json:"score"`
}

// Validate validates this report content body
func (o *ReportContentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReportContentBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *ReportContentBody) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"score", "body", o.Score); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReportContentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReportContentBody) UnmarshalBinary(b []byte) error {
	var res ReportContentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}