# TariffElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceComponents** | Pointer to **string** |  | [optional] 
**Restrictions** | Pointer to [**TariffElementsRestrictions**](TariffElementsRestrictions.md) |  | [optional] 

## Methods

### NewTariffElements

`func NewTariffElements() *TariffElements`

NewTariffElements instantiates a new TariffElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffElementsWithDefaults

`func NewTariffElementsWithDefaults() *TariffElements`

NewTariffElementsWithDefaults instantiates a new TariffElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceComponents

`func (o *TariffElements) GetPriceComponents() string`

GetPriceComponents returns the PriceComponents field if non-nil, zero value otherwise.

### GetPriceComponentsOk

`func (o *TariffElements) GetPriceComponentsOk() (*string, bool)`

GetPriceComponentsOk returns a tuple with the PriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceComponents

`func (o *TariffElements) SetPriceComponents(v string)`

SetPriceComponents sets PriceComponents field to given value.

### HasPriceComponents

`func (o *TariffElements) HasPriceComponents() bool`

HasPriceComponents returns a boolean if a field has been set.

### GetRestrictions

`func (o *TariffElements) GetRestrictions() TariffElementsRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *TariffElements) GetRestrictionsOk() (*TariffElementsRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *TariffElements) SetRestrictions(v TariffElementsRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *TariffElements) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


