# ServicesResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**map[string]Service**](Service.md) | Map of service id to service details. | [optional] 

## Methods

### NewServicesResponseSchema

`func NewServicesResponseSchema() *ServicesResponseSchema`

NewServicesResponseSchema instantiates a new ServicesResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesResponseSchemaWithDefaults

`func NewServicesResponseSchemaWithDefaults() *ServicesResponseSchema`

NewServicesResponseSchemaWithDefaults instantiates a new ServicesResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ServicesResponseSchema) GetResult() map[string]Service`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ServicesResponseSchema) GetResultOk() (*map[string]Service, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ServicesResponseSchema) SetResult(v map[string]Service)`

SetResult sets Result field to given value.

### HasResult

`func (o *ServicesResponseSchema) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


