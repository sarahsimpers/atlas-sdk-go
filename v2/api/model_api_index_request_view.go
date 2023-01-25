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

// ApiIndexRequestView struct for ApiIndexRequestView
type ApiIndexRequestView struct {
	Collation *Collation `json:"collation,omitempty"`
	// Human-readable label of the collection for which MongoDB Cloud creates an index.
	Collection string `json:"collection"`
	// Human-readable label of the database that holds the collection on which MongoDB Cloud creates an index.
	Db string `json:"db"`
	// List that contains one or more objects that describe the parameters that you want to index.
	Keys []map[string]string `json:"keys"`
	Options *IndexOptions `json:"options,omitempty"`
}

// NewApiIndexRequestView instantiates a new ApiIndexRequestView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIndexRequestView() *ApiIndexRequestView {
	this := ApiIndexRequestView{}
	return &this
}

// NewApiIndexRequestViewWithDefaults instantiates a new ApiIndexRequestView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIndexRequestViewWithDefaults() *ApiIndexRequestView {
	this := ApiIndexRequestView{}
	return &this
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *ApiIndexRequestView) GetCollation() Collation {
	if o == nil || o.Collation == nil {
		var ret Collation
		return ret
	}
	return *o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIndexRequestView) GetCollationOk() (*Collation, bool) {
	if o == nil || o.Collation == nil {
		return nil, false
	}
	return o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *ApiIndexRequestView) HasCollation() bool {
	if o != nil && o.Collation != nil {
		return true
	}

	return false
}

// SetCollation gets a reference to the given Collation and assigns it to the Collation field.
func (o *ApiIndexRequestView) SetCollation(v Collation) {
	o.Collation = &v
}

// GetCollection returns the Collection field value
func (o *ApiIndexRequestView) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *ApiIndexRequestView) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *ApiIndexRequestView) SetCollection(v string) {
	o.Collection = v
}

// GetDb returns the Db field value
func (o *ApiIndexRequestView) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *ApiIndexRequestView) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *ApiIndexRequestView) SetDb(v string) {
	o.Db = v
}

// GetKeys returns the Keys field value
func (o *ApiIndexRequestView) GetKeys() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ApiIndexRequestView) GetKeysOk() ([]map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *ApiIndexRequestView) SetKeys(v []map[string]string) {
	o.Keys = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ApiIndexRequestView) GetOptions() IndexOptions {
	if o == nil || o.Options == nil {
		var ret IndexOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIndexRequestView) GetOptionsOk() (*IndexOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApiIndexRequestView) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IndexOptions and assigns it to the Options field.
func (o *ApiIndexRequestView) SetOptions(v IndexOptions) {
	o.Options = &v
}

func (o ApiIndexRequestView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collation != nil {
		toSerialize["collation"] = o.Collation
	}
	if true {
		toSerialize["collection"] = o.Collection
	}
	if true {
		toSerialize["db"] = o.Db
	}
	if true {
		toSerialize["keys"] = o.Keys
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableApiIndexRequestView struct {
	value *ApiIndexRequestView
	isSet bool
}

func (v NullableApiIndexRequestView) Get() *ApiIndexRequestView {
	return v.value
}

func (v *NullableApiIndexRequestView) Set(val *ApiIndexRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIndexRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIndexRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIndexRequestView(val *ApiIndexRequestView) *NullableApiIndexRequestView {
	return &NullableApiIndexRequestView{value: val, isSet: true}
}

func (v NullableApiIndexRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIndexRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


