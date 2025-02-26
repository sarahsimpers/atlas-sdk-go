# CreateOrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**ApiUser**](ApiUser.md) |  | [optional] 
**OrgOwnerId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Atlas user that you want to assign the Organization Owner role. | [optional] [readonly] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewCreateOrganizationResponse

`func NewCreateOrganizationResponse() *CreateOrganizationResponse`

NewCreateOrganizationResponse instantiates a new CreateOrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationResponseWithDefaults

`func NewCreateOrganizationResponseWithDefaults() *CreateOrganizationResponse`

NewCreateOrganizationResponseWithDefaults instantiates a new CreateOrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateOrganizationResponse) GetApiKey() ApiUser`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateOrganizationResponse) GetApiKeyOk() (*ApiUser, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateOrganizationResponse) SetApiKey(v ApiUser)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateOrganizationResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetOrgOwnerId

`func (o *CreateOrganizationResponse) GetOrgOwnerId() string`

GetOrgOwnerId returns the OrgOwnerId field if non-nil, zero value otherwise.

### GetOrgOwnerIdOk

`func (o *CreateOrganizationResponse) GetOrgOwnerIdOk() (*string, bool)`

GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgOwnerId

`func (o *CreateOrganizationResponse) SetOrgOwnerId(v string)`

SetOrgOwnerId sets OrgOwnerId field to given value.

### HasOrgOwnerId

`func (o *CreateOrganizationResponse) HasOrgOwnerId() bool`

HasOrgOwnerId returns a boolean if a field has been set.

### GetOrganization

`func (o *CreateOrganizationResponse) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateOrganizationResponse) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateOrganizationResponse) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CreateOrganizationResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


