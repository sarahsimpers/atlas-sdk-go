/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the CloudProviderAccessAzureServicePrincipal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessAzureServicePrincipal{}

// CloudProviderAccessAzureServicePrincipal Details that describe the features linked to the Azure Service Principal.
type CloudProviderAccessAzureServicePrincipal struct {
	// Unique 24-hexadecimal digit string that identifies the Azure Service Principal in Atlas.
	Id *string `json:"_id,omitempty"`
	// Azure Active Directory Application ID of Atlas.
	AtlasAzureAppId *string `json:"atlasAzureAppId,omitempty"`
	// Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// List that contains application features associated with this Azure Service Principal.
	FeatureUsages []CloudProviderAccessFeatureUsage `json:"featureUsages,omitempty"`
	// Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// UUID string that identifies the Azure Service Principal.
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`
	// UUID String that identifies the Azure Active Directory Tenant ID.
	TenantId *string `json:"tenantId,omitempty"`
	// Human-readable label that identifies the cloud provider of the role.
	ProviderName string `json:"providerName"`
}

// NewCloudProviderAccessAzureServicePrincipal instantiates a new CloudProviderAccessAzureServicePrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessAzureServicePrincipal(providerName string) *CloudProviderAccessAzureServicePrincipal {
	this := CloudProviderAccessAzureServicePrincipal{}
	this.ProviderName = providerName
	return &this
}

// NewCloudProviderAccessAzureServicePrincipalWithDefaults instantiates a new CloudProviderAccessAzureServicePrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessAzureServicePrincipalWithDefaults() *CloudProviderAccessAzureServicePrincipal {
	this := CloudProviderAccessAzureServicePrincipal{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudProviderAccessAzureServicePrincipal) SetId(v string) {
	o.Id = &v
}

// GetAtlasAzureAppId returns the AtlasAzureAppId field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAzureAppId() string {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		var ret string
		return ret
	}
	return *o.AtlasAzureAppId
}

// GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetAtlasAzureAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		return nil, false
	}
	return o.AtlasAzureAppId, true
}

// HasAtlasAzureAppId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasAtlasAzureAppId() bool {
	if o != nil && !IsNil(o.AtlasAzureAppId) {
		return true
	}

	return false
}

// SetAtlasAzureAppId gets a reference to the given string and assigns it to the AtlasAzureAppId field.
func (o *CloudProviderAccessAzureServicePrincipal) SetAtlasAzureAppId(v string) {
	o.AtlasAzureAppId = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CloudProviderAccessAzureServicePrincipal) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetFeatureUsages() []CloudProviderAccessFeatureUsage {
	if o == nil || IsNil(o.FeatureUsages) {
		var ret []CloudProviderAccessFeatureUsage
		return ret
	}
	return o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetFeatureUsagesOk() ([]CloudProviderAccessFeatureUsage, bool) {
	if o == nil || IsNil(o.FeatureUsages) {
		return nil, false
	}
	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasFeatureUsages() bool {
	if o != nil && !IsNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []CloudProviderAccessFeatureUsage and assigns it to the FeatureUsages field.
func (o *CloudProviderAccessAzureServicePrincipal) SetFeatureUsages(v []CloudProviderAccessFeatureUsage) {
	o.FeatureUsages = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *CloudProviderAccessAzureServicePrincipal) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetServicePrincipalId returns the ServicePrincipalId field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetServicePrincipalId() string {
	if o == nil || IsNil(o.ServicePrincipalId) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalId
}

// GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetServicePrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePrincipalId) {
		return nil, false
	}
	return o.ServicePrincipalId, true
}

// HasServicePrincipalId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasServicePrincipalId() bool {
	if o != nil && !IsNil(o.ServicePrincipalId) {
		return true
	}

	return false
}

// SetServicePrincipalId gets a reference to the given string and assigns it to the ServicePrincipalId field.
func (o *CloudProviderAccessAzureServicePrincipal) SetServicePrincipalId(v string) {
	o.ServicePrincipalId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CloudProviderAccessAzureServicePrincipal) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipal) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CloudProviderAccessAzureServicePrincipal) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProviderName returns the ProviderName field value
func (o *CloudProviderAccessAzureServicePrincipal) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipal) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudProviderAccessAzureServicePrincipal) SetProviderName(v string) {
	o.ProviderName = v
}

func (o CloudProviderAccessAzureServicePrincipal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderAccessAzureServicePrincipal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: _id is readOnly
	if !IsNil(o.AtlasAzureAppId) {
		toSerialize["atlasAzureAppId"] = o.AtlasAzureAppId
	}
	// skip: createdDate is readOnly
	// skip: featureUsages is readOnly
	// skip: lastUpdatedDate is readOnly
	if !IsNil(o.ServicePrincipalId) {
		toSerialize["servicePrincipalId"] = o.ServicePrincipalId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableCloudProviderAccessAzureServicePrincipal struct {
	value *CloudProviderAccessAzureServicePrincipal
	isSet bool
}

func (v NullableCloudProviderAccessAzureServicePrincipal) Get() *CloudProviderAccessAzureServicePrincipal {
	return v.value
}

func (v *NullableCloudProviderAccessAzureServicePrincipal) Set(val *CloudProviderAccessAzureServicePrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessAzureServicePrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessAzureServicePrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessAzureServicePrincipal(val *CloudProviderAccessAzureServicePrincipal) *NullableCloudProviderAccessAzureServicePrincipal {
	return &NullableCloudProviderAccessAzureServicePrincipal{value: val, isSet: true}
}

func (v NullableCloudProviderAccessAzureServicePrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessAzureServicePrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


