# CdrCdrLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**CdrCdrLocationCoordinates**](CdrCdrLocationCoordinates.md) |  | [optional] 
**EvseUid** | Pointer to **string** |  | [optional] 
**EvseId** | Pointer to **string** |  | [optional] 
**ConnectorId** | Pointer to **string** |  | [optional] 
**ConnectorStandard** | Pointer to **string** |  | [optional] 
**ConnectorFormat** | Pointer to **string** |  | [optional] 
**ConnectorPowerType** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrCdrLocation

`func NewCdrCdrLocation() *CdrCdrLocation`

NewCdrCdrLocation instantiates a new CdrCdrLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrCdrLocationWithDefaults

`func NewCdrCdrLocationWithDefaults() *CdrCdrLocation`

NewCdrCdrLocationWithDefaults instantiates a new CdrCdrLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdrCdrLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrCdrLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrCdrLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdrCdrLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CdrCdrLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdrCdrLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdrCdrLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdrCdrLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *CdrCdrLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CdrCdrLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CdrCdrLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CdrCdrLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *CdrCdrLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CdrCdrLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CdrCdrLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CdrCdrLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *CdrCdrLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CdrCdrLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CdrCdrLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CdrCdrLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CdrCdrLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CdrCdrLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CdrCdrLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CdrCdrLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CdrCdrLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CdrCdrLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CdrCdrLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CdrCdrLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoordinates

`func (o *CdrCdrLocation) GetCoordinates() CdrCdrLocationCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CdrCdrLocation) GetCoordinatesOk() (*CdrCdrLocationCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CdrCdrLocation) SetCoordinates(v CdrCdrLocationCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *CdrCdrLocation) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetEvseUid

`func (o *CdrCdrLocation) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *CdrCdrLocation) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *CdrCdrLocation) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.

### HasEvseUid

`func (o *CdrCdrLocation) HasEvseUid() bool`

HasEvseUid returns a boolean if a field has been set.

### GetEvseId

`func (o *CdrCdrLocation) GetEvseId() string`

GetEvseId returns the EvseId field if non-nil, zero value otherwise.

### GetEvseIdOk

`func (o *CdrCdrLocation) GetEvseIdOk() (*string, bool)`

GetEvseIdOk returns a tuple with the EvseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseId

`func (o *CdrCdrLocation) SetEvseId(v string)`

SetEvseId sets EvseId field to given value.

### HasEvseId

`func (o *CdrCdrLocation) HasEvseId() bool`

HasEvseId returns a boolean if a field has been set.

### GetConnectorId

`func (o *CdrCdrLocation) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CdrCdrLocation) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CdrCdrLocation) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *CdrCdrLocation) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetConnectorStandard

`func (o *CdrCdrLocation) GetConnectorStandard() string`

GetConnectorStandard returns the ConnectorStandard field if non-nil, zero value otherwise.

### GetConnectorStandardOk

`func (o *CdrCdrLocation) GetConnectorStandardOk() (*string, bool)`

GetConnectorStandardOk returns a tuple with the ConnectorStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorStandard

`func (o *CdrCdrLocation) SetConnectorStandard(v string)`

SetConnectorStandard sets ConnectorStandard field to given value.

### HasConnectorStandard

`func (o *CdrCdrLocation) HasConnectorStandard() bool`

HasConnectorStandard returns a boolean if a field has been set.

### GetConnectorFormat

`func (o *CdrCdrLocation) GetConnectorFormat() string`

GetConnectorFormat returns the ConnectorFormat field if non-nil, zero value otherwise.

### GetConnectorFormatOk

`func (o *CdrCdrLocation) GetConnectorFormatOk() (*string, bool)`

GetConnectorFormatOk returns a tuple with the ConnectorFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorFormat

`func (o *CdrCdrLocation) SetConnectorFormat(v string)`

SetConnectorFormat sets ConnectorFormat field to given value.

### HasConnectorFormat

`func (o *CdrCdrLocation) HasConnectorFormat() bool`

HasConnectorFormat returns a boolean if a field has been set.

### GetConnectorPowerType

`func (o *CdrCdrLocation) GetConnectorPowerType() string`

GetConnectorPowerType returns the ConnectorPowerType field if non-nil, zero value otherwise.

### GetConnectorPowerTypeOk

`func (o *CdrCdrLocation) GetConnectorPowerTypeOk() (*string, bool)`

GetConnectorPowerTypeOk returns a tuple with the ConnectorPowerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPowerType

`func (o *CdrCdrLocation) SetConnectorPowerType(v string)`

SetConnectorPowerType sets ConnectorPowerType field to given value.

### HasConnectorPowerType

`func (o *CdrCdrLocation) HasConnectorPowerType() bool`

HasConnectorPowerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


