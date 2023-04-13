/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the DataProtectionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataProtectionSettings{}

// DataProtectionSettings struct for DataProtectionSettings
type DataProtectionSettings struct {
	// Email address of the user who authorized to updated the Backup Compliance Policy  settings.
	AuthorizedEmail *string `json:"authorizedEmail,omitempty"`
	// Flag that indicates whether to enable additional backup copies for the cluster. If unspecified, this value defaults to false.
	CopyProtectionEnabled *bool `json:"copyProtectionEnabled,omitempty"`
	// Flag that indicates whether Encryption at Rest using Customer Key  Management is required for all clusters with a Backup Compliance Policy. If unspecified, this value defaults to false.
	EncryptionAtRestEnabled *bool `json:"encryptionAtRestEnabled,omitempty"`
	OnDemandPolicyItem *PolicyItem `json:"onDemandPolicyItem,omitempty"`
	// Flag that indicates whether the cluster uses Continuous Cloud Backups with a Backup Compliance Policy. If unspecified, this value defaults to false.
	PitEnabled *bool `json:"pitEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project for the Backup Compliance Policy.
	ProjectId *string `json:"projectId,omitempty"`
	// Number of previous days that you can restore back to with Continuous Cloud Backup with a Backup Compliance Policy. You must specify a positive, non-zero integer, and the maximum retention window can't exceed the hourly retention time. This parameter applies only to Continuous Cloud Backups with a Backup Compliance Policy.
	RestoreWindowDays *int32 `json:"restoreWindowDays,omitempty"`
	// List that contains the specifications for one scheduled policy.
	ScheduledPolicyItems []PolicyItem `json:"scheduledPolicyItems,omitempty"`
	// Label that indicates the state of the Backup Compliance Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings.
	State *string `json:"state,omitempty"`
	// ISO 8601 timestamp format in UTC that indicates when the user updated the Data Protection Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings.
	UpdatedDate *time.Time `json:"updatedDate,omitempty"`
	// Email address that identifies the user who updated the Backup Compliance Policy settings. MongoDB Cloud ignores this email setting when you enable or update the Backup Compliance Policy settings.
	UpdatedUser *string `json:"updatedUser,omitempty"`
}

// NewDataProtectionSettings instantiates a new DataProtectionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataProtectionSettings() *DataProtectionSettings {
	this := DataProtectionSettings{}
	var copyProtectionEnabled bool = false
	this.CopyProtectionEnabled = &copyProtectionEnabled
	var encryptionAtRestEnabled bool = false
	this.EncryptionAtRestEnabled = &encryptionAtRestEnabled
	var pitEnabled bool = false
	this.PitEnabled = &pitEnabled
	return &this
}

// NewDataProtectionSettingsWithDefaults instantiates a new DataProtectionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataProtectionSettingsWithDefaults() *DataProtectionSettings {
	this := DataProtectionSettings{}
	var copyProtectionEnabled bool = false
	this.CopyProtectionEnabled = &copyProtectionEnabled
	var encryptionAtRestEnabled bool = false
	this.EncryptionAtRestEnabled = &encryptionAtRestEnabled
	var pitEnabled bool = false
	this.PitEnabled = &pitEnabled
	return &this
}

// GetAuthorizedEmail returns the AuthorizedEmail field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetAuthorizedEmail() string {
	if o == nil || IsNil(o.AuthorizedEmail) {
		var ret string
		return ret
	}
	return *o.AuthorizedEmail
}

// GetAuthorizedEmailOk returns a tuple with the AuthorizedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetAuthorizedEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizedEmail) {
		return nil, false
	}
	return o.AuthorizedEmail, true
}

// HasAuthorizedEmail returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasAuthorizedEmail() bool {
	if o != nil && !IsNil(o.AuthorizedEmail) {
		return true
	}

	return false
}

// SetAuthorizedEmail gets a reference to the given string and assigns it to the AuthorizedEmail field.
func (o *DataProtectionSettings) SetAuthorizedEmail(v string) {
	o.AuthorizedEmail = &v
}

// GetCopyProtectionEnabled returns the CopyProtectionEnabled field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetCopyProtectionEnabled() bool {
	if o == nil || IsNil(o.CopyProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.CopyProtectionEnabled
}

// GetCopyProtectionEnabledOk returns a tuple with the CopyProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetCopyProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyProtectionEnabled) {
		return nil, false
	}
	return o.CopyProtectionEnabled, true
}

// HasCopyProtectionEnabled returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasCopyProtectionEnabled() bool {
	if o != nil && !IsNil(o.CopyProtectionEnabled) {
		return true
	}

	return false
}

// SetCopyProtectionEnabled gets a reference to the given bool and assigns it to the CopyProtectionEnabled field.
func (o *DataProtectionSettings) SetCopyProtectionEnabled(v bool) {
	o.CopyProtectionEnabled = &v
}

// GetEncryptionAtRestEnabled returns the EncryptionAtRestEnabled field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetEncryptionAtRestEnabled() bool {
	if o == nil || IsNil(o.EncryptionAtRestEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionAtRestEnabled
}

// GetEncryptionAtRestEnabledOk returns a tuple with the EncryptionAtRestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetEncryptionAtRestEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptionAtRestEnabled) {
		return nil, false
	}
	return o.EncryptionAtRestEnabled, true
}

// HasEncryptionAtRestEnabled returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasEncryptionAtRestEnabled() bool {
	if o != nil && !IsNil(o.EncryptionAtRestEnabled) {
		return true
	}

	return false
}

// SetEncryptionAtRestEnabled gets a reference to the given bool and assigns it to the EncryptionAtRestEnabled field.
func (o *DataProtectionSettings) SetEncryptionAtRestEnabled(v bool) {
	o.EncryptionAtRestEnabled = &v
}

// GetOnDemandPolicyItem returns the OnDemandPolicyItem field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetOnDemandPolicyItem() PolicyItem {
	if o == nil || IsNil(o.OnDemandPolicyItem) {
		var ret PolicyItem
		return ret
	}
	return *o.OnDemandPolicyItem
}

// GetOnDemandPolicyItemOk returns a tuple with the OnDemandPolicyItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetOnDemandPolicyItemOk() (*PolicyItem, bool) {
	if o == nil || IsNil(o.OnDemandPolicyItem) {
		return nil, false
	}
	return o.OnDemandPolicyItem, true
}

// HasOnDemandPolicyItem returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasOnDemandPolicyItem() bool {
	if o != nil && !IsNil(o.OnDemandPolicyItem) {
		return true
	}

	return false
}

// SetOnDemandPolicyItem gets a reference to the given PolicyItem and assigns it to the OnDemandPolicyItem field.
func (o *DataProtectionSettings) SetOnDemandPolicyItem(v PolicyItem) {
	o.OnDemandPolicyItem = &v
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}
	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *DataProtectionSettings) SetPitEnabled(v bool) {
	o.PitEnabled = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DataProtectionSettings) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetRestoreWindowDays() int32 {
	if o == nil || IsNil(o.RestoreWindowDays) {
		var ret int32
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetRestoreWindowDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.RestoreWindowDays) {
		return nil, false
	}
	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasRestoreWindowDays() bool {
	if o != nil && !IsNil(o.RestoreWindowDays) {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int32 and assigns it to the RestoreWindowDays field.
func (o *DataProtectionSettings) SetRestoreWindowDays(v int32) {
	o.RestoreWindowDays = &v
}

// GetScheduledPolicyItems returns the ScheduledPolicyItems field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetScheduledPolicyItems() []PolicyItem {
	if o == nil || IsNil(o.ScheduledPolicyItems) {
		var ret []PolicyItem
		return ret
	}
	return o.ScheduledPolicyItems
}

// GetScheduledPolicyItemsOk returns a tuple with the ScheduledPolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetScheduledPolicyItemsOk() ([]PolicyItem, bool) {
	if o == nil || IsNil(o.ScheduledPolicyItems) {
		return nil, false
	}
	return o.ScheduledPolicyItems, true
}

// HasScheduledPolicyItems returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasScheduledPolicyItems() bool {
	if o != nil && !IsNil(o.ScheduledPolicyItems) {
		return true
	}

	return false
}

// SetScheduledPolicyItems gets a reference to the given []PolicyItem and assigns it to the ScheduledPolicyItems field.
func (o *DataProtectionSettings) SetScheduledPolicyItems(v []PolicyItem) {
	o.ScheduledPolicyItems = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DataProtectionSettings) SetState(v string) {
	o.State = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetUpdatedDate() time.Time {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given time.Time and assigns it to the UpdatedDate field.
func (o *DataProtectionSettings) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = &v
}

// GetUpdatedUser returns the UpdatedUser field value if set, zero value otherwise.
func (o *DataProtectionSettings) GetUpdatedUser() string {
	if o == nil || IsNil(o.UpdatedUser) {
		var ret string
		return ret
	}
	return *o.UpdatedUser
}

// GetUpdatedUserOk returns a tuple with the UpdatedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSettings) GetUpdatedUserOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedUser) {
		return nil, false
	}
	return o.UpdatedUser, true
}

// HasUpdatedUser returns a boolean if a field has been set.
func (o *DataProtectionSettings) HasUpdatedUser() bool {
	if o != nil && !IsNil(o.UpdatedUser) {
		return true
	}

	return false
}

// SetUpdatedUser gets a reference to the given string and assigns it to the UpdatedUser field.
func (o *DataProtectionSettings) SetUpdatedUser(v string) {
	o.UpdatedUser = &v
}

func (o DataProtectionSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataProtectionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyProtectionEnabled) {
		toSerialize["copyProtectionEnabled"] = o.CopyProtectionEnabled
	}
	if !IsNil(o.EncryptionAtRestEnabled) {
		toSerialize["encryptionAtRestEnabled"] = o.EncryptionAtRestEnabled
	}
	if !IsNil(o.OnDemandPolicyItem) {
		toSerialize["onDemandPolicyItem"] = o.OnDemandPolicyItem
	}
	if !IsNil(o.PitEnabled) {
		toSerialize["pitEnabled"] = o.PitEnabled
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.RestoreWindowDays) {
		toSerialize["restoreWindowDays"] = o.RestoreWindowDays
	}
	if !IsNil(o.ScheduledPolicyItems) {
		toSerialize["scheduledPolicyItems"] = o.ScheduledPolicyItems
	}
	return toSerialize, nil
}

type NullableDataProtectionSettings struct {
	value *DataProtectionSettings
	isSet bool
}

func (v NullableDataProtectionSettings) Get() *DataProtectionSettings {
	return v.value
}

func (v *NullableDataProtectionSettings) Set(val *DataProtectionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDataProtectionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDataProtectionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataProtectionSettings(val *DataProtectionSettings) *NullableDataProtectionSettings {
	return &NullableDataProtectionSettings{value: val, isSet: true}
}

func (v NullableDataProtectionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataProtectionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


