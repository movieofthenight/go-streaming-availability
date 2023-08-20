/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the ServiceImages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceImages{}

// ServiceImages Images of a service.
type ServiceImages struct {
	// Link to the logo of the service suitable for light themed background.
	LightThemeImage string `json:"lightThemeImage"`
	// Link to the logo of the service suitable for dark themed background.
	DarkThemeImage string `json:"darkThemeImage"`
	// Link to the logo of the service that is all white.
	WhiteImage string `json:"whiteImage"`
}

// NewServiceImages instantiates a new ServiceImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceImages(lightThemeImage string, darkThemeImage string, whiteImage string) *ServiceImages {
	this := ServiceImages{}
	this.LightThemeImage = lightThemeImage
	this.DarkThemeImage = darkThemeImage
	this.WhiteImage = whiteImage
	return &this
}

// NewServiceImagesWithDefaults instantiates a new ServiceImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceImagesWithDefaults() *ServiceImages {
	this := ServiceImages{}
	return &this
}

// GetLightThemeImage returns the LightThemeImage field value
func (o *ServiceImages) GetLightThemeImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LightThemeImage
}

// GetLightThemeImageOk returns a tuple with the LightThemeImage field value
// and a boolean to check if the value has been set.
func (o *ServiceImages) GetLightThemeImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LightThemeImage, true
}

// SetLightThemeImage sets field value
func (o *ServiceImages) SetLightThemeImage(v string) {
	o.LightThemeImage = v
}

// GetDarkThemeImage returns the DarkThemeImage field value
func (o *ServiceImages) GetDarkThemeImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DarkThemeImage
}

// GetDarkThemeImageOk returns a tuple with the DarkThemeImage field value
// and a boolean to check if the value has been set.
func (o *ServiceImages) GetDarkThemeImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkThemeImage, true
}

// SetDarkThemeImage sets field value
func (o *ServiceImages) SetDarkThemeImage(v string) {
	o.DarkThemeImage = v
}

// GetWhiteImage returns the WhiteImage field value
func (o *ServiceImages) GetWhiteImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WhiteImage
}

// GetWhiteImageOk returns a tuple with the WhiteImage field value
// and a boolean to check if the value has been set.
func (o *ServiceImages) GetWhiteImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhiteImage, true
}

// SetWhiteImage sets field value
func (o *ServiceImages) SetWhiteImage(v string) {
	o.WhiteImage = v
}

func (o ServiceImages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceImages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lightThemeImage"] = o.LightThemeImage
	toSerialize["darkThemeImage"] = o.DarkThemeImage
	toSerialize["whiteImage"] = o.WhiteImage
	return toSerialize, nil
}

type NullableServiceImages struct {
	value *ServiceImages
	isSet bool
}

func (v NullableServiceImages) Get() *ServiceImages {
	return v.value
}

func (v *NullableServiceImages) Set(val *ServiceImages) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceImages) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceImages(val *ServiceImages) *NullableServiceImages {
	return &NullableServiceImages{value: val, isSet: true}
}

func (v NullableServiceImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

