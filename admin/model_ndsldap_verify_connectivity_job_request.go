// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the NDSLDAPVerifyConnectivityJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NDSLDAPVerifyConnectivityJobRequest{}

// NDSLDAPVerifyConnectivityJobRequest struct for NDSLDAPVerifyConnectivityJobRequest
type NDSLDAPVerifyConnectivityJobRequest struct {
	// Unique 24-hexadecimal digit string that identifies the project associated with this Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration.
	GroupId *string `json:"groupId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links   []Link                                     `json:"links,omitempty"`
	Request *NDSLDAPVerifyConnectivityJobRequestParams `json:"request,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this request to verify an Lightweight Directory Access Protocol (LDAP) configuration.
	RequestId *string `json:"requestId,omitempty"`
	// Human-readable string that indicates the status of the Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration.
	Status *string `json:"status,omitempty"`
	// List that contains the validation messages related to the verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details. The list contains a document for each test that MongoDB Cloud runs. MongoDB Cloud stops running tests after the first failure.
	Validations []NDSLDAPVerifyConnectivityJobRequestValidation `json:"validations,omitempty"`
}

// NewNDSLDAPVerifyConnectivityJobRequest instantiates a new NDSLDAPVerifyConnectivityJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSLDAPVerifyConnectivityJobRequest() *NDSLDAPVerifyConnectivityJobRequest {
	this := NDSLDAPVerifyConnectivityJobRequest{}
	return &this
}

// NewNDSLDAPVerifyConnectivityJobRequestWithDefaults instantiates a new NDSLDAPVerifyConnectivityJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSLDAPVerifyConnectivityJobRequestWithDefaults() *NDSLDAPVerifyConnectivityJobRequest {
	this := NDSLDAPVerifyConnectivityJobRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetLinks(v []Link) {
	o.Links = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequest() NDSLDAPVerifyConnectivityJobRequestParams {
	if o == nil || IsNil(o.Request) {
		var ret NDSLDAPVerifyConnectivityJobRequestParams
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestOk() (*NDSLDAPVerifyConnectivityJobRequestParams, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given NDSLDAPVerifyConnectivityJobRequestParams and assigns it to the Request field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetRequest(v NDSLDAPVerifyConnectivityJobRequestParams) {
	o.Request = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetStatus(v string) {
	o.Status = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetValidations() []NDSLDAPVerifyConnectivityJobRequestValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []NDSLDAPVerifyConnectivityJobRequestValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) GetValidationsOk() ([]NDSLDAPVerifyConnectivityJobRequestValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequest) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []NDSLDAPVerifyConnectivityJobRequestValidation and assigns it to the Validations field.
func (o *NDSLDAPVerifyConnectivityJobRequest) SetValidations(v []NDSLDAPVerifyConnectivityJobRequestValidation) {
	o.Validations = v
}

func (o NDSLDAPVerifyConnectivityJobRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o NDSLDAPVerifyConnectivityJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return toSerialize, nil
}

type NullableNDSLDAPVerifyConnectivityJobRequest struct {
	value *NDSLDAPVerifyConnectivityJobRequest
	isSet bool
}

func (v NullableNDSLDAPVerifyConnectivityJobRequest) Get() *NDSLDAPVerifyConnectivityJobRequest {
	return v.value
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequest) Set(val *NDSLDAPVerifyConnectivityJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSLDAPVerifyConnectivityJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSLDAPVerifyConnectivityJobRequest(val *NDSLDAPVerifyConnectivityJobRequest) *NullableNDSLDAPVerifyConnectivityJobRequest {
	return &NullableNDSLDAPVerifyConnectivityJobRequest{value: val, isSet: true}
}

func (v NullableNDSLDAPVerifyConnectivityJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
