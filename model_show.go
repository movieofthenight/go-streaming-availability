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

// checks if the Show type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Show{}

// Show A show object represents a movie or a series. Type of the show is determined by the showType property, which is either movie or series. Based on this type, some properties are omitted, for example a movie does not have seasonCount and episodeCount properties.  Show object contains the details such as the title, overview, genres, cast, rating and images. You can find the streaming availability information under streamingOptions property. Each streaming option contains the service info, deep link, video quality, available audios and subtitles and more. It also includes the price if the show is available to buy or rent; and addon info if the show is available via an addon (such as Apple TV Channels, Prime Video Channels etc.).  You can also find the seasons of the series under the seasons property, and the episodes of a season under the episodes property of the season object. Via the streamingOptions property of seasons and episodes, you can get the individual streaming options of them. These streaming options include the same set of properties as the show streaming options, so you can use them to get deep links to episodes and seasons, and to see available audios and subtitles.  Note that seasons and episodes are not included in the search results unless you set the series_granularity parameter to seasons or episodes. For more info, check out the series_granularity parameter of the search endpoints.  There are multiple ways to retrieve shows. You can retrieve a show by its IMDb or TMDB id via [/shows/{id}](#get-a-show) endpoint; you can search shows by their title via [/shows/search/title](#search-shows-by-title) endpoint; and you can search shows by filters such as genres, release year, rating etc. via [/shows/search/filters](#search-shows-by-filters) endpoint. This endpoint also supports pagination and offers advanced ordering options such as ordering by popularity, rating, release year etc. 
type Show struct {
	// Type of the item. Always show.
	ItemType string `json:"itemType"`
	// Type of the show. Based on the type, some properties might be omitted.
	ShowType ShowType `json:"showType"`
	// Id of the show.
	Id string `json:"id"`
	// [IMDb](https://www.imdb.com/) id of the show.
	ImdbId string `json:"imdbId"`
	// [TMDB](https://www.themoviedb.org/) id of the show.
	TmdbId string `json:"tmdbId"`
	// Title of the show.
	Title string `json:"title"`
	// A brief overview of the overall plot of the show.
	Overview string `json:"overview"`
	// The year that the movie was released.
	ReleaseYear *int32 `json:"releaseYear,omitempty"`
	// The first year that the series aired.
	FirstAirYear *int32 `json:"firstAirYear,omitempty"`
	// The last year that the series aired.
	LastAirYear *int32 `json:"lastAirYear,omitempty"`
	// Original title of the show.
	OriginalTitle string `json:"originalTitle"`
	// Array of the genres of the show.
	Genres []Genre `json:"genres"`
	// Array of the directors of the movie.
	Directors []string `json:"directors,omitempty"`
	// Array of the creators of the series.
	Creators []string `json:"creators,omitempty"`
	// Array of the cast of the show.
	Cast []string `json:"cast"`
	// Rating of the show. This is calculated by taking the average of ratings found online from multiple sources.
	Rating int32 `json:"rating"`
	// Number of seasons that are either aired or announced for a series.
	SeasonCount *int32 `json:"seasonCount,omitempty"`
	// Number of episodes that are either aired or announced for a series.
	EpisodeCount *int32 `json:"episodeCount,omitempty"`
	// Runtime of the movie in minutes.
	Runtime *int32 `json:"runtime,omitempty"`
	// Image set of the show.
	ImageSet ShowImageSet `json:"imageSet"`
	// Map of the streaming options by the country code.
	StreamingOptions map[string][]StreamingOption `json:"streamingOptions"`
	// Array of the seasons belong to the series.
	Seasons []Season `json:"seasons,omitempty"`
}

type _Show Show

// NewShow instantiates a new Show object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShow(itemType string, showType ShowType, id string, imdbId string, tmdbId string, title string, overview string, originalTitle string, genres []Genre, cast []string, rating int32, imageSet ShowImageSet, streamingOptions map[string][]StreamingOption) *Show {
	this := Show{}
	this.ItemType = itemType
	this.ShowType = showType
	this.Id = id
	this.ImdbId = imdbId
	this.TmdbId = tmdbId
	this.Title = title
	this.Overview = overview
	this.OriginalTitle = originalTitle
	this.Genres = genres
	this.Cast = cast
	this.Rating = rating
	this.ImageSet = imageSet
	this.StreamingOptions = streamingOptions
	return &this
}

// NewShowWithDefaults instantiates a new Show object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowWithDefaults() *Show {
	this := Show{}
	return &this
}

// GetItemType returns the ItemType field value
func (o *Show) GetItemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *Show) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *Show) SetItemType(v string) {
	o.ItemType = v
}

// GetShowType returns the ShowType field value
func (o *Show) GetShowType() ShowType {
	if o == nil {
		var ret ShowType
		return ret
	}

	return o.ShowType
}

// GetShowTypeOk returns a tuple with the ShowType field value
// and a boolean to check if the value has been set.
func (o *Show) GetShowTypeOk() (*ShowType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowType, true
}

// SetShowType sets field value
func (o *Show) SetShowType(v ShowType) {
	o.ShowType = v
}

