# BusinessDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Website** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to [**BusinessDetailsLogo**](BusinessDetailsLogo.md) |  | [optional] 

## Methods

### NewBusinessDetails

`func NewBusinessDetails(name string, ) *BusinessDetails`

NewBusinessDetails instantiates a new BusinessDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDetailsWithDefaults

`func NewBusinessDetailsWithDefaults() *BusinessDetails`

NewBusinessDetailsWithDefaults instantiates a new BusinessDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BusinessDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessDetails) SetName(v string)`

SetName sets Name field to given value.


### GetWebsite

`func (o *BusinessDetails) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BusinessDetails) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BusinessDetails) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BusinessDetails) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLogo

`func (o *BusinessDetails) GetLogo() BusinessDetailsLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BusinessDetails) GetLogoOk() (*BusinessDetailsLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BusinessDetails) SetLogo(v BusinessDetailsLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BusinessDetails) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


