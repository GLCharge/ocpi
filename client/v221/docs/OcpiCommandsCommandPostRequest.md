# OcpiCommandsCommandPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseUrl** | **string** |  | 
**ReservationId** | **string** |  | 
**Token** | [**Token**](Token.md) |  | 
**ExpiryDate** | **string** |  | 
**LocationId** | **string** |  | 
**EvseUid** | **string** |  | 
**AuthorizationReference** | Pointer to **string** |  | [optional] 
**ConnectorId** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewOcpiCommandsCommandPostRequest

`func NewOcpiCommandsCommandPostRequest(responseUrl string, reservationId string, token Token, expiryDate string, locationId string, evseUid string, connectorId string, ) *OcpiCommandsCommandPostRequest`

NewOcpiCommandsCommandPostRequest instantiates a new OcpiCommandsCommandPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcpiCommandsCommandPostRequestWithDefaults

`func NewOcpiCommandsCommandPostRequestWithDefaults() *OcpiCommandsCommandPostRequest`

NewOcpiCommandsCommandPostRequestWithDefaults instantiates a new OcpiCommandsCommandPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseUrl

`func (o *OcpiCommandsCommandPostRequest) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *OcpiCommandsCommandPostRequest) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *OcpiCommandsCommandPostRequest) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.


### GetReservationId

`func (o *OcpiCommandsCommandPostRequest) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *OcpiCommandsCommandPostRequest) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *OcpiCommandsCommandPostRequest) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.


### GetToken

`func (o *OcpiCommandsCommandPostRequest) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OcpiCommandsCommandPostRequest) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OcpiCommandsCommandPostRequest) SetToken(v Token)`

SetToken sets Token field to given value.


### GetExpiryDate

`func (o *OcpiCommandsCommandPostRequest) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OcpiCommandsCommandPostRequest) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OcpiCommandsCommandPostRequest) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetLocationId

`func (o *OcpiCommandsCommandPostRequest) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OcpiCommandsCommandPostRequest) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OcpiCommandsCommandPostRequest) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetEvseUid

`func (o *OcpiCommandsCommandPostRequest) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *OcpiCommandsCommandPostRequest) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *OcpiCommandsCommandPostRequest) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.


### GetAuthorizationReference

`func (o *OcpiCommandsCommandPostRequest) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *OcpiCommandsCommandPostRequest) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *OcpiCommandsCommandPostRequest) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *OcpiCommandsCommandPostRequest) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetConnectorId

`func (o *OcpiCommandsCommandPostRequest) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *OcpiCommandsCommandPostRequest) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *OcpiCommandsCommandPostRequest) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetSessionId

`func (o *OcpiCommandsCommandPostRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *OcpiCommandsCommandPostRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *OcpiCommandsCommandPostRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *OcpiCommandsCommandPostRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


