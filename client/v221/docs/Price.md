# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclVat** | **float32** |  | 
**InclVat** | Pointer to **float32** |  | [optional] 

## Methods

### NewPrice

`func NewPrice(exclVat float32, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclVat

`func (o *Price) GetExclVat() float32`

GetExclVat returns the ExclVat field if non-nil, zero value otherwise.

### GetExclVatOk

`func (o *Price) GetExclVatOk() (*float32, bool)`

GetExclVatOk returns a tuple with the ExclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclVat

`func (o *Price) SetExclVat(v float32)`

SetExclVat sets ExclVat field to given value.


### GetInclVat

`func (o *Price) GetInclVat() float32`

GetInclVat returns the InclVat field if non-nil, zero value otherwise.

### GetInclVatOk

`func (o *Price) GetInclVatOk() (*float32, bool)`

GetInclVatOk returns a tuple with the InclVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclVat

`func (o *Price) SetInclVat(v float32)`

SetInclVat sets InclVat field to given value.

### HasInclVat

`func (o *Price) HasInclVat() bool`

HasInclVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


