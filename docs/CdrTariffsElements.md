# CdrTariffsElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceComponents** | Pointer to **string** |  | [optional] 
**Restrictions** | Pointer to [**CdrTariffsElementsRestrictions**](CdrTariffsElementsRestrictions.md) |  | [optional] 

## Methods

### NewCdrTariffsElements

`func NewCdrTariffsElements() *CdrTariffsElements`

NewCdrTariffsElements instantiates a new CdrTariffsElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrTariffsElementsWithDefaults

`func NewCdrTariffsElementsWithDefaults() *CdrTariffsElements`

NewCdrTariffsElementsWithDefaults instantiates a new CdrTariffsElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceComponents

`func (o *CdrTariffsElements) GetPriceComponents() string`

GetPriceComponents returns the PriceComponents field if non-nil, zero value otherwise.

### GetPriceComponentsOk

`func (o *CdrTariffsElements) GetPriceComponentsOk() (*string, bool)`

GetPriceComponentsOk returns a tuple with the PriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceComponents

`func (o *CdrTariffsElements) SetPriceComponents(v string)`

SetPriceComponents sets PriceComponents field to given value.

### HasPriceComponents

`func (o *CdrTariffsElements) HasPriceComponents() bool`

HasPriceComponents returns a boolean if a field has been set.

### GetRestrictions

`func (o *CdrTariffsElements) GetRestrictions() CdrTariffsElementsRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *CdrTariffsElements) GetRestrictionsOk() (*CdrTariffsElementsRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *CdrTariffsElements) SetRestrictions(v CdrTariffsElementsRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *CdrTariffsElements) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


