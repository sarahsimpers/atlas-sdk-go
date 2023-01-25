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

// ExampleResourceResponseView20221018 Dummy view.
type ExampleResourceResponseView20221018 struct {
	// Array that contains the dummy metadata.
	Data []string `json:"data"`
	// Dummy description added as response.
	Description string `json:"description"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewExampleResourceResponseView20221018 instantiates a new ExampleResourceResponseView20221018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleResourceResponseView20221018() *ExampleResourceResponseView20221018 {
	this := ExampleResourceResponseView20221018{}
	return &this
}

// NewExampleResourceResponseView20221018WithDefaults instantiates a new ExampleResourceResponseView20221018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleResourceResponseView20221018WithDefaults() *ExampleResourceResponseView20221018 {
	this := ExampleResourceResponseView20221018{}
	return &this
}

// GetData returns the Data field value
func (o *ExampleResourceResponseView20221018) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20221018) GetDataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ExampleResourceResponseView20221018) SetData(v []string) {
	o.Data = v
}

// GetDescription returns the Description field value
func (o *ExampleResourceResponseView20221018) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20221018) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExampleResourceResponseView20221018) SetDescription(v string) {
	o.Description = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExampleResourceResponseView20221018) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20221018) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExampleResourceResponseView20221018) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ExampleResourceResponseView20221018) SetLinks(v []Link) {
	o.Links = v
}

func (o ExampleResourceResponseView20221018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableExampleResourceResponseView20221018 struct {
	value *ExampleResourceResponseView20221018
	isSet bool
}

func (v NullableExampleResourceResponseView20221018) Get() *ExampleResourceResponseView20221018 {
	return v.value
}

func (v *NullableExampleResourceResponseView20221018) Set(val *ExampleResourceResponseView20221018) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleResourceResponseView20221018) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleResourceResponseView20221018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleResourceResponseView20221018(val *ExampleResourceResponseView20221018) *NullableExampleResourceResponseView20221018 {
	return &NullableExampleResourceResponseView20221018{value: val, isSet: true}
}

func (v NullableExampleResourceResponseView20221018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleResourceResponseView20221018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


