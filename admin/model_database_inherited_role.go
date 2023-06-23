// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DatabaseInheritedRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseInheritedRole{}

// DatabaseInheritedRole Role inherited from another context for this database user.
type DatabaseInheritedRole struct {
	// Human-readable label that identifies the database on which someone grants the action to one MongoDB user.
	Db string `json:"db"`
	// Human-readable label that identifies the role inherited. Set this value to `admin` for every role except `read` or `readWrite`.
	Role string `json:"role"`
}

// NewDatabaseInheritedRole instantiates a new DatabaseInheritedRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseInheritedRole(db string, role string) *DatabaseInheritedRole {
	this := DatabaseInheritedRole{}
	this.Db = db
	this.Role = role
	return &this
}

// NewDatabaseInheritedRoleWithDefaults instantiates a new DatabaseInheritedRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseInheritedRoleWithDefaults() *DatabaseInheritedRole {
	this := DatabaseInheritedRole{}
	return &this
}

// GetDb returns the Db field value
func (o *DatabaseInheritedRole) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *DatabaseInheritedRole) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *DatabaseInheritedRole) SetDb(v string) {
	o.Db = v
}

// GetRole returns the Role field value
func (o *DatabaseInheritedRole) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *DatabaseInheritedRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *DatabaseInheritedRole) SetRole(v string) {
	o.Role = v
}

func (o DatabaseInheritedRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseInheritedRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["db"] = o.Db
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableDatabaseInheritedRole struct {
	value *DatabaseInheritedRole
	isSet bool
}

func (v NullableDatabaseInheritedRole) Get() *DatabaseInheritedRole {
	return v.value
}

func (v *NullableDatabaseInheritedRole) Set(val *DatabaseInheritedRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseInheritedRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseInheritedRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseInheritedRole(val *DatabaseInheritedRole) *NullableDatabaseInheritedRole {
	return &NullableDatabaseInheritedRole{value: val, isSet: true}
}

func (v NullableDatabaseInheritedRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseInheritedRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}