# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**StartDateTime** | **string** |  | 
**EndDateTime** | Pointer to **string** |  | [optional] 
**Kwh** | **float32** |  | 
**CdrToken** | [**SessionCdrToken**](SessionCdrToken.md) |  | 
**AuthMethod** | **string** |  | 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**LocationId** | **string** |  | 
**EvseUid** | **string** |  | 
**ConnectorId** | **string** |  | 
**MeterId** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**ChargingPeriods** | Pointer to [**SessionChargingPeriods**](SessionChargingPeriods.md) |  | [optional] 
**TotalCosts** | Pointer to [**SessionTotalCosts**](SessionTotalCosts.md) |  | [optional] 
**Status** | **string** |  | 
**LastUpdated** | **string** |  | 

## Methods

### NewSession

`func NewSession(countryCode string, partyId string, id string, startDateTime string, kwh float32, cdrToken SessionCdrToken, authMethod string, locationId string, evseUid string, connectorId string, currency string, status string, lastUpdated string, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Session) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Session) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Session) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *Session) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *Session) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *Session) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetStartDateTime

`func (o *Session) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Session) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Session) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *Session) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Session) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Session) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Session) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetKwh

`func (o *Session) GetKwh() float32`

GetKwh returns the Kwh field if non-nil, zero value otherwise.

### GetKwhOk

`func (o *Session) GetKwhOk() (*float32, bool)`

GetKwhOk returns a tuple with the Kwh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwh

`func (o *Session) SetKwh(v float32)`

SetKwh sets Kwh field to given value.


### GetCdrToken

`func (o *Session) GetCdrToken() SessionCdrToken`

GetCdrToken returns the CdrToken field if non-nil, zero value otherwise.

### GetCdrTokenOk

`func (o *Session) GetCdrTokenOk() (*SessionCdrToken, bool)`

GetCdrTokenOk returns a tuple with the CdrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrToken

`func (o *Session) SetCdrToken(v SessionCdrToken)`

SetCdrToken sets CdrToken field to given value.


### GetAuthMethod

`func (o *Session) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Session) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Session) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetAuthorizationReference

`func (o *Session) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *Session) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *Session) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *Session) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetLocationId

`func (o *Session) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Session) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Session) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetEvseUid

`func (o *Session) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *Session) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *Session) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.


### GetConnectorId

`func (o *Session) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *Session) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *Session) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetMeterId

`func (o *Session) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *Session) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *Session) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *Session) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetCurrency

`func (o *Session) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Session) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Session) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChargingPeriods

`func (o *Session) GetChargingPeriods() SessionChargingPeriods`

GetChargingPeriods returns the ChargingPeriods field if non-nil, zero value otherwise.

### GetChargingPeriodsOk

`func (o *Session) GetChargingPeriodsOk() (*SessionChargingPeriods, bool)`

GetChargingPeriodsOk returns a tuple with the ChargingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingPeriods

`func (o *Session) SetChargingPeriods(v SessionChargingPeriods)`

SetChargingPeriods sets ChargingPeriods field to given value.

### HasChargingPeriods

`func (o *Session) HasChargingPeriods() bool`

HasChargingPeriods returns a boolean if a field has been set.

### GetTotalCosts

`func (o *Session) GetTotalCosts() SessionTotalCosts`

GetTotalCosts returns the TotalCosts field if non-nil, zero value otherwise.

### GetTotalCostsOk

`func (o *Session) GetTotalCostsOk() (*SessionTotalCosts, bool)`

GetTotalCostsOk returns a tuple with the TotalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCosts

`func (o *Session) SetTotalCosts(v SessionTotalCosts)`

SetTotalCosts sets TotalCosts field to given value.

### HasTotalCosts

`func (o *Session) HasTotalCosts() bool`

HasTotalCosts returns a boolean if a field has been set.

### GetStatus

`func (o *Session) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Session) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Session) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastUpdated

`func (o *Session) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Session) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Session) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


