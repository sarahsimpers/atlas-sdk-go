// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the DatabaseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseUser{}

// DatabaseUser struct for DatabaseUser
type DatabaseUser struct {
	// Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.
	AwsIAMType *string `json:"awsIAMType,omitempty"`
	// Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.
	DatabaseName string `json:"databaseName"`
	// Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.
	DeleteAfterDate *time.Time `json:"deleteAfterDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project.
	GroupId string `json:"groupId"`
	// List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console.
	Labels []NDSLabel `json:"labels,omitempty"`
	// Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host.
	LdapAuthType *string `json:"ldapAuthType,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Alphanumeric string that authenticates this database user against the database specified in `databaseName`. To authenticate with SCRAM-SHA, you must specify this parameter. This parameter doesn't appear in this response.
	Password *string `json:"password,omitempty"`
	// List that provides the pairings of one role with one applicable database.
	Roles []Role `json:"roles,omitempty"`
	// List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.
	Scopes []UserScope `json:"scopes,omitempty"`
	// Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | AWS IAM | awsType | USER | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
	Username string `json:"username"`
	// X.509 method that MongoDB Cloud uses to authenticate the database user.  - For application-managed X.509, specify `MANAGED`. - For self-managed X.509, specify `CUSTOMER`.  Users created with the `CUSTOMER` method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the `$external` database.
	X509Type *string `json:"x509Type,omitempty"`
}

// NewDatabaseUser instantiates a new DatabaseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUser(databaseName string, groupId string, username string) *DatabaseUser {
	this := DatabaseUser{}
	var awsIAMType string = "NONE"
	this.AwsIAMType = &awsIAMType
	this.DatabaseName = databaseName
	this.GroupId = groupId
	var ldapAuthType string = "NONE"
	this.LdapAuthType = &ldapAuthType
	this.Username = username
	var x509Type string = "NONE"
	this.X509Type = &x509Type
	return &this
}

// NewDatabaseUserWithDefaults instantiates a new DatabaseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUserWithDefaults() *DatabaseUser {
	this := DatabaseUser{}
	var awsIAMType string = "NONE"
	this.AwsIAMType = &awsIAMType
	var databaseName string = "admin"
	this.DatabaseName = databaseName
	var ldapAuthType string = "NONE"
	this.LdapAuthType = &ldapAuthType
	var x509Type string = "NONE"
	this.X509Type = &x509Type
	return &this
}

// GetAwsIAMType returns the AwsIAMType field value if set, zero value otherwise.
func (o *DatabaseUser) GetAwsIAMType() string {
	if o == nil || IsNil(o.AwsIAMType) {
		var ret string
		return ret
	}
	return *o.AwsIAMType
}

// GetAwsIAMTypeOk returns a tuple with the AwsIAMType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetAwsIAMTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AwsIAMType) {
		return nil, false
	}
	return o.AwsIAMType, true
}

// HasAwsIAMType returns a boolean if a field has been set.
func (o *DatabaseUser) HasAwsIAMType() bool {
	if o != nil && !IsNil(o.AwsIAMType) {
		return true
	}

	return false
}

// SetAwsIAMType gets a reference to the given string and assigns it to the AwsIAMType field.
func (o *DatabaseUser) SetAwsIAMType(v string) {
	o.AwsIAMType = &v
}

// GetDatabaseName returns the DatabaseName field value
func (o *DatabaseUser) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *DatabaseUser) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetDeleteAfterDate returns the DeleteAfterDate field value if set, zero value otherwise.
func (o *DatabaseUser) GetDeleteAfterDate() time.Time {
	if o == nil || IsNil(o.DeleteAfterDate) {
		var ret time.Time
		return ret
	}
	return *o.DeleteAfterDate
}

// GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetDeleteAfterDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeleteAfterDate) {
		return nil, false
	}
	return o.DeleteAfterDate, true
}

// HasDeleteAfterDate returns a boolean if a field has been set.
func (o *DatabaseUser) HasDeleteAfterDate() bool {
	if o != nil && !IsNil(o.DeleteAfterDate) {
		return true
	}

	return false
}

