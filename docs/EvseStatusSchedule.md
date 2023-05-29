# EvseStatusSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodBegin** | **string** |  | 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewEvseStatusSchedule

`func NewEvseStatusSchedule(periodBegin string, status string, ) *EvseStatusSchedule`

NewEvseStatusSchedule instantiates a new EvseStatusSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvseStatusScheduleWithDefaults

`func NewEvseStatusScheduleWithDefaults() *EvseStatusSchedule`

NewEvseStatusScheduleWithDefaults instantiates a new EvseStatusSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodBegin

`func (o *EvseStatusSchedule) GetPeriodBegin() string`

GetPeriodBegin returns the PeriodBegin field if non-nil, zero value otherwise.

### GetPeriodBeginOk

`func (o *EvseStatusSchedule) GetPeriodBeginOk() (*string, bool)`

GetPeriodBeginOk returns a tuple with the PeriodBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBegin

`func (o *EvseStatusSchedule) SetPeriodBegin(v string)`

SetPeriodBegin sets PeriodBegin field to given value.


### GetPeriodEnd

`func (o *EvseStatusSchedule) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *EvseStatusSchedule) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *EvseStatusSchedule) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *EvseStatusSchedule) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetStatus

`func (o *EvseStatusSchedule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvseStatusSchedule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvseStatusSchedule) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


