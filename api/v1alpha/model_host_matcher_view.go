/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// HostMatcherView Rules to apply when comparing an host against this alert configuration.
type HostMatcherView struct {
	FieldName *HostMatcherField `json:"fieldName,omitempty"`
	// Comparison operator to apply when checking the current metric value against **matcher[n].value**.
	Operator *string `json:"operator,omitempty"`
	Value *MatcherHostType `json:"value,omitempty"`
}

// NewHostMatcherView instantiates a new HostMatcherView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMatcherView() *HostMatcherView {
	this := HostMatcherView{}
	return &this
}

// NewHostMatcherViewWithDefaults instantiates a new HostMatcherView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMatcherViewWithDefaults() *HostMatcherView {
	this := HostMatcherView{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *HostMatcherView) GetFieldName() HostMatcherField {
	if o == nil || o.FieldName == nil {
		var ret HostMatcherField
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMatcherView) GetFieldNameOk() (*HostMatcherField, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *HostMatcherView) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given HostMatcherField and assigns it to the FieldName field.
func (o *HostMatcherView) SetFieldName(v HostMatcherField) {
	o.FieldName = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *HostMatcherView) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMatcherView) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *HostMatcherView) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *HostMatcherView) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HostMatcherView) GetValue() MatcherHostType {
	if o == nil || o.Value == nil {
		var ret MatcherHostType
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMatcherView) GetValueOk() (*MatcherHostType, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HostMatcherView) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given MatcherHostType and assigns it to the Value field.
func (o *HostMatcherView) SetValue(v MatcherHostType) {
	o.Value = &v
}

func (o HostMatcherView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHostMatcherView struct {
	value *HostMatcherView
	isSet bool
}

func (v NullableHostMatcherView) Get() *HostMatcherView {
	return v.value
}

func (v *NullableHostMatcherView) Set(val *HostMatcherView) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMatcherView) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMatcherView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMatcherView(val *HostMatcherView) *NullableHostMatcherView {
	return &NullableHostMatcherView{value: val, isSet: true}
}

func (v NullableHostMatcherView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMatcherView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


