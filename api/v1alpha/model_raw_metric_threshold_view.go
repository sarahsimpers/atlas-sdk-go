/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// RawMetricThresholdView struct for RawMetricThresholdView
type RawMetricThresholdView struct {
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.
	MetricName *string `json:"metricName,omitempty"`
	// MongoDB Cloud computes the current metric value as an average.
	Mode *string `json:"mode,omitempty"`
	Operator *Operator `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `json:"threshold,omitempty"`
	Units *RawMetricUnits `json:"units,omitempty"`
}

// NewRawMetricThresholdView instantiates a new RawMetricThresholdView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawMetricThresholdView() *RawMetricThresholdView {
	this := RawMetricThresholdView{}
	var units RawMetricUnits = RAWMETRICUNITS_RAW
	this.Units = &units
	return &this
}

// NewRawMetricThresholdViewWithDefaults instantiates a new RawMetricThresholdView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawMetricThresholdViewWithDefaults() *RawMetricThresholdView {
	this := RawMetricThresholdView{}
	var units RawMetricUnits = RAWMETRICUNITS_RAW
	this.Units = &units
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *RawMetricThresholdView) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricThresholdView) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *RawMetricThresholdView) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *RawMetricThresholdView) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RawMetricThresholdView) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricThresholdView) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RawMetricThresholdView) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RawMetricThresholdView) SetMode(v string) {
	o.Mode = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *RawMetricThresholdView) GetOperator() Operator {
	if o == nil || o.Operator == nil {
		var ret Operator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricThresholdView) GetOperatorOk() (*Operator, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *RawMetricThresholdView) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given Operator and assigns it to the Operator field.
func (o *RawMetricThresholdView) SetOperator(v Operator) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RawMetricThresholdView) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricThresholdView) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RawMetricThresholdView) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *RawMetricThresholdView) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *RawMetricThresholdView) GetUnits() RawMetricUnits {
	if o == nil || o.Units == nil {
		var ret RawMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawMetricThresholdView) GetUnitsOk() (*RawMetricUnits, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *RawMetricThresholdView) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given RawMetricUnits and assigns it to the Units field.
func (o *RawMetricThresholdView) SetUnits(v RawMetricUnits) {
	o.Units = &v
}

func (o RawMetricThresholdView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetricName != nil {
		toSerialize["metricName"] = o.MetricName
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableRawMetricThresholdView struct {
	value *RawMetricThresholdView
	isSet bool
}

func (v NullableRawMetricThresholdView) Get() *RawMetricThresholdView {
	return v.value
}

func (v *NullableRawMetricThresholdView) Set(val *RawMetricThresholdView) {
	v.value = val
	v.isSet = true
}

func (v NullableRawMetricThresholdView) IsSet() bool {
	return v.isSet
}

func (v *NullableRawMetricThresholdView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawMetricThresholdView(val *RawMetricThresholdView) *NullableRawMetricThresholdView {
	return &NullableRawMetricThresholdView{value: val, isSet: true}
}

func (v NullableRawMetricThresholdView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawMetricThresholdView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


