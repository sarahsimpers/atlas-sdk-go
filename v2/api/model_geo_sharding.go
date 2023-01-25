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

// GeoSharding struct for GeoSharding
type GeoSharding struct {
	// List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist.
	CustomZoneMapping *map[string]map[string]interface{} `json:"customZoneMapping,omitempty"`
	// List that contains a namespace for a Global Cluster. MongoDB Cloud manages this cluster.
	ManagedNamespaces []ManagedNamespaces `json:"managedNamespaces,omitempty"`
}

// NewGeoSharding instantiates a new GeoSharding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoSharding() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// NewGeoShardingWithDefaults instantiates a new GeoSharding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoShardingWithDefaults() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// GetCustomZoneMapping returns the CustomZoneMapping field value if set, zero value otherwise.
func (o *GeoSharding) GetCustomZoneMapping() map[string]map[string]interface{} {
	if o == nil || o.CustomZoneMapping == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.CustomZoneMapping
}

// GetCustomZoneMappingOk returns a tuple with the CustomZoneMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetCustomZoneMappingOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.CustomZoneMapping == nil {
		return nil, false
	}
	return o.CustomZoneMapping, true
}

// HasCustomZoneMapping returns a boolean if a field has been set.
func (o *GeoSharding) HasCustomZoneMapping() bool {
	if o != nil && o.CustomZoneMapping != nil {
		return true
	}

	return false
}

// SetCustomZoneMapping gets a reference to the given map[string]map[string]interface{} and assigns it to the CustomZoneMapping field.
func (o *GeoSharding) SetCustomZoneMapping(v map[string]map[string]interface{}) {
	o.CustomZoneMapping = &v
}

// GetManagedNamespaces returns the ManagedNamespaces field value if set, zero value otherwise.
func (o *GeoSharding) GetManagedNamespaces() []ManagedNamespaces {
	if o == nil || o.ManagedNamespaces == nil {
		var ret []ManagedNamespaces
		return ret
	}
	return o.ManagedNamespaces
}

// GetManagedNamespacesOk returns a tuple with the ManagedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetManagedNamespacesOk() ([]ManagedNamespaces, bool) {
	if o == nil || o.ManagedNamespaces == nil {
		return nil, false
	}
	return o.ManagedNamespaces, true
}

// HasManagedNamespaces returns a boolean if a field has been set.
func (o *GeoSharding) HasManagedNamespaces() bool {
	if o != nil && o.ManagedNamespaces != nil {
		return true
	}

	return false
}

// SetManagedNamespaces gets a reference to the given []ManagedNamespaces and assigns it to the ManagedNamespaces field.
func (o *GeoSharding) SetManagedNamespaces(v []ManagedNamespaces) {
	o.ManagedNamespaces = v
}

func (o GeoSharding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomZoneMapping != nil {
		toSerialize["customZoneMapping"] = o.CustomZoneMapping
	}
	if o.ManagedNamespaces != nil {
		toSerialize["managedNamespaces"] = o.ManagedNamespaces
	}
	return json.Marshal(toSerialize)
}

type NullableGeoSharding struct {
	value *GeoSharding
	isSet bool
}

func (v NullableGeoSharding) Get() *GeoSharding {
	return v.value
}

func (v *NullableGeoSharding) Set(val *GeoSharding) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoSharding) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoSharding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoSharding(val *GeoSharding) *NullableGeoSharding {
	return &NullableGeoSharding{value: val, isSet: true}
}

func (v NullableGeoSharding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoSharding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


