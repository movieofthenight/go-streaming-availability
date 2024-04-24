# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Name** | **string** | Name of the service. | 
**HomePage** | **string** | Link to the homepage of the service. | 
**ThemeColorCode** | **string** | Associated theme color hex code of the service. | 
**ImageSet** | [**ServiceImageSet**](ServiceImageSet.md) | Image set of the service. | 
**StreamingOptionTypes** | [**StreamingOptionTypes**](StreamingOptionTypes.md) |  | 
**Addons** |  | Map of the supported addons by their ids. | 

## Methods

### NewService

`func NewService(id string, name string, homePage string, themeColorCode string, imageSet ServiceImageSet, streamingOptionTypes StreamingOptionTypes, addons map[string]Addon, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetHomePage

`func (o *Service) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *Service) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *Service) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.


### GetThemeColorCode

`func (o *Service) GetThemeColorCode() string`

GetThemeColorCode returns the ThemeColorCode field if non-nil, zero value otherwise.

### GetThemeColorCodeOk

`func (o *Service) GetThemeColorCodeOk() (*string, bool)`

GetThemeColorCodeOk returns a tuple with the ThemeColorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColorCode

`func (o *Service) SetThemeColorCode(v string)`

SetThemeColorCode sets ThemeColorCode field to given value.


### GetImageSet

`func (o *Service) GetImageSet() ServiceImageSet`

GetImageSet returns the ImageSet field if non-nil, zero value otherwise.

### GetImageSetOk

`func (o *Service) GetImageSetOk() (*ServiceImageSet, bool)`

GetImageSetOk returns a tuple with the ImageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSet

`func (o *Service) SetImageSet(v ServiceImageSet)`

SetImageSet sets ImageSet field to given value.


### GetStreamingOptionTypes

`func (o *Service) GetStreamingOptionTypes() StreamingOptionTypes`

GetStreamingOptionTypes returns the StreamingOptionTypes field if non-nil, zero value otherwise.

### GetStreamingOptionTypesOk

`func (o *Service) GetStreamingOptionTypesOk() (*StreamingOptionTypes, bool)`

GetStreamingOptionTypesOk returns a tuple with the StreamingOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOptionTypes

`func (o *Service) SetStreamingOptionTypes(v StreamingOptionTypes)`

SetStreamingOptionTypes sets StreamingOptionTypes field to given value.


### GetAddons

`func (o *Service) GetAddons() map[string]Addon`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *Service) GetAddonsOk() (*map[string]Addon, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *Service) SetAddons(v map[string]Addon)`

SetAddons sets Addons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


