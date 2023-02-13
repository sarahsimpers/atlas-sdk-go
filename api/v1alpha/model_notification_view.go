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

// NotificationView - One target that MongoDB Cloud sends notifications when an alert triggers.
type NotificationView struct {
	DatadogNotificationView *DatadogNotificationView
	EmailNotificationView *EmailNotificationView
	GroupNotificationView *GroupNotificationView
	HipChatNotificationView *HipChatNotificationView
	MicrosoftTeamsNotificationView *MicrosoftTeamsNotificationView
	NDSNotificationView *NDSNotificationView
	OpsGenieNotificationView *OpsGenieNotificationView
	OrgNotificationView *OrgNotificationView
	PagerDutyNotificationView *PagerDutyNotificationView
	SMSNotificationView *SMSNotificationView
	SlackNotificationView *SlackNotificationView
	SummaryNotificationView *SummaryNotificationView
	TeamNotificationView *TeamNotificationView
	UserNotificationView *UserNotificationView
	VictorOpsNotificationView *VictorOpsNotificationView
	WebhookNotificationView *WebhookNotificationView
}

// DatadogNotificationViewAsNotificationView is a convenience function that returns DatadogNotificationView wrapped in NotificationView
func DatadogNotificationViewAsNotificationView(v *DatadogNotificationView) NotificationView {
	return NotificationView{
		DatadogNotificationView: v,
	}
}

// EmailNotificationViewAsNotificationView is a convenience function that returns EmailNotificationView wrapped in NotificationView
func EmailNotificationViewAsNotificationView(v *EmailNotificationView) NotificationView {
	return NotificationView{
		EmailNotificationView: v,
	}
}

// GroupNotificationViewAsNotificationView is a convenience function that returns GroupNotificationView wrapped in NotificationView
func GroupNotificationViewAsNotificationView(v *GroupNotificationView) NotificationView {
	return NotificationView{
		GroupNotificationView: v,
	}
}

// HipChatNotificationViewAsNotificationView is a convenience function that returns HipChatNotificationView wrapped in NotificationView
func HipChatNotificationViewAsNotificationView(v *HipChatNotificationView) NotificationView {
	return NotificationView{
		HipChatNotificationView: v,
	}
}

// MicrosoftTeamsNotificationViewAsNotificationView is a convenience function that returns MicrosoftTeamsNotificationView wrapped in NotificationView
func MicrosoftTeamsNotificationViewAsNotificationView(v *MicrosoftTeamsNotificationView) NotificationView {
	return NotificationView{
		MicrosoftTeamsNotificationView: v,
	}
}

// NDSNotificationViewAsNotificationView is a convenience function that returns NDSNotificationView wrapped in NotificationView
func NDSNotificationViewAsNotificationView(v *NDSNotificationView) NotificationView {
	return NotificationView{
		NDSNotificationView: v,
	}
}

// OpsGenieNotificationViewAsNotificationView is a convenience function that returns OpsGenieNotificationView wrapped in NotificationView
func OpsGenieNotificationViewAsNotificationView(v *OpsGenieNotificationView) NotificationView {
	return NotificationView{
		OpsGenieNotificationView: v,
	}
}

// OrgNotificationViewAsNotificationView is a convenience function that returns OrgNotificationView wrapped in NotificationView
func OrgNotificationViewAsNotificationView(v *OrgNotificationView) NotificationView {
	return NotificationView{
		OrgNotificationView: v,
	}
}

// PagerDutyNotificationViewAsNotificationView is a convenience function that returns PagerDutyNotificationView wrapped in NotificationView
func PagerDutyNotificationViewAsNotificationView(v *PagerDutyNotificationView) NotificationView {
	return NotificationView{
		PagerDutyNotificationView: v,
	}
}

// SMSNotificationViewAsNotificationView is a convenience function that returns SMSNotificationView wrapped in NotificationView
func SMSNotificationViewAsNotificationView(v *SMSNotificationView) NotificationView {
	return NotificationView{
		SMSNotificationView: v,
	}
}

