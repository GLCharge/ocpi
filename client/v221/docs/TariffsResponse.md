# TariffsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tariffs** | Pointer to [**Tariff**](Tariff.md) |  | [optional] 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewTariffsResponse

`func NewTariffsResponse(statusCode float32, ) *TariffsResponse`

NewTariffsResponse instantiates a new TariffsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffsResponseWithDefaults

`func NewTariffsResponseWithDefaults() *TariffsResponse`

NewTariffsResponseWithDefaults instantiates a new TariffsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTariffs

`func (o *TariffsResponse) GetTariffs() Tariff`

GetTariffs returns the Tariffs field if non-nil, zero value otherwise.

### GetTariffsOk

`func (o *TariffsResponse) GetTariffsOk() (*Tariff, bool)`

GetTariffsOk returns a tuple with the Tariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffs

`func (o *TariffsResponse) SetTariffs(v Tariff)`

SetTariffs sets Tariffs field to given value.

### HasTariffs

`func (o *TariffsResponse) HasTariffs() bool`

HasTariffs returns a boolean if a field has been set.

### GetStatusCode

`func (o *TariffsResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TariffsResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TariffsResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *TariffsResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *TariffsResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *TariffsResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *TariffsResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *TariffsResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *TariffsResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *TariffsResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *TariffsResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


