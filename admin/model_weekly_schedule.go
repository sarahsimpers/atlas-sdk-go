// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the WeeklySchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeeklySchedule{}

// WeeklySchedule struct for WeeklySchedule
type WeeklySchedule struct {
	// Day of the week when the scheduled archive starts. The week starts with Monday (`1`) and ends with Sunday (`7`).
	DayOfWeek *int `json:"dayOfWeek,omitempty"`
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *int `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *int `json:"endMinute,omitempty"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *int `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *int   `json:"startMinute,omitempty"`
	Type        string `json:"type"`
}

// NewWeeklySchedule instantiates a new WeeklySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklySchedule(type_ string) *WeeklySchedule {
	this := WeeklySchedule{}
	this.Type = type_
	return &this
}

// NewWeeklyScheduleWithDefaults instantiates a new WeeklySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyScheduleWithDefaults() *WeeklySchedule {
	this := WeeklySchedule{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *WeeklySchedule) GetDayOfWeek() int {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret int
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetDayOfWeekOk() (*int, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *WeeklySchedule) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given int and assigns it to the DayOfWeek field.
func (o *WeeklySchedule) SetDayOfWeek(v int) {
	o.DayOfWeek = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *WeeklySchedule) GetEndHour() int {
	if o == nil || IsNil(o.EndHour) {
		var ret int
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetEndHourOk() (*int, bool) {
	if o == nil || IsNil(o.EndHour) {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *WeeklySchedule) HasEndHour() bool {
	if o != nil && !IsNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int and assigns it to the EndHour field.
func (o *WeeklySchedule) SetEndHour(v int) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *WeeklySchedule) GetEndMinute() int {
	if o == nil || IsNil(o.EndMinute) {
		var ret int
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetEndMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.EndMinute) {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *WeeklySchedule) HasEndMinute() bool {
	if o != nil && !IsNil(o.EndMinute) {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int and assigns it to the EndMinute field.
func (o *WeeklySchedule) SetEndMinute(v int) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *WeeklySchedule) GetStartHour() int {
	if o == nil || IsNil(o.StartHour) {
		var ret int
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetStartHourOk() (*int, bool) {
	if o == nil || IsNil(o.StartHour) {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *WeeklySchedule) HasStartHour() bool {
	if o != nil && !IsNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int and assigns it to the StartHour field.
func (o *WeeklySchedule) SetStartHour(v int) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *WeeklySchedule) GetStartMinute() int {
	if o == nil || IsNil(o.StartMinute) {
		var ret int
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetStartMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.StartMinute) {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *WeeklySchedule) HasStartMinute() bool {
	if o != nil && !IsNil(o.StartMinute) {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int and assigns it to the StartMinute field.
func (o *WeeklySchedule) SetStartMinute(v int) {
	o.StartMinute = &v
}

// GetType returns the Type field value
func (o *WeeklySchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WeeklySchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WeeklySchedule) SetType(v string) {
	o.Type = v
}

func (o WeeklySchedule) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o WeeklySchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !IsNil(o.EndHour) {
		toSerialize["endHour"] = o.EndHour
	}
	if !IsNil(o.EndMinute) {
		toSerialize["endMinute"] = o.EndMinute
	}
	if !IsNil(o.StartHour) {
		toSerialize["startHour"] = o.StartHour
	}
	if !IsNil(o.StartMinute) {
		toSerialize["startMinute"] = o.StartMinute
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableWeeklySchedule struct {
	value *WeeklySchedule
	isSet bool
}

func (v NullableWeeklySchedule) Get() *WeeklySchedule {
	return v.value
}

func (v *NullableWeeklySchedule) Set(val *WeeklySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklySchedule(val *WeeklySchedule) *NullableWeeklySchedule {
	return &NullableWeeklySchedule{value: val, isSet: true}
}

func (v NullableWeeklySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
