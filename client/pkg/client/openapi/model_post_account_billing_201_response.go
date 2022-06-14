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

// PostAccountBilling201Response struct for PostAccountBilling201Response
type PostAccountBilling201Response struct {
	// URL user should be redirect to in order to manage their billing plan
	RedirectUrl *string `json:"redirect_url,omitempty"`
}

// NewPostAccountBilling201Response instantiates a new PostAccountBilling201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccountBilling201Response() *PostAccountBilling201Response {
	this := PostAccountBilling201Response{}
	return &this
}

// NewPostAccountBilling201ResponseWithDefaults instantiates a new PostAccountBilling201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccountBilling201ResponseWithDefaults() *PostAccountBilling201Response {
	this := PostAccountBilling201Response{}
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *PostAccountBilling201Response) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccountBilling201Response) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PostAccountBilling201Response) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *PostAccountBilling201Response) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o PostAccountBilling201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePostAccountBilling201Response struct {
	value *PostAccountBilling201Response
	isSet bool
}

func (v NullablePostAccountBilling201Response) Get() *PostAccountBilling201Response {
	return v.value
}

func (v *NullablePostAccountBilling201Response) Set(val *PostAccountBilling201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccountBilling201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccountBilling201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccountBilling201Response(val *PostAccountBilling201Response) *NullablePostAccountBilling201Response {
	return &NullablePostAccountBilling201Response{value: val, isSet: true}
}

func (v NullablePostAccountBilling201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccountBilling201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}