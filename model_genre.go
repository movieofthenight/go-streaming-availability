/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 60 countries!

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Genre type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Genre{}

// Genre Genres are used to categorize shows. Each genre object contains the id and name of the genre. When fetching genres via [/genres](#get-all-genres) endpoint, you can pass the output_language parameter to get the genre names in one of the supported languages.  You can use genre ids to filter shows in the search endpoints. 
type Genre struct {
	// Id of a genre.
	Id string `json:"id"`
	// Name of the genre
	Name string `json:"name"`
}

type _Genre Genre

// NewGenre instantiates a new Genre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenre(id string, name string) *Genre {
	this := Genre{}
	this.Id = id
	this.Name = name
	return &this
}

// NewGenreWithDefaults instantiates a new Genre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenreWithDefaults() *Genre {
	this := Genre{}
	return &this
}

// GetId returns the Id field value
func (o *Genre) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Genre) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Genre) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Genre) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Genre) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Genre) SetName(v string) {
	o.Name = v
}

func (o Genre) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Genre) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *Genre) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGenre := _Genre{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenre)

	if err != nil {
		return err
	}

	*o = Genre(varGenre)

	return err
}

type NullableGenre struct {
	value *Genre
	isSet bool
}

func (v NullableGenre) Get() *Genre {
	return v.value
}

func (v *NullableGenre) Set(val *Genre) {
	v.value = val
	v.isSet = true
}

func (v NullableGenre) IsSet() bool {
	return v.isSet
}

func (v *NullableGenre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenre(val *Genre) *NullableGenre {
	return &NullableGenre{value: val, isSet: true}
}

func (v NullableGenre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


