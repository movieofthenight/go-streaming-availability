# Show

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the show. | 
**Title** | **string** | Title of the show. | 
**Year** | Pointer to **int32** | The year that the movie was released. | [optional] 
**FirstAirYear** | Pointer to **int32** | The first year that the series aired. | [optional] 
**LastAirYear** | Pointer to **int32** | The last year that the series aired. | [optional] 
**ImdbId** | **string** | [IMDb](https://www.imdb.com/) ID of the show. | 
**TmdbId** | **int32** | [TMDb](https://www.themoviedb.org/) ID of the show. | 
**OriginalTitle** | **string** | Original title of the show. | 
**Genres** | [**[]Genre**](Genre.md) | Array of the genres associated with the show. | 
**Directors** | Pointer to **[]string** | Array of the directors of the movie. | [optional] 
**Creators** | Pointer to **[]string** | Array of the creators of the series. | [optional] 
**Status** | Pointer to [**SeriesStatus**](SeriesStatus.md) |  | [optional] 
**SeasonCount** | Pointer to **int32** | Number of seasons that are either available or announced. | [optional] 
**EpisodeCount** | Pointer to **int32** | Number of episodes that are either available or announced. | [optional] 
**StreamingInfo** | [**map[string][]StreamingOption**](array.md) | Country to streaming availability info mapping of a show. | 
**Seasons** | Pointer to [**[]Season**](Season.md) |  | [optional] 

## Methods

### NewShow

`func NewShow(type_ string, title string, imdbId string, tmdbId int32, originalTitle string, genres []Genre, streamingInfo map[string][]StreamingOption, ) *Show`

NewShow instantiates a new Show object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowWithDefaults

`func NewShowWithDefaults() *Show`

NewShowWithDefaults instantiates a new Show object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Show) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Show) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Show) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *Show) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Show) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Show) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetYear

`func (o *Show) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Show) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Show) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Show) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetFirstAirYear

`func (o *Show) GetFirstAirYear() int32`

GetFirstAirYear returns the FirstAirYear field if non-nil, zero value otherwise.

### GetFirstAirYearOk

`func (o *Show) GetFirstAirYearOk() (*int32, bool)`

GetFirstAirYearOk returns a tuple with the FirstAirYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirYear

`func (o *Show) SetFirstAirYear(v int32)`

SetFirstAirYear sets FirstAirYear field to given value.

### HasFirstAirYear

`func (o *Show) HasFirstAirYear() bool`

HasFirstAirYear returns a boolean if a field has been set.

### GetLastAirYear

`func (o *Show) GetLastAirYear() int32`

GetLastAirYear returns the LastAirYear field if non-nil, zero value otherwise.

### GetLastAirYearOk

`func (o *Show) GetLastAirYearOk() (*int32, bool)`

GetLastAirYearOk returns a tuple with the LastAirYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAirYear

`func (o *Show) SetLastAirYear(v int32)`

SetLastAirYear sets LastAirYear field to given value.

### HasLastAirYear

`func (o *Show) HasLastAirYear() bool`

HasLastAirYear returns a boolean if a field has been set.

### GetImdbId

`func (o *Show) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *Show) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *Show) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.


### GetTmdbId

`func (o *Show) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *Show) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *Show) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.


### GetOriginalTitle

`func (o *Show) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *Show) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *Show) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.


### GetGenres

`func (o *Show) GetGenres() []Genre`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *Show) GetGenresOk() (*[]Genre, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *Show) SetGenres(v []Genre)`

SetGenres sets Genres field to given value.


### GetDirectors

`func (o *Show) GetDirectors() []string`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *Show) GetDirectorsOk() (*[]string, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *Show) SetDirectors(v []string)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *Show) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### GetCreators

`func (o *Show) GetCreators() []string`

GetCreators returns the Creators field if non-nil, zero value otherwise.

### GetCreatorsOk

`func (o *Show) GetCreatorsOk() (*[]string, bool)`

GetCreatorsOk returns a tuple with the Creators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreators

`func (o *Show) SetCreators(v []string)`

SetCreators sets Creators field to given value.

### HasCreators

`func (o *Show) HasCreators() bool`

HasCreators returns a boolean if a field has been set.

### GetStatus

`func (o *Show) GetStatus() SeriesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Show) GetStatusOk() (*SeriesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Show) SetStatus(v SeriesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Show) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeasonCount

`func (o *Show) GetSeasonCount() int32`

GetSeasonCount returns the SeasonCount field if non-nil, zero value otherwise.

### GetSeasonCountOk

`func (o *Show) GetSeasonCountOk() (*int32, bool)`

GetSeasonCountOk returns a tuple with the SeasonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonCount

`func (o *Show) SetSeasonCount(v int32)`

SetSeasonCount sets SeasonCount field to given value.

### HasSeasonCount

`func (o *Show) HasSeasonCount() bool`

HasSeasonCount returns a boolean if a field has been set.

### GetEpisodeCount

`func (o *Show) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *Show) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *Show) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *Show) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetStreamingInfo

`func (o *Show) GetStreamingInfo() map[string][]StreamingOption`

GetStreamingInfo returns the StreamingInfo field if non-nil, zero value otherwise.

### GetStreamingInfoOk

`func (o *Show) GetStreamingInfoOk() (*map[string][]StreamingOption, bool)`

GetStreamingInfoOk returns a tuple with the StreamingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingInfo

`func (o *Show) SetStreamingInfo(v map[string][]StreamingOption)`

SetStreamingInfo sets StreamingInfo field to given value.


### GetSeasons

`func (o *Show) GetSeasons() []Season`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *Show) GetSeasonsOk() (*[]Season, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *Show) SetSeasons(v []Season)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *Show) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


