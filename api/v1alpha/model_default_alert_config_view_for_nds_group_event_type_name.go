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

// DefaultAlertConfigViewForNdsGroupEventTypeName - Incident that triggered this alert.
type DefaultAlertConfigViewForNdsGroupEventTypeName struct {
	String *string
}

// stringAsDefaultAlertConfigViewForNdsGroupEventTypeName is a convenience function that returns string wrapped in DefaultAlertConfigViewForNdsGroupEventTypeName
func StringAsDefaultAlertConfigViewForNdsGroupEventTypeName(v *string) DefaultAlertConfigViewForNdsGroupEventTypeName {
	return DefaultAlertConfigViewForNdsGroupEventTypeName{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultAlertConfigViewForNdsGroupEventTypeName) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into String
        err = json.Unmarshal(data, &dst.String)
        if err == nil {
                jsonstring, _ := json.Marshal(dst.String)
                if string(jsonstring) == "{}" { // empty struct
                        dst.String = nil
                } else {
                        match++
                }
        } else {
                dst.String = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.String = nil

                return fmt.Errorf("data matches more than one schema in oneOf(DefaultAlertConfigViewForNdsGroupEventTypeName)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("data failed to match schemas in oneOf(DefaultAlertConfigViewForNdsGroupEventTypeName)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultAlertConfigViewForNdsGroupEventTypeName) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultAlertConfigViewForNdsGroupEventTypeName) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDefaultAlertConfigViewForNdsGroupEventTypeName struct {
	value *DefaultAlertConfigViewForNdsGroupEventTypeName
	isSet bool
}

func (v NullableDefaultAlertConfigViewForNdsGroupEventTypeName) Get() *DefaultAlertConfigViewForNdsGroupEventTypeName {
	return v.value
}

func (v *NullableDefaultAlertConfigViewForNdsGroupEventTypeName) Set(val *DefaultAlertConfigViewForNdsGroupEventTypeName) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultAlertConfigViewForNdsGroupEventTypeName) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultAlertConfigViewForNdsGroupEventTypeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultAlertConfigViewForNdsGroupEventTypeName(val *DefaultAlertConfigViewForNdsGroupEventTypeName) *NullableDefaultAlertConfigViewForNdsGroupEventTypeName {
	return &NullableDefaultAlertConfigViewForNdsGroupEventTypeName{value: val, isSet: true}
}

func (v NullableDefaultAlertConfigViewForNdsGroupEventTypeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultAlertConfigViewForNdsGroupEventTypeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


