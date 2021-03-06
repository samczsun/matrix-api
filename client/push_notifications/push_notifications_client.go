// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new push notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for push notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePushRule deletes a push rule

This endpoint removes the push rule defined in the path.
*/
func (a *Client) DeletePushRule(params *DeletePushRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePushRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePushRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePushRule",
		Method:             "DELETE",
		PathPattern:        "/_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePushRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePushRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePushRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNotifications gets a list of events that the user has been notified about

This API is used to paginate through the list of events that the
user has been, or would have been notified about.
*/
func (a *Client) GetNotifications(params *GetNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNotifications",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPushRule retrieves a push rule

Retrieve a single specified push rule.
*/
func (a *Client) GetPushRule(params *GetPushRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetPushRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPushRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPushRule",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPushRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPushRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPushRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPushRules retrieves all push rulesets

Retrieve all push rulesets for this user. Clients can "drill-down" on
the rulesets by suffixing a ``scope`` to this path e.g.
``/pushrules/global/``. This will return a subset of this data under the
specified key e.g. the ``global`` key.
*/
func (a *Client) GetPushRules(params *GetPushRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPushRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPushRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPushRules",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/pushrules/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPushRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPushRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPushRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPushers gets the current pushers for the authenticated user

Gets all currently active pushers for the authenticated user.
*/
func (a *Client) GetPushers(params *GetPushersParams, authInfo runtime.ClientAuthInfoWriter) (*GetPushersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPushersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPushers",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/pushers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPushersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPushersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPushers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostPusher modifies a pusher for this user on the homeserver

This endpoint allows the creation, modification and deletion of `pushers`_
for this user ID. The behaviour of this endpoint varies depending on the
values in the JSON body.
*/
func (a *Client) PostPusher(params *PostPusherParams, authInfo runtime.ClientAuthInfoWriter) (*PostPusherOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPusherParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPusher",
		Method:             "POST",
		PathPattern:        "/_matrix/client/unstable/pushers/set",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPusherReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostPusherOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postPusher: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetPushRule adds or change a push rule

This endpoint allows the creation, modification and deletion of pushers
for this user ID. The behaviour of this endpoint varies depending on the
values in the JSON body.

When creating push rules, they MUST be enabled by default.
*/
func (a *Client) SetPushRule(params *SetPushRuleParams, authInfo runtime.ClientAuthInfoWriter) (*SetPushRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPushRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPushRule",
		Method:             "PUT",
		PathPattern:        "/_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetPushRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetPushRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setPushRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetPushRuleActions sets the actions for a push rule

This endpoint allows clients to change the actions of a push rule.
This can be used to change the actions of builtin rules.
*/
func (a *Client) SetPushRuleActions(params *SetPushRuleActionsParams, authInfo runtime.ClientAuthInfoWriter) (*SetPushRuleActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPushRuleActionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPushRuleActions",
		Method:             "PUT",
		PathPattern:        "/_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetPushRuleActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetPushRuleActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setPushRuleActions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetPushRuleEnabled enables or disable a push rule

This endpoint allows clients to enable or disable the specified push rule.
*/
func (a *Client) SetPushRuleEnabled(params *SetPushRuleEnabledParams, authInfo runtime.ClientAuthInfoWriter) (*SetPushRuleEnabledOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPushRuleEnabledParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setPushRuleEnabled",
		Method:             "PUT",
		PathPattern:        "/_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/enabled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetPushRuleEnabledReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetPushRuleEnabledOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setPushRuleEnabled: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
