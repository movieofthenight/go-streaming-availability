# ChangeSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]Change**](Change.md) |  | 
**Show** | [**Show**](Show.md) |  | 

## Methods

### NewChangeSet

`func NewChangeSet(changes []Change, show Show, ) *ChangeSet`

NewChangeSet instantiates a new ChangeSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSetWithDefaults

`func NewChangeSetWithDefaults() *ChangeSet`

NewChangeSetWithDefaults instantiates a new ChangeSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *ChangeSet) GetChanges() []Change`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChangeSet) GetChangesOk() (*[]Change, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChangeSet) SetChanges(v []Change)`

SetChanges sets Changes field to given value.


### GetShow

`func (o *ChangeSet) GetShow() Show`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *ChangeSet) GetShowOk() (*Show, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *ChangeSet) SetShow(v Show)`

SetShow sets Show field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


