# ChargingProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**ChargingProfileResponseType**](ChargingProfileResponseType.md) |  | 
**Timeout** | **int32** |  | 

## Methods

### NewChargingProfileResponse

`func NewChargingProfileResponse(result ChargingProfileResponseType, timeout int32, ) *ChargingProfileResponse`

NewChargingProfileResponse instantiates a new ChargingProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingProfileResponseWithDefaults

`func NewChargingProfileResponseWithDefaults() *ChargingProfileResponse`

NewChargingProfileResponseWithDefaults instantiates a new ChargingProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ChargingProfileResponse) GetResult() ChargingProfileResponseType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ChargingProfileResponse) GetResultOk() (*ChargingProfileResponseType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ChargingProfileResponse) SetResult(v ChargingProfileResponseType)`

SetResult sets Result field to given value.


### GetTimeout

`func (o *ChargingProfileResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ChargingProfileResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ChargingProfileResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


