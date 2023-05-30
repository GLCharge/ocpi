# Evse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**EvseId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**StatusSchedule** | Pointer to [**EvseStatusSchedule**](EvseStatusSchedule.md) |  | [optional] 
**Capabilities** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**Connector**](Connector.md) |  | [optional] 
**FloorLevel** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**CdrBodyCdrLocationCoordinates**](CdrBodyCdrLocationCoordinates.md) |  | [optional] 
**PhysicalReference** | Pointer to **string** |  | [optional] 
**Directions** | Pointer to [**CdrBodyTariffsTariffAltText**](CdrBodyTariffsTariffAltText.md) |  | [optional] 
**ParkingRestrictions** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**Image**](Image.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewEvse

`func NewEvse(uid string, status string, lastUpdated string, ) *Evse`

NewEvse instantiates a new Evse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvseWithDefaults

`func NewEvseWithDefaults() *Evse`

NewEvseWithDefaults instantiates a new Evse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Evse) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Evse) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Evse) SetUid(v string)`

SetUid sets Uid field to given value.


### GetEvseId

`func (o *Evse) GetEvseId() string`

GetEvseId returns the EvseId field if non-nil, zero value otherwise.

### GetEvseIdOk

`func (o *Evse) GetEvseIdOk() (*string, bool)`

GetEvseIdOk returns a tuple with the EvseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseId

`func (o *Evse) SetEvseId(v string)`

SetEvseId sets EvseId field to given value.

### HasEvseId

`func (o *Evse) HasEvseId() bool`

HasEvseId returns a boolean if a field has been set.

### GetStatus

`func (o *Evse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Evse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Evse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusSchedule

`func (o *Evse) GetStatusSchedule() EvseStatusSchedule`

GetStatusSchedule returns the StatusSchedule field if non-nil, zero value otherwise.

### GetStatusScheduleOk

`func (o *Evse) GetStatusScheduleOk() (*EvseStatusSchedule, bool)`

GetStatusScheduleOk returns a tuple with the StatusSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSchedule

`func (o *Evse) SetStatusSchedule(v EvseStatusSchedule)`

SetStatusSchedule sets StatusSchedule field to given value.

### HasStatusSchedule

`func (o *Evse) HasStatusSchedule() bool`

HasStatusSchedule returns a boolean if a field has been set.

### GetCapabilities

`func (o *Evse) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Evse) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Evse) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Evse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetConnectors

`func (o *Evse) GetConnectors() Connector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *Evse) GetConnectorsOk() (*Connector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *Evse) SetConnectors(v Connector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *Evse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetFloorLevel

`func (o *Evse) GetFloorLevel() string`

GetFloorLevel returns the FloorLevel field if non-nil, zero value otherwise.

### GetFloorLevelOk

`func (o *Evse) GetFloorLevelOk() (*string, bool)`

GetFloorLevelOk returns a tuple with the FloorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorLevel

`func (o *Evse) SetFloorLevel(v string)`

SetFloorLevel sets FloorLevel field to given value.

### HasFloorLevel

`func (o *Evse) HasFloorLevel() bool`

HasFloorLevel returns a boolean if a field has been set.

### GetCoordinates

`func (o *Evse) GetCoordinates() CdrBodyCdrLocationCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Evse) GetCoordinatesOk() (*CdrBodyCdrLocationCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Evse) SetCoordinates(v CdrBodyCdrLocationCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Evse) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetPhysicalReference

`func (o *Evse) GetPhysicalReference() string`

GetPhysicalReference returns the PhysicalReference field if non-nil, zero value otherwise.

### GetPhysicalReferenceOk

`func (o *Evse) GetPhysicalReferenceOk() (*string, bool)`

GetPhysicalReferenceOk returns a tuple with the PhysicalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalReference

`func (o *Evse) SetPhysicalReference(v string)`

SetPhysicalReference sets PhysicalReference field to given value.

### HasPhysicalReference

`func (o *Evse) HasPhysicalReference() bool`

HasPhysicalReference returns a boolean if a field has been set.

### GetDirections

`func (o *Evse) GetDirections() CdrBodyTariffsTariffAltText`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *Evse) GetDirectionsOk() (*CdrBodyTariffsTariffAltText, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *Evse) SetDirections(v CdrBodyTariffsTariffAltText)`

SetDirections sets Directions field to given value.

### HasDirections

`func (o *Evse) HasDirections() bool`

HasDirections returns a boolean if a field has been set.

### GetParkingRestrictions

`func (o *Evse) GetParkingRestrictions() string`

GetParkingRestrictions returns the ParkingRestrictions field if non-nil, zero value otherwise.

### GetParkingRestrictionsOk

`func (o *Evse) GetParkingRestrictionsOk() (*string, bool)`

GetParkingRestrictionsOk returns a tuple with the ParkingRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingRestrictions

`func (o *Evse) SetParkingRestrictions(v string)`

SetParkingRestrictions sets ParkingRestrictions field to given value.

### HasParkingRestrictions

`func (o *Evse) HasParkingRestrictions() bool`

HasParkingRestrictions returns a boolean if a field has been set.

### GetImages

`func (o *Evse) GetImages() Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Evse) GetImagesOk() (*Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Evse) SetImages(v Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Evse) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Evse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Evse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Evse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


