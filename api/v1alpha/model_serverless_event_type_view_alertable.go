/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"fmt"
)

// ServerlessEventTypeViewAlertable Event type that triggers an alert.
type ServerlessEventTypeViewAlertable string

// List of ServerlessEventTypeViewAlertable
const (
	SERVERLESSEVENTTYPEVIEWALERTABLE_OUTSIDE_SERVERLESS_METRIC_THRESHOLD ServerlessEventTypeViewAlertable = "OUTSIDE_SERVERLESS_METRIC_THRESHOLD"
)

// All allowed values of ServerlessEventTypeViewAlertable enum
var AllowedServerlessEventTypeViewAlertableEnumValues = []ServerlessEventTypeViewAlertable{
	"OUTSIDE_SERVERLESS_METRIC_THRESHOLD",
}

func (v *ServerlessEventTypeViewAlertable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServerlessEventTypeViewAlertable(value)
	for _, existing := range AllowedServerlessEventTypeViewAlertableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServerlessEventTypeViewAlertable", value)
}

// NewServerlessEventTypeViewAlertableFromValue returns a pointer to a valid ServerlessEventTypeViewAlertable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServerlessEventTypeViewAlertableFromValue(v string) (*ServerlessEventTypeViewAlertable, error) {
	ev := ServerlessEventTypeViewAlertable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServerlessEventTypeViewAlertable: valid values are %v", v, AllowedServerlessEventTypeViewAlertableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServerlessEventTypeViewAlertable) IsValid() bool {
	for _, existing := range AllowedServerlessEventTypeViewAlertableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServerlessEventTypeViewAlertable value
func (v ServerlessEventTypeViewAlertable) Ptr() *ServerlessEventTypeViewAlertable {
	return &v
}

type NullableServerlessEventTypeViewAlertable struct {
	value *ServerlessEventTypeViewAlertable
	isSet bool
}

func (v NullableServerlessEventTypeViewAlertable) Get() *ServerlessEventTypeViewAlertable {
	return v.value
}

func (v *NullableServerlessEventTypeViewAlertable) Set(val *ServerlessEventTypeViewAlertable) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessEventTypeViewAlertable) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessEventTypeViewAlertable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessEventTypeViewAlertable(val *ServerlessEventTypeViewAlertable) *NullableServerlessEventTypeViewAlertable {
	return &NullableServerlessEventTypeViewAlertable{value: val, isSet: true}
}

func (v NullableServerlessEventTypeViewAlertable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessEventTypeViewAlertable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

