# ServiceCountryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedStreamingTypes** | [**SupportedStreamingTypes**](SupportedStreamingTypes.md) |  | 
**Addons** | [**map[string]Addon**](Addon.md) | Map of id to details of the addons supported by the service in this country. | 

## Methods

### NewServiceCountryInfo

`func NewServiceCountryInfo(supportedStreamingTypes SupportedStreamingTypes, addons map[string]Addon, ) *ServiceCountryInfo`

NewServiceCountryInfo instantiates a new ServiceCountryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCountryInfoWithDefaults

`func NewServiceCountryInfoWithDefaults() *ServiceCountryInfo`

NewServiceCountryInfoWithDefaults instantiates a new ServiceCountryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedStreamingTypes

`func (o *ServiceCountryInfo) GetSupportedStreamingTypes() SupportedStreamingTypes`

GetSupportedStreamingTypes returns the SupportedStreamingTypes field if non-nil, zero value otherwise.

### GetSupportedStreamingTypesOk

`func (o *ServiceCountryInfo) GetSupportedStreamingTypesOk() (*SupportedStreamingTypes, bool)`

GetSupportedStreamingTypesOk returns a tuple with the SupportedStreamingTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedStreamingTypes

`func (o *ServiceCountryInfo) SetSupportedStreamingTypes(v SupportedStreamingTypes)`

SetSupportedStreamingTypes sets SupportedStreamingTypes field to given value.


### GetAddons

`func (o *ServiceCountryInfo) GetAddons() map[string]Addon`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *ServiceCountryInfo) GetAddonsOk() (*map[string]Addon, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *ServiceCountryInfo) SetAddons(v map[string]Addon)`

SetAddons sets Addons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


