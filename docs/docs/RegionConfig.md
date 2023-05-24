# RegionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAutoScaling** | Pointer to [**AutoScalingV15**](AutoScalingV15.md) |  | [optional] 
**AnalyticsSpecs** | Pointer to [**DedicatedHardwareSpec**](DedicatedHardwareSpec.md) |  | [optional] 
**AutoScaling** | Pointer to [**AutoScalingV15**](AutoScalingV15.md) |  | [optional] 
**ReadOnlySpecs** | Pointer to [**DedicatedHardwareSpec**](DedicatedHardwareSpec.md) |  | [optional] 
**ElectableSpecs** | Pointer to [**HardwareSpec**](HardwareSpec.md) |  | [optional] 
**Priority** | Pointer to **int** | Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to &#x60;0&#x60;. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is &#x60;7&#x60;.  **Example:** If you have three regions, their priorities would be &#x60;7&#x60;, &#x60;6&#x60;, and &#x60;5&#x60; respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be &#x60;4&#x60; and &#x60;3&#x60; respectively. | [optional] 
**ProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to &#x60;AWS&#x60;, &#x60;GCP&#x60;, &#x60;AZURE&#x60; or &#x60;TENANT&#x60;. | [optional] 
**RegionName** | Pointer to **string** | Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC. | [optional] 
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the multi-tenant cluster. The resource returns this parameter when **providerSettings.providerName** is &#x60;TENANT&#x60; and **providerSetting.instanceSizeName** is &#x60;M2&#x60; or &#x60;M5&#x60;. | [optional] 

## Methods

### NewRegionConfig

`func NewRegionConfig() *RegionConfig`

NewRegionConfig instantiates a new RegionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionConfigWithDefaults

`func NewRegionConfigWithDefaults() *RegionConfig`

NewRegionConfigWithDefaults instantiates a new RegionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAutoScaling

`func (o *RegionConfig) GetAnalyticsAutoScaling() AutoScalingV15`

GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field if non-nil, zero value otherwise.

### GetAnalyticsAutoScalingOk

`func (o *RegionConfig) GetAnalyticsAutoScalingOk() (*AutoScalingV15, bool)`

GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAutoScaling

`func (o *RegionConfig) SetAnalyticsAutoScaling(v AutoScalingV15)`

SetAnalyticsAutoScaling sets AnalyticsAutoScaling field to given value.

### HasAnalyticsAutoScaling

`func (o *RegionConfig) HasAnalyticsAutoScaling() bool`

HasAnalyticsAutoScaling returns a boolean if a field has been set.

### GetAnalyticsSpecs

`func (o *RegionConfig) GetAnalyticsSpecs() DedicatedHardwareSpec`

GetAnalyticsSpecs returns the AnalyticsSpecs field if non-nil, zero value otherwise.

### GetAnalyticsSpecsOk

`func (o *RegionConfig) GetAnalyticsSpecsOk() (*DedicatedHardwareSpec, bool)`

GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSpecs

`func (o *RegionConfig) SetAnalyticsSpecs(v DedicatedHardwareSpec)`

SetAnalyticsSpecs sets AnalyticsSpecs field to given value.

### HasAnalyticsSpecs

`func (o *RegionConfig) HasAnalyticsSpecs() bool`

HasAnalyticsSpecs returns a boolean if a field has been set.

### GetAutoScaling

`func (o *RegionConfig) GetAutoScaling() AutoScalingV15`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *RegionConfig) GetAutoScalingOk() (*AutoScalingV15, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *RegionConfig) SetAutoScaling(v AutoScalingV15)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *RegionConfig) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetReadOnlySpecs

`func (o *RegionConfig) GetReadOnlySpecs() DedicatedHardwareSpec`

GetReadOnlySpecs returns the ReadOnlySpecs field if non-nil, zero value otherwise.

### GetReadOnlySpecsOk

`func (o *RegionConfig) GetReadOnlySpecsOk() (*DedicatedHardwareSpec, bool)`

GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySpecs

`func (o *RegionConfig) SetReadOnlySpecs(v DedicatedHardwareSpec)`

SetReadOnlySpecs sets ReadOnlySpecs field to given value.

### HasReadOnlySpecs

`func (o *RegionConfig) HasReadOnlySpecs() bool`

HasReadOnlySpecs returns a boolean if a field has been set.

### GetElectableSpecs

`func (o *RegionConfig) GetElectableSpecs() HardwareSpec`

GetElectableSpecs returns the ElectableSpecs field if non-nil, zero value otherwise.

### GetElectableSpecsOk

`func (o *RegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool)`

GetElectableSpecsOk returns a tuple with the ElectableSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableSpecs

`func (o *RegionConfig) SetElectableSpecs(v HardwareSpec)`

SetElectableSpecs sets ElectableSpecs field to given value.

### HasElectableSpecs

`func (o *RegionConfig) HasElectableSpecs() bool`

HasElectableSpecs returns a boolean if a field has been set.

### GetPriority

`func (o *RegionConfig) GetPriority() int`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RegionConfig) GetPriorityOk() (*int, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RegionConfig) SetPriority(v int)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RegionConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *RegionConfig) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *RegionConfig) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *RegionConfig) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *RegionConfig) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetRegionName

`func (o *RegionConfig) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *RegionConfig) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *RegionConfig) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *RegionConfig) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetBackingProviderName

`func (o *RegionConfig) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *RegionConfig) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *RegionConfig) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *RegionConfig) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

