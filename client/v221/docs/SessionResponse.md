# SessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | [**Session**](Session.md) |  | 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionResponse

`func NewSessionResponse(session Session, statusCode float32, ) *SessionResponse`

NewSessionResponse instantiates a new SessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseWithDefaults

`func NewSessionResponseWithDefaults() *SessionResponse`

NewSessionResponseWithDefaults instantiates a new SessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *SessionResponse) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SessionResponse) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SessionResponse) SetSession(v Session)`

SetSession sets Session field to given value.


### GetStatusCode

`func (o *SessionResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SessionResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SessionResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *SessionResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SessionResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SessionResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SessionResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *SessionResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *SessionResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *SessionResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *SessionResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


