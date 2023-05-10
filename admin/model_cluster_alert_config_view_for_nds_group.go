// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the ClusterAlertConfigViewForNdsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterAlertConfigViewForNdsGroup{}

// ClusterAlertConfigViewForNdsGroup Cluster alert configuration allows to select which conditions of mongod cluster which trigger alerts and how users are notified.
type ClusterAlertConfigViewForNdsGroup struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled       *bool                         `json:"enabled,omitempty"`
	EventTypeName ClusterEventTypeViewAlertable `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.
	Matchers []ClusterMatcher `json:"matchers,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []NotificationViewForNdsGroup `json:"notifications,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewClusterAlertConfigViewForNdsGroup instantiates a new ClusterAlertConfigViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAlertConfigViewForNdsGroup(eventTypeName ClusterEventTypeViewAlertable) *ClusterAlertConfigViewForNdsGroup {
	this := ClusterAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	this.EventTypeName = eventTypeName
	return &this
}

// NewClusterAlertConfigViewForNdsGroupWithDefaults instantiates a new ClusterAlertConfigViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAlertConfigViewForNdsGroupWithDefaults() *ClusterAlertConfigViewForNdsGroup {
	this := ClusterAlertConfigViewForNdsGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ClusterAlertConfigViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterAlertConfigViewForNdsGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ClusterAlertConfigViewForNdsGroup) GetEventTypeName() ClusterEventTypeViewAlertable {
	if o == nil {
		var ret ClusterEventTypeViewAlertable
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetEventTypeNameOk() (*ClusterEventTypeViewAlertable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ClusterAlertConfigViewForNdsGroup) SetEventTypeName(v ClusterEventTypeViewAlertable) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ClusterAlertConfigViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterAlertConfigViewForNdsGroup) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ClusterAlertConfigViewForNdsGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetMatchers() []ClusterMatcher {
	if o == nil || IsNil(o.Matchers) {
		var ret []ClusterMatcher
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetMatchersOk() ([]ClusterMatcher, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []ClusterMatcher and assigns it to the Matchers field.
func (o *ClusterAlertConfigViewForNdsGroup) SetMatchers(v []ClusterMatcher) {
	o.Matchers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetNotifications() []NotificationViewForNdsGroup {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationViewForNdsGroup
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetNotificationsOk() ([]NotificationViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationViewForNdsGroup and assigns it to the Notifications field.
func (o *ClusterAlertConfigViewForNdsGroup) SetNotifications(v []NotificationViewForNdsGroup) {
	o.Notifications = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ClusterAlertConfigViewForNdsGroup) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertConfigViewForNdsGroup) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ClusterAlertConfigViewForNdsGroup) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ClusterAlertConfigViewForNdsGroup) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o ClusterAlertConfigViewForNdsGroup) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterAlertConfigViewForNdsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["eventTypeName"] = o.EventTypeName
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableClusterAlertConfigViewForNdsGroup struct {
	value *ClusterAlertConfigViewForNdsGroup
	isSet bool
}

func (v NullableClusterAlertConfigViewForNdsGroup) Get() *ClusterAlertConfigViewForNdsGroup {
	return v.value
}

func (v *NullableClusterAlertConfigViewForNdsGroup) Set(val *ClusterAlertConfigViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAlertConfigViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAlertConfigViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAlertConfigViewForNdsGroup(val *ClusterAlertConfigViewForNdsGroup) *NullableClusterAlertConfigViewForNdsGroup {
	return &NullableClusterAlertConfigViewForNdsGroup{value: val, isSet: true}
}

func (v NullableClusterAlertConfigViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAlertConfigViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}