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

// ItemType Type of an item.
type ItemType string

// List of itemType
const (
	SHOW ItemType = "show"
	SEASON ItemType = "season"
	EPISODE ItemType = "episode"
)

// All allowed values of ItemType enum
var AllowedItemTypeEnumValues = []ItemType{
	"show",
	"season",
	"episode",
}

func (v *ItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemType(value)
	for _, existing := range AllowedItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemType", value)
}

// NewItemTypeFromValue returns a pointer to a valid ItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemTypeFromValue(v string) (*ItemType, error) {
	ev := ItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemType: valid values are %v", v, AllowedItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemType) IsValid() bool {
	for _, existing := range AllowedItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to itemType value
func (v ItemType) Ptr() *ItemType {
	return &v
}

type NullableItemType struct {
	value *ItemType
	isSet bool
}

func (v NullableItemType) Get() *ItemType {
	return v.value
}

func (v *NullableItemType) Set(val *ItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemType(val *ItemType) *NullableItemType {
	return &NullableItemType{value: val, isSet: true}
}

func (v NullableItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
