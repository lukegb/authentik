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

// checks if the TenantAdminGroupRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAdminGroupRequestRequest{}

// TenantAdminGroupRequestRequest Tenant admin group creation request serializer
type TenantAdminGroupRequestRequest struct {
	User string `json:"user"`
}

// NewTenantAdminGroupRequestRequest instantiates a new TenantAdminGroupRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAdminGroupRequestRequest(user string) *TenantAdminGroupRequestRequest {
	this := TenantAdminGroupRequestRequest{}
	this.User = user
	return &this
}

// NewTenantAdminGroupRequestRequestWithDefaults instantiates a new TenantAdminGroupRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAdminGroupRequestRequestWithDefaults() *TenantAdminGroupRequestRequest {
	this := TenantAdminGroupRequestRequest{}
	return &this
}

// GetUser returns the User field value
func (o *TenantAdminGroupRequestRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TenantAdminGroupRequestRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TenantAdminGroupRequestRequest) SetUser(v string) {
	o.User = v
}

func (o TenantAdminGroupRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAdminGroupRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableTenantAdminGroupRequestRequest struct {
	value *TenantAdminGroupRequestRequest
	isSet bool
}

func (v NullableTenantAdminGroupRequestRequest) Get() *TenantAdminGroupRequestRequest {
	return v.value
}

func (v *NullableTenantAdminGroupRequestRequest) Set(val *TenantAdminGroupRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAdminGroupRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAdminGroupRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAdminGroupRequestRequest(val *TenantAdminGroupRequestRequest) *NullableTenantAdminGroupRequestRequest {
	return &NullableTenantAdminGroupRequestRequest{value: val, isSet: true}
}

func (v NullableTenantAdminGroupRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAdminGroupRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

