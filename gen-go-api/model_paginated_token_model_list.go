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

// checks if the PaginatedTokenModelList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedTokenModelList{}

// PaginatedTokenModelList struct for PaginatedTokenModelList
type PaginatedTokenModelList struct {
	Pagination Pagination `json:"pagination"`
	Results []TokenModel `json:"results"`
}

// NewPaginatedTokenModelList instantiates a new PaginatedTokenModelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTokenModelList(pagination Pagination, results []TokenModel) *PaginatedTokenModelList {
	this := PaginatedTokenModelList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedTokenModelListWithDefaults instantiates a new PaginatedTokenModelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTokenModelListWithDefaults() *PaginatedTokenModelList {
	this := PaginatedTokenModelList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedTokenModelList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedTokenModelList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedTokenModelList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedTokenModelList) GetResults() []TokenModel {
	if o == nil {
		var ret []TokenModel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedTokenModelList) GetResultsOk() ([]TokenModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedTokenModelList) SetResults(v []TokenModel) {
	o.Results = v
}

func (o PaginatedTokenModelList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedTokenModelList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedTokenModelList struct {
	value *PaginatedTokenModelList
	isSet bool
}

func (v NullablePaginatedTokenModelList) Get() *PaginatedTokenModelList {
	return v.value
}

func (v *NullablePaginatedTokenModelList) Set(val *PaginatedTokenModelList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTokenModelList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTokenModelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTokenModelList(val *PaginatedTokenModelList) *NullablePaginatedTokenModelList {
	return &NullablePaginatedTokenModelList{value: val, isSet: true}
}

func (v NullablePaginatedTokenModelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTokenModelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


