/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserProfileSettings User profile settings
type UserProfileSettings struct {
	Digest *UserDigestSettings `json:"digest,omitempty"`
	Subscriptions *UserSubscriptionSettings `json:"subscriptions,omitempty"`
}

// NewUserProfileSettings instantiates a new UserProfileSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfileSettings() *UserProfileSettings {
	this := UserProfileSettings{}
	return &this
}

// NewUserProfileSettingsWithDefaults instantiates a new UserProfileSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileSettingsWithDefaults() *UserProfileSettings {
	this := UserProfileSettings{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *UserProfileSettings) GetDigest() UserDigestSettings {
	if o == nil || o.Digest == nil {
		var ret UserDigestSettings
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileSettings) GetDigestOk() (*UserDigestSettings, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *UserProfileSettings) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given UserDigestSettings and assigns it to the Digest field.
func (o *UserProfileSettings) SetDigest(v UserDigestSettings) {
	o.Digest = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *UserProfileSettings) GetSubscriptions() UserSubscriptionSettings {
	if o == nil || o.Subscriptions == nil {
		var ret UserSubscriptionSettings
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileSettings) GetSubscriptionsOk() (*UserSubscriptionSettings, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *UserProfileSettings) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given UserSubscriptionSettings and assigns it to the Subscriptions field.
func (o *UserProfileSettings) SetSubscriptions(v UserSubscriptionSettings) {
	o.Subscriptions = &v
}

func (o UserProfileSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableUserProfileSettings struct {
	value *UserProfileSettings
	isSet bool
}

func (v NullableUserProfileSettings) Get() *UserProfileSettings {
	return v.value
}

func (v *NullableUserProfileSettings) Set(val *UserProfileSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileSettings(val *UserProfileSettings) *NullableUserProfileSettings {
	return &NullableUserProfileSettings{value: val, isSet: true}
}

func (v NullableUserProfileSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

