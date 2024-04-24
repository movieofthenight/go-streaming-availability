# ShowImageSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerticalPoster** | [**VerticalImage**](VerticalImage.md) | Vertical poster of the show. | 
**HorizontalPoster** | [**HorizontalImage**](HorizontalImage.md) | Horizontal poster of the show. | 
**VerticalBackdrop** | Pointer to [**VerticalImage**](VerticalImage.md) | Vertical backdrop of the show. | [optional] 
**HorizontalBackdrop** | Pointer to [**HorizontalImage**](HorizontalImage.md) | Horizontal backdrop of the show. | [optional] 

## Methods

### NewShowImageSet

`func NewShowImageSet(verticalPoster VerticalImage, horizontalPoster HorizontalImage, ) *ShowImageSet`

NewShowImageSet instantiates a new ShowImageSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowImageSetWithDefaults

`func NewShowImageSetWithDefaults() *ShowImageSet`

NewShowImageSetWithDefaults instantiates a new ShowImageSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerticalPoster

`func (o *ShowImageSet) GetVerticalPoster() VerticalImage`

GetVerticalPoster returns the VerticalPoster field if non-nil, zero value otherwise.

### GetVerticalPosterOk

`func (o *ShowImageSet) GetVerticalPosterOk() (*VerticalImage, bool)`

GetVerticalPosterOk returns a tuple with the VerticalPoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalPoster

`func (o *ShowImageSet) SetVerticalPoster(v VerticalImage)`

SetVerticalPoster sets VerticalPoster field to given value.


### GetHorizontalPoster

`func (o *ShowImageSet) GetHorizontalPoster() HorizontalImage`

GetHorizontalPoster returns the HorizontalPoster field if non-nil, zero value otherwise.

### GetHorizontalPosterOk

`func (o *ShowImageSet) GetHorizontalPosterOk() (*HorizontalImage, bool)`

GetHorizontalPosterOk returns a tuple with the HorizontalPoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalPoster

`func (o *ShowImageSet) SetHorizontalPoster(v HorizontalImage)`

SetHorizontalPoster sets HorizontalPoster field to given value.


### GetVerticalBackdrop

`func (o *ShowImageSet) GetVerticalBackdrop() VerticalImage`

GetVerticalBackdrop returns the VerticalBackdrop field if non-nil, zero value otherwise.

### GetVerticalBackdropOk

`func (o *ShowImageSet) GetVerticalBackdropOk() (*VerticalImage, bool)`

GetVerticalBackdropOk returns a tuple with the VerticalBackdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalBackdrop

`func (o *ShowImageSet) SetVerticalBackdrop(v VerticalImage)`

SetVerticalBackdrop sets VerticalBackdrop field to given value.

### HasVerticalBackdrop

`func (o *ShowImageSet) HasVerticalBackdrop() bool`

HasVerticalBackdrop returns a boolean if a field has been set.

### GetHorizontalBackdrop

`func (o *ShowImageSet) GetHorizontalBackdrop() HorizontalImage`

GetHorizontalBackdrop returns the HorizontalBackdrop field if non-nil, zero value otherwise.

### GetHorizontalBackdropOk

`func (o *ShowImageSet) GetHorizontalBackdropOk() (*HorizontalImage, bool)`

GetHorizontalBackdropOk returns a tuple with the HorizontalBackdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalBackdrop

`func (o *ShowImageSet) SetHorizontalBackdrop(v HorizontalImage)`

SetHorizontalBackdrop sets HorizontalBackdrop field to given value.

### HasHorizontalBackdrop

`func (o *ShowImageSet) HasHorizontalBackdrop() bool`

HasHorizontalBackdrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


