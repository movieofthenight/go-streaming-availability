/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 60 countries!

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ShowImageSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowImageSet{}

// ShowImageSet Image set of a show.
type ShowImageSet struct {
	// Vertical poster of the show.
	VerticalPoster VerticalImage `json:"verticalPoster"`
	// Horizontal poster of the show.
	HorizontalPoster HorizontalImage `json:"horizontalPoster"`
	// Vertical backdrop of the show.
	VerticalBackdrop *VerticalImage `json:"verticalBackdrop,omitempty"`
	// Horizontal backdrop of the show.
	HorizontalBackdrop *HorizontalImage `json:"horizontalBackdrop,omitempty"`
}

type _ShowImageSet ShowImageSet

// NewShowImageSet instantiates a new ShowImageSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowImageSet(verticalPoster VerticalImage, horizontalPoster HorizontalImage) *ShowImageSet {
	this := ShowImageSet{}
	this.VerticalPoster = verticalPoster
	this.HorizontalPoster = horizontalPoster
	return &this
}

// NewShowImageSetWithDefaults instantiates a new ShowImageSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowImageSetWithDefaults() *ShowImageSet {
	this := ShowImageSet{}
	return &this
}

// GetVerticalPoster returns the VerticalPoster field value
func (o *ShowImageSet) GetVerticalPoster() VerticalImage {
	if o == nil {
		var ret VerticalImage
		return ret
	}

	return o.VerticalPoster
}

// GetVerticalPosterOk returns a tuple with the VerticalPoster field value
// and a boolean to check if the value has been set.
func (o *ShowImageSet) GetVerticalPosterOk() (*VerticalImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerticalPoster, true
}

// SetVerticalPoster sets field value
func (o *ShowImageSet) SetVerticalPoster(v VerticalImage) {
	o.VerticalPoster = v
}

// GetHorizontalPoster returns the HorizontalPoster field value
func (o *ShowImageSet) GetHorizontalPoster() HorizontalImage {
	if o == nil {
		var ret HorizontalImage
		return ret
	}

	return o.HorizontalPoster
}

// GetHorizontalPosterOk returns a tuple with the HorizontalPoster field value
// and a boolean to check if the value has been set.
func (o *ShowImageSet) GetHorizontalPosterOk() (*HorizontalImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HorizontalPoster, true
}

// SetHorizontalPoster sets field value
func (o *ShowImageSet) SetHorizontalPoster(v HorizontalImage) {
	o.HorizontalPoster = v
}

// GetVerticalBackdrop returns the VerticalBackdrop field value if set, zero value otherwise.
func (o *ShowImageSet) GetVerticalBackdrop() VerticalImage {
	if o == nil || IsNil(o.VerticalBackdrop) {
		var ret VerticalImage
		return ret
	}
	return *o.VerticalBackdrop
}

// GetVerticalBackdropOk returns a tuple with the VerticalBackdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowImageSet) GetVerticalBackdropOk() (*VerticalImage, bool) {
	if o == nil || IsNil(o.VerticalBackdrop) {
		return nil, false
	}
	return o.VerticalBackdrop, true
}

// HasVerticalBackdrop returns a boolean if a field has been set.
func (o *ShowImageSet) HasVerticalBackdrop() bool {
	if o != nil && !IsNil(o.VerticalBackdrop) {
		return true
	}

	return false
}

// SetVerticalBackdrop gets a reference to the given VerticalImage and assigns it to the VerticalBackdrop field.
func (o *ShowImageSet) SetVerticalBackdrop(v VerticalImage) {
	o.VerticalBackdrop = &v
}

// GetHorizontalBackdrop returns the HorizontalBackdrop field value if set, zero value otherwise.
func (o *ShowImageSet) GetHorizontalBackdrop() HorizontalImage {
	if o == nil || IsNil(o.HorizontalBackdrop) {
		var ret HorizontalImage
		return ret
	}
	return *o.HorizontalBackdrop
}

// GetHorizontalBackdropOk returns a tuple with the HorizontalBackdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowImageSet) GetHorizontalBackdropOk() (*HorizontalImage, bool) {
	if o == nil || IsNil(o.HorizontalBackdrop) {
		return nil, false
	}
	return o.HorizontalBackdrop, true
}

// HasHorizontalBackdrop returns a boolean if a field has been set.
func (o *ShowImageSet) HasHorizontalBackdrop() bool {
	if o != nil && !IsNil(o.HorizontalBackdrop) {
		return true
	}

	return false
}

// SetHorizontalBackdrop gets a reference to the given HorizontalImage and assigns it to the HorizontalBackdrop field.
func (o *ShowImageSet) SetHorizontalBackdrop(v HorizontalImage) {
	o.HorizontalBackdrop = &v
}

func (o ShowImageSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowImageSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verticalPoster"] = o.VerticalPoster
	toSerialize["horizontalPoster"] = o.HorizontalPoster
	if !IsNil(o.VerticalBackdrop) {
		toSerialize["verticalBackdrop"] = o.VerticalBackdrop
	}
	if !IsNil(o.HorizontalBackdrop) {
		toSerialize["horizontalBackdrop"] = o.HorizontalBackdrop
	}
	return toSerialize, nil
}

func (o *ShowImageSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verticalPoster",
		"horizontalPoster",
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

	varShowImageSet := _ShowImageSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShowImageSet)

	if err != nil {
		return err
	}

	*o = ShowImageSet(varShowImageSet)

	return err
}

type NullableShowImageSet struct {
	value *ShowImageSet
	isSet bool
}

func (v NullableShowImageSet) Get() *ShowImageSet {
	return v.value
}

func (v *NullableShowImageSet) Set(val *ShowImageSet) {
	v.value = val
	v.isSet = true
}

func (v NullableShowImageSet) IsSet() bool {
	return v.isSet
}

func (v *NullableShowImageSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowImageSet(val *ShowImageSet) *NullableShowImageSet {
	return &NullableShowImageSet{value: val, isSet: true}
}

func (v NullableShowImageSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowImageSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


