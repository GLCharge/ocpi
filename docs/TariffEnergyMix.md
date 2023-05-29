# TariffEnergyMix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsGreenEnergy** | **bool** |  | 
**EnergySources** | Pointer to [**TariffEnergyMixEnergySources**](TariffEnergyMixEnergySources.md) |  | [optional] 
**EnvironImpact** | Pointer to [**TariffEnergyMixEnvironImpact**](TariffEnergyMixEnvironImpact.md) |  | [optional] 
**SupplierName** | Pointer to **string** |  | [optional] 
**EnergyProductName** | Pointer to **string** |  | [optional] 

## Methods

### NewTariffEnergyMix

`func NewTariffEnergyMix(isGreenEnergy bool, ) *TariffEnergyMix`

NewTariffEnergyMix instantiates a new TariffEnergyMix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffEnergyMixWithDefaults

`func NewTariffEnergyMixWithDefaults() *TariffEnergyMix`

NewTariffEnergyMixWithDefaults instantiates a new TariffEnergyMix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsGreenEnergy

`func (o *TariffEnergyMix) GetIsGreenEnergy() bool`

GetIsGreenEnergy returns the IsGreenEnergy field if non-nil, zero value otherwise.

### GetIsGreenEnergyOk

`func (o *TariffEnergyMix) GetIsGreenEnergyOk() (*bool, bool)`

GetIsGreenEnergyOk returns a tuple with the IsGreenEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGreenEnergy

`func (o *TariffEnergyMix) SetIsGreenEnergy(v bool)`

SetIsGreenEnergy sets IsGreenEnergy field to given value.


### GetEnergySources

`func (o *TariffEnergyMix) GetEnergySources() TariffEnergyMixEnergySources`

GetEnergySources returns the EnergySources field if non-nil, zero value otherwise.

### GetEnergySourcesOk

`func (o *TariffEnergyMix) GetEnergySourcesOk() (*TariffEnergyMixEnergySources, bool)`

GetEnergySourcesOk returns a tuple with the EnergySources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergySources

`func (o *TariffEnergyMix) SetEnergySources(v TariffEnergyMixEnergySources)`

SetEnergySources sets EnergySources field to given value.

### HasEnergySources

`func (o *TariffEnergyMix) HasEnergySources() bool`

HasEnergySources returns a boolean if a field has been set.

### GetEnvironImpact

`func (o *TariffEnergyMix) GetEnvironImpact() TariffEnergyMixEnvironImpact`

GetEnvironImpact returns the EnvironImpact field if non-nil, zero value otherwise.

### GetEnvironImpactOk

`func (o *TariffEnergyMix) GetEnvironImpactOk() (*TariffEnergyMixEnvironImpact, bool)`

GetEnvironImpactOk returns a tuple with the EnvironImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironImpact

`func (o *TariffEnergyMix) SetEnvironImpact(v TariffEnergyMixEnvironImpact)`

SetEnvironImpact sets EnvironImpact field to given value.

### HasEnvironImpact

`func (o *TariffEnergyMix) HasEnvironImpact() bool`

HasEnvironImpact returns a boolean if a field has been set.

### GetSupplierName

`func (o *TariffEnergyMix) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *TariffEnergyMix) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *TariffEnergyMix) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *TariffEnergyMix) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### GetEnergyProductName

`func (o *TariffEnergyMix) GetEnergyProductName() string`

GetEnergyProductName returns the EnergyProductName field if non-nil, zero value otherwise.

### GetEnergyProductNameOk

`func (o *TariffEnergyMix) GetEnergyProductNameOk() (*string, bool)`

GetEnergyProductNameOk returns a tuple with the EnergyProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyProductName

`func (o *TariffEnergyMix) SetEnergyProductName(v string)`

SetEnergyProductName sets EnergyProductName field to given value.

### HasEnergyProductName

`func (o *TariffEnergyMix) HasEnergyProductName() bool`

HasEnergyProductName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


