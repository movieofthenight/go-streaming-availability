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

// checks if the SeriesStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesStatus{}

// SeriesStatus The current status of a series.
type SeriesStatus struct {
	// Integer code of the status. 1: \"Returning Series\", 2: \"Planned\", 3: \"In Production\", 4: \"Ended\", 5: \"Cancelled\", 6: \"Pilot\", 
	StatusCode int32 `json:"statusCode"`
	// Textual representation of the status.
	StatusText string `json:"statusText"`
}

// NewSeriesStatus instantiates a new SeriesStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesStatus(statusCode int32, statusText string) *SeriesStatus {
	this := SeriesStatus{}
	this.StatusCode = statusCode
	this.StatusText = statusText
	return &this
}

// NewSeriesStatusWithDefaults instantiates a new SeriesStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesStatusWithDefaults() *SeriesStatus {
	this := SeriesStatus{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *SeriesStatus) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *SeriesStatus) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *SeriesStatus) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetStatusText returns the StatusText field value
func (o *SeriesStatus) GetStatusText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusText
}

// GetStatusTextOk returns a tuple with the StatusText field value
// and a boolean to check if the value has been set.
func (o *SeriesStatus) GetStatusTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusText, true
}

// SetStatusText sets field value
func (o *SeriesStatus) SetStatusText(v string) {
	o.StatusText = v
}

func (o SeriesStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["statusText"] = o.StatusText
	return toSerialize, nil
}

type NullableSeriesStatus struct {
	value *SeriesStatus
	isSet bool
}

func (v NullableSeriesStatus) Get() *SeriesStatus {
	return v.value
}

func (v *NullableSeriesStatus) Set(val *SeriesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesStatus(val *SeriesStatus) *NullableSeriesStatus {
	return &NullableSeriesStatus{value: val, isSet: true}
}

func (v NullableSeriesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


