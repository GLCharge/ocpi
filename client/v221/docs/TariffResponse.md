# TariffResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tariff** | [**Tariff**](Tariff.md) |  | 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewTariffResponse

`func NewTariffResponse(tariff Tariff, statusCode float32, ) *TariffResponse`

NewTariffResponse instantiates a new TariffResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffResponseWithDefaults

`func NewTariffResponseWithDefaults() *TariffResponse`

NewTariffResponseWithDefaults instantiates a new TariffResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTariff

`func (o *TariffResponse) GetTariff() Tariff`

GetTariff returns the Tariff field if non-nil, zero value otherwise.

### GetTariffOk

`func (o *TariffResponse) GetTariffOk() (*Tariff, bool)`

GetTariffOk returns a tuple with the Tariff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariff

`func (o *TariffResponse) SetTariff(v Tariff)`

SetTariff sets Tariff field to given value.


### GetStatusCode

`func (o *TariffResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TariffResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TariffResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *TariffResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *TariffResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *TariffResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *TariffResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *TariffResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *TariffResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *TariffResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *TariffResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

