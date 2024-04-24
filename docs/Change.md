# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | [**ChangeType**](ChangeType.md) | Type of the change. | 
**ItemType** | [**ItemType**](ItemType.md) | Type of the item affected from the change. | 
**ShowId** | **string** | Id of the show affected from the change. | 
**ShowType** | [**ShowType**](ShowType.md) | Type of the show affected from the change. | 
**Season** | Pointer to **int32** | Number of the season affected from the change. Omitted if item_type is not seasonor episode. | [optional] 
**Episode** | Pointer to **int32** | Number of the episode affected from the change. Omitted if item_type is not episode. | [optional] 
**Service** | [**ServiceInfo**](ServiceInfo.md) | Service affected from the change. | 
**StreamingOptionType** | [**StreamingOptionType**](StreamingOptionType.md) |  | 
**Addon** | Pointer to [**Addon**](Addon.md) | Addon info, if the streamingOptionType is addon. Otherwise omitted. | [optional] 
**Timestamp** | Pointer to **int64** | [Unix Time Stamp](https://www.unixtimestamp.com/) of the change. Past changes (new, updated, removed) will always have a timestamp. Future changes (expiring, upcoming) will have a timestamp if the exact date is known. If not, timestamp will be omitted, e.g. a show is known to be expiring soon, but the exact date is not known.  | [optional] 
**Link** | Pointer to **string** | Deep link to the affected streaming option&#39;s page in the web app of the streaming service. This field is guaranteed to be populated when changeType is new, updated, expiring or removed. When changeType is upcoming, this field might be populated or null depending on if the link of the future streaming option is known.  | [optional] 

## Methods

### NewChange

`func NewChange(changeType ChangeType, itemType ItemType, showId string, showType ShowType, service ServiceInfo, streamingOptionType StreamingOptionType, ) *Change`

NewChange instantiates a new Change object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWithDefaults

`func NewChangeWithDefaults() *Change`

NewChangeWithDefaults instantiates a new Change object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *Change) GetChangeType() ChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *Change) GetChangeTypeOk() (*ChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *Change) SetChangeType(v ChangeType)`

SetChangeType sets ChangeType field to given value.


### GetItemType

`func (o *Change) GetItemType() ItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Change) GetItemTypeOk() (*ItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Change) SetItemType(v ItemType)`

SetItemType sets ItemType field to given value.


### GetShowId

`func (o *Change) GetShowId() string`

GetShowId returns the ShowId field if non-nil, zero value otherwise.

### GetShowIdOk

`func (o *Change) GetShowIdOk() (*string, bool)`

GetShowIdOk returns a tuple with the ShowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowId

`func (o *Change) SetShowId(v string)`

SetShowId sets ShowId field to given value.


### GetShowType

`func (o *Change) GetShowType() ShowType`

GetShowType returns the ShowType field if non-nil, zero value otherwise.

### GetShowTypeOk

`func (o *Change) GetShowTypeOk() (*ShowType, bool)`

GetShowTypeOk returns a tuple with the ShowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowType

`func (o *Change) SetShowType(v ShowType)`

SetShowType sets ShowType field to given value.


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

`func (o *Change) GetService() ServiceInfo`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Change) GetServiceOk() (*ServiceInfo, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Change) SetService(v ServiceInfo)`

SetService sets Service field to given value.


### GetStreamingOptionType

`func (o *Change) GetStreamingOptionType() StreamingOptionType`

GetStreamingOptionType returns the StreamingOptionType field if non-nil, zero value otherwise.

### GetStreamingOptionTypeOk

`func (o *Change) GetStreamingOptionTypeOk() (*StreamingOptionType, bool)`

GetStreamingOptionTypeOk returns a tuple with the StreamingOptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOptionType

`func (o *Change) SetStreamingOptionType(v StreamingOptionType)`

SetStreamingOptionType sets StreamingOptionType field to given value.


### GetAddon

`func (o *Change) GetAddon() Addon`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *Change) GetAddonOk() (*Addon, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *Change) SetAddon(v Addon)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *Change) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetTimestamp

`func (o *Change) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Change) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Change) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Change) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLink

`func (o *Change) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Change) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Change) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Change) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


