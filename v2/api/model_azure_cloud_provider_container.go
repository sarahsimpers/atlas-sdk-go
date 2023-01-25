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

// AzureCloudProviderContainer Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type AzureCloudProviderContainer struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the `/24` and  `/21` ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of `/24` equals 27 three-node replica sets.
	AtlasCidrBlock string `json:"atlasCidrBlock"`
	// Unique string that identifies the Azure subscription in which the MongoDB Cloud VNet resides.
	AzureSubscriptionId *string `json:"azureSubscriptionId,omitempty"`
	// Azure region to which MongoDB Cloud deployed this network peering container.
	Region string `json:"region"`
	// Unique string that identifies the Azure VNet in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container.
	VnetName *string `json:"vnetName,omitempty"`
}

// NewAzureCloudProviderContainer instantiates a new AzureCloudProviderContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCloudProviderContainer() *AzureCloudProviderContainer {
	this := AzureCloudProviderContainer{}
	return &this
}

// NewAzureCloudProviderContainerWithDefaults instantiates a new AzureCloudProviderContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCloudProviderContainerWithDefaults() *AzureCloudProviderContainer {
	this := AzureCloudProviderContainer{}
	return &this
}

// GetAtlasCidrBlock returns the AtlasCidrBlock field value
func (o *AzureCloudProviderContainer) GetAtlasCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AtlasCidrBlock
}

// GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field value
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtlasCidrBlock, true
}

// SetAtlasCidrBlock sets field value
func (o *AzureCloudProviderContainer) SetAtlasCidrBlock(v string) {
	o.AtlasCidrBlock = v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetAzureSubscriptionId() string {
	if o == nil || o.AzureSubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil || o.AzureSubscriptionId == nil {
		return nil, false
	}
	return o.AzureSubscriptionId, true
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasAzureSubscriptionId() bool {
	if o != nil && o.AzureSubscriptionId != nil {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given string and assigns it to the AzureSubscriptionId field.
func (o *AzureCloudProviderContainer) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = &v
}

// GetRegion returns the Region field value
func (o *AzureCloudProviderContainer) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AzureCloudProviderContainer) SetRegion(v string) {
	o.Region = v
}

// GetVnetName returns the VnetName field value if set, zero value otherwise.
func (o *AzureCloudProviderContainer) GetVnetName() string {
	if o == nil || o.VnetName == nil {
		var ret string
		return ret
	}
	return *o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderContainer) GetVnetNameOk() (*string, bool) {
	if o == nil || o.VnetName == nil {
		return nil, false
	}
	return o.VnetName, true
}

// HasVnetName returns a boolean if a field has been set.
func (o *AzureCloudProviderContainer) HasVnetName() bool {
	if o != nil && o.VnetName != nil {
		return true
	}

	return false
}

// SetVnetName gets a reference to the given string and assigns it to the VnetName field.
func (o *AzureCloudProviderContainer) SetVnetName(v string) {
	o.VnetName = &v
}

func (o AzureCloudProviderContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["atlasCidrBlock"] = o.AtlasCidrBlock
	}
	if o.AzureSubscriptionId != nil {
		toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.VnetName != nil {
		toSerialize["vnetName"] = o.VnetName
	}
	return json.Marshal(toSerialize)
}

type NullableAzureCloudProviderContainer struct {
	value *AzureCloudProviderContainer
	isSet bool
}

func (v NullableAzureCloudProviderContainer) Get() *AzureCloudProviderContainer {
	return v.value
}

func (v *NullableAzureCloudProviderContainer) Set(val *AzureCloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCloudProviderContainer(val *AzureCloudProviderContainer) *NullableAzureCloudProviderContainer {
	return &NullableAzureCloudProviderContainer{value: val, isSet: true}
}

func (v NullableAzureCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


