# SearchFiltersResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**[]Show**](Show.md) | Array of the results matched with the query. | 
**HasMore** | **bool** | Whether there are more results to be loaded. | 
**NextCursor** | **string** | Cursor value to pass to get next set of the results. An empty string if \&quot;hasMore\&quot; is \&quot;false\&quot;. | 

## Methods

### NewSearchFiltersResponseSchema

`func NewSearchFiltersResponseSchema(result []Show, hasMore bool, nextCursor string, ) *SearchFiltersResponseSchema`

NewSearchFiltersResponseSchema instantiates a new SearchFiltersResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFiltersResponseSchemaWithDefaults

`func NewSearchFiltersResponseSchemaWithDefaults() *SearchFiltersResponseSchema`

NewSearchFiltersResponseSchemaWithDefaults instantiates a new SearchFiltersResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SearchFiltersResponseSchema) GetResult() []Show`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchFiltersResponseSchema) GetResultOk() (*[]Show, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchFiltersResponseSchema) SetResult(v []Show)`

SetResult sets Result field to given value.


### GetHasMore

`func (o *SearchFiltersResponseSchema) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SearchFiltersResponseSchema) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SearchFiltersResponseSchema) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *SearchFiltersResponseSchema) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SearchFiltersResponseSchema) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SearchFiltersResponseSchema) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


