// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the DBResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DBResource{}

// DBResource Namespace to which this database user has access.
type DBResource struct {
	// Flag that indicates whether to grant the action on the cluster resource. If `true`, MongoDB Cloud ignores the **actions.resources.collection** and **actions.resources.db** parameters.
	Cluster bool `json:"cluster"`
	// Human-readable label that identifies the collection on which you grant the action to one MongoDB user. If you don't set this parameter, you grant the action to all collections in the database specified in the **actions.resources.db** parameter. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Collection string `json:"collection"`
	// Human-readable label that identifies the database on which you grant the action to one MongoDB user. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Db string `json:"db"`
}

// NewDBResource instantiates a new DBResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBResource(cluster bool, collection string, db string) *DBResource {
	this := DBResource{}
	this.Cluster = cluster
	this.Collection = collection
	this.Db = db
	return &this
}

// NewDBResourceWithDefaults instantiates a new DBResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBResourceWithDefaults() *DBResource {
	this := DBResource{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *DBResource) GetCluster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *DBResource) GetClusterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *DBResource) SetCluster(v bool) {
	o.Cluster = v
}

// GetCollection returns the Collection field value
func (o *DBResource) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *DBResource) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *DBResource) SetCollection(v string) {
	o.Collection = v
}

// GetDb returns the Db field value
func (o *DBResource) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *DBResource) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *DBResource) SetDb(v string) {
	o.Db = v
}

func (o DBResource) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DBResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["collection"] = o.Collection
	toSerialize["db"] = o.Db
	return toSerialize, nil
}

type NullableDBResource struct {
	value *DBResource
	isSet bool
}

func (v NullableDBResource) Get() *DBResource {
	return v.value
}

func (v *NullableDBResource) Set(val *DBResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDBResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDBResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBResource(val *DBResource) *NullableDBResource {
	return &NullableDBResource{value: val, isSet: true}
}

func (v NullableDBResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
