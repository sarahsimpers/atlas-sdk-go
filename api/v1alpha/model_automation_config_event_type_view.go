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

// AutomationConfigEventTypeView Unique identifier of event type.
type AutomationConfigEventTypeView string

// List of AutomationConfigEventTypeView
const (
	AUTOMATIONCONFIGEVENTTYPEVIEW_AUTOMATION_CONFIG_PUBLISHED_AUDIT AutomationConfigEventTypeView = "AUTOMATION_CONFIG_PUBLISHED_AUDIT"
)

// All allowed values of AutomationConfigEventTypeView enum
var AllowedAutomationConfigEventTypeViewEnumValues = []AutomationConfigEventTypeView{
	"AUTOMATION_CONFIG_PUBLISHED_AUDIT",
}

func (v *AutomationConfigEventTypeView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutomationConfigEventTypeView(value)
	for _, existing := range AllowedAutomationConfigEventTypeViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutomationConfigEventTypeView", value)
}

// NewAutomationConfigEventTypeViewFromValue returns a pointer to a valid AutomationConfigEventTypeView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutomationConfigEventTypeViewFromValue(v string) (*AutomationConfigEventTypeView, error) {
	ev := AutomationConfigEventTypeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutomationConfigEventTypeView: valid values are %v", v, AllowedAutomationConfigEventTypeViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutomationConfigEventTypeView) IsValid() bool {
	for _, existing := range AllowedAutomationConfigEventTypeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutomationConfigEventTypeView value
func (v AutomationConfigEventTypeView) Ptr() *AutomationConfigEventTypeView {
	return &v
}

type NullableAutomationConfigEventTypeView struct {
	value *AutomationConfigEventTypeView
	isSet bool
}

func (v NullableAutomationConfigEventTypeView) Get() *AutomationConfigEventTypeView {
	return v.value
}

func (v *NullableAutomationConfigEventTypeView) Set(val *AutomationConfigEventTypeView) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationConfigEventTypeView) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationConfigEventTypeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationConfigEventTypeView(val *AutomationConfigEventTypeView) *NullableAutomationConfigEventTypeView {
	return &NullableAutomationConfigEventTypeView{value: val, isSet: true}
}

func (v NullableAutomationConfigEventTypeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationConfigEventTypeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

