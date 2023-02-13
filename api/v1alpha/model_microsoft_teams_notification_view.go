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

// MicrosoftTeamsNotificationView Microsoft Teams notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type MicrosoftTeamsNotificationView struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. The resource requires this parameter when `\"notifications.[n].typeName\" : \"MICROSOFT_TEAMS\"`. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted. 
	MicrosoftTeamsWebhookUrl *string `json:"microsoftTeamsWebhookUrl,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewMicrosoftTeamsNotificationView instantiates a new MicrosoftTeamsNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftTeamsNotificationView() *MicrosoftTeamsNotificationView {
	this := MicrosoftTeamsNotificationView{}
	return &this
}

// NewMicrosoftTeamsNotificationViewWithDefaults instantiates a new MicrosoftTeamsNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftTeamsNotificationViewWithDefaults() *MicrosoftTeamsNotificationView {
	this := MicrosoftTeamsNotificationView{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *MicrosoftTeamsNotificationView) GetDelayMin() int32 {
	if o == nil || o.DelayMin == nil {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || o.DelayMin == nil {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *MicrosoftTeamsNotificationView) HasDelayMin() bool {
	if o != nil && o.DelayMin != nil {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *MicrosoftTeamsNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *MicrosoftTeamsNotificationView) GetIntervalMin() int32 {
	if o == nil || o.IntervalMin == nil {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || o.IntervalMin == nil {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *MicrosoftTeamsNotificationView) HasIntervalMin() bool {
	if o != nil && o.IntervalMin != nil {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *MicrosoftTeamsNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field value if set, zero value otherwise.
func (o *MicrosoftTeamsNotificationView) GetMicrosoftTeamsWebhookUrl() string {
	if o == nil || o.MicrosoftTeamsWebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.MicrosoftTeamsWebhookUrl
}

// GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsNotificationView) GetMicrosoftTeamsWebhookUrlOk() (*string, bool) {
	if o == nil || o.MicrosoftTeamsWebhookUrl == nil {
		return nil, false
	}
	return o.MicrosoftTeamsWebhookUrl, true
}

// HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
func (o *MicrosoftTeamsNotificationView) HasMicrosoftTeamsWebhookUrl() bool {
	if o != nil && o.MicrosoftTeamsWebhookUrl != nil {
		return true
	}

	return false
}

// SetMicrosoftTeamsWebhookUrl gets a reference to the given string and assigns it to the MicrosoftTeamsWebhookUrl field.
func (o *MicrosoftTeamsNotificationView) SetMicrosoftTeamsWebhookUrl(v string) {
	o.MicrosoftTeamsWebhookUrl = &v
}

// GetTypeName returns the TypeName field value
func (o *MicrosoftTeamsNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *MicrosoftTeamsNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

func (o MicrosoftTeamsNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelayMin != nil {
		toSerialize["delayMin"] = o.DelayMin
	}
	if o.IntervalMin != nil {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if o.MicrosoftTeamsWebhookUrl != nil {
		toSerialize["microsoftTeamsWebhookUrl"] = o.MicrosoftTeamsWebhookUrl
	}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftTeamsNotificationView struct {
	value *MicrosoftTeamsNotificationView
	isSet bool
}

func (v NullableMicrosoftTeamsNotificationView) Get() *MicrosoftTeamsNotificationView {
	return v.value
}

func (v *NullableMicrosoftTeamsNotificationView) Set(val *MicrosoftTeamsNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftTeamsNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftTeamsNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftTeamsNotificationView(val *MicrosoftTeamsNotificationView) *NullableMicrosoftTeamsNotificationView {
	return &NullableMicrosoftTeamsNotificationView{value: val, isSet: true}
}

func (v NullableMicrosoftTeamsNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftTeamsNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


