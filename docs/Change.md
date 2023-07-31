# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Season** | Pointer to **int32** | Season number. Omitted if \&quot;target_type\&quot; is not \&quot;season\&quot;or \&quot;episode\&quot;. | [optional] 
**Episode** | Pointer to **int32** | Episode number. Omitted if \&quot;target_type\&quot; is not \&quot;episode\&quot;. | [optional] 
**Service** | **string** | Service id of the change. | 
**StreamingType** | [**StreamingType**](StreamingType.md) |  | 
**Addon** | Pointer to **string** | Addon id, if the \&quot;streamingType\&quot; is \&quot;addon\&quot;. Otherwise omitted. | [optional] 
**Time** | **int32** | [Unix Time Stamp](https://www.unixtimestamp.com/) of the change.  | 

## Methods

### NewChange

`func NewChange(service string, streamingType StreamingType, time int32, ) *Change`

NewChange instantiates a new Change object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWithDefaults

`func NewChangeWithDefaults() *Change`

NewChangeWithDefaults instantiates a new Change object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeason

`func (o *Change) GetSeason() int32`

GetSeason returns the Season field if non-nil, zero value otherwise.

### GetSeasonOk

`func (o *Change) GetSeasonOk() (*int32, bool)`

GetSeasonOk returns a tuple with the Season field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeason

`func (o *Change) SetSeason(v int32)`

SetSeason sets Season field to given value.

### HasSeason

`func (o *Change) HasSeason() bool`

HasSeason returns a boolean if a field has been set.

### GetEpisode

`func (o *Change) GetEpisode() int32`

GetEpisode returns the Episode field if non-nil, zero value otherwise.

### GetEpisodeOk

`func (o *Change) GetEpisodeOk() (*int32, bool)`

GetEpisodeOk returns a tuple with the Episode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisode

`func (o *Change) SetEpisode(v int32)`

SetEpisode sets Episode field to given value.

### HasEpisode

`func (o *Change) HasEpisode() bool`

HasEpisode returns a boolean if a field has been set.

### GetService

`func (o *Change) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Change) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Change) SetService(v string)`

SetService sets Service field to given value.


### GetStreamingType

`func (o *Change) GetStreamingType() StreamingType`

GetStreamingType returns the StreamingType field if non-nil, zero value otherwise.

### GetStreamingTypeOk

`func (o *Change) GetStreamingTypeOk() (*StreamingType, bool)`

GetStreamingTypeOk returns a tuple with the StreamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingType

`func (o *Change) SetStreamingType(v StreamingType)`

SetStreamingType sets StreamingType field to given value.


### GetAddon

`func (o *Change) GetAddon() string`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *Change) GetAddonOk() (*string, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *Change) SetAddon(v string)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *Change) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetTime

`func (o *Change) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Change) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Change) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


