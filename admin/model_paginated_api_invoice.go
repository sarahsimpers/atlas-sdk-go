// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PaginatedApiInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiInvoice{}

// PaginatedApiInvoice struct for PaginatedApiInvoice
type PaginatedApiInvoice struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []Invoice `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiInvoice instantiates a new PaginatedApiInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiInvoice() *PaginatedApiInvoice {
	this := PaginatedApiInvoice{}
	return &this
}

// NewPaginatedApiInvoiceWithDefaults instantiates a new PaginatedApiInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiInvoiceWithDefaults() *PaginatedApiInvoice {
	this := PaginatedApiInvoice{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiInvoice) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiInvoice) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiInvoice) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiInvoice) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiInvoice) GetResults() []Invoice {
	if o == nil || IsNil(o.Results) {
		var ret []Invoice
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiInvoice) GetResultsOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiInvoice) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Invoice and assigns it to the Results field.
func (o *PaginatedApiInvoice) SetResults(v []Invoice) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiInvoice) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiInvoice) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiInvoice) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiInvoice) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiInvoice) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedApiInvoice struct {
	value *PaginatedApiInvoice
	isSet bool
}

func (v NullablePaginatedApiInvoice) Get() *PaginatedApiInvoice {
	return v.value
}

func (v *NullablePaginatedApiInvoice) Set(val *PaginatedApiInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiInvoice(val *PaginatedApiInvoice) *NullablePaginatedApiInvoice {
	return &NullablePaginatedApiInvoice{value: val, isSet: true}
}

func (v NullablePaginatedApiInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
