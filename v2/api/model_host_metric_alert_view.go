/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// HostMetricAlertView Host Metric Alert notifies about changes of measurements or metrics for mongod host.
type HostMetricAlertView struct {
	// Date and time until which this alert has been acknowledged. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.
	AcknowledgedUntil time.Time `json:"acknowledgedUntil"`
	// Comment that a MongoDB Cloud user submitted when acknowledging the alert.
	AcknowledgementComment *string `json:"acknowledgementComment,omitempty"`
	// MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.
	AcknowledgingUsername *string `json:"acknowledgingUsername,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.
	AlertConfigId string `json:"alertConfigId"`
	// Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	CurrentValue *HostMetricValueView `json:"currentValue,omitempty"`
	EventTypeName HostMetricAlertableEventType `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert.
	GroupId *string `json:"groupId,omitempty"`
	// Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets.
	HostnameAndPort *string `json:"hostnameAndPort,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert.
	Id string `json:"id"`
	// Date and time that any notifications were last sent for this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.
	LastNotified *time.Time `json:"lastNotified,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Name of the metric against which Atlas checks the configured `metricThreshold.threshold`.  To learn more about the available metrics, see <a href=\"https://www.mongodb.com/docs/atlas/reference/alert-host-metrics/#std-label-measurement-types\" target=\"_blank\">Host Metrics</a>.  **NOTE**: If you set eventTypeName to OUTSIDE_SERVERLESS_METRIC_THRESHOLD, you can specify only metrics available for serverless. To learn more, see <a href=\"https://dochub.mongodb.org/core/alert-config-serverless-measurements\" target=\"_blank\">Serverless Measurements</a>.
	MetricName *string `json:"metricName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the organization that owns the project to which this alert applies.
	OrgId *string `json:"orgId,omitempty"`
	// Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Date and time that this alert changed to `\"status\" : \"CLOSED\"`. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC. The resource returns this parameter once `\"status\" : \"CLOSED\"`.
	Resolved *time.Time `json:"resolved,omitempty"`
	// State of this alert at the time you requested its details.
	Status string `json:"status"`
	// Date and time when someone last updated this alert. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated time.Time `json:"updated"`
}

// NewHostMetricAlertView instantiates a new HostMetricAlertView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMetricAlertView() *HostMetricAlertView {
	this := HostMetricAlertView{}
	return &this
}

// NewHostMetricAlertViewWithDefaults instantiates a new HostMetricAlertView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMetricAlertViewWithDefaults() *HostMetricAlertView {
	this := HostMetricAlertView{}
	return &this
}

// GetAcknowledgedUntil returns the AcknowledgedUntil field value
func (o *HostMetricAlertView) GetAcknowledgedUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AcknowledgedUntil
}

// GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetAcknowledgedUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgedUntil, true
}

// SetAcknowledgedUntil sets field value
func (o *HostMetricAlertView) SetAcknowledgedUntil(v time.Time) {
	o.AcknowledgedUntil = v
}

// GetAcknowledgementComment returns the AcknowledgementComment field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetAcknowledgementComment() string {
	if o == nil || o.AcknowledgementComment == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgementComment
}

// GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetAcknowledgementCommentOk() (*string, bool) {
	if o == nil || o.AcknowledgementComment == nil {
		return nil, false
	}
	return o.AcknowledgementComment, true
}

// HasAcknowledgementComment returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasAcknowledgementComment() bool {
	if o != nil && o.AcknowledgementComment != nil {
		return true
	}

	return false
}

// SetAcknowledgementComment gets a reference to the given string and assigns it to the AcknowledgementComment field.
func (o *HostMetricAlertView) SetAcknowledgementComment(v string) {
	o.AcknowledgementComment = &v
}

// GetAcknowledgingUsername returns the AcknowledgingUsername field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetAcknowledgingUsername() string {
	if o == nil || o.AcknowledgingUsername == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgingUsername
}

// GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetAcknowledgingUsernameOk() (*string, bool) {
	if o == nil || o.AcknowledgingUsername == nil {
		return nil, false
	}
	return o.AcknowledgingUsername, true
}

// HasAcknowledgingUsername returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasAcknowledgingUsername() bool {
	if o != nil && o.AcknowledgingUsername != nil {
		return true
	}

	return false
}

// SetAcknowledgingUsername gets a reference to the given string and assigns it to the AcknowledgingUsername field.
func (o *HostMetricAlertView) SetAcknowledgingUsername(v string) {
	o.AcknowledgingUsername = &v
}

// GetAlertConfigId returns the AlertConfigId field value
func (o *HostMetricAlertView) GetAlertConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetAlertConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertConfigId, true
}

// SetAlertConfigId sets field value
func (o *HostMetricAlertView) SetAlertConfigId(v string) {
	o.AlertConfigId = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HostMetricAlertView) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value
