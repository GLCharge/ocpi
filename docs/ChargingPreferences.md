# ChargingPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileType** | **string** |  | 
**DepartureTime** | Pointer to **string** |  | [optional] 
**EnergyNeed** | Pointer to **float32** |  | [optional] 

## Methods

### NewChargingPreferences

`func NewChargingPreferences(profileType string, ) *ChargingPreferences`

NewChargingPreferences instantiates a new ChargingPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingPreferencesWithDefaults

`func NewChargingPreferencesWithDefaults() *ChargingPreferences`

NewChargingPreferencesWithDefaults instantiates a new ChargingPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileType

`func (o *ChargingPreferences) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *ChargingPreferences) GetProfileTypeOk() (*string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *ChargingPreferences) SetProfileType(v string)`

SetProfileType sets ProfileType field to given value.


### GetDepartureTime

`func (o *ChargingPreferences) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *ChargingPreferences) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *ChargingPreferences) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *ChargingPreferences) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetEnergyNeed

`func (o *ChargingPreferences) GetEnergyNeed() float32`

GetEnergyNeed returns the EnergyNeed field if non-nil, zero value otherwise.

### GetEnergyNeedOk

`func (o *ChargingPreferences) GetEnergyNeedOk() (*float32, bool)`

GetEnergyNeedOk returns a tuple with the EnergyNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyNeed

`func (o *ChargingPreferences) SetEnergyNeed(v float32)`

SetEnergyNeed sets EnergyNeed field to given value.

### HasEnergyNeed

`func (o *ChargingPreferences) HasEnergyNeed() bool`

HasEnergyNeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


