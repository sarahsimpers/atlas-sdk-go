/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the ApiKeyView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyView{}

// ApiKeyView Details contained in one API key.
type ApiKeyView struct {
	// List of network addresses granted access to this API using this API key.
	AccessList []AccessListItemView `json:"accessList,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key.
	Id string `json:"id"`
	// Public API key value set for the specified organization API key.
	PublicKey string `json:"publicKey"`
	// List that contains roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the Cloud user.
	Roles []ApiRoleAssignmentView `json:"roles,omitempty"`
}

// NewApiKeyView instantiates a new ApiKeyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyView() *ApiKeyView {
	this := ApiKeyView{}
	return &this
}

// NewApiKeyViewWithDefaults instantiates a new ApiKeyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyViewWithDefaults() *ApiKeyView {
	this := ApiKeyView{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *ApiKeyView) GetAccessList() []AccessListItemView {
	if o == nil || IsNil(o.AccessList) {
		var ret []AccessListItemView
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyView) GetAccessListOk() ([]AccessListItemView, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ApiKeyView) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []AccessListItemView and assigns it to the AccessList field.
func (o *ApiKeyView) SetAccessList(v []AccessListItemView) {
	o.AccessList = v
}

// GetId returns the Id field value
func (o *ApiKeyView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKeyView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKeyView) SetId(v string) {
	o.Id = v
}

// GetPublicKey returns the PublicKey field value
func (o *ApiKeyView) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ApiKeyView) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ApiKeyView) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiKeyView) GetRoles() []ApiRoleAssignmentView {
	if o == nil || IsNil(o.Roles) {
		var ret []ApiRoleAssignmentView
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyView) GetRolesOk() ([]ApiRoleAssignmentView, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiKeyView) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ApiRoleAssignmentView and assigns it to the Roles field.
func (o *ApiKeyView) SetRoles(v []ApiRoleAssignmentView) {
	o.Roles = v
}

func (o ApiKeyView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: accessList is readOnly
	// skip: id is readOnly
	// skip: publicKey is readOnly
	// skip: roles is readOnly
	return toSerialize, nil
}

type NullableApiKeyView struct {
	value *ApiKeyView
	isSet bool
}

func (v NullableApiKeyView) Get() *ApiKeyView {
	return v.value
}

func (v *NullableApiKeyView) Set(val *ApiKeyView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyView(val *ApiKeyView) *NullableApiKeyView {
	return &NullableApiKeyView{value: val, isSet: true}
}

func (v NullableApiKeyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


