# CommandResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Message** | Pointer to [**CommandResponseMessage**](CommandResponseMessage.md) |  | [optional] 

## Methods

### NewCommandResult

`func NewCommandResult(result string, ) *CommandResult`

NewCommandResult instantiates a new CommandResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandResultWithDefaults

`func NewCommandResultWithDefaults() *CommandResult`

NewCommandResultWithDefaults instantiates a new CommandResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CommandResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CommandResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CommandResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetMessage

`func (o *CommandResult) GetMessage() CommandResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommandResult) GetMessageOk() (*CommandResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommandResult) SetMessage(v CommandResponseMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommandResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


