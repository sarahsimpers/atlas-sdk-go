// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the OrganizationInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInvitationRequest{}

// OrganizationInvitationRequest struct for OrganizationInvitationRequest
type OrganizationInvitationRequest struct {
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// List of teams to which you want to invite the desired MongoDB Cloud user.
	TeamIds []string `json:"teamIds,omitempty"`
	// Email address that belongs to the desired MongoDB Cloud user.
	Username *string `json:"username,omitempty"`
}

// NewOrganizationInvitationRequest instantiates a new OrganizationInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationRequest() *OrganizationInvitationRequest {
	this := OrganizationInvitationRequest{}
	return &this
}

// NewOrganizationInvitationRequestWithDefaults instantiates a new OrganizationInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationRequestWithDefaults() *OrganizationInvitationRequest {
	this := OrganizationInvitationRequest{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *OrganizationInvitationRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *OrganizationInvitationRequest) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationRequest) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *OrganizationInvitationRequest) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *OrganizationInvitationRequest) SetTeamIds(v []string) {
	o.TeamIds = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OrganizationInvitationRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OrganizationInvitationRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OrganizationInvitationRequest) SetUsername(v string) {
	o.Username = &v
}

func (o OrganizationInvitationRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrganizationInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.TeamIds) {
		toSerialize["teamIds"] = o.TeamIds
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableOrganizationInvitationRequest struct {
	value *OrganizationInvitationRequest
	isSet bool
}

func (v NullableOrganizationInvitationRequest) Get() *OrganizationInvitationRequest {
	return v.value
}

func (v *NullableOrganizationInvitationRequest) Set(val *OrganizationInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitationRequest(val *OrganizationInvitationRequest) *NullableOrganizationInvitationRequest {
	return &NullableOrganizationInvitationRequest{value: val, isSet: true}
}

func (v NullableOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
