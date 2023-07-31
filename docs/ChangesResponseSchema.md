# ChangesResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**[]ChangeSet**](ChangeSet.md) | Array of the results matched with the query. | 
**HasMore** | **bool** | Whether there are more results to be loaded. | 
**NextCursor** | **string** | Cursor value to pass to get next set of the results. An empty string if \&quot;hasMore\&quot; is \&quot;false\&quot;. | 

## Methods

### NewChangesResponseSchema

`func NewChangesResponseSchema(result []ChangeSet, hasMore bool, nextCursor string, ) *ChangesResponseSchema`

NewChangesResponseSchema instantiates a new ChangesResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangesResponseSchemaWithDefaults

`func NewChangesResponseSchemaWithDefaults() *ChangesResponseSchema`

NewChangesResponseSchemaWithDefaults instantiates a new ChangesResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ChangesResponseSchema) GetResult() []ChangeSet`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ChangesResponseSchema) GetResultOk() (*[]ChangeSet, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ChangesResponseSchema) SetResult(v []ChangeSet)`

SetResult sets Result field to given value.


### GetHasMore

`func (o *ChangesResponseSchema) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChangesResponseSchema) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChangesResponseSchema) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *ChangesResponseSchema) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ChangesResponseSchema) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ChangesResponseSchema) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


