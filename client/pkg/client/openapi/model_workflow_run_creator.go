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

// WorkflowRunCreator - struct for WorkflowRunCreator
type WorkflowRunCreator struct {
	EventWorkflowRunCreator *EventWorkflowRunCreator
	UserWorkflowRunCreator *UserWorkflowRunCreator
}

// EventWorkflowRunCreatorAsWorkflowRunCreator is a convenience function that returns EventWorkflowRunCreator wrapped in WorkflowRunCreator
func EventWorkflowRunCreatorAsWorkflowRunCreator(v *EventWorkflowRunCreator) WorkflowRunCreator {
	return WorkflowRunCreator{ EventWorkflowRunCreator: v}
}

// UserWorkflowRunCreatorAsWorkflowRunCreator is a convenience function that returns UserWorkflowRunCreator wrapped in WorkflowRunCreator
func UserWorkflowRunCreatorAsWorkflowRunCreator(v *UserWorkflowRunCreator) WorkflowRunCreator {
	return WorkflowRunCreator{ UserWorkflowRunCreator: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowRunCreator) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventWorkflowRunCreator
	err = json.Unmarshal(data, &dst.EventWorkflowRunCreator)
	if err == nil {
		jsonEventWorkflowRunCreator, _ := json.Marshal(dst.EventWorkflowRunCreator)
		if string(jsonEventWorkflowRunCreator) == "{}" { // empty struct
			dst.EventWorkflowRunCreator = nil
		} else {
			match++
		}
	} else {
		dst.EventWorkflowRunCreator = nil
	}

	// try to unmarshal data into UserWorkflowRunCreator
	err = json.Unmarshal(data, &dst.UserWorkflowRunCreator)
	if err == nil {
		jsonUserWorkflowRunCreator, _ := json.Marshal(dst.UserWorkflowRunCreator)
		if string(jsonUserWorkflowRunCreator) == "{}" { // empty struct
			dst.UserWorkflowRunCreator = nil
		} else {
			match++
		}
	} else {
		dst.UserWorkflowRunCreator = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EventWorkflowRunCreator = nil
		dst.UserWorkflowRunCreator = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WorkflowRunCreator)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WorkflowRunCreator)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowRunCreator) MarshalJSON() ([]byte, error) {
	if src.EventWorkflowRunCreator != nil {
		return json.Marshal(&src.EventWorkflowRunCreator)
	}

	if src.UserWorkflowRunCreator != nil {
		return json.Marshal(&src.UserWorkflowRunCreator)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowRunCreator) GetActualInstance() (interface{}) {
	if obj.EventWorkflowRunCreator != nil {
		return obj.EventWorkflowRunCreator
	}

	if obj.UserWorkflowRunCreator != nil {
		return obj.UserWorkflowRunCreator
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowRunCreator struct {
	value *WorkflowRunCreator
	isSet bool
}

func (v NullableWorkflowRunCreator) Get() *WorkflowRunCreator {
	return v.value
}

func (v *NullableWorkflowRunCreator) Set(val *WorkflowRunCreator) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunCreator) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunCreator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunCreator(val *WorkflowRunCreator) *NullableWorkflowRunCreator {
	return &NullableWorkflowRunCreator{value: val, isSet: true}
}

func (v NullableWorkflowRunCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


