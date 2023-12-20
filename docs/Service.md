# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Name** | **string** | Name of the service. | 
**HomePage** | **string** | Link to the homepage of the service. | 
**ThemeColorCode** | **string** | Associated theme color hex code of the service. | 
**Images** | [**ImageSet**](ImageSet.md) |  | 
**SupportedStreamingTypes** | [**SupportedStreamingTypes**](SupportedStreamingTypes.md) |  | 
**Addons** | [**map[string]Addon**](Addon.md) | Map of id to details of the addons supported by the service in this country. | 

## Methods

### NewService

`func NewService(id string, name string, homePage string, themeColorCode string, images ImageSet, supportedStreamingTypes SupportedStreamingTypes, addons map[string]Addon, ) *Service`

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


### GetImages

`func (o *Service) GetImages() ImageSet`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Service) GetImagesOk() (*ImageSet, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Service) SetImages(v ImageSet)`

SetImages sets Images field to given value.


### GetSupportedStreamingTypes

`func (o *Service) GetSupportedStreamingTypes() SupportedStreamingTypes`

GetSupportedStreamingTypes returns the SupportedStreamingTypes field if non-nil, zero value otherwise.

### GetSupportedStreamingTypesOk

`func (o *Service) GetSupportedStreamingTypesOk() (*SupportedStreamingTypes, bool)`

GetSupportedStreamingTypesOk returns a tuple with the SupportedStreamingTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedStreamingTypes

`func (o *Service) SetSupportedStreamingTypes(v SupportedStreamingTypes)`

SetSupportedStreamingTypes sets SupportedStreamingTypes field to given value.


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


