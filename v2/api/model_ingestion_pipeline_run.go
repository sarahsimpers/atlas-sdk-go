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

// IngestionPipelineRun Run details of a Data Lake Pipeline.
type IngestionPipelineRun struct {
	// Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
	Id *string `json:"_id,omitempty"`
	// Backup schedule interval of the Data Lake Pipeline.
	BackupFrequencyType *string `json:"backupFrequencyType,omitempty"`
	// Timestamp that indicates when the pipeline run was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Human-readable label that identifies the dataset that Atlas generates during this pipeline run. You can use this dataset as a `dataSource` in a Federated Database collection.
	DatasetName *string `json:"datasetName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
	// Timestamp that indicates the last time that the pipeline run was updated.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// Processing phase of the Data Lake Pipeline.
	Phase *string `json:"phase,omitempty"`
	// Unique 24-hexadecimal character string that identifies a Data Lake Pipeline.
	PipelineId *string `json:"pipelineId,omitempty"`
	// Unique 24-hexadecimal character string that identifies the snapshot of a cluster.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// State of the pipeline run.
	State *string `json:"state,omitempty"`
	Stats *PipelineRunStats `json:"stats,omitempty"`
}

// NewIngestionPipelineRun instantiates a new IngestionPipelineRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionPipelineRun() *IngestionPipelineRun {
	this := IngestionPipelineRun{}
	return &this
}

// NewIngestionPipelineRunWithDefaults instantiates a new IngestionPipelineRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionPipelineRunWithDefaults() *IngestionPipelineRun {
	this := IngestionPipelineRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IngestionPipelineRun) SetId(v string) {
	o.Id = &v
}

// GetBackupFrequencyType returns the BackupFrequencyType field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetBackupFrequencyType() string {
	if o == nil || o.BackupFrequencyType == nil {
		var ret string
		return ret
	}
	return *o.BackupFrequencyType
}

// GetBackupFrequencyTypeOk returns a tuple with the BackupFrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetBackupFrequencyTypeOk() (*string, bool) {
	if o == nil || o.BackupFrequencyType == nil {
		return nil, false
	}
	return o.BackupFrequencyType, true
}

// HasBackupFrequencyType returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasBackupFrequencyType() bool {
	if o != nil && o.BackupFrequencyType != nil {
		return true
	}

	return false
}

// SetBackupFrequencyType gets a reference to the given string and assigns it to the BackupFrequencyType field.
func (o *IngestionPipelineRun) SetBackupFrequencyType(v string) {
	o.BackupFrequencyType = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetCreatedDate() time.Time {
	if o == nil || o.CreatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *IngestionPipelineRun) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDatasetName returns the DatasetName field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetDatasetName() string {
	if o == nil || o.DatasetName == nil {
		var ret string
		return ret
	}
	return *o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetDatasetNameOk() (*string, bool) {
	if o == nil || o.DatasetName == nil {
		return nil, false
	}
	return o.DatasetName, true
}

// HasDatasetName returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasDatasetName() bool {
	if o != nil && o.DatasetName != nil {
		return true
	}

	return false
}

// SetDatasetName gets a reference to the given string and assigns it to the DatasetName field.
func (o *IngestionPipelineRun) SetDatasetName(v string) {
	o.DatasetName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *IngestionPipelineRun) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetLastUpdatedDate() time.Time {
	if o == nil || o.LastUpdatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedDate == nil {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasLastUpdatedDate() bool {
	if o != nil && o.LastUpdatedDate != nil {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *IngestionPipelineRun) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *IngestionPipelineRun) SetPhase(v string) {
	o.Phase = &v
}

// GetPipelineId returns the PipelineId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetPipelineId() string {
	if o == nil || o.PipelineId == nil {
		var ret string
		return ret
	}
	return *o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetPipelineIdOk() (*string, bool) {
	if o == nil || o.PipelineId == nil {
		return nil, false
	}
	return o.PipelineId, true
}

// HasPipelineId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasPipelineId() bool {
	if o != nil && o.PipelineId != nil {
		return true
	}

	return false
}

// SetPipelineId gets a reference to the given string and assigns it to the PipelineId field.
func (o *IngestionPipelineRun) SetPipelineId(v string) {
	o.PipelineId = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *IngestionPipelineRun) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IngestionPipelineRun) SetState(v string) {
	o.State = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *IngestionPipelineRun) GetStats() PipelineRunStats {
	if o == nil || o.Stats == nil {
		var ret PipelineRunStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionPipelineRun) GetStatsOk() (*PipelineRunStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *IngestionPipelineRun) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given PipelineRunStats and assigns it to the Stats field.
func (o *IngestionPipelineRun) SetStats(v PipelineRunStats) {
	o.Stats = &v
}

func (o IngestionPipelineRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.BackupFrequencyType != nil {
		toSerialize["backupFrequencyType"] = o.BackupFrequencyType
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.DatasetName != nil {
		toSerialize["datasetName"] = o.DatasetName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.LastUpdatedDate != nil {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.PipelineId != nil {
		toSerialize["pipelineId"] = o.PipelineId
	}
	if o.SnapshotId != nil {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	return json.Marshal(toSerialize)
}

type NullableIngestionPipelineRun struct {
	value *IngestionPipelineRun
	isSet bool
}

func (v NullableIngestionPipelineRun) Get() *IngestionPipelineRun {
	return v.value
}

func (v *NullableIngestionPipelineRun) Set(val *IngestionPipelineRun) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionPipelineRun) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionPipelineRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionPipelineRun(val *IngestionPipelineRun) *NullableIngestionPipelineRun {
	return &NullableIngestionPipelineRun{value: val, isSet: true}
}

func (v NullableIngestionPipelineRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionPipelineRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


