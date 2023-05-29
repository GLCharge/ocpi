# CdrTariffs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**Currency** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**TariffAltText** | Pointer to [**CdrTariffsTariffAltText**](CdrTariffsTariffAltText.md) |  | [optional] 
**TariffAltUrl** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**MaxPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**Elements** | Pointer to [**CdrTariffsElements**](CdrTariffsElements.md) |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**EnergyMix** | Pointer to [**CdrTariffsEnergyMix**](CdrTariffsEnergyMix.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewCdrTariffs

`func NewCdrTariffs(countryCode string, partyId string, id string, currency string, lastUpdated string, ) *CdrTariffs`

NewCdrTariffs instantiates a new CdrTariffs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrTariffsWithDefaults

`func NewCdrTariffsWithDefaults() *CdrTariffs`

NewCdrTariffsWithDefaults instantiates a new CdrTariffs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CdrTariffs) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CdrTariffs) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CdrTariffs) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *CdrTariffs) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *CdrTariffs) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *CdrTariffs) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *CdrTariffs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrTariffs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrTariffs) SetId(v string)`

SetId sets Id field to given value.


### GetCurrency

`func (o *CdrTariffs) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CdrTariffs) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CdrTariffs) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetType

`func (o *CdrTariffs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdrTariffs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdrTariffs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CdrTariffs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTariffAltText

`func (o *CdrTariffs) GetTariffAltText() CdrTariffsTariffAltText`

GetTariffAltText returns the TariffAltText field if non-nil, zero value otherwise.

### GetTariffAltTextOk

`func (o *CdrTariffs) GetTariffAltTextOk() (*CdrTariffsTariffAltText, bool)`

GetTariffAltTextOk returns a tuple with the TariffAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltText

`func (o *CdrTariffs) SetTariffAltText(v CdrTariffsTariffAltText)`

SetTariffAltText sets TariffAltText field to given value.

### HasTariffAltText

`func (o *CdrTariffs) HasTariffAltText() bool`

HasTariffAltText returns a boolean if a field has been set.

### GetTariffAltUrl

`func (o *CdrTariffs) GetTariffAltUrl() string`

GetTariffAltUrl returns the TariffAltUrl field if non-nil, zero value otherwise.

### GetTariffAltUrlOk

`func (o *CdrTariffs) GetTariffAltUrlOk() (*string, bool)`

GetTariffAltUrlOk returns a tuple with the TariffAltUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltUrl

`func (o *CdrTariffs) SetTariffAltUrl(v string)`

SetTariffAltUrl sets TariffAltUrl field to given value.

### HasTariffAltUrl

`func (o *CdrTariffs) HasTariffAltUrl() bool`

HasTariffAltUrl returns a boolean if a field has been set.

### GetMinPrice

`func (o *CdrTariffs) GetMinPrice() Price`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *CdrTariffs) GetMinPriceOk() (*Price, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *CdrTariffs) SetMinPrice(v Price)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *CdrTariffs) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *CdrTariffs) GetMaxPrice() Price`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *CdrTariffs) GetMaxPriceOk() (*Price, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *CdrTariffs) SetMaxPrice(v Price)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *CdrTariffs) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetElements

`func (o *CdrTariffs) GetElements() CdrTariffsElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CdrTariffs) GetElementsOk() (*CdrTariffsElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CdrTariffs) SetElements(v CdrTariffsElements)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CdrTariffs) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetStartDateTime

`func (o *CdrTariffs) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CdrTariffs) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CdrTariffs) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CdrTariffs) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *CdrTariffs) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CdrTariffs) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CdrTariffs) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CdrTariffs) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEnergyMix

`func (o *CdrTariffs) GetEnergyMix() CdrTariffsEnergyMix`

GetEnergyMix returns the EnergyMix field if non-nil, zero value otherwise.

### GetEnergyMixOk

`func (o *CdrTariffs) GetEnergyMixOk() (*CdrTariffsEnergyMix, bool)`

GetEnergyMixOk returns a tuple with the EnergyMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyMix

`func (o *CdrTariffs) SetEnergyMix(v CdrTariffsEnergyMix)`

SetEnergyMix sets EnergyMix field to given value.

### HasEnergyMix

`func (o *CdrTariffs) HasEnergyMix() bool`

HasEnergyMix returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CdrTariffs) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CdrTariffs) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CdrTariffs) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


