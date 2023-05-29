# ActiveChargingProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Profile** | Pointer to [**ActiveChargingProfileResultProfile**](ActiveChargingProfileResultProfile.md) |  | [optional] 

## Methods

### NewActiveChargingProfileResult

`func NewActiveChargingProfileResult(result string, ) *ActiveChargingProfileResult`

NewActiveChargingProfileResult instantiates a new ActiveChargingProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveChargingProfileResultWithDefaults

`func NewActiveChargingProfileResultWithDefaults() *ActiveChargingProfileResult`

NewActiveChargingProfileResultWithDefaults instantiates a new ActiveChargingProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ActiveChargingProfileResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ActiveChargingProfileResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ActiveChargingProfileResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetProfile

`func (o *ActiveChargingProfileResult) GetProfile() ActiveChargingProfileResultProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ActiveChargingProfileResult) GetProfileOk() (*ActiveChargingProfileResultProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ActiveChargingProfileResult) SetProfile(v ActiveChargingProfileResultProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ActiveChargingProfileResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


