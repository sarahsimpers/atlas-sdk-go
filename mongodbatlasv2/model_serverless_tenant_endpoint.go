/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// ServerlessTenantEndpoint - struct for ServerlessTenantEndpoint
type ServerlessTenantEndpoint struct {
	ServerlessAWSTenantEndpoint *ServerlessAWSTenantEndpoint
	ServerlessAzureTenantEndpoint *ServerlessAzureTenantEndpoint
}

// ServerlessAWSTenantEndpointAsServerlessTenantEndpoint is a convenience function that returns ServerlessAWSTenantEndpoint wrapped in ServerlessTenantEndpoint
func ServerlessAWSTenantEndpointAsServerlessTenantEndpoint(v *ServerlessAWSTenantEndpoint) ServerlessTenantEndpoint {
	return ServerlessTenantEndpoint{
		ServerlessAWSTenantEndpoint: v,
	}
}

// ServerlessAzureTenantEndpointAsServerlessTenantEndpoint is a convenience function that returns ServerlessAzureTenantEndpoint wrapped in ServerlessTenantEndpoint
func ServerlessAzureTenantEndpointAsServerlessTenantEndpoint(v *ServerlessAzureTenantEndpoint) ServerlessTenantEndpoint {
	return ServerlessTenantEndpoint{
		ServerlessAzureTenantEndpoint: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerlessTenantEndpoint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServerlessAWSTenantEndpoint
	err = json.Unmarshal(data, &dst.ServerlessAWSTenantEndpoint)
	if err == nil {
		jsonServerlessAWSTenantEndpoint, _ := json.Marshal(dst.ServerlessAWSTenantEndpoint)
		if string(jsonServerlessAWSTenantEndpoint) == "{}" { // empty struct
			dst.ServerlessAWSTenantEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.ServerlessAWSTenantEndpoint = nil
	}

	// try to unmarshal data into ServerlessAzureTenantEndpoint
	err = json.Unmarshal(data, &dst.ServerlessAzureTenantEndpoint)
	if err == nil {
		jsonServerlessAzureTenantEndpoint, _ := json.Marshal(dst.ServerlessAzureTenantEndpoint)
		if string(jsonServerlessAzureTenantEndpoint) == "{}" { // empty struct
			dst.ServerlessAzureTenantEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.ServerlessAzureTenantEndpoint = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ServerlessAWSTenantEndpoint = nil
		dst.ServerlessAzureTenantEndpoint = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServerlessTenantEndpoint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServerlessTenantEndpoint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerlessTenantEndpoint) MarshalJSON() ([]byte, error) {
	if src.ServerlessAWSTenantEndpoint != nil {
		return json.Marshal(&src.ServerlessAWSTenantEndpoint)
	}

	if src.ServerlessAzureTenantEndpoint != nil {
		return json.Marshal(&src.ServerlessAzureTenantEndpoint)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerlessTenantEndpoint) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ServerlessAWSTenantEndpoint != nil {
		return obj.ServerlessAWSTenantEndpoint
	}

	if obj.ServerlessAzureTenantEndpoint != nil {
		return obj.ServerlessAzureTenantEndpoint
	}

	// all schemas are nil
	return nil
}

type NullableServerlessTenantEndpoint struct {
	value *ServerlessTenantEndpoint
	isSet bool
}

func (v NullableServerlessTenantEndpoint) Get() *ServerlessTenantEndpoint {
	return v.value
}

func (v *NullableServerlessTenantEndpoint) Set(val *ServerlessTenantEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessTenantEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessTenantEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessTenantEndpoint(val *ServerlessTenantEndpoint) *NullableServerlessTenantEndpoint {
	return &NullableServerlessTenantEndpoint{value: val, isSet: true}
}

func (v NullableServerlessTenantEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessTenantEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

