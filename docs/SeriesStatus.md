# SeriesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | Integer code of the status. 1: \&quot;Returning Series\&quot;, 2: \&quot;Planned\&quot;, 3: \&quot;In Production\&quot;, 4: \&quot;Ended\&quot;, 5: \&quot;Cancelled\&quot;, 6: \&quot;Pilot\&quot;,  | 
**StatusText** | **string** | Textual representation of the status. | 

## Methods

### NewSeriesStatus

`func NewSeriesStatus(statusCode int32, statusText string, ) *SeriesStatus`

NewSeriesStatus instantiates a new SeriesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesStatusWithDefaults

`func NewSeriesStatusWithDefaults() *SeriesStatus`

NewSeriesStatusWithDefaults instantiates a new SeriesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *SeriesStatus) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SeriesStatus) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SeriesStatus) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusText

`func (o *SeriesStatus) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *SeriesStatus) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *SeriesStatus) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


