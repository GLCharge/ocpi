# CredentialsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Url** | **string** |  | 
**Roles** | Pointer to [**CredentialsDataRoles**](CredentialsDataRoles.md) |  | [optional] 

## Methods

### NewCredentialsData

`func NewCredentialsData(token string, url string, ) *CredentialsData`

NewCredentialsData instantiates a new CredentialsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsDataWithDefaults

`func NewCredentialsDataWithDefaults() *CredentialsData`

NewCredentialsDataWithDefaults instantiates a new CredentialsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CredentialsData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CredentialsData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CredentialsData) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *CredentialsData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CredentialsData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CredentialsData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRoles

`func (o *CredentialsData) GetRoles() CredentialsDataRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CredentialsData) GetRolesOk() (*CredentialsDataRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CredentialsData) SetRoles(v CredentialsDataRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CredentialsData) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


