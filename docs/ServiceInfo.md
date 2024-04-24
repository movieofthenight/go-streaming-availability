# ServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Name** | **string** | Name of the service. | 
**HomePage** | **string** | Link to the homepage of the service. | 
**ThemeColorCode** | **string** | Associated theme color hex code of the service. | 
**ImageSet** | [**ServiceImageSet**](ServiceImageSet.md) | Image set of the service. | 

## Methods

### NewServiceInfo

`func NewServiceInfo(id string, name string, homePage string, themeColorCode string, imageSet ServiceImageSet, ) *ServiceInfo`

NewServiceInfo instantiates a new ServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInfoWithDefaults

`func NewServiceInfoWithDefaults() *ServiceInfo`

NewServiceInfoWithDefaults instantiates a new ServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServiceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceInfo) SetName(v string)`

SetName sets Name field to given value.


### GetHomePage

`func (o *ServiceInfo) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *ServiceInfo) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *ServiceInfo) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.


### GetThemeColorCode

`func (o *ServiceInfo) GetThemeColorCode() string`

GetThemeColorCode returns the ThemeColorCode field if non-nil, zero value otherwise.

### GetThemeColorCodeOk

`func (o *ServiceInfo) GetThemeColorCodeOk() (*string, bool)`

GetThemeColorCodeOk returns a tuple with the ThemeColorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColorCode

`func (o *ServiceInfo) SetThemeColorCode(v string)`

SetThemeColorCode sets ThemeColorCode field to given value.


### GetImageSet

`func (o *ServiceInfo) GetImageSet() ServiceImageSet`

GetImageSet returns the ImageSet field if non-nil, zero value otherwise.

### GetImageSetOk

`func (o *ServiceInfo) GetImageSetOk() (*ServiceImageSet, bool)`

GetImageSetOk returns a tuple with the ImageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSet

`func (o *ServiceInfo) SetImageSet(v ServiceImageSet)`

SetImageSet sets ImageSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


