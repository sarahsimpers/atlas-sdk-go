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

// AzureProviderSettings struct for AzureProviderSettings
type AzureProviderSettings struct {
	AutoScaling *AzureAutoScaling `json:"autoScaling,omitempty"`
	// Disk type that corresponds to the host's root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies.
	DiskTypeName *string `json:"diskTypeName,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Microsoft Azure Regions.
	RegionName *string `json:"regionName,omitempty"`
}

// NewAzureProviderSettings instantiates a new AzureProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProviderSettings() *AzureProviderSettings {
	this := AzureProviderSettings{}
	return &this
}

// NewAzureProviderSettingsWithDefaults instantiates a new AzureProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProviderSettingsWithDefaults() *AzureProviderSettings {
	this := AzureProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *AzureProviderSettings) GetAutoScaling() AzureAutoScaling {
	if o == nil || o.AutoScaling == nil {
		var ret AzureAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProviderSettings) GetAutoScalingOk() (*AzureAutoScaling, bool) {
	if o == nil || o.AutoScaling == nil {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *AzureProviderSettings) HasAutoScaling() bool {
	if o != nil && o.AutoScaling != nil {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given AzureAutoScaling and assigns it to the AutoScaling field.
func (o *AzureProviderSettings) SetAutoScaling(v AzureAutoScaling) {
	o.AutoScaling = &v
}

// GetDiskTypeName returns the DiskTypeName field value if set, zero value otherwise.
func (o *AzureProviderSettings) GetDiskTypeName() string {
	if o == nil || o.DiskTypeName == nil {
		var ret string
		return ret
	}
	return *o.DiskTypeName
}

// GetDiskTypeNameOk returns a tuple with the DiskTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProviderSettings) GetDiskTypeNameOk() (*string, bool) {
	if o == nil || o.DiskTypeName == nil {
		return nil, false
	}
	return o.DiskTypeName, true
}

// HasDiskTypeName returns a boolean if a field has been set.
func (o *AzureProviderSettings) HasDiskTypeName() bool {
	if o != nil && o.DiskTypeName != nil {
		return true
	}

	return false
}

// SetDiskTypeName gets a reference to the given string and assigns it to the DiskTypeName field.
func (o *AzureProviderSettings) SetDiskTypeName(v string) {
	o.DiskTypeName = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *AzureProviderSettings) GetInstanceSizeName() string {
	if o == nil || o.InstanceSizeName == nil {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || o.InstanceSizeName == nil {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *AzureProviderSettings) HasInstanceSizeName() bool {
	if o != nil && o.InstanceSizeName != nil {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *AzureProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AzureProviderSettings) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AzureProviderSettings) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AzureProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

func (o AzureProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoScaling != nil {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if o.DiskTypeName != nil {
		toSerialize["diskTypeName"] = o.DiskTypeName
	}
	if o.InstanceSizeName != nil {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	return json.Marshal(toSerialize)
}

type NullableAzureProviderSettings struct {
	value *AzureProviderSettings
	isSet bool
}

func (v NullableAzureProviderSettings) Get() *AzureProviderSettings {
	return v.value
}

func (v *NullableAzureProviderSettings) Set(val *AzureProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProviderSettings(val *AzureProviderSettings) *NullableAzureProviderSettings {
	return &NullableAzureProviderSettings{value: val, isSet: true}
}

func (v NullableAzureProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


