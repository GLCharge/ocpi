# DetailsDataEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**ModuleIDType**](ModuleIDType.md) |  | 
**Role** | [**InterfaceRoleType**](InterfaceRoleType.md) |  | 
**Url** | **string** | URL to the endpoint. | 

## Methods

### NewDetailsDataEndpoints

`func NewDetailsDataEndpoints(identifier ModuleIDType, role InterfaceRoleType, url string, ) *DetailsDataEndpoints`

NewDetailsDataEndpoints instantiates a new DetailsDataEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsDataEndpointsWithDefaults

`func NewDetailsDataEndpointsWithDefaults() *DetailsDataEndpoints`

NewDetailsDataEndpointsWithDefaults instantiates a new DetailsDataEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DetailsDataEndpoints) GetIdentifier() ModuleIDType`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DetailsDataEndpoints) GetIdentifierOk() (*ModuleIDType, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DetailsDataEndpoints) SetIdentifier(v ModuleIDType)`

SetIdentifier sets Identifier field to given value.


### GetRole

`func (o *DetailsDataEndpoints) GetRole() InterfaceRoleType`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DetailsDataEndpoints) GetRoleOk() (*InterfaceRoleType, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DetailsDataEndpoints) SetRole(v InterfaceRoleType)`

SetRole sets Role field to given value.


### GetUrl

`func (o *DetailsDataEndpoints) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DetailsDataEndpoints) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DetailsDataEndpoints) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


