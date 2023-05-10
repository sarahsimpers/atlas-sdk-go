// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ServerlessInstanceDescriptionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessInstanceDescriptionUpdate{}

// ServerlessInstanceDescriptionUpdate Settings that you can update when you request a serverless cluster update.
type ServerlessInstanceDescriptionUpdate struct {
	ServerlessBackupOptions *ServerlessBackupOptions `json:"serverlessBackupOptions,omitempty"`
	// Flag that indicates whether termination protection is enabled on the serverless instance. If set to `true`, MongoDB Cloud won't delete the serverless instance. If set to `false`, MongoDB Cloud will delete the serverless instance.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
}

// NewServerlessInstanceDescriptionUpdate instantiates a new ServerlessInstanceDescriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessInstanceDescriptionUpdate() *ServerlessInstanceDescriptionUpdate {
	this := ServerlessInstanceDescriptionUpdate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// NewServerlessInstanceDescriptionUpdateWithDefaults instantiates a new ServerlessInstanceDescriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessInstanceDescriptionUpdateWithDefaults() *ServerlessInstanceDescriptionUpdate {
	this := ServerlessInstanceDescriptionUpdate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// GetServerlessBackupOptions returns the ServerlessBackupOptions field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptions() ServerlessBackupOptions {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		var ret ServerlessBackupOptions
		return ret
	}
	return *o.ServerlessBackupOptions
}

// GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptionsOk() (*ServerlessBackupOptions, bool) {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		return nil, false
	}
	return o.ServerlessBackupOptions, true
}

// HasServerlessBackupOptions returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionUpdate) HasServerlessBackupOptions() bool {
	if o != nil && !IsNil(o.ServerlessBackupOptions) {
		return true
	}

	return false
}

// SetServerlessBackupOptions gets a reference to the given ServerlessBackupOptions and assigns it to the ServerlessBackupOptions field.
func (o *ServerlessInstanceDescriptionUpdate) SetServerlessBackupOptions(v ServerlessBackupOptions) {
	o.ServerlessBackupOptions = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise.
func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}
	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionUpdate) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *ServerlessInstanceDescriptionUpdate) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

func (o ServerlessInstanceDescriptionUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessInstanceDescriptionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerlessBackupOptions) {
		toSerialize["serverlessBackupOptions"] = o.ServerlessBackupOptions
	}
	if !IsNil(o.TerminationProtectionEnabled) {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	return toSerialize, nil
}

type NullableServerlessInstanceDescriptionUpdate struct {
	value *ServerlessInstanceDescriptionUpdate
	isSet bool
}

func (v NullableServerlessInstanceDescriptionUpdate) Get() *ServerlessInstanceDescriptionUpdate {
	return v.value
}

func (v *NullableServerlessInstanceDescriptionUpdate) Set(val *ServerlessInstanceDescriptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessInstanceDescriptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessInstanceDescriptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessInstanceDescriptionUpdate(val *ServerlessInstanceDescriptionUpdate) *NullableServerlessInstanceDescriptionUpdate {
	return &NullableServerlessInstanceDescriptionUpdate{value: val, isSet: true}
}

func (v NullableServerlessInstanceDescriptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessInstanceDescriptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}