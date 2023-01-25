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

// ApiUserRoleAssignment struct for ApiUserRoleAssignment
type ApiUserRoleAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the organization API key.
	ApiUserId *string `json:"apiUserId,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles []string `json:"roles,omitempty"`
}

// NewApiUserRoleAssignment instantiates a new ApiUserRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserRoleAssignment() *ApiUserRoleAssignment {
	this := ApiUserRoleAssignment{}
	return &this
}

// NewApiUserRoleAssignmentWithDefaults instantiates a new ApiUserRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserRoleAssignmentWithDefaults() *ApiUserRoleAssignment {
	this := ApiUserRoleAssignment{}
	return &this
}

// GetApiUserId returns the ApiUserId field value if set, zero value otherwise.
func (o *ApiUserRoleAssignment) GetApiUserId() string {
	if o == nil || o.ApiUserId == nil {
		var ret string
		return ret
	}
	return *o.ApiUserId
}

// GetApiUserIdOk returns a tuple with the ApiUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserRoleAssignment) GetApiUserIdOk() (*string, bool) {
	if o == nil || o.ApiUserId == nil {
		return nil, false
	}
	return o.ApiUserId, true
}

// HasApiUserId returns a boolean if a field has been set.
func (o *ApiUserRoleAssignment) HasApiUserId() bool {
	if o != nil && o.ApiUserId != nil {
		return true
	}

	return false
}

// SetApiUserId gets a reference to the given string and assigns it to the ApiUserId field.
func (o *ApiUserRoleAssignment) SetApiUserId(v string) {
	o.ApiUserId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiUserRoleAssignment) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserRoleAssignment) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiUserRoleAssignment) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiUserRoleAssignment) SetRoles(v []string) {
	o.Roles = v
}

func (o ApiUserRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiUserId != nil {
		toSerialize["apiUserId"] = o.ApiUserId
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableApiUserRoleAssignment struct {
	value *ApiUserRoleAssignment
	isSet bool
}

func (v NullableApiUserRoleAssignment) Get() *ApiUserRoleAssignment {
	return v.value
}

func (v *NullableApiUserRoleAssignment) Set(val *ApiUserRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserRoleAssignment(val *ApiUserRoleAssignment) *NullableApiUserRoleAssignment {
	return &NullableApiUserRoleAssignment{value: val, isSet: true}
}

func (v NullableApiUserRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


