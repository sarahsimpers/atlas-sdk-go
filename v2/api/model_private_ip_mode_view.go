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

// PrivateIPModeView struct for PrivateIPModeView
type PrivateIPModeView struct {
	// Flag that indicates whether someone enabled **Connect via Peering Only** mode for the specified project.
	Enabled bool `json:"enabled"`
}

// NewPrivateIPModeView instantiates a new PrivateIPModeView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateIPModeView() *PrivateIPModeView {
	this := PrivateIPModeView{}
	return &this
}

// NewPrivateIPModeViewWithDefaults instantiates a new PrivateIPModeView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateIPModeViewWithDefaults() *PrivateIPModeView {
	this := PrivateIPModeView{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *PrivateIPModeView) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PrivateIPModeView) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PrivateIPModeView) SetEnabled(v bool) {
	o.Enabled = v
}

func (o PrivateIPModeView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateIPModeView struct {
	value *PrivateIPModeView
	isSet bool
}

func (v NullablePrivateIPModeView) Get() *PrivateIPModeView {
	return v.value
}

func (v *NullablePrivateIPModeView) Set(val *PrivateIPModeView) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateIPModeView) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateIPModeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateIPModeView(val *PrivateIPModeView) *NullablePrivateIPModeView {
	return &NullablePrivateIPModeView{value: val, isSet: true}
}

func (v NullablePrivateIPModeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateIPModeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


