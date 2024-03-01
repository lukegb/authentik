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

// checks if the SMSDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMSDeviceRequest{}

// SMSDeviceRequest Serializer for sms authenticator devices
type SMSDeviceRequest struct {
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewSMSDeviceRequest instantiates a new SMSDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSDeviceRequest(name string) *SMSDeviceRequest {
	this := SMSDeviceRequest{}
	this.Name = name
	return &this
}

// NewSMSDeviceRequestWithDefaults instantiates a new SMSDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSDeviceRequestWithDefaults() *SMSDeviceRequest {
	this := SMSDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SMSDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SMSDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SMSDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o SMSDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMSDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSMSDeviceRequest struct {
	value *SMSDeviceRequest
	isSet bool
}

func (v NullableSMSDeviceRequest) Get() *SMSDeviceRequest {
	return v.value
}

func (v *NullableSMSDeviceRequest) Set(val *SMSDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSDeviceRequest(val *SMSDeviceRequest) *NullableSMSDeviceRequest {
	return &NullableSMSDeviceRequest{value: val, isSet: true}
}

func (v NullableSMSDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