// SlackNotificationViewAsNotificationView is a convenience function that returns SlackNotificationView wrapped in NotificationView
func SlackNotificationViewAsNotificationView(v *SlackNotificationView) NotificationView {
	return NotificationView{
		SlackNotificationView: v,
	}
}

// SummaryNotificationViewAsNotificationView is a convenience function that returns SummaryNotificationView wrapped in NotificationView
func SummaryNotificationViewAsNotificationView(v *SummaryNotificationView) NotificationView {
	return NotificationView{
		SummaryNotificationView: v,
	}
}

// TeamNotificationViewAsNotificationView is a convenience function that returns TeamNotificationView wrapped in NotificationView
func TeamNotificationViewAsNotificationView(v *TeamNotificationView) NotificationView {
	return NotificationView{
		TeamNotificationView: v,
	}
}

// UserNotificationViewAsNotificationView is a convenience function that returns UserNotificationView wrapped in NotificationView
func UserNotificationViewAsNotificationView(v *UserNotificationView) NotificationView {
	return NotificationView{
		UserNotificationView: v,
	}
}

// VictorOpsNotificationViewAsNotificationView is a convenience function that returns VictorOpsNotificationView wrapped in NotificationView
func VictorOpsNotificationViewAsNotificationView(v *VictorOpsNotificationView) NotificationView {
	return NotificationView{
		VictorOpsNotificationView: v,
	}
}

