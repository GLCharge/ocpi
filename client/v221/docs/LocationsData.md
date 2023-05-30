# LocationsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] 
**PartyId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Publish** | Pointer to **bool** |  | [optional] 
**PublishAllowedTo** | Pointer to [**LocationsDataPublishAllowedTo**](LocationsDataPublishAllowedTo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Address** | **string** |  | 
**City** | **string** |  | 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | **string** |  | 
**Coordinates** | [**CdrBodyCdrLocationCoordinates**](CdrBodyCdrLocationCoordinates.md) |  | 
**RelatedLocations** | Pointer to [**LocationsDataRelatedLocations**](LocationsDataRelatedLocations.md) |  | [optional] 
**ParkingType** | Pointer to **string** |  | [optional] 
**Evses** | Pointer to [**Evse**](Evse.md) |  | [optional] 
**Directions** | Pointer to [**CdrBodyTariffsTariffAltText**](CdrBodyTariffsTariffAltText.md) |  | [optional] 
**Operator** | Pointer to [**BusinessDetails**](BusinessDetails.md) |  | [optional] 
**Suboperator** | Pointer to [**BusinessDetails**](BusinessDetails.md) |  | [optional] 
**Owner** | Pointer to [**BusinessDetails**](BusinessDetails.md) |  | [optional] 
**Facilities** | Pointer to **string** |  | [optional] 
**TimeZone** | **string** |  | 
**OpeningTimes** | Pointer to [**LocationsDataOpeningTimes**](LocationsDataOpeningTimes.md) |  | [optional] 
**ChargingWhenClosed** | Pointer to **bool** |  | [optional] 
**Images** | Pointer to [**Image**](Image.md) |  | [optional] 
**EnergyMix** | Pointer to [**CdrBodyTariffsEnergyMix**](CdrBodyTariffsEnergyMix.md) |  | [optional] 
**LastUpdated** | **string** |  | 

## Methods

### NewLocationsData

`func NewLocationsData(address string, city string, country string, coordinates CdrBodyCdrLocationCoordinates, timeZone string, lastUpdated string, ) *LocationsData`

NewLocationsData instantiates a new LocationsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsDataWithDefaults

`func NewLocationsDataWithDefaults() *LocationsData`

NewLocationsDataWithDefaults instantiates a new LocationsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *LocationsData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LocationsData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LocationsData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *LocationsData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPartyId

`func (o *LocationsData) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *LocationsData) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *LocationsData) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *LocationsData) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetId

`func (o *LocationsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationsData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationsData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublish

`func (o *LocationsData) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *LocationsData) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *LocationsData) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *LocationsData) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetPublishAllowedTo

`func (o *LocationsData) GetPublishAllowedTo() LocationsDataPublishAllowedTo`

GetPublishAllowedTo returns the PublishAllowedTo field if non-nil, zero value otherwise.

### GetPublishAllowedToOk

`func (o *LocationsData) GetPublishAllowedToOk() (*LocationsDataPublishAllowedTo, bool)`

GetPublishAllowedToOk returns a tuple with the PublishAllowedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAllowedTo

`func (o *LocationsData) SetPublishAllowedTo(v LocationsDataPublishAllowedTo)`

SetPublishAllowedTo sets PublishAllowedTo field to given value.

### HasPublishAllowedTo

`func (o *LocationsData) HasPublishAllowedTo() bool`

HasPublishAllowedTo returns a boolean if a field has been set.

### GetName

`func (o *LocationsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *LocationsData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationsData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationsData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *LocationsData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationsData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationsData) SetCity(v string)`

SetCity sets City field to given value.


### GetPostalCode

`func (o *LocationsData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LocationsData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LocationsData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LocationsData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *LocationsData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationsData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationsData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationsData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *LocationsData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationsData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationsData) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCoordinates

`func (o *LocationsData) GetCoordinates() CdrBodyCdrLocationCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *LocationsData) GetCoordinatesOk() (*CdrBodyCdrLocationCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *LocationsData) SetCoordinates(v CdrBodyCdrLocationCoordinates)`

SetCoordinates sets Coordinates field to given value.


### GetRelatedLocations

`func (o *LocationsData) GetRelatedLocations() LocationsDataRelatedLocations`

GetRelatedLocations returns the RelatedLocations field if non-nil, zero value otherwise.

### GetRelatedLocationsOk

`func (o *LocationsData) GetRelatedLocationsOk() (*LocationsDataRelatedLocations, bool)`

GetRelatedLocationsOk returns a tuple with the RelatedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedLocations

`func (o *LocationsData) SetRelatedLocations(v LocationsDataRelatedLocations)`

SetRelatedLocations sets RelatedLocations field to given value.

### HasRelatedLocations

`func (o *LocationsData) HasRelatedLocations() bool`

HasRelatedLocations returns a boolean if a field has been set.

### GetParkingType

`func (o *LocationsData) GetParkingType() string`

GetParkingType returns the ParkingType field if non-nil, zero value otherwise.

### GetParkingTypeOk

`func (o *LocationsData) GetParkingTypeOk() (*string, bool)`

GetParkingTypeOk returns a tuple with the ParkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingType

`func (o *LocationsData) SetParkingType(v string)`

SetParkingType sets ParkingType field to given value.

### HasParkingType

`func (o *LocationsData) HasParkingType() bool`

HasParkingType returns a boolean if a field has been set.

### GetEvses

`func (o *LocationsData) GetEvses() Evse`

GetEvses returns the Evses field if non-nil, zero value otherwise.

### GetEvsesOk

`func (o *LocationsData) GetEvsesOk() (*Evse, bool)`

GetEvsesOk returns a tuple with the Evses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvses

`func (o *LocationsData) SetEvses(v Evse)`

SetEvses sets Evses field to given value.

### HasEvses

`func (o *LocationsData) HasEvses() bool`

HasEvses returns a boolean if a field has been set.

### GetDirections

`func (o *LocationsData) GetDirections() CdrBodyTariffsTariffAltText`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *LocationsData) GetDirectionsOk() (*CdrBodyTariffsTariffAltText, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *LocationsData) SetDirections(v CdrBodyTariffsTariffAltText)`

SetDirections sets Directions field to given value.

### HasDirections

`func (o *LocationsData) HasDirections() bool`

HasDirections returns a boolean if a field has been set.

### GetOperator

`func (o *LocationsData) GetOperator() BusinessDetails`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LocationsData) GetOperatorOk() (*BusinessDetails, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LocationsData) SetOperator(v BusinessDetails)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LocationsData) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetSuboperator

`func (o *LocationsData) GetSuboperator() BusinessDetails`

GetSuboperator returns the Suboperator field if non-nil, zero value otherwise.

### GetSuboperatorOk

`func (o *LocationsData) GetSuboperatorOk() (*BusinessDetails, bool)`

GetSuboperatorOk returns a tuple with the Suboperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuboperator

`func (o *LocationsData) SetSuboperator(v BusinessDetails)`

SetSuboperator sets Suboperator field to given value.

### HasSuboperator

`func (o *LocationsData) HasSuboperator() bool`

HasSuboperator returns a boolean if a field has been set.

### GetOwner

`func (o *LocationsData) GetOwner() BusinessDetails`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *LocationsData) GetOwnerOk() (*BusinessDetails, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *LocationsData) SetOwner(v BusinessDetails)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *LocationsData) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFacilities

`func (o *LocationsData) GetFacilities() string`

GetFacilities returns the Facilities field if non-nil, zero value otherwise.

### GetFacilitiesOk

`func (o *LocationsData) GetFacilitiesOk() (*string, bool)`

GetFacilitiesOk returns a tuple with the Facilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilities

`func (o *LocationsData) SetFacilities(v string)`

SetFacilities sets Facilities field to given value.

### HasFacilities

`func (o *LocationsData) HasFacilities() bool`

HasFacilities returns a boolean if a field has been set.

### GetTimeZone

`func (o *LocationsData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *LocationsData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *LocationsData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetOpeningTimes

`func (o *LocationsData) GetOpeningTimes() LocationsDataOpeningTimes`

GetOpeningTimes returns the OpeningTimes field if non-nil, zero value otherwise.

### GetOpeningTimesOk

`func (o *LocationsData) GetOpeningTimesOk() (*LocationsDataOpeningTimes, bool)`

GetOpeningTimesOk returns a tuple with the OpeningTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningTimes

`func (o *LocationsData) SetOpeningTimes(v LocationsDataOpeningTimes)`

SetOpeningTimes sets OpeningTimes field to given value.

### HasOpeningTimes

`func (o *LocationsData) HasOpeningTimes() bool`

HasOpeningTimes returns a boolean if a field has been set.

### GetChargingWhenClosed

`func (o *LocationsData) GetChargingWhenClosed() bool`

GetChargingWhenClosed returns the ChargingWhenClosed field if non-nil, zero value otherwise.

### GetChargingWhenClosedOk

`func (o *LocationsData) GetChargingWhenClosedOk() (*bool, bool)`

GetChargingWhenClosedOk returns a tuple with the ChargingWhenClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingWhenClosed

`func (o *LocationsData) SetChargingWhenClosed(v bool)`

SetChargingWhenClosed sets ChargingWhenClosed field to given value.

### HasChargingWhenClosed

`func (o *LocationsData) HasChargingWhenClosed() bool`

HasChargingWhenClosed returns a boolean if a field has been set.

### GetImages

`func (o *LocationsData) GetImages() Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *LocationsData) GetImagesOk() (*Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *LocationsData) SetImages(v Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *LocationsData) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetEnergyMix

`func (o *LocationsData) GetEnergyMix() CdrBodyTariffsEnergyMix`

GetEnergyMix returns the EnergyMix field if non-nil, zero value otherwise.

### GetEnergyMixOk

`func (o *LocationsData) GetEnergyMixOk() (*CdrBodyTariffsEnergyMix, bool)`

GetEnergyMixOk returns a tuple with the EnergyMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyMix

`func (o *LocationsData) SetEnergyMix(v CdrBodyTariffsEnergyMix)`

SetEnergyMix sets EnergyMix field to given value.

### HasEnergyMix

`func (o *LocationsData) HasEnergyMix() bool`

HasEnergyMix returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LocationsData) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LocationsData) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LocationsData) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


