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

// checks if the ChangesResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangesResponseSchema{}

// ChangesResponseSchema struct for ChangesResponseSchema
type ChangesResponseSchema struct {
	// Array of the results matched with the query.
	Result []ChangeSet `json:"result"`
	// Whether there are more results to be loaded.
	HasMore bool `json:"hasMore"`
	// Cursor value to pass to get next set of the results. An empty string if \"hasMore\" is \"false\".
	NextCursor string `json:"nextCursor"`
}

// NewChangesResponseSchema instantiates a new ChangesResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangesResponseSchema(result []ChangeSet, hasMore bool, nextCursor string) *ChangesResponseSchema {
	this := ChangesResponseSchema{}
	this.Result = result
	this.HasMore = hasMore
	this.NextCursor = nextCursor
	return &this
}

// NewChangesResponseSchemaWithDefaults instantiates a new ChangesResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangesResponseSchemaWithDefaults() *ChangesResponseSchema {
	this := ChangesResponseSchema{}
	return &this
}

// GetResult returns the Result field value
func (o *ChangesResponseSchema) GetResult() []ChangeSet {
	if o == nil {
		var ret []ChangeSet
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ChangesResponseSchema) GetResultOk() ([]ChangeSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ChangesResponseSchema) SetResult(v []ChangeSet) {
	o.Result = v
}

// GetHasMore returns the HasMore field value
func (o *ChangesResponseSchema) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ChangesResponseSchema) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ChangesResponseSchema) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value
func (o *ChangesResponseSchema) GetNextCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
func (o *ChangesResponseSchema) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextCursor, true
}

// SetNextCursor sets field value
func (o *ChangesResponseSchema) SetNextCursor(v string) {
	o.NextCursor = v
}

func (o ChangesResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangesResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["hasMore"] = o.HasMore
	toSerialize["nextCursor"] = o.NextCursor
	return toSerialize, nil
}

type NullableChangesResponseSchema struct {
	value *ChangesResponseSchema
	isSet bool
}

func (v NullableChangesResponseSchema) Get() *ChangesResponseSchema {
	return v.value
}

func (v *NullableChangesResponseSchema) Set(val *ChangesResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangesResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangesResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangesResponseSchema(val *ChangesResponseSchema) *NullableChangesResponseSchema {
	return &NullableChangesResponseSchema{value: val, isSet: true}
}

func (v NullableChangesResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangesResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


