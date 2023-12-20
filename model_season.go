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

// checks if the Season type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Season{}

// Season Details of a season.
type Season struct {
	// Type of the item. Always \"season\".
	Type string `json:"type"`
	// Title of the season.
	Title string `json:"title"`
	// The first year that the season aired.
	FirstAirYear int32 `json:"firstAirYear"`
	// The last year that the season aired.
	LastAirYear int32 `json:"lastAirYear"`
	// Country to streaming availability info mapping of a show.
	StreamingInfo map[string][]StreamingOption `json:"streamingInfo"`
	// Array of the episodes belong to the season.
	Episodes []Episode `json:"episodes"`
}

// NewSeason instantiates a new Season object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeason(type_ string, title string, firstAirYear int32, lastAirYear int32, streamingInfo map[string][]StreamingOption, episodes []Episode) *Season {
	this := Season{}
	this.Type = type_
	this.Title = title
	this.FirstAirYear = firstAirYear
	this.LastAirYear = lastAirYear
	this.StreamingInfo = streamingInfo
	this.Episodes = episodes
	return &this
}

// NewSeasonWithDefaults instantiates a new Season object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeasonWithDefaults() *Season {
	this := Season{}
	return &this
}

// GetType returns the Type field value
func (o *Season) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Season) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Season) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *Season) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Season) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Season) SetTitle(v string) {
	o.Title = v
}

// GetFirstAirYear returns the FirstAirYear field value
func (o *Season) GetFirstAirYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstAirYear
}

// GetFirstAirYearOk returns a tuple with the FirstAirYear field value
// and a boolean to check if the value has been set.
func (o *Season) GetFirstAirYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstAirYear, true
}

// SetFirstAirYear sets field value
func (o *Season) SetFirstAirYear(v int32) {
	o.FirstAirYear = v
}

// GetLastAirYear returns the LastAirYear field value
func (o *Season) GetLastAirYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastAirYear
}

// GetLastAirYearOk returns a tuple with the LastAirYear field value
// and a boolean to check if the value has been set.
func (o *Season) GetLastAirYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAirYear, true
}

// SetLastAirYear sets field value
func (o *Season) SetLastAirYear(v int32) {
	o.LastAirYear = v
}

// GetStreamingInfo returns the StreamingInfo field value
func (o *Season) GetStreamingInfo() map[string][]StreamingOption {
	if o == nil {
		var ret map[string][]StreamingOption
		return ret
	}

	return o.StreamingInfo
}

// GetStreamingInfoOk returns a tuple with the StreamingInfo field value
// and a boolean to check if the value has been set.
func (o *Season) GetStreamingInfoOk() (*map[string][]StreamingOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingInfo, true
}

// SetStreamingInfo sets field value
func (o *Season) SetStreamingInfo(v map[string][]StreamingOption) {
	o.StreamingInfo = v
}

// GetEpisodes returns the Episodes field value
func (o *Season) GetEpisodes() []Episode {
	if o == nil {
		var ret []Episode
		return ret
	}

	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value
// and a boolean to check if the value has been set.
func (o *Season) GetEpisodesOk() ([]Episode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Episodes, true
}

// SetEpisodes sets field value
func (o *Season) SetEpisodes(v []Episode) {
	o.Episodes = v
}

func (o Season) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Season) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["firstAirYear"] = o.FirstAirYear
	toSerialize["lastAirYear"] = o.LastAirYear
	toSerialize["streamingInfo"] = o.StreamingInfo
	toSerialize["episodes"] = o.Episodes
	return toSerialize, nil
}

type NullableSeason struct {
	value *Season
	isSet bool
}

func (v NullableSeason) Get() *Season {
	return v.value
}

func (v *NullableSeason) Set(val *Season) {
	v.value = val
	v.isSet = true
}

func (v NullableSeason) IsSet() bool {
	return v.isSet
}

func (v *NullableSeason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeason(val *Season) *NullableSeason {
	return &NullableSeason{value: val, isSet: true}
}

func (v NullableSeason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


