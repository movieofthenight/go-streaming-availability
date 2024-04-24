/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
	"fmt"
)

// ShowType Type of a show.
type ShowType string

// List of showType
const (
	MOVIE ShowType = "movie"
	SERIES ShowType = "series"
)

// All allowed values of ShowType enum
var AllowedShowTypeEnumValues = []ShowType{
	"movie",
	"series",
}

func (v *ShowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShowType(value)
	for _, existing := range AllowedShowTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShowType", value)
}

// NewShowTypeFromValue returns a pointer to a valid ShowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShowTypeFromValue(v string) (*ShowType, error) {
	ev := ShowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShowType: valid values are %v", v, AllowedShowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShowType) IsValid() bool {
	for _, existing := range AllowedShowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to showType value
func (v ShowType) Ptr() *ShowType {
	return &v
}

type NullableShowType struct {
	value *ShowType
	isSet bool
}

func (v NullableShowType) Get() *ShowType {
	return v.value
}

func (v *NullableShowType) Set(val *ShowType) {
	v.value = val
	v.isSet = true
}

func (v NullableShowType) IsSet() bool {
	return v.isSet
}

func (v *NullableShowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowType(val *ShowType) *NullableShowType {
	return &NullableShowType{value: val, isSet: true}
}

func (v NullableShowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

