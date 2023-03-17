/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ApiAtlasPartitionFieldView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAtlasPartitionFieldView{}

// ApiAtlasPartitionFieldView Partition Field in the Data Lake Storage provider for a Data Lake Pipeline.
type ApiAtlasPartitionFieldView struct {
	// Human-readable label that identifies the field name used to partition data.
	FieldName string `json:"fieldName"`
	// Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero.
	Order int32 `json:"order"`
}

// NewApiAtlasPartitionFieldView instantiates a new ApiAtlasPartitionFieldView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasPartitionFieldView(fieldName string, order int32) *ApiAtlasPartitionFieldView {
	this := ApiAtlasPartitionFieldView{}
	this.FieldName = fieldName
	this.Order = order
	return &this
}

// NewApiAtlasPartitionFieldViewWithDefaults instantiates a new ApiAtlasPartitionFieldView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasPartitionFieldViewWithDefaults() *ApiAtlasPartitionFieldView {
	this := ApiAtlasPartitionFieldView{}
	var order int32 = 0
	this.Order = order
	return &this
}

// GetFieldName returns the FieldName field value
func (o *ApiAtlasPartitionFieldView) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasPartitionFieldView) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *ApiAtlasPartitionFieldView) SetFieldName(v string) {
	o.FieldName = v
}

// GetOrder returns the Order field value
func (o *ApiAtlasPartitionFieldView) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasPartitionFieldView) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ApiAtlasPartitionFieldView) SetOrder(v int32) {
	o.Order = v
}

func (o ApiAtlasPartitionFieldView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAtlasPartitionFieldView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableApiAtlasPartitionFieldView struct {
	value *ApiAtlasPartitionFieldView
	isSet bool
}

func (v NullableApiAtlasPartitionFieldView) Get() *ApiAtlasPartitionFieldView {
	return v.value
}

func (v *NullableApiAtlasPartitionFieldView) Set(val *ApiAtlasPartitionFieldView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasPartitionFieldView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasPartitionFieldView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasPartitionFieldView(val *ApiAtlasPartitionFieldView) *NullableApiAtlasPartitionFieldView {
	return &NullableApiAtlasPartitionFieldView{value: val, isSet: true}
}

func (v NullableApiAtlasPartitionFieldView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasPartitionFieldView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

