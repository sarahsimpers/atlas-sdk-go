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

// ClusterDescriptionConnectionStrings Collection of Uniform Resource Locators that point to the MongoDB database.
type ClusterDescriptionConnectionStrings struct {
	// Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related `mongodb://` connection string that you use to connect to MongoDB Cloud through the interface endpoint that the key names.
	AwsPrivateLink *map[string]string `json:"awsPrivateLink,omitempty"`
	// Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related `mongodb://` connection string that you use to connect to Atlas through the interface endpoint that the key names.
	AwsPrivateLinkSrv *map[string]string `json:"awsPrivateLinkSrv,omitempty"`
	// Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the `mongodb+srv://` protocol. The resource returns this parameter once someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this resource returns this parameter only if you enable custom DNS.
	Private *string `json:"private,omitempty"`
	// List of private endpoint-aware connection strings that you can use to connect to this cluster through a private endpoint. This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.
	PrivateEndpoint []ClusterDescriptionConnectionStringsPrivateEndpoint `json:"privateEndpoint,omitempty"`
	// Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the `mongodb+srv://` protocol. The resource returns this parameter when someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your driver supports it. If it doesn't, use `connectionStrings.private`. For Amazon Web Services (AWS) clusters, this parameter returns only if you [enable custom DNS](https://docs.atlas.mongodb.com/reference/api/aws-custom-dns-update/).
	PrivateSrv *string `json:"privateSrv,omitempty"`
	// Public connection string that you can use to connect to this cluster. This connection string uses the `mongodb://` protocol.
	Standard *string `json:"standard,omitempty"`
	// Public connection string that you can use to connect to this cluster. This connection string uses the `mongodb+srv://` protocol.
	StandardSrv *string `json:"standardSrv,omitempty"`
}

// NewClusterDescriptionConnectionStrings instantiates a new ClusterDescriptionConnectionStrings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDescriptionConnectionStrings() *ClusterDescriptionConnectionStrings {
	this := ClusterDescriptionConnectionStrings{}
	return &this
}

// NewClusterDescriptionConnectionStringsWithDefaults instantiates a new ClusterDescriptionConnectionStrings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDescriptionConnectionStringsWithDefaults() *ClusterDescriptionConnectionStrings {
	this := ClusterDescriptionConnectionStrings{}
	return &this
}

// GetAwsPrivateLink returns the AwsPrivateLink field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetAwsPrivateLink() map[string]string {
	if o == nil || o.AwsPrivateLink == nil {
		var ret map[string]string
		return ret
	}
	return *o.AwsPrivateLink
}

// GetAwsPrivateLinkOk returns a tuple with the AwsPrivateLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetAwsPrivateLinkOk() (*map[string]string, bool) {
	if o == nil || o.AwsPrivateLink == nil {
		return nil, false
	}
	return o.AwsPrivateLink, true
}

// HasAwsPrivateLink returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasAwsPrivateLink() bool {
	if o != nil && o.AwsPrivateLink != nil {
		return true
	}

	return false
}

// SetAwsPrivateLink gets a reference to the given map[string]string and assigns it to the AwsPrivateLink field.
func (o *ClusterDescriptionConnectionStrings) SetAwsPrivateLink(v map[string]string) {
	o.AwsPrivateLink = &v
}

// GetAwsPrivateLinkSrv returns the AwsPrivateLinkSrv field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetAwsPrivateLinkSrv() map[string]string {
	if o == nil || o.AwsPrivateLinkSrv == nil {
		var ret map[string]string
		return ret
	}
	return *o.AwsPrivateLinkSrv
}

// GetAwsPrivateLinkSrvOk returns a tuple with the AwsPrivateLinkSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetAwsPrivateLinkSrvOk() (*map[string]string, bool) {
	if o == nil || o.AwsPrivateLinkSrv == nil {
		return nil, false
	}
	return o.AwsPrivateLinkSrv, true
}

// HasAwsPrivateLinkSrv returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasAwsPrivateLinkSrv() bool {
	if o != nil && o.AwsPrivateLinkSrv != nil {
		return true
	}

	return false
}

