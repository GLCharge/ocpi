# ReserveNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseUrl** | **string** |  | 
**Token** | [**Token**](Token.md) |  | 
**ExpiryDate** | **string** |  | 
**ReservationId** | **string** |  | 
**LocationId** | **string** |  | 
**EvseUid** | Pointer to **string** |  | [optional] 
**AuthorizationReference** | Pointer to **string** |  | [optional] 

## Methods

### NewReserveNow

`func NewReserveNow(responseUrl string, token Token, expiryDate string, reservationId string, locationId string, ) *ReserveNow`

NewReserveNow instantiates a new ReserveNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReserveNowWithDefaults

`func NewReserveNowWithDefaults() *ReserveNow`

NewReserveNowWithDefaults instantiates a new ReserveNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseUrl

`func (o *ReserveNow) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *ReserveNow) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *ReserveNow) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.


### GetToken

`func (o *ReserveNow) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReserveNow) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReserveNow) SetToken(v Token)`

SetToken sets Token field to given value.


### GetExpiryDate

`func (o *ReserveNow) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ReserveNow) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ReserveNow) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetReservationId

`func (o *ReserveNow) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *ReserveNow) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *ReserveNow) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.


### GetLocationId

`func (o *ReserveNow) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ReserveNow) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ReserveNow) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetEvseUid

`func (o *ReserveNow) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *ReserveNow) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *ReserveNow) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.

### HasEvseUid

`func (o *ReserveNow) HasEvseUid() bool`

HasEvseUid returns a boolean if a field has been set.

### GetAuthorizationReference

`func (o *ReserveNow) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *ReserveNow) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *ReserveNow) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *ReserveNow) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


