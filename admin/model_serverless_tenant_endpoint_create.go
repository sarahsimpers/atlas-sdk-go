// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ServerlessTenantEndpointCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessTenantEndpointCreate{}

// ServerlessTenantEndpointCreate struct for ServerlessTenantEndpointCreate
type ServerlessTenantEndpointCreate struct {
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
}

// NewServerlessTenantEndpointCreate instantiates a new ServerlessTenantEndpointCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessTenantEndpointCreate() *ServerlessTenantEndpointCreate {
	this := ServerlessTenantEndpointCreate{}
	return &this
}

// NewServerlessTenantEndpointCreateWithDefaults instantiates a new ServerlessTenantEndpointCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessTenantEndpointCreateWithDefaults() *ServerlessTenantEndpointCreate {
	this := ServerlessTenantEndpointCreate{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessTenantEndpointCreate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessTenantEndpointCreate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessTenantEndpointCreate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessTenantEndpointCreate) SetComment(v string) {
	o.Comment = &v
}

func (o ServerlessTenantEndpointCreate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessTenantEndpointCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableServerlessTenantEndpointCreate struct {
	value *ServerlessTenantEndpointCreate
	isSet bool
}

func (v NullableServerlessTenantEndpointCreate) Get() *ServerlessTenantEndpointCreate {
	return v.value
}

func (v *NullableServerlessTenantEndpointCreate) Set(val *ServerlessTenantEndpointCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessTenantEndpointCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessTenantEndpointCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessTenantEndpointCreate(val *ServerlessTenantEndpointCreate) *NullableServerlessTenantEndpointCreate {
	return &NullableServerlessTenantEndpointCreate{value: val, isSet: true}
}

func (v NullableServerlessTenantEndpointCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessTenantEndpointCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
