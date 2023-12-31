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

// checks if the UpcomingChangeSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpcomingChangeSet{}

// UpcomingChangeSet struct for UpcomingChangeSet
type UpcomingChangeSet struct {
	UpcomingChanges []Change `json:"upcomingChanges"`
	Show Show `json:"show"`
}

// NewUpcomingChangeSet instantiates a new UpcomingChangeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpcomingChangeSet(upcomingChanges []Change, show Show) *UpcomingChangeSet {
	this := UpcomingChangeSet{}
	this.UpcomingChanges = upcomingChanges
	this.Show = show
	return &this
}

// NewUpcomingChangeSetWithDefaults instantiates a new UpcomingChangeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpcomingChangeSetWithDefaults() *UpcomingChangeSet {
	this := UpcomingChangeSet{}
	return &this
}

// GetUpcomingChanges returns the UpcomingChanges field value
func (o *UpcomingChangeSet) GetUpcomingChanges() []Change {
	if o == nil {
		var ret []Change
		return ret
	}

	return o.UpcomingChanges
}

// GetUpcomingChangesOk returns a tuple with the UpcomingChanges field value
// and a boolean to check if the value has been set.
func (o *UpcomingChangeSet) GetUpcomingChangesOk() ([]Change, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpcomingChanges, true
}

// SetUpcomingChanges sets field value
func (o *UpcomingChangeSet) SetUpcomingChanges(v []Change) {
	o.UpcomingChanges = v
}

// GetShow returns the Show field value
func (o *UpcomingChangeSet) GetShow() Show {
	if o == nil {
		var ret Show
		return ret
	}

	return o.Show
}

// GetShowOk returns a tuple with the Show field value
// and a boolean to check if the value has been set.
func (o *UpcomingChangeSet) GetShowOk() (*Show, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Show, true
}

// SetShow sets field value
func (o *UpcomingChangeSet) SetShow(v Show) {
	o.Show = v
}

func (o UpcomingChangeSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpcomingChangeSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upcomingChanges"] = o.UpcomingChanges
	toSerialize["show"] = o.Show
	return toSerialize, nil
}

type NullableUpcomingChangeSet struct {
	value *UpcomingChangeSet
	isSet bool
}

func (v NullableUpcomingChangeSet) Get() *UpcomingChangeSet {
	return v.value
}

func (v *NullableUpcomingChangeSet) Set(val *UpcomingChangeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableUpcomingChangeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableUpcomingChangeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpcomingChangeSet(val *UpcomingChangeSet) *NullableUpcomingChangeSet {
	return &NullableUpcomingChangeSet{value: val, isSet: true}
}

func (v NullableUpcomingChangeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpcomingChangeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


