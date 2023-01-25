/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RawMetricValueView Measurement of the **metricName** recorded at the time of the event.
type RawMetricValueView struct {
	// Amount of the **metricName** recorded at the time of the event. This value triggered the alert.
	Number *float64 `json:"number,omitempty"`
	Units *RawMetricUnits `json:"units,omitempty"`
}

// NewRawMetricValueView instantiates a new RawMetricValueView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawMetricValueView() *RawMetricValueView {
	this := RawMetricValueView{}
	return &this
}

// NewRawMetricValueViewWithDefaults instantiates a new RawMetricValueView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawMetricValueViewWithDefaults() *RawMetricValueView {
	this := RawMetricValueView{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *RawMetricValueView) GetNumber() float64 {
	if o == nil || o.Number == nil {
		var ret float64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricValueView) GetNumberOk() (*float64, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *RawMetricValueView) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float64 and assigns it to the Number field.
func (o *RawMetricValueView) SetNumber(v float64) {
	o.Number = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *RawMetricValueView) GetUnits() RawMetricUnits {
	if o == nil || o.Units == nil {
		var ret RawMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricValueView) GetUnitsOk() (*RawMetricUnits, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *RawMetricValueView) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given RawMetricUnits and assigns it to the Units field.
func (o *RawMetricValueView) SetUnits(v RawMetricUnits) {
	o.Units = &v
}

func (o RawMetricValueView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableRawMetricValueView struct {
	value *RawMetricValueView
	isSet bool
}

func (v NullableRawMetricValueView) Get() *RawMetricValueView {
	return v.value
}

func (v *NullableRawMetricValueView) Set(val *RawMetricValueView) {
	v.value = val
	v.isSet = true
}

func (v NullableRawMetricValueView) IsSet() bool {
	return v.isSet
}

func (v *NullableRawMetricValueView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawMetricValueView(val *RawMetricValueView) *NullableRawMetricValueView {
	return &NullableRawMetricValueView{value: val, isSet: true}
}

func (v NullableRawMetricValueView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawMetricValueView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


