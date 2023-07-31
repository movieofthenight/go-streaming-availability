# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Numerical amount of the price. | 
**Currency** | **string** | ISO 4217 alphabetic code of the currency. | 
**Formatted** | **string** | Formatted price, including both the amount and the currency. | 

## Methods

### NewPrice

`func NewPrice(amount string, currency string, formatted string, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Price) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Price) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Price) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFormatted

`func (o *Price) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *Price) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *Price) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


