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

// ApiPrometheusView Details to integrate one Prometheus account with one MongoDB Cloud project.
type ApiPrometheusView struct {
	// Flag that indicates whether someone has activated the Prometheus integration.
	Enabled bool `json:"enabled"`
	// Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics.
	ListenAddress *string `json:"listenAddress,omitempty"`
	Password *string `json:"password,omitempty"`
	RateLimitInterval *int32 `json:"rateLimitInterval,omitempty"`
	// Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.
	Scheme string `json:"scheme"`
	// Desired method to discover the Prometheus service.
	ServiceDiscovery string `json:"serviceDiscovery"`
	// Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host.
	TlsPemPath *string `json:"tlsPemPath,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
	// Human-readable label that identifies your Prometheus incoming webhook.
	Username string `json:"username"`
}

// NewApiPrometheusView instantiates a new ApiPrometheusView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrometheusView() *ApiPrometheusView {
	this := ApiPrometheusView{}
	var listenAddress string = ":9216"
	this.ListenAddress = &listenAddress
	return &this
}

// NewApiPrometheusViewWithDefaults instantiates a new ApiPrometheusView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrometheusViewWithDefaults() *ApiPrometheusView {
	this := ApiPrometheusView{}
	var listenAddress string = ":9216"
	this.ListenAddress = &listenAddress
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ApiPrometheusView) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApiPrometheusView) SetEnabled(v bool) {
	o.Enabled = v
}

// GetListenAddress returns the ListenAddress field value if set, zero value otherwise.
func (o *ApiPrometheusView) GetListenAddress() string {
	if o == nil || o.ListenAddress == nil {
		var ret string
		return ret
	}
	return *o.ListenAddress
}

// GetListenAddressOk returns a tuple with the ListenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetListenAddressOk() (*string, bool) {
	if o == nil || o.ListenAddress == nil {
		return nil, false
	}
	return o.ListenAddress, true
}

// HasListenAddress returns a boolean if a field has been set.
func (o *ApiPrometheusView) HasListenAddress() bool {
	if o != nil && o.ListenAddress != nil {
		return true
	}

	return false
}

// SetListenAddress gets a reference to the given string and assigns it to the ListenAddress field.
func (o *ApiPrometheusView) SetListenAddress(v string) {
	o.ListenAddress = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiPrometheusView) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiPrometheusView) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiPrometheusView) SetPassword(v string) {
	o.Password = &v
}

// GetRateLimitInterval returns the RateLimitInterval field value if set, zero value otherwise.
func (o *ApiPrometheusView) GetRateLimitInterval() int32 {
	if o == nil || o.RateLimitInterval == nil {
		var ret int32
		return ret
	}
	return *o.RateLimitInterval
}

// GetRateLimitIntervalOk returns a tuple with the RateLimitInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetRateLimitIntervalOk() (*int32, bool) {
	if o == nil || o.RateLimitInterval == nil {
		return nil, false
	}
	return o.RateLimitInterval, true
}

// HasRateLimitInterval returns a boolean if a field has been set.
func (o *ApiPrometheusView) HasRateLimitInterval() bool {
	if o != nil && o.RateLimitInterval != nil {
		return true
	}

	return false
}

// SetRateLimitInterval gets a reference to the given int32 and assigns it to the RateLimitInterval field.
func (o *ApiPrometheusView) SetRateLimitInterval(v int32) {
	o.RateLimitInterval = &v
}

// GetScheme returns the Scheme field value
func (o *ApiPrometheusView) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *ApiPrometheusView) SetScheme(v string) {
	o.Scheme = v
}

// GetServiceDiscovery returns the ServiceDiscovery field value
func (o *ApiPrometheusView) GetServiceDiscovery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetServiceDiscoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDiscovery, true
}

// SetServiceDiscovery sets field value
func (o *ApiPrometheusView) SetServiceDiscovery(v string) {
	o.ServiceDiscovery = v
}

// GetTlsPemPath returns the TlsPemPath field value if set, zero value otherwise.
func (o *ApiPrometheusView) GetTlsPemPath() string {
	if o == nil || o.TlsPemPath == nil {
		var ret string
		return ret
	}
	return *o.TlsPemPath
}

// GetTlsPemPathOk returns a tuple with the TlsPemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetTlsPemPathOk() (*string, bool) {
	if o == nil || o.TlsPemPath == nil {
		return nil, false
	}
	return o.TlsPemPath, true
}

// HasTlsPemPath returns a boolean if a field has been set.
func (o *ApiPrometheusView) HasTlsPemPath() bool {
	if o != nil && o.TlsPemPath != nil {
		return true
	}

	return false
}

// SetTlsPemPath gets a reference to the given string and assigns it to the TlsPemPath field.
func (o *ApiPrometheusView) SetTlsPemPath(v string) {
	o.TlsPemPath = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiPrometheusView) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiPrometheusView) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiPrometheusView) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value
func (o *ApiPrometheusView) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiPrometheusView) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiPrometheusView) SetUsername(v string) {
	o.Username = v
}

func (o ApiPrometheusView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ListenAddress != nil {
		toSerialize["listenAddress"] = o.ListenAddress
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RateLimitInterval != nil {
		toSerialize["rateLimitInterval"] = o.RateLimitInterval
	}
	if true {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["serviceDiscovery"] = o.ServiceDiscovery
	}
	if o.TlsPemPath != nil {
		toSerialize["tlsPemPath"] = o.TlsPemPath
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiPrometheusView struct {
	value *ApiPrometheusView
	isSet bool
}

func (v NullableApiPrometheusView) Get() *ApiPrometheusView {
	return v.value
}

func (v *NullableApiPrometheusView) Set(val *ApiPrometheusView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPrometheusView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPrometheusView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPrometheusView(val *ApiPrometheusView) *NullableApiPrometheusView {
	return &NullableApiPrometheusView{value: val, isSet: true}
}

func (v NullableApiPrometheusView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPrometheusView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


