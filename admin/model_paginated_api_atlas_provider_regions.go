// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PaginatedApiAtlasProviderRegions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiAtlasProviderRegions{}

// PaginatedApiAtlasProviderRegions struct for PaginatedApiAtlasProviderRegions
type PaginatedApiAtlasProviderRegions struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ProviderRegions `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiAtlasProviderRegions instantiates a new PaginatedApiAtlasProviderRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiAtlasProviderRegions() *PaginatedApiAtlasProviderRegions {
	this := PaginatedApiAtlasProviderRegions{}
	return &this
}

// NewPaginatedApiAtlasProviderRegionsWithDefaults instantiates a new PaginatedApiAtlasProviderRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiAtlasProviderRegionsWithDefaults() *PaginatedApiAtlasProviderRegions {
	this := PaginatedApiAtlasProviderRegions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiAtlasProviderRegions) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasProviderRegions) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiAtlasProviderRegions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiAtlasProviderRegions) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiAtlasProviderRegions) GetResults() []ProviderRegions {
	if o == nil || IsNil(o.Results) {
		var ret []ProviderRegions
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasProviderRegions) GetResultsOk() ([]ProviderRegions, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiAtlasProviderRegions) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProviderRegions and assigns it to the Results field.
func (o *PaginatedApiAtlasProviderRegions) SetResults(v []ProviderRegions) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiAtlasProviderRegions) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasProviderRegions) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiAtlasProviderRegions) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiAtlasProviderRegions) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiAtlasProviderRegions) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiAtlasProviderRegions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedApiAtlasProviderRegions struct {
	value *PaginatedApiAtlasProviderRegions
	isSet bool
}

func (v NullablePaginatedApiAtlasProviderRegions) Get() *PaginatedApiAtlasProviderRegions {
	return v.value
}

func (v *NullablePaginatedApiAtlasProviderRegions) Set(val *PaginatedApiAtlasProviderRegions) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiAtlasProviderRegions) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiAtlasProviderRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiAtlasProviderRegions(val *PaginatedApiAtlasProviderRegions) *NullablePaginatedApiAtlasProviderRegions {
	return &NullablePaginatedApiAtlasProviderRegions{value: val, isSet: true}
}

func (v NullablePaginatedApiAtlasProviderRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiAtlasProviderRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
