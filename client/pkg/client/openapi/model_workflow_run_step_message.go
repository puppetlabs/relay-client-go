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

// WorkflowRunStepMessage - An error or notice produced by a step
type WorkflowRunStepMessage struct {
	LogWorkflowRunStepMessage            *LogWorkflowRunStepMessage
	SpecValidationWorkflowRunStepMessage *SpecValidationWorkflowRunStepMessage
	WhenEvaluationWorkflowRunStepMessage *WhenEvaluationWorkflowRunStepMessage
}

// LogWorkflowRunStepMessageAsWorkflowRunStepMessage is a convenience function that returns LogWorkflowRunStepMessage wrapped in WorkflowRunStepMessage
func LogWorkflowRunStepMessageAsWorkflowRunStepMessage(v *LogWorkflowRunStepMessage) WorkflowRunStepMessage {
	return WorkflowRunStepMessage{
		LogWorkflowRunStepMessage: v,
	}
}

// SpecValidationWorkflowRunStepMessageAsWorkflowRunStepMessage is a convenience function that returns SpecValidationWorkflowRunStepMessage wrapped in WorkflowRunStepMessage
func SpecValidationWorkflowRunStepMessageAsWorkflowRunStepMessage(v *SpecValidationWorkflowRunStepMessage) WorkflowRunStepMessage {
	return WorkflowRunStepMessage{
		SpecValidationWorkflowRunStepMessage: v,
	}
}

// WhenEvaluationWorkflowRunStepMessageAsWorkflowRunStepMessage is a convenience function that returns WhenEvaluationWorkflowRunStepMessage wrapped in WorkflowRunStepMessage
func WhenEvaluationWorkflowRunStepMessageAsWorkflowRunStepMessage(v *WhenEvaluationWorkflowRunStepMessage) WorkflowRunStepMessage {
	return WorkflowRunStepMessage{
		WhenEvaluationWorkflowRunStepMessage: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowRunStepMessage) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'LogWorkflowRunStepMessage'
	if jsonDict["type"] == "LogWorkflowRunStepMessage" {
		// try to unmarshal JSON data into LogWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.LogWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.LogWorkflowRunStepMessage, return on the first match
		} else {
			dst.LogWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as LogWorkflowRunStepMessage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SpecValidationWorkflowRunStepMessage'
	if jsonDict["type"] == "SpecValidationWorkflowRunStepMessage" {
		// try to unmarshal JSON data into SpecValidationWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.SpecValidationWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.SpecValidationWorkflowRunStepMessage, return on the first match
		} else {
			dst.SpecValidationWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as SpecValidationWorkflowRunStepMessage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WhenEvaluationWorkflowRunStepMessage'
	if jsonDict["type"] == "WhenEvaluationWorkflowRunStepMessage" {
		// try to unmarshal JSON data into WhenEvaluationWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.WhenEvaluationWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.WhenEvaluationWorkflowRunStepMessage, return on the first match
		} else {
			dst.WhenEvaluationWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as WhenEvaluationWorkflowRunStepMessage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'log'
	if jsonDict["type"] == "log" {
		// try to unmarshal JSON data into LogWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.LogWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.LogWorkflowRunStepMessage, return on the first match
		} else {
			dst.LogWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as LogWorkflowRunStepMessage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'spec-validation'
	if jsonDict["type"] == "spec-validation" {
		// try to unmarshal JSON data into SpecValidationWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.SpecValidationWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.SpecValidationWorkflowRunStepMessage, return on the first match
		} else {
			dst.SpecValidationWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as SpecValidationWorkflowRunStepMessage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'when-evaluation'
	if jsonDict["type"] == "when-evaluation" {
		// try to unmarshal JSON data into WhenEvaluationWorkflowRunStepMessage
		err = json.Unmarshal(data, &dst.WhenEvaluationWorkflowRunStepMessage)
		if err == nil {
			return nil // data stored in dst.WhenEvaluationWorkflowRunStepMessage, return on the first match
		} else {
			dst.WhenEvaluationWorkflowRunStepMessage = nil
			return fmt.Errorf("Failed to unmarshal WorkflowRunStepMessage as WhenEvaluationWorkflowRunStepMessage: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowRunStepMessage) MarshalJSON() ([]byte, error) {
	if src.LogWorkflowRunStepMessage != nil {
		return json.Marshal(&src.LogWorkflowRunStepMessage)
	}

	if src.SpecValidationWorkflowRunStepMessage != nil {
		return json.Marshal(&src.SpecValidationWorkflowRunStepMessage)
	}

	if src.WhenEvaluationWorkflowRunStepMessage != nil {
		return json.Marshal(&src.WhenEvaluationWorkflowRunStepMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowRunStepMessage) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LogWorkflowRunStepMessage != nil {
		return obj.LogWorkflowRunStepMessage
	}

	if obj.SpecValidationWorkflowRunStepMessage != nil {
		return obj.SpecValidationWorkflowRunStepMessage
	}

	if obj.WhenEvaluationWorkflowRunStepMessage != nil {
		return obj.WhenEvaluationWorkflowRunStepMessage
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowRunStepMessage struct {
	value *WorkflowRunStepMessage
	isSet bool
}

func (v NullableWorkflowRunStepMessage) Get() *WorkflowRunStepMessage {
	return v.value
}

func (v *NullableWorkflowRunStepMessage) Set(val *WorkflowRunStepMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunStepMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunStepMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunStepMessage(val *WorkflowRunStepMessage) *NullableWorkflowRunStepMessage {
	return &NullableWorkflowRunStepMessage{value: val, isSet: true}
}

func (v NullableWorkflowRunStepMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunStepMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
