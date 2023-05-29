# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Standard** | **string** |  | 
**Format** | **string** |  | 
**PowerType** | **string** |  | 
**MaxVoltage** | **int32** |  | 
**MaxAmperage** | **int32** |  | 
**MaxElectricPower** | Pointer to **int32** |  | [optional] 
**TariffIds** | Pointer to **string** |  | [optional] 
**TermsAndConditions** | Pointer to **string** |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewConnector

`func NewConnector(id string, standard string, format string, powerType string, maxVoltage int32, maxAmperage int32, lastUpdated string, ) *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v string)`

SetId sets Id field to given value.


### GetStandard

`func (o *Connector) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Connector) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Connector) SetStandard(v string)`

SetStandard sets Standard field to given value.


### GetFormat

`func (o *Connector) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Connector) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Connector) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetPowerType

`func (o *Connector) GetPowerType() string`

GetPowerType returns the PowerType field if non-nil, zero value otherwise.

### GetPowerTypeOk

`func (o *Connector) GetPowerTypeOk() (*string, bool)`

GetPowerTypeOk returns a tuple with the PowerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerType

`func (o *Connector) SetPowerType(v string)`

SetPowerType sets PowerType field to given value.


### GetMaxVoltage

`func (o *Connector) GetMaxVoltage() int32`

GetMaxVoltage returns the MaxVoltage field if non-nil, zero value otherwise.

### GetMaxVoltageOk

`func (o *Connector) GetMaxVoltageOk() (*int32, bool)`

GetMaxVoltageOk returns a tuple with the MaxVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVoltage

`func (o *Connector) SetMaxVoltage(v int32)`

SetMaxVoltage sets MaxVoltage field to given value.


### GetMaxAmperage

`func (o *Connector) GetMaxAmperage() int32`

GetMaxAmperage returns the MaxAmperage field if non-nil, zero value otherwise.

### GetMaxAmperageOk

`func (o *Connector) GetMaxAmperageOk() (*int32, bool)`

GetMaxAmperageOk returns a tuple with the MaxAmperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmperage

`func (o *Connector) SetMaxAmperage(v int32)`

SetMaxAmperage sets MaxAmperage field to given value.


### GetMaxElectricPower

`func (o *Connector) GetMaxElectricPower() int32`

GetMaxElectricPower returns the MaxElectricPower field if non-nil, zero value otherwise.

### GetMaxElectricPowerOk

`func (o *Connector) GetMaxElectricPowerOk() (*int32, bool)`

GetMaxElectricPowerOk returns a tuple with the MaxElectricPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElectricPower

`func (o *Connector) SetMaxElectricPower(v int32)`

SetMaxElectricPower sets MaxElectricPower field to given value.

### HasMaxElectricPower

`func (o *Connector) HasMaxElectricPower() bool`

HasMaxElectricPower returns a boolean if a field has been set.

### GetTariffIds

`func (o *Connector) GetTariffIds() string`

GetTariffIds returns the TariffIds field if non-nil, zero value otherwise.

### GetTariffIdsOk

`func (o *Connector) GetTariffIdsOk() (*string, bool)`

GetTariffIdsOk returns a tuple with the TariffIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffIds

`func (o *Connector) SetTariffIds(v string)`

SetTariffIds sets TariffIds field to given value.

### HasTariffIds

`func (o *Connector) HasTariffIds() bool`

HasTariffIds returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *Connector) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *Connector) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *Connector) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *Connector) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Connector) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Connector) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Connector) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


