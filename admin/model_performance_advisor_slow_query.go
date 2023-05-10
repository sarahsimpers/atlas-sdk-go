// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PerformanceAdvisorSlowQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceAdvisorSlowQuery{}

// PerformanceAdvisorSlowQuery Details of one slow query that the Performance Advisor detected.
type PerformanceAdvisorSlowQuery struct {
	// Text of the MongoDB log related to this slow query.
	Line *string `json:"line,omitempty"`
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	Namespace *string `json:"namespace,omitempty"`
}

// NewPerformanceAdvisorSlowQuery instantiates a new PerformanceAdvisorSlowQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorSlowQuery() *PerformanceAdvisorSlowQuery {
	this := PerformanceAdvisorSlowQuery{}
	return &this
}

// NewPerformanceAdvisorSlowQueryWithDefaults instantiates a new PerformanceAdvisorSlowQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorSlowQueryWithDefaults() *PerformanceAdvisorSlowQuery {
	this := PerformanceAdvisorSlowQuery{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *PerformanceAdvisorSlowQuery) GetLine() string {
	if o == nil || IsNil(o.Line) {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *PerformanceAdvisorSlowQuery) SetLine(v string) {
	o.Line = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *PerformanceAdvisorSlowQuery) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQuery) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQuery) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *PerformanceAdvisorSlowQuery) SetNamespace(v string) {
	o.Namespace = &v
}

func (o PerformanceAdvisorSlowQuery) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PerformanceAdvisorSlowQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePerformanceAdvisorSlowQuery struct {
	value *PerformanceAdvisorSlowQuery
	isSet bool
}

func (v NullablePerformanceAdvisorSlowQuery) Get() *PerformanceAdvisorSlowQuery {
	return v.value
}

func (v *NullablePerformanceAdvisorSlowQuery) Set(val *PerformanceAdvisorSlowQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceAdvisorSlowQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceAdvisorSlowQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceAdvisorSlowQuery(val *PerformanceAdvisorSlowQuery) *NullablePerformanceAdvisorSlowQuery {
	return &NullablePerformanceAdvisorSlowQuery{value: val, isSet: true}
}

func (v NullablePerformanceAdvisorSlowQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceAdvisorSlowQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}