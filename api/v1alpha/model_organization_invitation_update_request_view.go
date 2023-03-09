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

// checks if the OrganizationInvitationUpdateRequestView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInvitationUpdateRequestView{}

// OrganizationInvitationUpdateRequestView struct for OrganizationInvitationUpdateRequestView
type OrganizationInvitationUpdateRequestView struct {
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// List of teams to which you want to invite the desired MongoDB Cloud user.
	TeamIds []string `json:"teamIds,omitempty"`
}

// NewOrganizationInvitationUpdateRequestView instantiates a new OrganizationInvitationUpdateRequestView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationUpdateRequestView() *OrganizationInvitationUpdateRequestView {
	this := OrganizationInvitationUpdateRequestView{}
	return &this
}

// NewOrganizationInvitationUpdateRequestViewWithDefaults instantiates a new OrganizationInvitationUpdateRequestView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationUpdateRequestViewWithDefaults() *OrganizationInvitationUpdateRequestView {
	this := OrganizationInvitationUpdateRequestView{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *OrganizationInvitationUpdateRequestView) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequestView) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequestView) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationUpdateRequestView) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *OrganizationInvitationUpdateRequestView) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequestView) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequestView) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *OrganizationInvitationUpdateRequestView) SetTeamIds(v []string) {
	o.TeamIds = v
}

func (o OrganizationInvitationUpdateRequestView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInvitationUpdateRequestView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.TeamIds) {
		toSerialize["teamIds"] = o.TeamIds
	}
	return toSerialize, nil
}

type NullableOrganizationInvitationUpdateRequestView struct {
	value *OrganizationInvitationUpdateRequestView
	isSet bool
}

func (v NullableOrganizationInvitationUpdateRequestView) Get() *OrganizationInvitationUpdateRequestView {
	return v.value
}

func (v *NullableOrganizationInvitationUpdateRequestView) Set(val *OrganizationInvitationUpdateRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitationUpdateRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitationUpdateRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitationUpdateRequestView(val *OrganizationInvitationUpdateRequestView) *NullableOrganizationInvitationUpdateRequestView {
	return &NullableOrganizationInvitationUpdateRequestView{value: val, isSet: true}
}

func (v NullableOrganizationInvitationUpdateRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitationUpdateRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

