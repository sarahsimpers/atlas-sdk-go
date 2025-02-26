// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the OrganizationInvitationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInvitationUpdateRequest{}

// OrganizationInvitationUpdateRequest struct for OrganizationInvitationUpdateRequest
type OrganizationInvitationUpdateRequest struct {
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// List of teams to which you want to invite the desired MongoDB Cloud user.
	TeamIds []string `json:"teamIds,omitempty"`
}

// NewOrganizationInvitationUpdateRequest instantiates a new OrganizationInvitationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationUpdateRequest() *OrganizationInvitationUpdateRequest {
	this := OrganizationInvitationUpdateRequest{}
	return &this
}

// NewOrganizationInvitationUpdateRequestWithDefaults instantiates a new OrganizationInvitationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationUpdateRequestWithDefaults() *OrganizationInvitationUpdateRequest {
	this := OrganizationInvitationUpdateRequest{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *OrganizationInvitationUpdateRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationUpdateRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *OrganizationInvitationUpdateRequest) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationUpdateRequest) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *OrganizationInvitationUpdateRequest) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *OrganizationInvitationUpdateRequest) SetTeamIds(v []string) {
	o.TeamIds = v
}

func (o OrganizationInvitationUpdateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrganizationInvitationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.TeamIds) {
		toSerialize["teamIds"] = o.TeamIds
	}
	return toSerialize, nil
}

type NullableOrganizationInvitationUpdateRequest struct {
	value *OrganizationInvitationUpdateRequest
	isSet bool
}

func (v NullableOrganizationInvitationUpdateRequest) Get() *OrganizationInvitationUpdateRequest {
	return v.value
}

func (v *NullableOrganizationInvitationUpdateRequest) Set(val *OrganizationInvitationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitationUpdateRequest(val *OrganizationInvitationUpdateRequest) *NullableOrganizationInvitationUpdateRequest {
	return &NullableOrganizationInvitationUpdateRequest{value: val, isSet: true}
}

func (v NullableOrganizationInvitationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
