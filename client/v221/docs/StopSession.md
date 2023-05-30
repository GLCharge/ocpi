# StopSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseUrl** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewStopSession

`func NewStopSession(responseUrl string, ) *StopSession`

NewStopSession instantiates a new StopSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopSessionWithDefaults

`func NewStopSessionWithDefaults() *StopSession`

NewStopSessionWithDefaults instantiates a new StopSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseUrl

`func (o *StopSession) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *StopSession) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *StopSession) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.


### GetSessionId

`func (o *StopSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *StopSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *StopSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *StopSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


