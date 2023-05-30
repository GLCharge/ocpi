# OcpiLocationsLocationIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LocationsData**](LocationsData.md) |  | [optional] 
**StatusCode** | **float32** |  | 
**StatusMessage** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 
**Uid** | **string** |  | 
**EvseId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**StatusSchedule** | Pointer to [**EvseStatusSchedule**](EvseStatusSchedule.md) |  | [optional] 
**Capabilities** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**Connector**](Connector.md) |  | [optional] 
**FloorLevel** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**CdrCdrLocationCoordinates**](CdrCdrLocationCoordinates.md) |  | [optional] 
**PhysicalReference** | Pointer to **string** |  | [optional] 
**Directions** | Pointer to [**CdrTariffsTariffAltText**](CdrTariffsTariffAltText.md) |  | [optional] 
**ParkingRestrictions** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**Image**](Image.md) |  | [optional] 
**LastUpdated** | **string** |  | 
**Id** | **string** |  | 
**Standard** | **string** |  | 
**Format** | **string** |  | 
**PowerType** | **string** |  | 
**MaxVoltage** | **int32** |  | 
**MaxAmperage** | **int32** |  | 
**MaxElectricPower** | Pointer to **int32** |  | [optional] 
**TariffIds** | Pointer to **string** |  | [optional] 
**TermsAndConditions** | Pointer to **string** |  | [optional] 

## Methods

### NewOcpiLocationsLocationIdGet200Response

`func NewOcpiLocationsLocationIdGet200Response(statusCode float32, uid string, status string, lastUpdated string, id string, standard string, format string, powerType string, maxVoltage int32, maxAmperage int32, ) *OcpiLocationsLocationIdGet200Response`

NewOcpiLocationsLocationIdGet200Response instantiates a new OcpiLocationsLocationIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcpiLocationsLocationIdGet200ResponseWithDefaults

`func NewOcpiLocationsLocationIdGet200ResponseWithDefaults() *OcpiLocationsLocationIdGet200Response`

NewOcpiLocationsLocationIdGet200ResponseWithDefaults instantiates a new OcpiLocationsLocationIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OcpiLocationsLocationIdGet200Response) GetData() LocationsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OcpiLocationsLocationIdGet200Response) GetDataOk() (*LocationsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OcpiLocationsLocationIdGet200Response) SetData(v LocationsData)`

SetData sets Data field to given value.

### HasData

`func (o *OcpiLocationsLocationIdGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatusCode

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OcpiLocationsLocationIdGet200Response) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *OcpiLocationsLocationIdGet200Response) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *OcpiLocationsLocationIdGet200Response) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTimeStamp

`func (o *OcpiLocationsLocationIdGet200Response) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OcpiLocationsLocationIdGet200Response) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OcpiLocationsLocationIdGet200Response) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *OcpiLocationsLocationIdGet200Response) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetUid

`func (o *OcpiLocationsLocationIdGet200Response) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OcpiLocationsLocationIdGet200Response) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OcpiLocationsLocationIdGet200Response) SetUid(v string)`

SetUid sets Uid field to given value.


### GetEvseId

`func (o *OcpiLocationsLocationIdGet200Response) GetEvseId() string`

GetEvseId returns the EvseId field if non-nil, zero value otherwise.

### GetEvseIdOk

`func (o *OcpiLocationsLocationIdGet200Response) GetEvseIdOk() (*string, bool)`

GetEvseIdOk returns a tuple with the EvseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvseId

`func (o *OcpiLocationsLocationIdGet200Response) SetEvseId(v string)`

SetEvseId sets EvseId field to given value.

### HasEvseId

`func (o *OcpiLocationsLocationIdGet200Response) HasEvseId() bool`

HasEvseId returns a boolean if a field has been set.

### GetStatus

`func (o *OcpiLocationsLocationIdGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OcpiLocationsLocationIdGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusSchedule

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusSchedule() EvseStatusSchedule`

GetStatusSchedule returns the StatusSchedule field if non-nil, zero value otherwise.

### GetStatusScheduleOk

`func (o *OcpiLocationsLocationIdGet200Response) GetStatusScheduleOk() (*EvseStatusSchedule, bool)`

GetStatusScheduleOk returns a tuple with the StatusSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSchedule

`func (o *OcpiLocationsLocationIdGet200Response) SetStatusSchedule(v EvseStatusSchedule)`

SetStatusSchedule sets StatusSchedule field to given value.

### HasStatusSchedule

`func (o *OcpiLocationsLocationIdGet200Response) HasStatusSchedule() bool`

HasStatusSchedule returns a boolean if a field has been set.

### GetCapabilities

`func (o *OcpiLocationsLocationIdGet200Response) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OcpiLocationsLocationIdGet200Response) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OcpiLocationsLocationIdGet200Response) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OcpiLocationsLocationIdGet200Response) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetConnectors

