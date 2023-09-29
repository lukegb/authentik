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

// checks if the PermissionAssignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionAssignRequest{}

// PermissionAssignRequest Request to assign a new permission
type PermissionAssignRequest struct {
	Permissions []string `json:"permissions"`
	Model *ModelEnum `json:"model,omitempty"`
	ObjectPk *string `json:"object_pk,omitempty"`
}

// NewPermissionAssignRequest instantiates a new PermissionAssignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionAssignRequest(permissions []string) *PermissionAssignRequest {
	this := PermissionAssignRequest{}
	this.Permissions = permissions
	return &this
}

// NewPermissionAssignRequestWithDefaults instantiates a new PermissionAssignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionAssignRequestWithDefaults() *PermissionAssignRequest {
	this := PermissionAssignRequest{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *PermissionAssignRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionAssignRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionAssignRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *PermissionAssignRequest) GetModel() ModelEnum {
	if o == nil || IsNil(o.Model) {
		var ret ModelEnum
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionAssignRequest) GetModelOk() (*ModelEnum, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *PermissionAssignRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given ModelEnum and assigns it to the Model field.
func (o *PermissionAssignRequest) SetModel(v ModelEnum) {
	o.Model = &v
}

// GetObjectPk returns the ObjectPk field value if set, zero value otherwise.
func (o *PermissionAssignRequest) GetObjectPk() string {
	if o == nil || IsNil(o.ObjectPk) {
		var ret string
		return ret
	}
	return *o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionAssignRequest) GetObjectPkOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectPk) {
		return nil, false
	}
	return o.ObjectPk, true
}

// HasObjectPk returns a boolean if a field has been set.
func (o *PermissionAssignRequest) HasObjectPk() bool {
	if o != nil && !IsNil(o.ObjectPk) {
		return true
	}

	return false
}

// SetObjectPk gets a reference to the given string and assigns it to the ObjectPk field.
func (o *PermissionAssignRequest) SetObjectPk(v string) {
	o.ObjectPk = &v
}

func (o PermissionAssignRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionAssignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.ObjectPk) {
		toSerialize["object_pk"] = o.ObjectPk
	}
	return toSerialize, nil
}

type NullablePermissionAssignRequest struct {
	value *PermissionAssignRequest
	isSet bool
}

func (v NullablePermissionAssignRequest) Get() *PermissionAssignRequest {
	return v.value
}

func (v *NullablePermissionAssignRequest) Set(val *PermissionAssignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionAssignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionAssignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionAssignRequest(val *PermissionAssignRequest) *NullablePermissionAssignRequest {
	return &NullablePermissionAssignRequest{value: val, isSet: true}
}

func (v NullablePermissionAssignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionAssignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


