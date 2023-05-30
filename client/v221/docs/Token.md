# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**PartyId** | **string** |  | 
**Uid** | **string** |  | 
**Type** | **string** |  | 
**ContractId** | **string** |  | 
**VisualNumber** | Pointer to **string** |  | [optional] 
**Issuer** | **string** |  | 
**GroupId** | Pointer to **string** |  | [optional] 
**Valid** | **bool** |  | 
**Whitelist** | **string** |  | 
**Language** | Pointer to **string** |  | [optional] 
**DefaultProfileType** | Pointer to **string** |  | [optional] 
**EnergyContract** | Pointer to [**TokenEnergyContract**](TokenEnergyContract.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewToken

`func NewToken(countryCode string, partyId string, uid string, type_ string, contractId string, issuer string, valid bool, whitelist string, lastUpdated string, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Token) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Token) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Token) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPartyId

`func (o *Token) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *Token) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *Token) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetUid

`func (o *Token) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Token) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Token) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *Token) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Token) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Token) SetType(v string)`

SetType sets Type field to given value.


### GetContractId

`func (o *Token) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Token) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Token) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetVisualNumber

`func (o *Token) GetVisualNumber() string`

GetVisualNumber returns the VisualNumber field if non-nil, zero value otherwise.

### GetVisualNumberOk

`func (o *Token) GetVisualNumberOk() (*string, bool)`

GetVisualNumberOk returns a tuple with the VisualNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualNumber

`func (o *Token) SetVisualNumber(v string)`

SetVisualNumber sets VisualNumber field to given value.

### HasVisualNumber

`func (o *Token) HasVisualNumber() bool`

HasVisualNumber returns a boolean if a field has been set.

### GetIssuer

`func (o *Token) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Token) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Token) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetGroupId

`func (o *Token) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Token) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Token) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Token) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetValid

`func (o *Token) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Token) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Token) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetWhitelist

`func (o *Token) GetWhitelist() string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *Token) GetWhitelistOk() (*string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *Token) SetWhitelist(v string)`

SetWhitelist sets Whitelist field to given value.


### GetLanguage

`func (o *Token) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Token) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Token) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Token) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetDefaultProfileType

`func (o *Token) GetDefaultProfileType() string`

GetDefaultProfileType returns the DefaultProfileType field if non-nil, zero value otherwise.

### GetDefaultProfileTypeOk

`func (o *Token) GetDefaultProfileTypeOk() (*string, bool)`

GetDefaultProfileTypeOk returns a tuple with the DefaultProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfileType

`func (o *Token) SetDefaultProfileType(v string)`

SetDefaultProfileType sets DefaultProfileType field to given value.

### HasDefaultProfileType

`func (o *Token) HasDefaultProfileType() bool`

HasDefaultProfileType returns a boolean if a field has been set.

### GetEnergyContract

`func (o *Token) GetEnergyContract() TokenEnergyContract`

GetEnergyContract returns the EnergyContract field if non-nil, zero value otherwise.

### GetEnergyContractOk

`func (o *Token) GetEnergyContractOk() (*TokenEnergyContract, bool)`

GetEnergyContractOk returns a tuple with the EnergyContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyContract

`func (o *Token) SetEnergyContract(v TokenEnergyContract)`

SetEnergyContract sets EnergyContract field to given value.

### HasEnergyContract

`func (o *Token) HasEnergyContract() bool`

HasEnergyContract returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Token) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Token) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Token) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


