# LocationReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** |  | 
**EvseUids** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationReferences

`func NewLocationReferences(locationId string, ) *LocationReferences`

NewLocationReferences instantiates a new LocationReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReferencesWithDefaults

`func NewLocationReferencesWithDefaults() *LocationReferences`

NewLocationReferencesWithDefaults instantiates a new LocationReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *LocationReferences) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *LocationReferences) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *LocationReferences) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetEvseUids

`func (o *LocationReferences) GetEvseUids() string`

GetEvseUids returns the EvseUids field if non-nil, zero value otherwise.

### GetEvseUidsOk

`func (o *LocationReferences) GetEvseUidsOk() (*string, bool)`

GetEvseUidsOk returns a tuple with the EvseUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUids

`func (o *LocationReferences) SetEvseUids(v string)`

SetEvseUids sets EvseUids field to given value.

### HasEvseUids

`func (o *LocationReferences) HasEvseUids() bool`

HasEvseUids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


