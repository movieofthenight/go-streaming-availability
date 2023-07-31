# StreamingOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | Id of the streaming service. | 
**StreamingType** | [**StreamingType**](StreamingType.md) |  | 
**Addon** | Pointer to **string** | Addon id, if the \&quot;streamingType\&quot; is \&quot;addon\&quot;. Otherwise omitted. | [optional] 
**Link** | **string** | Deep link to the streaming option&#39;s page in the streaming service. Guaranteed to be populated. | 
**VideoLink** | Pointer to **string** | Deep link to the video associated with the streaming option. Omitted if there&#39;s no direct link to the video. Might have the same value as \&quot;link\&quot;.  | [optional] 
**Quality** | Pointer to **string** | Maximum video quality of the streaming option. Omitted if the quality is unknown. | [optional] 
**Audios** | [**[]Locale**](Locale.md) | Array of the available audios. | 
**Subtitles** | [**[]Subtitle**](Subtitle.md) | Array of the available subtitles. | 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**Leaving** | Pointer to **int64** | [Unix Time Stamp](https://www.unixtimestamp.com/) of the date that this streaming option is expiring. In other words, last day to watch. A value of 1 means the streaming option is expiring soon, but there&#39;s no specific date info is found. Omitted if there&#39;s no known expiry date.  | [optional] 
**AvailableSince** | **int64** | [Unix Time Stamp](https://www.unixtimestamp.com/) of the date that this streaming option was found on the service.  | 

## Methods

### NewStreamingOption

`func NewStreamingOption(service string, streamingType StreamingType, link string, audios []Locale, subtitles []Subtitle, availableSince int64, ) *StreamingOption`

NewStreamingOption instantiates a new StreamingOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingOptionWithDefaults

`func NewStreamingOptionWithDefaults() *StreamingOption`

NewStreamingOptionWithDefaults instantiates a new StreamingOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *StreamingOption) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *StreamingOption) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *StreamingOption) SetService(v string)`

SetService sets Service field to given value.


### GetStreamingType

`func (o *StreamingOption) GetStreamingType() StreamingType`

GetStreamingType returns the StreamingType field if non-nil, zero value otherwise.

### GetStreamingTypeOk

`func (o *StreamingOption) GetStreamingTypeOk() (*StreamingType, bool)`

GetStreamingTypeOk returns a tuple with the StreamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingType

`func (o *StreamingOption) SetStreamingType(v StreamingType)`

SetStreamingType sets StreamingType field to given value.


### GetAddon

`func (o *StreamingOption) GetAddon() string`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *StreamingOption) GetAddonOk() (*string, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *StreamingOption) SetAddon(v string)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *StreamingOption) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetLink

`func (o *StreamingOption) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *StreamingOption) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *StreamingOption) SetLink(v string)`

SetLink sets Link field to given value.


### GetVideoLink

`func (o *StreamingOption) GetVideoLink() string`

GetVideoLink returns the VideoLink field if non-nil, zero value otherwise.

### GetVideoLinkOk

`func (o *StreamingOption) GetVideoLinkOk() (*string, bool)`

GetVideoLinkOk returns a tuple with the VideoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLink

`func (o *StreamingOption) SetVideoLink(v string)`

SetVideoLink sets VideoLink field to given value.

### HasVideoLink

`func (o *StreamingOption) HasVideoLink() bool`

HasVideoLink returns a boolean if a field has been set.

### GetQuality

`func (o *StreamingOption) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *StreamingOption) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *StreamingOption) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *StreamingOption) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetAudios

`func (o *StreamingOption) GetAudios() []Locale`

GetAudios returns the Audios field if non-nil, zero value otherwise.

### GetAudiosOk

`func (o *StreamingOption) GetAudiosOk() (*[]Locale, bool)`

GetAudiosOk returns a tuple with the Audios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudios

`func (o *StreamingOption) SetAudios(v []Locale)`

SetAudios sets Audios field to given value.


### GetSubtitles

`func (o *StreamingOption) GetSubtitles() []Subtitle`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *StreamingOption) GetSubtitlesOk() (*[]Subtitle, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *StreamingOption) SetSubtitles(v []Subtitle)`

SetSubtitles sets Subtitles field to given value.


### GetPrice

`func (o *StreamingOption) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StreamingOption) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *StreamingOption) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *StreamingOption) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLeaving

`func (o *StreamingOption) GetLeaving() int64`

GetLeaving returns the Leaving field if non-nil, zero value otherwise.

### GetLeavingOk

`func (o *StreamingOption) GetLeavingOk() (*int64, bool)`

GetLeavingOk returns a tuple with the Leaving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaving

`func (o *StreamingOption) SetLeaving(v int64)`

SetLeaving sets Leaving field to given value.

### HasLeaving

`func (o *StreamingOption) HasLeaving() bool`

HasLeaving returns a boolean if a field has been set.

### GetAvailableSince

`func (o *StreamingOption) GetAvailableSince() int64`

GetAvailableSince returns the AvailableSince field if non-nil, zero value otherwise.

### GetAvailableSinceOk

`func (o *StreamingOption) GetAvailableSinceOk() (*int64, bool)`

GetAvailableSinceOk returns a tuple with the AvailableSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSince

`func (o *StreamingOption) SetAvailableSince(v int64)`

SetAvailableSince sets AvailableSince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


