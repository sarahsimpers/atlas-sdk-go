// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

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
