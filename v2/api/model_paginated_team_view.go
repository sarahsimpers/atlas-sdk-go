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

// PaginatedTeamView struct for PaginatedTeamView
type PaginatedTeamView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiTeamResponseView `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedTeamView instantiates a new PaginatedTeamView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTeamView() *PaginatedTeamView {
	this := PaginatedTeamView{}
	return &this
}

// NewPaginatedTeamViewWithDefaults instantiates a new PaginatedTeamView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTeamViewWithDefaults() *PaginatedTeamView {
	this := PaginatedTeamView{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedTeamView) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamView) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedTeamView) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedTeamView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTeamView) GetResults() []ApiTeamResponseView {
	if o == nil || o.Results == nil {
		var ret []ApiTeamResponseView
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamView) GetResultsOk() ([]ApiTeamResponseView, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTeamView) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApiTeamResponseView and assigns it to the Results field.
func (o *PaginatedTeamView) SetResults(v []ApiTeamResponseView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedTeamView) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamView) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedTeamView) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedTeamView) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedTeamView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
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

type NullablePaginatedTeamView struct {
	value *PaginatedTeamView
	isSet bool
}

func (v NullablePaginatedTeamView) Get() *PaginatedTeamView {
	return v.value
}

func (v *NullablePaginatedTeamView) Set(val *PaginatedTeamView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTeamView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTeamView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTeamView(val *PaginatedTeamView) *NullablePaginatedTeamView {
	return &NullablePaginatedTeamView{value: val, isSet: true}
}

func (v NullablePaginatedTeamView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTeamView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


