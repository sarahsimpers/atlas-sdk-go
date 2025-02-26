// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the UserSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSecurity{}

// UserSecurity struct for UserSecurity
type UserSecurity struct {
	CustomerX509 *CustomerX509 `json:"customerX509,omitempty"`
	Ldap         *NDSLDAP      `json:"ldap,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewUserSecurity instantiates a new UserSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSecurity() *UserSecurity {
	this := UserSecurity{}
	return &this
}

// NewUserSecurityWithDefaults instantiates a new UserSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSecurityWithDefaults() *UserSecurity {
	this := UserSecurity{}
	return &this
}

// GetCustomerX509 returns the CustomerX509 field value if set, zero value otherwise.
func (o *UserSecurity) GetCustomerX509() CustomerX509 {
	if o == nil || IsNil(o.CustomerX509) {
		var ret CustomerX509
		return ret
	}
	return *o.CustomerX509
}

// GetCustomerX509Ok returns a tuple with the CustomerX509 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetCustomerX509Ok() (*CustomerX509, bool) {
	if o == nil || IsNil(o.CustomerX509) {
		return nil, false
	}
	return o.CustomerX509, true
}

// HasCustomerX509 returns a boolean if a field has been set.
func (o *UserSecurity) HasCustomerX509() bool {
	if o != nil && !IsNil(o.CustomerX509) {
		return true
	}

	return false
}

// SetCustomerX509 gets a reference to the given CustomerX509 and assigns it to the CustomerX509 field.
func (o *UserSecurity) SetCustomerX509(v CustomerX509) {
	o.CustomerX509 = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *UserSecurity) GetLdap() NDSLDAP {
	if o == nil || IsNil(o.Ldap) {
		var ret NDSLDAP
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetLdapOk() (*NDSLDAP, bool) {
	if o == nil || IsNil(o.Ldap) {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *UserSecurity) HasLdap() bool {
	if o != nil && !IsNil(o.Ldap) {
		return true
	}

	return false
}

// SetLdap gets a reference to the given NDSLDAP and assigns it to the Ldap field.
func (o *UserSecurity) SetLdap(v NDSLDAP) {
	o.Ldap = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserSecurity) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserSecurity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *UserSecurity) SetLinks(v []Link) {
	o.Links = v
}

func (o UserSecurity) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerX509) {
		toSerialize["customerX509"] = o.CustomerX509
	}
	if !IsNil(o.Ldap) {
		toSerialize["ldap"] = o.Ldap
	}
	return toSerialize, nil
}

type NullableUserSecurity struct {
	value *UserSecurity
	isSet bool
}

func (v NullableUserSecurity) Get() *UserSecurity {
	return v.value
}

func (v *NullableUserSecurity) Set(val *UserSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSecurity(val *UserSecurity) *NullableUserSecurity {
	return &NullableUserSecurity{value: val, isSet: true}
}

func (v NullableUserSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
