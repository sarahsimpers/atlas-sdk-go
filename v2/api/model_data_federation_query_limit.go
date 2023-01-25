/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// DataFederationQueryLimit Details of a query limit for Data Federation. Query limit is the limit on the amount of usage during a time period based on cost.
type DataFederationQueryLimit struct {
	// Amount that indicates the current usage of the limit.
	CurrentUsage *int64 `json:"currentUsage,omitempty"`
	// Default value of the limit.
	DefaultLimit *int64 `json:"defaultLimit,omitempty"`
	// Only used for Data Federation limits. Timestamp that indicates when this usage limit was last modified. This field uses the ISO 8601 timestamp format in UTC.
	LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
	// Maximum value of the limit.
	MaximumLimit *int64 `json:"maximumLimit,omitempty"`
	// Human-readable label that identifies the user-managed limit to modify.
	Name string `json:"name"`
	// Only used for Data Federation limits. Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit.
	OverrunPolicy *string `json:"overrunPolicy,omitempty"`
	// Amount to set the limit to.
	Value int64 `json:"value"`
}

// NewDataFederationQueryLimit instantiates a new DataFederationQueryLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFederationQueryLimit() *DataFederationQueryLimit {
	this := DataFederationQueryLimit{}
	return &this
}

// NewDataFederationQueryLimitWithDefaults instantiates a new DataFederationQueryLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFederationQueryLimitWithDefaults() *DataFederationQueryLimit {
	this := DataFederationQueryLimit{}
	return &this
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetCurrentUsage() int64 {
	if o == nil || o.CurrentUsage == nil {
		var ret int64
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetCurrentUsageOk() (*int64, bool) {
	if o == nil || o.CurrentUsage == nil {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasCurrentUsage() bool {
	if o != nil && o.CurrentUsage != nil {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given int64 and assigns it to the CurrentUsage field.
func (o *DataFederationQueryLimit) SetCurrentUsage(v int64) {
	o.CurrentUsage = &v
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetDefaultLimit() int64 {
	if o == nil || o.DefaultLimit == nil {
		var ret int64
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetDefaultLimitOk() (*int64, bool) {
	if o == nil || o.DefaultLimit == nil {
		return nil, false
	}
	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasDefaultLimit() bool {
	if o != nil && o.DefaultLimit != nil {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int64 and assigns it to the DefaultLimit field.
func (o *DataFederationQueryLimit) SetDefaultLimit(v int64) {
	o.DefaultLimit = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetLastModifiedDate() time.Time {
	if o == nil || o.LastModifiedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDate == nil {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *DataFederationQueryLimit) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetMaximumLimit() int64 {
	if o == nil || o.MaximumLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *DataFederationQueryLimit) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetName returns the Name field value
func (o *DataFederationQueryLimit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataFederationQueryLimit) SetName(v string) {
	o.Name = v
}

// GetOverrunPolicy returns the OverrunPolicy field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetOverrunPolicy() string {
	if o == nil || o.OverrunPolicy == nil {
		var ret string
		return ret
	}
	return *o.OverrunPolicy
}

// GetOverrunPolicyOk returns a tuple with the OverrunPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetOverrunPolicyOk() (*string, bool) {
	if o == nil || o.OverrunPolicy == nil {
		return nil, false
	}
	return o.OverrunPolicy, true
}

// HasOverrunPolicy returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasOverrunPolicy() bool {
	if o != nil && o.OverrunPolicy != nil {
		return true
	}

	return false
}

// SetOverrunPolicy gets a reference to the given string and assigns it to the OverrunPolicy field.
func (o *DataFederationQueryLimit) SetOverrunPolicy(v string) {
	o.OverrunPolicy = &v
}

// GetValue returns the Value field value
func (o *DataFederationQueryLimit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DataFederationQueryLimit) SetValue(v int64) {
	o.Value = v
}

func (o DataFederationQueryLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentUsage != nil {
		toSerialize["currentUsage"] = o.CurrentUsage
	}
	if o.DefaultLimit != nil {
		toSerialize["defaultLimit"] = o.DefaultLimit
	}
	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if o.MaximumLimit != nil {
		toSerialize["maximumLimit"] = o.MaximumLimit
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OverrunPolicy != nil {
		toSerialize["overrunPolicy"] = o.OverrunPolicy
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDataFederationQueryLimit struct {
	value *DataFederationQueryLimit
	isSet bool
}

func (v NullableDataFederationQueryLimit) Get() *DataFederationQueryLimit {
	return v.value
}

func (v *NullableDataFederationQueryLimit) Set(val *DataFederationQueryLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFederationQueryLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFederationQueryLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFederationQueryLimit(val *DataFederationQueryLimit) *NullableDataFederationQueryLimit {
	return &NullableDataFederationQueryLimit{value: val, isSet: true}
}

func (v NullableDataFederationQueryLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFederationQueryLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


