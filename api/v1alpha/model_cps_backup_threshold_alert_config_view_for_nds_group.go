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

// checks if the CpsBackupThresholdAlertConfigViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpsBackupThresholdAlertConfigViewForNdsGroup{}

// CpsBackupThresholdAlertConfigViewForNdsGroup Cps Backup threshold alert configuration allows to select thresholds for conditions of CPS backup or oplogs anomalies which trigger alerts and how users are notified.
type CpsBackupThresholdAlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	EventTypeName CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	Matchers []MatcherView `json:"matchers,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	Threshold *GreaterThanTimeThresholdView `json:"threshold,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewCpsBackupThresholdAlertConfigViewForNdsGroup instantiates a new CpsBackupThresholdAlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpsBackupThresholdAlertConfigViewForNdsGroup() *CpsBackupThresholdAlertConfigViewForNdsGroup {
	this := CpsBackupThresholdAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewCpsBackupThresholdAlertConfigViewForNdsGroupWithDefaults instantiates a new CpsBackupThresholdAlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpsBackupThresholdAlertConfigViewForNdsGroupWithDefaults() *CpsBackupThresholdAlertConfigViewForNdsGroup {
	this := CpsBackupThresholdAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetEventTypeName() CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold {
	if o == nil {
		var ret CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetEventTypeName(v CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetMatchers() []MatcherView {
	if o == nil || IsNil(o.Matchers) {
		var ret []MatcherView
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetMatchersOk() ([]MatcherView, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []MatcherView and assigns it to the Matchers field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetMatchers(v []MatcherView) {
	o.Matchers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetThreshold() GreaterThanTimeThresholdView {
	if o == nil || IsNil(o.Threshold) {
		var ret GreaterThanTimeThresholdView
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetThresholdOk() (*GreaterThanTimeThresholdView, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given GreaterThanTimeThresholdView and assigns it to the Threshold field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetThreshold(v GreaterThanTimeThresholdView) {
	o.Threshold = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *CpsBackupThresholdAlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o CpsBackupThresholdAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpsBackupThresholdAlertConfigViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created is readOnly
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["eventTypeName"] = o.EventTypeName
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: links is readOnly
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	// skip: updated is readOnly
	return toSerialize, nil
}

type NullableCpsBackupThresholdAlertConfigViewForNdsGroup struct {
	value *CpsBackupThresholdAlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableCpsBackupThresholdAlertConfigViewForNdsGroup) Get() *CpsBackupThresholdAlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableCpsBackupThresholdAlertConfigViewForNdsGroup) Set(val *CpsBackupThresholdAlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCpsBackupThresholdAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCpsBackupThresholdAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpsBackupThresholdAlertConfigViewForNdsGroup(val *CpsBackupThresholdAlertConfigViewForNdsGroup) *NullableCpsBackupThresholdAlertConfigViewForNdsGroup {
	return &NullableCpsBackupThresholdAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableCpsBackupThresholdAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpsBackupThresholdAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


