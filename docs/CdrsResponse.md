# CdrsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdrs** | Pointer to [**Cdr**](Cdr.md) |  | [optional] 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrsResponse

`func NewCdrsResponse(statusCode float32, ) *CdrsResponse`

NewCdrsResponse instantiates a new CdrsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrsResponseWithDefaults

`func NewCdrsResponseWithDefaults() *CdrsResponse`

NewCdrsResponseWithDefaults instantiates a new CdrsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdrs

`func (o *CdrsResponse) GetCdrs() Cdr`

GetCdrs returns the Cdrs field if non-nil, zero value otherwise.

### GetCdrsOk

`func (o *CdrsResponse) GetCdrsOk() (*Cdr, bool)`

GetCdrsOk returns a tuple with the Cdrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrs

`func (o *CdrsResponse) SetCdrs(v Cdr)`

SetCdrs sets Cdrs field to given value.

### HasCdrs

`func (o *CdrsResponse) HasCdrs() bool`

HasCdrs returns a boolean if a field has been set.

### GetStatusCode

`func (o *CdrsResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CdrsResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CdrsResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *CdrsResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CdrsResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CdrsResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CdrsResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *CdrsResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *CdrsResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *CdrsResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *CdrsResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


