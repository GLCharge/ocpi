# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **string** |  | 
**Token** | [**Token**](Token.md) |  | 
**Location** | Pointer to [**LocationReferences**](LocationReferences.md) |  | [optional] 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**CommandResponseMessage**](CommandResponseMessage.md) |  | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization(allowed string, token Token, ) *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *Authorization) GetAllowed() string`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *Authorization) GetAllowedOk() (*string, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *Authorization) SetAllowed(v string)`

SetAllowed sets Allowed field to given value.


### GetToken

`func (o *Authorization) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Authorization) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Authorization) SetToken(v Token)`

SetToken sets Token field to given value.


### GetLocation

`func (o *Authorization) GetLocation() LocationReferences`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Authorization) GetLocationOk() (*LocationReferences, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Authorization) SetLocation(v LocationReferences)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Authorization) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAuthorizationReference

`func (o *Authorization) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *Authorization) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *Authorization) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *Authorization) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetInfo

`func (o *Authorization) GetInfo() CommandResponseMessage`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Authorization) GetInfoOk() (*CommandResponseMessage, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Authorization) SetInfo(v CommandResponseMessage)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Authorization) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


