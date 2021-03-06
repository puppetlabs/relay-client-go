/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TokenWithSecret - A newly created token with secret
type TokenWithSecret struct {
	UserTokenWithSecret *UserTokenWithSecret
}

// UserTokenWithSecretAsTokenWithSecret is a convenience function that returns UserTokenWithSecret wrapped in TokenWithSecret
func UserTokenWithSecretAsTokenWithSecret(v *UserTokenWithSecret) TokenWithSecret {
	return TokenWithSecret{
		UserTokenWithSecret: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TokenWithSecret) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'UserTokenWithSecret'
	if jsonDict["type"] == "UserTokenWithSecret" {
		// try to unmarshal JSON data into UserTokenWithSecret
		err = json.Unmarshal(data, &dst.UserTokenWithSecret)
		if err == nil {
			return nil // data stored in dst.UserTokenWithSecret, return on the first match
		} else {
			dst.UserTokenWithSecret = nil
			return fmt.Errorf("Failed to unmarshal TokenWithSecret as UserTokenWithSecret: %s", err.Error())
		}
	}

	// check if the discriminator value is 'user'
	if jsonDict["type"] == "user" {
		// try to unmarshal JSON data into UserTokenWithSecret
		err = json.Unmarshal(data, &dst.UserTokenWithSecret)
		if err == nil {
			return nil // data stored in dst.UserTokenWithSecret, return on the first match
		} else {
			dst.UserTokenWithSecret = nil
			return fmt.Errorf("Failed to unmarshal TokenWithSecret as UserTokenWithSecret: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TokenWithSecret) MarshalJSON() ([]byte, error) {
	if src.UserTokenWithSecret != nil {
		return json.Marshal(&src.UserTokenWithSecret)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TokenWithSecret) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserTokenWithSecret != nil {
		return obj.UserTokenWithSecret
	}

	// all schemas are nil
	return nil
}

type NullableTokenWithSecret struct {
	value *TokenWithSecret
	isSet bool
}

func (v NullableTokenWithSecret) Get() *TokenWithSecret {
	return v.value
}

func (v *NullableTokenWithSecret) Set(val *TokenWithSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenWithSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenWithSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenWithSecret(val *TokenWithSecret) *NullableTokenWithSecret {
	return &NullableTokenWithSecret{value: val, isSet: true}
}

func (v NullableTokenWithSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenWithSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
