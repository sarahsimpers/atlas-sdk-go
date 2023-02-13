/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// TimeMetricEventView struct for TimeMetricEventView
type TimeMetricEventView struct {
	// (**For audit only**), Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	CurrentValue *TimeMetricValueView `json:"currentValue,omitempty"`
	EventTypeName HostMetricEventTypeView `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// (**For audit only**), Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Human-readable label of the metric associated with the **alertId**. This field may change type of **currentValue** field.
	MetricName *string `json:"metricName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// IANA port on which the MongoDB process listens for requests.
	Port *int32 `json:"port,omitempty"`
	// (**For audit only**), Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// (**For audit only**), IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Human-readable label of the replica set associated with the event.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Human-readable label of the shard associated with the event.
	ShardName *string `json:"shardName,omitempty"`
	// (**For audit only**), Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// (**For audit only**), Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
}

// NewTimeMetricEventView instantiates a new TimeMetricEventView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeMetricEventView() *TimeMetricEventView {
	this := TimeMetricEventView{}
	return &this
}

// NewTimeMetricEventViewWithDefaults instantiates a new TimeMetricEventView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeMetricEventViewWithDefaults() *TimeMetricEventView {
	this := TimeMetricEventView{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetApiKeyId() string {
	if o == nil || o.ApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetApiKeyIdOk() (*string, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasApiKeyId() bool {
	if o != nil && o.ApiKeyId != nil {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *TimeMetricEventView) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value
func (o *TimeMetricEventView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TimeMetricEventView) SetCreated(v time.Time) {
	o.Created = v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetCurrentValue() TimeMetricValueView {
	if o == nil || o.CurrentValue == nil {
		var ret TimeMetricValueView
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetCurrentValueOk() (*TimeMetricValueView, bool) {
	if o == nil || o.CurrentValue == nil {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasCurrentValue() bool {
	if o != nil && o.CurrentValue != nil {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given TimeMetricValueView and assigns it to the CurrentValue field.
func (o *TimeMetricEventView) SetCurrentValue(v TimeMetricValueView) {
	o.CurrentValue = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *TimeMetricEventView) GetEventTypeName() HostMetricEventTypeView {
	if o == nil {
		var ret HostMetricEventTypeView
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetEventTypeNameOk() (*HostMetricEventTypeView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *TimeMetricEventView) SetEventTypeName(v HostMetricEventTypeView) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *TimeMetricEventView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *TimeMetricEventView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TimeMetricEventView) SetId(v string) {
	o.Id = v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetIsGlobalAdmin() bool {
	if o == nil || o.IsGlobalAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || o.IsGlobalAdmin == nil {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasIsGlobalAdmin() bool {
	if o != nil && o.IsGlobalAdmin != nil {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *TimeMetricEventView) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value
func (o *TimeMetricEventView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *TimeMetricEventView) SetLinks(v []Link) {
	o.Links = v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *TimeMetricEventView) SetMetricName(v string) {
	o.MetricName = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *TimeMetricEventView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TimeMetricEventView) SetPort(v int32) {
	o.Port = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *TimeMetricEventView) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetRaw() Raw {
	if o == nil || o.Raw == nil {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetRawOk() (*Raw, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *TimeMetricEventView) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *TimeMetricEventView) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetReplicaSetName() string {
	if o == nil || o.ReplicaSetName == nil {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || o.ReplicaSetName == nil {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasReplicaSetName() bool {
	if o != nil && o.ReplicaSetName != nil {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *TimeMetricEventView) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetShardName() string {
	if o == nil || o.ShardName == nil {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetShardNameOk() (*string, bool) {
	if o == nil || o.ShardName == nil {
		return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasShardName() bool {
	if o != nil && o.ShardName != nil {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *TimeMetricEventView) SetShardName(v string) {
	o.ShardName = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *TimeMetricEventView) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TimeMetricEventView) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricEventView) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TimeMetricEventView) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TimeMetricEventView) SetUsername(v string) {
	o.Username = &v
}

func (o TimeMetricEventView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKeyId != nil {
		toSerialize["apiKeyId"] = o.ApiKeyId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.CurrentValue != nil {
		toSerialize["currentValue"] = o.CurrentValue
	}
	if true {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IsGlobalAdmin != nil {
		toSerialize["isGlobalAdmin"] = o.IsGlobalAdmin
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.MetricName != nil {
		toSerialize["metricName"] = o.MetricName
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.RemoteAddress != nil {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if o.ReplicaSetName != nil {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if o.ShardName != nil {
		toSerialize["shardName"] = o.ShardName
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableTimeMetricEventView struct {
	value *TimeMetricEventView
	isSet bool
}

func (v NullableTimeMetricEventView) Get() *TimeMetricEventView {
	return v.value
}

func (v *NullableTimeMetricEventView) Set(val *TimeMetricEventView) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeMetricEventView) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeMetricEventView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeMetricEventView(val *TimeMetricEventView) *NullableTimeMetricEventView {
	return &NullableTimeMetricEventView{value: val, isSet: true}
}

func (v NullableTimeMetricEventView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeMetricEventView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


