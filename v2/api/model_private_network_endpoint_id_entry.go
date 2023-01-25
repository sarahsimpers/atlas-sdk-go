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

// PrivateNetworkEndpointIdEntry struct for PrivateNetworkEndpointIdEntry
type PrivateNetworkEndpointIdEntry struct {
	// Human-readable string to associate with this private endpoint.
	Comment *string `json:"comment,omitempty"`
	// Unique 22-character alphanumeric string that identifies the private endpoint.
	EndpointId string `json:"endpointId"`
	// Human-readable label that identifies the cloud service provider. Atlas Data Lake supports Amazon Web Services only.
	Provider *string `json:"provider,omitempty"`
	// Human-readable label that identifies the resource type associated with this private endpoint.
	Type *string `json:"type,omitempty"`
}

// NewPrivateNetworkEndpointIdEntry instantiates a new PrivateNetworkEndpointIdEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkEndpointIdEntry() *PrivateNetworkEndpointIdEntry {
	this := PrivateNetworkEndpointIdEntry{}
	var provider string = "AWS"
	this.Provider = &provider
	var type_ string = "DATA_LAKE"
	this.Type = &type_
	return &this
}

// NewPrivateNetworkEndpointIdEntryWithDefaults instantiates a new PrivateNetworkEndpointIdEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkEndpointIdEntryWithDefaults() *PrivateNetworkEndpointIdEntry {
	this := PrivateNetworkEndpointIdEntry{}
	var provider string = "AWS"
	this.Provider = &provider
	var type_ string = "DATA_LAKE"
	this.Type = &type_
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PrivateNetworkEndpointIdEntry) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkEndpointIdEntry) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PrivateNetworkEndpointIdEntry) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *PrivateNetworkEndpointIdEntry) SetComment(v string) {
	o.Comment = &v
}

// GetEndpointId returns the EndpointId field value
func (o *PrivateNetworkEndpointIdEntry) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkEndpointIdEntry) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *PrivateNetworkEndpointIdEntry) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PrivateNetworkEndpointIdEntry) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkEndpointIdEntry) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PrivateNetworkEndpointIdEntry) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PrivateNetworkEndpointIdEntry) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PrivateNetworkEndpointIdEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkEndpointIdEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PrivateNetworkEndpointIdEntry) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PrivateNetworkEndpointIdEntry) SetType(v string) {
	o.Type = &v
}

func (o PrivateNetworkEndpointIdEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkEndpointIdEntry struct {
	value *PrivateNetworkEndpointIdEntry
	isSet bool
}

func (v NullablePrivateNetworkEndpointIdEntry) Get() *PrivateNetworkEndpointIdEntry {
	return v.value
}

func (v *NullablePrivateNetworkEndpointIdEntry) Set(val *PrivateNetworkEndpointIdEntry) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkEndpointIdEntry) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkEndpointIdEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkEndpointIdEntry(val *PrivateNetworkEndpointIdEntry) *NullablePrivateNetworkEndpointIdEntry {
	return &NullablePrivateNetworkEndpointIdEntry{value: val, isSet: true}
}

func (v NullablePrivateNetworkEndpointIdEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkEndpointIdEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


