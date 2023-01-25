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

// SampleDatasetStatus struct for SampleDatasetStatus
type SampleDatasetStatus struct {
	// Unique 24-hexadecimal character string that identifies this sample dataset.
	Id *string `json:"_id,omitempty"`
	// Human-readable label that identifies the cluster into which you loaded the sample dataset.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when the sample dataset load job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	CompleteDate *time.Time `json:"completeDate,omitempty"`
	// Date and time when you started the sample dataset load job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// Details of the error returned when MongoDB Cloud loads the sample dataset. This endpoint returns null if state has a value other than FAILED.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Status of the sample dataset load job.
	State *string `json:"state,omitempty"`
}

// NewSampleDatasetStatus instantiates a new SampleDatasetStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleDatasetStatus() *SampleDatasetStatus {
	this := SampleDatasetStatus{}
	return &this
}

// NewSampleDatasetStatusWithDefaults instantiates a new SampleDatasetStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleDatasetStatusWithDefaults() *SampleDatasetStatus {
	this := SampleDatasetStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SampleDatasetStatus) SetId(v string) {
	o.Id = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *SampleDatasetStatus) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCompleteDate returns the CompleteDate field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetCompleteDate() time.Time {
	if o == nil || o.CompleteDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CompleteDate
}

// GetCompleteDateOk returns a tuple with the CompleteDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetCompleteDateOk() (*time.Time, bool) {
	if o == nil || o.CompleteDate == nil {
		return nil, false
	}
	return o.CompleteDate, true
}

// HasCompleteDate returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasCompleteDate() bool {
	if o != nil && o.CompleteDate != nil {
		return true
	}

	return false
}

// SetCompleteDate gets a reference to the given time.Time and assigns it to the CompleteDate field.
func (o *SampleDatasetStatus) SetCompleteDate(v time.Time) {
	o.CompleteDate = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetCreateDate() time.Time {
	if o == nil || o.CreateDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || o.CreateDate == nil {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasCreateDate() bool {
	if o != nil && o.CreateDate != nil {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *SampleDatasetStatus) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SampleDatasetStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SampleDatasetStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleDatasetStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SampleDatasetStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SampleDatasetStatus) SetState(v string) {
	o.State = &v
}

func (o SampleDatasetStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.CompleteDate != nil {
		toSerialize["completeDate"] = o.CompleteDate
	}
	if o.CreateDate != nil {
		toSerialize["createDate"] = o.CreateDate
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableSampleDatasetStatus struct {
	value *SampleDatasetStatus
	isSet bool
}

func (v NullableSampleDatasetStatus) Get() *SampleDatasetStatus {
	return v.value
}

func (v *NullableSampleDatasetStatus) Set(val *SampleDatasetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleDatasetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleDatasetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleDatasetStatus(val *SampleDatasetStatus) *NullableSampleDatasetStatus {
	return &NullableSampleDatasetStatus{value: val, isSet: true}
}

func (v NullableSampleDatasetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleDatasetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


