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

// checks if the UserSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSetting{}

// UserSetting Serializer for User settings for stages and sources
type UserSetting struct {
	ObjectUid string `json:"object_uid"`
	Component string `json:"component"`
	Title string `json:"title"`
	ConfigureUrl *string `json:"configure_url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

// NewUserSetting instantiates a new UserSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSetting(objectUid string, component string, title string) *UserSetting {
	this := UserSetting{}
	this.ObjectUid = objectUid
	this.Component = component
	this.Title = title
	return &this
}

// NewUserSettingWithDefaults instantiates a new UserSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingWithDefaults() *UserSetting {
	this := UserSetting{}
	return &this
}

// GetObjectUid returns the ObjectUid field value
func (o *UserSetting) GetObjectUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectUid
}

// GetObjectUidOk returns a tuple with the ObjectUid field value
// and a boolean to check if the value has been set.
func (o *UserSetting) GetObjectUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectUid, true
}

// SetObjectUid sets field value
func (o *UserSetting) SetObjectUid(v string) {
	o.ObjectUid = v
}

// GetComponent returns the Component field value
func (o *UserSetting) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *UserSetting) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *UserSetting) SetComponent(v string) {
	o.Component = v
}

// GetTitle returns the Title field value
func (o *UserSetting) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UserSetting) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UserSetting) SetTitle(v string) {
	o.Title = v
}

// GetConfigureUrl returns the ConfigureUrl field value if set, zero value otherwise.
func (o *UserSetting) GetConfigureUrl() string {
	if o == nil || IsNil(o.ConfigureUrl) {
		var ret string
		return ret
	}
	return *o.ConfigureUrl
}

// GetConfigureUrlOk returns a tuple with the ConfigureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSetting) GetConfigureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigureUrl) {
		return nil, false
	}
	return o.ConfigureUrl, true
}

// HasConfigureUrl returns a boolean if a field has been set.
func (o *UserSetting) HasConfigureUrl() bool {
	if o != nil && !IsNil(o.ConfigureUrl) {
		return true
	}

	return false
}

// SetConfigureUrl gets a reference to the given string and assigns it to the ConfigureUrl field.
func (o *UserSetting) SetConfigureUrl(v string) {
	o.ConfigureUrl = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *UserSetting) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSetting) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *UserSetting) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *UserSetting) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o UserSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_uid"] = o.ObjectUid
	toSerialize["component"] = o.Component
	toSerialize["title"] = o.Title
	if !IsNil(o.ConfigureUrl) {
		toSerialize["configure_url"] = o.ConfigureUrl
	}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	return toSerialize, nil
}

type NullableUserSetting struct {
	value *UserSetting
	isSet bool
}

func (v NullableUserSetting) Get() *UserSetting {
	return v.value
}

func (v *NullableUserSetting) Set(val *UserSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSetting(val *UserSetting) *NullableUserSetting {
	return &NullableUserSetting{value: val, isSet: true}
}

func (v NullableUserSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


