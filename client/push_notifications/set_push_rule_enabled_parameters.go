// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetPushRuleEnabledParams creates a new SetPushRuleEnabledParams object
// with the default values initialized.
func NewSetPushRuleEnabledParams() *SetPushRuleEnabledParams {
	var ()
	return &SetPushRuleEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPushRuleEnabledParamsWithTimeout creates a new SetPushRuleEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPushRuleEnabledParamsWithTimeout(timeout time.Duration) *SetPushRuleEnabledParams {
	var ()
	return &SetPushRuleEnabledParams{

		timeout: timeout,
	}
}

// NewSetPushRuleEnabledParamsWithContext creates a new SetPushRuleEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPushRuleEnabledParamsWithContext(ctx context.Context) *SetPushRuleEnabledParams {
	var ()
	return &SetPushRuleEnabledParams{

		Context: ctx,
	}
}

// NewSetPushRuleEnabledParamsWithHTTPClient creates a new SetPushRuleEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPushRuleEnabledParamsWithHTTPClient(client *http.Client) *SetPushRuleEnabledParams {
	var ()
	return &SetPushRuleEnabledParams{
		HTTPClient: client,
	}
}

/*SetPushRuleEnabledParams contains all the parameters to send to the API endpoint
for the set push rule enabled operation typically these are written to a http.Request
*/
type SetPushRuleEnabledParams struct {

	/*Body
	  Whether the push rule is enabled or not.


	*/
	Body SetPushRuleEnabledBody
	/*Kind
	  The kind of rule


	*/
	Kind string
	/*RuleID
	  The identifier for the rule.


	*/
	RuleID string
	/*Scope
	  ``global`` to specify global rules.

	*/
	Scope string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithTimeout(timeout time.Duration) *SetPushRuleEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithContext(ctx context.Context) *SetPushRuleEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithHTTPClient(client *http.Client) *SetPushRuleEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithBody(body SetPushRuleEnabledBody) *SetPushRuleEnabledParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetBody(body SetPushRuleEnabledBody) {
	o.Body = body
}

// WithKind adds the kind to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithKind(kind string) *SetPushRuleEnabledParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetKind(kind string) {
	o.Kind = kind
}

// WithRuleID adds the ruleID to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithRuleID(ruleID string) *SetPushRuleEnabledParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithScope adds the scope to the set push rule enabled params
func (o *SetPushRuleEnabledParams) WithScope(scope string) *SetPushRuleEnabledParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the set push rule enabled params
func (o *SetPushRuleEnabledParams) SetScope(scope string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *SetPushRuleEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	// path param scope
	if err := r.SetPathParam("scope", o.Scope); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
