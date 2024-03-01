/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FlowImportResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowImportResult{}

// FlowImportResult Logs of an attempted flow import
type FlowImportResult struct {
	Logs []map[string]interface{} `json:"logs"`
	Success bool `json:"success"`
}

// NewFlowImportResult instantiates a new FlowImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowImportResult(logs []map[string]interface{}, success bool) *FlowImportResult {
	this := FlowImportResult{}
	this.Logs = logs
	this.Success = success
	return &this
}

// NewFlowImportResultWithDefaults instantiates a new FlowImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowImportResultWithDefaults() *FlowImportResult {
	this := FlowImportResult{}
	return &this
}

// GetLogs returns the Logs field value
func (o *FlowImportResult) GetLogs() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *FlowImportResult) GetLogsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *FlowImportResult) SetLogs(v []map[string]interface{}) {
	o.Logs = v
}

// GetSuccess returns the Success field value
func (o *FlowImportResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *FlowImportResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *FlowImportResult) SetSuccess(v bool) {
	o.Success = v
}

func (o FlowImportResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowImportResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: logs is readOnly
	// skip: success is readOnly
	return toSerialize, nil
}

type NullableFlowImportResult struct {
	value *FlowImportResult
	isSet bool
}

func (v NullableFlowImportResult) Get() *FlowImportResult {
	return v.value
}

func (v *NullableFlowImportResult) Set(val *FlowImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowImportResult(val *FlowImportResult) *NullableFlowImportResult {
	return &NullableFlowImportResult{value: val, isSet: true}
}

func (v NullableFlowImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

