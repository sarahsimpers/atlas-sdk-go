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

// DiskBackupOnDemandSnapshotRequest struct for DiskBackupOnDemandSnapshotRequest
type DiskBackupOnDemandSnapshotRequest struct {
	// Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `\"status\" : \"onDemand\"`.
	Description *string `json:"description,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1**
	RetentionInDays *int32 `json:"retentionInDays,omitempty"`
}

// NewDiskBackupOnDemandSnapshotRequest instantiates a new DiskBackupOnDemandSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupOnDemandSnapshotRequest() *DiskBackupOnDemandSnapshotRequest {
	this := DiskBackupOnDemandSnapshotRequest{}
	return &this
}

// NewDiskBackupOnDemandSnapshotRequestWithDefaults instantiates a new DiskBackupOnDemandSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupOnDemandSnapshotRequestWithDefaults() *DiskBackupOnDemandSnapshotRequest {
	this := DiskBackupOnDemandSnapshotRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DiskBackupOnDemandSnapshotRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupOnDemandSnapshotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DiskBackupOnDemandSnapshotRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DiskBackupOnDemandSnapshotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiskBackupOnDemandSnapshotRequest) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupOnDemandSnapshotRequest) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupOnDemandSnapshotRequest) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupOnDemandSnapshotRequest) SetLinks(v []Link) {
	o.Links = v
}

// GetRetentionInDays returns the RetentionInDays field value if set, zero value otherwise.
func (o *DiskBackupOnDemandSnapshotRequest) GetRetentionInDays() int32 {
	if o == nil || o.RetentionInDays == nil {
		var ret int32
		return ret
	}
	return *o.RetentionInDays
}

// GetRetentionInDaysOk returns a tuple with the RetentionInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupOnDemandSnapshotRequest) GetRetentionInDaysOk() (*int32, bool) {
	if o == nil || o.RetentionInDays == nil {
		return nil, false
	}
	return o.RetentionInDays, true
}

// HasRetentionInDays returns a boolean if a field has been set.
func (o *DiskBackupOnDemandSnapshotRequest) HasRetentionInDays() bool {
	if o != nil && o.RetentionInDays != nil {
		return true
	}

	return false
}

// SetRetentionInDays gets a reference to the given int32 and assigns it to the RetentionInDays field.
func (o *DiskBackupOnDemandSnapshotRequest) SetRetentionInDays(v int32) {
	o.RetentionInDays = &v
}

func (o DiskBackupOnDemandSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.RetentionInDays != nil {
		toSerialize["retentionInDays"] = o.RetentionInDays
	}
	return json.Marshal(toSerialize)
}

type NullableDiskBackupOnDemandSnapshotRequest struct {
	value *DiskBackupOnDemandSnapshotRequest
	isSet bool
}

func (v NullableDiskBackupOnDemandSnapshotRequest) Get() *DiskBackupOnDemandSnapshotRequest {
	return v.value
}

func (v *NullableDiskBackupOnDemandSnapshotRequest) Set(val *DiskBackupOnDemandSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupOnDemandSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupOnDemandSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupOnDemandSnapshotRequest(val *DiskBackupOnDemandSnapshotRequest) *NullableDiskBackupOnDemandSnapshotRequest {
	return &NullableDiskBackupOnDemandSnapshotRequest{value: val, isSet: true}
}

func (v NullableDiskBackupOnDemandSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupOnDemandSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


