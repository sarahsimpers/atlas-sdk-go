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

// Group struct for Group
type Group struct {
	// Quantity of MongoDB Cloud clusters deployed in this project.
	ClusterCount int64 `json:"clusterCount"`
	// Date and time when MongoDB Cloud created this project. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created time.Time `json:"created"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud project.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Human-readable label that identifies the project included in the MongoDB Cloud organization.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud organization to which the project belongs.
	OrgId string `json:"orgId"`
	// Flag that indicates whether to create the project with default alert settings.
	WithDefaultAlertsSettings *bool `json:"withDefaultAlertsSettings,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetClusterCount returns the ClusterCount field value
func (o *Group) GetClusterCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value
// and a boolean to check if the value has been set.
func (o *Group) GetClusterCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterCount, true
}

// SetClusterCount sets field value
func (o *Group) SetClusterCount(v int64) {
	o.ClusterCount = v
}

// GetCreated returns the Created field value
func (o *Group) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Group) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Group) SetLinks(v []Link) {
	o.Links = v
}

// GetName returns the Name field value
func (o *Group) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Group) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *Group) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *Group) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *Group) SetOrgId(v string) {
	o.OrgId = v
}

// GetWithDefaultAlertsSettings returns the WithDefaultAlertsSettings field value if set, zero value otherwise.
func (o *Group) GetWithDefaultAlertsSettings() bool {
	if o == nil || o.WithDefaultAlertsSettings == nil {
		var ret bool
		return ret
	}
	return *o.WithDefaultAlertsSettings
}

// GetWithDefaultAlertsSettingsOk returns a tuple with the WithDefaultAlertsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetWithDefaultAlertsSettingsOk() (*bool, bool) {
	if o == nil || o.WithDefaultAlertsSettings == nil {
		return nil, false
	}
	return o.WithDefaultAlertsSettings, true
}

// HasWithDefaultAlertsSettings returns a boolean if a field has been set.
func (o *Group) HasWithDefaultAlertsSettings() bool {
	if o != nil && o.WithDefaultAlertsSettings != nil {
		return true
	}

	return false
}

// SetWithDefaultAlertsSettings gets a reference to the given bool and assigns it to the WithDefaultAlertsSettings field.
func (o *Group) SetWithDefaultAlertsSettings(v bool) {
	o.WithDefaultAlertsSettings = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clusterCount"] = o.ClusterCount
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["orgId"] = o.OrgId
	}
	if o.WithDefaultAlertsSettings != nil {
		toSerialize["withDefaultAlertsSettings"] = o.WithDefaultAlertsSettings
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


