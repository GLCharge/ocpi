# StartSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseUrl** | **string** |  | 
**Token** | [**Token**](Token.md) |  | 
**LocationId** | **string** |  | 
**EvseUid** | Pointer to **string** |  | [optional] 
**ConnectorId** | Pointer to **string** |  | [optional] 
**AuthorizationReference** | Pointer to **string** |  | [optional] 

## Methods

### NewStartSession

`func NewStartSession(responseUrl string, token Token, locationId string, ) *StartSession`

NewStartSession instantiates a new StartSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartSessionWithDefaults

`func NewStartSessionWithDefaults() *StartSession`

NewStartSessionWithDefaults instantiates a new StartSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseUrl

`func (o *StartSession) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *StartSession) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *StartSession) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.


### GetToken

`func (o *StartSession) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StartSession) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StartSession) SetToken(v Token)`

SetToken sets Token field to given value.


### GetLocationId

`func (o *StartSession) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *StartSession) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *StartSession) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetEvseUid

`func (o *StartSession) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *StartSession) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *StartSession) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.

### HasEvseUid

`func (o *StartSession) HasEvseUid() bool`

HasEvseUid returns a boolean if a field has been set.

### GetConnectorId

`func (o *StartSession) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *StartSession) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *StartSession) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *StartSession) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetAuthorizationReference

`func (o *StartSession) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *StartSession) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *StartSession) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *StartSession) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


