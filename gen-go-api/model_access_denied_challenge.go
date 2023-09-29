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

// checks if the AccessDeniedChallenge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDeniedChallenge{}

// AccessDeniedChallenge Challenge when a flow's active stage calls `stage_invalid()`.
type AccessDeniedChallenge struct {
	Type ChallengeChoices `json:"type"`
	FlowInfo *ContextualFlowInfo `json:"flow_info,omitempty"`
	Component *string `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser string `json:"pending_user"`
	PendingUserAvatar string `json:"pending_user_avatar"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewAccessDeniedChallenge instantiates a new AccessDeniedChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDeniedChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string) *AccessDeniedChallenge {
	this := AccessDeniedChallenge{}
	this.Type = type_
	var component string = "ak-stage-access-denied"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	return &this
}

// NewAccessDeniedChallengeWithDefaults instantiates a new AccessDeniedChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDeniedChallengeWithDefaults() *AccessDeniedChallenge {
	this := AccessDeniedChallenge{}
	var component string = "ak-stage-access-denied"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *AccessDeniedChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessDeniedChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AccessDeniedChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || IsNil(o.FlowInfo) {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || IsNil(o.FlowInfo) {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AccessDeniedChallenge) HasFlowInfo() bool {
	if o != nil && !IsNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AccessDeniedChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AccessDeniedChallenge) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AccessDeniedChallenge) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AccessDeniedChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AccessDeniedChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || IsNil(o.ResponseErrors) {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || IsNil(o.ResponseErrors) {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AccessDeniedChallenge) HasResponseErrors() bool {
	if o != nil && !IsNil(o.ResponseErrors) {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AccessDeniedChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AccessDeniedChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AccessDeniedChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AccessDeniedChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AccessDeniedChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AccessDeniedChallenge) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedChallenge) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AccessDeniedChallenge) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AccessDeniedChallenge) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o AccessDeniedChallenge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDeniedChallenge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.FlowInfo) {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.ResponseErrors) {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	toSerialize["pending_user"] = o.PendingUser
	toSerialize["pending_user_avatar"] = o.PendingUserAvatar
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableAccessDeniedChallenge struct {
	value *AccessDeniedChallenge
	isSet bool
}

func (v NullableAccessDeniedChallenge) Get() *AccessDeniedChallenge {
	return v.value
}

func (v *NullableAccessDeniedChallenge) Set(val *AccessDeniedChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDeniedChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDeniedChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDeniedChallenge(val *AccessDeniedChallenge) *NullableAccessDeniedChallenge {
	return &NullableAccessDeniedChallenge{value: val, isSet: true}
}

func (v NullableAccessDeniedChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDeniedChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


