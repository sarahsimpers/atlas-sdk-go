/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Checkpoint struct for Checkpoint
type Checkpoint struct {
	// Unique 24-hexadecimal digit string that identifies the cluster that contains the checkpoint.
	ClusterId *string `json:"clusterId,omitempty"`
	// Date and time when the checkpoint completed and the balancer restarted. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Completed *time.Time `json:"completed,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns the checkpoints.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies checkpoint.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host.
	Parts []ApiCheckpointPartView `json:"parts,omitempty"`
	// Flag that indicates whether MongoDB Cloud can use the checkpoint for a restore.
	Restorable *bool `json:"restorable,omitempty"`
	// Date and time when the balancer stopped and began the checkpoint. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Started *time.Time `json:"started,omitempty"`
	// Date and time to which the checkpoint restores. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewCheckpoint instantiates a new Checkpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpoint() *Checkpoint {
	this := Checkpoint{}
	return &this
}

// NewCheckpointWithDefaults instantiates a new Checkpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointWithDefaults() *Checkpoint {
	this := Checkpoint{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Checkpoint) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Checkpoint) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Checkpoint) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *Checkpoint) GetCompleted() time.Time {
	if o == nil || o.Completed == nil {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetCompletedOk() (*time.Time, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *Checkpoint) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *Checkpoint) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Checkpoint) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Checkpoint) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Checkpoint) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Checkpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Checkpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Checkpoint) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Checkpoint) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Checkpoint) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Checkpoint) SetLinks(v []Link) {
	o.Links = v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *Checkpoint) GetParts() []ApiCheckpointPartView {
	if o == nil || o.Parts == nil {
		var ret []ApiCheckpointPartView
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetPartsOk() ([]ApiCheckpointPartView, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *Checkpoint) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given []ApiCheckpointPartView and assigns it to the Parts field.
func (o *Checkpoint) SetParts(v []ApiCheckpointPartView) {
	o.Parts = v
}

// GetRestorable returns the Restorable field value if set, zero value otherwise.
func (o *Checkpoint) GetRestorable() bool {
	if o == nil || o.Restorable == nil {
		var ret bool
		return ret
	}
	return *o.Restorable
}

// GetRestorableOk returns a tuple with the Restorable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetRestorableOk() (*bool, bool) {
	if o == nil || o.Restorable == nil {
		return nil, false
	}
	return o.Restorable, true
}

// HasRestorable returns a boolean if a field has been set.
func (o *Checkpoint) HasRestorable() bool {
	if o != nil && o.Restorable != nil {
		return true
	}

	return false
}

// SetRestorable gets a reference to the given bool and assigns it to the Restorable field.
func (o *Checkpoint) SetRestorable(v bool) {
	o.Restorable = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *Checkpoint) GetStarted() time.Time {
	if o == nil || o.Started == nil {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetStartedOk() (*time.Time, bool) {
	if o == nil || o.Started == nil {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *Checkpoint) HasStarted() bool {
	if o != nil && o.Started != nil {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *Checkpoint) SetStarted(v time.Time) {
	o.Started = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Checkpoint) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Checkpoint) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Checkpoint) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o Checkpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	if o.Restorable != nil {
		toSerialize["restorable"] = o.Restorable
	}
	if o.Started != nil {
		toSerialize["started"] = o.Started
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpoint struct {
	value *Checkpoint
	isSet bool
}

func (v NullableCheckpoint) Get() *Checkpoint {
	return v.value
}

func (v *NullableCheckpoint) Set(val *Checkpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpoint(val *Checkpoint) *NullableCheckpoint {
	return &NullableCheckpoint{value: val, isSet: true}
}

func (v NullableCheckpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


