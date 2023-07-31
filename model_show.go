/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the Show type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Show{}

// Show Details of a show.
type Show struct {
	// Type of the show.
	Type string `json:"type"`
	// Title of the show.
	Title string `json:"title"`
	// The year that the movie was released.
	Year *int32 `json:"year,omitempty"`
	// The first year that the series aired.
	FirstAirYear *int32 `json:"firstAirYear,omitempty"`
	// The last year that the series aired.
	LastAirYear *int32 `json:"lastAirYear,omitempty"`
	// [IMDb](https://www.imdb.com/) ID of the show.
	ImdbId string `json:"imdbId"`
	// [TMDb](https://www.themoviedb.org/) ID of the show.
	TmdbId int32 `json:"tmdbId"`
	// Original title of the show.
	OriginalTitle string `json:"originalTitle"`
	// Array of the genres associated with the show.
	Genres []Genre `json:"genres"`
	// Array of the directors of the movie.
	Directors []string `json:"directors,omitempty"`
	// Array of the creators of the series.
	Creators []string `json:"creators,omitempty"`
	Status *SeriesStatus `json:"status,omitempty"`
	// Number of seasons that are either available or announced.
	SeasonCount *int32 `json:"seasonCount,omitempty"`
	// Number of episodes that are either available or announced.
	EpisodeCount *int32 `json:"episodeCount,omitempty"`
	// Country to streaming availability info mapping of a show.
	StreamingInfo map[string][]StreamingOption `json:"streamingInfo"`
	Seasons []Season `json:"seasons,omitempty"`
}

// NewShow instantiates a new Show object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShow(type_ string, title string, imdbId string, tmdbId int32, originalTitle string, genres []Genre, streamingInfo map[string][]StreamingOption) *Show {
	this := Show{}
	this.Type = type_
	this.Title = title
	this.ImdbId = imdbId
	this.TmdbId = tmdbId
	this.OriginalTitle = originalTitle
	this.Genres = genres
	this.StreamingInfo = streamingInfo
	return &this
}

// NewShowWithDefaults instantiates a new Show object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowWithDefaults() *Show {
	this := Show{}
	return &this
}

// GetType returns the Type field value
func (o *Show) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Show) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Show) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *Show) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Show) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Show) SetTitle(v string) {
	o.Title = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Show) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Show) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *Show) SetYear(v int32) {
	o.Year = &v
}

// GetFirstAirYear returns the FirstAirYear field value if set, zero value otherwise.
func (o *Show) GetFirstAirYear() int32 {
	if o == nil || IsNil(o.FirstAirYear) {
		var ret int32
		return ret
	}
	return *o.FirstAirYear
}

// GetFirstAirYearOk returns a tuple with the FirstAirYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetFirstAirYearOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstAirYear) {
		return nil, false
	}
	return o.FirstAirYear, true
}

// HasFirstAirYear returns a boolean if a field has been set.
func (o *Show) HasFirstAirYear() bool {
	if o != nil && !IsNil(o.FirstAirYear) {
		return true
	}

	return false
}

// SetFirstAirYear gets a reference to the given int32 and assigns it to the FirstAirYear field.
func (o *Show) SetFirstAirYear(v int32) {
	o.FirstAirYear = &v
}

// GetLastAirYear returns the LastAirYear field value if set, zero value otherwise.
func (o *Show) GetLastAirYear() int32 {
	if o == nil || IsNil(o.LastAirYear) {
		var ret int32
		return ret
	}
	return *o.LastAirYear
}

// GetLastAirYearOk returns a tuple with the LastAirYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetLastAirYearOk() (*int32, bool) {
	if o == nil || IsNil(o.LastAirYear) {
		return nil, false
	}
	return o.LastAirYear, true
}

// HasLastAirYear returns a boolean if a field has been set.
func (o *Show) HasLastAirYear() bool {
	if o != nil && !IsNil(o.LastAirYear) {
		return true
	}

	return false
}

// SetLastAirYear gets a reference to the given int32 and assigns it to the LastAirYear field.
func (o *Show) SetLastAirYear(v int32) {
	o.LastAirYear = &v
}

// GetImdbId returns the ImdbId field value
func (o *Show) GetImdbId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImdbId
}

// GetImdbIdOk returns a tuple with the ImdbId field value
// and a boolean to check if the value has been set.
func (o *Show) GetImdbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImdbId, true
}

// SetImdbId sets field value
func (o *Show) SetImdbId(v string) {
	o.ImdbId = v
}

// GetTmdbId returns the TmdbId field value
func (o *Show) GetTmdbId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value
// and a boolean to check if the value has been set.
func (o *Show) GetTmdbIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmdbId, true
}

// SetTmdbId sets field value
func (o *Show) SetTmdbId(v int32) {
	o.TmdbId = v
}

// GetOriginalTitle returns the OriginalTitle field value
func (o *Show) GetOriginalTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTitle
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value
// and a boolean to check if the value has been set.
func (o *Show) GetOriginalTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTitle, true
}

// SetOriginalTitle sets field value
func (o *Show) SetOriginalTitle(v string) {
	o.OriginalTitle = v
}

// GetGenres returns the Genres field value
func (o *Show) GetGenres() []Genre {
	if o == nil {
		var ret []Genre
		return ret
	}

	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value
// and a boolean to check if the value has been set.
func (o *Show) GetGenresOk() ([]Genre, bool) {
	if o == nil {
		return nil, false
	}
	return o.Genres, true
}

// SetGenres sets field value
func (o *Show) SetGenres(v []Genre) {
	o.Genres = v
}

// GetDirectors returns the Directors field value if set, zero value otherwise.
func (o *Show) GetDirectors() []string {
	if o == nil || IsNil(o.Directors) {
		var ret []string
		return ret
	}
	return o.Directors
}

// GetDirectorsOk returns a tuple with the Directors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetDirectorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Directors) {
		return nil, false
	}
	return o.Directors, true
}

