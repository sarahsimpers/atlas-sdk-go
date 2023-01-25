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

// TokenFiltertrim Filter that trims leading and trailing whitespace from tokens.
type TokenFiltertrim struct {
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltertrim instantiates a new TokenFiltertrim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltertrim() *TokenFiltertrim {
	this := TokenFiltertrim{}
	return &this
}

// NewTokenFiltertrimWithDefaults instantiates a new TokenFiltertrim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltertrimWithDefaults() *TokenFiltertrim {
	this := TokenFiltertrim{}
	return &this
}

// GetType returns the Type field value
func (o *TokenFiltertrim) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltertrim) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltertrim) SetType(v string) {
	o.Type = v
}

func (o TokenFiltertrim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTokenFiltertrim struct {
	value *TokenFiltertrim
	isSet bool
}

func (v NullableTokenFiltertrim) Get() *TokenFiltertrim {
	return v.value
}

func (v *NullableTokenFiltertrim) Set(val *TokenFiltertrim) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltertrim) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltertrim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltertrim(val *TokenFiltertrim) *NullableTokenFiltertrim {
	return &NullableTokenFiltertrim{value: val, isSet: true}
}

func (v NullableTokenFiltertrim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltertrim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


