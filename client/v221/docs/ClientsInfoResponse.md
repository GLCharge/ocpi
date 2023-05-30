# ClientsInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientInfo** | [**ClientInfo**](ClientInfo.md) |  | 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewClientsInfoResponse

`func NewClientsInfoResponse(clientInfo ClientInfo, statusCode float32, ) *ClientsInfoResponse`

NewClientsInfoResponse instantiates a new ClientsInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientsInfoResponseWithDefaults

`func NewClientsInfoResponseWithDefaults() *ClientsInfoResponse`

NewClientsInfoResponseWithDefaults instantiates a new ClientsInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientInfo

`func (o *ClientsInfoResponse) GetClientInfo() ClientInfo`

GetClientInfo returns the ClientInfo field if non-nil, zero value otherwise.

### GetClientInfoOk

`func (o *ClientsInfoResponse) GetClientInfoOk() (*ClientInfo, bool)`

GetClientInfoOk returns a tuple with the ClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInfo

`func (o *ClientsInfoResponse) SetClientInfo(v ClientInfo)`

SetClientInfo sets ClientInfo field to given value.


### GetStatusCode

`func (o *ClientsInfoResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ClientsInfoResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ClientsInfoResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *ClientsInfoResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ClientsInfoResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ClientsInfoResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ClientsInfoResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ClientsInfoResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ClientsInfoResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ClientsInfoResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ClientsInfoResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


