// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DataLakeHTTPStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeHTTPStore{}

// DataLakeHTTPStore struct for DataLakeHTTPStore
type DataLakeHTTPStore struct {
	// Flag that validates the scheme in the specified URLs. If `true`, allows insecure `HTTP` scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If `false`, allows secure `HTTPS` scheme only.
	AllowInsecure *bool `json:"allowInsecure,omitempty"`
	// Default format that Data Lake assumes if it encounters a file without an extension while searching the `storeName`. If omitted, Data Lake attempts to detect the file type by processing a few bytes of the file. The specified format only applies to the URLs specified in the **databases.[n].collections.[n].dataSources** object.
	DefaultFormat *string `json:"defaultFormat,omitempty"`
	// Comma-separated list of publicly accessible HTTP URLs where data is stored. You can't specify URLs that require authentication.
	Urls []string `json:"urls,omitempty"`
	// Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an `M10` or higher cluster.
	Name     *string `json:"name,omitempty"`
	Provider string  `json:"provider"`
}

// NewDataLakeHTTPStore instantiates a new DataLakeHTTPStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeHTTPStore(provider string) *DataLakeHTTPStore {
	this := DataLakeHTTPStore{}
	var allowInsecure bool = false
	this.AllowInsecure = &allowInsecure
	this.Provider = provider
	return &this
}

// NewDataLakeHTTPStoreWithDefaults instantiates a new DataLakeHTTPStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeHTTPStoreWithDefaults() *DataLakeHTTPStore {
	this := DataLakeHTTPStore{}
	var allowInsecure bool = false
	this.AllowInsecure = &allowInsecure
	return &this
}

// GetAllowInsecure returns the AllowInsecure field value if set, zero value otherwise.
func (o *DataLakeHTTPStore) GetAllowInsecure() bool {
	if o == nil || IsNil(o.AllowInsecure) {
		var ret bool
		return ret
	}
	return *o.AllowInsecure
}

// GetAllowInsecureOk returns a tuple with the AllowInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStore) GetAllowInsecureOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowInsecure) {
		return nil, false
	}
	return o.AllowInsecure, true
}

// HasAllowInsecure returns a boolean if a field has been set.
func (o *DataLakeHTTPStore) HasAllowInsecure() bool {
	if o != nil && !IsNil(o.AllowInsecure) {
		return true
	}

	return false
}

// SetAllowInsecure gets a reference to the given bool and assigns it to the AllowInsecure field.
func (o *DataLakeHTTPStore) SetAllowInsecure(v bool) {
	o.AllowInsecure = &v
}

// GetDefaultFormat returns the DefaultFormat field value if set, zero value otherwise.
func (o *DataLakeHTTPStore) GetDefaultFormat() string {
	if o == nil || IsNil(o.DefaultFormat) {
		var ret string
		return ret
	}
	return *o.DefaultFormat
}

// GetDefaultFormatOk returns a tuple with the DefaultFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStore) GetDefaultFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultFormat) {
		return nil, false
	}
	return o.DefaultFormat, true
}

// HasDefaultFormat returns a boolean if a field has been set.
func (o *DataLakeHTTPStore) HasDefaultFormat() bool {
	if o != nil && !IsNil(o.DefaultFormat) {
		return true
	}

	return false
}

// SetDefaultFormat gets a reference to the given string and assigns it to the DefaultFormat field.
func (o *DataLakeHTTPStore) SetDefaultFormat(v string) {
	o.DefaultFormat = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *DataLakeHTTPStore) GetUrls() []string {
	if o == nil || IsNil(o.Urls) {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStore) GetUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *DataLakeHTTPStore) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *DataLakeHTTPStore) SetUrls(v []string) {
	o.Urls = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeHTTPStore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeHTTPStore) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeHTTPStore) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value
func (o *DataLakeHTTPStore) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DataLakeHTTPStore) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DataLakeHTTPStore) SetProvider(v string) {
	o.Provider = v
}

func (o DataLakeHTTPStore) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeHTTPStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowInsecure) {
		toSerialize["allowInsecure"] = o.AllowInsecure
	}
	if !IsNil(o.DefaultFormat) {
		toSerialize["defaultFormat"] = o.DefaultFormat
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

type NullableDataLakeHTTPStore struct {
	value *DataLakeHTTPStore
	isSet bool
}

func (v NullableDataLakeHTTPStore) Get() *DataLakeHTTPStore {
	return v.value
}

func (v *NullableDataLakeHTTPStore) Set(val *DataLakeHTTPStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeHTTPStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeHTTPStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeHTTPStore(val *DataLakeHTTPStore) *NullableDataLakeHTTPStore {
	return &NullableDataLakeHTTPStore{value: val, isSet: true}
}

func (v NullableDataLakeHTTPStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeHTTPStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
