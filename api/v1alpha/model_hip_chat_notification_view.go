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

// HipChatNotificationView HipChat notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type HipChatNotificationView struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when `\"notifications.[n].typeName\" : \"HIP_CHAT\"`\". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	NotificationToken *string `json:"notificationToken,omitempty"`
	// HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"HIP_CHAT\"`\".
	RoomName *string `json:"roomName,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewHipChatNotificationView instantiates a new HipChatNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHipChatNotificationView() *HipChatNotificationView {
	this := HipChatNotificationView{}
	return &this
}

// NewHipChatNotificationViewWithDefaults instantiates a new HipChatNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHipChatNotificationViewWithDefaults() *HipChatNotificationView {
	this := HipChatNotificationView{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *HipChatNotificationView) GetDelayMin() int32 {
	if o == nil || o.DelayMin == nil {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || o.DelayMin == nil {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *HipChatNotificationView) HasDelayMin() bool {
	if o != nil && o.DelayMin != nil {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *HipChatNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *HipChatNotificationView) GetIntervalMin() int32 {
	if o == nil || o.IntervalMin == nil {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || o.IntervalMin == nil {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *HipChatNotificationView) HasIntervalMin() bool {
	if o != nil && o.IntervalMin != nil {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *HipChatNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetNotificationToken returns the NotificationToken field value if set, zero value otherwise.
func (o *HipChatNotificationView) GetNotificationToken() string {
	if o == nil || o.NotificationToken == nil {
		var ret string
		return ret
	}
	return *o.NotificationToken
}

// GetNotificationTokenOk returns a tuple with the NotificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotificationView) GetNotificationTokenOk() (*string, bool) {
	if o == nil || o.NotificationToken == nil {
		return nil, false
	}
	return o.NotificationToken, true
}

// HasNotificationToken returns a boolean if a field has been set.
func (o *HipChatNotificationView) HasNotificationToken() bool {
	if o != nil && o.NotificationToken != nil {
		return true
	}

	return false
}

// SetNotificationToken gets a reference to the given string and assigns it to the NotificationToken field.
func (o *HipChatNotificationView) SetNotificationToken(v string) {
	o.NotificationToken = &v
}

// GetRoomName returns the RoomName field value if set, zero value otherwise.
func (o *HipChatNotificationView) GetRoomName() string {
	if o == nil || o.RoomName == nil {
		var ret string
		return ret
	}
	return *o.RoomName
}

// GetRoomNameOk returns a tuple with the RoomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChatNotificationView) GetRoomNameOk() (*string, bool) {
	if o == nil || o.RoomName == nil {
		return nil, false
	}
	return o.RoomName, true
}

// HasRoomName returns a boolean if a field has been set.
func (o *HipChatNotificationView) HasRoomName() bool {
	if o != nil && o.RoomName != nil {
		return true
	}

	return false
}

// SetRoomName gets a reference to the given string and assigns it to the RoomName field.
func (o *HipChatNotificationView) SetRoomName(v string) {
	o.RoomName = &v
}

// GetTypeName returns the TypeName field value
func (o *HipChatNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *HipChatNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *HipChatNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

func (o HipChatNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelayMin != nil {
		toSerialize["delayMin"] = o.DelayMin
	}
	if o.IntervalMin != nil {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if o.NotificationToken != nil {
		toSerialize["notificationToken"] = o.NotificationToken
	}
	if o.RoomName != nil {
		toSerialize["roomName"] = o.RoomName
	}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableHipChatNotificationView struct {
	value *HipChatNotificationView
	isSet bool
}

func (v NullableHipChatNotificationView) Get() *HipChatNotificationView {
	return v.value
}

func (v *NullableHipChatNotificationView) Set(val *HipChatNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableHipChatNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableHipChatNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHipChatNotificationView(val *HipChatNotificationView) *NullableHipChatNotificationView {
	return &NullableHipChatNotificationView{value: val, isSet: true}
}

func (v NullableHipChatNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHipChatNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


