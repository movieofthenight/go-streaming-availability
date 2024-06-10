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

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price Price of the renting or buying the item.  A movie and an episode that is available to buy or rent will always have a price.  A series or a season that is available to buy or rent may have a price or not. If the price is available, that means the entire series or the season can be bought or rented as a whole for the given price. If the price is null, that means sub-items of the series or the season are available to buy or rent, but it is not possible to buy or rent the entire series or the season as a whole at once. In this case, the price of the sub-items can be found in the episodes or seasons array. 
type Price struct {
	// Numerical amount of the price.
	Amount string `json:"amount"`
	// [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) alphabetic code of the currency.
	Currency string `json:"currency"`
	// Formatted price, including both the amount and the currency.
	Formatted string `json:"formatted"`
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(amount string, currency string, formatted string) *Price {
	this := Price{}
	this.Amount = amount
	this.Currency = currency
	this.Formatted = formatted
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Price) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Price) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Price) SetAmount(v string) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *Price) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Price) SetCurrency(v string) {
	o.Currency = v
}

// GetFormatted returns the Formatted field value
func (o *Price) GetFormatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value
// and a boolean to check if the value has been set.
func (o *Price) GetFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formatted, true
}

// SetFormatted sets field value
func (o *Price) SetFormatted(v string) {
	o.Formatted = v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["formatted"] = o.Formatted
	return toSerialize, nil
}

func (o *Price) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"formatted",
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

	varPrice := _Price{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrice)

	if err != nil {
		return err
	}

	*o = Price(varPrice)

	return err
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


