/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// BillingEventTypeForOrg Unique identifier of event type.
type BillingEventTypeForOrg string

// List of BillingEventTypeForOrg
const (
	BILLINGEVENTTYPEFORORG_CHARGE_SUCCEEDED BillingEventTypeForOrg = "CHARGE_SUCCEEDED"
	BILLINGEVENTTYPEFORORG_CHARGE_FAILED BillingEventTypeForOrg = "CHARGE_FAILED"
	BILLINGEVENTTYPEFORORG_CHARGE_PROCESSING BillingEventTypeForOrg = "CHARGE_PROCESSING"
	BILLINGEVENTTYPEFORORG_CHARGE_PENDING_REVERSAL BillingEventTypeForOrg = "CHARGE_PENDING_REVERSAL"
	BILLINGEVENTTYPEFORORG_BRAINTREE_CHARGE_FAILED BillingEventTypeForOrg = "BRAINTREE_CHARGE_FAILED"
	BILLINGEVENTTYPEFORORG_INVOICE_CLOSED BillingEventTypeForOrg = "INVOICE_CLOSED"
	BILLINGEVENTTYPEFORORG_CHECK_PAYMENT_RECEIVED BillingEventTypeForOrg = "CHECK_PAYMENT_RECEIVED"
	BILLINGEVENTTYPEFORORG_WIRE_TRANSFER_PAYMENT_RECEIVED BillingEventTypeForOrg = "WIRE_TRANSFER_PAYMENT_RECEIVED"
	BILLINGEVENTTYPEFORORG_DISCOUNT_APPLIED BillingEventTypeForOrg = "DISCOUNT_APPLIED"
	BILLINGEVENTTYPEFORORG_CREDIT_ISSUED BillingEventTypeForOrg = "CREDIT_ISSUED"
	BILLINGEVENTTYPEFORORG_CREDIT_PULLED_FWD BillingEventTypeForOrg = "CREDIT_PULLED_FWD"
	BILLINGEVENTTYPEFORORG_CREDIT_END_DATE_MODIFIED BillingEventTypeForOrg = "CREDIT_END_DATE_MODIFIED"
	BILLINGEVENTTYPEFORORG_PROMO_CODE_APPLIED BillingEventTypeForOrg = "PROMO_CODE_APPLIED"
	BILLINGEVENTTYPEFORORG_PAYMENT_FORGIVEN BillingEventTypeForOrg = "PAYMENT_FORGIVEN"
	BILLINGEVENTTYPEFORORG_REFUND_ISSUED BillingEventTypeForOrg = "REFUND_ISSUED"
	BILLINGEVENTTYPEFORORG_ACCOUNT_DOWNGRADED BillingEventTypeForOrg = "ACCOUNT_DOWNGRADED"
	BILLINGEVENTTYPEFORORG_ACCOUNT_UPGRADED BillingEventTypeForOrg = "ACCOUNT_UPGRADED"
	BILLINGEVENTTYPEFORORG_ACCOUNT_MODIFIED BillingEventTypeForOrg = "ACCOUNT_MODIFIED"
	BILLINGEVENTTYPEFORORG_SUPPORT_PLAN_ACTIVATED BillingEventTypeForOrg = "SUPPORT_PLAN_ACTIVATED"
	BILLINGEVENTTYPEFORORG_SUPPORT_PLAN_CANCELLED BillingEventTypeForOrg = "SUPPORT_PLAN_CANCELLED"
	BILLINGEVENTTYPEFORORG_SUPPORT_PLAN_CANCELLATION_SCHEDULED BillingEventTypeForOrg = "SUPPORT_PLAN_CANCELLATION_SCHEDULED"
	BILLINGEVENTTYPEFORORG_INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC BillingEventTypeForOrg = "INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC"
	BILLINGEVENTTYPEFORORG_INVOICE_ADDRESS_CHANGED BillingEventTypeForOrg = "INVOICE_ADDRESS_CHANGED"
	BILLINGEVENTTYPEFORORG_INVOICE_ADDRESS_ADDED BillingEventTypeForOrg = "INVOICE_ADDRESS_ADDED"
	BILLINGEVENTTYPEFORORG_PREPAID_PLAN_ACTIVATED BillingEventTypeForOrg = "PREPAID_PLAN_ACTIVATED"
	BILLINGEVENTTYPEFORORG_ELASTIC_INVOICING_MODE_ACTIVATED BillingEventTypeForOrg = "ELASTIC_INVOICING_MODE_ACTIVATED"
	BILLINGEVENTTYPEFORORG_ELASTIC_INVOICING_MODE_DEACTIVATED BillingEventTypeForOrg = "ELASTIC_INVOICING_MODE_DEACTIVATED"
	BILLINGEVENTTYPEFORORG_TERMINATE_PAID_SERVICES BillingEventTypeForOrg = "TERMINATE_PAID_SERVICES"
	BILLINGEVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_ADDED BillingEventTypeForOrg = "BILLING_EMAIL_ADDRESS_ADDED"
	BILLINGEVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_CHANGED BillingEventTypeForOrg = "BILLING_EMAIL_ADDRESS_CHANGED"
	BILLINGEVENTTYPEFORORG_BILLING_EMAIL_ADDRESS_REMOVED BillingEventTypeForOrg = "BILLING_EMAIL_ADDRESS_REMOVED"
	BILLINGEVENTTYPEFORORG_AWS_BILLING_ACCOUNT_CREDIT_ISSUED BillingEventTypeForOrg = "AWS_BILLING_ACCOUNT_CREDIT_ISSUED"
	BILLINGEVENTTYPEFORORG_GCP_BILLING_ACCOUNT_CREDIT_ISSUED BillingEventTypeForOrg = "GCP_BILLING_ACCOUNT_CREDIT_ISSUED"
	BILLINGEVENTTYPEFORORG_CREDIT_SFOLI_MODIFIED BillingEventTypeForOrg = "CREDIT_SFOLI_MODIFIED"
	BILLINGEVENTTYPEFORORG_CREDIT_AMOUNT_MODIFIED BillingEventTypeForOrg = "CREDIT_AMOUNT_MODIFIED"
	BILLINGEVENTTYPEFORORG_PREPAID_PLAN_MODIFIED BillingEventTypeForOrg = "PREPAID_PLAN_MODIFIED"
	BILLINGEVENTTYPEFORORG_AWS_USAGE_REPORTED BillingEventTypeForOrg = "AWS_USAGE_REPORTED"
	BILLINGEVENTTYPEFORORG_AZURE_USAGE_REPORTED BillingEventTypeForOrg = "AZURE_USAGE_REPORTED"
	BILLINGEVENTTYPEFORORG_GCP_USAGE_REPORTED BillingEventTypeForOrg = "GCP_USAGE_REPORTED"
	BILLINGEVENTTYPEFORORG_BECAME_PAYING_ORG BillingEventTypeForOrg = "BECAME_PAYING_ORG"
	BILLINGEVENTTYPEFORORG_NEW_LINKED_ORG BillingEventTypeForOrg = "NEW_LINKED_ORG"
	BILLINGEVENTTYPEFORORG_UNLINKED_ORG BillingEventTypeForOrg = "UNLINKED_ORG"
	BILLINGEVENTTYPEFORORG_ORG_LINKED_TO_PAYING_ORG BillingEventTypeForOrg = "ORG_LINKED_TO_PAYING_ORG"
	BILLINGEVENTTYPEFORORG_ORG_UNLINKED_FROM_PAYING_ORG BillingEventTypeForOrg = "ORG_UNLINKED_FROM_PAYING_ORG"
	BILLINGEVENTTYPEFORORG_PAYMENT_UPDATED_THROUGH_API BillingEventTypeForOrg = "PAYMENT_UPDATED_THROUGH_API"
)

