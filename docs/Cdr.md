# Cdr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**StartDateTime** | **string** |  | 
**EndDateTime** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**CdrToken** | [**CdrCdrToken**](CdrCdrToken.md) |  | 
**AuthMethod** | **string** |  | 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**CdrLocation** | [**CdrCdrLocation**](CdrCdrLocation.md) |  | 
**MeterId** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**Tariffs** | Pointer to [**CdrTariffs**](CdrTariffs.md) |  | [optional] 
**ChargingPeriods** | Pointer to [**CdrChargingPeriods**](CdrChargingPeriods.md) |  | [optional] 
**SignedData** | Pointer to [**CdrSignedData**](CdrSignedData.md) |  | [optional] 
**TotalCost** | [**Price**](Price.md) |  | 
**TotalFixedCost** | Pointer to [**Price**](Price.md) |  | [optional] 
**TotalEnergy** | **float32** |  | 
**TotalEnergyCost** | Pointer to [**Price**](Price.md) |  | [optional] 
**TotalTime** | **float32** |  | 
**TotalTimeCost** | Pointer to [**Price**](Price.md) |  | [optional] 
**TotalParkingTime** | Pointer to **float32** |  | [optional] 
**TotalParkingCost** | Pointer to [**Price**](Price.md) |  | [optional] 
**TotalReservationCost** | Pointer to [**Price**](Price.md) |  | [optional] 
**Remark** | Pointer to **string** |  | [optional] 
**InvoiceReferenceId** | Pointer to **string** |  | [optional] 
**Credit** | Pointer to **bool** |  | [optional] 
**CreditReferenceId** | Pointer to **string** |  | [optional] 
**HomeChargingCompensation** | Pointer to **bool** |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewCdr

`func NewCdr(countryCode string, partyId string, id string, startDateTime string, endDateTime string, cdrToken CdrCdrToken, authMethod string, cdrLocation CdrCdrLocation, currency string, totalCost Price, totalEnergy float32, totalTime float32, lastUpdated string, ) *Cdr`

NewCdr instantiates a new Cdr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrWithDefaults

`func NewCdrWithDefaults() *Cdr`

NewCdrWithDefaults instantiates a new Cdr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Cdr) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Cdr) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Cdr) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *Cdr) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *Cdr) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *Cdr) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *Cdr) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cdr) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cdr) SetId(v string)`

SetId sets Id field to given value.


### GetStartDateTime

`func (o *Cdr) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Cdr) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Cdr) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *Cdr) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Cdr) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Cdr) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.


### GetSessionId

`func (o *Cdr) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Cdr) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Cdr) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Cdr) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCdrToken

`func (o *Cdr) GetCdrToken() CdrCdrToken`

GetCdrToken returns the CdrToken field if non-nil, zero value otherwise.

### GetCdrTokenOk

`func (o *Cdr) GetCdrTokenOk() (*CdrCdrToken, bool)`

GetCdrTokenOk returns a tuple with the CdrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrToken

`func (o *Cdr) SetCdrToken(v CdrCdrToken)`

SetCdrToken sets CdrToken field to given value.


### GetAuthMethod

`func (o *Cdr) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Cdr) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Cdr) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetAuthorizationReference

`func (o *Cdr) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *Cdr) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *Cdr) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *Cdr) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetCdrLocation

`func (o *Cdr) GetCdrLocation() CdrCdrLocation`

GetCdrLocation returns the CdrLocation field if non-nil, zero value otherwise.

### GetCdrLocationOk

`func (o *Cdr) GetCdrLocationOk() (*CdrCdrLocation, bool)`

GetCdrLocationOk returns a tuple with the CdrLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrLocation

`func (o *Cdr) SetCdrLocation(v CdrCdrLocation)`

SetCdrLocation sets CdrLocation field to given value.


### GetMeterId

`func (o *Cdr) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *Cdr) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *Cdr) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *Cdr) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetCurrency

