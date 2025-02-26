// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DataLakeStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeStorage{}

// DataLakeStorage Configuration information for each data store and its mapping to MongoDB Cloud databases.
type DataLakeStorage struct {
	// Array that contains the queryable databases and collections for this data lake.
	Databases []DataLakeDatabase `json:"databases,omitempty"`
	// Array that contains the data stores for the data lake.
	Stores []DataLakeStore `json:"stores,omitempty"`
}

// NewDataLakeStorage instantiates a new DataLakeStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeStorage() *DataLakeStorage {
	this := DataLakeStorage{}
	return &this
}

// NewDataLakeStorageWithDefaults instantiates a new DataLakeStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeStorageWithDefaults() *DataLakeStorage {
	this := DataLakeStorage{}
	return &this
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *DataLakeStorage) GetDatabases() []DataLakeDatabase {
	if o == nil || IsNil(o.Databases) {
		var ret []DataLakeDatabase
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeStorage) GetDatabasesOk() ([]DataLakeDatabase, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *DataLakeStorage) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []DataLakeDatabase and assigns it to the Databases field.
func (o *DataLakeStorage) SetDatabases(v []DataLakeDatabase) {
	o.Databases = v
}

// GetStores returns the Stores field value if set, zero value otherwise.
func (o *DataLakeStorage) GetStores() []DataLakeStore {
	if o == nil || IsNil(o.Stores) {
		var ret []DataLakeStore
		return ret
	}
	return o.Stores
}

// GetStoresOk returns a tuple with the Stores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeStorage) GetStoresOk() ([]DataLakeStore, bool) {
	if o == nil || IsNil(o.Stores) {
		return nil, false
	}
	return o.Stores, true
}

// HasStores returns a boolean if a field has been set.
func (o *DataLakeStorage) HasStores() bool {
	if o != nil && !IsNil(o.Stores) {
		return true
	}

	return false
}

// SetStores gets a reference to the given []DataLakeStore and assigns it to the Stores field.
func (o *DataLakeStorage) SetStores(v []DataLakeStore) {
	o.Stores = v
}

func (o DataLakeStorage) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}
	if !IsNil(o.Stores) {
		toSerialize["stores"] = o.Stores
	}
	return toSerialize, nil
}

type NullableDataLakeStorage struct {
	value *DataLakeStorage
	isSet bool
}

func (v NullableDataLakeStorage) Get() *DataLakeStorage {
	return v.value
}

func (v *NullableDataLakeStorage) Set(val *DataLakeStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeStorage(val *DataLakeStorage) *NullableDataLakeStorage {
	return &NullableDataLakeStorage{value: val, isSet: true}
}

func (v NullableDataLakeStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
