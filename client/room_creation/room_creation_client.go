// Code generated by go-swagger; DO NOT EDIT.

package room_creation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new room creation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for room creation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRoom creates a new room

Create a new room with various configuration options.

The server MUST apply the normal state resolution rules when creating
the new room, including checking power levels for each event. It MUST
apply the events implied by the request in the following order:

0. A default ``m.room.power_levels`` event, giving the room creator
   (and not other members) permission to send state events. Overridden
   by the ``power_level_content_override`` parameter.

1. Events set by the ``preset``. Currently these are the ``m.room.join_rules``,
   ``m.room.history_visibility``, and ``m.room.guest_access`` state events.

2. Events listed in ``initial_state``, in the order that they are
   listed.

3. Events implied by ``name`` and ``topic`` (``m.room.name`` and ``m.room.topic``
   state events).

4. Invite events implied by ``invite`` and ``invite_3pid`` (``m.room.member`` with
   ``membership: invite`` and ``m.room.third_party_invite``).

The available presets do the following with respect to room state:

========================  ==============  ======================  ================  =========
         Preset           ``join_rules``  ``history_visibility``  ``guest_access``  Other
========================  ==============  ======================  ================  =========
``private_chat``          ``invite``      ``shared``              ``can_join``
``trusted_private_chat``  ``invite``      ``shared``              ``can_join``      All invitees are given the same power level as the room creator.
``public_chat``           ``public``      ``shared``              ``forbidden``
========================  ==============  ======================  ================  =========

The server will create a ``m.room.create`` event in the room with the
requesting user as the creator, alongside other keys provided in the
``creation_content``.
*/
func (a *Client) CreateRoom(params *CreateRoomParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRoomOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoomParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRoom",
		Method:             "POST",
		PathPattern:        "/_matrix/client/unstable/createRoom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRoomReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRoomOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRoom: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
