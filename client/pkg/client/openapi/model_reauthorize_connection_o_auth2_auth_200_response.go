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

// ReauthorizeConnectionOAuth2Auth200Response The response type for reauthorization an OAuth 2.0 connection
type ReauthorizeConnectionOAuth2Auth200Response struct {
	// The URL to forward the user to
	Url string `json:"url"`
}

// NewReauthorizeConnectionOAuth2Auth200Response instantiates a new ReauthorizeConnectionOAuth2Auth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReauthorizeConnectionOAuth2Auth200Response(url string) *ReauthorizeConnectionOAuth2Auth200Response {
	this := ReauthorizeConnectionOAuth2Auth200Response{}
	this.Url = url
	return &this
}

// NewReauthorizeConnectionOAuth2Auth200ResponseWithDefaults instantiates a new ReauthorizeConnectionOAuth2Auth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReauthorizeConnectionOAuth2Auth200ResponseWithDefaults() *ReauthorizeConnectionOAuth2Auth200Response {
	this := ReauthorizeConnectionOAuth2Auth200Response{}
	return &this
}

// GetUrl returns the Url field value
func (o *ReauthorizeConnectionOAuth2Auth200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ReauthorizeConnectionOAuth2Auth200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ReauthorizeConnectionOAuth2Auth200Response) SetUrl(v string) {
	o.Url = v
}

func (o ReauthorizeConnectionOAuth2Auth200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableReauthorizeConnectionOAuth2Auth200Response struct {
	value *ReauthorizeConnectionOAuth2Auth200Response
	isSet bool
}

func (v NullableReauthorizeConnectionOAuth2Auth200Response) Get() *ReauthorizeConnectionOAuth2Auth200Response {
	return v.value
}

func (v *NullableReauthorizeConnectionOAuth2Auth200Response) Set(val *ReauthorizeConnectionOAuth2Auth200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReauthorizeConnectionOAuth2Auth200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReauthorizeConnectionOAuth2Auth200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReauthorizeConnectionOAuth2Auth200Response(val *ReauthorizeConnectionOAuth2Auth200Response) *NullableReauthorizeConnectionOAuth2Auth200Response {
	return &NullableReauthorizeConnectionOAuth2Auth200Response{value: val, isSet: true}
}

func (v NullableReauthorizeConnectionOAuth2Auth200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReauthorizeConnectionOAuth2Auth200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
