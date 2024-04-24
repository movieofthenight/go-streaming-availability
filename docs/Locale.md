# Locale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** | [ISO 639-2](https://en.wikipedia.org/wiki/ISO_639-2) code of the associated language with the locale. | 
**Region** | Pointer to **string** | [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) code of the country, or [UN M49](https://en.wikipedia.org/wiki/UN_M49) code of the area associated with the locale.  | [optional] 

## Methods

### NewLocale

`func NewLocale(language string, ) *Locale`

NewLocale instantiates a new Locale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocaleWithDefaults

`func NewLocaleWithDefaults() *Locale`

NewLocaleWithDefaults instantiates a new Locale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *Locale) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Locale) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Locale) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetRegion

`func (o *Locale) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Locale) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Locale) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Locale) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