// All allowed values of BillingEventTypeForOrg enum
var AllowedBillingEventTypeForOrgEnumValues = []BillingEventTypeForOrg{
	"CHARGE_SUCCEEDED",
	"CHARGE_FAILED",
	"CHARGE_PROCESSING",
	"CHARGE_PENDING_REVERSAL",
	"BRAINTREE_CHARGE_FAILED",
	"INVOICE_CLOSED",
	"CHECK_PAYMENT_RECEIVED",
	"WIRE_TRANSFER_PAYMENT_RECEIVED",
	"DISCOUNT_APPLIED",
	"CREDIT_ISSUED",
	"CREDIT_PULLED_FWD",
	"CREDIT_END_DATE_MODIFIED",
	"PROMO_CODE_APPLIED",
	"PAYMENT_FORGIVEN",
	"REFUND_ISSUED",
	"ACCOUNT_DOWNGRADED",
	"ACCOUNT_UPGRADED",
	"ACCOUNT_MODIFIED",
	"SUPPORT_PLAN_ACTIVATED",
	"SUPPORT_PLAN_CANCELLED",
	"SUPPORT_PLAN_CANCELLATION_SCHEDULED",
	"INITIATE_SALESFORCE_SERVICE_CLOUD_SYNC",
	"INVOICE_ADDRESS_CHANGED",
	"INVOICE_ADDRESS_ADDED",
	"PREPAID_PLAN_ACTIVATED",
	"ELASTIC_INVOICING_MODE_ACTIVATED",
	"ELASTIC_INVOICING_MODE_DEACTIVATED",
	"TERMINATE_PAID_SERVICES",
	"BILLING_EMAIL_ADDRESS_ADDED",
	"BILLING_EMAIL_ADDRESS_CHANGED",
	"BILLING_EMAIL_ADDRESS_REMOVED",
	"AWS_BILLING_ACCOUNT_CREDIT_ISSUED",
	"GCP_BILLING_ACCOUNT_CREDIT_ISSUED",
	"CREDIT_SFOLI_MODIFIED",
	"CREDIT_AMOUNT_MODIFIED",
	"PREPAID_PLAN_MODIFIED",
	"AWS_USAGE_REPORTED",
	"AZURE_USAGE_REPORTED",
	"GCP_USAGE_REPORTED",
	"BECAME_PAYING_ORG",
	"NEW_LINKED_ORG",
	"UNLINKED_ORG",
	"ORG_LINKED_TO_PAYING_ORG",
	"ORG_UNLINKED_FROM_PAYING_ORG",
	"PAYMENT_UPDATED_THROUGH_API",
}

