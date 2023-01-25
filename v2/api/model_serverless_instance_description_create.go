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

// ServerlessInstanceDescriptionCreate Settings that you can specify when you create a serverless instance.
type ServerlessInstanceDescriptionCreate struct {
	// Human-readable label that identifies the serverless instance.
	Name string `json:"name"`
	ProviderSettings ServerlessProviderSettings `json:"providerSettings"`
	ServerlessBackupOptions *ServerlessBackupOptions `json:"serverlessBackupOptions,omitempty"`
	// Human-readable label that indicates the current operating condition of the serverless instance.
	StateName *string `json:"stateName,omitempty"`
	// Flag that indicates whether termination protection is enabled on the serverless instance. If set to `true`, MongoDB Cloud won't delete the serverless instance. If set to `false`, MongoDB Cloud will delete the serverless instance.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
}

// NewServerlessInstanceDescriptionCreate instantiates a new ServerlessInstanceDescriptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessInstanceDescriptionCreate() *ServerlessInstanceDescriptionCreate {
	this := ServerlessInstanceDescriptionCreate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// NewServerlessInstanceDescriptionCreateWithDefaults instantiates a new ServerlessInstanceDescriptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessInstanceDescriptionCreateWithDefaults() *ServerlessInstanceDescriptionCreate {
	this := ServerlessInstanceDescriptionCreate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// GetName returns the Name field value
func (o *ServerlessInstanceDescriptionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerlessInstanceDescriptionCreate) SetName(v string) {
	o.Name = v
}

// GetProviderSettings returns the ProviderSettings field value
func (o *ServerlessInstanceDescriptionCreate) GetProviderSettings() ServerlessProviderSettings {
	if o == nil {
		var ret ServerlessProviderSettings
		return ret
	}

	return o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetProviderSettingsOk() (*ServerlessProviderSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderSettings, true
}

// SetProviderSettings sets field value
func (o *ServerlessInstanceDescriptionCreate) SetProviderSettings(v ServerlessProviderSettings) {
	o.ProviderSettings = v
}

// GetServerlessBackupOptions returns the ServerlessBackupOptions field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptions() ServerlessBackupOptions {
	if o == nil || o.ServerlessBackupOptions == nil {
		var ret ServerlessBackupOptions
		return ret
	}
	return *o.ServerlessBackupOptions
}

// GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptionsOk() (*ServerlessBackupOptions, bool) {
	if o == nil || o.ServerlessBackupOptions == nil {
		return nil, false
	}
	return o.ServerlessBackupOptions, true
}

// HasServerlessBackupOptions returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasServerlessBackupOptions() bool {
	if o != nil && o.ServerlessBackupOptions != nil {
		return true
	}

	return false
}

// SetServerlessBackupOptions gets a reference to the given ServerlessBackupOptions and assigns it to the ServerlessBackupOptions field.
func (o *ServerlessInstanceDescriptionCreate) SetServerlessBackupOptions(v ServerlessBackupOptions) {
	o.ServerlessBackupOptions = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionCreate) GetStateName() string {
	if o == nil || o.StateName == nil {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetStateNameOk() (*string, bool) {
	if o == nil || o.StateName == nil {
		return nil, false
	}
	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasStateName() bool {
	if o != nil && o.StateName != nil {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *ServerlessInstanceDescriptionCreate) SetStateName(v string) {
	o.StateName = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabled() bool {
	if o == nil || o.TerminationProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || o.TerminationProtectionEnabled == nil {
		return nil, false
	}
	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasTerminationProtectionEnabled() bool {
	if o != nil && o.TerminationProtectionEnabled != nil {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *ServerlessInstanceDescriptionCreate) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

func (o ServerlessInstanceDescriptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["providerSettings"] = o.ProviderSettings
	}
	if o.ServerlessBackupOptions != nil {
		toSerialize["serverlessBackupOptions"] = o.ServerlessBackupOptions
	}
	if o.StateName != nil {
		toSerialize["stateName"] = o.StateName
	}
	if o.TerminationProtectionEnabled != nil {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableServerlessInstanceDescriptionCreate struct {
	value *ServerlessInstanceDescriptionCreate
	isSet bool
}

func (v NullableServerlessInstanceDescriptionCreate) Get() *ServerlessInstanceDescriptionCreate {
	return v.value
}

func (v *NullableServerlessInstanceDescriptionCreate) Set(val *ServerlessInstanceDescriptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessInstanceDescriptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessInstanceDescriptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessInstanceDescriptionCreate(val *ServerlessInstanceDescriptionCreate) *NullableServerlessInstanceDescriptionCreate {
	return &NullableServerlessInstanceDescriptionCreate{value: val, isSet: true}
}

func (v NullableServerlessInstanceDescriptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessInstanceDescriptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


