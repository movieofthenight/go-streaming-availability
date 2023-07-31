# SearchTitleResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**[]Show**](Show.md) | Array of the shows matched with the title. | 

## Methods

### NewSearchTitleResponseSchema

`func NewSearchTitleResponseSchema(result []Show, ) *SearchTitleResponseSchema`

NewSearchTitleResponseSchema instantiates a new SearchTitleResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTitleResponseSchemaWithDefaults

`func NewSearchTitleResponseSchemaWithDefaults() *SearchTitleResponseSchema`

NewSearchTitleResponseSchemaWithDefaults instantiates a new SearchTitleResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SearchTitleResponseSchema) GetResult() []Show`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchTitleResponseSchema) GetResultOk() (*[]Show, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchTitleResponseSchema) SetResult(v []Show)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


