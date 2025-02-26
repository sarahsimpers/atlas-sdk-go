// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the AzurePrivateLinkConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePrivateLinkConnection{}

// AzurePrivateLinkConnection Group of Private Endpoint Service settings.
type AzurePrivateLinkConnection struct {
	// Cloud service provider that serves the requested endpoint service.
	CloudProvider string `json:"cloudProvider"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	Id *string `json:"id,omitempty"`
	// List of private endpoints assigned to this Azure Private Link Service.
	PrivateEndpoints []string `json:"privateEndpoints,omitempty"`
	// Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages.
	PrivateLinkServiceName *string `json:"privateLinkServiceName,omitempty"`
	// Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet.
	PrivateLinkServiceResourceId *string `json:"privateLinkServiceResourceId,omitempty"`
	// Cloud provider region that manages this Private Endpoint Service.
	RegionName *string `json:"regionName,omitempty"`
	// State of the Private Endpoint Service connection when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewAzurePrivateLinkConnection instantiates a new AzurePrivateLinkConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePrivateLinkConnection(cloudProvider string) *AzurePrivateLinkConnection {
	this := AzurePrivateLinkConnection{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewAzurePrivateLinkConnectionWithDefaults instantiates a new AzurePrivateLinkConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePrivateLinkConnectionWithDefaults() *AzurePrivateLinkConnection {
	this := AzurePrivateLinkConnection{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *AzurePrivateLinkConnection) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *AzurePrivateLinkConnection) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AzurePrivateLinkConnection) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzurePrivateLinkConnection) SetId(v string) {
	o.Id = &v
}

// GetPrivateEndpoints returns the PrivateEndpoints field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetPrivateEndpoints() []string {
	if o == nil || IsNil(o.PrivateEndpoints) {
		var ret []string
		return ret
	}
	return o.PrivateEndpoints
}

// GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetPrivateEndpointsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivateEndpoints) {
		return nil, false
	}
	return o.PrivateEndpoints, true
}

// HasPrivateEndpoints returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasPrivateEndpoints() bool {
	if o != nil && !IsNil(o.PrivateEndpoints) {
		return true
	}

	return false
}

// SetPrivateEndpoints gets a reference to the given []string and assigns it to the PrivateEndpoints field.
func (o *AzurePrivateLinkConnection) SetPrivateEndpoints(v []string) {
	o.PrivateEndpoints = v
}

// GetPrivateLinkServiceName returns the PrivateLinkServiceName field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceName() string {
	if o == nil || IsNil(o.PrivateLinkServiceName) {
		var ret string
		return ret
	}
	return *o.PrivateLinkServiceName
}

// GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkServiceName) {
		return nil, false
	}
	return o.PrivateLinkServiceName, true
}

// HasPrivateLinkServiceName returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasPrivateLinkServiceName() bool {
	if o != nil && !IsNil(o.PrivateLinkServiceName) {
		return true
	}

	return false
}

// SetPrivateLinkServiceName gets a reference to the given string and assigns it to the PrivateLinkServiceName field.
func (o *AzurePrivateLinkConnection) SetPrivateLinkServiceName(v string) {
	o.PrivateLinkServiceName = &v
}

// GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceResourceId() string {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		var ret string
		return ret
	}
	return *o.PrivateLinkServiceResourceId
}

// GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		return nil, false
	}
	return o.PrivateLinkServiceResourceId, true
}

// HasPrivateLinkServiceResourceId returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasPrivateLinkServiceResourceId() bool {
	if o != nil && !IsNil(o.PrivateLinkServiceResourceId) {
		return true
	}

	return false
}

// SetPrivateLinkServiceResourceId gets a reference to the given string and assigns it to the PrivateLinkServiceResourceId field.
func (o *AzurePrivateLinkConnection) SetPrivateLinkServiceResourceId(v string) {
	o.PrivateLinkServiceResourceId = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AzurePrivateLinkConnection) SetRegionName(v string) {
	o.RegionName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AzurePrivateLinkConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateLinkConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AzurePrivateLinkConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AzurePrivateLinkConnection) SetStatus(v string) {
	o.Status = &v
}

func (o AzurePrivateLinkConnection) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzurePrivateLinkConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAzurePrivateLinkConnection struct {
	value *AzurePrivateLinkConnection
	isSet bool
}

func (v NullableAzurePrivateLinkConnection) Get() *AzurePrivateLinkConnection {
	return v.value
}

func (v *NullableAzurePrivateLinkConnection) Set(val *AzurePrivateLinkConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePrivateLinkConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePrivateLinkConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePrivateLinkConnection(val *AzurePrivateLinkConnection) *NullableAzurePrivateLinkConnection {
	return &NullableAzurePrivateLinkConnection{value: val, isSet: true}
}

func (v NullableAzurePrivateLinkConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePrivateLinkConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
