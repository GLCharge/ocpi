# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | [**Token**](Token.md) |  | 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse(token Token, statusCode float32, ) *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenResponse) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenResponse) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenResponse) SetToken(v Token)`

SetToken sets Token field to given value.


### GetStatusCode

`func (o *TokenResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TokenResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TokenResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *TokenResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *TokenResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *TokenResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *TokenResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *TokenResponse) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *TokenResponse) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *TokenResponse) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *TokenResponse) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


