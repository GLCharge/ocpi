# DetailsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Endpoints** | Pointer to [**DetailsDataEndpoints**](DetailsDataEndpoints.md) |  | [optional] 

## Methods

### NewDetailsData

`func NewDetailsData(version string, ) *DetailsData`

NewDetailsData instantiates a new DetailsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsDataWithDefaults

`func NewDetailsDataWithDefaults() *DetailsData`

NewDetailsDataWithDefaults instantiates a new DetailsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DetailsData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DetailsData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DetailsData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEndpoints

`func (o *DetailsData) GetEndpoints() DetailsDataEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *DetailsData) GetEndpointsOk() (*DetailsDataEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *DetailsData) SetEndpoints(v DetailsDataEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *DetailsData) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


