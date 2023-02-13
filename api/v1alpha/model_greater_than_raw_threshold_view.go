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

// GreaterThanRawThresholdView A Limit that triggers an alert when greater than a number.
type GreaterThanRawThresholdView struct {
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *int32 `json:"threshold,omitempty"`
	Units *RawMetricUnits `json:"units,omitempty"`
}

// NewGreaterThanRawThresholdView instantiates a new GreaterThanRawThresholdView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterThanRawThresholdView() *GreaterThanRawThresholdView {
	this := GreaterThanRawThresholdView{}
	var units RawMetricUnits = RAWMETRICUNITS_RAW
	this.Units = &units
	return &this
}

// NewGreaterThanRawThresholdViewWithDefaults instantiates a new GreaterThanRawThresholdView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterThanRawThresholdViewWithDefaults() *GreaterThanRawThresholdView {
	this := GreaterThanRawThresholdView{}
	var units RawMetricUnits = RAWMETRICUNITS_RAW
	this.Units = &units
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *GreaterThanRawThresholdView) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThresholdView) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *GreaterThanRawThresholdView) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *GreaterThanRawThresholdView) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *GreaterThanRawThresholdView) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThresholdView) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *GreaterThanRawThresholdView) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *GreaterThanRawThresholdView) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GreaterThanRawThresholdView) GetUnits() RawMetricUnits {
	if o == nil || o.Units == nil {
		var ret RawMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThresholdView) GetUnitsOk() (*RawMetricUnits, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GreaterThanRawThresholdView) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given RawMetricUnits and assigns it to the Units field.
func (o *GreaterThanRawThresholdView) SetUnits(v RawMetricUnits) {
	o.Units = &v
}

func (o GreaterThanRawThresholdView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableGreaterThanRawThresholdView struct {
	value *GreaterThanRawThresholdView
	isSet bool
}

func (v NullableGreaterThanRawThresholdView) Get() *GreaterThanRawThresholdView {
	return v.value
}

func (v *NullableGreaterThanRawThresholdView) Set(val *GreaterThanRawThresholdView) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterThanRawThresholdView) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterThanRawThresholdView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterThanRawThresholdView(val *GreaterThanRawThresholdView) *NullableGreaterThanRawThresholdView {
	return &NullableGreaterThanRawThresholdView{value: val, isSet: true}
}

func (v NullableGreaterThanRawThresholdView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterThanRawThresholdView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


