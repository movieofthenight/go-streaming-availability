# Addon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the addon. | 
**DisplayName** | **string** | Name of the addon. | 
**HomePage** | **string** | Link to the homepage of the addon. | 
**ThemeColorCode** | **string** | Associated theme color hex code of the addon. | 
**Image** | **string** | Link to the logo of the addon. | 

## Methods

### NewAddon

`func NewAddon(id string, displayName string, homePage string, themeColorCode string, image string, ) *Addon`

NewAddon instantiates a new Addon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonWithDefaults

`func NewAddonWithDefaults() *Addon`

NewAddonWithDefaults instantiates a new Addon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Addon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Addon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Addon) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *Addon) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Addon) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Addon) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHomePage

`func (o *Addon) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *Addon) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *Addon) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.


### GetThemeColorCode

`func (o *Addon) GetThemeColorCode() string`

GetThemeColorCode returns the ThemeColorCode field if non-nil, zero value otherwise.

### GetThemeColorCodeOk

`func (o *Addon) GetThemeColorCodeOk() (*string, bool)`

GetThemeColorCodeOk returns a tuple with the ThemeColorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColorCode

`func (o *Addon) SetThemeColorCode(v string)`

SetThemeColorCode sets ThemeColorCode field to given value.


### GetImage

`func (o *Addon) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Addon) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Addon) SetImage(v string)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


