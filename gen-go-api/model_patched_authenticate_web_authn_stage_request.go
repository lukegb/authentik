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

// checks if the PatchedAuthenticateWebAuthnStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedAuthenticateWebAuthnStageRequest{}

// PatchedAuthenticateWebAuthnStageRequest AuthenticateWebAuthnStage Serializer
type PatchedAuthenticateWebAuthnStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName NullableString `json:"friendly_name,omitempty"`
	UserVerification *UserVerificationEnum `json:"user_verification,omitempty"`
	AuthenticatorAttachment NullableAuthenticatorAttachmentEnum `json:"authenticator_attachment,omitempty"`
	ResidentKeyRequirement *ResidentKeyRequirementEnum `json:"resident_key_requirement,omitempty"`
}

// NewPatchedAuthenticateWebAuthnStageRequest instantiates a new PatchedAuthenticateWebAuthnStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticateWebAuthnStageRequest() *PatchedAuthenticateWebAuthnStageRequest {
	this := PatchedAuthenticateWebAuthnStageRequest{}
	return &this
}

// NewPatchedAuthenticateWebAuthnStageRequestWithDefaults instantiates a new PatchedAuthenticateWebAuthnStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticateWebAuthnStageRequestWithDefaults() *PatchedAuthenticateWebAuthnStageRequest {
	this := PatchedAuthenticateWebAuthnStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlow() string {
	if o == nil || IsNil(o.ConfigureFlow.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}
// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName.Get()) {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticateWebAuthnStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerification() UserVerificationEnum {
	if o == nil || IsNil(o.UserVerification) {
		var ret UserVerificationEnum
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given UserVerificationEnum and assigns it to the UserVerification field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetUserVerification(v UserVerificationEnum) {
	o.UserVerification = &v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticateWebAuthnStageRequest) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum {
	if o == nil || IsNil(o.AuthenticatorAttachment.Get()) {
		var ret AuthenticatorAttachmentEnum
		return ret
	}
	return *o.AuthenticatorAttachment.Get()
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticateWebAuthnStageRequest) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticatorAttachment.Get(), o.AuthenticatorAttachment.IsSet()
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasAuthenticatorAttachment() bool {
	if o != nil && o.AuthenticatorAttachment.IsSet() {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given NullableAuthenticatorAttachmentEnum and assigns it to the AuthenticatorAttachment field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum) {
	o.AuthenticatorAttachment.Set(&v)
}
// SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) SetAuthenticatorAttachmentNil() {
	o.AuthenticatorAttachment.Set(nil)
}

// UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetAuthenticatorAttachment() {
	o.AuthenticatorAttachment.Unset()
}

// GetResidentKeyRequirement returns the ResidentKeyRequirement field value if set, zero value otherwise.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetResidentKeyRequirement() ResidentKeyRequirementEnum {
	if o == nil || IsNil(o.ResidentKeyRequirement) {
		var ret ResidentKeyRequirementEnum
		return ret
	}
	return *o.ResidentKeyRequirement
}

// GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool) {
	if o == nil || IsNil(o.ResidentKeyRequirement) {
		return nil, false
	}
	return o.ResidentKeyRequirement, true
}

// HasResidentKeyRequirement returns a boolean if a field has been set.
func (o *PatchedAuthenticateWebAuthnStageRequest) HasResidentKeyRequirement() bool {
	if o != nil && !IsNil(o.ResidentKeyRequirement) {
		return true
	}

	return false
}

// SetResidentKeyRequirement gets a reference to the given ResidentKeyRequirementEnum and assigns it to the ResidentKeyRequirement field.
func (o *PatchedAuthenticateWebAuthnStageRequest) SetResidentKeyRequirement(v ResidentKeyRequirementEnum) {
	o.ResidentKeyRequirement = &v
}

func (o PatchedAuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedAuthenticateWebAuthnStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if !IsNil(o.UserVerification) {
		toSerialize["user_verification"] = o.UserVerification
	}
	if o.AuthenticatorAttachment.IsSet() {
		toSerialize["authenticator_attachment"] = o.AuthenticatorAttachment.Get()
	}
	if !IsNil(o.ResidentKeyRequirement) {
		toSerialize["resident_key_requirement"] = o.ResidentKeyRequirement
	}
	return toSerialize, nil
}

type NullablePatchedAuthenticateWebAuthnStageRequest struct {
	value *PatchedAuthenticateWebAuthnStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) Get() *PatchedAuthenticateWebAuthnStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) Set(val *PatchedAuthenticateWebAuthnStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticateWebAuthnStageRequest(val *PatchedAuthenticateWebAuthnStageRequest) *NullablePatchedAuthenticateWebAuthnStageRequest {
	return &NullablePatchedAuthenticateWebAuthnStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticateWebAuthnStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


