# ActiveChargingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | **string** |  | 
**ChargingProfile** | [**ChargingProfile**](ChargingProfile.md) |  | 

## Methods

### NewActiveChargingProfile

`func NewActiveChargingProfile(startDateTime string, chargingProfile ChargingProfile, ) *ActiveChargingProfile`

NewActiveChargingProfile instantiates a new ActiveChargingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveChargingProfileWithDefaults

`func NewActiveChargingProfileWithDefaults() *ActiveChargingProfile`

NewActiveChargingProfileWithDefaults instantiates a new ActiveChargingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *ActiveChargingProfile) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ActiveChargingProfile) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ActiveChargingProfile) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetChargingProfile

`func (o *ActiveChargingProfile) GetChargingProfile() ChargingProfile`

GetChargingProfile returns the ChargingProfile field if non-nil, zero value otherwise.

### GetChargingProfileOk

`func (o *ActiveChargingProfile) GetChargingProfileOk() (*ChargingProfile, bool)`

GetChargingProfileOk returns a tuple with the ChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingProfile

`func (o *ActiveChargingProfile) SetChargingProfile(v ChargingProfile)`

SetChargingProfile sets ChargingProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


