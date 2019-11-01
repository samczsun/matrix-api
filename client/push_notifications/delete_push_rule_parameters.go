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

// NewDeletePushRuleParams creates a new DeletePushRuleParams object
// with the default values initialized.
func NewDeletePushRuleParams() *DeletePushRuleParams {
	var ()
	return &DeletePushRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePushRuleParamsWithTimeout creates a new DeletePushRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePushRuleParamsWithTimeout(timeout time.Duration) *DeletePushRuleParams {
	var ()
	return &DeletePushRuleParams{

		timeout: timeout,
	}
}

// NewDeletePushRuleParamsWithContext creates a new DeletePushRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePushRuleParamsWithContext(ctx context.Context) *DeletePushRuleParams {
	var ()
	return &DeletePushRuleParams{

		Context: ctx,
	}
}

// NewDeletePushRuleParamsWithHTTPClient creates a new DeletePushRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePushRuleParamsWithHTTPClient(client *http.Client) *DeletePushRuleParams {
	var ()
	return &DeletePushRuleParams{
		HTTPClient: client,
	}
}

/*DeletePushRuleParams contains all the parameters to send to the API endpoint
for the delete push rule operation typically these are written to a http.Request
*/
type DeletePushRuleParams struct {

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

// WithTimeout adds the timeout to the delete push rule params
func (o *DeletePushRuleParams) WithTimeout(timeout time.Duration) *DeletePushRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete push rule params
func (o *DeletePushRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete push rule params
func (o *DeletePushRuleParams) WithContext(ctx context.Context) *DeletePushRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete push rule params
func (o *DeletePushRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete push rule params
func (o *DeletePushRuleParams) WithHTTPClient(client *http.Client) *DeletePushRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete push rule params
func (o *DeletePushRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKind adds the kind to the delete push rule params
func (o *DeletePushRuleParams) WithKind(kind string) *DeletePushRuleParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the delete push rule params
func (o *DeletePushRuleParams) SetKind(kind string) {
	o.Kind = kind
}

// WithRuleID adds the ruleID to the delete push rule params
func (o *DeletePushRuleParams) WithRuleID(ruleID string) *DeletePushRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete push rule params
func (o *DeletePushRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithScope adds the scope to the delete push rule params
func (o *DeletePushRuleParams) WithScope(scope string) *DeletePushRuleParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete push rule params
func (o *DeletePushRuleParams) SetScope(scope string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePushRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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