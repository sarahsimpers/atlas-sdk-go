// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DLSIngestionSink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DLSIngestionSink{}

// DLSIngestionSink Atlas Data Lake Storage as the destination for a Data Lake Pipeline.
type DLSIngestionSink struct {
	// Target cloud provider for this Data Lake Pipeline.
	MetadataProvider *string `json:"metadataProvider,omitempty"`
	// Target cloud provider region for this Data Lake Pipeline.
	MetadataRegion *string `json:"metadataRegion,omitempty"`
	// Ordered fields used to physically organize data in the destination.
	PartitionFields []PartitionField `json:"partitionFields,omitempty"`
	// Type of ingestion destination of this Data Lake Pipeline.
	Type *string `json:"type,omitempty"`
}

// NewDLSIngestionSink instantiates a new DLSIngestionSink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDLSIngestionSink() *DLSIngestionSink {
	this := DLSIngestionSink{}
	return &this
}

// NewDLSIngestionSinkWithDefaults instantiates a new DLSIngestionSink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDLSIngestionSinkWithDefaults() *DLSIngestionSink {
	this := DLSIngestionSink{}
	return &this
}

// GetMetadataProvider returns the MetadataProvider field value if set, zero value otherwise.
func (o *DLSIngestionSink) GetMetadataProvider() string {
	if o == nil || IsNil(o.MetadataProvider) {
		var ret string
		return ret
	}
	return *o.MetadataProvider
}

// GetMetadataProviderOk returns a tuple with the MetadataProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSink) GetMetadataProviderOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataProvider) {
		return nil, false
	}
	return o.MetadataProvider, true
}

// HasMetadataProvider returns a boolean if a field has been set.
func (o *DLSIngestionSink) HasMetadataProvider() bool {
	if o != nil && !IsNil(o.MetadataProvider) {
		return true
	}

	return false
}

// SetMetadataProvider gets a reference to the given string and assigns it to the MetadataProvider field.
func (o *DLSIngestionSink) SetMetadataProvider(v string) {
	o.MetadataProvider = &v
}

// GetMetadataRegion returns the MetadataRegion field value if set, zero value otherwise.
func (o *DLSIngestionSink) GetMetadataRegion() string {
	if o == nil || IsNil(o.MetadataRegion) {
		var ret string
		return ret
	}
	return *o.MetadataRegion
}

// GetMetadataRegionOk returns a tuple with the MetadataRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSink) GetMetadataRegionOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataRegion) {
		return nil, false
	}
	return o.MetadataRegion, true
}

// HasMetadataRegion returns a boolean if a field has been set.
func (o *DLSIngestionSink) HasMetadataRegion() bool {
	if o != nil && !IsNil(o.MetadataRegion) {
		return true
	}

	return false
}

// SetMetadataRegion gets a reference to the given string and assigns it to the MetadataRegion field.
func (o *DLSIngestionSink) SetMetadataRegion(v string) {
	o.MetadataRegion = &v
}

// GetPartitionFields returns the PartitionFields field value if set, zero value otherwise.
func (o *DLSIngestionSink) GetPartitionFields() []PartitionField {
	if o == nil || IsNil(o.PartitionFields) {
		var ret []PartitionField
		return ret
	}
	return o.PartitionFields
}

// GetPartitionFieldsOk returns a tuple with the PartitionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSink) GetPartitionFieldsOk() ([]PartitionField, bool) {
	if o == nil || IsNil(o.PartitionFields) {
		return nil, false
	}
	return o.PartitionFields, true
}

// HasPartitionFields returns a boolean if a field has been set.
func (o *DLSIngestionSink) HasPartitionFields() bool {
	if o != nil && !IsNil(o.PartitionFields) {
		return true
	}

	return false
}

// SetPartitionFields gets a reference to the given []PartitionField and assigns it to the PartitionFields field.
func (o *DLSIngestionSink) SetPartitionFields(v []PartitionField) {
	o.PartitionFields = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DLSIngestionSink) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSink) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DLSIngestionSink) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DLSIngestionSink) SetType(v string) {
	o.Type = &v
}

func (o DLSIngestionSink) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DLSIngestionSink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataProvider) {
		toSerialize["metadataProvider"] = o.MetadataProvider
	}
	if !IsNil(o.MetadataRegion) {
		toSerialize["metadataRegion"] = o.MetadataRegion
	}
	if !IsNil(o.PartitionFields) {
		toSerialize["partitionFields"] = o.PartitionFields
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDLSIngestionSink struct {
	value *DLSIngestionSink
	isSet bool
}

func (v NullableDLSIngestionSink) Get() *DLSIngestionSink {
	return v.value
}

func (v *NullableDLSIngestionSink) Set(val *DLSIngestionSink) {
	v.value = val
	v.isSet = true
}

func (v NullableDLSIngestionSink) IsSet() bool {
	return v.isSet
}

func (v *NullableDLSIngestionSink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDLSIngestionSink(val *DLSIngestionSink) *NullableDLSIngestionSink {
	return &NullableDLSIngestionSink{value: val, isSet: true}
}

func (v NullableDLSIngestionSink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDLSIngestionSink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
