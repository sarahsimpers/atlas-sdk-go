// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the TokenFiltersnowballStemming type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFiltersnowballStemming{}

// TokenFiltersnowballStemming Filter that stems tokens using a Snowball-generated stemmer.
type TokenFiltersnowballStemming struct {
	// Snowball-generated stemmer to use.
	StemmerName string `json:"stemmerName"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltersnowballStemming instantiates a new TokenFiltersnowballStemming object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltersnowballStemming(stemmerName string, type_ string) *TokenFiltersnowballStemming {
	this := TokenFiltersnowballStemming{}
	this.StemmerName = stemmerName
	this.Type = type_
	return &this
}

// NewTokenFiltersnowballStemmingWithDefaults instantiates a new TokenFiltersnowballStemming object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltersnowballStemmingWithDefaults() *TokenFiltersnowballStemming {
	this := TokenFiltersnowballStemming{}
	return &this
}

// GetStemmerName returns the StemmerName field value
func (o *TokenFiltersnowballStemming) GetStemmerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StemmerName
}

// GetStemmerNameOk returns a tuple with the StemmerName field value
// and a boolean to check if the value has been set.
func (o *TokenFiltersnowballStemming) GetStemmerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StemmerName, true
}

// SetStemmerName sets field value
func (o *TokenFiltersnowballStemming) SetStemmerName(v string) {
	o.StemmerName = v
}

// GetType returns the Type field value
func (o *TokenFiltersnowballStemming) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltersnowballStemming) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltersnowballStemming) SetType(v string) {
	o.Type = v
}

func (o TokenFiltersnowballStemming) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFiltersnowballStemming) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stemmerName"] = o.StemmerName
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFiltersnowballStemming struct {
	value *TokenFiltersnowballStemming
	isSet bool
}

func (v NullableTokenFiltersnowballStemming) Get() *TokenFiltersnowballStemming {
	return v.value
}

func (v *NullableTokenFiltersnowballStemming) Set(val *TokenFiltersnowballStemming) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltersnowballStemming) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltersnowballStemming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltersnowballStemming(val *TokenFiltersnowballStemming) *NullableTokenFiltersnowballStemming {
	return &NullableTokenFiltersnowballStemming{value: val, isSet: true}
}

func (v NullableTokenFiltersnowballStemming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltersnowballStemming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
