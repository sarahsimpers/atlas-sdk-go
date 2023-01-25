/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrgPaginatedEventView struct for OrgPaginatedEventView
type OrgPaginatedEventView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []EventViewForOrg `json:"results,omitempty"`
	// Number of documents returned in this response if **includeCount** query param is true.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewOrgPaginatedEventView instantiates a new OrgPaginatedEventView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgPaginatedEventView() *OrgPaginatedEventView {
	this := OrgPaginatedEventView{}
	return &this
}

// NewOrgPaginatedEventViewWithDefaults instantiates a new OrgPaginatedEventView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgPaginatedEventViewWithDefaults() *OrgPaginatedEventView {
	this := OrgPaginatedEventView{}
	return &this
}

// GetLinks returns the Links field value
func (o *OrgPaginatedEventView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *OrgPaginatedEventView) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *OrgPaginatedEventView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OrgPaginatedEventView) GetResults() []EventViewForOrg {
	if o == nil || o.Results == nil {
		var ret []EventViewForOrg
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPaginatedEventView) GetResultsOk() ([]EventViewForOrg, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OrgPaginatedEventView) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EventViewForOrg and assigns it to the Results field.
func (o *OrgPaginatedEventView) SetResults(v []EventViewForOrg) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OrgPaginatedEventView) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPaginatedEventView) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OrgPaginatedEventView) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OrgPaginatedEventView) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o OrgPaginatedEventView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableOrgPaginatedEventView struct {
	value *OrgPaginatedEventView
	isSet bool
}

func (v NullableOrgPaginatedEventView) Get() *OrgPaginatedEventView {
	return v.value
}

func (v *NullableOrgPaginatedEventView) Set(val *OrgPaginatedEventView) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgPaginatedEventView) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgPaginatedEventView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgPaginatedEventView(val *OrgPaginatedEventView) *NullableOrgPaginatedEventView {
	return &NullableOrgPaginatedEventView{value: val, isSet: true}
}

func (v NullableOrgPaginatedEventView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgPaginatedEventView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