`func (o *Cdr) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Cdr) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Cdr) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTariffs

`func (o *Cdr) GetTariffs() CdrTariffs`

GetTariffs returns the Tariffs field if non-nil, zero value otherwise.

### GetTariffsOk

`func (o *Cdr) GetTariffsOk() (*CdrTariffs, bool)`

GetTariffsOk returns a tuple with the Tariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffs

`func (o *Cdr) SetTariffs(v CdrTariffs)`

SetTariffs sets Tariffs field to given value.

### HasTariffs

`func (o *Cdr) HasTariffs() bool`

HasTariffs returns a boolean if a field has been set.

### GetChargingPeriods

`func (o *Cdr) GetChargingPeriods() CdrChargingPeriods`

GetChargingPeriods returns the ChargingPeriods field if non-nil, zero value otherwise.

### GetChargingPeriodsOk

`func (o *Cdr) GetChargingPeriodsOk() (*CdrChargingPeriods, bool)`

GetChargingPeriodsOk returns a tuple with the ChargingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingPeriods

`func (o *Cdr) SetChargingPeriods(v CdrChargingPeriods)`

SetChargingPeriods sets ChargingPeriods field to given value.

### HasChargingPeriods

`func (o *Cdr) HasChargingPeriods() bool`

HasChargingPeriods returns a boolean if a field has been set.

### GetSignedData

`func (o *Cdr) GetSignedData() CdrSignedData`

GetSignedData returns the SignedData field if non-nil, zero value otherwise.

### GetSignedDataOk

`func (o *Cdr) GetSignedDataOk() (*CdrSignedData, bool)`

GetSignedDataOk returns a tuple with the SignedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedData

`func (o *Cdr) SetSignedData(v CdrSignedData)`

SetSignedData sets SignedData field to given value.

### HasSignedData

`func (o *Cdr) HasSignedData() bool`

HasSignedData returns a boolean if a field has been set.

### GetTotalCost

`func (o *Cdr) GetTotalCost() Price`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *Cdr) GetTotalCostOk() (*Price, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *Cdr) SetTotalCost(v Price)`

SetTotalCost sets TotalCost field to given value.


### GetTotalFixedCost

`func (o *Cdr) GetTotalFixedCost() Price`

GetTotalFixedCost returns the TotalFixedCost field if non-nil, zero value otherwise.

### GetTotalFixedCostOk

`func (o *Cdr) GetTotalFixedCostOk() (*Price, bool)`

GetTotalFixedCostOk returns a tuple with the TotalFixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFixedCost

`func (o *Cdr) SetTotalFixedCost(v Price)`

SetTotalFixedCost sets TotalFixedCost field to given value.

### HasTotalFixedCost

`func (o *Cdr) HasTotalFixedCost() bool`

HasTotalFixedCost returns a boolean if a field has been set.

### GetTotalEnergy

`func (o *Cdr) GetTotalEnergy() float32`

GetTotalEnergy returns the TotalEnergy field if non-nil, zero value otherwise.

### GetTotalEnergyOk

`func (o *Cdr) GetTotalEnergyOk() (*float32, bool)`

GetTotalEnergyOk returns a tuple with the TotalEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnergy

`func (o *Cdr) SetTotalEnergy(v float32)`

SetTotalEnergy sets TotalEnergy field to given value.


### GetTotalEnergyCost

`func (o *Cdr) GetTotalEnergyCost() Price`

GetTotalEnergyCost returns the TotalEnergyCost field if non-nil, zero value otherwise.

### GetTotalEnergyCostOk

`func (o *Cdr) GetTotalEnergyCostOk() (*Price, bool)`

GetTotalEnergyCostOk returns a tuple with the TotalEnergyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnergyCost

`func (o *Cdr) SetTotalEnergyCost(v Price)`

SetTotalEnergyCost sets TotalEnergyCost field to given value.

### HasTotalEnergyCost

`func (o *Cdr) HasTotalEnergyCost() bool`

HasTotalEnergyCost returns a boolean if a field has been set.

### GetTotalTime

`func (o *Cdr) GetTotalTime() float32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *Cdr) GetTotalTimeOk() (*float32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *Cdr) SetTotalTime(v float32)`

SetTotalTime sets TotalTime field to given value.


### GetTotalTimeCost

`func (o *Cdr) GetTotalTimeCost() Price`

GetTotalTimeCost returns the TotalTimeCost field if non-nil, zero value otherwise.

### GetTotalTimeCostOk

`func (o *Cdr) GetTotalTimeCostOk() (*Price, bool)`

GetTotalTimeCostOk returns a tuple with the TotalTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeCost

`func (o *Cdr) SetTotalTimeCost(v Price)`

SetTotalTimeCost sets TotalTimeCost field to given value.

### HasTotalTimeCost

`func (o *Cdr) HasTotalTimeCost() bool`

HasTotalTimeCost returns a boolean if a field has been set.

### GetTotalParkingTime

