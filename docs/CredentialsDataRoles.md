# CredentialsDataRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**BusinessDetails** | [**CredentialsDataRolesBusinessDetails**](CredentialsDataRolesBusinessDetails.md) |  | 
**PartyId** | **string** |  | 
**CountryCode** | **string** |  | 

## Methods

### NewCredentialsDataRoles

`func NewCredentialsDataRoles(role string, businessDetails CredentialsDataRolesBusinessDetails, partyId string, countryCode string, ) *CredentialsDataRoles`

NewCredentialsDataRoles instantiates a new CredentialsDataRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsDataRolesWithDefaults

`func NewCredentialsDataRolesWithDefaults() *CredentialsDataRoles`

NewCredentialsDataRolesWithDefaults instantiates a new CredentialsDataRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *CredentialsDataRoles) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CredentialsDataRoles) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CredentialsDataRoles) SetRole(v string)`

SetRole sets Role field to given value.


### GetBusinessDetails

`func (o *CredentialsDataRoles) GetBusinessDetails() CredentialsDataRolesBusinessDetails`

GetBusinessDetails returns the BusinessDetails field if non-nil, zero value otherwise.

### GetBusinessDetailsOk

`func (o *CredentialsDataRoles) GetBusinessDetailsOk() (*CredentialsDataRolesBusinessDetails, bool)`

GetBusinessDetailsOk returns a tuple with the BusinessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDetails

`func (o *CredentialsDataRoles) SetBusinessDetails(v CredentialsDataRolesBusinessDetails)`

SetBusinessDetails sets BusinessDetails field to given value.


### GetPartyId

`func (o *CredentialsDataRoles) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *CredentialsDataRoles) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *CredentialsDataRoles) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.


### GetCountryCode

`func (o *CredentialsDataRoles) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CredentialsDataRoles) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CredentialsDataRoles) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


