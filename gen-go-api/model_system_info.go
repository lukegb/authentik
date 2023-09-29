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
	"time"
)

// checks if the SystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInfo{}

// SystemInfo Get system information.
type SystemInfo struct {
	// Get HTTP Request headers
	HttpHeaders map[string]string `json:"http_headers"`
	// Get HTTP host
	HttpHost string `json:"http_host"`
	// Get HTTP Secure flag
	HttpIsSecure bool `json:"http_is_secure"`
	Runtime SystemInfoRuntime `json:"runtime"`
	// Currently active brand
	Brand string `json:"brand"`
	// Current server time
	ServerTime time.Time `json:"server_time"`
	// Whether the embedded outpost is disabled
	EmbeddedOutpostDisabled bool `json:"embedded_outpost_disabled"`
	// Get the FQDN configured on the embedded outpost
	EmbeddedOutpostHost string `json:"embedded_outpost_host"`
}

// NewSystemInfo instantiates a new SystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfo(httpHeaders map[string]string, httpHost string, httpIsSecure bool, runtime SystemInfoRuntime, brand string, serverTime time.Time, embeddedOutpostDisabled bool, embeddedOutpostHost string) *SystemInfo {
	this := SystemInfo{}
	this.HttpHeaders = httpHeaders
	this.HttpHost = httpHost
	this.HttpIsSecure = httpIsSecure
	this.Runtime = runtime
	this.Brand = brand
	this.ServerTime = serverTime
	this.EmbeddedOutpostDisabled = embeddedOutpostDisabled
	this.EmbeddedOutpostHost = embeddedOutpostHost
	return &this
}

// NewSystemInfoWithDefaults instantiates a new SystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoWithDefaults() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// GetHttpHeaders returns the HttpHeaders field value
func (o *SystemInfo) GetHttpHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.HttpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetHttpHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpHeaders, true
}

// SetHttpHeaders sets field value
func (o *SystemInfo) SetHttpHeaders(v map[string]string) {
	o.HttpHeaders = v
}

// GetHttpHost returns the HttpHost field value
func (o *SystemInfo) GetHttpHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpHost
}

// GetHttpHostOk returns a tuple with the HttpHost field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetHttpHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpHost, true
}

// SetHttpHost sets field value
func (o *SystemInfo) SetHttpHost(v string) {
	o.HttpHost = v
}

// GetHttpIsSecure returns the HttpIsSecure field value
func (o *SystemInfo) GetHttpIsSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HttpIsSecure
}

// GetHttpIsSecureOk returns a tuple with the HttpIsSecure field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetHttpIsSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpIsSecure, true
}

// SetHttpIsSecure sets field value
func (o *SystemInfo) SetHttpIsSecure(v bool) {
	o.HttpIsSecure = v
}

// GetRuntime returns the Runtime field value
func (o *SystemInfo) GetRuntime() SystemInfoRuntime {
	if o == nil {
		var ret SystemInfoRuntime
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetRuntimeOk() (*SystemInfoRuntime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *SystemInfo) SetRuntime(v SystemInfoRuntime) {
	o.Runtime = v
}

// GetBrand returns the Brand field value
func (o *SystemInfo) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *SystemInfo) SetBrand(v string) {
	o.Brand = v
}

// GetServerTime returns the ServerTime field value
func (o *SystemInfo) GetServerTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetServerTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerTime, true
}

// SetServerTime sets field value
func (o *SystemInfo) SetServerTime(v time.Time) {
	o.ServerTime = v
}

// GetEmbeddedOutpostDisabled returns the EmbeddedOutpostDisabled field value
func (o *SystemInfo) GetEmbeddedOutpostDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmbeddedOutpostDisabled
}

// GetEmbeddedOutpostDisabledOk returns a tuple with the EmbeddedOutpostDisabled field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetEmbeddedOutpostDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbeddedOutpostDisabled, true
}

// SetEmbeddedOutpostDisabled sets field value
func (o *SystemInfo) SetEmbeddedOutpostDisabled(v bool) {
	o.EmbeddedOutpostDisabled = v
}

// GetEmbeddedOutpostHost returns the EmbeddedOutpostHost field value
func (o *SystemInfo) GetEmbeddedOutpostHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmbeddedOutpostHost
}

// GetEmbeddedOutpostHostOk returns a tuple with the EmbeddedOutpostHost field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetEmbeddedOutpostHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbeddedOutpostHost, true
}

// SetEmbeddedOutpostHost sets field value
func (o *SystemInfo) SetEmbeddedOutpostHost(v string) {
	o.EmbeddedOutpostHost = v
}

func (o SystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: http_headers is readOnly
	// skip: http_host is readOnly
	// skip: http_is_secure is readOnly
	toSerialize["runtime"] = o.Runtime
	// skip: brand is readOnly
	// skip: server_time is readOnly
	// skip: embedded_outpost_disabled is readOnly
	// skip: embedded_outpost_host is readOnly
	return toSerialize, nil
}

type NullableSystemInfo struct {
	value *SystemInfo
	isSet bool
}

func (v NullableSystemInfo) Get() *SystemInfo {
	return v.value
}

func (v *NullableSystemInfo) Set(val *SystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfo(val *SystemInfo) *NullableSystemInfo {
	return &NullableSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


