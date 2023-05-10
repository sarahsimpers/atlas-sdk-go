// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// AlertAuditType Unique identifier of event type.
type AlertAuditType string

// List of AlertAuditType
const (
	ALERTAUDITTYPE_ACKNOWLEDGED_AUDIT   AlertAuditType = "ALERT_ACKNOWLEDGED_AUDIT"
	ALERTAUDITTYPE_UNACKNOWLEDGED_AUDIT AlertAuditType = "ALERT_UNACKNOWLEDGED_AUDIT"
)

// All allowed values of AlertAuditType enum
var AllowedAlertAuditTypeEnumValues = []AlertAuditType{
	"ALERT_ACKNOWLEDGED_AUDIT",
	"ALERT_UNACKNOWLEDGED_AUDIT",
}

func (v *AlertAuditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertAuditType(value)
	for _, existing := range AllowedAlertAuditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertAuditType", value)
}

// NewAlertAuditTypeFromValue returns a pointer to a valid AlertAuditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertAuditTypeFromValue(v string) (*AlertAuditType, error) {
	ev := AlertAuditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertAuditType: valid values are %v", v, AllowedAlertAuditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertAuditType) IsValid() bool {
	for _, existing := range AllowedAlertAuditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertAuditType value
func (v AlertAuditType) Ptr() *AlertAuditType {
	return &v
}