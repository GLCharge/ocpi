# CdrBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**StartDateTime** | **string** |  | 
**EndDateTime** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**CdrToken** | [**CdrBodyCdrToken**](CdrBodyCdrToken.md) |  | 
**AuthMethod** | **string** |  | 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**CdrLocation** | [**CdrBodyCdrLocation**](CdrBodyCdrLocation.md) |  | 
**MeterId** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**Tariffs** | Pointer to [**CdrBodyTariffs**](CdrBodyTariffs.md) |  | [optional] 
**ChargingPeriods** | Pointer to [**CdrBodyChargingPeriods**](CdrBodyChargingPeriods.md) |  | [optional] 
**SignedData** | Pointer to [**CdrBodySignedData**](CdrBodySignedData.md) |  | [optional] 
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

### NewCdrBody

`func NewCdrBody(countryCode string, partyId string, id string, startDateTime string, endDateTime string, cdrToken CdrBodyCdrToken, authMethod string, cdrLocation CdrBodyCdrLocation, currency string, totalCost Price, totalEnergy float32, totalTime float32, lastUpdated string, ) *CdrBody`

NewCdrBody instantiates a new CdrBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrBodyWithDefaults

`func NewCdrBodyWithDefaults() *CdrBody`

NewCdrBodyWithDefaults instantiates a new CdrBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CdrBody) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CdrBody) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CdrBody) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *CdrBody) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *CdrBody) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *CdrBody) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *CdrBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrBody) SetId(v string)`

SetId sets Id field to given value.


### GetStartDateTime

`func (o *CdrBody) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CdrBody) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CdrBody) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *CdrBody) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CdrBody) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CdrBody) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.


### GetSessionId

`func (o *CdrBody) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CdrBody) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CdrBody) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CdrBody) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCdrToken

`func (o *CdrBody) GetCdrToken() CdrBodyCdrToken`

GetCdrToken returns the CdrToken field if non-nil, zero value otherwise.

### GetCdrTokenOk

`func (o *CdrBody) GetCdrTokenOk() (*CdrBodyCdrToken, bool)`

GetCdrTokenOk returns a tuple with the CdrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrToken

`func (o *CdrBody) SetCdrToken(v CdrBodyCdrToken)`

SetCdrToken sets CdrToken field to given value.


### GetAuthMethod

`func (o *CdrBody) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CdrBody) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CdrBody) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetAuthorizationReference

`func (o *CdrBody) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *CdrBody) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *CdrBody) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *CdrBody) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetCdrLocation

`func (o *CdrBody) GetCdrLocation() CdrBodyCdrLocation`

GetCdrLocation returns the CdrLocation field if non-nil, zero value otherwise.

### GetCdrLocationOk

`func (o *CdrBody) GetCdrLocationOk() (*CdrBodyCdrLocation, bool)`

GetCdrLocationOk returns a tuple with the CdrLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrLocation

`func (o *CdrBody) SetCdrLocation(v CdrBodyCdrLocation)`

SetCdrLocation sets CdrLocation field to given value.


### GetMeterId

`func (o *CdrBody) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *CdrBody) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *CdrBody) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *CdrBody) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetCurrency

`func (o *CdrBody) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CdrBody) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CdrBody) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTariffs

`func (o *CdrBody) GetTariffs() CdrBodyTariffs`

GetTariffs returns the Tariffs field if non-nil, zero value otherwise.

### GetTariffsOk

`func (o *CdrBody) GetTariffsOk() (*CdrBodyTariffs, bool)`

GetTariffsOk returns a tuple with the Tariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffs

`func (o *CdrBody) SetTariffs(v CdrBodyTariffs)`

SetTariffs sets Tariffs field to given value.

### HasTariffs

`func (o *CdrBody) HasTariffs() bool`

HasTariffs returns a boolean if a field has been set.

### GetChargingPeriods

`func (o *CdrBody) GetChargingPeriods() CdrBodyChargingPeriods`

GetChargingPeriods returns the ChargingPeriods field if non-nil, zero value otherwise.

### GetChargingPeriodsOk

`func (o *CdrBody) GetChargingPeriodsOk() (*CdrBodyChargingPeriods, bool)`

GetChargingPeriodsOk returns a tuple with the ChargingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingPeriods

`func (o *CdrBody) SetChargingPeriods(v CdrBodyChargingPeriods)`

SetChargingPeriods sets ChargingPeriods field to given value.

### HasChargingPeriods

`func (o *CdrBody) HasChargingPeriods() bool`

HasChargingPeriods returns a boolean if a field has been set.

### GetSignedData

`func (o *CdrBody) GetSignedData() CdrBodySignedData`

GetSignedData returns the SignedData field if non-nil, zero value otherwise.

### GetSignedDataOk

`func (o *CdrBody) GetSignedDataOk() (*CdrBodySignedData, bool)`

GetSignedDataOk returns a tuple with the SignedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedData

`func (o *CdrBody) SetSignedData(v CdrBodySignedData)`

SetSignedData sets SignedData field to given value.

### HasSignedData

`func (o *CdrBody) HasSignedData() bool`

HasSignedData returns a boolean if a field has been set.

### GetTotalCost

`func (o *CdrBody) GetTotalCost() Price`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *CdrBody) GetTotalCostOk() (*Price, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *CdrBody) SetTotalCost(v Price)`

