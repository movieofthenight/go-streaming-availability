# Subtitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosedCaptions** | **bool** | Whether closed captions are available for the subtitle. | 
**Locale** | [**Locale**](Locale.md) |  | 

## Methods

### NewSubtitle

`func NewSubtitle(closedCaptions bool, locale Locale, ) *Subtitle`

NewSubtitle instantiates a new Subtitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtitleWithDefaults

`func NewSubtitleWithDefaults() *Subtitle`

NewSubtitleWithDefaults instantiates a new Subtitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosedCaptions

`func (o *Subtitle) GetClosedCaptions() bool`

GetClosedCaptions returns the ClosedCaptions field if non-nil, zero value otherwise.

### GetClosedCaptionsOk

`func (o *Subtitle) GetClosedCaptionsOk() (*bool, bool)`

GetClosedCaptionsOk returns a tuple with the ClosedCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedCaptions

`func (o *Subtitle) SetClosedCaptions(v bool)`

SetClosedCaptions sets ClosedCaptions field to given value.


### GetLocale

`func (o *Subtitle) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Subtitle) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Subtitle) SetLocale(v Locale)`

SetLocale sets Locale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


