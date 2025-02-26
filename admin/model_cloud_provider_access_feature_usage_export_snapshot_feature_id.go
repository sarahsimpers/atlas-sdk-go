// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAccessFeatureUsageExportSnapshotFeatureId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessFeatureUsageExportSnapshotFeatureId{}

// CloudProviderAccessFeatureUsageExportSnapshotFeatureId Identifying characteristics about the Amazon Web Services (AWS) Simple Storage Service (S3) export bucket linked to this AWS Identity and Access Management (IAM) role.
type CloudProviderAccessFeatureUsageExportSnapshotFeatureId struct {
	// Unique 24-hexadecimal digit string that identifies the AWS S3 bucket to which you export your snapshots.
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies your project.
	GroupId *string `json:"groupId,omitempty"`
}

// NewCloudProviderAccessFeatureUsageExportSnapshotFeatureId instantiates a new CloudProviderAccessFeatureUsageExportSnapshotFeatureId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessFeatureUsageExportSnapshotFeatureId() *CloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	this := CloudProviderAccessFeatureUsageExportSnapshotFeatureId{}
	return &this
}

// NewCloudProviderAccessFeatureUsageExportSnapshotFeatureIdWithDefaults instantiates a new CloudProviderAccessFeatureUsageExportSnapshotFeatureId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessFeatureUsageExportSnapshotFeatureIdWithDefaults() *CloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	this := CloudProviderAccessFeatureUsageExportSnapshotFeatureId{}
	return &this
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetExportBucketId() string {
	if o == nil || IsNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetExportBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportBucketId) {
		return nil, false
	}
	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) HasExportBucketId() bool {
	if o != nil && !IsNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) SetExportBucketId(v string) {
	o.ExportBucketId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) SetGroupId(v string) {
	o.GroupId = &v
}

func (o CloudProviderAccessFeatureUsageExportSnapshotFeatureId) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessFeatureUsageExportSnapshotFeatureId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId struct {
	value *CloudProviderAccessFeatureUsageExportSnapshotFeatureId
	isSet bool
}

func (v NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) Get() *CloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	return v.value
}

func (v *NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) Set(val *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId(val *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) *NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	return &NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId{value: val, isSet: true}
}

func (v NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessFeatureUsageExportSnapshotFeatureId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
