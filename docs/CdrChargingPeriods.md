# CdrChargingPeriods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | **string** |  | 
**Dimensions** | Pointer to [**CdrChargingPeriodsDimensions**](CdrChargingPeriodsDimensions.md) |  | [optional] 
**TariffId** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrChargingPeriods

`func NewCdrChargingPeriods(startDateTime string, ) *CdrChargingPeriods`

NewCdrChargingPeriods instantiates a new CdrChargingPeriods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrChargingPeriodsWithDefaults

`func NewCdrChargingPeriodsWithDefaults() *CdrChargingPeriods`

NewCdrChargingPeriodsWithDefaults instantiates a new CdrChargingPeriods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *CdrChargingPeriods) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CdrChargingPeriods) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CdrChargingPeriods) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetDimensions

`func (o *CdrChargingPeriods) GetDimensions() CdrChargingPeriodsDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CdrChargingPeriods) GetDimensionsOk() (*CdrChargingPeriodsDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CdrChargingPeriods) SetDimensions(v CdrChargingPeriodsDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CdrChargingPeriods) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetTariffId

`func (o *CdrChargingPeriods) GetTariffId() string`

GetTariffId returns the TariffId field if non-nil, zero value otherwise.

### GetTariffIdOk

`func (o *CdrChargingPeriods) GetTariffIdOk() (*string, bool)`

GetTariffIdOk returns a tuple with the TariffId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffId

`func (o *CdrChargingPeriods) SetTariffId(v string)`

SetTariffId sets TariffId field to given value.

### HasTariffId

`func (o *CdrChargingPeriods) HasTariffId() bool`

HasTariffId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


