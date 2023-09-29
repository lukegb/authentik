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

// checks if the PatchedRACPropertyMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRACPropertyMappingRequest{}

// PatchedRACPropertyMappingRequest RACPropertyMapping Serializer
type PatchedRACPropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
	Name *string `json:"name,omitempty"`
	Expression *string `json:"expression,omitempty"`
	StaticSettings map[string]interface{} `json:"static_settings,omitempty"`
}

// NewPatchedRACPropertyMappingRequest instantiates a new PatchedRACPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRACPropertyMappingRequest() *PatchedRACPropertyMappingRequest {
	this := PatchedRACPropertyMappingRequest{}
	return &this
}

// NewPatchedRACPropertyMappingRequestWithDefaults instantiates a new PatchedRACPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRACPropertyMappingRequestWithDefaults() *PatchedRACPropertyMappingRequest {
	this := PatchedRACPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRACPropertyMappingRequest) GetManaged() string {
	if o == nil || IsNil(o.Managed.Get()) {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRACPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedRACPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedRACPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}
// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedRACPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedRACPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRACPropertyMappingRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRACPropertyMappingRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRACPropertyMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedRACPropertyMappingRequest) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedRACPropertyMappingRequest) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedRACPropertyMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetStaticSettings returns the StaticSettings field value if set, zero value otherwise.
func (o *PatchedRACPropertyMappingRequest) GetStaticSettings() map[string]interface{} {
	if o == nil || IsNil(o.StaticSettings) {
		var ret map[string]interface{}
		return ret
	}
	return o.StaticSettings
}

// GetStaticSettingsOk returns a tuple with the StaticSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACPropertyMappingRequest) GetStaticSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StaticSettings) {
		return map[string]interface{}{}, false
	}
	return o.StaticSettings, true
}

// HasStaticSettings returns a boolean if a field has been set.
func (o *PatchedRACPropertyMappingRequest) HasStaticSettings() bool {
	if o != nil && !IsNil(o.StaticSettings) {
		return true
	}

	return false
}

// SetStaticSettings gets a reference to the given map[string]interface{} and assigns it to the StaticSettings field.
func (o *PatchedRACPropertyMappingRequest) SetStaticSettings(v map[string]interface{}) {
	o.StaticSettings = v
}

func (o PatchedRACPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRACPropertyMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.StaticSettings) {
		toSerialize["static_settings"] = o.StaticSettings
	}
	return toSerialize, nil
}

type NullablePatchedRACPropertyMappingRequest struct {
	value *PatchedRACPropertyMappingRequest
	isSet bool
}

func (v NullablePatchedRACPropertyMappingRequest) Get() *PatchedRACPropertyMappingRequest {
	return v.value
}

func (v *NullablePatchedRACPropertyMappingRequest) Set(val *PatchedRACPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRACPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRACPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRACPropertyMappingRequest(val *PatchedRACPropertyMappingRequest) *NullablePatchedRACPropertyMappingRequest {
	return &NullablePatchedRACPropertyMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedRACPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRACPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


