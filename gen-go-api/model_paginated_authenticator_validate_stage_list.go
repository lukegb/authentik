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

// checks if the PaginatedAuthenticatorValidateStageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAuthenticatorValidateStageList{}

// PaginatedAuthenticatorValidateStageList struct for PaginatedAuthenticatorValidateStageList
type PaginatedAuthenticatorValidateStageList struct {
	Pagination Pagination `json:"pagination"`
	Results []AuthenticatorValidateStage `json:"results"`
}

// NewPaginatedAuthenticatorValidateStageList instantiates a new PaginatedAuthenticatorValidateStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorValidateStageList(pagination Pagination, results []AuthenticatorValidateStage) *PaginatedAuthenticatorValidateStageList {
	this := PaginatedAuthenticatorValidateStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorValidateStageListWithDefaults instantiates a new PaginatedAuthenticatorValidateStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorValidateStageListWithDefaults() *PaginatedAuthenticatorValidateStageList {
	this := PaginatedAuthenticatorValidateStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorValidateStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorValidateStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorValidateStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorValidateStageList) GetResults() []AuthenticatorValidateStage {
	if o == nil {
		var ret []AuthenticatorValidateStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorValidateStageList) GetResultsOk() ([]AuthenticatorValidateStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorValidateStageList) SetResults(v []AuthenticatorValidateStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorValidateStageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedAuthenticatorValidateStageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedAuthenticatorValidateStageList struct {
	value *PaginatedAuthenticatorValidateStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorValidateStageList) Get() *PaginatedAuthenticatorValidateStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorValidateStageList) Set(val *PaginatedAuthenticatorValidateStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorValidateStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorValidateStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorValidateStageList(val *PaginatedAuthenticatorValidateStageList) *NullablePaginatedAuthenticatorValidateStageList {
	return &NullablePaginatedAuthenticatorValidateStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorValidateStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorValidateStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

