# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | 2-letter ISO code of the country. | 
**Name** | **string** | Name of the country. | 
**Services** | [**map[string]Service**](Service.md) | Map of service id to details of the service in this country. | 

## Methods

### NewCountry

`func NewCountry(countryCode string, name string, services map[string]Service, ) *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Country) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Country) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Country) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *Country) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Country) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Country) SetName(v string)`

SetName sets Name field to given value.


### GetServices

`func (o *Country) GetServices() map[string]Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Country) GetServicesOk() (*map[string]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Country) SetServices(v map[string]Service)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


