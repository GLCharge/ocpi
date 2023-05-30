# CdrBodyCdrLocation

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
**Coordinates** | Pointer to [**CdrBodyCdrLocationCoordinates**](CdrBodyCdrLocationCoordinates.md) |  | [optional] 
**EvseUid** | Pointer to **string** |  | [optional] 
**EvseId** | Pointer to **string** |  | [optional] 
**ConnectorId** | Pointer to **string** |  | [optional] 
**ConnectorStandard** | Pointer to **string** |  | [optional] 
**ConnectorFormat** | Pointer to **string** |  | [optional] 
**ConnectorPowerType** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrBodyCdrLocation

`func NewCdrBodyCdrLocation() *CdrBodyCdrLocation`

NewCdrBodyCdrLocation instantiates a new CdrBodyCdrLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrBodyCdrLocationWithDefaults

`func NewCdrBodyCdrLocationWithDefaults() *CdrBodyCdrLocation`

NewCdrBodyCdrLocationWithDefaults instantiates a new CdrBodyCdrLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdrBodyCdrLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrBodyCdrLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrBodyCdrLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdrBodyCdrLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CdrBodyCdrLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdrBodyCdrLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdrBodyCdrLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdrBodyCdrLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *CdrBodyCdrLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CdrBodyCdrLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CdrBodyCdrLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CdrBodyCdrLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *CdrBodyCdrLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CdrBodyCdrLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CdrBodyCdrLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CdrBodyCdrLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *CdrBodyCdrLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CdrBodyCdrLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CdrBodyCdrLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CdrBodyCdrLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CdrBodyCdrLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CdrBodyCdrLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CdrBodyCdrLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CdrBodyCdrLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CdrBodyCdrLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CdrBodyCdrLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CdrBodyCdrLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CdrBodyCdrLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoordinates

`func (o *CdrBodyCdrLocation) GetCoordinates() CdrBodyCdrLocationCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CdrBodyCdrLocation) GetCoordinatesOk() (*CdrBodyCdrLocationCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CdrBodyCdrLocation) SetCoordinates(v CdrBodyCdrLocationCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *CdrBodyCdrLocation) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetEvseUid

`func (o *CdrBodyCdrLocation) GetEvseUid() string`

GetEvseUid returns the EvseUid field if non-nil, zero value otherwise.

### GetEvseUidOk

`func (o *CdrBodyCdrLocation) GetEvseUidOk() (*string, bool)`

GetEvseUidOk returns a tuple with the EvseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseUid

`func (o *CdrBodyCdrLocation) SetEvseUid(v string)`

SetEvseUid sets EvseUid field to given value.

### HasEvseUid

`func (o *CdrBodyCdrLocation) HasEvseUid() bool`

HasEvseUid returns a boolean if a field has been set.

### GetEvseId

`func (o *CdrBodyCdrLocation) GetEvseId() string`

GetEvseId returns the EvseId field if non-nil, zero value otherwise.

### GetEvseIdOk

`func (o *CdrBodyCdrLocation) GetEvseIdOk() (*string, bool)`

GetEvseIdOk returns a tuple with the EvseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseId

`func (o *CdrBodyCdrLocation) SetEvseId(v string)`

SetEvseId sets EvseId field to given value.

### HasEvseId

`func (o *CdrBodyCdrLocation) HasEvseId() bool`

HasEvseId returns a boolean if a field has been set.

### GetConnectorId

`func (o *CdrBodyCdrLocation) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CdrBodyCdrLocation) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CdrBodyCdrLocation) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *CdrBodyCdrLocation) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetConnectorStandard

`func (o *CdrBodyCdrLocation) GetConnectorStandard() string`

GetConnectorStandard returns the ConnectorStandard field if non-nil, zero value otherwise.

### GetConnectorStandardOk

`func (o *CdrBodyCdrLocation) GetConnectorStandardOk() (*string, bool)`

GetConnectorStandardOk returns a tuple with the ConnectorStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorStandard

`func (o *CdrBodyCdrLocation) SetConnectorStandard(v string)`

SetConnectorStandard sets ConnectorStandard field to given value.

### HasConnectorStandard

`func (o *CdrBodyCdrLocation) HasConnectorStandard() bool`

HasConnectorStandard returns a boolean if a field has been set.

### GetConnectorFormat

`func (o *CdrBodyCdrLocation) GetConnectorFormat() string`

GetConnectorFormat returns the ConnectorFormat field if non-nil, zero value otherwise.

### GetConnectorFormatOk

`func (o *CdrBodyCdrLocation) GetConnectorFormatOk() (*string, bool)`

GetConnectorFormatOk returns a tuple with the ConnectorFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorFormat

`func (o *CdrBodyCdrLocation) SetConnectorFormat(v string)`

SetConnectorFormat sets ConnectorFormat field to given value.

### HasConnectorFormat

`func (o *CdrBodyCdrLocation) HasConnectorFormat() bool`

HasConnectorFormat returns a boolean if a field has been set.

### GetConnectorPowerType

`func (o *CdrBodyCdrLocation) GetConnectorPowerType() string`

GetConnectorPowerType returns the ConnectorPowerType field if non-nil, zero value otherwise.

### GetConnectorPowerTypeOk

`func (o *CdrBodyCdrLocation) GetConnectorPowerTypeOk() (*string, bool)`

GetConnectorPowerTypeOk returns a tuple with the ConnectorPowerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPowerType

`func (o *CdrBodyCdrLocation) SetConnectorPowerType(v string)`

SetConnectorPowerType sets ConnectorPowerType field to given value.

### HasConnectorPowerType

`func (o *CdrBodyCdrLocation) HasConnectorPowerType() bool`

HasConnectorPowerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


