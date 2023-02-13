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

// NDSServerlessInstanceAuditView NDS serverless instance audit identifies any activities around serverless instance.
type NDSServerlessInstanceAuditView struct {
	// Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	EventTypeName NDSServerlessInstanceAuditTypeView `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
}

// NewNDSServerlessInstanceAuditView instantiates a new NDSServerlessInstanceAuditView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSServerlessInstanceAuditView() *NDSServerlessInstanceAuditView {
	this := NDSServerlessInstanceAuditView{}
	return &this
}

// NewNDSServerlessInstanceAuditViewWithDefaults instantiates a new NDSServerlessInstanceAuditView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSServerlessInstanceAuditViewWithDefaults() *NDSServerlessInstanceAuditView {
	this := NDSServerlessInstanceAuditView{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetApiKeyId() string {
	if o == nil || o.ApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetApiKeyIdOk() (*string, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasApiKeyId() bool {
	if o != nil && o.ApiKeyId != nil {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *NDSServerlessInstanceAuditView) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value
func (o *NDSServerlessInstanceAuditView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *NDSServerlessInstanceAuditView) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *NDSServerlessInstanceAuditView) GetEventTypeName() NDSServerlessInstanceAuditTypeView {
	if o == nil {
		var ret NDSServerlessInstanceAuditTypeView
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetEventTypeNameOk() (*NDSServerlessInstanceAuditTypeView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *NDSServerlessInstanceAuditView) SetEventTypeName(v NDSServerlessInstanceAuditTypeView) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *NDSServerlessInstanceAuditView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *NDSServerlessInstanceAuditView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NDSServerlessInstanceAuditView) SetId(v string) {
	o.Id = v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetIsGlobalAdmin() bool {
	if o == nil || o.IsGlobalAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || o.IsGlobalAdmin == nil {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasIsGlobalAdmin() bool {
	if o != nil && o.IsGlobalAdmin != nil {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *NDSServerlessInstanceAuditView) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value
func (o *NDSServerlessInstanceAuditView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *NDSServerlessInstanceAuditView) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *NDSServerlessInstanceAuditView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *NDSServerlessInstanceAuditView) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetRaw() Raw {
	if o == nil || o.Raw == nil {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetRawOk() (*Raw, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *NDSServerlessInstanceAuditView) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *NDSServerlessInstanceAuditView) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NDSServerlessInstanceAuditView) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NDSServerlessInstanceAuditView) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSServerlessInstanceAuditView) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NDSServerlessInstanceAuditView) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NDSServerlessInstanceAuditView) SetUsername(v string) {
	o.Username = &v
}

func (o NDSServerlessInstanceAuditView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKeyId != nil {
		toSerialize["apiKeyId"] = o.ApiKeyId
	}
	if true {
		toSerialize["created"] = o.Created
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
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
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
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableNDSServerlessInstanceAuditView struct {
	value *NDSServerlessInstanceAuditView
	isSet bool
}

func (v NullableNDSServerlessInstanceAuditView) Get() *NDSServerlessInstanceAuditView {
	return v.value
}

func (v *NullableNDSServerlessInstanceAuditView) Set(val *NDSServerlessInstanceAuditView) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSServerlessInstanceAuditView) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSServerlessInstanceAuditView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSServerlessInstanceAuditView(val *NDSServerlessInstanceAuditView) *NullableNDSServerlessInstanceAuditView {
	return &NullableNDSServerlessInstanceAuditView{value: val, isSet: true}
}

func (v NullableNDSServerlessInstanceAuditView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSServerlessInstanceAuditView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


