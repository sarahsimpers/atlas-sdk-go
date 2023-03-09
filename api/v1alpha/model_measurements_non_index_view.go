/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the MeasurementsNonIndexView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasurementsNonIndexView{}

// MeasurementsNonIndexView struct for MeasurementsNonIndexView
type MeasurementsNonIndexView struct {
	// Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	End *time.Time `json:"end,omitempty"`
	// Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**.
	Granularity *string `json:"granularity,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the `mongod` or `mongos`.
	GroupId *string `json:"groupId,omitempty"`
	// List that contains the Atlas Search hardware measurements.
	HardwareMeasurements []MeasurementView `json:"hardwareMeasurements,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	ProcessId *string `json:"processId,omitempty"`
	// Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Start *time.Time `json:"start,omitempty"`
	// List that contains the Atlas Search status measurements.
	StatusMeasurements []MeasurementView `json:"statusMeasurements,omitempty"`
}

// NewMeasurementsNonIndexView instantiates a new MeasurementsNonIndexView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementsNonIndexView() *MeasurementsNonIndexView {
	this := MeasurementsNonIndexView{}
	return &this
}

// NewMeasurementsNonIndexViewWithDefaults instantiates a new MeasurementsNonIndexView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementsNonIndexViewWithDefaults() *MeasurementsNonIndexView {
	this := MeasurementsNonIndexView{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *MeasurementsNonIndexView) SetEnd(v time.Time) {
	o.End = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *MeasurementsNonIndexView) SetGranularity(v string) {
	o.Granularity = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *MeasurementsNonIndexView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHardwareMeasurements returns the HardwareMeasurements field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetHardwareMeasurements() []MeasurementView {
	if o == nil || IsNil(o.HardwareMeasurements) {
		var ret []MeasurementView
		return ret
	}
	return o.HardwareMeasurements
}

// GetHardwareMeasurementsOk returns a tuple with the HardwareMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetHardwareMeasurementsOk() ([]MeasurementView, bool) {
	if o == nil || IsNil(o.HardwareMeasurements) {
		return nil, false
	}
	return o.HardwareMeasurements, true
}

// HasHardwareMeasurements returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasHardwareMeasurements() bool {
	if o != nil && !IsNil(o.HardwareMeasurements) {
		return true
	}

	return false
}

// SetHardwareMeasurements gets a reference to the given []MeasurementView and assigns it to the HardwareMeasurements field.
func (o *MeasurementsNonIndexView) SetHardwareMeasurements(v []MeasurementView) {
	o.HardwareMeasurements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *MeasurementsNonIndexView) SetLinks(v []Link) {
	o.Links = v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *MeasurementsNonIndexView) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *MeasurementsNonIndexView) SetStart(v time.Time) {
	o.Start = &v
}

// GetStatusMeasurements returns the StatusMeasurements field value if set, zero value otherwise.
func (o *MeasurementsNonIndexView) GetStatusMeasurements() []MeasurementView {
	if o == nil || IsNil(o.StatusMeasurements) {
		var ret []MeasurementView
		return ret
	}
	return o.StatusMeasurements
}

// GetStatusMeasurementsOk returns a tuple with the StatusMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsNonIndexView) GetStatusMeasurementsOk() ([]MeasurementView, bool) {
	if o == nil || IsNil(o.StatusMeasurements) {
		return nil, false
	}
	return o.StatusMeasurements, true
}

// HasStatusMeasurements returns a boolean if a field has been set.
func (o *MeasurementsNonIndexView) HasStatusMeasurements() bool {
	if o != nil && !IsNil(o.StatusMeasurements) {
		return true
	}

	return false
}

// SetStatusMeasurements gets a reference to the given []MeasurementView and assigns it to the StatusMeasurements field.
func (o *MeasurementsNonIndexView) SetStatusMeasurements(v []MeasurementView) {
	o.StatusMeasurements = v
}

func (o MeasurementsNonIndexView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasurementsNonIndexView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: end is readOnly
	// skip: granularity is readOnly
	// skip: groupId is readOnly
	// skip: hardwareMeasurements is readOnly
	// skip: links is readOnly
	// skip: processId is readOnly
	// skip: start is readOnly
	// skip: statusMeasurements is readOnly
	return toSerialize, nil
}

type NullableMeasurementsNonIndexView struct {
	value *MeasurementsNonIndexView
	isSet bool
}

func (v NullableMeasurementsNonIndexView) Get() *MeasurementsNonIndexView {
	return v.value
}

func (v *NullableMeasurementsNonIndexView) Set(val *MeasurementsNonIndexView) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementsNonIndexView) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementsNonIndexView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementsNonIndexView(val *MeasurementsNonIndexView) *NullableMeasurementsNonIndexView {
	return &NullableMeasurementsNonIndexView{value: val, isSet: true}
}

func (v NullableMeasurementsNonIndexView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementsNonIndexView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

