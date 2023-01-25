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

// DailyScheduleView struct for DailyScheduleView
type DailyScheduleView struct {
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *int32 `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *int32 `json:"endMinute,omitempty"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *int32 `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *int32 `json:"startMinute,omitempty"`
}

// NewDailyScheduleView instantiates a new DailyScheduleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyScheduleView() *DailyScheduleView {
	this := DailyScheduleView{}
	return &this
}

// NewDailyScheduleViewWithDefaults instantiates a new DailyScheduleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyScheduleViewWithDefaults() *DailyScheduleView {
	this := DailyScheduleView{}
	return &this
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *DailyScheduleView) GetEndHour() int32 {
	if o == nil || o.EndHour == nil {
		var ret int32
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyScheduleView) GetEndHourOk() (*int32, bool) {
	if o == nil || o.EndHour == nil {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *DailyScheduleView) HasEndHour() bool {
	if o != nil && o.EndHour != nil {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int32 and assigns it to the EndHour field.
func (o *DailyScheduleView) SetEndHour(v int32) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *DailyScheduleView) GetEndMinute() int32 {
	if o == nil || o.EndMinute == nil {
		var ret int32
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyScheduleView) GetEndMinuteOk() (*int32, bool) {
	if o == nil || o.EndMinute == nil {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *DailyScheduleView) HasEndMinute() bool {
	if o != nil && o.EndMinute != nil {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int32 and assigns it to the EndMinute field.
func (o *DailyScheduleView) SetEndMinute(v int32) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *DailyScheduleView) GetStartHour() int32 {
	if o == nil || o.StartHour == nil {
		var ret int32
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyScheduleView) GetStartHourOk() (*int32, bool) {
	if o == nil || o.StartHour == nil {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *DailyScheduleView) HasStartHour() bool {
	if o != nil && o.StartHour != nil {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int32 and assigns it to the StartHour field.
func (o *DailyScheduleView) SetStartHour(v int32) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *DailyScheduleView) GetStartMinute() int32 {
	if o == nil || o.StartMinute == nil {
		var ret int32
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyScheduleView) GetStartMinuteOk() (*int32, bool) {
	if o == nil || o.StartMinute == nil {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *DailyScheduleView) HasStartMinute() bool {
	if o != nil && o.StartMinute != nil {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int32 and assigns it to the StartMinute field.
func (o *DailyScheduleView) SetStartMinute(v int32) {
	o.StartMinute = &v
}

func (o DailyScheduleView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndHour != nil {
		toSerialize["endHour"] = o.EndHour
	}
	if o.EndMinute != nil {
		toSerialize["endMinute"] = o.EndMinute
	}
	if o.StartHour != nil {
		toSerialize["startHour"] = o.StartHour
	}
	if o.StartMinute != nil {
		toSerialize["startMinute"] = o.StartMinute
	}
	return json.Marshal(toSerialize)
}

type NullableDailyScheduleView struct {
	value *DailyScheduleView
	isSet bool
}

func (v NullableDailyScheduleView) Get() *DailyScheduleView {
	return v.value
}

func (v *NullableDailyScheduleView) Set(val *DailyScheduleView) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyScheduleView) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyScheduleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyScheduleView(val *DailyScheduleView) *NullableDailyScheduleView {
	return &NullableDailyScheduleView{value: val, isSet: true}
}

func (v NullableDailyScheduleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyScheduleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


