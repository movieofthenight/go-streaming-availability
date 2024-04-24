/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Addon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Addon{}

// Addon Details of an addon.
type Addon struct {
	// Id of the addon.
	Id string `json:"id"`
	// Name of the addon.
	Name string `json:"name"`
	// Link to the homepage of the addon.
	HomePage string `json:"homePage"`
	// Associated theme color hex code of the addon.
	ThemeColorCode string `json:"themeColorCode"`
	// Image set of the addon.
	ImageSet ServiceImageSet `json:"imageSet"`
}

type _Addon Addon

// NewAddon instantiates a new Addon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddon(id string, name string, homePage string, themeColorCode string, imageSet ServiceImageSet) *Addon {
	this := Addon{}
	this.Id = id
	this.Name = name
	this.HomePage = homePage
	this.ThemeColorCode = themeColorCode
	this.ImageSet = imageSet
	return &this
}

// NewAddonWithDefaults instantiates a new Addon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonWithDefaults() *Addon {
	this := Addon{}
	return &this
}

// GetId returns the Id field value
func (o *Addon) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Addon) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Addon) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Addon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Addon) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Addon) SetName(v string) {
	o.Name = v
}

// GetHomePage returns the HomePage field value
func (o *Addon) GetHomePage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value
// and a boolean to check if the value has been set.
func (o *Addon) GetHomePageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomePage, true
}

// SetHomePage sets field value
func (o *Addon) SetHomePage(v string) {
	o.HomePage = v
}

// GetThemeColorCode returns the ThemeColorCode field value
func (o *Addon) GetThemeColorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeColorCode
}

// GetThemeColorCodeOk returns a tuple with the ThemeColorCode field value
// and a boolean to check if the value has been set.
func (o *Addon) GetThemeColorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeColorCode, true
}

// SetThemeColorCode sets field value
func (o *Addon) SetThemeColorCode(v string) {
	o.ThemeColorCode = v
}

// GetImageSet returns the ImageSet field value
func (o *Addon) GetImageSet() ServiceImageSet {
	if o == nil {
		var ret ServiceImageSet
		return ret
	}

	return o.ImageSet
}

// GetImageSetOk returns a tuple with the ImageSet field value
// and a boolean to check if the value has been set.
func (o *Addon) GetImageSetOk() (*ServiceImageSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageSet, true
}

// SetImageSet sets field value
func (o *Addon) SetImageSet(v ServiceImageSet) {
	o.ImageSet = v
}

func (o Addon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Addon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["homePage"] = o.HomePage
	toSerialize["themeColorCode"] = o.ThemeColorCode
	toSerialize["imageSet"] = o.ImageSet
	return toSerialize, nil
}

func (o *Addon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"homePage",
		"themeColorCode",
		"imageSet",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddon := _Addon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddon)

	if err != nil {
		return err
	}

	*o = Addon(varAddon)

	return err
}

type NullableAddon struct {
	value *Addon
	isSet bool
}

func (v NullableAddon) Get() *Addon {
	return v.value
}

func (v *NullableAddon) Set(val *Addon) {
	v.value = val
	v.isSet = true
}

func (v NullableAddon) IsSet() bool {
	return v.isSet
}

func (v *NullableAddon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddon(val *Addon) *NullableAddon {
	return &NullableAddon{value: val, isSet: true}
}

func (v NullableAddon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


