# ClientInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsInfo** | Pointer to [**ClientInfo**](ClientInfo.md) |  | [optional] 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewClientInfoResponse

`func NewClientInfoResponse(statusCode float32, ) *ClientInfoResponse`

NewClientInfoResponse instantiates a new ClientInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoResponseWithDefaults

`func NewClientInfoResponseWithDefaults() *ClientInfoResponse`

NewClientInfoResponseWithDefaults instantiates a new ClientInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsInfo

`func (o *ClientInfoResponse) GetClientsInfo() ClientInfo`

GetClientsInfo returns the ClientsInfo field if non-nil, zero value otherwise.

### GetClientsInfoOk

`func (o *ClientInfoResponse) GetClientsInfoOk() (*ClientInfo, bool)`

GetClientsInfoOk returns a tuple with the ClientsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsInfo

`func (o *ClientInfoResponse) SetClientsInfo(v ClientInfo)`

SetClientsInfo sets ClientsInfo field to given value.

### HasClientsInfo

`func (o *ClientInfoResponse) HasClientsInfo() bool`

HasClientsInfo returns a boolean if a field has been set.

### GetStatusCode

`func (o *ClientInfoResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ClientInfoResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ClientInfoResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *ClientInfoResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ClientInfoResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ClientInfoResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ClientInfoResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ClientInfoResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ClientInfoResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ClientInfoResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ClientInfoResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