// SetAwsPrivateLinkSrv gets a reference to the given map[string]string and assigns it to the AwsPrivateLinkSrv field.
func (o *ClusterDescriptionConnectionStrings) SetAwsPrivateLinkSrv(v map[string]string) {
	o.AwsPrivateLinkSrv = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetPrivate() string {
	if o == nil || o.Private == nil {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetPrivateOk() (*string, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *ClusterDescriptionConnectionStrings) SetPrivate(v string) {
	o.Private = &v
}

// GetPrivateEndpoint returns the PrivateEndpoint field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetPrivateEndpoint() []ClusterDescriptionConnectionStringsPrivateEndpoint {
	if o == nil || o.PrivateEndpoint == nil {
		var ret []ClusterDescriptionConnectionStringsPrivateEndpoint
		return ret
	}
	return o.PrivateEndpoint
}

// GetPrivateEndpointOk returns a tuple with the PrivateEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetPrivateEndpointOk() ([]ClusterDescriptionConnectionStringsPrivateEndpoint, bool) {
	if o == nil || o.PrivateEndpoint == nil {
		return nil, false
	}
	return o.PrivateEndpoint, true
}

// HasPrivateEndpoint returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasPrivateEndpoint() bool {
	if o != nil && o.PrivateEndpoint != nil {
		return true
	}

	return false
}

// SetPrivateEndpoint gets a reference to the given []ClusterDescriptionConnectionStringsPrivateEndpoint and assigns it to the PrivateEndpoint field.
func (o *ClusterDescriptionConnectionStrings) SetPrivateEndpoint(v []ClusterDescriptionConnectionStringsPrivateEndpoint) {
	o.PrivateEndpoint = v
}

// GetPrivateSrv returns the PrivateSrv field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetPrivateSrv() string {
	if o == nil || o.PrivateSrv == nil {
		var ret string
		return ret
	}
	return *o.PrivateSrv
}

// GetPrivateSrvOk returns a tuple with the PrivateSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetPrivateSrvOk() (*string, bool) {
	if o == nil || o.PrivateSrv == nil {
		return nil, false
	}
	return o.PrivateSrv, true
}

// HasPrivateSrv returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasPrivateSrv() bool {
	if o != nil && o.PrivateSrv != nil {
		return true
	}

	return false
}

// SetPrivateSrv gets a reference to the given string and assigns it to the PrivateSrv field.
func (o *ClusterDescriptionConnectionStrings) SetPrivateSrv(v string) {
	o.PrivateSrv = &v
}

// GetStandard returns the Standard field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetStandard() string {
	if o == nil || o.Standard == nil {
		var ret string
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetStandardOk() (*string, bool) {
	if o == nil || o.Standard == nil {
		return nil, false
	}
	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasStandard() bool {
	if o != nil && o.Standard != nil {
		return true
	}

	return false
}

// SetStandard gets a reference to the given string and assigns it to the Standard field.
func (o *ClusterDescriptionConnectionStrings) SetStandard(v string) {
	o.Standard = &v
}

// GetStandardSrv returns the StandardSrv field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStrings) GetStandardSrv() string {
	if o == nil || o.StandardSrv == nil {
		var ret string
		return ret
	}
	return *o.StandardSrv
}

// GetStandardSrvOk returns a tuple with the StandardSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStrings) GetStandardSrvOk() (*string, bool) {
	if o == nil || o.StandardSrv == nil {
		return nil, false
	}
	return o.StandardSrv, true
}

// HasStandardSrv returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStrings) HasStandardSrv() bool {
	if o != nil && o.StandardSrv != nil {
		return true
	}

	return false
}

// SetStandardSrv gets a reference to the given string and assigns it to the StandardSrv field.
func (o *ClusterDescriptionConnectionStrings) SetStandardSrv(v string) {
	o.StandardSrv = &v
}

func (o ClusterDescriptionConnectionStrings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsPrivateLink != nil {
		toSerialize["awsPrivateLink"] = o.AwsPrivateLink
	}
	if o.AwsPrivateLinkSrv != nil {
		toSerialize["awsPrivateLinkSrv"] = o.AwsPrivateLinkSrv
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.PrivateEndpoint != nil {
		toSerialize["privateEndpoint"] = o.PrivateEndpoint
	}
	if o.PrivateSrv != nil {
		toSerialize["privateSrv"] = o.PrivateSrv
	}
	if o.Standard != nil {
		toSerialize["standard"] = o.Standard
	}
	if o.StandardSrv != nil {
		toSerialize["standardSrv"] = o.StandardSrv
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDescriptionConnectionStrings struct {
	value *ClusterDescriptionConnectionStrings
	isSet bool
}

func (v NullableClusterDescriptionConnectionStrings) Get() *ClusterDescriptionConnectionStrings {
	return v.value
}

func (v *NullableClusterDescriptionConnectionStrings) Set(val *ClusterDescriptionConnectionStrings) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDescriptionConnectionStrings) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDescriptionConnectionStrings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDescriptionConnectionStrings(val *ClusterDescriptionConnectionStrings) *NullableClusterDescriptionConnectionStrings {
	return &NullableClusterDescriptionConnectionStrings{value: val, isSet: true}
}

func (v NullableClusterDescriptionConnectionStrings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDescriptionConnectionStrings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


