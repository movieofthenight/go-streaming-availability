# Show

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | **string** | Type of the item. Always show. | 
**ShowType** | [**ShowType**](ShowType.md) | Type of the show. Based on the type, some properties might be omitted. | 
**Id** | **string** | Id of the show. | 
**ImdbId** | **string** | [IMDb](https://www.imdb.com/) id of the show. | 
**TmdbId** | **string** | [TMDB](https://www.themoviedb.org/) id of the show. | 
**Title** | **string** | Title of the show. | 
**Overview** | **string** | A brief overview of the overall plot of the show. | 
**ReleaseYear** | Pointer to **int32** | The year that the movie was released. | [optional] 
**FirstAirYear** | Pointer to **int32** | The first year that the series aired. | [optional] 
**LastAirYear** | Pointer to **int32** | The last year that the series aired. | [optional] 
**OriginalTitle** | **string** | Original title of the show. | 
**Genres** | [**[]Genre**](Genre.md) | Array of the genres of the show. | 
**Directors** | Pointer to **[]string** | Array of the directors of the movie. | [optional] 
**Creators** | Pointer to **[]string** | Array of the creators of the series. | [optional] 
**Cast** | **[]string** | Array of the cast of the show. | 
**Rating** | **int32** | Rating of the show. This is calculated by taking the average of ratings found online from multiple sources. | 
**SeasonCount** | Pointer to **int32** | Number of seasons that are either aired or announced for a series. | [optional] 
**EpisodeCount** | Pointer to **int32** | Number of episodes that are either aired or announced for a series. | [optional] 
**Runtime** | Pointer to **int32** | Runtime of the movie in minutes. | [optional] 
**ImageSet** | [**ShowImageSet**](ShowImageSet.md) | Image set of the show. | 
**StreamingOptions** |  | Map of the streaming options by the country code. | 
**Seasons** | Pointer to [**[]Season**](Season.md) | Array of the seasons belong to the series. | [optional] 

## Methods

### NewShow

`func NewShow(itemType string, showType ShowType, id string, imdbId string, tmdbId string, title string, overview string, originalTitle string, genres []Genre, cast []string, rating int32, imageSet ShowImageSet, streamingOptions map[string][]StreamingOption, ) *Show`

NewShow instantiates a new Show object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowWithDefaults

`func NewShowWithDefaults() *Show`

NewShowWithDefaults instantiates a new Show object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *Show) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Show) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Show) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetShowType

`func (o *Show) GetShowType() ShowType`

GetShowType returns the ShowType field if non-nil, zero value otherwise.

### GetShowTypeOk

`func (o *Show) GetShowTypeOk() (*ShowType, bool)`

GetShowTypeOk returns a tuple with the ShowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowType

`func (o *Show) SetShowType(v ShowType)`

SetShowType sets ShowType field to given value.


### GetId

`func (o *Show) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Show) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Show) SetId(v string)`

SetId sets Id field to given value.


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

`func (o *Show) GetTmdbId() string`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *Show) GetTmdbIdOk() (*string, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *Show) SetTmdbId(v string)`

SetTmdbId sets TmdbId field to given value.


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


### GetOverview

`func (o *Show) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Show) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Show) SetOverview(v string)`

SetOverview sets Overview field to given value.


### GetReleaseYear

`func (o *Show) GetReleaseYear() int32`

GetReleaseYear returns the ReleaseYear field if non-nil, zero value otherwise.

### GetReleaseYearOk

`func (o *Show) GetReleaseYearOk() (*int32, bool)`

GetReleaseYearOk returns a tuple with the ReleaseYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseYear

`func (o *Show) SetReleaseYear(v int32)`

SetReleaseYear sets ReleaseYear field to given value.

### HasReleaseYear

`func (o *Show) HasReleaseYear() bool`

HasReleaseYear returns a boolean if a field has been set.

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

### GetCast

`func (o *Show) GetCast() []string`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *Show) GetCastOk() (*[]string, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *Show) SetCast(v []string)`

SetCast sets Cast field to given value.


### GetRating

`func (o *Show) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Show) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Show) SetRating(v int32)`

SetRating sets Rating field to given value.


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

### GetRuntime

`func (o *Show) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Show) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Show) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Show) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetImageSet

`func (o *Show) GetImageSet() ShowImageSet`

GetImageSet returns the ImageSet field if non-nil, zero value otherwise.

### GetImageSetOk

`func (o *Show) GetImageSetOk() (*ShowImageSet, bool)`

GetImageSetOk returns a tuple with the ImageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSet

`func (o *Show) SetImageSet(v ShowImageSet)`

SetImageSet sets ImageSet field to given value.


### GetStreamingOptions

`func (o *Show) GetStreamingOptions() map[string][]StreamingOption`

GetStreamingOptions returns the StreamingOptions field if non-nil, zero value otherwise.

### GetStreamingOptionsOk

`func (o *Show) GetStreamingOptionsOk() (*map[string][]StreamingOption, bool)`

GetStreamingOptionsOk returns a tuple with the StreamingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOptions

`func (o *Show) SetStreamingOptions(v map[string][]StreamingOption)`

SetStreamingOptions sets StreamingOptions field to given value.


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


