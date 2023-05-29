# CdrSignedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingMethod** | **string** |  | 
**EncodingMethodVersion** | Pointer to **int32** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**SignedValues** | Pointer to [**CdrSignedDataSignedValues**](CdrSignedDataSignedValues.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrSignedData

`func NewCdrSignedData(encodingMethod string, ) *CdrSignedData`

NewCdrSignedData instantiates a new CdrSignedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrSignedDataWithDefaults

`func NewCdrSignedDataWithDefaults() *CdrSignedData`

NewCdrSignedDataWithDefaults instantiates a new CdrSignedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingMethod

`func (o *CdrSignedData) GetEncodingMethod() string`

GetEncodingMethod returns the EncodingMethod field if non-nil, zero value otherwise.

### GetEncodingMethodOk

`func (o *CdrSignedData) GetEncodingMethodOk() (*string, bool)`

GetEncodingMethodOk returns a tuple with the EncodingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingMethod

`func (o *CdrSignedData) SetEncodingMethod(v string)`

SetEncodingMethod sets EncodingMethod field to given value.


### GetEncodingMethodVersion

`func (o *CdrSignedData) GetEncodingMethodVersion() int32`

GetEncodingMethodVersion returns the EncodingMethodVersion field if non-nil, zero value otherwise.

### GetEncodingMethodVersionOk

`func (o *CdrSignedData) GetEncodingMethodVersionOk() (*int32, bool)`

GetEncodingMethodVersionOk returns a tuple with the EncodingMethodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingMethodVersion

`func (o *CdrSignedData) SetEncodingMethodVersion(v int32)`

SetEncodingMethodVersion sets EncodingMethodVersion field to given value.

### HasEncodingMethodVersion

`func (o *CdrSignedData) HasEncodingMethodVersion() bool`

HasEncodingMethodVersion returns a boolean if a field has been set.

### GetPublicKey

`func (o *CdrSignedData) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CdrSignedData) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CdrSignedData) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CdrSignedData) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSignedValues

`func (o *CdrSignedData) GetSignedValues() CdrSignedDataSignedValues`

GetSignedValues returns the SignedValues field if non-nil, zero value otherwise.

### GetSignedValuesOk

`func (o *CdrSignedData) GetSignedValuesOk() (*CdrSignedDataSignedValues, bool)`

GetSignedValuesOk returns a tuple with the SignedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedValues

`func (o *CdrSignedData) SetSignedValues(v CdrSignedDataSignedValues)`

SetSignedValues sets SignedValues field to given value.

### HasSignedValues

`func (o *CdrSignedData) HasSignedValues() bool`

HasSignedValues returns a boolean if a field has been set.

### GetUrl

`func (o *CdrSignedData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CdrSignedData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CdrSignedData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CdrSignedData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


