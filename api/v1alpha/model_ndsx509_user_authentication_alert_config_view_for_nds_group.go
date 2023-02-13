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

// NDSX509UserAuthenticationAlertConfigViewForNdsGroup NDS X509 User Authentication alert configuration allows to select thresholds for expiration of client, CA certificates and CRL which trigger alerts and how users are notified.
type NDSX509UserAuthenticationAlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	EventTypeName NDSX509UserAuthenticationEventTypeViewAlertable `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	Matchers []MatcherView `json:"matchers,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	Threshold *LessThanDaysThresholdView `json:"threshold,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewNDSX509UserAuthenticationAlertConfigViewForNdsGroup instantiates a new NDSX509UserAuthenticationAlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSX509UserAuthenticationAlertConfigViewForNdsGroup() *NDSX509UserAuthenticationAlertConfigViewForNdsGroup {
	this := NDSX509UserAuthenticationAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewNDSX509UserAuthenticationAlertConfigViewForNdsGroupWithDefaults instantiates a new NDSX509UserAuthenticationAlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSX509UserAuthenticationAlertConfigViewForNdsGroupWithDefaults() *NDSX509UserAuthenticationAlertConfigViewForNdsGroup {
	this := NDSX509UserAuthenticationAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetEventTypeName() NDSX509UserAuthenticationEventTypeViewAlertable {
	if o == nil {
		var ret NDSX509UserAuthenticationEventTypeViewAlertable
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*NDSX509UserAuthenticationEventTypeViewAlertable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetEventTypeName(v NDSX509UserAuthenticationEventTypeViewAlertable) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetMatchers() []MatcherView {
	if o == nil || o.Matchers == nil {
		var ret []MatcherView
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetMatchersOk() ([]MatcherView, bool) {
	if o == nil || o.Matchers == nil {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && o.Matchers != nil {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []MatcherView and assigns it to the Matchers field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetMatchers(v []MatcherView) {
	o.Matchers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || o.Notifications == nil {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetThreshold() LessThanDaysThresholdView {
	if o == nil || o.Threshold == nil {
		var ret LessThanDaysThresholdView
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetThresholdOk() (*LessThanDaysThresholdView, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given LessThanDaysThresholdView and assigns it to the Threshold field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetThreshold(v LessThanDaysThresholdView) {
	o.Threshold = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o NDSX509UserAuthenticationAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.Matchers != nil {
		toSerialize["matchers"] = o.Matchers
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup struct {
	value *NDSX509UserAuthenticationAlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) Get() *NDSX509UserAuthenticationAlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) Set(val *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup(val *NDSX509UserAuthenticationAlertConfigViewForNdsGroup) *NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup {
	return &NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSX509UserAuthenticationAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


