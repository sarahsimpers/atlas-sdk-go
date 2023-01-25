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

// CreateGCPEndpointGroupRequest Group of Private Endpoint settings.
type CreateGCPEndpointGroupRequest struct {
	// Human-readable label that identifies a set of endpoints.
	EndpointGroupName *string `json:"endpointGroupName,omitempty"`
	// List of individual private endpoints that comprise this endpoint group.
	Endpoints []CreateGCPForwardingRuleRequest `json:"endpoints,omitempty"`
	// Unique string that identifies the Google Cloud project in which you created the endpoints.
	GcpProjectId *string `json:"gcpProjectId,omitempty"`
}

// NewCreateGCPEndpointGroupRequest instantiates a new CreateGCPEndpointGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGCPEndpointGroupRequest() *CreateGCPEndpointGroupRequest {
	this := CreateGCPEndpointGroupRequest{}
	return &this
}

// NewCreateGCPEndpointGroupRequestWithDefaults instantiates a new CreateGCPEndpointGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGCPEndpointGroupRequestWithDefaults() *CreateGCPEndpointGroupRequest {
	this := CreateGCPEndpointGroupRequest{}
	return &this
}

// GetEndpointGroupName returns the EndpointGroupName field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequest) GetEndpointGroupName() string {
	if o == nil || o.EndpointGroupName == nil {
		var ret string
		return ret
	}
	return *o.EndpointGroupName
}

// GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetEndpointGroupNameOk() (*string, bool) {
	if o == nil || o.EndpointGroupName == nil {
		return nil, false
	}
	return o.EndpointGroupName, true
}

// HasEndpointGroupName returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequest) HasEndpointGroupName() bool {
	if o != nil && o.EndpointGroupName != nil {
		return true
	}

	return false
}

// SetEndpointGroupName gets a reference to the given string and assigns it to the EndpointGroupName field.
func (o *CreateGCPEndpointGroupRequest) SetEndpointGroupName(v string) {
	o.EndpointGroupName = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequest) GetEndpoints() []CreateGCPForwardingRuleRequest {
	if o == nil || o.Endpoints == nil {
		var ret []CreateGCPForwardingRuleRequest
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetEndpointsOk() ([]CreateGCPForwardingRuleRequest, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequest) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []CreateGCPForwardingRuleRequest and assigns it to the Endpoints field.
func (o *CreateGCPEndpointGroupRequest) SetEndpoints(v []CreateGCPForwardingRuleRequest) {
	o.Endpoints = v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise.
func (o *CreateGCPEndpointGroupRequest) GetGcpProjectId() string {
	if o == nil || o.GcpProjectId == nil {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGCPEndpointGroupRequest) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || o.GcpProjectId == nil {
		return nil, false
	}
	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *CreateGCPEndpointGroupRequest) HasGcpProjectId() bool {
	if o != nil && o.GcpProjectId != nil {
		return true
	}

	return false
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *CreateGCPEndpointGroupRequest) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

func (o CreateGCPEndpointGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndpointGroupName != nil {
		toSerialize["endpointGroupName"] = o.EndpointGroupName
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.GcpProjectId != nil {
		toSerialize["gcpProjectId"] = o.GcpProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGCPEndpointGroupRequest struct {
	value *CreateGCPEndpointGroupRequest
	isSet bool
}

func (v NullableCreateGCPEndpointGroupRequest) Get() *CreateGCPEndpointGroupRequest {
	return v.value
}

func (v *NullableCreateGCPEndpointGroupRequest) Set(val *CreateGCPEndpointGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGCPEndpointGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGCPEndpointGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGCPEndpointGroupRequest(val *CreateGCPEndpointGroupRequest) *NullableCreateGCPEndpointGroupRequest {
	return &NullableCreateGCPEndpointGroupRequest{value: val, isSet: true}
}

func (v NullableCreateGCPEndpointGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGCPEndpointGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


