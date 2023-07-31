/*
Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
)

// checks if the StreamingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamingOption{}

// StreamingOption A streaming option.
type StreamingOption struct {
	// Id of the streaming service.
	Service string `json:"service"`
	StreamingType StreamingType `json:"streamingType"`
	// Addon id, if the \"streamingType\" is \"addon\". Otherwise omitted.
	Addon *string `json:"addon,omitempty"`
	// Deep link to the streaming option's page in the streaming service. Guaranteed to be populated.
	Link string `json:"link"`
	// Deep link to the video associated with the streaming option. Omitted if there's no direct link to the video. Might have the same value as \"link\". 
	VideoLink *string `json:"videoLink,omitempty"`
	// Maximum video quality of the streaming option. Omitted if the quality is unknown.
	Quality *string `json:"quality,omitempty"`
	// Array of the available audios.
	Audios []Locale `json:"audios"`
	// Array of the available subtitles.
	Subtitles []Subtitle `json:"subtitles"`
	Price *Price `json:"price,omitempty"`
	// [Unix Time Stamp](https://www.unixtimestamp.com/) of the date that this streaming option is expiring. In other words, last day to watch. A value of 1 means the streaming option is expiring soon, but there's no specific date info is found. Omitted if there's no known expiry date. 
	Leaving *int64 `json:"leaving,omitempty"`
	// [Unix Time Stamp](https://www.unixtimestamp.com/) of the date that this streaming option was found on the service. 
	AvailableSince int64 `json:"availableSince"`
}

// NewStreamingOption instantiates a new StreamingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamingOption(service string, streamingType StreamingType, link string, audios []Locale, subtitles []Subtitle, availableSince int64) *StreamingOption {
	this := StreamingOption{}
	this.Service = service
	this.StreamingType = streamingType
	this.Link = link
	this.Audios = audios
	this.Subtitles = subtitles
	this.AvailableSince = availableSince
	return &this
}

// NewStreamingOptionWithDefaults instantiates a new StreamingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamingOptionWithDefaults() *StreamingOption {
	this := StreamingOption{}
	return &this
}

// GetService returns the Service field value
func (o *StreamingOption) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *StreamingOption) SetService(v string) {
	o.Service = v
}

// GetStreamingType returns the StreamingType field value
func (o *StreamingOption) GetStreamingType() StreamingType {
	if o == nil {
		var ret StreamingType
		return ret
	}

	return o.StreamingType
}

// GetStreamingTypeOk returns a tuple with the StreamingType field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetStreamingTypeOk() (*StreamingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingType, true
}

// SetStreamingType sets field value
func (o *StreamingOption) SetStreamingType(v StreamingType) {
	o.StreamingType = v
}

// GetAddon returns the Addon field value if set, zero value otherwise.
func (o *StreamingOption) GetAddon() string {
	if o == nil || IsNil(o.Addon) {
		var ret string
		return ret
	}
	return *o.Addon
}

// GetAddonOk returns a tuple with the Addon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetAddonOk() (*string, bool) {
	if o == nil || IsNil(o.Addon) {
		return nil, false
	}
	return o.Addon, true
}

// HasAddon returns a boolean if a field has been set.
func (o *StreamingOption) HasAddon() bool {
	if o != nil && !IsNil(o.Addon) {
		return true
	}

	return false
}

// SetAddon gets a reference to the given string and assigns it to the Addon field.
func (o *StreamingOption) SetAddon(v string) {
	o.Addon = &v
}

// GetLink returns the Link field value
func (o *StreamingOption) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *StreamingOption) SetLink(v string) {
	o.Link = v
}

// GetVideoLink returns the VideoLink field value if set, zero value otherwise.
func (o *StreamingOption) GetVideoLink() string {
	if o == nil || IsNil(o.VideoLink) {
		var ret string
		return ret
	}
	return *o.VideoLink
}

// GetVideoLinkOk returns a tuple with the VideoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetVideoLinkOk() (*string, bool) {
	if o == nil || IsNil(o.VideoLink) {
		return nil, false
	}
	return o.VideoLink, true
}

// HasVideoLink returns a boolean if a field has been set.
func (o *StreamingOption) HasVideoLink() bool {
	if o != nil && !IsNil(o.VideoLink) {
		return true
	}

	return false
}

// SetVideoLink gets a reference to the given string and assigns it to the VideoLink field.
func (o *StreamingOption) SetVideoLink(v string) {
	o.VideoLink = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *StreamingOption) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *StreamingOption) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *StreamingOption) SetQuality(v string) {
	o.Quality = &v
}

// GetAudios returns the Audios field value
func (o *StreamingOption) GetAudios() []Locale {
	if o == nil {
		var ret []Locale
		return ret
	}

	return o.Audios
}

// GetAudiosOk returns a tuple with the Audios field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetAudiosOk() ([]Locale, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audios, true
}

// SetAudios sets field value
func (o *StreamingOption) SetAudios(v []Locale) {
	o.Audios = v
}

// GetSubtitles returns the Subtitles field value
func (o *StreamingOption) GetSubtitles() []Subtitle {
	if o == nil {
		var ret []Subtitle
		return ret
	}

	return o.Subtitles
}

// GetSubtitlesOk returns a tuple with the Subtitles field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetSubtitlesOk() ([]Subtitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtitles, true
}

// SetSubtitles sets field value
func (o *StreamingOption) SetSubtitles(v []Subtitle) {
	o.Subtitles = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *StreamingOption) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *StreamingOption) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *StreamingOption) SetPrice(v Price) {
	o.Price = &v
}

// GetLeaving returns the Leaving field value if set, zero value otherwise.
func (o *StreamingOption) GetLeaving() int64 {
	if o == nil || IsNil(o.Leaving) {
		var ret int64
		return ret
	}
	return *o.Leaving
}

// GetLeavingOk returns a tuple with the Leaving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetLeavingOk() (*int64, bool) {
	if o == nil || IsNil(o.Leaving) {
		return nil, false
	}
	return o.Leaving, true
}

// HasLeaving returns a boolean if a field has been set.
func (o *StreamingOption) HasLeaving() bool {
	if o != nil && !IsNil(o.Leaving) {
		return true
	}

	return false
}

// SetLeaving gets a reference to the given int64 and assigns it to the Leaving field.
func (o *StreamingOption) SetLeaving(v int64) {
	o.Leaving = &v
}

// GetAvailableSince returns the AvailableSince field value
func (o *StreamingOption) GetAvailableSince() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailableSince
}

// GetAvailableSinceOk returns a tuple with the AvailableSince field value
// and a boolean to check if the value has been set.
func (o *StreamingOption) GetAvailableSinceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableSince, true
}

// SetAvailableSince sets field value
func (o *StreamingOption) SetAvailableSince(v int64) {
	o.AvailableSince = v
}

func (o StreamingOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	toSerialize["streamingType"] = o.StreamingType
	if !IsNil(o.Addon) {
		toSerialize["addon"] = o.Addon
	}
	toSerialize["link"] = o.Link
	if !IsNil(o.VideoLink) {
		toSerialize["videoLink"] = o.VideoLink
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	toSerialize["audios"] = o.Audios
	toSerialize["subtitles"] = o.Subtitles
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Leaving) {
		toSerialize["leaving"] = o.Leaving
	}
	toSerialize["availableSince"] = o.AvailableSince
	return toSerialize, nil
}

type NullableStreamingOption struct {
	value *StreamingOption
	isSet bool
}

func (v NullableStreamingOption) Get() *StreamingOption {
	return v.value
}

func (v *NullableStreamingOption) Set(val *StreamingOption) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamingOption) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamingOption(val *StreamingOption) *NullableStreamingOption {
	return &NullableStreamingOption{value: val, isSet: true}
}

func (v NullableStreamingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


