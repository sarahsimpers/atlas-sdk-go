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

// DataLakeAWSCloudProviderConfig Name of the cloud service that hosts the data lake's data stores.
type DataLakeAWSCloudProviderConfig struct {
	// Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores.
	ExternalId *string `json:"externalId,omitempty"`
	// Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores.
	IamAssumedRoleARN *string `json:"iamAssumedRoleARN,omitempty"`
	// Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores.
	IamUserARN *string `json:"iamUserARN,omitempty"`
	// Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig.
	RoleId string `json:"roleId"`
	// Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.
	TestS3Bucket string `json:"testS3Bucket"`
}

// NewDataLakeAWSCloudProviderConfig instantiates a new DataLakeAWSCloudProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeAWSCloudProviderConfig() *DataLakeAWSCloudProviderConfig {
	this := DataLakeAWSCloudProviderConfig{}
	return &this
}

// NewDataLakeAWSCloudProviderConfigWithDefaults instantiates a new DataLakeAWSCloudProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeAWSCloudProviderConfigWithDefaults() *DataLakeAWSCloudProviderConfig {
	this := DataLakeAWSCloudProviderConfig{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *DataLakeAWSCloudProviderConfig) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAWSCloudProviderConfig) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *DataLakeAWSCloudProviderConfig) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *DataLakeAWSCloudProviderConfig) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIamAssumedRoleARN returns the IamAssumedRoleARN field value if set, zero value otherwise.
func (o *DataLakeAWSCloudProviderConfig) GetIamAssumedRoleARN() string {
	if o == nil || o.IamAssumedRoleARN == nil {
		var ret string
		return ret
	}
	return *o.IamAssumedRoleARN
}

// GetIamAssumedRoleARNOk returns a tuple with the IamAssumedRoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAWSCloudProviderConfig) GetIamAssumedRoleARNOk() (*string, bool) {
	if o == nil || o.IamAssumedRoleARN == nil {
		return nil, false
	}
	return o.IamAssumedRoleARN, true
}

// HasIamAssumedRoleARN returns a boolean if a field has been set.
func (o *DataLakeAWSCloudProviderConfig) HasIamAssumedRoleARN() bool {
	if o != nil && o.IamAssumedRoleARN != nil {
		return true
	}

	return false
}

// SetIamAssumedRoleARN gets a reference to the given string and assigns it to the IamAssumedRoleARN field.
func (o *DataLakeAWSCloudProviderConfig) SetIamAssumedRoleARN(v string) {
	o.IamAssumedRoleARN = &v
}

// GetIamUserARN returns the IamUserARN field value if set, zero value otherwise.
func (o *DataLakeAWSCloudProviderConfig) GetIamUserARN() string {
	if o == nil || o.IamUserARN == nil {
		var ret string
		return ret
	}
	return *o.IamUserARN
}

// GetIamUserARNOk returns a tuple with the IamUserARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAWSCloudProviderConfig) GetIamUserARNOk() (*string, bool) {
	if o == nil || o.IamUserARN == nil {
		return nil, false
	}
	return o.IamUserARN, true
}

// HasIamUserARN returns a boolean if a field has been set.
func (o *DataLakeAWSCloudProviderConfig) HasIamUserARN() bool {
	if o != nil && o.IamUserARN != nil {
		return true
	}

	return false
}

// SetIamUserARN gets a reference to the given string and assigns it to the IamUserARN field.
func (o *DataLakeAWSCloudProviderConfig) SetIamUserARN(v string) {
	o.IamUserARN = &v
}

// GetRoleId returns the RoleId field value
func (o *DataLakeAWSCloudProviderConfig) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *DataLakeAWSCloudProviderConfig) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *DataLakeAWSCloudProviderConfig) SetRoleId(v string) {
	o.RoleId = v
}

// GetTestS3Bucket returns the TestS3Bucket field value
func (o *DataLakeAWSCloudProviderConfig) GetTestS3Bucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestS3Bucket
}

// GetTestS3BucketOk returns a tuple with the TestS3Bucket field value
// and a boolean to check if the value has been set.
func (o *DataLakeAWSCloudProviderConfig) GetTestS3BucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestS3Bucket, true
}

// SetTestS3Bucket sets field value
func (o *DataLakeAWSCloudProviderConfig) SetTestS3Bucket(v string) {
	o.TestS3Bucket = v
}

func (o DataLakeAWSCloudProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.IamAssumedRoleARN != nil {
		toSerialize["iamAssumedRoleARN"] = o.IamAssumedRoleARN
	}
	if o.IamUserARN != nil {
		toSerialize["iamUserARN"] = o.IamUserARN
	}
	if true {
		toSerialize["roleId"] = o.RoleId
	}
	if true {
		toSerialize["testS3Bucket"] = o.TestS3Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableDataLakeAWSCloudProviderConfig struct {
	value *DataLakeAWSCloudProviderConfig
	isSet bool
}

func (v NullableDataLakeAWSCloudProviderConfig) Get() *DataLakeAWSCloudProviderConfig {
	return v.value
}

func (v *NullableDataLakeAWSCloudProviderConfig) Set(val *DataLakeAWSCloudProviderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeAWSCloudProviderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeAWSCloudProviderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeAWSCloudProviderConfig(val *DataLakeAWSCloudProviderConfig) *NullableDataLakeAWSCloudProviderConfig {
	return &NullableDataLakeAWSCloudProviderConfig{value: val, isSet: true}
}

func (v NullableDataLakeAWSCloudProviderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeAWSCloudProviderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