`func (o *OcpiLocationsLocationIdGet200Response) GetConnectors() Connector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *OcpiLocationsLocationIdGet200Response) GetConnectorsOk() (*Connector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *OcpiLocationsLocationIdGet200Response) SetConnectors(v Connector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *OcpiLocationsLocationIdGet200Response) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetFloorLevel

`func (o *OcpiLocationsLocationIdGet200Response) GetFloorLevel() string`

GetFloorLevel returns the FloorLevel field if non-nil, zero value otherwise.

### GetFloorLevelOk

`func (o *OcpiLocationsLocationIdGet200Response) GetFloorLevelOk() (*string, bool)`

GetFloorLevelOk returns a tuple with the FloorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorLevel

`func (o *OcpiLocationsLocationIdGet200Response) SetFloorLevel(v string)`

SetFloorLevel sets FloorLevel field to given value.

### HasFloorLevel

`func (o *OcpiLocationsLocationIdGet200Response) HasFloorLevel() bool`

HasFloorLevel returns a boolean if a field has been set.

### GetCoordinates

`func (o *OcpiLocationsLocationIdGet200Response) GetCoordinates() CdrCdrLocationCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *OcpiLocationsLocationIdGet200Response) GetCoordinatesOk() (*CdrCdrLocationCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *OcpiLocationsLocationIdGet200Response) SetCoordinates(v CdrCdrLocationCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *OcpiLocationsLocationIdGet200Response) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetPhysicalReference

`func (o *OcpiLocationsLocationIdGet200Response) GetPhysicalReference() string`

GetPhysicalReference returns the PhysicalReference field if non-nil, zero value otherwise.

### GetPhysicalReferenceOk

`func (o *OcpiLocationsLocationIdGet200Response) GetPhysicalReferenceOk() (*string, bool)`

GetPhysicalReferenceOk returns a tuple with the PhysicalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalReference

`func (o *OcpiLocationsLocationIdGet200Response) SetPhysicalReference(v string)`

SetPhysicalReference sets PhysicalReference field to given value.

### HasPhysicalReference

`func (o *OcpiLocationsLocationIdGet200Response) HasPhysicalReference() bool`

HasPhysicalReference returns a boolean if a field has been set.

### GetDirections

`func (o *OcpiLocationsLocationIdGet200Response) GetDirections() CdrTariffsTariffAltText`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *OcpiLocationsLocationIdGet200Response) GetDirectionsOk() (*CdrTariffsTariffAltText, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *OcpiLocationsLocationIdGet200Response) SetDirections(v CdrTariffsTariffAltText)`

SetDirections sets Directions field to given value.

### HasDirections

`func (o *OcpiLocationsLocationIdGet200Response) HasDirections() bool`

HasDirections returns a boolean if a field has been set.

### GetParkingRestrictions

`func (o *OcpiLocationsLocationIdGet200Response) GetParkingRestrictions() string`

GetParkingRestrictions returns the ParkingRestrictions field if non-nil, zero value otherwise.

### GetParkingRestrictionsOk

`func (o *OcpiLocationsLocationIdGet200Response) GetParkingRestrictionsOk() (*string, bool)`

GetParkingRestrictionsOk returns a tuple with the ParkingRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingRestrictions

`func (o *OcpiLocationsLocationIdGet200Response) SetParkingRestrictions(v string)`

SetParkingRestrictions sets ParkingRestrictions field to given value.

### HasParkingRestrictions

`func (o *OcpiLocationsLocationIdGet200Response) HasParkingRestrictions() bool`

HasParkingRestrictions returns a boolean if a field has been set.

### GetImages

`func (o *OcpiLocationsLocationIdGet200Response) GetImages() Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *OcpiLocationsLocationIdGet200Response) GetImagesOk() (*Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *OcpiLocationsLocationIdGet200Response) SetImages(v Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *OcpiLocationsLocationIdGet200Response) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OcpiLocationsLocationIdGet200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OcpiLocationsLocationIdGet200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OcpiLocationsLocationIdGet200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetId

`func (o *OcpiLocationsLocationIdGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OcpiLocationsLocationIdGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OcpiLocationsLocationIdGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetStandard

`func (o *OcpiLocationsLocationIdGet200Response) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *OcpiLocationsLocationIdGet200Response) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *OcpiLocationsLocationIdGet200Response) SetStandard(v string)`

SetStandard sets Standard field to given value.


### GetFormat

`func (o *OcpiLocationsLocationIdGet200Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OcpiLocationsLocationIdGet200Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OcpiLocationsLocationIdGet200Response) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetPowerType

`func (o *OcpiLocationsLocationIdGet200Response) GetPowerType() string`

GetPowerType returns the PowerType field if non-nil, zero value otherwise.

### GetPowerTypeOk

`func (o *OcpiLocationsLocationIdGet200Response) GetPowerTypeOk() (*string, bool)`

GetPowerTypeOk returns a tuple with the PowerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerType

`func (o *OcpiLocationsLocationIdGet200Response) SetPowerType(v string)`

SetPowerType sets PowerType field to given value.


### GetMaxVoltage

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxVoltage() int32`

GetMaxVoltage returns the MaxVoltage field if non-nil, zero value otherwise.

### GetMaxVoltageOk

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxVoltageOk() (*int32, bool)`

GetMaxVoltageOk returns a tuple with the MaxVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVoltage

`func (o *OcpiLocationsLocationIdGet200Response) SetMaxVoltage(v int32)`

SetMaxVoltage sets MaxVoltage field to given value.


### GetMaxAmperage

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxAmperage() int32`

GetMaxAmperage returns the MaxAmperage field if non-nil, zero value otherwise.

### GetMaxAmperageOk

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxAmperageOk() (*int32, bool)`

GetMaxAmperageOk returns a tuple with the MaxAmperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmperage

`func (o *OcpiLocationsLocationIdGet200Response) SetMaxAmperage(v int32)`

SetMaxAmperage sets MaxAmperage field to given value.


### GetMaxElectricPower

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxElectricPower() int32`

GetMaxElectricPower returns the MaxElectricPower field if non-nil, zero value otherwise.

### GetMaxElectricPowerOk

`func (o *OcpiLocationsLocationIdGet200Response) GetMaxElectricPowerOk() (*int32, bool)`

GetMaxElectricPowerOk returns a tuple with the MaxElectricPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElectricPower

`func (o *OcpiLocationsLocationIdGet200Response) SetMaxElectricPower(v int32)`

SetMaxElectricPower sets MaxElectricPower field to given value.

### HasMaxElectricPower

`func (o *OcpiLocationsLocationIdGet200Response) HasMaxElectricPower() bool`

HasMaxElectricPower returns a boolean if a field has been set.

### GetTariffIds

`func (o *OcpiLocationsLocationIdGet200Response) GetTariffIds() string`

GetTariffIds returns the TariffIds field if non-nil, zero value otherwise.

### GetTariffIdsOk

`func (o *OcpiLocationsLocationIdGet200Response) GetTariffIdsOk() (*string, bool)`

GetTariffIdsOk returns a tuple with the TariffIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffIds

`func (o *OcpiLocationsLocationIdGet200Response) SetTariffIds(v string)`

SetTariffIds sets TariffIds field to given value.

### HasTariffIds

`func (o *OcpiLocationsLocationIdGet200Response) HasTariffIds() bool`

HasTariffIds returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *OcpiLocationsLocationIdGet200Response) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *OcpiLocationsLocationIdGet200Response) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *OcpiLocationsLocationIdGet200Response) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *OcpiLocationsLocationIdGet200Response) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


