/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 3.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the SearchFiltersResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchFiltersResponseSchema{}

// SearchFiltersResponseSchema struct for SearchFiltersResponseSchema
type SearchFiltersResponseSchema struct {
	// Array of the results matched with the query.
	Result []Show `json:"result"`
	// Whether there are more results to be loaded.
	HasMore bool `json:"hasMore"`
	// Cursor value to pass to get next set of the results. An empty string if \"hasMore\" is \"false\".
	NextCursor string `json:"nextCursor"`
}

// NewSearchFiltersResponseSchema instantiates a new SearchFiltersResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchFiltersResponseSchema(result []Show, hasMore bool, nextCursor string) *SearchFiltersResponseSchema {
	this := SearchFiltersResponseSchema{}
	this.Result = result
	this.HasMore = hasMore
	this.NextCursor = nextCursor
	return &this
}

// NewSearchFiltersResponseSchemaWithDefaults instantiates a new SearchFiltersResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchFiltersResponseSchemaWithDefaults() *SearchFiltersResponseSchema {
	this := SearchFiltersResponseSchema{}
	return &this
}

// GetResult returns the Result field value
func (o *SearchFiltersResponseSchema) GetResult() []Show {
	if o == nil {
		var ret []Show
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SearchFiltersResponseSchema) GetResultOk() ([]Show, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *SearchFiltersResponseSchema) SetResult(v []Show) {
	o.Result = v
}

// GetHasMore returns the HasMore field value
func (o *SearchFiltersResponseSchema) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *SearchFiltersResponseSchema) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *SearchFiltersResponseSchema) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value
func (o *SearchFiltersResponseSchema) GetNextCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
func (o *SearchFiltersResponseSchema) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextCursor, true
}

// SetNextCursor sets field value
func (o *SearchFiltersResponseSchema) SetNextCursor(v string) {
	o.NextCursor = v
}

func (o SearchFiltersResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchFiltersResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["hasMore"] = o.HasMore
	toSerialize["nextCursor"] = o.NextCursor
	return toSerialize, nil
}

type NullableSearchFiltersResponseSchema struct {
	value *SearchFiltersResponseSchema
	isSet bool
}

func (v NullableSearchFiltersResponseSchema) Get() *SearchFiltersResponseSchema {
	return v.value
}

func (v *NullableSearchFiltersResponseSchema) Set(val *SearchFiltersResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFiltersResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFiltersResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFiltersResponseSchema(val *SearchFiltersResponseSchema) *NullableSearchFiltersResponseSchema {
	return &NullableSearchFiltersResponseSchema{value: val, isSet: true}
}

func (v NullableSearchFiltersResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFiltersResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


