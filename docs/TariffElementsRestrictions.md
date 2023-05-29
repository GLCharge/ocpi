# TariffElementsRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**MinKwh** | Pointer to **float32** |  | [optional] 
**MaxKwh** | Pointer to **float32** |  | [optional] 
**MinCurrent** | Pointer to **float32** |  | [optional] 
**MaxCurrent** | Pointer to **float32** |  | [optional] 
**MinPower** | Pointer to **float32** |  | [optional] 
**MaxPower** | Pointer to **float32** |  | [optional] 
**MinDuration** | Pointer to **int32** |  | [optional] 
**MaxDuration** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | Pointer to **string** |  | [optional] 
**Reservation** | Pointer to [**ReservationRestrictionType**](ReservationRestrictionType.md) |  | [optional] 

## Methods

### NewTariffElementsRestrictions

`func NewTariffElementsRestrictions() *TariffElementsRestrictions`

NewTariffElementsRestrictions instantiates a new TariffElementsRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffElementsRestrictionsWithDefaults

`func NewTariffElementsRestrictionsWithDefaults() *TariffElementsRestrictions`

NewTariffElementsRestrictionsWithDefaults instantiates a new TariffElementsRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *TariffElementsRestrictions) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TariffElementsRestrictions) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TariffElementsRestrictions) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TariffElementsRestrictions) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TariffElementsRestrictions) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TariffElementsRestrictions) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TariffElementsRestrictions) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TariffElementsRestrictions) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *TariffElementsRestrictions) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TariffElementsRestrictions) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TariffElementsRestrictions) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TariffElementsRestrictions) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TariffElementsRestrictions) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TariffElementsRestrictions) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TariffElementsRestrictions) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TariffElementsRestrictions) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMinKwh

`func (o *TariffElementsRestrictions) GetMinKwh() float32`

GetMinKwh returns the MinKwh field if non-nil, zero value otherwise.

### GetMinKwhOk

`func (o *TariffElementsRestrictions) GetMinKwhOk() (*float32, bool)`

GetMinKwhOk returns a tuple with the MinKwh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinKwh

`func (o *TariffElementsRestrictions) SetMinKwh(v float32)`

SetMinKwh sets MinKwh field to given value.

### HasMinKwh

`func (o *TariffElementsRestrictions) HasMinKwh() bool`

HasMinKwh returns a boolean if a field has been set.

### GetMaxKwh

`func (o *TariffElementsRestrictions) GetMaxKwh() float32`

GetMaxKwh returns the MaxKwh field if non-nil, zero value otherwise.

### GetMaxKwhOk

`func (o *TariffElementsRestrictions) GetMaxKwhOk() (*float32, bool)`

GetMaxKwhOk returns a tuple with the MaxKwh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKwh

`func (o *TariffElementsRestrictions) SetMaxKwh(v float32)`

SetMaxKwh sets MaxKwh field to given value.

### HasMaxKwh

`func (o *TariffElementsRestrictions) HasMaxKwh() bool`

HasMaxKwh returns a boolean if a field has been set.

### GetMinCurrent

`func (o *TariffElementsRestrictions) GetMinCurrent() float32`

GetMinCurrent returns the MinCurrent field if non-nil, zero value otherwise.

### GetMinCurrentOk

`func (o *TariffElementsRestrictions) GetMinCurrentOk() (*float32, bool)`

GetMinCurrentOk returns a tuple with the MinCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCurrent

`func (o *TariffElementsRestrictions) SetMinCurrent(v float32)`

SetMinCurrent sets MinCurrent field to given value.

### HasMinCurrent

`func (o *TariffElementsRestrictions) HasMinCurrent() bool`

HasMinCurrent returns a boolean if a field has been set.

### GetMaxCurrent

`func (o *TariffElementsRestrictions) GetMaxCurrent() float32`

GetMaxCurrent returns the MaxCurrent field if non-nil, zero value otherwise.

### GetMaxCurrentOk

`func (o *TariffElementsRestrictions) GetMaxCurrentOk() (*float32, bool)`

GetMaxCurrentOk returns a tuple with the MaxCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCurrent

`func (o *TariffElementsRestrictions) SetMaxCurrent(v float32)`

SetMaxCurrent sets MaxCurrent field to given value.

### HasMaxCurrent

`func (o *TariffElementsRestrictions) HasMaxCurrent() bool`

HasMaxCurrent returns a boolean if a field has been set.

### GetMinPower

`func (o *TariffElementsRestrictions) GetMinPower() float32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *TariffElementsRestrictions) GetMinPowerOk() (*float32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *TariffElementsRestrictions) SetMinPower(v float32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *TariffElementsRestrictions) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMaxPower

`func (o *TariffElementsRestrictions) GetMaxPower() float32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *TariffElementsRestrictions) GetMaxPowerOk() (*float32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *TariffElementsRestrictions) SetMaxPower(v float32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *TariffElementsRestrictions) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinDuration

`func (o *TariffElementsRestrictions) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *TariffElementsRestrictions) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *TariffElementsRestrictions) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *TariffElementsRestrictions) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetMaxDuration

`func (o *TariffElementsRestrictions) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *TariffElementsRestrictions) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *TariffElementsRestrictions) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *TariffElementsRestrictions) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *TariffElementsRestrictions) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TariffElementsRestrictions) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TariffElementsRestrictions) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TariffElementsRestrictions) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetReservation

`func (o *TariffElementsRestrictions) GetReservation() ReservationRestrictionType`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *TariffElementsRestrictions) GetReservationOk() (*ReservationRestrictionType, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *TariffElementsRestrictions) SetReservation(v ReservationRestrictionType)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *TariffElementsRestrictions) HasReservation() bool`

HasReservation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


