// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ProviderInstanceSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderInstanceSize{}

// ProviderInstanceSize List of instances sizes that this cloud provider supports.
type ProviderInstanceSize struct {
	// List of regions that this cloud provider supports for this instance size.
	AvailableRegions []AvailableRegion `json:"availableRegions,omitempty"`
	// Human-readable label that identifies the instance size or cluster tier.
	Name *string `json:"name,omitempty"`
}

// NewProviderInstanceSize instantiates a new ProviderInstanceSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderInstanceSize() *ProviderInstanceSize {
	this := ProviderInstanceSize{}
	return &this
}

// NewProviderInstanceSizeWithDefaults instantiates a new ProviderInstanceSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderInstanceSizeWithDefaults() *ProviderInstanceSize {
	this := ProviderInstanceSize{}
	return &this
}

// GetAvailableRegions returns the AvailableRegions field value if set, zero value otherwise.
func (o *ProviderInstanceSize) GetAvailableRegions() []AvailableRegion {
	if o == nil || IsNil(o.AvailableRegions) {
		var ret []AvailableRegion
		return ret
	}
	return o.AvailableRegions
}

// GetAvailableRegionsOk returns a tuple with the AvailableRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInstanceSize) GetAvailableRegionsOk() ([]AvailableRegion, bool) {
	if o == nil || IsNil(o.AvailableRegions) {
		return nil, false
	}
	return o.AvailableRegions, true
}

// HasAvailableRegions returns a boolean if a field has been set.
func (o *ProviderInstanceSize) HasAvailableRegions() bool {
	if o != nil && !IsNil(o.AvailableRegions) {
		return true
	}

	return false
}

// SetAvailableRegions gets a reference to the given []AvailableRegion and assigns it to the AvailableRegions field.
func (o *ProviderInstanceSize) SetAvailableRegions(v []AvailableRegion) {
	o.AvailableRegions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderInstanceSize) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInstanceSize) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderInstanceSize) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProviderInstanceSize) SetName(v string) {
	o.Name = &v
}

func (o ProviderInstanceSize) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ProviderInstanceSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableProviderInstanceSize struct {
	value *ProviderInstanceSize
	isSet bool
}

func (v NullableProviderInstanceSize) Get() *ProviderInstanceSize {
	return v.value
}

func (v *NullableProviderInstanceSize) Set(val *ProviderInstanceSize) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderInstanceSize) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderInstanceSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderInstanceSize(val *ProviderInstanceSize) *NullableProviderInstanceSize {
	return &NullableProviderInstanceSize{value: val, isSet: true}
}

func (v NullableProviderInstanceSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderInstanceSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
