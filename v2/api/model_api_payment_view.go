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

// ApiPaymentView Funds transferred to MongoDB to cover the specified service in this invoice.
type ApiPaymentView struct {
	// Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar) and calculates its value as **subtotalCents** + **salesTaxCents** - **startingBalanceCents**.
	AmountBilledCents *int64 `json:"amountBilledCents,omitempty"`
	// Sum that the specified organization paid toward the associated invoice. This parameter expresses its value in cents (100ths of one US Dollar).
	AmountPaidCents *int64 `json:"amountPaidCents,omitempty"`
	// Date and time when the customer made this payment attempt. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this payment toward the associated invoice.
	Id *string `json:"id,omitempty"`
	// Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar).
	SalesTaxCents *int64 `json:"salesTaxCents,omitempty"`
	// Phase of payment processing for the associated invoice when you made this request.  These phases include:  | Phase Value | Reason | |---|---| | `CANCELLED` | Customer or MongoDB cancelled the payment. | | `ERROR` | Issue arose when attempting to complete payment. | | `FAILED` | MongoDB tried to charge the credit card without success. | | `FAILED_AUTHENTICATION` | Strong Customer Authentication has failed. Confirm that your payment method is authenticated. | | `FORGIVEN` | Customer initiated payment which MongoDB later forgave. | | `INVOICED` | MongoDB issued an invoice that included this line item. | | `NEW` | Customer provided a method of payment, but MongoDB hasn't tried to charge the credit card. | | `PAID` | Customer submitted a successful payment. | | `PARTIAL_PAID` | Customer paid for part of this line item. | 
	StatusName *string `json:"statusName,omitempty"`
	// Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar).
	SubtotalCents *int64 `json:"subtotalCents,omitempty"`
	// Date and time when the customer made an update to this payment attempt. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewApiPaymentView instantiates a new ApiPaymentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPaymentView() *ApiPaymentView {
	this := ApiPaymentView{}
	return &this
}

// NewApiPaymentViewWithDefaults instantiates a new ApiPaymentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPaymentViewWithDefaults() *ApiPaymentView {
	this := ApiPaymentView{}
	return &this
}

// GetAmountBilledCents returns the AmountBilledCents field value if set, zero value otherwise.
func (o *ApiPaymentView) GetAmountBilledCents() int64 {
	if o == nil || o.AmountBilledCents == nil {
		var ret int64
		return ret
	}
	return *o.AmountBilledCents
}

// GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetAmountBilledCentsOk() (*int64, bool) {
	if o == nil || o.AmountBilledCents == nil {
		return nil, false
	}
	return o.AmountBilledCents, true
}

// HasAmountBilledCents returns a boolean if a field has been set.
func (o *ApiPaymentView) HasAmountBilledCents() bool {
	if o != nil && o.AmountBilledCents != nil {
		return true
	}

	return false
}

// SetAmountBilledCents gets a reference to the given int64 and assigns it to the AmountBilledCents field.
func (o *ApiPaymentView) SetAmountBilledCents(v int64) {
	o.AmountBilledCents = &v
}

// GetAmountPaidCents returns the AmountPaidCents field value if set, zero value otherwise.
func (o *ApiPaymentView) GetAmountPaidCents() int64 {
	if o == nil || o.AmountPaidCents == nil {
		var ret int64
		return ret
	}
	return *o.AmountPaidCents
}

// GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetAmountPaidCentsOk() (*int64, bool) {
	if o == nil || o.AmountPaidCents == nil {
		return nil, false
	}
	return o.AmountPaidCents, true
}

// HasAmountPaidCents returns a boolean if a field has been set.
func (o *ApiPaymentView) HasAmountPaidCents() bool {
	if o != nil && o.AmountPaidCents != nil {
		return true
	}

	return false
}

// SetAmountPaidCents gets a reference to the given int64 and assigns it to the AmountPaidCents field.
func (o *ApiPaymentView) SetAmountPaidCents(v int64) {
	o.AmountPaidCents = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiPaymentView) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiPaymentView) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiPaymentView) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiPaymentView) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiPaymentView) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiPaymentView) SetId(v string) {
	o.Id = &v
}

// GetSalesTaxCents returns the SalesTaxCents field value if set, zero value otherwise.
func (o *ApiPaymentView) GetSalesTaxCents() int64 {
	if o == nil || o.SalesTaxCents == nil {
		var ret int64
		return ret
	}
	return *o.SalesTaxCents
}

// GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetSalesTaxCentsOk() (*int64, bool) {
	if o == nil || o.SalesTaxCents == nil {
		return nil, false
	}
	return o.SalesTaxCents, true
}

// HasSalesTaxCents returns a boolean if a field has been set.
func (o *ApiPaymentView) HasSalesTaxCents() bool {
	if o != nil && o.SalesTaxCents != nil {
		return true
	}

	return false
}

// SetSalesTaxCents gets a reference to the given int64 and assigns it to the SalesTaxCents field.
func (o *ApiPaymentView) SetSalesTaxCents(v int64) {
	o.SalesTaxCents = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *ApiPaymentView) GetStatusName() string {
	if o == nil || o.StatusName == nil {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetStatusNameOk() (*string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *ApiPaymentView) HasStatusName() bool {
	if o != nil && o.StatusName != nil {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *ApiPaymentView) SetStatusName(v string) {
	o.StatusName = &v
}

// GetSubtotalCents returns the SubtotalCents field value if set, zero value otherwise.
func (o *ApiPaymentView) GetSubtotalCents() int64 {
	if o == nil || o.SubtotalCents == nil {
		var ret int64
		return ret
	}
	return *o.SubtotalCents
}

// GetSubtotalCentsOk returns a tuple with the SubtotalCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetSubtotalCentsOk() (*int64, bool) {
	if o == nil || o.SubtotalCents == nil {
		return nil, false
	}
	return o.SubtotalCents, true
}

// HasSubtotalCents returns a boolean if a field has been set.
func (o *ApiPaymentView) HasSubtotalCents() bool {
	if o != nil && o.SubtotalCents != nil {
		return true
	}

	return false
}

// SetSubtotalCents gets a reference to the given int64 and assigns it to the SubtotalCents field.
func (o *ApiPaymentView) SetSubtotalCents(v int64) {
	o.SubtotalCents = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ApiPaymentView) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPaymentView) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ApiPaymentView) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ApiPaymentView) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o ApiPaymentView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountBilledCents != nil {
		toSerialize["amountBilledCents"] = o.AmountBilledCents
	}
	if o.AmountPaidCents != nil {
		toSerialize["amountPaidCents"] = o.AmountPaidCents
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SalesTaxCents != nil {
		toSerialize["salesTaxCents"] = o.SalesTaxCents
	}
	if o.StatusName != nil {
		toSerialize["statusName"] = o.StatusName
	}
	if o.SubtotalCents != nil {
		toSerialize["subtotalCents"] = o.SubtotalCents
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableApiPaymentView struct {
	value *ApiPaymentView
	isSet bool
}

func (v NullableApiPaymentView) Get() *ApiPaymentView {
	return v.value
}

func (v *NullableApiPaymentView) Set(val *ApiPaymentView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPaymentView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPaymentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPaymentView(val *ApiPaymentView) *NullableApiPaymentView {
	return &NullableApiPaymentView{value: val, isSet: true}
}

func (v NullableApiPaymentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPaymentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


