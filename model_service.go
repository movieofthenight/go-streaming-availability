/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 3.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service Details of a service.
type Service struct {
	// Id of the service.
	Id string `json:"id"`
	// Name of the service.
	Name string `json:"name"`
	// Link to the homepage of the service.
	HomePage string `json:"homePage"`
	// Associated theme color hex code of the service.
	ThemeColorCode string `json:"themeColorCode"`
	Images ImageSet `json:"images"`
	SupportedStreamingTypes SupportedStreamingTypes `json:"supportedStreamingTypes"`
	// Map of id to details of the addons supported by the service in this country.
	Addons map[string]Addon `json:"addons"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id string, name string, homePage string, themeColorCode string, images ImageSet, supportedStreamingTypes SupportedStreamingTypes, addons map[string]Addon) *Service {
	this := Service{}
	this.Id = id
	this.Name = name
	this.HomePage = homePage
	this.ThemeColorCode = themeColorCode
	this.Images = images
	this.SupportedStreamingTypes = supportedStreamingTypes
	this.Addons = addons
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetHomePage returns the HomePage field value
func (o *Service) GetHomePage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value
// and a boolean to check if the value has been set.
func (o *Service) GetHomePageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomePage, true
}

// SetHomePage sets field value
func (o *Service) SetHomePage(v string) {
	o.HomePage = v
}

// GetThemeColorCode returns the ThemeColorCode field value
func (o *Service) GetThemeColorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeColorCode
}

// GetThemeColorCodeOk returns a tuple with the ThemeColorCode field value
// and a boolean to check if the value has been set.
func (o *Service) GetThemeColorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeColorCode, true
}

// SetThemeColorCode sets field value
func (o *Service) SetThemeColorCode(v string) {
	o.ThemeColorCode = v
}

// GetImages returns the Images field value
func (o *Service) GetImages() ImageSet {
	if o == nil {
		var ret ImageSet
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *Service) GetImagesOk() (*ImageSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Images, true
}

// SetImages sets field value
func (o *Service) SetImages(v ImageSet) {
	o.Images = v
}

// GetSupportedStreamingTypes returns the SupportedStreamingTypes field value
func (o *Service) GetSupportedStreamingTypes() SupportedStreamingTypes {
	if o == nil {
		var ret SupportedStreamingTypes
		return ret
	}

	return o.SupportedStreamingTypes
}

// GetSupportedStreamingTypesOk returns a tuple with the SupportedStreamingTypes field value
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedStreamingTypesOk() (*SupportedStreamingTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedStreamingTypes, true
}

// SetSupportedStreamingTypes sets field value
func (o *Service) SetSupportedStreamingTypes(v SupportedStreamingTypes) {
	o.SupportedStreamingTypes = v
}

// GetAddons returns the Addons field value
func (o *Service) GetAddons() map[string]Addon {
	if o == nil {
		var ret map[string]Addon
		return ret
	}

	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value
// and a boolean to check if the value has been set.
func (o *Service) GetAddonsOk() (*map[string]Addon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addons, true
}

// SetAddons sets field value
func (o *Service) SetAddons(v map[string]Addon) {
	o.Addons = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["homePage"] = o.HomePage
	toSerialize["themeColorCode"] = o.ThemeColorCode
	toSerialize["images"] = o.Images
	toSerialize["supportedStreamingTypes"] = o.SupportedStreamingTypes
	toSerialize["addons"] = o.Addons
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