func (o *HostMetricAlertView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *HostMetricAlertView) SetCreated(v time.Time) {
	o.Created = v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetCurrentValue() HostMetricValueView {
	if o == nil || o.CurrentValue == nil {
		var ret HostMetricValueView
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetCurrentValueOk() (*HostMetricValueView, bool) {
	if o == nil || o.CurrentValue == nil {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasCurrentValue() bool {
	if o != nil && o.CurrentValue != nil {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given HostMetricValueView and assigns it to the CurrentValue field.
func (o *HostMetricAlertView) SetCurrentValue(v HostMetricValueView) {
	o.CurrentValue = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *HostMetricAlertView) GetEventTypeName() HostMetricAlertableEventType {
	if o == nil {
		var ret HostMetricAlertableEventType
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetEventTypeNameOk() (*HostMetricAlertableEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *HostMetricAlertView) SetEventTypeName(v HostMetricAlertableEventType) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *HostMetricAlertView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHostnameAndPort returns the HostnameAndPort field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetHostnameAndPort() string {
	if o == nil || o.HostnameAndPort == nil {
		var ret string
		return ret
	}
	return *o.HostnameAndPort
}

// GetHostnameAndPortOk returns a tuple with the HostnameAndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetHostnameAndPortOk() (*string, bool) {
	if o == nil || o.HostnameAndPort == nil {
		return nil, false
	}
	return o.HostnameAndPort, true
}

// HasHostnameAndPort returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasHostnameAndPort() bool {
	if o != nil && o.HostnameAndPort != nil {
		return true
	}

	return false
}

// SetHostnameAndPort gets a reference to the given string and assigns it to the HostnameAndPort field.
func (o *HostMetricAlertView) SetHostnameAndPort(v string) {
	o.HostnameAndPort = &v
}

// GetId returns the Id field value
func (o *HostMetricAlertView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HostMetricAlertView) SetId(v string) {
	o.Id = v
}

// GetLastNotified returns the LastNotified field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetLastNotified() time.Time {
	if o == nil || o.LastNotified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastNotified
}

// GetLastNotifiedOk returns a tuple with the LastNotified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetLastNotifiedOk() (*time.Time, bool) {
	if o == nil || o.LastNotified == nil {
		return nil, false
	}
	return o.LastNotified, true
}

// HasLastNotified returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasLastNotified() bool {
	if o != nil && o.LastNotified != nil {
		return true
	}

	return false
}

// SetLastNotified gets a reference to the given time.Time and assigns it to the LastNotified field.
func (o *HostMetricAlertView) SetLastNotified(v time.Time) {
	o.LastNotified = &v
}

// GetLinks returns the Links field value
func (o *HostMetricAlertView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *HostMetricAlertView) SetLinks(v []Link) {
	o.Links = v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *HostMetricAlertView) SetMetricName(v string) {
	o.MetricName = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *HostMetricAlertView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetReplicaSetName() string {
	if o == nil || o.ReplicaSetName == nil {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || o.ReplicaSetName == nil {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasReplicaSetName() bool {
	if o != nil && o.ReplicaSetName != nil {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *HostMetricAlertView) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *HostMetricAlertView) GetResolved() time.Time {
	if o == nil || o.Resolved == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetResolvedOk() (*time.Time, bool) {
	if o == nil || o.Resolved == nil {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *HostMetricAlertView) HasResolved() bool {
	if o != nil && o.Resolved != nil {
		return true
	}

	return false
}

// SetResolved gets a reference to the given time.Time and assigns it to the Resolved field.
func (o *HostMetricAlertView) SetResolved(v time.Time) {
	o.Resolved = &v
}

// GetStatus returns the Status field value
func (o *HostMetricAlertView) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HostMetricAlertView) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *HostMetricAlertView) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *HostMetricAlertView) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *HostMetricAlertView) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o HostMetricAlertView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["acknowledgedUntil"] = o.AcknowledgedUntil
	}
	if o.AcknowledgementComment != nil {
		toSerialize["acknowledgementComment"] = o.AcknowledgementComment
	}
	if o.AcknowledgingUsername != nil {
		toSerialize["acknowledgingUsername"] = o.AcknowledgingUsername
	}
	if true {
		toSerialize["alertConfigId"] = o.AlertConfigId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.CurrentValue != nil {
		toSerialize["currentValue"] = o.CurrentValue
	}
	if true {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.HostnameAndPort != nil {
		toSerialize["hostnameAndPort"] = o.HostnameAndPort
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastNotified != nil {
		toSerialize["lastNotified"] = o.LastNotified
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if o.MetricName != nil {
		toSerialize["metricName"] = o.MetricName
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.ReplicaSetName != nil {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if o.Resolved != nil {
		toSerialize["resolved"] = o.Resolved
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableHostMetricAlertView struct {
	value *HostMetricAlertView
	isSet bool
}

func (v NullableHostMetricAlertView) Get() *HostMetricAlertView {
	return v.value
}

func (v *NullableHostMetricAlertView) Set(val *HostMetricAlertView) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMetricAlertView) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMetricAlertView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMetricAlertView(val *HostMetricAlertView) *NullableHostMetricAlertView {
	return &NullableHostMetricAlertView{value: val, isSet: true}
}

func (v NullableHostMetricAlertView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMetricAlertView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


