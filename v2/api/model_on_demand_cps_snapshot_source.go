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

// OnDemandCpsSnapshotSource On-Demand Cloud Provider Snapshots as Source for a Data Lake Pipeline.
type OnDemandCpsSnapshotSource struct {
	// Human-readable name that identifies the cluster.
	ClusterName *string `json:"clusterName,omitempty"`
	// Human-readable name that identifies the collection.
	CollectionName *string `json:"collectionName,omitempty"`
	// Human-readable name that identifies the database.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
}

// NewOnDemandCpsSnapshotSource instantiates a new OnDemandCpsSnapshotSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandCpsSnapshotSource() *OnDemandCpsSnapshotSource {
	this := OnDemandCpsSnapshotSource{}
	return &this
}

// NewOnDemandCpsSnapshotSourceWithDefaults instantiates a new OnDemandCpsSnapshotSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandCpsSnapshotSourceWithDefaults() *OnDemandCpsSnapshotSource {
	this := OnDemandCpsSnapshotSource{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *OnDemandCpsSnapshotSource) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandCpsSnapshotSource) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *OnDemandCpsSnapshotSource) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *OnDemandCpsSnapshotSource) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *OnDemandCpsSnapshotSource) GetCollectionName() string {
	if o == nil || o.CollectionName == nil {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandCpsSnapshotSource) GetCollectionNameOk() (*string, bool) {
	if o == nil || o.CollectionName == nil {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *OnDemandCpsSnapshotSource) HasCollectionName() bool {
	if o != nil && o.CollectionName != nil {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *OnDemandCpsSnapshotSource) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *OnDemandCpsSnapshotSource) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandCpsSnapshotSource) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *OnDemandCpsSnapshotSource) HasDatabaseName() bool {
	if o != nil && o.DatabaseName != nil {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *OnDemandCpsSnapshotSource) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *OnDemandCpsSnapshotSource) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandCpsSnapshotSource) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *OnDemandCpsSnapshotSource) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *OnDemandCpsSnapshotSource) SetGroupId(v string) {
	o.GroupId = &v
}

func (o OnDemandCpsSnapshotSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.CollectionName != nil {
		toSerialize["collectionName"] = o.CollectionName
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableOnDemandCpsSnapshotSource struct {
	value *OnDemandCpsSnapshotSource
	isSet bool
}

func (v NullableOnDemandCpsSnapshotSource) Get() *OnDemandCpsSnapshotSource {
	return v.value
}

func (v *NullableOnDemandCpsSnapshotSource) Set(val *OnDemandCpsSnapshotSource) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandCpsSnapshotSource) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandCpsSnapshotSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandCpsSnapshotSource(val *OnDemandCpsSnapshotSource) *NullableOnDemandCpsSnapshotSource {
	return &NullableOnDemandCpsSnapshotSource{value: val, isSet: true}
}

func (v NullableOnDemandCpsSnapshotSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandCpsSnapshotSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


