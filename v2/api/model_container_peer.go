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

// ContainerPeer struct for ContainerPeer
type ContainerPeer struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
}

// NewContainerPeer instantiates a new ContainerPeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerPeer() *ContainerPeer {
	this := ContainerPeer{}
	return &this
}

// NewContainerPeerWithDefaults instantiates a new ContainerPeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerPeerWithDefaults() *ContainerPeer {
	this := ContainerPeer{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *ContainerPeer) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *ContainerPeer) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *ContainerPeer) SetContainerId(v string) {
	o.ContainerId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerPeer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPeer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerPeer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerPeer) SetId(v string) {
	o.Id = &v
}

func (o ContainerPeer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerId"] = o.ContainerId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableContainerPeer struct {
	value *ContainerPeer
	isSet bool
}

func (v NullableContainerPeer) Get() *ContainerPeer {
	return v.value
}

func (v *NullableContainerPeer) Set(val *ContainerPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerPeer(val *ContainerPeer) *NullableContainerPeer {
	return &NullableContainerPeer{value: val, isSet: true}
}

func (v NullableContainerPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


