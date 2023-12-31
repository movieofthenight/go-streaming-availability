/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 59 countries!

API version: 3.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the SupportedStreamingTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedStreamingTypes{}

// SupportedStreamingTypes Details about the supported streaming types for a service-country combo.
type SupportedStreamingTypes struct {
	// Whether there are addons/channels available.
	Addon bool `json:"addon"`
	// Whether buying shows is supported.
	Buy bool `json:"buy"`
	// Whether renting shows is supported.
	Rent bool `json:"rent"`
	// Whether there are free shows available.
	Free bool `json:"free"`
	// Whether signing up for a subscription plan is available.
	Subscription bool `json:"subscription"`
}

// NewSupportedStreamingTypes instantiates a new SupportedStreamingTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedStreamingTypes(addon bool, buy bool, rent bool, free bool, subscription bool) *SupportedStreamingTypes {
	this := SupportedStreamingTypes{}
	this.Addon = addon
	this.Buy = buy
	this.Rent = rent
	this.Free = free
	this.Subscription = subscription
	return &this
}

// NewSupportedStreamingTypesWithDefaults instantiates a new SupportedStreamingTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedStreamingTypesWithDefaults() *SupportedStreamingTypes {
	this := SupportedStreamingTypes{}
	return &this
}

// GetAddon returns the Addon field value
func (o *SupportedStreamingTypes) GetAddon() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Addon
}

// GetAddonOk returns a tuple with the Addon field value
// and a boolean to check if the value has been set.
func (o *SupportedStreamingTypes) GetAddonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addon, true
}

// SetAddon sets field value
func (o *SupportedStreamingTypes) SetAddon(v bool) {
	o.Addon = v
}

// GetBuy returns the Buy field value
func (o *SupportedStreamingTypes) GetBuy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Buy
}

// GetBuyOk returns a tuple with the Buy field value
// and a boolean to check if the value has been set.
func (o *SupportedStreamingTypes) GetBuyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buy, true
}

// SetBuy sets field value
func (o *SupportedStreamingTypes) SetBuy(v bool) {
	o.Buy = v
}

// GetRent returns the Rent field value
func (o *SupportedStreamingTypes) GetRent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rent
}

// GetRentOk returns a tuple with the Rent field value
// and a boolean to check if the value has been set.
func (o *SupportedStreamingTypes) GetRentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rent, true
}

// SetRent sets field value
func (o *SupportedStreamingTypes) SetRent(v bool) {
	o.Rent = v
}

// GetFree returns the Free field value
func (o *SupportedStreamingTypes) GetFree() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *SupportedStreamingTypes) GetFreeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *SupportedStreamingTypes) SetFree(v bool) {
	o.Free = v
}

// GetSubscription returns the Subscription field value
func (o *SupportedStreamingTypes) GetSubscription() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SupportedStreamingTypes) GetSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SupportedStreamingTypes) SetSubscription(v bool) {
	o.Subscription = v
}

func (o SupportedStreamingTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedStreamingTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addon"] = o.Addon
	toSerialize["buy"] = o.Buy
	toSerialize["rent"] = o.Rent
	toSerialize["free"] = o.Free
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

type NullableSupportedStreamingTypes struct {
	value *SupportedStreamingTypes
	isSet bool
}

func (v NullableSupportedStreamingTypes) Get() *SupportedStreamingTypes {
	return v.value
}

func (v *NullableSupportedStreamingTypes) Set(val *SupportedStreamingTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedStreamingTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedStreamingTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedStreamingTypes(val *SupportedStreamingTypes) *NullableSupportedStreamingTypes {
	return &NullableSupportedStreamingTypes{value: val, isSet: true}
}

func (v NullableSupportedStreamingTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedStreamingTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


