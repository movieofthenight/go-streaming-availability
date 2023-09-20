/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

API version: 3.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the SearchTitleResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchTitleResponseSchema{}

// SearchTitleResponseSchema struct for SearchTitleResponseSchema
type SearchTitleResponseSchema struct {
	// Array of the shows matched with the title.
	Result []Show `json:"result"`
}

// NewSearchTitleResponseSchema instantiates a new SearchTitleResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTitleResponseSchema(result []Show) *SearchTitleResponseSchema {
	this := SearchTitleResponseSchema{}
	this.Result = result
	return &this
}

// NewSearchTitleResponseSchemaWithDefaults instantiates a new SearchTitleResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTitleResponseSchemaWithDefaults() *SearchTitleResponseSchema {
	this := SearchTitleResponseSchema{}
	return &this
}

// GetResult returns the Result field value
func (o *SearchTitleResponseSchema) GetResult() []Show {
	if o == nil {
		var ret []Show
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SearchTitleResponseSchema) GetResultOk() ([]Show, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *SearchTitleResponseSchema) SetResult(v []Show) {
	o.Result = v
}

func (o SearchTitleResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchTitleResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableSearchTitleResponseSchema struct {
	value *SearchTitleResponseSchema
	isSet bool
}

func (v NullableSearchTitleResponseSchema) Get() *SearchTitleResponseSchema {
	return v.value
}

func (v *NullableSearchTitleResponseSchema) Set(val *SearchTitleResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTitleResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTitleResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTitleResponseSchema(val *SearchTitleResponseSchema) *NullableSearchTitleResponseSchema {
	return &NullableSearchTitleResponseSchema{value: val, isSet: true}
}

func (v NullableSearchTitleResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTitleResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


