# Season

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | **string** | Type of the item. Always \&quot;season\&quot;. | 
**Title** | **string** | Title of the season. | 
**FirstAirYear** | **int32** | The first year that the season aired. | 
**LastAirYear** | **int32** | The last year that the season aired. | 
**StreamingOptions** |  | Map of the streaming options by the country code. | 
**Episodes** | Pointer to [**[]Episode**](Episode.md) | Array of the episodes belong to the season. | [optional] 

## Methods

### NewSeason

`func NewSeason(itemType string, title string, firstAirYear int32, lastAirYear int32, streamingOptions map[string][]StreamingOption, ) *Season`

NewSeason instantiates a new Season object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonWithDefaults

`func NewSeasonWithDefaults() *Season`

NewSeasonWithDefaults instantiates a new Season object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *Season) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Season) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Season) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetTitle

`func (o *Season) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Season) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Season) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFirstAirYear

`func (o *Season) GetFirstAirYear() int32`

GetFirstAirYear returns the FirstAirYear field if non-nil, zero value otherwise.

### GetFirstAirYearOk

`func (o *Season) GetFirstAirYearOk() (*int32, bool)`

GetFirstAirYearOk returns a tuple with the FirstAirYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirYear

`func (o *Season) SetFirstAirYear(v int32)`

SetFirstAirYear sets FirstAirYear field to given value.


### GetLastAirYear

`func (o *Season) GetLastAirYear() int32`

GetLastAirYear returns the LastAirYear field if non-nil, zero value otherwise.

### GetLastAirYearOk

`func (o *Season) GetLastAirYearOk() (*int32, bool)`

GetLastAirYearOk returns a tuple with the LastAirYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAirYear

`func (o *Season) SetLastAirYear(v int32)`

SetLastAirYear sets LastAirYear field to given value.


### GetStreamingOptions

`func (o *Season) GetStreamingOptions() map[string][]StreamingOption`

GetStreamingOptions returns the StreamingOptions field if non-nil, zero value otherwise.

### GetStreamingOptionsOk

`func (o *Season) GetStreamingOptionsOk() (*map[string][]StreamingOption, bool)`

GetStreamingOptionsOk returns a tuple with the StreamingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOptions

`func (o *Season) SetStreamingOptions(v map[string][]StreamingOption)`

SetStreamingOptions sets StreamingOptions field to given value.


### GetEpisodes

`func (o *Season) GetEpisodes() []Episode`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *Season) GetEpisodesOk() (*[]Episode, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *Season) SetEpisodes(v []Episode)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *Season) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


