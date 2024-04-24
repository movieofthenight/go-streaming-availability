# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shows** | [**[]Show**](Show.md) | Array of shows. | 
**HasMore** | **bool** | Whether there are more shows available. | 
**NextCursor** | Pointer to **string** | Cursor value to pass to get the next set of shows. | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult(shows []Show, hasMore bool, ) *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShows

`func (o *SearchResult) GetShows() []Show`

GetShows returns the Shows field if non-nil, zero value otherwise.

### GetShowsOk

`func (o *SearchResult) GetShowsOk() (*[]Show, bool)`

GetShowsOk returns a tuple with the Shows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShows

`func (o *SearchResult) SetShows(v []Show)`

SetShows sets Shows field to given value.


### GetHasMore

`func (o *SearchResult) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SearchResult) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SearchResult) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *SearchResult) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SearchResult) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SearchResult) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *SearchResult) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


