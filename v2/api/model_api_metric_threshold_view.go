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

// ApiMetricThresholdView Threshold for the metric that, when exceeded, triggers an alert. The resource returns this parameter when `\"eventTypeName\" : \"OUTSIDE_METRIC_THRESHOLD\"`.
type ApiMetricThresholdView struct {
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. To learn more about the available metrics, see [Host Metrics](https://dochub.mongodb.org/core/atlas-host-metrics).  **NOTE**: If you set eventTypeName to OUTSIDE_SERVERLESS_METRIC_THRESHOLD, you can specify only metrics available for serverless. To learn more, see <a href=\"https://dochub.mongodb.org/core/alert-config-serverless-measurements\" target=\"_blank\">Serverless Measurements</a>.
	MetricName *string `json:"metricName,omitempty"`
	// MongoDB Cloud computes the current metric value as an average.
	Mode *string `json:"mode,omitempty"`
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `json:"threshold,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.  **EXAMPLE**: A metric that measures memory consumption would have a byte measurement, while a metric that measures time would have a time unit.
	Units *string `json:"units,omitempty"`
}

// NewApiMetricThresholdView instantiates a new ApiMetricThresholdView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMetricThresholdView() *ApiMetricThresholdView {
	this := ApiMetricThresholdView{}
	var mode string = "AVERAGE"
	this.Mode = &mode
	return &this
}

// NewApiMetricThresholdViewWithDefaults instantiates a new ApiMetricThresholdView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMetricThresholdViewWithDefaults() *ApiMetricThresholdView {
	this := ApiMetricThresholdView{}
	var mode string = "AVERAGE"
	this.Mode = &mode
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ApiMetricThresholdView) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricThresholdView) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ApiMetricThresholdView) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ApiMetricThresholdView) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ApiMetricThresholdView) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricThresholdView) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ApiMetricThresholdView) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ApiMetricThresholdView) SetMode(v string) {
	o.Mode = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ApiMetricThresholdView) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricThresholdView) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ApiMetricThresholdView) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ApiMetricThresholdView) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ApiMetricThresholdView) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricThresholdView) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ApiMetricThresholdView) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *ApiMetricThresholdView) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ApiMetricThresholdView) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricThresholdView) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ApiMetricThresholdView) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *ApiMetricThresholdView) SetUnits(v string) {
	o.Units = &v
}

func (o ApiMetricThresholdView) MarshalJSON() ([]byte, error) {
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

type NullableApiMetricThresholdView struct {
	value *ApiMetricThresholdView
	isSet bool
}

func (v NullableApiMetricThresholdView) Get() *ApiMetricThresholdView {
	return v.value
}

func (v *NullableApiMetricThresholdView) Set(val *ApiMetricThresholdView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMetricThresholdView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMetricThresholdView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMetricThresholdView(val *ApiMetricThresholdView) *NullableApiMetricThresholdView {
	return &NullableApiMetricThresholdView{value: val, isSet: true}
}

func (v NullableApiMetricThresholdView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMetricThresholdView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


