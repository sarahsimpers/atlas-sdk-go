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

// GCPCloudProviderContainerAllOf struct for GCPCloudProviderContainerAllOf
type GCPCloudProviderContainerAllOf struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. GCP further limits the block to a lower bound of the `/18` range.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Google Cloud (GCP) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of `/24` equals 27 three-node replica sets.
	AtlasCidrBlock *string `json:"atlasCidrBlock,omitempty"`
	// Unique string that identifies the GCP project in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container.
	GcpProjectId *string `json:"gcpProjectId,omitempty"`
	// Human-readable label that identifies the network in which MongoDB Cloud clusters in this network peering container exist. MongoDB Cloud returns **null** if no clusters exist in this network peering container.
	NetworkName *string `json:"networkName,omitempty"`
	// List of GCP regions to which you want to deploy this MongoDB Cloud network peering container.  In this MongoDB Cloud project, you can deploy clusters only to the GCP regions in this list. To deploy MongoDB Cloud clusters to other GCP regions, create additional projects.
	Regions []string `json:"regions,omitempty"`
}

// NewGCPCloudProviderContainerAllOf instantiates a new GCPCloudProviderContainerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPCloudProviderContainerAllOf() *GCPCloudProviderContainerAllOf {
	this := GCPCloudProviderContainerAllOf{}
	return &this
}

// NewGCPCloudProviderContainerAllOfWithDefaults instantiates a new GCPCloudProviderContainerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPCloudProviderContainerAllOfWithDefaults() *GCPCloudProviderContainerAllOf {
	this := GCPCloudProviderContainerAllOf{}
	return &this
}

// GetAtlasCidrBlock returns the AtlasCidrBlock field value if set, zero value otherwise.
func (o *GCPCloudProviderContainerAllOf) GetAtlasCidrBlock() string {
	if o == nil || o.AtlasCidrBlock == nil {
		var ret string
		return ret
	}
	return *o.AtlasCidrBlock
}

// GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainerAllOf) GetAtlasCidrBlockOk() (*string, bool) {
	if o == nil || o.AtlasCidrBlock == nil {
		return nil, false
	}
	return o.AtlasCidrBlock, true
}

// HasAtlasCidrBlock returns a boolean if a field has been set.
func (o *GCPCloudProviderContainerAllOf) HasAtlasCidrBlock() bool {
	if o != nil && o.AtlasCidrBlock != nil {
		return true
	}

	return false
}

// SetAtlasCidrBlock gets a reference to the given string and assigns it to the AtlasCidrBlock field.
func (o *GCPCloudProviderContainerAllOf) SetAtlasCidrBlock(v string) {
	o.AtlasCidrBlock = &v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise.
func (o *GCPCloudProviderContainerAllOf) GetGcpProjectId() string {
	if o == nil || o.GcpProjectId == nil {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainerAllOf) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || o.GcpProjectId == nil {
		return nil, false
	}
	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *GCPCloudProviderContainerAllOf) HasGcpProjectId() bool {
	if o != nil && o.GcpProjectId != nil {
		return true
	}

	return false
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *GCPCloudProviderContainerAllOf) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *GCPCloudProviderContainerAllOf) GetNetworkName() string {
	if o == nil || o.NetworkName == nil {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainerAllOf) GetNetworkNameOk() (*string, bool) {
	if o == nil || o.NetworkName == nil {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *GCPCloudProviderContainerAllOf) HasNetworkName() bool {
	if o != nil && o.NetworkName != nil {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *GCPCloudProviderContainerAllOf) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *GCPCloudProviderContainerAllOf) GetRegions() []string {
	if o == nil || o.Regions == nil {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainerAllOf) GetRegionsOk() ([]string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *GCPCloudProviderContainerAllOf) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *GCPCloudProviderContainerAllOf) SetRegions(v []string) {
	o.Regions = v
}

func (o GCPCloudProviderContainerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AtlasCidrBlock != nil {
		toSerialize["atlasCidrBlock"] = o.AtlasCidrBlock
	}
	if o.GcpProjectId != nil {
		toSerialize["gcpProjectId"] = o.GcpProjectId
	}
	if o.NetworkName != nil {
		toSerialize["networkName"] = o.NetworkName
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableGCPCloudProviderContainerAllOf struct {
	value *GCPCloudProviderContainerAllOf
	isSet bool
}

func (v NullableGCPCloudProviderContainerAllOf) Get() *GCPCloudProviderContainerAllOf {
	return v.value
}

func (v *NullableGCPCloudProviderContainerAllOf) Set(val *GCPCloudProviderContainerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPCloudProviderContainerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPCloudProviderContainerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPCloudProviderContainerAllOf(val *GCPCloudProviderContainerAllOf) *NullableGCPCloudProviderContainerAllOf {
	return &NullableGCPCloudProviderContainerAllOf{value: val, isSet: true}
}

func (v NullableGCPCloudProviderContainerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPCloudProviderContainerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


