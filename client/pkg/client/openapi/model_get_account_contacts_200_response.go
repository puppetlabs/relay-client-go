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

// GetAccountContacts200Response struct for GetAccountContacts200Response
type GetAccountContacts200Response struct {
	// The list of contacts for this account
	Contacts []AccountContact `json:"contacts,omitempty"`
}

// NewGetAccountContacts200Response instantiates a new GetAccountContacts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountContacts200Response() *GetAccountContacts200Response {
	this := GetAccountContacts200Response{}
	return &this
}

// NewGetAccountContacts200ResponseWithDefaults instantiates a new GetAccountContacts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountContacts200ResponseWithDefaults() *GetAccountContacts200Response {
	this := GetAccountContacts200Response{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *GetAccountContacts200Response) GetContacts() []AccountContact {
	if o == nil || o.Contacts == nil {
		var ret []AccountContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountContacts200Response) GetContactsOk() ([]AccountContact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *GetAccountContacts200Response) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []AccountContact and assigns it to the Contacts field.
func (o *GetAccountContacts200Response) SetContacts(v []AccountContact) {
	o.Contacts = v
}

func (o GetAccountContacts200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountContacts200Response struct {
	value *GetAccountContacts200Response
	isSet bool
}

func (v NullableGetAccountContacts200Response) Get() *GetAccountContacts200Response {
	return v.value
}

func (v *NullableGetAccountContacts200Response) Set(val *GetAccountContacts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountContacts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountContacts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountContacts200Response(val *GetAccountContacts200Response) *NullableGetAccountContacts200Response {
	return &NullableGetAccountContacts200Response{value: val, isSet: true}
}

func (v NullableGetAccountContacts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountContacts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}