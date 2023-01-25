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

// DataLakeHTTPStoreAllOf struct for DataLakeHTTPStoreAllOf
type DataLakeHTTPStoreAllOf struct {
	// Flag that validates the scheme in the specified URLs. If `true`, allows insecure `HTTP` scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If `false`, allows secure `HTTPS` scheme only.
	AllowInsecure *bool `json:"allowInsecure,omitempty"`
	// Default format that Data Lake assumes if it encounters a file without an extension while searching the `storeName`. If omitted, Data Lake attempts to detect the file type by processing a few bytes of the file. The specified format only applies to the URLs specified in the **databases.[n].collections.[n].dataSources** object.
	DefaultFormat *string `json:"defaultFormat,omitempty"`
	// Comma-separated list of publicly accessible HTTP URLs where data is stored. You can't specify URLs that require authentication.
	Urls []string `json:"urls,omitempty"`
}

// NewDataLakeHTTPStoreAllOf instantiates a new DataLakeHTTPStoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeHTTPStoreAllOf() *DataLakeHTTPStoreAllOf {
	this := DataLakeHTTPStoreAllOf{}
	var allowInsecure bool = false
	this.AllowInsecure = &allowInsecure
	return &this
}

// NewDataLakeHTTPStoreAllOfWithDefaults instantiates a new DataLakeHTTPStoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeHTTPStoreAllOfWithDefaults() *DataLakeHTTPStoreAllOf {
	this := DataLakeHTTPStoreAllOf{}
	var allowInsecure bool = false
	this.AllowInsecure = &allowInsecure
	return &this
}

// GetAllowInsecure returns the AllowInsecure field value if set, zero value otherwise.
func (o *DataLakeHTTPStoreAllOf) GetAllowInsecure() bool {
	if o == nil || o.AllowInsecure == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsecure
}

// GetAllowInsecureOk returns a tuple with the AllowInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStoreAllOf) GetAllowInsecureOk() (*bool, bool) {
	if o == nil || o.AllowInsecure == nil {
		return nil, false
	}
	return o.AllowInsecure, true
}

// HasAllowInsecure returns a boolean if a field has been set.
func (o *DataLakeHTTPStoreAllOf) HasAllowInsecure() bool {
	if o != nil && o.AllowInsecure != nil {
		return true
	}

	return false
}

// SetAllowInsecure gets a reference to the given bool and assigns it to the AllowInsecure field.
func (o *DataLakeHTTPStoreAllOf) SetAllowInsecure(v bool) {
	o.AllowInsecure = &v
}

// GetDefaultFormat returns the DefaultFormat field value if set, zero value otherwise.
func (o *DataLakeHTTPStoreAllOf) GetDefaultFormat() string {
	if o == nil || o.DefaultFormat == nil {
		var ret string
		return ret
	}
	return *o.DefaultFormat
}

// GetDefaultFormatOk returns a tuple with the DefaultFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStoreAllOf) GetDefaultFormatOk() (*string, bool) {
	if o == nil || o.DefaultFormat == nil {
		return nil, false
	}
	return o.DefaultFormat, true
}

// HasDefaultFormat returns a boolean if a field has been set.
func (o *DataLakeHTTPStoreAllOf) HasDefaultFormat() bool {
	if o != nil && o.DefaultFormat != nil {
		return true
	}

	return false
}

// SetDefaultFormat gets a reference to the given string and assigns it to the DefaultFormat field.
func (o *DataLakeHTTPStoreAllOf) SetDefaultFormat(v string) {
	o.DefaultFormat = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *DataLakeHTTPStoreAllOf) GetUrls() []string {
	if o == nil || o.Urls == nil {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStoreAllOf) GetUrlsOk() ([]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *DataLakeHTTPStoreAllOf) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *DataLakeHTTPStoreAllOf) SetUrls(v []string) {
	o.Urls = v
}

func (o DataLakeHTTPStoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowInsecure != nil {
		toSerialize["allowInsecure"] = o.AllowInsecure
	}
	if o.DefaultFormat != nil {
		toSerialize["defaultFormat"] = o.DefaultFormat
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableDataLakeHTTPStoreAllOf struct {
	value *DataLakeHTTPStoreAllOf
	isSet bool
}

func (v NullableDataLakeHTTPStoreAllOf) Get() *DataLakeHTTPStoreAllOf {
	return v.value
}

func (v *NullableDataLakeHTTPStoreAllOf) Set(val *DataLakeHTTPStoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeHTTPStoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeHTTPStoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeHTTPStoreAllOf(val *DataLakeHTTPStoreAllOf) *NullableDataLakeHTTPStoreAllOf {
	return &NullableDataLakeHTTPStoreAllOf{value: val, isSet: true}
}

func (v NullableDataLakeHTTPStoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeHTTPStoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


