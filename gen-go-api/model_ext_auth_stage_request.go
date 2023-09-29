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

// checks if the ExtAuthStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtAuthStageRequest{}

// ExtAuthStageRequest ExtAuthStage Serializer
type ExtAuthStageRequest struct {
	Name string `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
}

// NewExtAuthStageRequest instantiates a new ExtAuthStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtAuthStageRequest(name string) *ExtAuthStageRequest {
	this := ExtAuthStageRequest{}
	this.Name = name
	return &this
}

// NewExtAuthStageRequestWithDefaults instantiates a new ExtAuthStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtAuthStageRequestWithDefaults() *ExtAuthStageRequest {
	this := ExtAuthStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ExtAuthStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtAuthStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtAuthStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *ExtAuthStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAuthStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *ExtAuthStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *ExtAuthStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

func (o ExtAuthStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtAuthStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	return toSerialize, nil
}

type NullableExtAuthStageRequest struct {
	value *ExtAuthStageRequest
	isSet bool
}

func (v NullableExtAuthStageRequest) Get() *ExtAuthStageRequest {
	return v.value
}

func (v *NullableExtAuthStageRequest) Set(val *ExtAuthStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExtAuthStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExtAuthStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtAuthStageRequest(val *ExtAuthStageRequest) *NullableExtAuthStageRequest {
	return &NullableExtAuthStageRequest{value: val, isSet: true}
}

func (v NullableExtAuthStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtAuthStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


