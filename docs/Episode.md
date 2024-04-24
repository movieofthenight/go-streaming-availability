# Episode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | **string** | Type of the item. Always \&quot;episode\&quot;. | 
**Title** | **string** | Title of the episode. | 
**AirYear** | **int32** | The year that the episode aired. | 
**StreamingOptions** |  | Map of the streaming options by the country code. | 

## Methods

### NewEpisode

`func NewEpisode(itemType string, title string, airYear int32, streamingOptions map[string][]StreamingOption, ) *Episode`

NewEpisode instantiates a new Episode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeWithDefaults

`func NewEpisodeWithDefaults() *Episode`

NewEpisodeWithDefaults instantiates a new Episode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *Episode) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Episode) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Episode) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetTitle

`func (o *Episode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Episode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Episode) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAirYear

`func (o *Episode) GetAirYear() int32`

GetAirYear returns the AirYear field if non-nil, zero value otherwise.

### GetAirYearOk

`func (o *Episode) GetAirYearOk() (*int32, bool)`

GetAirYearOk returns a tuple with the AirYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirYear

`func (o *Episode) SetAirYear(v int32)`

SetAirYear sets AirYear field to given value.


### GetStreamingOptions

`func (o *Episode) GetStreamingOptions() map[string][]StreamingOption`

GetStreamingOptions returns the StreamingOptions field if non-nil, zero value otherwise.

### GetStreamingOptionsOk

`func (o *Episode) GetStreamingOptionsOk() (*map[string][]StreamingOption, bool)`

GetStreamingOptionsOk returns a tuple with the StreamingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOptions

`func (o *Episode) SetStreamingOptions(v map[string][]StreamingOption)`

SetStreamingOptions sets StreamingOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


