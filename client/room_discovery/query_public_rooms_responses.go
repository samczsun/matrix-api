// Code generated by go-swagger; DO NOT EDIT.

package room_discovery

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

// QueryPublicRoomsReader is a Reader for the QueryPublicRooms structure.
type QueryPublicRoomsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPublicRoomsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPublicRoomsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryPublicRoomsOK creates a QueryPublicRoomsOK with default headers values
func NewQueryPublicRoomsOK() *QueryPublicRoomsOK {
	return &QueryPublicRoomsOK{}
}

/*QueryPublicRoomsOK handles this case with default header values.

A list of the rooms on the server.
*/
type QueryPublicRoomsOK struct {
	Payload *QueryPublicRoomsOKBody
}

func (o *QueryPublicRoomsOK) Error() string {
	return fmt.Sprintf("[POST /_matrix/client/unstable/publicRooms][%d] queryPublicRoomsOK  %+v", 200, o.Payload)
}

func (o *QueryPublicRoomsOK) GetPayload() *QueryPublicRoomsOKBody {
	return o.Payload
}

func (o *QueryPublicRoomsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryPublicRoomsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QueryPublicRoomsBody query public rooms body
swagger:model QueryPublicRoomsBody
*/
type QueryPublicRoomsBody struct {

	// filter
	Filter *QueryPublicRoomsParamsBodyFilter `json:"filter,omitempty"`

	// Whether or not to include all known networks/protocols from
	// application services on the homeserver. Defaults to false.
	IncludeAllNetworks bool `json:"include_all_networks,omitempty"`

	// Limit the number of results returned.
	Limit int64 `json:"limit,omitempty"`

	// A pagination token from a previous request, allowing clients
	// to get the next (or previous) batch of rooms.  The direction
	// of pagination is specified solely by which token is supplied,
	// rather than via an explicit flag.
	Since string `json:"since,omitempty"`

	// The specific third party network/protocol to request from the
	// homeserver. Can only be used if ``include_all_networks`` is false.
	ThirdPartyInstanceID string `json:"third_party_instance_id,omitempty"`
}

// Validate validates this query public rooms body
func (o *QueryPublicRoomsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryPublicRoomsBody) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(o.Filter) { // not required
		return nil
	}

	if o.Filter != nil {
		if err := o.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryPublicRoomsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryPublicRoomsBody) UnmarshalBinary(b []byte) error {
	var res QueryPublicRoomsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryPublicRoomsOKBody A list of the rooms on the server.
swagger:model QueryPublicRoomsOKBody
*/
type QueryPublicRoomsOKBody struct {

	// PublicRoomsChunks
	//
	// A paginated chunk of public rooms.
	// Required: true
	Chunk []*ChunkItems0 `json:"chunk"`

	// A pagination token for the response. The absence of this token
	// means there are no more results to fetch and the client should
	// stop paginating.
	NextBatch string `json:"next_batch,omitempty"`

	// A pagination token that allows fetching previous results. The
	// absence of this token means there are no results before this
	// batch, i.e. this is the first batch.
	PrevBatch string `json:"prev_batch,omitempty"`

	// An estimate on the total number of public rooms, if the
	// server has an estimate.
	TotalRoomCountEstimate int64 `json:"total_room_count_estimate,omitempty"`
}

// Validate validates this query public rooms o k body
func (o *QueryPublicRoomsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChunk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryPublicRoomsOKBody) validateChunk(formats strfmt.Registry) error {

	if err := validate.Required("queryPublicRoomsOK"+"."+"chunk", "body", o.Chunk); err != nil {
		return err
	}

	for i := 0; i < len(o.Chunk); i++ {
		if swag.IsZero(o.Chunk[i]) { // not required
			continue
		}

		if o.Chunk[i] != nil {
			if err := o.Chunk[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryPublicRoomsOK" + "." + "chunk" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryPublicRoomsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryPublicRoomsOKBody) UnmarshalBinary(b []byte) error {
	var res QueryPublicRoomsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QueryPublicRoomsParamsBodyFilter Filter
//
// Filter to apply to the results.
swagger:model QueryPublicRoomsParamsBodyFilter
*/
type QueryPublicRoomsParamsBodyFilter struct {

	// A string to search for in the room metadata, e.g. name,
	// topic, canonical alias etc. (Optional).
	GenericSearchTerm string `json:"generic_search_term,omitempty"`
}

// Validate validates this query public rooms params body filter
func (o *QueryPublicRoomsParamsBodyFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryPublicRoomsParamsBodyFilter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryPublicRoomsParamsBodyFilter) UnmarshalBinary(b []byte) error {
	var res QueryPublicRoomsParamsBodyFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
