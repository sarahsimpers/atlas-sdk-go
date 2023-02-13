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

// AppServiceEventTypeView Unique identifier of event type.
type AppServiceEventTypeView string

// List of AppServiceEventTypeView
const (
	APPSERVICEEVENTTYPEVIEW_URL_CONFIRMATION AppServiceEventTypeView = "URL_CONFIRMATION"
	APPSERVICEEVENTTYPEVIEW_SUCCESSFUL_DEPLOY AppServiceEventTypeView = "SUCCESSFUL_DEPLOY"
	APPSERVICEEVENTTYPEVIEW_DEPLOYMENT_FAILURE AppServiceEventTypeView = "DEPLOYMENT_FAILURE"
	APPSERVICEEVENTTYPEVIEW_DEPLOYMENT_MODEL_CHANGE_SUCCESS AppServiceEventTypeView = "DEPLOYMENT_MODEL_CHANGE_SUCCESS"
	APPSERVICEEVENTTYPEVIEW_DEPLOYMENT_MODEL_CHANGE_FAILURE AppServiceEventTypeView = "DEPLOYMENT_MODEL_CHANGE_FAILURE"
	APPSERVICEEVENTTYPEVIEW_REQUEST_RATE_LIMIT AppServiceEventTypeView = "REQUEST_RATE_LIMIT"
	APPSERVICEEVENTTYPEVIEW_LOG_FORWARDER_FAILURE AppServiceEventTypeView = "LOG_FORWARDER_FAILURE"
	APPSERVICEEVENTTYPEVIEW_INSIDE_REALM_METRIC_THRESHOLD AppServiceEventTypeView = "INSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEEVENTTYPEVIEW_OUTSIDE_REALM_METRIC_THRESHOLD AppServiceEventTypeView = "OUTSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEEVENTTYPEVIEW_SYNC_FAILURE AppServiceEventTypeView = "SYNC_FAILURE"
	APPSERVICEEVENTTYPEVIEW_TRIGGER_FAILURE AppServiceEventTypeView = "TRIGGER_FAILURE"
	APPSERVICEEVENTTYPEVIEW_TRIGGER_AUTO_RESUMED AppServiceEventTypeView = "TRIGGER_AUTO_RESUMED"
)

// All allowed values of AppServiceEventTypeView enum
var AllowedAppServiceEventTypeViewEnumValues = []AppServiceEventTypeView{
	"URL_CONFIRMATION",
	"SUCCESSFUL_DEPLOY",
	"DEPLOYMENT_FAILURE",
	"DEPLOYMENT_MODEL_CHANGE_SUCCESS",
	"DEPLOYMENT_MODEL_CHANGE_FAILURE",
	"REQUEST_RATE_LIMIT",
	"LOG_FORWARDER_FAILURE",
	"INSIDE_REALM_METRIC_THRESHOLD",
	"OUTSIDE_REALM_METRIC_THRESHOLD",
	"SYNC_FAILURE",
	"TRIGGER_FAILURE",
	"TRIGGER_AUTO_RESUMED",
}

func (v *AppServiceEventTypeView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppServiceEventTypeView(value)
	for _, existing := range AllowedAppServiceEventTypeViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppServiceEventTypeView", value)
}

// NewAppServiceEventTypeViewFromValue returns a pointer to a valid AppServiceEventTypeView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppServiceEventTypeViewFromValue(v string) (*AppServiceEventTypeView, error) {
	ev := AppServiceEventTypeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppServiceEventTypeView: valid values are %v", v, AllowedAppServiceEventTypeViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppServiceEventTypeView) IsValid() bool {
	for _, existing := range AllowedAppServiceEventTypeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppServiceEventTypeView value
func (v AppServiceEventTypeView) Ptr() *AppServiceEventTypeView {
	return &v
}

type NullableAppServiceEventTypeView struct {
	value *AppServiceEventTypeView
	isSet bool
}

func (v NullableAppServiceEventTypeView) Get() *AppServiceEventTypeView {
	return v.value
}

func (v *NullableAppServiceEventTypeView) Set(val *AppServiceEventTypeView) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceEventTypeView) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceEventTypeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceEventTypeView(val *AppServiceEventTypeView) *NullableAppServiceEventTypeView {
	return &NullableAppServiceEventTypeView{value: val, isSet: true}
}

func (v NullableAppServiceEventTypeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceEventTypeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

