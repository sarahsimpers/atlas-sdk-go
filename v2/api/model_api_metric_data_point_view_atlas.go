/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// ApiMetricDataPointViewAtlas value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the `value` parameter returns `null`.
type ApiMetricDataPointViewAtlas struct {
	// Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Value that comprises this data point.
	Value *float32 `json:"value,omitempty"`
}

// NewApiMetricDataPointViewAtlas instantiates a new ApiMetricDataPointViewAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMetricDataPointViewAtlas() *ApiMetricDataPointViewAtlas {
	this := ApiMetricDataPointViewAtlas{}
	return &this
}

// NewApiMetricDataPointViewAtlasWithDefaults instantiates a new ApiMetricDataPointViewAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMetricDataPointViewAtlasWithDefaults() *ApiMetricDataPointViewAtlas {
	this := ApiMetricDataPointViewAtlas{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiMetricDataPointViewAtlas) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricDataPointViewAtlas) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiMetricDataPointViewAtlas) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ApiMetricDataPointViewAtlas) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiMetricDataPointViewAtlas) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricDataPointViewAtlas) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiMetricDataPointViewAtlas) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ApiMetricDataPointViewAtlas) SetValue(v float32) {
	o.Value = &v
}

func (o ApiMetricDataPointViewAtlas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiMetricDataPointViewAtlas struct {
	value *ApiMetricDataPointViewAtlas
	isSet bool
}

func (v NullableApiMetricDataPointViewAtlas) Get() *ApiMetricDataPointViewAtlas {
	return v.value
}

func (v *NullableApiMetricDataPointViewAtlas) Set(val *ApiMetricDataPointViewAtlas) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMetricDataPointViewAtlas) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMetricDataPointViewAtlas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMetricDataPointViewAtlas(val *ApiMetricDataPointViewAtlas) *NullableApiMetricDataPointViewAtlas {
	return &NullableApiMetricDataPointViewAtlas{value: val, isSet: true}
}

func (v NullableApiMetricDataPointViewAtlas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMetricDataPointViewAtlas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


