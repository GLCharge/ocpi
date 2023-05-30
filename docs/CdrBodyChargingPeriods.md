# CdrBodyChargingPeriods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | **string** |  | 
**Dimensions** | Pointer to [**CdrBodyChargingPeriodsDimensions**](CdrBodyChargingPeriodsDimensions.md) |  | [optional] 
**TariffId** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrBodyChargingPeriods

`func NewCdrBodyChargingPeriods(startDateTime string, ) *CdrBodyChargingPeriods`

NewCdrBodyChargingPeriods instantiates a new CdrBodyChargingPeriods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrBodyChargingPeriodsWithDefaults

`func NewCdrBodyChargingPeriodsWithDefaults() *CdrBodyChargingPeriods`

NewCdrBodyChargingPeriodsWithDefaults instantiates a new CdrBodyChargingPeriods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *CdrBodyChargingPeriods) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CdrBodyChargingPeriods) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CdrBodyChargingPeriods) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetDimensions

`func (o *CdrBodyChargingPeriods) GetDimensions() CdrBodyChargingPeriodsDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CdrBodyChargingPeriods) GetDimensionsOk() (*CdrBodyChargingPeriodsDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CdrBodyChargingPeriods) SetDimensions(v CdrBodyChargingPeriodsDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CdrBodyChargingPeriods) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetTariffId

`func (o *CdrBodyChargingPeriods) GetTariffId() string`

GetTariffId returns the TariffId field if non-nil, zero value otherwise.

### GetTariffIdOk

`func (o *CdrBodyChargingPeriods) GetTariffIdOk() (*string, bool)`

GetTariffIdOk returns a tuple with the TariffId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffId

`func (o *CdrBodyChargingPeriods) SetTariffId(v string)`

SetTariffId sets TariffId field to given value.

### HasTariffId

`func (o *CdrBodyChargingPeriods) HasTariffId() bool`

HasTariffId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


