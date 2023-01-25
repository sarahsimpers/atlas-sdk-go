/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LiveMigrationResponseView struct for LiveMigrationResponseView
type LiveMigrationResponseView struct {
	// Unique 24-hexadecimal digit string that identifies the migration job.
	Id *string `json:"_id,omitempty"`
	// Replication lag between the source and destination clusters. Atlas returns this setting only during an active migration, before the cutover phase.
	LagTimeSeconds NullableInt64 `json:"lagTimeSeconds,omitempty"`
	// List of hosts running MongoDB Agents. These Agents can transfer your MongoDB data between one source and one target cluster.
	MigrationHosts []string `json:"migrationHosts,omitempty"`
	// Flag that indicates the migrated cluster can be cut over to MongoDB Atlas.
	ReadyForCutover *bool `json:"readyForCutover,omitempty"`
	// Progress made in migrating one cluster to MongoDB Atlas.  | Status   | Explanation | |----------|-------------| | NEW      | Someone scheduled a local cluster migration to MongoDB Atlas. | | FAILED   | The cluster migration to MongoDB Atlas failed.                | | COMPLETE | The cluster migration to MongoDB Atlas succeeded.             | | EXPIRED  | MongoDB Atlas prepares to begin the cut over of the migrating cluster when source and target clusters have almost synchronized. If `\"readyForCutover\" : true`, this synchronization starts a timer of 120 hours. You can extend this timer. If the timer expires, MongoDB Atlas returns this status. | | WORKING  | The cluster migration to MongoDB Atlas is performing one of the following tasks:<ul><li>Preparing connections to source and target clusters</li><li>Replicating data from source to target</li><li>Verifying MongoDB Atlas connection settings</li><li>Stopping replication after the cut over</li></ul> | 
	Status *string `json:"status,omitempty"`
}

// NewLiveMigrationResponseView instantiates a new LiveMigrationResponseView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveMigrationResponseView() *LiveMigrationResponseView {
	this := LiveMigrationResponseView{}
	return &this
}

// NewLiveMigrationResponseViewWithDefaults instantiates a new LiveMigrationResponseView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveMigrationResponseViewWithDefaults() *LiveMigrationResponseView {
	this := LiveMigrationResponseView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveMigrationResponseView) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponseView) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveMigrationResponseView) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveMigrationResponseView) SetId(v string) {
	o.Id = &v
}

// GetLagTimeSeconds returns the LagTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveMigrationResponseView) GetLagTimeSeconds() int64 {
	if o == nil || o.LagTimeSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LagTimeSeconds.Get()
}

// GetLagTimeSecondsOk returns a tuple with the LagTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveMigrationResponseView) GetLagTimeSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LagTimeSeconds.Get(), o.LagTimeSeconds.IsSet()
}

// HasLagTimeSeconds returns a boolean if a field has been set.
func (o *LiveMigrationResponseView) HasLagTimeSeconds() bool {
	if o != nil && o.LagTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetLagTimeSeconds gets a reference to the given NullableInt64 and assigns it to the LagTimeSeconds field.
func (o *LiveMigrationResponseView) SetLagTimeSeconds(v int64) {
	o.LagTimeSeconds.Set(&v)
}
// SetLagTimeSecondsNil sets the value for LagTimeSeconds to be an explicit nil
func (o *LiveMigrationResponseView) SetLagTimeSecondsNil() {
	o.LagTimeSeconds.Set(nil)
}

// UnsetLagTimeSeconds ensures that no value is present for LagTimeSeconds, not even an explicit nil
func (o *LiveMigrationResponseView) UnsetLagTimeSeconds() {
	o.LagTimeSeconds.Unset()
}

// GetMigrationHosts returns the MigrationHosts field value if set, zero value otherwise.
func (o *LiveMigrationResponseView) GetMigrationHosts() []string {
	if o == nil || o.MigrationHosts == nil {
		var ret []string
		return ret
	}
	return o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponseView) GetMigrationHostsOk() ([]string, bool) {
	if o == nil || o.MigrationHosts == nil {
		return nil, false
	}
	return o.MigrationHosts, true
}

// HasMigrationHosts returns a boolean if a field has been set.
func (o *LiveMigrationResponseView) HasMigrationHosts() bool {
	if o != nil && o.MigrationHosts != nil {
		return true
	}

	return false
}

// SetMigrationHosts gets a reference to the given []string and assigns it to the MigrationHosts field.
func (o *LiveMigrationResponseView) SetMigrationHosts(v []string) {
	o.MigrationHosts = v
}

// GetReadyForCutover returns the ReadyForCutover field value if set, zero value otherwise.
func (o *LiveMigrationResponseView) GetReadyForCutover() bool {
	if o == nil || o.ReadyForCutover == nil {
		var ret bool
		return ret
	}
	return *o.ReadyForCutover
}

// GetReadyForCutoverOk returns a tuple with the ReadyForCutover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponseView) GetReadyForCutoverOk() (*bool, bool) {
	if o == nil || o.ReadyForCutover == nil {
		return nil, false
	}
	return o.ReadyForCutover, true
}

// HasReadyForCutover returns a boolean if a field has been set.
func (o *LiveMigrationResponseView) HasReadyForCutover() bool {
	if o != nil && o.ReadyForCutover != nil {
		return true
	}

	return false
}

// SetReadyForCutover gets a reference to the given bool and assigns it to the ReadyForCutover field.
func (o *LiveMigrationResponseView) SetReadyForCutover(v bool) {
	o.ReadyForCutover = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveMigrationResponseView) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponseView) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveMigrationResponseView) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiveMigrationResponseView) SetStatus(v string) {
	o.Status = &v
}

func (o LiveMigrationResponseView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.LagTimeSeconds.IsSet() {
		toSerialize["lagTimeSeconds"] = o.LagTimeSeconds.Get()
	}
	if o.MigrationHosts != nil {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	if o.ReadyForCutover != nil {
		toSerialize["readyForCutover"] = o.ReadyForCutover
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableLiveMigrationResponseView struct {
	value *LiveMigrationResponseView
	isSet bool
}

func (v NullableLiveMigrationResponseView) Get() *LiveMigrationResponseView {
	return v.value
}

func (v *NullableLiveMigrationResponseView) Set(val *LiveMigrationResponseView) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveMigrationResponseView) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveMigrationResponseView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveMigrationResponseView(val *LiveMigrationResponseView) *NullableLiveMigrationResponseView {
	return &NullableLiveMigrationResponseView{value: val, isSet: true}
}

func (v NullableLiveMigrationResponseView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveMigrationResponseView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


