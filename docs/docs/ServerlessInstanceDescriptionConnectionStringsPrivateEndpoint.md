# ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint**](ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint.md) | List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**. | [optional] [readonly] 
**SrvConnectionString** | Pointer to **string** | Private endpoint-aware connection string that uses the &#x60;mongodb+srv://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. The &#x60;mongodb+srv&#x60; protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). | [optional] [readonly] 
**Type** | Pointer to **string** | MongoDB process type to which your application connects. | [optional] [readonly] 

## Methods

### NewServerlessInstanceDescriptionConnectionStringsPrivateEndpoint

`func NewServerlessInstanceDescriptionConnectionStringsPrivateEndpoint() *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint`

NewServerlessInstanceDescriptionConnectionStringsPrivateEndpoint instantiates a new ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceDescriptionConnectionStringsPrivateEndpointWithDefaults

`func NewServerlessInstanceDescriptionConnectionStringsPrivateEndpointWithDefaults() *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint`

NewServerlessInstanceDescriptionConnectionStringsPrivateEndpointWithDefaults instantiates a new ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetEndpoints() []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetEndpointsOk() (*[]ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetEndpoints(v []ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetSrvConnectionString

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionString() string`

GetSrvConnectionString returns the SrvConnectionString field if non-nil, zero value otherwise.

### GetSrvConnectionStringOk

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetSrvConnectionStringOk() (*string, bool)`

GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvConnectionString

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetSrvConnectionString(v string)`

SetSrvConnectionString sets SrvConnectionString field to given value.

### HasSrvConnectionString

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasSrvConnectionString() bool`

HasSrvConnectionString returns a boolean if a field has been set.

### GetType

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


