# ServicesResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**map[string]DeprecatedServiceInfo**](DeprecatedServiceInfo.md) | Map of service id to service details. | 

## Methods

### NewServicesResponseSchema

`func NewServicesResponseSchema(result map[string]DeprecatedServiceInfo, ) *ServicesResponseSchema`

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

`func (o *ServicesResponseSchema) GetResult() map[string]DeprecatedServiceInfo`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ServicesResponseSchema) GetResultOk() (*map[string]DeprecatedServiceInfo, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ServicesResponseSchema) SetResult(v map[string]DeprecatedServiceInfo)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


