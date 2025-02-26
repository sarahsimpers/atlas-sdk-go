// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the GroupNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupNotification{}

// GroupNotification Group notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type GroupNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int `json:"delayMin,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.[n].typeName\" : \"ORG\"` - `\"notifications.[n].typeName\" : \"GROUP\"` - `\"notifications.[n].typeName\" : \"USER\"`
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int `json:"intervalMin,omitempty"`
	// List that contains the one or more [organization](https://dochub.mongodb.org/core/atlas-org-roles) or [project roles](https://dochub.mongodb.org/core/atlas-proj-roles) that receive the configured alert. The resource requires this parameter when `\"notifications.[n].typeName\" : \"GROUP\"` or `\"notifications.[n].typeName\" : \"ORG\"`. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role.
	Roles []string `json:"roles,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.[n].typeName\" : \"ORG\"` - `\"notifications.[n].typeName\" : \"GROUP\"` - `\"notifications.[n].typeName\" : \"USER\"`
	SmsEnabled *bool `json:"smsEnabled,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewGroupNotification instantiates a new GroupNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupNotification(typeName string) *GroupNotification {
	this := GroupNotification{}
	this.TypeName = typeName
	return &this
}

// NewGroupNotificationWithDefaults instantiates a new GroupNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupNotificationWithDefaults() *GroupNotification {
	this := GroupNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *GroupNotification) GetDelayMin() int {
	if o == nil || IsNil(o.DelayMin) {
		var ret int
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetDelayMinOk() (*int, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *GroupNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int and assigns it to the DelayMin field.
func (o *GroupNotification) SetDelayMin(v int) {
	o.DelayMin = &v
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *GroupNotification) GetEmailEnabled() bool {
	if o == nil || IsNil(o.EmailEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEnabled) {
		return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *GroupNotification) HasEmailEnabled() bool {
	if o != nil && !IsNil(o.EmailEnabled) {
		return true
	}

	return false
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *GroupNotification) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *GroupNotification) GetIntervalMin() int {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetIntervalMinOk() (*int, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *GroupNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int and assigns it to the IntervalMin field.
func (o *GroupNotification) SetIntervalMin(v int) {
	o.IntervalMin = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GroupNotification) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupNotification) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupNotification) SetRoles(v []string) {
	o.Roles = v
}

// GetSmsEnabled returns the SmsEnabled field value if set, zero value otherwise.
func (o *GroupNotification) GetSmsEnabled() bool {
	if o == nil || IsNil(o.SmsEnabled) {
		var ret bool
		return ret
	}
	return *o.SmsEnabled
}

// GetSmsEnabledOk returns a tuple with the SmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetSmsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SmsEnabled) {
		return nil, false
	}
	return o.SmsEnabled, true
}

// HasSmsEnabled returns a boolean if a field has been set.
func (o *GroupNotification) HasSmsEnabled() bool {
	if o != nil && !IsNil(o.SmsEnabled) {
		return true
	}

	return false
}

// SetSmsEnabled gets a reference to the given bool and assigns it to the SmsEnabled field.
func (o *GroupNotification) SetSmsEnabled(v bool) {
	o.SmsEnabled = &v
}

// GetTypeName returns the TypeName field value
func (o *GroupNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *GroupNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *GroupNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o GroupNotification) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.EmailEnabled) {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.SmsEnabled) {
		toSerialize["smsEnabled"] = o.SmsEnabled
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableGroupNotification struct {
	value *GroupNotification
	isSet bool
}

func (v NullableGroupNotification) Get() *GroupNotification {
	return v.value
}

func (v *NullableGroupNotification) Set(val *GroupNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupNotification(val *GroupNotification) *NullableGroupNotification {
	return &NullableGroupNotification{value: val, isSet: true}
}

func (v NullableGroupNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
