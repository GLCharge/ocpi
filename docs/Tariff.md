# Tariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Id** | **string** |  | 
**Currency** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**TariffAltText** | Pointer to [**TariffTariffAltText**](TariffTariffAltText.md) |  | [optional] 
**TariffAltUrl** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**MaxPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**Elements** | Pointer to [**TariffElements**](TariffElements.md) |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**EnergyMix** | Pointer to [**TariffEnergyMix**](TariffEnergyMix.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewTariff

`func NewTariff(countryCode string, partyId string, id string, currency string, lastUpdated string, ) *Tariff`

NewTariff instantiates a new Tariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffWithDefaults

`func NewTariffWithDefaults() *Tariff`

NewTariffWithDefaults instantiates a new Tariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Tariff) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Tariff) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Tariff) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *Tariff) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *Tariff) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *Tariff) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetId

`func (o *Tariff) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tariff) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tariff) SetId(v string)`

SetId sets Id field to given value.


### GetCurrency

`func (o *Tariff) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Tariff) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Tariff) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetType

`func (o *Tariff) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tariff) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tariff) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Tariff) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTariffAltText

`func (o *Tariff) GetTariffAltText() TariffTariffAltText`

GetTariffAltText returns the TariffAltText field if non-nil, zero value otherwise.

### GetTariffAltTextOk

`func (o *Tariff) GetTariffAltTextOk() (*TariffTariffAltText, bool)`

GetTariffAltTextOk returns a tuple with the TariffAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltText

`func (o *Tariff) SetTariffAltText(v TariffTariffAltText)`

SetTariffAltText sets TariffAltText field to given value.

### HasTariffAltText

`func (o *Tariff) HasTariffAltText() bool`

HasTariffAltText returns a boolean if a field has been set.

### GetTariffAltUrl

`func (o *Tariff) GetTariffAltUrl() string`

GetTariffAltUrl returns the TariffAltUrl field if non-nil, zero value otherwise.

### GetTariffAltUrlOk

`func (o *Tariff) GetTariffAltUrlOk() (*string, bool)`

GetTariffAltUrlOk returns a tuple with the TariffAltUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffAltUrl

`func (o *Tariff) SetTariffAltUrl(v string)`

SetTariffAltUrl sets TariffAltUrl field to given value.

### HasTariffAltUrl

`func (o *Tariff) HasTariffAltUrl() bool`

HasTariffAltUrl returns a boolean if a field has been set.

### GetMinPrice

`func (o *Tariff) GetMinPrice() Price`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *Tariff) GetMinPriceOk() (*Price, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *Tariff) SetMinPrice(v Price)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *Tariff) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *Tariff) GetMaxPrice() Price`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *Tariff) GetMaxPriceOk() (*Price, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *Tariff) SetMaxPrice(v Price)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *Tariff) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetElements

`func (o *Tariff) GetElements() TariffElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Tariff) GetElementsOk() (*TariffElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Tariff) SetElements(v TariffElements)`

SetElements sets Elements field to given value.

### HasElements

`func (o *Tariff) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetStartDateTime

`func (o *Tariff) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Tariff) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Tariff) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *Tariff) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *Tariff) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Tariff) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Tariff) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Tariff) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetEnergyMix

`func (o *Tariff) GetEnergyMix() TariffEnergyMix`

GetEnergyMix returns the EnergyMix field if non-nil, zero value otherwise.

### GetEnergyMixOk

`func (o *Tariff) GetEnergyMixOk() (*TariffEnergyMix, bool)`

GetEnergyMixOk returns a tuple with the EnergyMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyMix

`func (o *Tariff) SetEnergyMix(v TariffEnergyMix)`

SetEnergyMix sets EnergyMix field to given value.

### HasEnergyMix

`func (o *Tariff) HasEnergyMix() bool`

HasEnergyMix returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Tariff) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Tariff) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Tariff) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


