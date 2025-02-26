// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PaginatedApiUserAccessList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiUserAccessList{}

// PaginatedApiUserAccessList struct for PaginatedApiUserAccessList
type PaginatedApiUserAccessList struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []UserAccessList `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiUserAccessList instantiates a new PaginatedApiUserAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiUserAccessList() *PaginatedApiUserAccessList {
	this := PaginatedApiUserAccessList{}
	return &this
}

// NewPaginatedApiUserAccessListWithDefaults instantiates a new PaginatedApiUserAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiUserAccessListWithDefaults() *PaginatedApiUserAccessList {
	this := PaginatedApiUserAccessList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiUserAccessList) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiUserAccessList) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiUserAccessList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiUserAccessList) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiUserAccessList) GetResults() []UserAccessList {
	if o == nil || IsNil(o.Results) {
		var ret []UserAccessList
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiUserAccessList) GetResultsOk() ([]UserAccessList, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiUserAccessList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []UserAccessList and assigns it to the Results field.
func (o *PaginatedApiUserAccessList) SetResults(v []UserAccessList) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiUserAccessList) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiUserAccessList) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiUserAccessList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiUserAccessList) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiUserAccessList) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiUserAccessList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedApiUserAccessList struct {
	value *PaginatedApiUserAccessList
	isSet bool
}

func (v NullablePaginatedApiUserAccessList) Get() *PaginatedApiUserAccessList {
	return v.value
}

func (v *NullablePaginatedApiUserAccessList) Set(val *PaginatedApiUserAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiUserAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiUserAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiUserAccessList(val *PaginatedApiUserAccessList) *NullablePaginatedApiUserAccessList {
	return &NullablePaginatedApiUserAccessList{value: val, isSet: true}
}

func (v NullablePaginatedApiUserAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiUserAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
