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

// checks if the DeprecatedServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedServiceInfo{}

// DeprecatedServiceInfo Details of a service.
type DeprecatedServiceInfo struct {
	// Id of the service.
	Id string `json:"id"`
	// Map of 2-letter country ISO code to details of the service in that country.
	Countries map[string]ServiceCountryInfo `json:"countries"`
}

// NewDeprecatedServiceInfo instantiates a new DeprecatedServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedServiceInfo(id string, countries map[string]ServiceCountryInfo) *DeprecatedServiceInfo {
	this := DeprecatedServiceInfo{}
	this.Id = id
	this.Countries = countries
	return &this
}

// NewDeprecatedServiceInfoWithDefaults instantiates a new DeprecatedServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedServiceInfoWithDefaults() *DeprecatedServiceInfo {
	this := DeprecatedServiceInfo{}
	return &this
}

// GetId returns the Id field value
func (o *DeprecatedServiceInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServiceInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeprecatedServiceInfo) SetId(v string) {
	o.Id = v
}

// GetCountries returns the Countries field value
func (o *DeprecatedServiceInfo) GetCountries() map[string]ServiceCountryInfo {
	if o == nil {
		var ret map[string]ServiceCountryInfo
		return ret
	}

	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value
// and a boolean to check if the value has been set.
func (o *DeprecatedServiceInfo) GetCountriesOk() (*map[string]ServiceCountryInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Countries, true
}

// SetCountries sets field value
func (o *DeprecatedServiceInfo) SetCountries(v map[string]ServiceCountryInfo) {
	o.Countries = v
}

func (o DeprecatedServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["countries"] = o.Countries
	return toSerialize, nil
}

type NullableDeprecatedServiceInfo struct {
	value *DeprecatedServiceInfo
	isSet bool
}

func (v NullableDeprecatedServiceInfo) Get() *DeprecatedServiceInfo {
	return v.value
}

func (v *NullableDeprecatedServiceInfo) Set(val *DeprecatedServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedServiceInfo(val *DeprecatedServiceInfo) *NullableDeprecatedServiceInfo {
	return &NullableDeprecatedServiceInfo{value: val, isSet: true}
}

func (v NullableDeprecatedServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