SetTotalCost sets TotalCost field to given value.


### GetTotalFixedCost

`func (o *CdrBody) GetTotalFixedCost() Price`

GetTotalFixedCost returns the TotalFixedCost field if non-nil, zero value otherwise.

### GetTotalFixedCostOk

`func (o *CdrBody) GetTotalFixedCostOk() (*Price, bool)`

GetTotalFixedCostOk returns a tuple with the TotalFixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFixedCost

`func (o *CdrBody) SetTotalFixedCost(v Price)`

SetTotalFixedCost sets TotalFixedCost field to given value.

### HasTotalFixedCost

`func (o *CdrBody) HasTotalFixedCost() bool`

HasTotalFixedCost returns a boolean if a field has been set.

### GetTotalEnergy

`func (o *CdrBody) GetTotalEnergy() float32`

GetTotalEnergy returns the TotalEnergy field if non-nil, zero value otherwise.

### GetTotalEnergyOk

`func (o *CdrBody) GetTotalEnergyOk() (*float32, bool)`

GetTotalEnergyOk returns a tuple with the TotalEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnergy

`func (o *CdrBody) SetTotalEnergy(v float32)`

SetTotalEnergy sets TotalEnergy field to given value.


### GetTotalEnergyCost

`func (o *CdrBody) GetTotalEnergyCost() Price`

GetTotalEnergyCost returns the TotalEnergyCost field if non-nil, zero value otherwise.

### GetTotalEnergyCostOk

`func (o *CdrBody) GetTotalEnergyCostOk() (*Price, bool)`

GetTotalEnergyCostOk returns a tuple with the TotalEnergyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnergyCost

`func (o *CdrBody) SetTotalEnergyCost(v Price)`

SetTotalEnergyCost sets TotalEnergyCost field to given value.

### HasTotalEnergyCost

`func (o *CdrBody) HasTotalEnergyCost() bool`

HasTotalEnergyCost returns a boolean if a field has been set.

### GetTotalTime

`func (o *CdrBody) GetTotalTime() float32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *CdrBody) GetTotalTimeOk() (*float32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *CdrBody) SetTotalTime(v float32)`

SetTotalTime sets TotalTime field to given value.


### GetTotalTimeCost

`func (o *CdrBody) GetTotalTimeCost() Price`

GetTotalTimeCost returns the TotalTimeCost field if non-nil, zero value otherwise.

### GetTotalTimeCostOk

`func (o *CdrBody) GetTotalTimeCostOk() (*Price, bool)`

GetTotalTimeCostOk returns a tuple with the TotalTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeCost

`func (o *CdrBody) SetTotalTimeCost(v Price)`

SetTotalTimeCost sets TotalTimeCost field to given value.

### HasTotalTimeCost

`func (o *CdrBody) HasTotalTimeCost() bool`

HasTotalTimeCost returns a boolean if a field has been set.

### GetTotalParkingTime

`func (o *CdrBody) GetTotalParkingTime() float32`

GetTotalParkingTime returns the TotalParkingTime field if non-nil, zero value otherwise.

### GetTotalParkingTimeOk

`func (o *CdrBody) GetTotalParkingTimeOk() (*float32, bool)`