// SetDeleteAfterDate gets a reference to the given time.Time and assigns it to the DeleteAfterDate field.
func (o *DatabaseUser) SetDeleteAfterDate(v time.Time) {
	o.DeleteAfterDate = &v
}

// GetGroupId returns the GroupId field value
func (o *DatabaseUser) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *DatabaseUser) SetGroupId(v string) {
	o.GroupId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DatabaseUser) GetLabels() []NDSLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []NDSLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetLabelsOk() ([]NDSLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DatabaseUser) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []NDSLabel and assigns it to the Labels field.
func (o *DatabaseUser) SetLabels(v []NDSLabel) {
	o.Labels = v
}

// GetLdapAuthType returns the LdapAuthType field value if set, zero value otherwise.
func (o *DatabaseUser) GetLdapAuthType() string {
	if o == nil || IsNil(o.LdapAuthType) {
		var ret string
		return ret
	}
	return *o.LdapAuthType
}

// GetLdapAuthTypeOk returns a tuple with the LdapAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetLdapAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LdapAuthType) {
		return nil, false
	}
	return o.LdapAuthType, true
}

// HasLdapAuthType returns a boolean if a field has been set.
func (o *DatabaseUser) HasLdapAuthType() bool {
	if o != nil && !IsNil(o.LdapAuthType) {
		return true
	}

	return false
}

// SetLdapAuthType gets a reference to the given string and assigns it to the LdapAuthType field.
func (o *DatabaseUser) SetLdapAuthType(v string) {
	o.LdapAuthType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DatabaseUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DatabaseUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DatabaseUser) SetLinks(v []Link) {
	o.Links = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DatabaseUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DatabaseUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DatabaseUser) SetPassword(v string) {
	o.Password = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *DatabaseUser) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *DatabaseUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *DatabaseUser) SetRoles(v []Role) {
	o.Roles = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DatabaseUser) GetScopes() []UserScope {
	if o == nil || IsNil(o.Scopes) {
		var ret []UserScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetScopesOk() ([]UserScope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DatabaseUser) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []UserScope and assigns it to the Scopes field.
func (o *DatabaseUser) SetScopes(v []UserScope) {
	o.Scopes = v
}

// GetUsername returns the Username field value
func (o *DatabaseUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *DatabaseUser) SetUsername(v string) {
	o.Username = v
}

// GetX509Type returns the X509Type field value if set, zero value otherwise.
func (o *DatabaseUser) GetX509Type() string {
	if o == nil || IsNil(o.X509Type) {
		var ret string
		return ret
	}
	return *o.X509Type
}

// GetX509TypeOk returns a tuple with the X509Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUser) GetX509TypeOk() (*string, bool) {
	if o == nil || IsNil(o.X509Type) {
		return nil, false
	}
	return o.X509Type, true
}

// HasX509Type returns a boolean if a field has been set.
func (o *DatabaseUser) HasX509Type() bool {
	if o != nil && !IsNil(o.X509Type) {
		return true
	}

	return false
}

// SetX509Type gets a reference to the given string and assigns it to the X509Type field.
func (o *DatabaseUser) SetX509Type(v string) {
	o.X509Type = &v
}

func (o DatabaseUser) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsIAMType) {
		toSerialize["awsIAMType"] = o.AwsIAMType
	}
	toSerialize["databaseName"] = o.DatabaseName
	if !IsNil(o.DeleteAfterDate) {
		toSerialize["deleteAfterDate"] = o.DeleteAfterDate
	}
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.LdapAuthType) {
		toSerialize["ldapAuthType"] = o.LdapAuthType
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["username"] = o.Username
	if !IsNil(o.X509Type) {
		toSerialize["x509Type"] = o.X509Type
	}
	return toSerialize, nil
}

type NullableDatabaseUser struct {
	value *DatabaseUser
	isSet bool
}

func (v NullableDatabaseUser) Get() *DatabaseUser {
	return v.value
}

func (v *NullableDatabaseUser) Set(val *DatabaseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseUser(val *DatabaseUser) *NullableDatabaseUser {
	return &NullableDatabaseUser{value: val, isSet: true}
}

func (v NullableDatabaseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
