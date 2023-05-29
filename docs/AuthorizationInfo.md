# AuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **string** |  | 
**Token** | [**Token**](Token.md) |  | 
**Location** | Pointer to [**LocationReferences**](LocationReferences.md) |  | [optional] 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**AuthorizationInfoInfo**](AuthorizationInfoInfo.md) |  | [optional] 

## Methods

### NewAuthorizationInfo

`func NewAuthorizationInfo(allowed string, token Token, ) *AuthorizationInfo`

NewAuthorizationInfo instantiates a new AuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationInfoWithDefaults

`func NewAuthorizationInfoWithDefaults() *AuthorizationInfo`

NewAuthorizationInfoWithDefaults instantiates a new AuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *AuthorizationInfo) GetAllowed() string`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AuthorizationInfo) GetAllowedOk() (*string, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AuthorizationInfo) SetAllowed(v string)`

SetAllowed sets Allowed field to given value.


### GetToken

`func (o *AuthorizationInfo) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthorizationInfo) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthorizationInfo) SetToken(v Token)`

SetToken sets Token field to given value.


### GetLocation

`func (o *AuthorizationInfo) GetLocation() LocationReferences`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AuthorizationInfo) GetLocationOk() (*LocationReferences, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AuthorizationInfo) SetLocation(v LocationReferences)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AuthorizationInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAuthorizationReference

`func (o *AuthorizationInfo) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *AuthorizationInfo) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *AuthorizationInfo) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *AuthorizationInfo) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetInfo

`func (o *AuthorizationInfo) GetInfo() AuthorizationInfoInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthorizationInfo) GetInfoOk() (*AuthorizationInfoInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthorizationInfo) SetInfo(v AuthorizationInfoInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AuthorizationInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


