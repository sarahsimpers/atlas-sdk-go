// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ComputeAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputeAutoScaling{}

// ComputeAutoScaling Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type ComputeAutoScaling struct {
	// Flag that indicates whether cluster tier auto-scaling is enabled. Set to `true` to enable cluster tier auto-scaling. If enabled, you must specify a value for **providerSettings.autoScaling.compute.maxInstanceSize** also. Set to `false` to disable cluster tier auto-scaling.
	Enabled *bool `json:"enabled,omitempty"`
	// Flag that indicates whether the cluster tier can scale down. This is required if **autoScaling.compute.enabled** is `true`. If you enable this option, specify a value for **providerSettings.autoScaling.compute.minInstanceSize**.
	ScaleDownEnabled *bool `json:"scaleDownEnabled,omitempty"`
}

// NewComputeAutoScaling instantiates a new ComputeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAutoScaling() *ComputeAutoScaling {
	this := ComputeAutoScaling{}
	return &this
}

// NewComputeAutoScalingWithDefaults instantiates a new ComputeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAutoScalingWithDefaults() *ComputeAutoScaling {
	this := ComputeAutoScaling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComputeAutoScaling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScaling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComputeAutoScaling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComputeAutoScaling) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetScaleDownEnabled returns the ScaleDownEnabled field value if set, zero value otherwise.
func (o *ComputeAutoScaling) GetScaleDownEnabled() bool {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		var ret bool
		return ret
	}
	return *o.ScaleDownEnabled
}

// GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		return nil, false
	}
	return o.ScaleDownEnabled, true
}

// HasScaleDownEnabled returns a boolean if a field has been set.
func (o *ComputeAutoScaling) HasScaleDownEnabled() bool {
	if o != nil && !IsNil(o.ScaleDownEnabled) {
		return true
	}

	return false
}

// SetScaleDownEnabled gets a reference to the given bool and assigns it to the ScaleDownEnabled field.
func (o *ComputeAutoScaling) SetScaleDownEnabled(v bool) {
	o.ScaleDownEnabled = &v
}

func (o ComputeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ComputeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ScaleDownEnabled) {
		toSerialize["scaleDownEnabled"] = o.ScaleDownEnabled
	}
	return toSerialize, nil
}

type NullableComputeAutoScaling struct {
	value *ComputeAutoScaling
	isSet bool
}

func (v NullableComputeAutoScaling) Get() *ComputeAutoScaling {
	return v.value
}

func (v *NullableComputeAutoScaling) Set(val *ComputeAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeAutoScaling(val *ComputeAutoScaling) *NullableComputeAutoScaling {
	return &NullableComputeAutoScaling{value: val, isSet: true}
}

func (v NullableComputeAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
