// Code generated by go-swagger; DO NOT EDIT.

package user_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// SearchUserDirectoryReader is a Reader for the SearchUserDirectory structure.
type SearchUserDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUserDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewSearchUserDirectoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchUserDirectoryOK creates a SearchUserDirectoryOK with default headers values
func NewSearchUserDirectoryOK() *SearchUserDirectoryOK {
	return &SearchUserDirectoryOK{}
}

/*SearchUserDirectoryOK handles this case with default header values.

The results of the search.
*/
type SearchUserDirectoryOK struct {
	Payload *SearchUserDirectoryOKBody
}

func (o *SearchUserDirectoryOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/user_directory/search][%d] searchUserDirectoryOK  %+v", 200, o.Payload)
}

func (o *SearchUserDirectoryOK) GetPayload() *SearchUserDirectoryOKBody {
	return o.Payload
}

func (o *SearchUserDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchUserDirectoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserDirectoryTooManyRequests creates a SearchUserDirectoryTooManyRequests with default headers values
func NewSearchUserDirectoryTooManyRequests() *SearchUserDirectoryTooManyRequests {
	return &SearchUserDirectoryTooManyRequests{}
}

/*SearchUserDirectoryTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type SearchUserDirectoryTooManyRequests struct {
	Payload *SearchUserDirectoryTooManyRequestsBody
}

func (o *SearchUserDirectoryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/user_directory/search][%d] searchUserDirectoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchUserDirectoryTooManyRequests) GetPayload() *SearchUserDirectoryTooManyRequestsBody {
	return o.Payload
}

func (o *SearchUserDirectoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchUserDirectoryTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ResultsItems0 User
swagger:model ResultsItems0
*/
type ResultsItems0 struct {

	// The avatar url, as an MXC, if one exists.
	AvatarURL string `json:"avatar_url,omitempty"`

	// The display name of the user, if one exists.
	DisplayName string `json:"display_name,omitempty"`

	// The user's matrix user ID.
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this results items0
func (o *ResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ResultsItems0) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultsItems0) UnmarshalBinary(b []byte) error {
	var res ResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SearchUserDirectoryBody search user directory body
swagger:model SearchUserDirectoryBody
*/
type SearchUserDirectoryBody struct {

	// The maximum number of results to return. Defaults to 10.
	Limit int64 `json:"limit,omitempty"`

	// The term to search for
	// Required: true
	SearchTerm *string `json:"search_term"`
}

// Validate validates this search user directory body
func (o *SearchUserDirectoryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSearchTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchUserDirectoryBody) validateSearchTerm(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"search_term", "body", o.SearchTerm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchUserDirectoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchUserDirectoryBody) UnmarshalBinary(b []byte) error {
	var res SearchUserDirectoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SearchUserDirectoryOKBody search user directory o k body
swagger:model SearchUserDirectoryOKBody
*/
type SearchUserDirectoryOKBody struct {

	// Indicates if the result list has been truncated by the limit.
	// Required: true
	Limited *bool `json:"limited"`

	// Ordered by rank and then whether or not profile info is available.
	// Required: true
	Results []*ResultsItems0 `json:"results"`
}

// Validate validates this search user directory o k body
func (o *SearchUserDirectoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLimited(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchUserDirectoryOKBody) validateLimited(formats strfmt.Registry) error {

	if err := validate.Required("searchUserDirectoryOK"+"."+"limited", "body", o.Limited); err != nil {
		return err
	}

	return nil
}

func (o *SearchUserDirectoryOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("searchUserDirectoryOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchUserDirectoryOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchUserDirectoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchUserDirectoryOKBody) UnmarshalBinary(b []byte) error {
	var res SearchUserDirectoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SearchUserDirectoryTooManyRequestsBody The rate limit was reached for this request
swagger:model SearchUserDirectoryTooManyRequestsBody
*/
type SearchUserDirectoryTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this search user directory too many requests body
func (o *SearchUserDirectoryTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchUserDirectoryTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("searchUserDirectoryTooManyRequests"+"."+"errcode", "body", o.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchUserDirectoryTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchUserDirectoryTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res SearchUserDirectoryTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
