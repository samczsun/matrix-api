// Code generated by go-swagger; DO NOT EDIT.

package room_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new room discovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for room discovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPublicRooms lists the public rooms on the server

Lists the public rooms on the server.

This API returns paginated responses. The rooms are ordered by the number
of joined members, with the largest rooms first.
*/
func (a *Client) GetPublicRooms(params *GetPublicRoomsParams) (*GetPublicRoomsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicRoomsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicRooms",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/publicRooms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublicRoomsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicRoomsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPublicRooms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryPublicRooms lists the public rooms on the server with optional filter

Lists the public rooms on the server, with optional filter.

This API returns paginated responses. The rooms are ordered by the number
of joined members, with the largest rooms first.
*/
func (a *Client) QueryPublicRooms(params *QueryPublicRoomsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPublicRoomsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryPublicRoomsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryPublicRooms",
		Method:             "POST",
		PathPattern:        "/_matrix/client/unstable/publicRooms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryPublicRoomsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryPublicRoomsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryPublicRooms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
