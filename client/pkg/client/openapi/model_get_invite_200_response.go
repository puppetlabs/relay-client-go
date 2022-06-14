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

// GetInvite200Response struct for GetInvite200Response
type GetInvite200Response struct {
	Access *EntityAccess `json:"access,omitempty"`
	Invite *Invite       `json:"invite,omitempty"`
}

// NewGetInvite200Response instantiates a new GetInvite200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvite200Response() *GetInvite200Response {
	this := GetInvite200Response{}
	return &this
}

// NewGetInvite200ResponseWithDefaults instantiates a new GetInvite200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvite200ResponseWithDefaults() *GetInvite200Response {
	this := GetInvite200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetInvite200Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvite200Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetInvite200Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *GetInvite200Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetInvite returns the Invite field value if set, zero value otherwise.
func (o *GetInvite200Response) GetInvite() Invite {
	if o == nil || o.Invite == nil {
		var ret Invite
		return ret
	}
	return *o.Invite
}

// GetInviteOk returns a tuple with the Invite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvite200Response) GetInviteOk() (*Invite, bool) {
	if o == nil || o.Invite == nil {
		return nil, false
	}
	return o.Invite, true
}

// HasInvite returns a boolean if a field has been set.
func (o *GetInvite200Response) HasInvite() bool {
	if o != nil && o.Invite != nil {
		return true
	}

	return false
}

// SetInvite gets a reference to the given Invite and assigns it to the Invite field.
func (o *GetInvite200Response) SetInvite(v Invite) {
	o.Invite = &v
}

func (o GetInvite200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Invite != nil {
		toSerialize["invite"] = o.Invite
	}
	return json.Marshal(toSerialize)
}

type NullableGetInvite200Response struct {
	value *GetInvite200Response
	isSet bool
}

func (v NullableGetInvite200Response) Get() *GetInvite200Response {
	return v.value
}

func (v *NullableGetInvite200Response) Set(val *GetInvite200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvite200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvite200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvite200Response(val *GetInvite200Response) *NullableGetInvite200Response {
	return &NullableGetInvite200Response{value: val, isSet: true}
}

func (v NullableGetInvite200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvite200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}