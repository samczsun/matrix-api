// Code generated by go-swagger; DO NOT EDIT.

package push_notifications

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

// GetPushRuleReader is a Reader for the GetPushRule structure.
type GetPushRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPushRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPushRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPushRuleOK creates a GetPushRuleOK with default headers values
func NewGetPushRuleOK() *GetPushRuleOK {
	return &GetPushRuleOK{}
}

/*GetPushRuleOK handles this case with default header values.

The specific push rule. This will also include keys specific to the
rule itself such as the rule's ``actions`` and ``conditions`` if set.
*/
type GetPushRuleOK struct {
	Payload *GetPushRuleOKBody
}

func (o *GetPushRuleOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}][%d] getPushRuleOK  %+v", 200, o.Payload)
}

func (o *GetPushRuleOK) GetPayload() *GetPushRuleOKBody {
	return o.Payload
}

func (o *GetPushRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPushRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConditionsItems0 conditions items0
swagger:model ConditionsItems0
*/
type ConditionsItems0 struct {

	// Required for ``room_member_count`` conditions. A decimal integer
	// optionally prefixed by one of, ==, <, >, >= or <=. A prefix of < matches
	// rooms where the member count is strictly less than the given number and
	// so forth. If no prefix is present, this parameter defaults to ==.
	Is string `json:"is,omitempty"`

	// Required for ``event_match`` conditions. The dot-separated field of the
	// event to match.
	//
	// Required for ``sender_notification_permission`` conditions. The field in
	// the power level event the user needs a minimum power level for. Fields
	// must be specified under the ``notifications`` property in the power level
	// event's ``content``.
	Key string `json:"key,omitempty"`

	// The kind of condition to apply. See `conditions <#conditions>`_ for
	// more information on the allowed kinds and how they work.
	// Required: true
	Kind *string `json:"kind"`

	// Required for ``event_match`` conditions. The glob-style pattern to
	// match against. Patterns with no special glob characters should be
	// treated as having asterisks prepended and appended when testing the
	// condition.
	Pattern string `json:"pattern,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ConditionsItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Is string `json:"is,omitempty"`

		Key string `json:"key,omitempty"`

		Kind *string `json:"kind"`

		Pattern string `json:"pattern,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.Is = dataAO0.Is

	o.Key = dataAO0.Key

	o.Kind = dataAO0.Kind

	o.Pattern = dataAO0.Pattern

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ConditionsItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Is string `json:"is,omitempty"`

		Key string `json:"key,omitempty"`

		Kind *string `json:"kind"`

		Pattern string `json:"pattern,omitempty"`
	}

	dataAO0.Is = o.Is

	dataAO0.Key = o.Key

	dataAO0.Kind = o.Kind

	dataAO0.Pattern = o.Pattern

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this conditions items0
func (o *ConditionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConditionsItems0) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConditionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConditionsItems0) UnmarshalBinary(b []byte) error {
	var res ConditionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPushRuleOKBody The push rule.
swagger:model GetPushRuleOKBody
*/
type GetPushRuleOKBody struct {

	// The actions to perform when this rule is matched.
	// Required: true
	Actions []interface{} `json:"actions"`

	// The conditions that must hold true for an event in order for a rule to be
	// applied to an event. A rule with no conditions always matches. Only
	// applicable to ``underride`` and ``override`` rules.
	Conditions []*ConditionsItems0 `json:"conditions"`

	// Whether this is a default rule, or has been set explicitly.
	// Required: true
	Default *bool `json:"default"`

	// Whether the push rule is enabled or not.
	// Required: true
	Enabled *bool `json:"enabled"`

	// The glob-style pattern to match against.  Only applicable to ``content``
	// rules.
	Pattern string `json:"pattern,omitempty"`

	// The ID of this rule.
	// Required: true
	RuleID *string `json:"rule_id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPushRuleOKBody) UnmarshalJSON(raw []byte) error {
	// GetPushRuleOKBodyAO0
	var dataGetPushRuleOKBodyAO0 struct {
		Actions []interface{} `json:"actions"`

		Conditions []*ConditionsItems0 `json:"conditions"`

		Default *bool `json:"default"`

		Enabled *bool `json:"enabled"`

		Pattern string `json:"pattern,omitempty"`

		RuleID *string `json:"rule_id"`
	}
	if err := swag.ReadJSON(raw, &dataGetPushRuleOKBodyAO0); err != nil {
		return err
	}

	o.Actions = dataGetPushRuleOKBodyAO0.Actions

	o.Conditions = dataGetPushRuleOKBodyAO0.Conditions

	o.Default = dataGetPushRuleOKBodyAO0.Default

	o.Enabled = dataGetPushRuleOKBodyAO0.Enabled

	o.Pattern = dataGetPushRuleOKBodyAO0.Pattern

	o.RuleID = dataGetPushRuleOKBodyAO0.RuleID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPushRuleOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataGetPushRuleOKBodyAO0 struct {
		Actions []interface{} `json:"actions"`

		Conditions []*ConditionsItems0 `json:"conditions"`

		Default *bool `json:"default"`

		Enabled *bool `json:"enabled"`

		Pattern string `json:"pattern,omitempty"`

		RuleID *string `json:"rule_id"`
	}

	dataGetPushRuleOKBodyAO0.Actions = o.Actions

	dataGetPushRuleOKBodyAO0.Conditions = o.Conditions

	dataGetPushRuleOKBodyAO0.Default = o.Default

	dataGetPushRuleOKBodyAO0.Enabled = o.Enabled

	dataGetPushRuleOKBodyAO0.Pattern = o.Pattern

	dataGetPushRuleOKBodyAO0.RuleID = o.RuleID

	jsonDataGetPushRuleOKBodyAO0, errGetPushRuleOKBodyAO0 := swag.WriteJSON(dataGetPushRuleOKBodyAO0)
	if errGetPushRuleOKBodyAO0 != nil {
		return nil, errGetPushRuleOKBodyAO0
	}
	_parts = append(_parts, jsonDataGetPushRuleOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get push rule o k body
func (o *GetPushRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPushRuleOKBody) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("getPushRuleOK"+"."+"actions", "body", o.Actions); err != nil {
		return err
	}

	return nil
}

func (o *GetPushRuleOKBody) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(o.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(o.Conditions); i++ {
		if swag.IsZero(o.Conditions[i]) { // not required
			continue
		}

		if o.Conditions[i] != nil {
			if err := o.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPushRuleOK" + "." + "conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPushRuleOKBody) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("getPushRuleOK"+"."+"default", "body", o.Default); err != nil {
		return err
	}

	return nil
}

func (o *GetPushRuleOKBody) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("getPushRuleOK"+"."+"enabled", "body", o.Enabled); err != nil {
		return err
	}

	return nil
}

func (o *GetPushRuleOKBody) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("getPushRuleOK"+"."+"rule_id", "body", o.RuleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPushRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPushRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetPushRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}