GetTotalParkingTimeOk returns a tuple with the TotalParkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParkingTime

`func (o *CdrBody) SetTotalParkingTime(v float32)`

SetTotalParkingTime sets TotalParkingTime field to given value.

### HasTotalParkingTime

`func (o *CdrBody) HasTotalParkingTime() bool`

HasTotalParkingTime returns a boolean if a field has been set.

### GetTotalParkingCost

`func (o *CdrBody) GetTotalParkingCost() Price`

GetTotalParkingCost returns the TotalParkingCost field if non-nil, zero value otherwise.

### GetTotalParkingCostOk

`func (o *CdrBody) GetTotalParkingCostOk() (*Price, bool)`

GetTotalParkingCostOk returns a tuple with the TotalParkingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParkingCost

`func (o *CdrBody) SetTotalParkingCost(v Price)`

SetTotalParkingCost sets TotalParkingCost field to given value.

### HasTotalParkingCost

`func (o *CdrBody) HasTotalParkingCost() bool`

HasTotalParkingCost returns a boolean if a field has been set.

### GetTotalReservationCost

`func (o *CdrBody) GetTotalReservationCost() Price`

GetTotalReservationCost returns the TotalReservationCost field if non-nil, zero value otherwise.

### GetTotalReservationCostOk

`func (o *CdrBody) GetTotalReservationCostOk() (*Price, bool)`

GetTotalReservationCostOk returns a tuple with the TotalReservationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservationCost

`func (o *CdrBody) SetTotalReservationCost(v Price)`

SetTotalReservationCost sets TotalReservationCost field to given value.

### HasTotalReservationCost

`func (o *CdrBody) HasTotalReservationCost() bool`

HasTotalReservationCost returns a boolean if a field has been set.

### GetRemark

`func (o *CdrBody) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CdrBody) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CdrBody) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CdrBody) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetInvoiceReferenceId

`func (o *CdrBody) GetInvoiceReferenceId() string`

GetInvoiceReferenceId returns the InvoiceReferenceId field if non-nil, zero value otherwise.

### GetInvoiceReferenceIdOk

`func (o *CdrBody) GetInvoiceReferenceIdOk() (*string, bool)`

GetInvoiceReferenceIdOk returns a tuple with the InvoiceReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReferenceId

`func (o *CdrBody) SetInvoiceReferenceId(v string)`

SetInvoiceReferenceId sets InvoiceReferenceId field to given value.

### HasInvoiceReferenceId

`func (o *CdrBody) HasInvoiceReferenceId() bool`

HasInvoiceReferenceId returns a boolean if a field has been set.

### GetCredit

`func (o *CdrBody) GetCredit() bool`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *CdrBody) GetCreditOk() (*bool, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *CdrBody) SetCredit(v bool)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *CdrBody) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditReferenceId

`func (o *CdrBody) GetCreditReferenceId() string`

GetCreditReferenceId returns the CreditReferenceId field if non-nil, zero value otherwise.

### GetCreditReferenceIdOk

`func (o *CdrBody) GetCreditReferenceIdOk() (*string, bool)`

GetCreditReferenceIdOk returns a tuple with the CreditReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditReferenceId

`func (o *CdrBody) SetCreditReferenceId(v string)`

SetCreditReferenceId sets CreditReferenceId field to given value.

### HasCreditReferenceId

`func (o *CdrBody) HasCreditReferenceId() bool`

HasCreditReferenceId returns a boolean if a field has been set.

### GetHomeChargingCompensation

`func (o *CdrBody) GetHomeChargingCompensation() bool`

GetHomeChargingCompensation returns the HomeChargingCompensation field if non-nil, zero value otherwise.

### GetHomeChargingCompensationOk

`func (o *CdrBody) GetHomeChargingCompensationOk() (*bool, bool)`

GetHomeChargingCompensationOk returns a tuple with the HomeChargingCompensation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChargingCompensation

`func (o *CdrBody) SetHomeChargingCompensation(v bool)`

SetHomeChargingCompensation sets HomeChargingCompensation field to given value.

### HasHomeChargingCompensation

`func (o *CdrBody) HasHomeChargingCompensation() bool`

HasHomeChargingCompensation returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CdrBody) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CdrBody) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CdrBody) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


