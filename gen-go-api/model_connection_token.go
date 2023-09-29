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

// checks if the ConnectionToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionToken{}

// ConnectionToken ConnectionToken Serializer
type ConnectionToken struct {
	Pk string `json:"pk"`
	Provider int32 `json:"provider"`
	ProviderObj RACProvider `json:"provider_obj"`
	Endpoint string `json:"endpoint"`
	EndpointObj Endpoint `json:"endpoint_obj"`
	User GroupMember `json:"user"`
}

// NewConnectionToken instantiates a new ConnectionToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionToken(pk string, provider int32, providerObj RACProvider, endpoint string, endpointObj Endpoint, user GroupMember) *ConnectionToken {
	this := ConnectionToken{}
	this.Pk = pk
	this.Provider = provider
	this.ProviderObj = providerObj
	this.Endpoint = endpoint
	this.EndpointObj = endpointObj
	this.User = user
	return &this
}

// NewConnectionTokenWithDefaults instantiates a new ConnectionToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionTokenWithDefaults() *ConnectionToken {
	this := ConnectionToken{}
	return &this
}

// GetPk returns the Pk field value
func (o *ConnectionToken) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ConnectionToken) SetPk(v string) {
	o.Pk = v
}

// GetProvider returns the Provider field value
func (o *ConnectionToken) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ConnectionToken) SetProvider(v int32) {
	o.Provider = v
}

// GetProviderObj returns the ProviderObj field value
func (o *ConnectionToken) GetProviderObj() RACProvider {
	if o == nil {
		var ret RACProvider
		return ret
	}

	return o.ProviderObj
}

// GetProviderObjOk returns a tuple with the ProviderObj field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetProviderObjOk() (*RACProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderObj, true
}

// SetProviderObj sets field value
func (o *ConnectionToken) SetProviderObj(v RACProvider) {
	o.ProviderObj = v
}

// GetEndpoint returns the Endpoint field value
func (o *ConnectionToken) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *ConnectionToken) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetEndpointObj returns the EndpointObj field value
func (o *ConnectionToken) GetEndpointObj() Endpoint {
	if o == nil {
		var ret Endpoint
		return ret
	}

	return o.EndpointObj
}

// GetEndpointObjOk returns a tuple with the EndpointObj field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetEndpointObjOk() (*Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointObj, true
}

// SetEndpointObj sets field value
func (o *ConnectionToken) SetEndpointObj(v Endpoint) {
	o.EndpointObj = v
}

// GetUser returns the User field value
func (o *ConnectionToken) GetUser() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ConnectionToken) GetUserOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ConnectionToken) SetUser(v GroupMember) {
	o.User = v
}

func (o ConnectionToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pk is readOnly
	toSerialize["provider"] = o.Provider
	// skip: provider_obj is readOnly
	// skip: endpoint is readOnly
	// skip: endpoint_obj is readOnly
	// skip: user is readOnly
	return toSerialize, nil
}

type NullableConnectionToken struct {
	value *ConnectionToken
	isSet bool
}

func (v NullableConnectionToken) Get() *ConnectionToken {
	return v.value
}

func (v *NullableConnectionToken) Set(val *ConnectionToken) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionToken) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionToken(val *ConnectionToken) *NullableConnectionToken {
	return &NullableConnectionToken{value: val, isSet: true}
}

func (v NullableConnectionToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


