# InheritedRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | **string** | Human-readable label that identifies the database on which someone grants the action to one MongoDB user. | 
**Role** | **string** | Human-readable label that identifies the role inherited. Set this value to &#x60;admin&#x60; for every role except &#x60;read&#x60; or &#x60;readWrite&#x60;. | 

## Methods

### NewInheritedRole

`func NewInheritedRole(db string, role string, ) *InheritedRole`

NewInheritedRole instantiates a new InheritedRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedRoleWithDefaults

`func NewInheritedRoleWithDefaults() *InheritedRole`

NewInheritedRoleWithDefaults instantiates a new InheritedRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *InheritedRole) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *InheritedRole) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *InheritedRole) SetDb(v string)`

SetDb sets Db field to given value.


### GetRole

`func (o *InheritedRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InheritedRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InheritedRole) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


