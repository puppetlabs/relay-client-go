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
	"fmt"
)

// RevisionSource - The source of a revision
type RevisionSource struct {
	RevisionRelaySource *RevisionRelaySource
	RevisionRepositorySource *RevisionRepositorySource
}

// RevisionRelaySourceAsRevisionSource is a convenience function that returns RevisionRelaySource wrapped in RevisionSource
func RevisionRelaySourceAsRevisionSource(v *RevisionRelaySource) RevisionSource {
	return RevisionSource{ RevisionRelaySource: v}
}

// RevisionRepositorySourceAsRevisionSource is a convenience function that returns RevisionRepositorySource wrapped in RevisionSource
func RevisionRepositorySourceAsRevisionSource(v *RevisionRepositorySource) RevisionSource {
	return RevisionSource{ RevisionRepositorySource: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RevisionSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RevisionRelaySource
	err = json.Unmarshal(data, &dst.RevisionRelaySource)
	if err == nil {
		jsonRevisionRelaySource, _ := json.Marshal(dst.RevisionRelaySource)
		if string(jsonRevisionRelaySource) == "{}" { // empty struct
			dst.RevisionRelaySource = nil
		} else {
			match++
		}
	} else {
		dst.RevisionRelaySource = nil
	}

	// try to unmarshal data into RevisionRepositorySource
	err = json.Unmarshal(data, &dst.RevisionRepositorySource)
	if err == nil {
		jsonRevisionRepositorySource, _ := json.Marshal(dst.RevisionRepositorySource)
		if string(jsonRevisionRepositorySource) == "{}" { // empty struct
			dst.RevisionRepositorySource = nil
		} else {
			match++
		}
	} else {
		dst.RevisionRepositorySource = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RevisionRelaySource = nil
		dst.RevisionRepositorySource = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(RevisionSource)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(RevisionSource)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RevisionSource) MarshalJSON() ([]byte, error) {
	if src.RevisionRelaySource != nil {
		return json.Marshal(&src.RevisionRelaySource)
	}

	if src.RevisionRepositorySource != nil {
		return json.Marshal(&src.RevisionRepositorySource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RevisionSource) GetActualInstance() (interface{}) {
	if obj.RevisionRelaySource != nil {
		return obj.RevisionRelaySource
	}

	if obj.RevisionRepositorySource != nil {
		return obj.RevisionRepositorySource
	}

	// all schemas are nil
	return nil
}

type NullableRevisionSource struct {
	value *RevisionSource
	isSet bool
}

func (v NullableRevisionSource) Get() *RevisionSource {
	return v.value
}

func (v *NullableRevisionSource) Set(val *RevisionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionSource(val *RevisionSource) *NullableRevisionSource {
	return &NullableRevisionSource{value: val, isSet: true}
}

func (v NullableRevisionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

