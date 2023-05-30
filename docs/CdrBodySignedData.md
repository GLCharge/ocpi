# CdrBodySignedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingMethod** | **string** |  | 
**EncodingMethodVersion** | Pointer to **int32** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**SignedValues** | Pointer to [**CdrBodySignedDataSignedValues**](CdrBodySignedDataSignedValues.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrBodySignedData

`func NewCdrBodySignedData(encodingMethod string, ) *CdrBodySignedData`

NewCdrBodySignedData instantiates a new CdrBodySignedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrBodySignedDataWithDefaults

`func NewCdrBodySignedDataWithDefaults() *CdrBodySignedData`

NewCdrBodySignedDataWithDefaults instantiates a new CdrBodySignedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingMethod

`func (o *CdrBodySignedData) GetEncodingMethod() string`

GetEncodingMethod returns the EncodingMethod field if non-nil, zero value otherwise.

### GetEncodingMethodOk

`func (o *CdrBodySignedData) GetEncodingMethodOk() (*string, bool)`

GetEncodingMethodOk returns a tuple with the EncodingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingMethod

`func (o *CdrBodySignedData) SetEncodingMethod(v string)`

SetEncodingMethod sets EncodingMethod field to given value.


### GetEncodingMethodVersion

`func (o *CdrBodySignedData) GetEncodingMethodVersion() int32`

GetEncodingMethodVersion returns the EncodingMethodVersion field if non-nil, zero value otherwise.

### GetEncodingMethodVersionOk

`func (o *CdrBodySignedData) GetEncodingMethodVersionOk() (*int32, bool)`

GetEncodingMethodVersionOk returns a tuple with the EncodingMethodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingMethodVersion

`func (o *CdrBodySignedData) SetEncodingMethodVersion(v int32)`

SetEncodingMethodVersion sets EncodingMethodVersion field to given value.

### HasEncodingMethodVersion

`func (o *CdrBodySignedData) HasEncodingMethodVersion() bool`

HasEncodingMethodVersion returns a boolean if a field has been set.

### GetPublicKey

`func (o *CdrBodySignedData) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CdrBodySignedData) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CdrBodySignedData) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CdrBodySignedData) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSignedValues

`func (o *CdrBodySignedData) GetSignedValues() CdrBodySignedDataSignedValues`

GetSignedValues returns the SignedValues field if non-nil, zero value otherwise.

### GetSignedValuesOk

`func (o *CdrBodySignedData) GetSignedValuesOk() (*CdrBodySignedDataSignedValues, bool)`

GetSignedValuesOk returns a tuple with the SignedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedValues

`func (o *CdrBodySignedData) SetSignedValues(v CdrBodySignedDataSignedValues)`

SetSignedValues sets SignedValues field to given value.

### HasSignedValues

`func (o *CdrBodySignedData) HasSignedValues() bool`

HasSignedValues returns a boolean if a field has been set.

### GetUrl

`func (o *CdrBodySignedData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CdrBodySignedData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CdrBodySignedData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CdrBodySignedData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


