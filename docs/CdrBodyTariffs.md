# CdrBodyTariffs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**Currency** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**TariffAltText** | Pointer to [**CdrBodyTariffsTariffAltText**](CdrBodyTariffsTariffAltText.md) |  | [optional] 
**TariffAltUrl** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**MaxPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**Elements** | Pointer to [**CdrBodyTariffsElements**](CdrBodyTariffsElements.md) |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**EnergyMix** | Pointer to [**CdrBodyTariffsEnergyMix**](CdrBodyTariffsEnergyMix.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewCdrBodyTariffs

`func NewCdrBodyTariffs(countryCode string, partyId string, id string, currency string, lastUpdated string, ) *CdrBodyTariffs`

NewCdrBodyTariffs instantiates a new CdrBodyTariffs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrBodyTariffsWithDefaults

`func NewCdrBodyTariffsWithDefaults() *CdrBodyTariffs`

NewCdrBodyTariffsWithDefaults instantiates a new CdrBodyTariffs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CdrBodyTariffs) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CdrBodyTariffs) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CdrBodyTariffs) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *CdrBodyTariffs) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *CdrBodyTariffs) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *CdrBodyTariffs) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *CdrBodyTariffs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrBodyTariffs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrBodyTariffs) SetId(v string)`

SetId sets Id field to given value.


### GetCurrency

`func (o *CdrBodyTariffs) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CdrBodyTariffs) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CdrBodyTariffs) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetType

`func (o *CdrBodyTariffs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdrBodyTariffs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdrBodyTariffs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CdrBodyTariffs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTariffAltText

`func (o *CdrBodyTariffs) GetTariffAltText() CdrBodyTariffsTariffAltText`

GetTariffAltText returns the TariffAltText field if non-nil, zero value otherwise.

### GetTariffAltTextOk

`func (o *CdrBodyTariffs) GetTariffAltTextOk() (*CdrBodyTariffsTariffAltText, bool)`

GetTariffAltTextOk returns a tuple with the TariffAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltText

`func (o *CdrBodyTariffs) SetTariffAltText(v CdrBodyTariffsTariffAltText)`

SetTariffAltText sets TariffAltText field to given value.

### HasTariffAltText

`func (o *CdrBodyTariffs) HasTariffAltText() bool`

HasTariffAltText returns a boolean if a field has been set.

### GetTariffAltUrl

`func (o *CdrBodyTariffs) GetTariffAltUrl() string`

GetTariffAltUrl returns the TariffAltUrl field if non-nil, zero value otherwise.

### GetTariffAltUrlOk

`func (o *CdrBodyTariffs) GetTariffAltUrlOk() (*string, bool)`

GetTariffAltUrlOk returns a tuple with the TariffAltUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltUrl

`func (o *CdrBodyTariffs) SetTariffAltUrl(v string)`

SetTariffAltUrl sets TariffAltUrl field to given value.

### HasTariffAltUrl

`func (o *CdrBodyTariffs) HasTariffAltUrl() bool`

HasTariffAltUrl returns a boolean if a field has been set.

### GetMinPrice

`func (o *CdrBodyTariffs) GetMinPrice() Price`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *CdrBodyTariffs) GetMinPriceOk() (*Price, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *CdrBodyTariffs) SetMinPrice(v Price)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *CdrBodyTariffs) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *CdrBodyTariffs) GetMaxPrice() Price`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *CdrBodyTariffs) GetMaxPriceOk() (*Price, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *CdrBodyTariffs) SetMaxPrice(v Price)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *CdrBodyTariffs) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetElements

`func (o *CdrBodyTariffs) GetElements() CdrBodyTariffsElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CdrBodyTariffs) GetElementsOk() (*CdrBodyTariffsElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CdrBodyTariffs) SetElements(v CdrBodyTariffsElements)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CdrBodyTariffs) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetStartDateTime

`func (o *CdrBodyTariffs) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CdrBodyTariffs) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CdrBodyTariffs) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CdrBodyTariffs) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *CdrBodyTariffs) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CdrBodyTariffs) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CdrBodyTariffs) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CdrBodyTariffs) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEnergyMix

`func (o *CdrBodyTariffs) GetEnergyMix() CdrBodyTariffsEnergyMix`

GetEnergyMix returns the EnergyMix field if non-nil, zero value otherwise.

### GetEnergyMixOk

`func (o *CdrBodyTariffs) GetEnergyMixOk() (*CdrBodyTariffsEnergyMix, bool)`

GetEnergyMixOk returns a tuple with the EnergyMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyMix

`func (o *CdrBodyTariffs) SetEnergyMix(v CdrBodyTariffsEnergyMix)`

SetEnergyMix sets EnergyMix field to given value.

### HasEnergyMix

`func (o *CdrBodyTariffs) HasEnergyMix() bool`

HasEnergyMix returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CdrBodyTariffs) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CdrBodyTariffs) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CdrBodyTariffs) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