// HasDirectors returns a boolean if a field has been set.
func (o *Show) HasDirectors() bool {
	if o != nil && !IsNil(o.Directors) {
		return true
	}

	return false
}

// SetDirectors gets a reference to the given []string and assigns it to the Directors field.
func (o *Show) SetDirectors(v []string) {
	o.Directors = v
}

// GetCreators returns the Creators field value if set, zero value otherwise.
func (o *Show) GetCreators() []string {
	if o == nil || IsNil(o.Creators) {
		var ret []string
		return ret
	}
	return o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetCreatorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Creators) {
		return nil, false
	}
	return o.Creators, true
}

// HasCreators returns a boolean if a field has been set.
func (o *Show) HasCreators() bool {
	if o != nil && !IsNil(o.Creators) {
		return true
	}

	return false
}

// SetCreators gets a reference to the given []string and assigns it to the Creators field.
func (o *Show) SetCreators(v []string) {
	o.Creators = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Show) GetStatus() SeriesStatus {
	if o == nil || IsNil(o.Status) {
		var ret SeriesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetStatusOk() (*SeriesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Show) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SeriesStatus and assigns it to the Status field.
func (o *Show) SetStatus(v SeriesStatus) {
	o.Status = &v
}

// GetSeasonCount returns the SeasonCount field value if set, zero value otherwise.
func (o *Show) GetSeasonCount() int32 {
	if o == nil || IsNil(o.SeasonCount) {
		var ret int32
		return ret
	}
	return *o.SeasonCount
}

// GetSeasonCountOk returns a tuple with the SeasonCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetSeasonCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeasonCount) {
		return nil, false
	}
	return o.SeasonCount, true
}

// HasSeasonCount returns a boolean if a field has been set.
func (o *Show) HasSeasonCount() bool {
	if o != nil && !IsNil(o.SeasonCount) {
		return true
	}

	return false
}

// SetSeasonCount gets a reference to the given int32 and assigns it to the SeasonCount field.
func (o *Show) SetSeasonCount(v int32) {
	o.SeasonCount = &v
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *Show) GetEpisodeCount() int32 {
	if o == nil || IsNil(o.EpisodeCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetEpisodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EpisodeCount) {
		return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *Show) HasEpisodeCount() bool {
	if o != nil && !IsNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given int32 and assigns it to the EpisodeCount field.
func (o *Show) SetEpisodeCount(v int32) {
	o.EpisodeCount = &v
}

// GetStreamingInfo returns the StreamingInfo field value
func (o *Show) GetStreamingInfo() map[string][]StreamingOption {
	if o == nil {
		var ret map[string][]StreamingOption
		return ret
	}

	return o.StreamingInfo
}

// GetStreamingInfoOk returns a tuple with the StreamingInfo field value
// and a boolean to check if the value has been set.
func (o *Show) GetStreamingInfoOk() (*map[string][]StreamingOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingInfo, true
}

// SetStreamingInfo sets field value
func (o *Show) SetStreamingInfo(v map[string][]StreamingOption) {
	o.StreamingInfo = v
}

// GetSeasons returns the Seasons field value if set, zero value otherwise.
func (o *Show) GetSeasons() []Season {
	if o == nil || IsNil(o.Seasons) {
		var ret []Season
		return ret
	}
	return o.Seasons
}

// GetSeasonsOk returns a tuple with the Seasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetSeasonsOk() ([]Season, bool) {
	if o == nil || IsNil(o.Seasons) {
		return nil, false
	}
	return o.Seasons, true
}

// HasSeasons returns a boolean if a field has been set.
func (o *Show) HasSeasons() bool {
	if o != nil && !IsNil(o.Seasons) {
		return true
	}

	return false
}

// SetSeasons gets a reference to the given []Season and assigns it to the Seasons field.
func (o *Show) SetSeasons(v []Season) {
	o.Seasons = v
}

func (o Show) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Show) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.FirstAirYear) {
		toSerialize["firstAirYear"] = o.FirstAirYear
	}
	if !IsNil(o.LastAirYear) {
		toSerialize["lastAirYear"] = o.LastAirYear
	}
	toSerialize["imdbId"] = o.ImdbId
	toSerialize["tmdbId"] = o.TmdbId
	toSerialize["originalTitle"] = o.OriginalTitle
	toSerialize["genres"] = o.Genres
	if !IsNil(o.Directors) {
		toSerialize["directors"] = o.Directors
	}
	if !IsNil(o.Creators) {
		toSerialize["creators"] = o.Creators
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SeasonCount) {
		toSerialize["seasonCount"] = o.SeasonCount
	}
	if !IsNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	toSerialize["streamingInfo"] = o.StreamingInfo
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	return toSerialize, nil
}

type NullableShow struct {
	value *Show
	isSet bool
}

func (v NullableShow) Get() *Show {
	return v.value
}

func (v *NullableShow) Set(val *Show) {
	v.value = val
	v.isSet = true
}

func (v NullableShow) IsSet() bool {
	return v.isSet
}

func (v *NullableShow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShow(val *Show) *NullableShow {
	return &NullableShow{value: val, isSet: true}
}

func (v NullableShow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


