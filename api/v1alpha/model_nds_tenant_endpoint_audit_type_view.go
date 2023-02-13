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

// NDSTenantEndpointAuditTypeView Unique identifier of event type.
type NDSTenantEndpointAuditTypeView string

// List of NDSTenantEndpointAuditTypeView
const (
	NDSTENANTENDPOINTAUDITTYPEVIEW_CREATED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_CREATED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_RESERVED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_RESERVED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_RESERVATION_FAILED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_RESERVATION_FAILED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_UPDATED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_UPDATED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_INITIATING NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_INITIATING"
	NDSTENANTENDPOINTAUDITTYPEVIEW_AVAILABLE NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_AVAILABLE"
	NDSTENANTENDPOINTAUDITTYPEVIEW_FAILED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_FAILED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_DELETING NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_DELETING"
	NDSTENANTENDPOINTAUDITTYPEVIEW_DELETED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_DELETED"
	NDSTENANTENDPOINTAUDITTYPEVIEW_EXPIRED NDSTenantEndpointAuditTypeView = "TENANT_ENDPOINT_EXPIRED"
)

// All allowed values of NDSTenantEndpointAuditTypeView enum
var AllowedNDSTenantEndpointAuditTypeViewEnumValues = []NDSTenantEndpointAuditTypeView{
	"TENANT_ENDPOINT_CREATED",
	"TENANT_ENDPOINT_RESERVED",
	"TENANT_ENDPOINT_RESERVATION_FAILED",
	"TENANT_ENDPOINT_UPDATED",
	"TENANT_ENDPOINT_INITIATING",
	"TENANT_ENDPOINT_AVAILABLE",
	"TENANT_ENDPOINT_FAILED",
	"TENANT_ENDPOINT_DELETING",
	"TENANT_ENDPOINT_DELETED",
	"TENANT_ENDPOINT_EXPIRED",
}

func (v *NDSTenantEndpointAuditTypeView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSTenantEndpointAuditTypeView(value)
	for _, existing := range AllowedNDSTenantEndpointAuditTypeViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSTenantEndpointAuditTypeView", value)
}

// NewNDSTenantEndpointAuditTypeViewFromValue returns a pointer to a valid NDSTenantEndpointAuditTypeView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSTenantEndpointAuditTypeViewFromValue(v string) (*NDSTenantEndpointAuditTypeView, error) {
	ev := NDSTenantEndpointAuditTypeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSTenantEndpointAuditTypeView: valid values are %v", v, AllowedNDSTenantEndpointAuditTypeViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSTenantEndpointAuditTypeView) IsValid() bool {
	for _, existing := range AllowedNDSTenantEndpointAuditTypeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSTenantEndpointAuditTypeView value
func (v NDSTenantEndpointAuditTypeView) Ptr() *NDSTenantEndpointAuditTypeView {
	return &v
}

type NullableNDSTenantEndpointAuditTypeView struct {
	value *NDSTenantEndpointAuditTypeView
	isSet bool
}

func (v NullableNDSTenantEndpointAuditTypeView) Get() *NDSTenantEndpointAuditTypeView {
	return v.value
}

func (v *NullableNDSTenantEndpointAuditTypeView) Set(val *NDSTenantEndpointAuditTypeView) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSTenantEndpointAuditTypeView) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSTenantEndpointAuditTypeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSTenantEndpointAuditTypeView(val *NDSTenantEndpointAuditTypeView) *NullableNDSTenantEndpointAuditTypeView {
	return &NullableNDSTenantEndpointAuditTypeView{value: val, isSet: true}
}

func (v NullableNDSTenantEndpointAuditTypeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSTenantEndpointAuditTypeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

