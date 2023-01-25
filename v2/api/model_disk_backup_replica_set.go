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

// DiskBackupReplicaSet Details of the replica set snapshot that MongoDB Cloud created.
type DiskBackupReplicaSet struct {
	// Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when `\"type\": \"replicaSet\".`
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// List that identifies the regions to which MongoDB Cloud copies the snapshot.
	CopyRegions []string `json:"copyRegions,omitempty"`
	// Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `\"status\": \"onDemand\"`.
	Description *string `json:"description,omitempty"`
	// Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Human-readable label that identifies how often this snapshot triggers.
	FrequencyType *string `json:"frequencyType,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when `\"encryptionEnabled\" : true`.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Version of the MongoDB host that this snapshot backs up.
	MongodVersion *string `json:"mongodVersion,omitempty"`
	// List that contains unique identifiers for the policy items.
	PolicyItems []string `json:"policyItems,omitempty"`
	// Human-readable label that identifies the replica set from which MongoDB Cloud took this snapshot. The resource returns this parameter when `\"type\": \"replicaSet\"`
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Human-readable label that identifies when this snapshot triggers.
	SnapshotType *string `json:"snapshotType,omitempty"`
	// Human-readable label that indicates the stage of the backup process for this snapshot.
	Status *string `json:"status,omitempty"`
	// Number of bytes taken to store the backup snapshot.
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty"`
	// Human-readable label that categorizes the cluster as a replica set or sharded cluster.
	Type *string `json:"type,omitempty"`
}

// NewDiskBackupReplicaSet instantiates a new DiskBackupReplicaSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupReplicaSet() *DiskBackupReplicaSet {
	this := DiskBackupReplicaSet{}
	return &this
}

// NewDiskBackupReplicaSetWithDefaults instantiates a new DiskBackupReplicaSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupReplicaSetWithDefaults() *DiskBackupReplicaSet {
	this := DiskBackupReplicaSet{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DiskBackupReplicaSet) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetCopyRegions returns the CopyRegions field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetCopyRegions() []string {
	if o == nil || o.CopyRegions == nil {
		var ret []string
		return ret
	}
	return o.CopyRegions
}

// GetCopyRegionsOk returns a tuple with the CopyRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetCopyRegionsOk() ([]string, bool) {
	if o == nil || o.CopyRegions == nil {
		return nil, false
	}
	return o.CopyRegions, true
}

// HasCopyRegions returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasCopyRegions() bool {
	if o != nil && o.CopyRegions != nil {
		return true
	}

	return false
}

// SetCopyRegions gets a reference to the given []string and assigns it to the CopyRegions field.
func (o *DiskBackupReplicaSet) SetCopyRegions(v []string) {
	o.CopyRegions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DiskBackupReplicaSet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DiskBackupReplicaSet) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *DiskBackupReplicaSet) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetFrequencyType() string {
	if o == nil || o.FrequencyType == nil {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || o.FrequencyType == nil {
		return nil, false
	}
	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasFrequencyType() bool {
	if o != nil && o.FrequencyType != nil {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *DiskBackupReplicaSet) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskBackupReplicaSet) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupReplicaSet) SetLinks(v []Link) {
	o.Links = v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetMasterKeyUUID() string {
	if o == nil || o.MasterKeyUUID == nil {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || o.MasterKeyUUID == nil {
		return nil, false
	}
	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasMasterKeyUUID() bool {
	if o != nil && o.MasterKeyUUID != nil {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *DiskBackupReplicaSet) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetMongodVersion() string {
	if o == nil || o.MongodVersion == nil {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetMongodVersionOk() (*string, bool) {
	if o == nil || o.MongodVersion == nil {
		return nil, false
	}
	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasMongodVersion() bool {
	if o != nil && o.MongodVersion != nil {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *DiskBackupReplicaSet) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetPolicyItems() []string {
	if o == nil || o.PolicyItems == nil {
		var ret []string
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetPolicyItemsOk() ([]string, bool) {
	if o == nil || o.PolicyItems == nil {
		return nil, false
	}
	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasPolicyItems() bool {
	if o != nil && o.PolicyItems != nil {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []string and assigns it to the PolicyItems field.
func (o *DiskBackupReplicaSet) SetPolicyItems(v []string) {
	o.PolicyItems = v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetReplicaSetName() string {
	if o == nil || o.ReplicaSetName == nil {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || o.ReplicaSetName == nil {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasReplicaSetName() bool {
	if o != nil && o.ReplicaSetName != nil {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *DiskBackupReplicaSet) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetSnapshotType returns the SnapshotType field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetSnapshotType() string {
	if o == nil || o.SnapshotType == nil {
		var ret string
		return ret
	}
	return *o.SnapshotType
}

// GetSnapshotTypeOk returns a tuple with the SnapshotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetSnapshotTypeOk() (*string, bool) {
	if o == nil || o.SnapshotType == nil {
		return nil, false
	}
	return o.SnapshotType, true
}

// HasSnapshotType returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasSnapshotType() bool {
	if o != nil && o.SnapshotType != nil {
		return true
	}

	return false
}

// SetSnapshotType gets a reference to the given string and assigns it to the SnapshotType field.
func (o *DiskBackupReplicaSet) SetSnapshotType(v string) {
	o.SnapshotType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiskBackupReplicaSet) SetStatus(v string) {
	o.Status = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetStorageSizeBytes() int64 {
	if o == nil || o.StorageSizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetStorageSizeBytesOk() (*int64, bool) {
	if o == nil || o.StorageSizeBytes == nil {
		return nil, false
	}
	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasStorageSizeBytes() bool {
	if o != nil && o.StorageSizeBytes != nil {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given int64 and assigns it to the StorageSizeBytes field.
func (o *DiskBackupReplicaSet) SetStorageSizeBytes(v int64) {
	o.StorageSizeBytes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DiskBackupReplicaSet) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupReplicaSet) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DiskBackupReplicaSet) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DiskBackupReplicaSet) SetType(v string) {
	o.Type = &v
}

func (o DiskBackupReplicaSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudProvider != nil {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if o.CopyRegions != nil {
		toSerialize["copyRegions"] = o.CopyRegions
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FrequencyType != nil {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.MasterKeyUUID != nil {
		toSerialize["masterKeyUUID"] = o.MasterKeyUUID
	}
	if o.MongodVersion != nil {
		toSerialize["mongodVersion"] = o.MongodVersion
	}
	if o.PolicyItems != nil {
		toSerialize["policyItems"] = o.PolicyItems
	}
	if o.ReplicaSetName != nil {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if o.SnapshotType != nil {
		toSerialize["snapshotType"] = o.SnapshotType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StorageSizeBytes != nil {
		toSerialize["storageSizeBytes"] = o.StorageSizeBytes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDiskBackupReplicaSet struct {
	value *DiskBackupReplicaSet
	isSet bool
}

func (v NullableDiskBackupReplicaSet) Get() *DiskBackupReplicaSet {
	return v.value
}

func (v *NullableDiskBackupReplicaSet) Set(val *DiskBackupReplicaSet) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupReplicaSet) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupReplicaSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupReplicaSet(val *DiskBackupReplicaSet) *NullableDiskBackupReplicaSet {
	return &NullableDiskBackupReplicaSet{value: val, isSet: true}
}

func (v NullableDiskBackupReplicaSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupReplicaSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


