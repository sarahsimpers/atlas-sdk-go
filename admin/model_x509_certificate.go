// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the X509Certificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X509Certificate{}

// X509Certificate struct for X509Certificate
type X509Certificate struct {
	Content   *string    `json:"content,omitempty"`
	NotAfter  *time.Time `json:"notAfter,omitempty"`
	NotBefore *time.Time `json:"notBefore,omitempty"`
}

// NewX509Certificate instantiates a new X509Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate() *X509Certificate {
	this := X509Certificate{}
	return &this
}

// NewX509CertificateWithDefaults instantiates a new X509Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateWithDefaults() *X509Certificate {
	this := X509Certificate{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *X509Certificate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *X509Certificate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *X509Certificate) SetContent(v string) {
	o.Content = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509Certificate) GetNotAfter() time.Time {
	if o == nil || IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509Certificate) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509Certificate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509Certificate) GetNotBefore() time.Time {
	if o == nil || IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509Certificate) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509Certificate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

func (o X509Certificate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o X509Certificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	return toSerialize, nil
}

type NullableX509Certificate struct {
	value *X509Certificate
	isSet bool
}

func (v NullableX509Certificate) Get() *X509Certificate {
	return v.value
}

func (v *NullableX509Certificate) Set(val *X509Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableX509Certificate) IsSet() bool {
	return v.isSet
}

func (v *NullableX509Certificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509Certificate(val *X509Certificate) *NullableX509Certificate {
	return &NullableX509Certificate{value: val, isSet: true}
}

func (v NullableX509Certificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509Certificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