// WebhookNotificationViewAsNotificationView is a convenience function that returns WebhookNotificationView wrapped in NotificationView
func WebhookNotificationViewAsNotificationView(v *WebhookNotificationView) NotificationView {
	return NotificationView{
		WebhookNotificationView: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationView) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into DatadogNotificationView
        err = json.Unmarshal(data, &dst.DatadogNotificationView)
        if err == nil {
                jsonDatadogNotificationView, _ := json.Marshal(dst.DatadogNotificationView)
                if string(jsonDatadogNotificationView) == "{}" { // empty struct
                        dst.DatadogNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.DatadogNotificationView = nil
        }

        // try to unmarshal data into EmailNotificationView
        err = json.Unmarshal(data, &dst.EmailNotificationView)
        if err == nil {
                jsonEmailNotificationView, _ := json.Marshal(dst.EmailNotificationView)
                if string(jsonEmailNotificationView) == "{}" { // empty struct
                        dst.EmailNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.EmailNotificationView = nil
        }

        // try to unmarshal data into GroupNotificationView
        err = json.Unmarshal(data, &dst.GroupNotificationView)
        if err == nil {
                jsonGroupNotificationView, _ := json.Marshal(dst.GroupNotificationView)
                if string(jsonGroupNotificationView) == "{}" { // empty struct
                        dst.GroupNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.GroupNotificationView = nil
        }

        // try to unmarshal data into HipChatNotificationView
        err = json.Unmarshal(data, &dst.HipChatNotificationView)
        if err == nil {
                jsonHipChatNotificationView, _ := json.Marshal(dst.HipChatNotificationView)
                if string(jsonHipChatNotificationView) == "{}" { // empty struct
                        dst.HipChatNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.HipChatNotificationView = nil
        }

        // try to unmarshal data into MicrosoftTeamsNotificationView
        err = json.Unmarshal(data, &dst.MicrosoftTeamsNotificationView)
        if err == nil {
                jsonMicrosoftTeamsNotificationView, _ := json.Marshal(dst.MicrosoftTeamsNotificationView)
                if string(jsonMicrosoftTeamsNotificationView) == "{}" { // empty struct
                        dst.MicrosoftTeamsNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.MicrosoftTeamsNotificationView = nil
        }

        // try to unmarshal data into NDSNotificationView
        err = json.Unmarshal(data, &dst.NDSNotificationView)
        if err == nil {
                jsonNDSNotificationView, _ := json.Marshal(dst.NDSNotificationView)
                if string(jsonNDSNotificationView) == "{}" { // empty struct
                        dst.NDSNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.NDSNotificationView = nil
        }

        // try to unmarshal data into OpsGenieNotificationView
        err = json.Unmarshal(data, &dst.OpsGenieNotificationView)
        if err == nil {
                jsonOpsGenieNotificationView, _ := json.Marshal(dst.OpsGenieNotificationView)
                if string(jsonOpsGenieNotificationView) == "{}" { // empty struct
                        dst.OpsGenieNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.OpsGenieNotificationView = nil
        }

        // try to unmarshal data into OrgNotificationView
        err = json.Unmarshal(data, &dst.OrgNotificationView)
        if err == nil {
                jsonOrgNotificationView, _ := json.Marshal(dst.OrgNotificationView)
                if string(jsonOrgNotificationView) == "{}" { // empty struct
                        dst.OrgNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.OrgNotificationView = nil
        }

        // try to unmarshal data into PagerDutyNotificationView
        err = json.Unmarshal(data, &dst.PagerDutyNotificationView)
        if err == nil {
                jsonPagerDutyNotificationView, _ := json.Marshal(dst.PagerDutyNotificationView)
                if string(jsonPagerDutyNotificationView) == "{}" { // empty struct
                        dst.PagerDutyNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.PagerDutyNotificationView = nil
        }

        // try to unmarshal data into SMSNotificationView
        err = json.Unmarshal(data, &dst.SMSNotificationView)
        if err == nil {
                jsonSMSNotificationView, _ := json.Marshal(dst.SMSNotificationView)
                if string(jsonSMSNotificationView) == "{}" { // empty struct
                        dst.SMSNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.SMSNotificationView = nil
        }

        // try to unmarshal data into SlackNotificationView
        err = json.Unmarshal(data, &dst.SlackNotificationView)
        if err == nil {
                jsonSlackNotificationView, _ := json.Marshal(dst.SlackNotificationView)
                if string(jsonSlackNotificationView) == "{}" { // empty struct
                        dst.SlackNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.SlackNotificationView = nil
        }

        // try to unmarshal data into SummaryNotificationView
        err = json.Unmarshal(data, &dst.SummaryNotificationView)
        if err == nil {
                jsonSummaryNotificationView, _ := json.Marshal(dst.SummaryNotificationView)
                if string(jsonSummaryNotificationView) == "{}" { // empty struct
                        dst.SummaryNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.SummaryNotificationView = nil
        }

        // try to unmarshal data into TeamNotificationView
        err = json.Unmarshal(data, &dst.TeamNotificationView)
        if err == nil {
                jsonTeamNotificationView, _ := json.Marshal(dst.TeamNotificationView)
                if string(jsonTeamNotificationView) == "{}" { // empty struct
                        dst.TeamNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.TeamNotificationView = nil
        }

        // try to unmarshal data into UserNotificationView
        err = json.Unmarshal(data, &dst.UserNotificationView)
        if err == nil {
                jsonUserNotificationView, _ := json.Marshal(dst.UserNotificationView)
                if string(jsonUserNotificationView) == "{}" { // empty struct
                        dst.UserNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.UserNotificationView = nil
        }

        // try to unmarshal data into VictorOpsNotificationView
        err = json.Unmarshal(data, &dst.VictorOpsNotificationView)
        if err == nil {
                jsonVictorOpsNotificationView, _ := json.Marshal(dst.VictorOpsNotificationView)
                if string(jsonVictorOpsNotificationView) == "{}" { // empty struct
                        dst.VictorOpsNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.VictorOpsNotificationView = nil
        }

        // try to unmarshal data into WebhookNotificationView
        err = json.Unmarshal(data, &dst.WebhookNotificationView)
        if err == nil {
                jsonWebhookNotificationView, _ := json.Marshal(dst.WebhookNotificationView)
                if string(jsonWebhookNotificationView) == "{}" { // empty struct
                        dst.WebhookNotificationView = nil
                } else {
                        match++
                }
        } else {
                dst.WebhookNotificationView = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.DatadogNotificationView = nil
                dst.EmailNotificationView = nil
                dst.GroupNotificationView = nil
                dst.HipChatNotificationView = nil
                dst.MicrosoftTeamsNotificationView = nil
                dst.NDSNotificationView = nil
                dst.OpsGenieNotificationView = nil
                dst.OrgNotificationView = nil
                dst.PagerDutyNotificationView = nil
                dst.SMSNotificationView = nil
                dst.SlackNotificationView = nil
                dst.SummaryNotificationView = nil
                dst.TeamNotificationView = nil
                dst.UserNotificationView = nil
                dst.VictorOpsNotificationView = nil
                dst.WebhookNotificationView = nil

                return fmt.Errorf("data matches more than one schema in oneOf(NotificationView)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("data failed to match schemas in oneOf(NotificationView)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationView) MarshalJSON() ([]byte, error) {
	if src.DatadogNotificationView != nil {
		return json.Marshal(&src.DatadogNotificationView)
	}

	if src.EmailNotificationView != nil {
		return json.Marshal(&src.EmailNotificationView)
	}

	if src.GroupNotificationView != nil {
		return json.Marshal(&src.GroupNotificationView)
	}

	if src.HipChatNotificationView != nil {
		return json.Marshal(&src.HipChatNotificationView)
	}

	if src.MicrosoftTeamsNotificationView != nil {
		return json.Marshal(&src.MicrosoftTeamsNotificationView)
	}

	if src.NDSNotificationView != nil {
		return json.Marshal(&src.NDSNotificationView)
	}

	if src.OpsGenieNotificationView != nil {
		return json.Marshal(&src.OpsGenieNotificationView)
	}

	if src.OrgNotificationView != nil {
		return json.Marshal(&src.OrgNotificationView)
	}

	if src.PagerDutyNotificationView != nil {
		return json.Marshal(&src.PagerDutyNotificationView)
	}

	if src.SMSNotificationView != nil {
		return json.Marshal(&src.SMSNotificationView)
	}

	if src.SlackNotificationView != nil {
		return json.Marshal(&src.SlackNotificationView)
	}

	if src.SummaryNotificationView != nil {
		return json.Marshal(&src.SummaryNotificationView)
	}

	if src.TeamNotificationView != nil {
		return json.Marshal(&src.TeamNotificationView)
	}

	if src.UserNotificationView != nil {
		return json.Marshal(&src.UserNotificationView)
	}

	if src.VictorOpsNotificationView != nil {
		return json.Marshal(&src.VictorOpsNotificationView)
	}

	if src.WebhookNotificationView != nil {
		return json.Marshal(&src.WebhookNotificationView)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationView) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DatadogNotificationView != nil {
		return obj.DatadogNotificationView
	}

	if obj.EmailNotificationView != nil {
		return obj.EmailNotificationView
	}

	if obj.GroupNotificationView != nil {
		return obj.GroupNotificationView
	}

	if obj.HipChatNotificationView != nil {
		return obj.HipChatNotificationView
	}

	if obj.MicrosoftTeamsNotificationView != nil {
		return obj.MicrosoftTeamsNotificationView
	}

	if obj.NDSNotificationView != nil {
		return obj.NDSNotificationView
	}

	if obj.OpsGenieNotificationView != nil {
		return obj.OpsGenieNotificationView
	}

	if obj.OrgNotificationView != nil {
		return obj.OrgNotificationView
	}

	if obj.PagerDutyNotificationView != nil {
		return obj.PagerDutyNotificationView
	}

	if obj.SMSNotificationView != nil {
		return obj.SMSNotificationView
	}

	if obj.SlackNotificationView != nil {
		return obj.SlackNotificationView
	}

	if obj.SummaryNotificationView != nil {
		return obj.SummaryNotificationView
	}

	if obj.TeamNotificationView != nil {
		return obj.TeamNotificationView
	}

	if obj.UserNotificationView != nil {
		return obj.UserNotificationView
	}

	if obj.VictorOpsNotificationView != nil {
		return obj.VictorOpsNotificationView
	}

	if obj.WebhookNotificationView != nil {
		return obj.WebhookNotificationView
	}

	// all schemas are nil
	return nil
}

type NullableNotificationView struct {
	value *NotificationView
	isSet bool
}

func (v NullableNotificationView) Get() *NotificationView {
	return v.value
}

func (v *NullableNotificationView) Set(val *NotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationView(val *NotificationView) *NullableNotificationView {
	return &NullableNotificationView{value: val, isSet: true}
}

func (v NullableNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


