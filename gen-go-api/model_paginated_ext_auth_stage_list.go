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

// checks if the PaginatedExtAuthStageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedExtAuthStageList{}

// PaginatedExtAuthStageList struct for PaginatedExtAuthStageList
type PaginatedExtAuthStageList struct {
	Pagination Pagination `json:"pagination"`
	Results []ExtAuthStage `json:"results"`
}

// NewPaginatedExtAuthStageList instantiates a new PaginatedExtAuthStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedExtAuthStageList(pagination Pagination, results []ExtAuthStage) *PaginatedExtAuthStageList {
	this := PaginatedExtAuthStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedExtAuthStageListWithDefaults instantiates a new PaginatedExtAuthStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedExtAuthStageListWithDefaults() *PaginatedExtAuthStageList {
	this := PaginatedExtAuthStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedExtAuthStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtAuthStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedExtAuthStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedExtAuthStageList) GetResults() []ExtAuthStage {
	if o == nil {
		var ret []ExtAuthStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtAuthStageList) GetResultsOk() ([]ExtAuthStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedExtAuthStageList) SetResults(v []ExtAuthStage) {
	o.Results = v
}

func (o PaginatedExtAuthStageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedExtAuthStageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedExtAuthStageList struct {
	value *PaginatedExtAuthStageList
	isSet bool
}

func (v NullablePaginatedExtAuthStageList) Get() *PaginatedExtAuthStageList {
	return v.value
}

func (v *NullablePaginatedExtAuthStageList) Set(val *PaginatedExtAuthStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedExtAuthStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedExtAuthStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedExtAuthStageList(val *PaginatedExtAuthStageList) *NullablePaginatedExtAuthStageList {
	return &NullablePaginatedExtAuthStageList{value: val, isSet: true}
}

func (v NullablePaginatedExtAuthStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedExtAuthStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


