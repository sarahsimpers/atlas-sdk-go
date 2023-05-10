// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PaginatedCloudBackupRestoreJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCloudBackupRestoreJob{}

// PaginatedCloudBackupRestoreJob struct for PaginatedCloudBackupRestoreJob
type PaginatedCloudBackupRestoreJob struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []DiskBackupRestoreJob `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedCloudBackupRestoreJob instantiates a new PaginatedCloudBackupRestoreJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudBackupRestoreJob() *PaginatedCloudBackupRestoreJob {
	this := PaginatedCloudBackupRestoreJob{}
	return &this
}

// NewPaginatedCloudBackupRestoreJobWithDefaults instantiates a new PaginatedCloudBackupRestoreJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudBackupRestoreJobWithDefaults() *PaginatedCloudBackupRestoreJob {
	this := PaginatedCloudBackupRestoreJob{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedCloudBackupRestoreJob) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupRestoreJob) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedCloudBackupRestoreJob) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedCloudBackupRestoreJob) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedCloudBackupRestoreJob) GetResults() []DiskBackupRestoreJob {
	if o == nil || IsNil(o.Results) {
		var ret []DiskBackupRestoreJob
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupRestoreJob) GetResultsOk() ([]DiskBackupRestoreJob, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedCloudBackupRestoreJob) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiskBackupRestoreJob and assigns it to the Results field.
func (o *PaginatedCloudBackupRestoreJob) SetResults(v []DiskBackupRestoreJob) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedCloudBackupRestoreJob) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupRestoreJob) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedCloudBackupRestoreJob) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedCloudBackupRestoreJob) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedCloudBackupRestoreJob) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedCloudBackupRestoreJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedCloudBackupRestoreJob struct {
	value *PaginatedCloudBackupRestoreJob
	isSet bool
}

func (v NullablePaginatedCloudBackupRestoreJob) Get() *PaginatedCloudBackupRestoreJob {
	return v.value
}

func (v *NullablePaginatedCloudBackupRestoreJob) Set(val *PaginatedCloudBackupRestoreJob) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudBackupRestoreJob) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudBackupRestoreJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudBackupRestoreJob(val *PaginatedCloudBackupRestoreJob) *NullablePaginatedCloudBackupRestoreJob {
	return &NullablePaginatedCloudBackupRestoreJob{value: val, isSet: true}
}

func (v NullablePaginatedCloudBackupRestoreJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudBackupRestoreJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}