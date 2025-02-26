// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DBAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DBAction{}

// DBAction Privilege action that the role grants.
type DBAction struct {
	// Human-readable label that identifies the privilege action.
	Action string `json:"action"`
	// List of resources on which you grant the action.
	Resources []DBResource `json:"resources,omitempty"`
}

// NewDBAction instantiates a new DBAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBAction(action string) *DBAction {
	this := DBAction{}
	this.Action = action
	return &this
}

// NewDBActionWithDefaults instantiates a new DBAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBActionWithDefaults() *DBAction {
	this := DBAction{}
	return &this
}

// GetAction returns the Action field value
func (o *DBAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DBAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *DBAction) SetAction(v string) {
	o.Action = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DBAction) GetResources() []DBResource {
	if o == nil || IsNil(o.Resources) {
		var ret []DBResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBAction) GetResourcesOk() ([]DBResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DBAction) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []DBResource and assigns it to the Resources field.
func (o *DBAction) SetResources(v []DBResource) {
	o.Resources = v
}

func (o DBAction) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DBAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableDBAction struct {
	value *DBAction
	isSet bool
}

func (v NullableDBAction) Get() *DBAction {
	return v.value
}

func (v *NullableDBAction) Set(val *DBAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDBAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDBAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBAction(val *DBAction) *NullableDBAction {
	return &NullableDBAction{value: val, isSet: true}
}

func (v NullableDBAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
