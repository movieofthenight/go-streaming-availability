# GenresResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **map[string]string** | Map of genre id to genre name. | 

## Methods

### NewGenresResponseSchema

`func NewGenresResponseSchema(result map[string]string, ) *GenresResponseSchema`

NewGenresResponseSchema instantiates a new GenresResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenresResponseSchemaWithDefaults

`func NewGenresResponseSchemaWithDefaults() *GenresResponseSchema`

NewGenresResponseSchemaWithDefaults instantiates a new GenresResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GenresResponseSchema) GetResult() map[string]string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GenresResponseSchema) GetResultOk() (*map[string]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GenresResponseSchema) SetResult(v map[string]string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


