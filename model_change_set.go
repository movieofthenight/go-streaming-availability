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

// checks if the ChangeSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeSet{}

// ChangeSet struct for ChangeSet
type ChangeSet struct {
	Changes []Change `json:"changes"`
	Show Show `json:"show"`
}

// NewChangeSet instantiates a new ChangeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeSet(changes []Change, show Show) *ChangeSet {
	this := ChangeSet{}
	this.Changes = changes
	this.Show = show
	return &this
}

// NewChangeSetWithDefaults instantiates a new ChangeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeSetWithDefaults() *ChangeSet {
	this := ChangeSet{}
	return &this
}

// GetChanges returns the Changes field value
func (o *ChangeSet) GetChanges() []Change {
	if o == nil {
		var ret []Change
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetChangesOk() ([]Change, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *ChangeSet) SetChanges(v []Change) {
	o.Changes = v
}

// GetShow returns the Show field value
func (o *ChangeSet) GetShow() Show {
	if o == nil {
		var ret Show
		return ret
	}

	return o.Show
}

// GetShowOk returns a tuple with the Show field value
// and a boolean to check if the value has been set.
func (o *ChangeSet) GetShowOk() (*Show, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Show, true
}

// SetShow sets field value
func (o *ChangeSet) SetShow(v Show) {
	o.Show = v
}

func (o ChangeSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["changes"] = o.Changes
	toSerialize["show"] = o.Show
	return toSerialize, nil
}

type NullableChangeSet struct {
	value *ChangeSet
	isSet bool
}

func (v NullableChangeSet) Get() *ChangeSet {
	return v.value
}

func (v *NullableChangeSet) Set(val *ChangeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeSet(val *ChangeSet) *NullableChangeSet {
	return &NullableChangeSet{value: val, isSet: true}
}

func (v NullableChangeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