func (v *BillingEventTypeForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingEventTypeForOrg(value)
	for _, existing := range AllowedBillingEventTypeForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingEventTypeForOrg", value)
}

// NewBillingEventTypeForOrgFromValue returns a pointer to a valid BillingEventTypeForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingEventTypeForOrgFromValue(v string) (*BillingEventTypeForOrg, error) {
	ev := BillingEventTypeForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingEventTypeForOrg: valid values are %v", v, AllowedBillingEventTypeForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingEventTypeForOrg) IsValid() bool {
	for _, existing := range AllowedBillingEventTypeForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingEventTypeForOrg value
func (v BillingEventTypeForOrg) Ptr() *BillingEventTypeForOrg {
	return &v
}

type NullableBillingEventTypeForOrg struct {
	value *BillingEventTypeForOrg
	isSet bool
}

func (v NullableBillingEventTypeForOrg) Get() *BillingEventTypeForOrg {
	return v.value
}

func (v *NullableBillingEventTypeForOrg) Set(val *BillingEventTypeForOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingEventTypeForOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingEventTypeForOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingEventTypeForOrg(val *BillingEventTypeForOrg) *NullableBillingEventTypeForOrg {
	return &NullableBillingEventTypeForOrg{value: val, isSet: true}
}

func (v NullableBillingEventTypeForOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingEventTypeForOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

