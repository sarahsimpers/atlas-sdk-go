// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterstopword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterstopword{}

// TokenFilterstopword Filter that removes tokens that correspond to the specified stop words. This token filter doesn't analyze the stop words that you specify.
type TokenFilterstopword struct {
	// Flag that indicates whether to ignore the case of stop words when filtering the tokens to remove.
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// The stop words that correspond to the tokens to remove. Value must be one or more stop words.
	Tokens []string `json:"tokens"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterstopword instantiates a new TokenFilterstopword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterstopword(tokens []string, type_ string) *TokenFilterstopword {
	this := TokenFilterstopword{}
	var ignoreCase bool = true
	this.IgnoreCase = &ignoreCase
	this.Tokens = tokens
	this.Type = type_
	return &this
}

// NewTokenFilterstopwordWithDefaults instantiates a new TokenFilterstopword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterstopwordWithDefaults() *TokenFilterstopword {
	this := TokenFilterstopword{}
	var ignoreCase bool = true
	this.IgnoreCase = &ignoreCase
	return &this
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *TokenFilterstopword) GetIgnoreCase() bool {
	if o == nil || IsNil(o.IgnoreCase) {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterstopword) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCase) {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *TokenFilterstopword) HasIgnoreCase() bool {
	if o != nil && !IsNil(o.IgnoreCase) {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *TokenFilterstopword) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

// GetTokens returns the Tokens field value
func (o *TokenFilterstopword) GetTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *TokenFilterstopword) GetTokensOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *TokenFilterstopword) SetTokens(v []string) {
	o.Tokens = v
}

// GetType returns the Type field value
func (o *TokenFilterstopword) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterstopword) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterstopword) SetType(v string) {
	o.Type = v
}

func (o TokenFilterstopword) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterstopword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoreCase) {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	toSerialize["tokens"] = o.Tokens
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterstopword struct {
	value *TokenFilterstopword
	isSet bool
}

func (v NullableTokenFilterstopword) Get() *TokenFilterstopword {
	return v.value
}

func (v *NullableTokenFilterstopword) Set(val *TokenFilterstopword) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterstopword) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterstopword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterstopword(val *TokenFilterstopword) *NullableTokenFilterstopword {
	return &NullableTokenFilterstopword{value: val, isSet: true}
}

func (v NullableTokenFilterstopword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterstopword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
