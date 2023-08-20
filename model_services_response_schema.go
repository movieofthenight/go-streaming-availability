/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the ServicesResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesResponseSchema{}

// ServicesResponseSchema struct for ServicesResponseSchema
type ServicesResponseSchema struct {
	// Map of service id to service details.
	Result map[string]DeprecatedServiceInfo `json:"result"`
}

// NewServicesResponseSchema instantiates a new ServicesResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesResponseSchema(result map[string]DeprecatedServiceInfo) *ServicesResponseSchema {
	this := ServicesResponseSchema{}
	this.Result = result
	return &this
}

// NewServicesResponseSchemaWithDefaults instantiates a new ServicesResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesResponseSchemaWithDefaults() *ServicesResponseSchema {
	this := ServicesResponseSchema{}
	return &this
}

// GetResult returns the Result field value
func (o *ServicesResponseSchema) GetResult() map[string]DeprecatedServiceInfo {
	if o == nil {
		var ret map[string]DeprecatedServiceInfo
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ServicesResponseSchema) GetResultOk() (*map[string]DeprecatedServiceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ServicesResponseSchema) SetResult(v map[string]DeprecatedServiceInfo) {
	o.Result = v
}

func (o ServicesResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableServicesResponseSchema struct {
	value *ServicesResponseSchema
	isSet bool
}

func (v NullableServicesResponseSchema) Get() *ServicesResponseSchema {
	return v.value
}

func (v *NullableServicesResponseSchema) Set(val *ServicesResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesResponseSchema(val *ServicesResponseSchema) *NullableServicesResponseSchema {
	return &NullableServicesResponseSchema{value: val, isSet: true}
}

func (v NullableServicesResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


