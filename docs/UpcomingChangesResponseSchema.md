# UpcomingChangesResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**[]UpcomingChangeSet**](UpcomingChangeSet.md) | Array of the results matched with the query. | 
**HasMore** | **bool** | Whether there are more results to be loaded. | 
**NextCursor** | **string** | Cursor value to pass to get next set of the results. An empty string if \&quot;hasMore\&quot; is \&quot;false\&quot;. | 

## Methods

### NewUpcomingChangesResponseSchema

`func NewUpcomingChangesResponseSchema(result []UpcomingChangeSet, hasMore bool, nextCursor string, ) *UpcomingChangesResponseSchema`

NewUpcomingChangesResponseSchema instantiates a new UpcomingChangesResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpcomingChangesResponseSchemaWithDefaults

`func NewUpcomingChangesResponseSchemaWithDefaults() *UpcomingChangesResponseSchema`

NewUpcomingChangesResponseSchemaWithDefaults instantiates a new UpcomingChangesResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpcomingChangesResponseSchema) GetResult() []UpcomingChangeSet`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpcomingChangesResponseSchema) GetResultOk() (*[]UpcomingChangeSet, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpcomingChangesResponseSchema) SetResult(v []UpcomingChangeSet)`

SetResult sets Result field to given value.


### GetHasMore

`func (o *UpcomingChangesResponseSchema) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UpcomingChangesResponseSchema) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UpcomingChangesResponseSchema) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *UpcomingChangesResponseSchema) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *UpcomingChangesResponseSchema) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *UpcomingChangesResponseSchema) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


