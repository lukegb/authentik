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

// checks if the PatchedUserSAMLSourceConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUserSAMLSourceConnectionRequest{}

// PatchedUserSAMLSourceConnectionRequest SAML Source Serializer
type PatchedUserSAMLSourceConnectionRequest struct {
	User *int32 `json:"user,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewPatchedUserSAMLSourceConnectionRequest instantiates a new PatchedUserSAMLSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserSAMLSourceConnectionRequest() *PatchedUserSAMLSourceConnectionRequest {
	this := PatchedUserSAMLSourceConnectionRequest{}
	return &this
}

// NewPatchedUserSAMLSourceConnectionRequestWithDefaults instantiates a new PatchedUserSAMLSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserSAMLSourceConnectionRequestWithDefaults() *PatchedUserSAMLSourceConnectionRequest {
	this := PatchedUserSAMLSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedUserSAMLSourceConnectionRequest) GetUser() int32 {
	if o == nil || IsNil(o.User) {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSAMLSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedUserSAMLSourceConnectionRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedUserSAMLSourceConnectionRequest) SetUser(v int32) {
	o.User = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedUserSAMLSourceConnectionRequest) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSAMLSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedUserSAMLSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedUserSAMLSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o PatchedUserSAMLSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUserSAMLSourceConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	return toSerialize, nil
}

type NullablePatchedUserSAMLSourceConnectionRequest struct {
	value *PatchedUserSAMLSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedUserSAMLSourceConnectionRequest) Get() *PatchedUserSAMLSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedUserSAMLSourceConnectionRequest) Set(val *PatchedUserSAMLSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserSAMLSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserSAMLSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserSAMLSourceConnectionRequest(val *PatchedUserSAMLSourceConnectionRequest) *NullablePatchedUserSAMLSourceConnectionRequest {
	return &NullablePatchedUserSAMLSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedUserSAMLSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserSAMLSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


