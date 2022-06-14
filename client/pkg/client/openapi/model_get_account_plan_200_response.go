/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetAccountPlan200Response struct for GetAccountPlan200Response
type GetAccountPlan200Response struct {
	Access  *EntityAccess  `json:"access,omitempty"`
	Account *Account       `json:"account,omitempty"`
	Quotas  *AccountQuotas `json:"quotas,omitempty"`
}

// NewGetAccountPlan200Response instantiates a new GetAccountPlan200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountPlan200Response() *GetAccountPlan200Response {
	this := GetAccountPlan200Response{}
	return &this
}

// NewGetAccountPlan200ResponseWithDefaults instantiates a new GetAccountPlan200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountPlan200ResponseWithDefaults() *GetAccountPlan200Response {
	this := GetAccountPlan200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetAccountPlan200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPlan200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetAccountPlan200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetAccountPlan200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GetAccountPlan200Response) GetAccount() Account {
	if o == nil || o.Account == nil {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPlan200Response) GetAccountOk() (*Account, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *GetAccountPlan200Response) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *GetAccountPlan200Response) SetAccount(v Account) {
	o.Account = &v
}

// GetQuotas returns the Quotas field value if set, zero value otherwise.
func (o *GetAccountPlan200Response) GetQuotas() AccountQuotas {
	if o == nil || o.Quotas == nil {
		var ret AccountQuotas
		return ret
	}
	return *o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPlan200Response) GetQuotasOk() (*AccountQuotas, bool) {
	if o == nil || o.Quotas == nil {
		return nil, false
	}
	return o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *GetAccountPlan200Response) HasQuotas() bool {
	if o != nil && o.Quotas != nil {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given AccountQuotas and assigns it to the Quotas field.
func (o *GetAccountPlan200Response) SetQuotas(v AccountQuotas) {
	o.Quotas = &v
}

func (o GetAccountPlan200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Quotas != nil {
		toSerialize["quotas"] = o.Quotas
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountPlan200Response struct {
	value *GetAccountPlan200Response
	isSet bool
}

func (v NullableGetAccountPlan200Response) Get() *GetAccountPlan200Response {
	return v.value
}

func (v *NullableGetAccountPlan200Response) Set(val *GetAccountPlan200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountPlan200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountPlan200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountPlan200Response(val *GetAccountPlan200Response) *NullableGetAccountPlan200Response {
	return &NullableGetAccountPlan200Response{value: val, isSet: true}
}

func (v NullableGetAccountPlan200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountPlan200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}