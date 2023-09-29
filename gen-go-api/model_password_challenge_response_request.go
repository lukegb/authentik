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

// checks if the PasswordChallengeResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordChallengeResponseRequest{}

// PasswordChallengeResponseRequest Password challenge response
type PasswordChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Password string `json:"password"`
}

// NewPasswordChallengeResponseRequest instantiates a new PasswordChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordChallengeResponseRequest(password string) *PasswordChallengeResponseRequest {
	this := PasswordChallengeResponseRequest{}
	var component string = "ak-stage-password"
	this.Component = &component
	this.Password = password
	return &this
}

// NewPasswordChallengeResponseRequestWithDefaults instantiates a new PasswordChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordChallengeResponseRequestWithDefaults() *PasswordChallengeResponseRequest {
	this := PasswordChallengeResponseRequest{}
	var component string = "ak-stage-password"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PasswordChallengeResponseRequest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PasswordChallengeResponseRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *PasswordChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetPassword returns the Password field value
func (o *PasswordChallengeResponseRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PasswordChallengeResponseRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PasswordChallengeResponseRequest) SetPassword(v string) {
	o.Password = v
}

func (o PasswordChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordChallengeResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullablePasswordChallengeResponseRequest struct {
	value *PasswordChallengeResponseRequest
	isSet bool
}

func (v NullablePasswordChallengeResponseRequest) Get() *PasswordChallengeResponseRequest {
	return v.value
}

func (v *NullablePasswordChallengeResponseRequest) Set(val *PasswordChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordChallengeResponseRequest(val *PasswordChallengeResponseRequest) *NullablePasswordChallengeResponseRequest {
	return &NullablePasswordChallengeResponseRequest{value: val, isSet: true}
}

func (v NullablePasswordChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


