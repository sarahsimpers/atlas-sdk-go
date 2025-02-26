// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the MetricDataPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDataPoint{}

// MetricDataPoint value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the `value` parameter returns `null`.
type MetricDataPoint struct {
	// Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Value that comprises this data point.
	Value *float32 `json:"value,omitempty"`
}

// NewMetricDataPoint instantiates a new MetricDataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDataPoint() *MetricDataPoint {
	this := MetricDataPoint{}
	return &this
}

// NewMetricDataPointWithDefaults instantiates a new MetricDataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDataPointWithDefaults() *MetricDataPoint {
	this := MetricDataPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MetricDataPoint) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetricDataPoint) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MetricDataPoint) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricDataPoint) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPoint) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricDataPoint) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *MetricDataPoint) SetValue(v float32) {
	o.Value = &v
}

func (o MetricDataPoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o MetricDataPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableMetricDataPoint struct {
	value *MetricDataPoint
	isSet bool
}

func (v NullableMetricDataPoint) Get() *MetricDataPoint {
	return v.value
}

func (v *NullableMetricDataPoint) Set(val *MetricDataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDataPoint(val *MetricDataPoint) *NullableMetricDataPoint {
	return &NullableMetricDataPoint{value: val, isSet: true}
}

func (v NullableMetricDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
