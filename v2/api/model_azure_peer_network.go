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

// AzurePeerNetwork Group of Network Peering connection settings.
type AzurePeerNetwork struct {
	// Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides.
	AzureDirectoryId string `json:"azureDirectoryId"`
	// Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides.
	AzureSubscriptionId string `json:"azureSubscriptionId"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Error message returned when a requested Azure network peering resource returns `\"status\" : \"FAILED\"`. The resource returns `null` if the request succeeded.
	ErrorState *string `json:"errorState,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides.
	ResourceGroupName string `json:"resourceGroupName"`
	// State of the network peering connection at the time you made the request.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet.
	VnetName string `json:"vnetName"`
}

// NewAzurePeerNetwork instantiates a new AzurePeerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePeerNetwork() *AzurePeerNetwork {
	this := AzurePeerNetwork{}
	return &this
}

// NewAzurePeerNetworkWithDefaults instantiates a new AzurePeerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePeerNetworkWithDefaults() *AzurePeerNetwork {
	this := AzurePeerNetwork{}
	return &this
}

// GetAzureDirectoryId returns the AzureDirectoryId field value
func (o *AzurePeerNetwork) GetAzureDirectoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureDirectoryId
}

// GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetAzureDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureDirectoryId, true
}

// SetAzureDirectoryId sets field value
func (o *AzurePeerNetwork) SetAzureDirectoryId(v string) {
	o.AzureDirectoryId = v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value
func (o *AzurePeerNetwork) GetAzureSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureSubscriptionId, true
}

// SetAzureSubscriptionId sets field value
func (o *AzurePeerNetwork) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = v
}

// GetContainerId returns the ContainerId field value
func (o *AzurePeerNetwork) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AzurePeerNetwork) SetContainerId(v string) {
	o.ContainerId = v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetErrorState() string {
	if o == nil || o.ErrorState == nil {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetErrorStateOk() (*string, bool) {
	if o == nil || o.ErrorState == nil {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasErrorState() bool {
	if o != nil && o.ErrorState != nil {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *AzurePeerNetwork) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzurePeerNetwork) SetId(v string) {
	o.Id = &v
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *AzurePeerNetwork) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *AzurePeerNetwork) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AzurePeerNetwork) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AzurePeerNetwork) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AzurePeerNetwork) SetStatus(v string) {
	o.Status = &v
}

// GetVnetName returns the VnetName field value
func (o *AzurePeerNetwork) GetVnetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VnetName
}

// GetVnetNameOk returns a tuple with the VnetName field value
// and a boolean to check if the value has been set.
func (o *AzurePeerNetwork) GetVnetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnetName, true
}

// SetVnetName sets field value
func (o *AzurePeerNetwork) SetVnetName(v string) {
	o.VnetName = v
}

func (o AzurePeerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["azureDirectoryId"] = o.AzureDirectoryId
	}
	if true {
		toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId
	}
	if true {
		toSerialize["containerId"] = o.ContainerId
	}
	if o.ErrorState != nil {
		toSerialize["errorState"] = o.ErrorState
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["resourceGroupName"] = o.ResourceGroupName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["vnetName"] = o.VnetName
	}
	return json.Marshal(toSerialize)
}

type NullableAzurePeerNetwork struct {
	value *AzurePeerNetwork
	isSet bool
}

func (v NullableAzurePeerNetwork) Get() *AzurePeerNetwork {
	return v.value
}

func (v *NullableAzurePeerNetwork) Set(val *AzurePeerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePeerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePeerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePeerNetwork(val *AzurePeerNetwork) *NullableAzurePeerNetwork {
	return &NullableAzurePeerNetwork{value: val, isSet: true}
}

func (v NullableAzurePeerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePeerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


