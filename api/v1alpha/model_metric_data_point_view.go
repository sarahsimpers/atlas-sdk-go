/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the MetricDataPointView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDataPointView{}

// MetricDataPointView value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the `value` parameter returns `null`.
type MetricDataPointView struct {
	// Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Value that comprises this data point.
	Value *float32 `json:"value,omitempty"`
}

// NewMetricDataPointView instantiates a new MetricDataPointView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDataPointView() *MetricDataPointView {
	this := MetricDataPointView{}
	return &this
}

// NewMetricDataPointViewWithDefaults instantiates a new MetricDataPointView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDataPointViewWithDefaults() *MetricDataPointView {
	this := MetricDataPointView{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MetricDataPointView) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointView) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetricDataPointView) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MetricDataPointView) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricDataPointView) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointView) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricDataPointView) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *MetricDataPointView) SetValue(v float32) {
	o.Value = &v
}

func (o MetricDataPointView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricDataPointView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: timestamp is readOnly
	// skip: value is readOnly
	return toSerialize, nil
}

type NullableMetricDataPointView struct {
	value *MetricDataPointView
	isSet bool
}

func (v NullableMetricDataPointView) Get() *MetricDataPointView {
	return v.value
}

func (v *NullableMetricDataPointView) Set(val *MetricDataPointView) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDataPointView) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDataPointView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDataPointView(val *MetricDataPointView) *NullableMetricDataPointView {
	return &NullableMetricDataPointView{value: val, isSet: true}
}

func (v NullableMetricDataPointView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDataPointView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

