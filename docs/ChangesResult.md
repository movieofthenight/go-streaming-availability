# ChangesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]Change**](Change.md) | Array of the changes. | 
**Shows** |  | Map of the shows by their ids. | 
**HasMore** | **bool** | Whether there are more changes available. | 
**NextCursor** | Pointer to **string** | Cursor value to pass to get the next set of changes. | [optional] 

## Methods

### NewChangesResult

`func NewChangesResult(changes []Change, shows map[string]Show, hasMore bool, ) *ChangesResult`

NewChangesResult instantiates a new ChangesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangesResultWithDefaults

`func NewChangesResultWithDefaults() *ChangesResult`

NewChangesResultWithDefaults instantiates a new ChangesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *ChangesResult) GetChanges() []Change`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChangesResult) GetChangesOk() (*[]Change, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChangesResult) SetChanges(v []Change)`

SetChanges sets Changes field to given value.


### GetShows

`func (o *ChangesResult) GetShows() map[string]Show`

GetShows returns the Shows field if non-nil, zero value otherwise.

### GetShowsOk

`func (o *ChangesResult) GetShowsOk() (*map[string]Show, bool)`

GetShowsOk returns a tuple with the Shows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShows

`func (o *ChangesResult) SetShows(v map[string]Show)`

SetShows sets Shows field to given value.


### GetHasMore

`func (o *ChangesResult) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChangesResult) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChangesResult) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *ChangesResult) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ChangesResult) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ChangesResult) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ChangesResult) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


