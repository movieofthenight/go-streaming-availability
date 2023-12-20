/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 3.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the NotFoundResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundResponseSchema{}

// NotFoundResponseSchema struct for NotFoundResponseSchema
type NotFoundResponseSchema struct {
	Message string `json:"message"`
}

// NewNotFoundResponseSchema instantiates a new NotFoundResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundResponseSchema(message string) *NotFoundResponseSchema {
	this := NotFoundResponseSchema{}
	this.Message = message
	return &this
}

// NewNotFoundResponseSchemaWithDefaults instantiates a new NotFoundResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundResponseSchemaWithDefaults() *NotFoundResponseSchema {
	this := NotFoundResponseSchema{}
	return &this
}

// GetMessage returns the Message field value
func (o *NotFoundResponseSchema) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotFoundResponseSchema) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotFoundResponseSchema) SetMessage(v string) {
	o.Message = v
}

func (o NotFoundResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableNotFoundResponseSchema struct {
	value *NotFoundResponseSchema
	isSet bool
}

func (v NullableNotFoundResponseSchema) Get() *NotFoundResponseSchema {
	return v.value
}

func (v *NullableNotFoundResponseSchema) Set(val *NotFoundResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundResponseSchema(val *NotFoundResponseSchema) *NullableNotFoundResponseSchema {
	return &NullableNotFoundResponseSchema{value: val, isSet: true}
}

func (v NullableNotFoundResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


