# ChargingPreferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargingPreferences** | **string** |  | 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewChargingPreferencesResponse

`func NewChargingPreferencesResponse(chargingPreferences string, statusCode float32, ) *ChargingPreferencesResponse`

NewChargingPreferencesResponse instantiates a new ChargingPreferencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingPreferencesResponseWithDefaults

`func NewChargingPreferencesResponseWithDefaults() *ChargingPreferencesResponse`

NewChargingPreferencesResponseWithDefaults instantiates a new ChargingPreferencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargingPreferences

`func (o *ChargingPreferencesResponse) GetChargingPreferences() string`

GetChargingPreferences returns the ChargingPreferences field if non-nil, zero value otherwise.

### GetChargingPreferencesOk

`func (o *ChargingPreferencesResponse) GetChargingPreferencesOk() (*string, bool)`

GetChargingPreferencesOk returns a tuple with the ChargingPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingPreferences

`func (o *ChargingPreferencesResponse) SetChargingPreferences(v string)`

SetChargingPreferences sets ChargingPreferences field to given value.


### GetStatusCode

`func (o *ChargingPreferencesResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ChargingPreferencesResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ChargingPreferencesResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *ChargingPreferencesResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ChargingPreferencesResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ChargingPreferencesResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ChargingPreferencesResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ChargingPreferencesResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ChargingPreferencesResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ChargingPreferencesResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ChargingPreferencesResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


