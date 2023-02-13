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

// NDSAuditTypeViewForOrg Unique identifier of event type.
type NDSAuditTypeViewForOrg string

// List of NDSAuditTypeViewForOrg
const (
	NDSAUDITTYPEVIEWFORORG_ORG_LIMIT_UPDATED NDSAuditTypeViewForOrg = "ORG_LIMIT_UPDATED"
)

// All allowed values of NDSAuditTypeViewForOrg enum
var AllowedNDSAuditTypeViewForOrgEnumValues = []NDSAuditTypeViewForOrg{
	"ORG_LIMIT_UPDATED",
}

func (v *NDSAuditTypeViewForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSAuditTypeViewForOrg(value)
	for _, existing := range AllowedNDSAuditTypeViewForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSAuditTypeViewForOrg", value)
}

// NewNDSAuditTypeViewForOrgFromValue returns a pointer to a valid NDSAuditTypeViewForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSAuditTypeViewForOrgFromValue(v string) (*NDSAuditTypeViewForOrg, error) {
	ev := NDSAuditTypeViewForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSAuditTypeViewForOrg: valid values are %v", v, AllowedNDSAuditTypeViewForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSAuditTypeViewForOrg) IsValid() bool {
	for _, existing := range AllowedNDSAuditTypeViewForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSAuditTypeViewForOrg value
func (v NDSAuditTypeViewForOrg) Ptr() *NDSAuditTypeViewForOrg {
	return &v
}

type NullableNDSAuditTypeViewForOrg struct {
	value *NDSAuditTypeViewForOrg
	isSet bool
}

func (v NullableNDSAuditTypeViewForOrg) Get() *NDSAuditTypeViewForOrg {
	return v.value
}

func (v *NullableNDSAuditTypeViewForOrg) Set(val *NDSAuditTypeViewForOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSAuditTypeViewForOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSAuditTypeViewForOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSAuditTypeViewForOrg(val *NDSAuditTypeViewForOrg) *NullableNDSAuditTypeViewForOrg {
	return &NullableNDSAuditTypeViewForOrg{value: val, isSet: true}
}

func (v NullableNDSAuditTypeViewForOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSAuditTypeViewForOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

