// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the CreateApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiKey{}

// CreateApiKey struct for CreateApiKey
type CreateApiKey struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization or project.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateApiKey instantiates a new CreateApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKey() *CreateApiKey {
	this := CreateApiKey{}
	return &this
}

// NewCreateApiKeyWithDefaults instantiates a new CreateApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyWithDefaults() *CreateApiKey {
	this := CreateApiKey{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *CreateApiKey) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKey) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *CreateApiKey) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *CreateApiKey) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateApiKey) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKey) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateApiKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateApiKey) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateApiKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCreateApiKey struct {
	value *CreateApiKey
	isSet bool
}

func (v NullableCreateApiKey) Get() *CreateApiKey {
	return v.value
}

func (v *NullableCreateApiKey) Set(val *CreateApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKey(val *CreateApiKey) *NullableCreateApiKey {
	return &NullableCreateApiKey{value: val, isSet: true}
}

func (v NullableCreateApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
