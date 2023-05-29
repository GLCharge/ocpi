# CommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Timeout** | **int32** |  | 
**Message** | Pointer to [**CommandResponseMessage**](CommandResponseMessage.md) |  | [optional] 

## Methods

### NewCommandResponse

`func NewCommandResponse(result string, timeout int32, ) *CommandResponse`

NewCommandResponse instantiates a new CommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandResponseWithDefaults

`func NewCommandResponseWithDefaults() *CommandResponse`

NewCommandResponseWithDefaults instantiates a new CommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CommandResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CommandResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CommandResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetTimeout

`func (o *CommandResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CommandResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CommandResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetMessage

`func (o *CommandResponse) GetMessage() CommandResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommandResponse) GetMessageOk() (*CommandResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommandResponse) SetMessage(v CommandResponseMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommandResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


