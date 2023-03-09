/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the AtlasFTSAnalyzersViewManual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtlasFTSAnalyzersViewManual{}

// AtlasFTSAnalyzersViewManual Settings that describe one Atlas Search custom analyzer.
type AtlasFTSAnalyzersViewManual struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters []AtlasFTSAnalyzersViewManualCharFiltersInner `json:"charFilters,omitempty"`
	// Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - `lucene.` - `builtin.` - `mongodb.`
	Name string `json:"name"`
	// Filter that performs operations such as:  - Stemming, which reduces related words, such as \"talking\", \"talked\", and \"talks\" to their root word \"talk\".  - Redaction, the removal of sensitive information from public documents.
	TokenFilters []AtlasFTSAnalyzersViewManualTokenFiltersInner `json:"tokenFilters,omitempty"`
	Tokenizer AtlasFTSAnalyzersViewManualTokenizer `json:"tokenizer"`
}

// NewAtlasFTSAnalyzersViewManual instantiates a new AtlasFTSAnalyzersViewManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasFTSAnalyzersViewManual(name string, tokenizer AtlasFTSAnalyzersViewManualTokenizer) *AtlasFTSAnalyzersViewManual {
	this := AtlasFTSAnalyzersViewManual{}
	this.Name = name
	this.Tokenizer = tokenizer
	return &this
}

// NewAtlasFTSAnalyzersViewManualWithDefaults instantiates a new AtlasFTSAnalyzersViewManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasFTSAnalyzersViewManualWithDefaults() *AtlasFTSAnalyzersViewManual {
	this := AtlasFTSAnalyzersViewManual{}
	return &this
}

// GetCharFilters returns the CharFilters field value if set, zero value otherwise.
func (o *AtlasFTSAnalyzersViewManual) GetCharFilters() []AtlasFTSAnalyzersViewManualCharFiltersInner {
	if o == nil || IsNil(o.CharFilters) {
		var ret []AtlasFTSAnalyzersViewManualCharFiltersInner
		return ret
	}
	return o.CharFilters
}

// GetCharFiltersOk returns a tuple with the CharFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasFTSAnalyzersViewManual) GetCharFiltersOk() ([]AtlasFTSAnalyzersViewManualCharFiltersInner, bool) {
	if o == nil || IsNil(o.CharFilters) {
		return nil, false
	}
	return o.CharFilters, true
}

// HasCharFilters returns a boolean if a field has been set.
func (o *AtlasFTSAnalyzersViewManual) HasCharFilters() bool {
	if o != nil && !IsNil(o.CharFilters) {
		return true
	}

	return false
}

// SetCharFilters gets a reference to the given []AtlasFTSAnalyzersViewManualCharFiltersInner and assigns it to the CharFilters field.
func (o *AtlasFTSAnalyzersViewManual) SetCharFilters(v []AtlasFTSAnalyzersViewManualCharFiltersInner) {
	o.CharFilters = v
}

// GetName returns the Name field value
func (o *AtlasFTSAnalyzersViewManual) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AtlasFTSAnalyzersViewManual) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AtlasFTSAnalyzersViewManual) SetName(v string) {
	o.Name = v
}

// GetTokenFilters returns the TokenFilters field value if set, zero value otherwise.
func (o *AtlasFTSAnalyzersViewManual) GetTokenFilters() []AtlasFTSAnalyzersViewManualTokenFiltersInner {
	if o == nil || IsNil(o.TokenFilters) {
		var ret []AtlasFTSAnalyzersViewManualTokenFiltersInner
		return ret
	}
	return o.TokenFilters
}

// GetTokenFiltersOk returns a tuple with the TokenFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasFTSAnalyzersViewManual) GetTokenFiltersOk() ([]AtlasFTSAnalyzersViewManualTokenFiltersInner, bool) {
	if o == nil || IsNil(o.TokenFilters) {
		return nil, false
	}
	return o.TokenFilters, true
}

// HasTokenFilters returns a boolean if a field has been set.
func (o *AtlasFTSAnalyzersViewManual) HasTokenFilters() bool {
	if o != nil && !IsNil(o.TokenFilters) {
		return true
	}

	return false
}

// SetTokenFilters gets a reference to the given []AtlasFTSAnalyzersViewManualTokenFiltersInner and assigns it to the TokenFilters field.
func (o *AtlasFTSAnalyzersViewManual) SetTokenFilters(v []AtlasFTSAnalyzersViewManualTokenFiltersInner) {
	o.TokenFilters = v
}

// GetTokenizer returns the Tokenizer field value
func (o *AtlasFTSAnalyzersViewManual) GetTokenizer() AtlasFTSAnalyzersViewManualTokenizer {
	if o == nil {
		var ret AtlasFTSAnalyzersViewManualTokenizer
		return ret
	}

	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value
// and a boolean to check if the value has been set.
func (o *AtlasFTSAnalyzersViewManual) GetTokenizerOk() (*AtlasFTSAnalyzersViewManualTokenizer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokenizer, true
}

// SetTokenizer sets field value
func (o *AtlasFTSAnalyzersViewManual) SetTokenizer(v AtlasFTSAnalyzersViewManualTokenizer) {
	o.Tokenizer = v
}

func (o AtlasFTSAnalyzersViewManual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtlasFTSAnalyzersViewManual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CharFilters) {
		toSerialize["charFilters"] = o.CharFilters
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TokenFilters) {
		toSerialize["tokenFilters"] = o.TokenFilters
	}
	toSerialize["tokenizer"] = o.Tokenizer
	return toSerialize, nil
}

type NullableAtlasFTSAnalyzersViewManual struct {
	value *AtlasFTSAnalyzersViewManual
	isSet bool
}

func (v NullableAtlasFTSAnalyzersViewManual) Get() *AtlasFTSAnalyzersViewManual {
	return v.value
}

func (v *NullableAtlasFTSAnalyzersViewManual) Set(val *AtlasFTSAnalyzersViewManual) {
	v.value = val
	v.isSet = true
}

func (v NullableAtlasFTSAnalyzersViewManual) IsSet() bool {
	return v.isSet
}

func (v *NullableAtlasFTSAnalyzersViewManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtlasFTSAnalyzersViewManual(val *AtlasFTSAnalyzersViewManual) *NullableAtlasFTSAnalyzersViewManual {
	return &NullableAtlasFTSAnalyzersViewManual{value: val, isSet: true}
}

func (v NullableAtlasFTSAnalyzersViewManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtlasFTSAnalyzersViewManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

