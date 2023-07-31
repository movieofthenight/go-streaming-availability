# Episode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the item. Always \&quot;episode\&quot;. | 
**Title** | **string** | Title of the episode. | 
**Year** | **int32** | The year that the movie was released. | 
**StreamingInfo** | [**map[string][]StreamingOption**](array.md) | Country to streaming availability info mapping of a show. | 

## Methods

### NewEpisode

`func NewEpisode(type_ string, title string, year int32, streamingInfo map[string][]StreamingOption, ) *Episode`

NewEpisode instantiates a new Episode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeWithDefaults

`func NewEpisodeWithDefaults() *Episode`

NewEpisodeWithDefaults instantiates a new Episode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Episode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Episode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Episode) SetType(v string)`

SetType sets Type field to given value.


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


### GetYear

`func (o *Episode) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Episode) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Episode) SetYear(v int32)`

SetYear sets Year field to given value.


### GetStreamingInfo

`func (o *Episode) GetStreamingInfo() map[string][]StreamingOption`

GetStreamingInfo returns the StreamingInfo field if non-nil, zero value otherwise.

### GetStreamingInfoOk

`func (o *Episode) GetStreamingInfoOk() (*map[string][]StreamingOption, bool)`

GetStreamingInfoOk returns a tuple with the StreamingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingInfo

`func (o *Episode) SetStreamingInfo(v map[string][]StreamingOption)`

SetStreamingInfo sets StreamingInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