`func (o *Cdr) GetTotalParkingTime() float32`

GetTotalParkingTime returns the TotalParkingTime field if non-nil, zero value otherwise.

### GetTotalParkingTimeOk

`func (o *Cdr) GetTotalParkingTimeOk() (*float32, bool)`

GetTotalParkingTimeOk returns a tuple with the TotalParkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParkingTime

`func (o *Cdr) SetTotalParkingTime(v float32)`

SetTotalParkingTime sets TotalParkingTime field to given value.

### HasTotalParkingTime

`func (o *Cdr) HasTotalParkingTime() bool`

HasTotalParkingTime returns a boolean if a field has been set.

### GetTotalParkingCost

`func (o *Cdr) GetTotalParkingCost() Price`

GetTotalParkingCost returns the TotalParkingCost field if non-nil, zero value otherwise.

### GetTotalParkingCostOk

`func (o *Cdr) GetTotalParkingCostOk() (*Price, bool)`

GetTotalParkingCostOk returns a tuple with the TotalParkingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParkingCost

`func (o *Cdr) SetTotalParkingCost(v Price)`

SetTotalParkingCost sets TotalParkingCost field to given value.

### HasTotalParkingCost

`func (o *Cdr) HasTotalParkingCost() bool`

HasTotalParkingCost returns a boolean if a field has been set.

### GetTotalReservationCost

`func (o *Cdr) GetTotalReservationCost() Price`

GetTotalReservationCost returns the TotalReservationCost field if non-nil, zero value otherwise.

### GetTotalReservationCostOk

`func (o *Cdr) GetTotalReservationCostOk() (*Price, bool)`

GetTotalReservationCostOk returns a tuple with the TotalReservationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservationCost

`func (o *Cdr) SetTotalReservationCost(v Price)`

SetTotalReservationCost sets TotalReservationCost field to given value.

### HasTotalReservationCost

`func (o *Cdr) HasTotalReservationCost() bool`

HasTotalReservationCost returns a boolean if a field has been set.

### GetRemark

`func (o *Cdr) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *Cdr) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *Cdr) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *Cdr) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetInvoiceReferenceId

`func (o *Cdr) GetInvoiceReferenceId() string`

GetInvoiceReferenceId returns the InvoiceReferenceId field if non-nil, zero value otherwise.

### GetInvoiceReferenceIdOk

`func (o *Cdr) GetInvoiceReferenceIdOk() (*string, bool)`

GetInvoiceReferenceIdOk returns a tuple with the InvoiceReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReferenceId

`func (o *Cdr) SetInvoiceReferenceId(v string)`

SetInvoiceReferenceId sets InvoiceReferenceId field to given value.

### HasInvoiceReferenceId

`func (o *Cdr) HasInvoiceReferenceId() bool`

HasInvoiceReferenceId returns a boolean if a field has been set.

### GetCredit

`func (o *Cdr) GetCredit() bool`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *Cdr) GetCreditOk() (*bool, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *Cdr) SetCredit(v bool)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *Cdr) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditReferenceId

`func (o *Cdr) GetCreditReferenceId() string`

GetCreditReferenceId returns the CreditReferenceId field if non-nil, zero value otherwise.

### GetCreditReferenceIdOk

`func (o *Cdr) GetCreditReferenceIdOk() (*string, bool)`

GetCreditReferenceIdOk returns a tuple with the CreditReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditReferenceId

`func (o *Cdr) SetCreditReferenceId(v string)`

SetCreditReferenceId sets CreditReferenceId field to given value.

### HasCreditReferenceId

`func (o *Cdr) HasCreditReferenceId() bool`

HasCreditReferenceId returns a boolean if a field has been set.

### GetHomeChargingCompensation

`func (o *Cdr) GetHomeChargingCompensation() bool`

GetHomeChargingCompensation returns the HomeChargingCompensation field if non-nil, zero value otherwise.

### GetHomeChargingCompensationOk

`func (o *Cdr) GetHomeChargingCompensationOk() (*bool, bool)`

GetHomeChargingCompensationOk returns a tuple with the HomeChargingCompensation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChargingCompensation

`func (o *Cdr) SetHomeChargingCompensation(v bool)`

SetHomeChargingCompensation sets HomeChargingCompensation field to given value.

### HasHomeChargingCompensation

`func (o *Cdr) HasHomeChargingCompensation() bool`

HasHomeChargingCompensation returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Cdr) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Cdr) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Cdr) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


