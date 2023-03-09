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

// checks if the TeamView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamView{}

// TeamView struct for TeamView
type TeamView struct {
	// Unique 24-hexadecimal digit string that identifies this team.
	Id string `json:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Human-readable label that identifies the team.
	Name string `json:"name"`
	// List that contains the MongoDB Cloud users in this team.
	Usernames []string `json:"usernames,omitempty"`
}

// NewTeamView instantiates a new TeamView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamView(id string, name string) *TeamView {
	this := TeamView{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTeamViewWithDefaults instantiates a new TeamView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamViewWithDefaults() *TeamView {
	this := TeamView{}
	return &this
}

// GetId returns the Id field value
func (o *TeamView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamView) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *TeamView) SetLinks(v []Link) {
	o.Links = v
}

// GetName returns the Name field value
func (o *TeamView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamView) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamView) SetName(v string) {
	o.Name = v
}

// GetUsernames returns the Usernames field value if set, zero value otherwise.
func (o *TeamView) GetUsernames() []string {
	if o == nil || IsNil(o.Usernames) {
		var ret []string
		return ret
	}
	return o.Usernames
}

// GetUsernamesOk returns a tuple with the Usernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamView) GetUsernamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Usernames) {
		return nil, false
	}
	return o.Usernames, true
}

// HasUsernames returns a boolean if a field has been set.
func (o *TeamView) HasUsernames() bool {
	if o != nil && !IsNil(o.Usernames) {
		return true
	}

	return false
}

// SetUsernames gets a reference to the given []string and assigns it to the Usernames field.
func (o *TeamView) SetUsernames(v []string) {
	o.Usernames = v
}

func (o TeamView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: links is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.Usernames) {
		toSerialize["usernames"] = o.Usernames
	}
	return toSerialize, nil
}

type NullableTeamView struct {
	value *TeamView
	isSet bool
}

func (v NullableTeamView) Get() *TeamView {
	return v.value
}

func (v *NullableTeamView) Set(val *TeamView) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamView) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamView(val *TeamView) *NullableTeamView {
	return &NullableTeamView{value: val, isSet: true}
}

func (v NullableTeamView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

