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

// EmailNotificationView Email notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type EmailNotificationView struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"EMAIL\"`. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (`\"notifications.[n].typeName\" : \"USER\"`) - MongoDB Cloud users with specific project roles (`\"notifications.[n].typeName\" : \"GROUP\"`) - MongoDB Cloud users with specific organization roles (`\"notifications.[n].typeName\" : \"ORG\"`) - MongoDB Cloud teams (`\"notifications.[n].typeName\" : \"TEAM\"`)  To send emails to one MongoDB Cloud user or grouping of users, set the `notifications.[n].emailEnabled` parameter.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewEmailNotificationView instantiates a new EmailNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationView() *EmailNotificationView {
	this := EmailNotificationView{}
	return &this
}

// NewEmailNotificationViewWithDefaults instantiates a new EmailNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationViewWithDefaults() *EmailNotificationView {
	this := EmailNotificationView{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *EmailNotificationView) GetDelayMin() int32 {
	if o == nil || o.DelayMin == nil {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || o.DelayMin == nil {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *EmailNotificationView) HasDelayMin() bool {
	if o != nil && o.DelayMin != nil {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *EmailNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *EmailNotificationView) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationView) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *EmailNotificationView) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *EmailNotificationView) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *EmailNotificationView) GetIntervalMin() int32 {
	if o == nil || o.IntervalMin == nil {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || o.IntervalMin == nil {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *EmailNotificationView) HasIntervalMin() bool {
	if o != nil && o.IntervalMin != nil {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *EmailNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *EmailNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *EmailNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

func (o EmailNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelayMin != nil {
		toSerialize["delayMin"] = o.DelayMin
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.IntervalMin != nil {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationView struct {
	value *EmailNotificationView
	isSet bool
}

func (v NullableEmailNotificationView) Get() *EmailNotificationView {
	return v.value
}

func (v *NullableEmailNotificationView) Set(val *EmailNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationView(val *EmailNotificationView) *NullableEmailNotificationView {
	return &NullableEmailNotificationView{value: val, isSet: true}
}

func (v NullableEmailNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


