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

// ApiDatadogViewAllOf struct for ApiDatadogViewAllOf
type ApiDatadogViewAllOf struct {
	// Key that allows MongoDB Cloud to access your Datadog account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey *string `json:"apiKey,omitempty"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.  To learn more about Datadog's regions, see <a href=\"https://docs.datadoghq.com/getting_started/site/\" target=\"_blank\" rel=\"noopener noreferrer\">Datadog Sites.</a>
	Region *string `json:"region,omitempty"`
}

// NewApiDatadogViewAllOf instantiates a new ApiDatadogViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatadogViewAllOf() *ApiDatadogViewAllOf {
	this := ApiDatadogViewAllOf{}
	return &this
}

// NewApiDatadogViewAllOfWithDefaults instantiates a new ApiDatadogViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatadogViewAllOfWithDefaults() *ApiDatadogViewAllOf {
	this := ApiDatadogViewAllOf{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *ApiDatadogViewAllOf) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatadogViewAllOf) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *ApiDatadogViewAllOf) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *ApiDatadogViewAllOf) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ApiDatadogViewAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatadogViewAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ApiDatadogViewAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ApiDatadogViewAllOf) SetRegion(v string) {
	o.Region = &v
}

func (o ApiDatadogViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableApiDatadogViewAllOf struct {
	value *ApiDatadogViewAllOf
	isSet bool
}

func (v NullableApiDatadogViewAllOf) Get() *ApiDatadogViewAllOf {
	return v.value
}

func (v *NullableApiDatadogViewAllOf) Set(val *ApiDatadogViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDatadogViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDatadogViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDatadogViewAllOf(val *ApiDatadogViewAllOf) *NullableApiDatadogViewAllOf {
	return &NullableApiDatadogViewAllOf{value: val, isSet: true}
}

func (v NullableApiDatadogViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDatadogViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


