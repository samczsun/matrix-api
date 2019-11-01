// Code generated by go-swagger; DO NOT EDIT.

package server_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server administration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server administration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetVersions gets the versions of the specification supported by the server

Gets the versions of the specification supported by the server.

Values will take the form ``rX.Y.Z``.

Only the latest ``Z`` value will be reported for each supported ``X.Y`` value.
i.e. if the server implements ``r0.0.0``, ``r0.0.1``, and ``r1.2.0``, it will report ``r0.0.1`` and ``r1.2.0``.

The server may additionally advertise experimental features it supports
through ``unstable_features``. These features should be namespaced and
may optionally include version information within their name if desired.
Features listed here are not for optionally toggling parts of the Matrix
specification and should only be used to advertise support for a feature
which has not yet landed in the spec. For example, a feature currently
undergoing the proposal process may appear here and eventually be taken
off this list once the feature lands in the spec and the server deems it
reasonable to do so. Servers may wish to keep advertising features here
after they've been released into the spec to give clients a chance to
upgrade appropriately. Additionally, clients should avoid using unstable
features in their stable releases.
*/
func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersions",
		Method:             "GET",
		PathPattern:        "/_matrix/client/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWellknown gets matrix server discovery information about the domain

Gets discovery information about the domain. The file may include
additional keys, which MUST follow the Java package naming convention,
e.g. ``com.example.myapp.property``. This ensures property names are
suitably namespaced for each application and reduces the risk of
clashes.

Note that this endpoint is not necessarily handled by the homeserver,
but by another webserver, to be used for discovering the homeserver URL.
*/
func (a *Client) GetWellknown(params *GetWellknownParams) (*GetWellknownOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWellknownParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWellknown",
		Method:             "GET",
		PathPattern:        "/.well-known/matrix/client",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWellknownReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWellknownOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWellknown: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWhoIs gets information about a particular user

Gets information about a particular user.

This API may be restricted to only be called by the user being looked
up, or by a server admin. Server-local administrator privileges are not
specified in this document.
*/
func (a *Client) GetWhoIs(params *GetWhoIsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWhoIsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWhoIsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWhoIs",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/admin/whois/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWhoIsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWhoIsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWhoIs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}