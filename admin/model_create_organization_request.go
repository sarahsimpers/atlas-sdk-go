// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the CreateOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationRequest{}

// CreateOrganizationRequest struct for CreateOrganizationRequest
type CreateOrganizationRequest struct {
	ApiKey *CreateApiKey `json:"apiKey,omitempty"`
	// Human-readable label that identifies the organization.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the Atlas user that you want to assign the Organization Owner role. This user must be a member of the same organization as the calling API key. This is only required when authenticating with Programmatic API Keys.
	OrgOwnerId *string `json:"orgOwnerId,omitempty"`
}

// NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationRequest(name string) *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	this.Name = name
	return &this
}

// NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetApiKey() CreateApiKey {
	if o == nil || IsNil(o.ApiKey) {
		var ret CreateApiKey
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetApiKeyOk() (*CreateApiKey, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given CreateApiKey and assigns it to the ApiKey field.
func (o *CreateOrganizationRequest) SetApiKey(v CreateApiKey) {
	o.ApiKey = &v
}

// GetName returns the Name field value
func (o *CreateOrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationRequest) SetName(v string) {
	o.Name = v
}

// GetOrgOwnerId returns the OrgOwnerId field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetOrgOwnerId() string {
	if o == nil || IsNil(o.OrgOwnerId) {
		var ret string
		return ret
	}
	return *o.OrgOwnerId
}

// GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetOrgOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgOwnerId) {
		return nil, false
	}
	return o.OrgOwnerId, true
}

// HasOrgOwnerId returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasOrgOwnerId() bool {
	if o != nil && !IsNil(o.OrgOwnerId) {
		return true
	}

	return false
}

// SetOrgOwnerId gets a reference to the given string and assigns it to the OrgOwnerId field.
func (o *CreateOrganizationRequest) SetOrgOwnerId(v string) {
	o.OrgOwnerId = &v
}

func (o CreateOrganizationRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgOwnerId) {
		toSerialize["orgOwnerId"] = o.OrgOwnerId
	}
	return toSerialize, nil
}

type NullableCreateOrganizationRequest struct {
	value *CreateOrganizationRequest
	isSet bool
}

func (v NullableCreateOrganizationRequest) Get() *CreateOrganizationRequest {
	return v.value
}

func (v *NullableCreateOrganizationRequest) Set(val *CreateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationRequest(val *CreateOrganizationRequest) *NullableCreateOrganizationRequest {
	return &NullableCreateOrganizationRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
