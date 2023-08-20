# DeprecatedServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Countries** | [**map[string]ServiceCountryInfo**](ServiceCountryInfo.md) | Map of 2-letter country ISO code to details of the service in that country. | 

## Methods

### NewDeprecatedServiceInfo

`func NewDeprecatedServiceInfo(id string, countries map[string]ServiceCountryInfo, ) *DeprecatedServiceInfo`

NewDeprecatedServiceInfo instantiates a new DeprecatedServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedServiceInfoWithDefaults

`func NewDeprecatedServiceInfoWithDefaults() *DeprecatedServiceInfo`

NewDeprecatedServiceInfoWithDefaults instantiates a new DeprecatedServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedServiceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedServiceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedServiceInfo) SetId(v string)`

SetId sets Id field to given value.


### GetCountries

`func (o *DeprecatedServiceInfo) GetCountries() map[string]ServiceCountryInfo`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *DeprecatedServiceInfo) GetCountriesOk() (*map[string]ServiceCountryInfo, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *DeprecatedServiceInfo) SetCountries(v map[string]ServiceCountryInfo)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


