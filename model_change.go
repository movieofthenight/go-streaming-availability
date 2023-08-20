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

// checks if the Change type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Change{}

// Change A past or upcoming change in a catalog.
type Change struct {
	// Season number. Omitted if \"target_type\" is not \"season\"or \"episode\".
	Season *int32 `json:"season,omitempty"`
	// Episode number. Omitted if \"target_type\" is not \"episode\".
	Episode *int32 `json:"episode,omitempty"`
	// Service id of the change.
	Service string `json:"service"`
	StreamingType StreamingType `json:"streamingType"`
	// Addon id, if the \"streamingType\" is \"addon\". Otherwise omitted.
	Addon *string `json:"addon,omitempty"`
	// [Unix Time Stamp](https://www.unixtimestamp.com/) of the change. 
	Time int32 `json:"time"`
}

// NewChange instantiates a new Change object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChange(service string, streamingType StreamingType, time int32) *Change {
	this := Change{}
	this.Service = service
	this.StreamingType = streamingType
	this.Time = time
	return &this
}

// NewChangeWithDefaults instantiates a new Change object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWithDefaults() *Change {
	this := Change{}
	return &this
}

// GetSeason returns the Season field value if set, zero value otherwise.
func (o *Change) GetSeason() int32 {
	if o == nil || IsNil(o.Season) {
		var ret int32
		return ret
	}
	return *o.Season
}

// GetSeasonOk returns a tuple with the Season field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetSeasonOk() (*int32, bool) {
	if o == nil || IsNil(o.Season) {
		return nil, false
	}
	return o.Season, true
}

// HasSeason returns a boolean if a field has been set.
func (o *Change) HasSeason() bool {
	if o != nil && !IsNil(o.Season) {
		return true
	}

	return false
}

// SetSeason gets a reference to the given int32 and assigns it to the Season field.
func (o *Change) SetSeason(v int32) {
	o.Season = &v
}

// GetEpisode returns the Episode field value if set, zero value otherwise.
func (o *Change) GetEpisode() int32 {
	if o == nil || IsNil(o.Episode) {
		var ret int32
		return ret
	}
	return *o.Episode
}

// GetEpisodeOk returns a tuple with the Episode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetEpisodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Episode) {
		return nil, false
	}
	return o.Episode, true
}

// HasEpisode returns a boolean if a field has been set.
func (o *Change) HasEpisode() bool {
	if o != nil && !IsNil(o.Episode) {
		return true
	}

	return false
}

// SetEpisode gets a reference to the given int32 and assigns it to the Episode field.
func (o *Change) SetEpisode(v int32) {
	o.Episode = &v
}

// GetService returns the Service field value
func (o *Change) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *Change) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *Change) SetService(v string) {
	o.Service = v
}

// GetStreamingType returns the StreamingType field value
func (o *Change) GetStreamingType() StreamingType {
	if o == nil {
		var ret StreamingType
		return ret
	}

	return o.StreamingType
}

// GetStreamingTypeOk returns a tuple with the StreamingType field value
// and a boolean to check if the value has been set.
func (o *Change) GetStreamingTypeOk() (*StreamingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingType, true
}

// SetStreamingType sets field value
func (o *Change) SetStreamingType(v StreamingType) {
	o.StreamingType = v
}

// GetAddon returns the Addon field value if set, zero value otherwise.
func (o *Change) GetAddon() string {
	if o == nil || IsNil(o.Addon) {
		var ret string
		return ret
	}
	return *o.Addon
}

// GetAddonOk returns a tuple with the Addon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetAddonOk() (*string, bool) {
	if o == nil || IsNil(o.Addon) {
		return nil, false
	}
	return o.Addon, true
}

// HasAddon returns a boolean if a field has been set.
func (o *Change) HasAddon() bool {
	if o != nil && !IsNil(o.Addon) {
		return true
	}

	return false
}

// SetAddon gets a reference to the given string and assigns it to the Addon field.
func (o *Change) SetAddon(v string) {
	o.Addon = &v
}

// GetTime returns the Time field value
func (o *Change) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *Change) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *Change) SetTime(v int32) {
	o.Time = v
}

func (o Change) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Change) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Season) {
		toSerialize["season"] = o.Season
	}
	if !IsNil(o.Episode) {
		toSerialize["episode"] = o.Episode
	}
	toSerialize["service"] = o.Service
	toSerialize["streamingType"] = o.StreamingType
	if !IsNil(o.Addon) {
		toSerialize["addon"] = o.Addon
	}
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

type NullableChange struct {
	value *Change
	isSet bool
}

func (v NullableChange) Get() *Change {
	return v.value
}

func (v *NullableChange) Set(val *Change) {
	v.value = val
	v.isSet = true
}

func (v NullableChange) IsSet() bool {
	return v.isSet
}

func (v *NullableChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChange(val *Change) *NullableChange {
	return &NullableChange{value: val, isSet: true}
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


