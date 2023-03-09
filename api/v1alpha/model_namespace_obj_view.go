/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the NamespaceObjView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceObjView{}

// NamespaceObjView Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
type NamespaceObjView struct {
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	Namespace *string `json:"namespace,omitempty"`
	// Human-readable label that identifies the type of namespace.
	Type *string `json:"type,omitempty"`
}

// NewNamespaceObjView instantiates a new NamespaceObjView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceObjView() *NamespaceObjView {
	this := NamespaceObjView{}
	return &this
}

// NewNamespaceObjViewWithDefaults instantiates a new NamespaceObjView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceObjViewWithDefaults() *NamespaceObjView {
	this := NamespaceObjView{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *NamespaceObjView) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceObjView) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *NamespaceObjView) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *NamespaceObjView) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NamespaceObjView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceObjView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NamespaceObjView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NamespaceObjView) SetType(v string) {
	o.Type = &v
}

func (o NamespaceObjView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceObjView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: namespace is readOnly
	// skip: type is readOnly
	return toSerialize, nil
}

type NullableNamespaceObjView struct {
	value *NamespaceObjView
	isSet bool
}

func (v NullableNamespaceObjView) Get() *NamespaceObjView {
	return v.value
}

func (v *NullableNamespaceObjView) Set(val *NamespaceObjView) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceObjView) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceObjView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceObjView(val *NamespaceObjView) *NullableNamespaceObjView {
	return &NullableNamespaceObjView{value: val, isSet: true}
}

func (v NullableNamespaceObjView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceObjView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

