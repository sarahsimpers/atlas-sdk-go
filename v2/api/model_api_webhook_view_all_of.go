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

// ApiWebhookViewAllOf struct for ApiWebhookViewAllOf
type ApiWebhookViewAllOf struct {
	// An optional field returned if your webhook is configured with a secret.  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted.
	Secret *string `json:"secret,omitempty"`
	// Endpoint web address to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a webhook notification, the URL appears partially redacted.
	Url *string `json:"url,omitempty"`
}

// NewApiWebhookViewAllOf instantiates a new ApiWebhookViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWebhookViewAllOf() *ApiWebhookViewAllOf {
	this := ApiWebhookViewAllOf{}
	return &this
}

// NewApiWebhookViewAllOfWithDefaults instantiates a new ApiWebhookViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWebhookViewAllOfWithDefaults() *ApiWebhookViewAllOf {
	this := ApiWebhookViewAllOf{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ApiWebhookViewAllOf) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWebhookViewAllOf) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ApiWebhookViewAllOf) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ApiWebhookViewAllOf) SetSecret(v string) {
	o.Secret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiWebhookViewAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWebhookViewAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiWebhookViewAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiWebhookViewAllOf) SetUrl(v string) {
	o.Url = &v
}

func (o ApiWebhookViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableApiWebhookViewAllOf struct {
	value *ApiWebhookViewAllOf
	isSet bool
}

func (v NullableApiWebhookViewAllOf) Get() *ApiWebhookViewAllOf {
	return v.value
}

func (v *NullableApiWebhookViewAllOf) Set(val *ApiWebhookViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWebhookViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWebhookViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWebhookViewAllOf(val *ApiWebhookViewAllOf) *NullableApiWebhookViewAllOf {
	return &NullableApiWebhookViewAllOf{value: val, isSet: true}
}

func (v NullableApiWebhookViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWebhookViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


