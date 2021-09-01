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

// Token - An access token
type Token struct {
	UserToken *UserToken
}

// UserTokenAsToken is a convenience function that returns UserToken wrapped in Token
func UserTokenAsToken(v *UserToken) Token {
	return Token{ UserToken: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Token) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'UserToken'
	if jsonDict["type"] == "UserToken" {
		// try to unmarshal JSON data into UserToken
		err = json.Unmarshal(data, &dst.UserToken)
		if err == nil {
			return nil // data stored in dst.UserToken, return on the first match
		} else {
			dst.UserToken = nil
			return fmt.Errorf("Failed to unmarshal Token as UserToken: %s", err.Error())
		}
	}

	// check if the discriminator value is 'user'
	if jsonDict["type"] == "user" {
		// try to unmarshal JSON data into UserToken
		err = json.Unmarshal(data, &dst.UserToken)
		if err == nil {
			return nil // data stored in dst.UserToken, return on the first match
		} else {
			dst.UserToken = nil
			return fmt.Errorf("Failed to unmarshal Token as UserToken: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Token) MarshalJSON() ([]byte, error) {
	if src.UserToken != nil {
		return json.Marshal(&src.UserToken)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Token) GetActualInstance() (interface{}) {
	if obj.UserToken != nil {
		return obj.UserToken
	}

	// all schemas are nil
	return nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


