# ChargingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**ChargingRateUnit** | **string** |  | 
**MinChargingRate** | Pointer to **float32** |  | [optional] 
**ChargingProfilePeriod** | Pointer to [**ChargingProfileChargingProfilePeriod**](ChargingProfileChargingProfilePeriod.md) |  | [optional] 

## Methods

### NewChargingProfile

`func NewChargingProfile(chargingRateUnit string, ) *ChargingProfile`

NewChargingProfile instantiates a new ChargingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingProfileWithDefaults

`func NewChargingProfileWithDefaults() *ChargingProfile`

NewChargingProfileWithDefaults instantiates a new ChargingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *ChargingProfile) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ChargingProfile) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ChargingProfile) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ChargingProfile) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetDuration

`func (o *ChargingProfile) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ChargingProfile) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ChargingProfile) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ChargingProfile) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetChargingRateUnit

`func (o *ChargingProfile) GetChargingRateUnit() string`

GetChargingRateUnit returns the ChargingRateUnit field if non-nil, zero value otherwise.

### GetChargingRateUnitOk

`func (o *ChargingProfile) GetChargingRateUnitOk() (*string, bool)`

GetChargingRateUnitOk returns a tuple with the ChargingRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingRateUnit

`func (o *ChargingProfile) SetChargingRateUnit(v string)`

SetChargingRateUnit sets ChargingRateUnit field to given value.


### GetMinChargingRate

`func (o *ChargingProfile) GetMinChargingRate() float32`

GetMinChargingRate returns the MinChargingRate field if non-nil, zero value otherwise.

### GetMinChargingRateOk

`func (o *ChargingProfile) GetMinChargingRateOk() (*float32, bool)`

GetMinChargingRateOk returns a tuple with the MinChargingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinChargingRate

`func (o *ChargingProfile) SetMinChargingRate(v float32)`

SetMinChargingRate sets MinChargingRate field to given value.

### HasMinChargingRate

`func (o *ChargingProfile) HasMinChargingRate() bool`

HasMinChargingRate returns a boolean if a field has been set.

### GetChargingProfilePeriod

`func (o *ChargingProfile) GetChargingProfilePeriod() ChargingProfileChargingProfilePeriod`

GetChargingProfilePeriod returns the ChargingProfilePeriod field if non-nil, zero value otherwise.

### GetChargingProfilePeriodOk

`func (o *ChargingProfile) GetChargingProfilePeriodOk() (*ChargingProfileChargingProfilePeriod, bool)`

GetChargingProfilePeriodOk returns a tuple with the ChargingProfilePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingProfilePeriod

`func (o *ChargingProfile) SetChargingProfilePeriod(v ChargingProfileChargingProfilePeriod)`

SetChargingProfilePeriod sets ChargingProfilePeriod field to given value.

### HasChargingProfilePeriod

`func (o *ChargingProfile) HasChargingProfilePeriod() bool`

HasChargingProfilePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


