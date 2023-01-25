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

// ProjectSettingItemView struct for ProjectSettingItemView
type ProjectSettingItemView struct {
	// Flag that indicates whether someone enabled the regionalized private endpoint setting for the specified project.  - Set this value to `true` to enable regionalized private endpoints. This allows you to create more than one private endpoint in a cloud provider region. You need to enable this setting to connect to multi-region and global MongoDB Cloud sharded clusters. Enabling regionalized private endpoints introduces the following limitations:   - Your applications must use the new connection strings for existing multi-region and global sharded clusters. This might cause downtime.   - Your MongoDB Cloud project can't contain replica sets nor can you create new replica sets in this project.    - You can't disable this setting if you have:     - more than one private endpoint in more than one region     - more than one private endpoint in one region and one private endpoint in one or more regions.  - Set this value to `false` to disable regionalized private endpoints.
	Enabled bool `json:"enabled"`
}

// NewProjectSettingItemView instantiates a new ProjectSettingItemView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSettingItemView() *ProjectSettingItemView {
	this := ProjectSettingItemView{}
	return &this
}

// NewProjectSettingItemViewWithDefaults instantiates a new ProjectSettingItemView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSettingItemViewWithDefaults() *ProjectSettingItemView {
	this := ProjectSettingItemView{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ProjectSettingItemView) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProjectSettingItemView) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProjectSettingItemView) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ProjectSettingItemView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableProjectSettingItemView struct {
	value *ProjectSettingItemView
	isSet bool
}

func (v NullableProjectSettingItemView) Get() *ProjectSettingItemView {
	return v.value
}

func (v *NullableProjectSettingItemView) Set(val *ProjectSettingItemView) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSettingItemView) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSettingItemView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSettingItemView(val *ProjectSettingItemView) *NullableProjectSettingItemView {
	return &NullableProjectSettingItemView{value: val, isSet: true}
}

func (v NullableProjectSettingItemView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSettingItemView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


