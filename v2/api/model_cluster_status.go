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

// ClusterStatus struct for ClusterStatus
type ClusterStatus struct {
	// State of cluster at the time of this request.
	ChangeStatus *string `json:"changeStatus,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewClusterStatus instantiates a new ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// NewClusterStatusWithDefaults instantiates a new ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusWithDefaults() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// GetChangeStatus returns the ChangeStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetChangeStatus() string {
	if o == nil || o.ChangeStatus == nil {
		var ret string
		return ret
	}
	return *o.ChangeStatus
}

// GetChangeStatusOk returns a tuple with the ChangeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetChangeStatusOk() (*string, bool) {
	if o == nil || o.ChangeStatus == nil {
		return nil, false
	}
	return o.ChangeStatus, true
}

// HasChangeStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasChangeStatus() bool {
	if o != nil && o.ChangeStatus != nil {
		return true
	}

	return false
}

// SetChangeStatus gets a reference to the given string and assigns it to the ChangeStatus field.
func (o *ClusterStatus) SetChangeStatus(v string) {
	o.ChangeStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClusterStatus) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClusterStatus) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ClusterStatus) SetLinks(v []Link) {
	o.Links = v
}

func (o ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeStatus != nil {
		toSerialize["changeStatus"] = o.ChangeStatus
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