// GetId returns the Id field value
func (o *Show) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Show) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Show) SetId(v string) {
	o.Id = v
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
func (o *Show) GetTmdbId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value
// and a boolean to check if the value has been set.
func (o *Show) GetTmdbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmdbId, true
}

// SetTmdbId sets field value
func (o *Show) SetTmdbId(v string) {
	o.TmdbId = v
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

// GetOverview returns the Overview field value
func (o *Show) GetOverview() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value
// and a boolean to check if the value has been set.
func (o *Show) GetOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overview, true
}

// SetOverview sets field value
func (o *Show) SetOverview(v string) {
	o.Overview = v
}

// GetReleaseYear returns the ReleaseYear field value if set, zero value otherwise.
func (o *Show) GetReleaseYear() int32 {
	if o == nil || IsNil(o.ReleaseYear) {
		var ret int32
		return ret
	}
	return *o.ReleaseYear
}

// GetReleaseYearOk returns a tuple with the ReleaseYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetReleaseYearOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleaseYear) {
		return nil, false
	}
	return o.ReleaseYear, true
}

// HasReleaseYear returns a boolean if a field has been set.
func (o *Show) HasReleaseYear() bool {
	if o != nil && !IsNil(o.ReleaseYear) {
		return true
	}

	return false
}

// SetReleaseYear gets a reference to the given int32 and assigns it to the ReleaseYear field.
func (o *Show) SetReleaseYear(v int32) {
	o.ReleaseYear = &v
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

// GetCast returns the Cast field value
func (o *Show) GetCast() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value
// and a boolean to check if the value has been set.
func (o *Show) GetCastOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cast, true
}

// SetCast sets field value
func (o *Show) SetCast(v []string) {
	o.Cast = v
}

// GetRating returns the Rating field value
func (o *Show) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *Show) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *Show) SetRating(v int32) {
	o.Rating = v
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

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *Show) GetRuntime() int32 {
	if o == nil || IsNil(o.Runtime) {
		var ret int32
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Show) GetRuntimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *Show) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given int32 and assigns it to the Runtime field.
func (o *Show) SetRuntime(v int32) {
	o.Runtime = &v
}

// GetImageSet returns the ImageSet field value
func (o *Show) GetImageSet() ShowImageSet {
	if o == nil {
		var ret ShowImageSet
		return ret
	}

	return o.ImageSet
}

// GetImageSetOk returns a tuple with the ImageSet field value
// and a boolean to check if the value has been set.
func (o *Show) GetImageSetOk() (*ShowImageSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageSet, true
}

// SetImageSet sets field value
func (o *Show) SetImageSet(v ShowImageSet) {
	o.ImageSet = v
}

// GetStreamingOptions returns the StreamingOptions field value
func (o *Show) GetStreamingOptions() map[string][]StreamingOption {
	if o == nil {
		var ret map[string][]StreamingOption
		return ret
	}

	return o.StreamingOptions
}

// GetStreamingOptionsOk returns a tuple with the StreamingOptions field value
// and a boolean to check if the value has been set.
func (o *Show) GetStreamingOptionsOk() (map[string][]StreamingOption, bool) {
	if o == nil {
		return map[string][]StreamingOption{}, false
	}
	return o.StreamingOptions, true
}

// SetStreamingOptions sets field value
func (o *Show) SetStreamingOptions(v map[string][]StreamingOption) {
	o.StreamingOptions = v
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
	toSerialize["itemType"] = o.ItemType
	toSerialize["showType"] = o.ShowType
	toSerialize["id"] = o.Id
	toSerialize["imdbId"] = o.ImdbId
	toSerialize["tmdbId"] = o.TmdbId
	toSerialize["title"] = o.Title
	toSerialize["overview"] = o.Overview
	if !IsNil(o.ReleaseYear) {
		toSerialize["releaseYear"] = o.ReleaseYear
	}
	if !IsNil(o.FirstAirYear) {
		toSerialize["firstAirYear"] = o.FirstAirYear
	}
	if !IsNil(o.LastAirYear) {
		toSerialize["lastAirYear"] = o.LastAirYear
	}
	toSerialize["originalTitle"] = o.OriginalTitle
	toSerialize["genres"] = o.Genres
	if !IsNil(o.Directors) {
		toSerialize["directors"] = o.Directors
	}
	if !IsNil(o.Creators) {
		toSerialize["creators"] = o.Creators
	}
	toSerialize["cast"] = o.Cast
	toSerialize["rating"] = o.Rating
	if !IsNil(o.SeasonCount) {
		toSerialize["seasonCount"] = o.SeasonCount
	}
	if !IsNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	toSerialize["imageSet"] = o.ImageSet
	toSerialize["streamingOptions"] = o.StreamingOptions
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	return toSerialize, nil
}

func (o *Show) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemType",
		"showType",
		"id",
		"imdbId",
		"tmdbId",
		"title",
		"overview",
		"originalTitle",
		"genres",
		"cast",
		"rating",
		"imageSet",
		"streamingOptions",
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

	varShow := _Show{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShow)

	if err != nil {
		return err
	}

	*o = Show(varShow)

	return err
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


