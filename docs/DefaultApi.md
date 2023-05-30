# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ocpi221CredentialsDelete**](DefaultApi.md#Ocpi221CredentialsDelete) | **Delete** /ocpi/2.2.1/credentials | Informs that credentials are invalid
[**Ocpi221CredentialsGet**](DefaultApi.md#Ocpi221CredentialsGet) | **Get** /ocpi/2.2.1/credentials | Get the credentials object
[**Ocpi221CredentialsPost**](DefaultApi.md#Ocpi221CredentialsPost) | **Post** /ocpi/2.2.1/credentials | Provides the credentials object
[**Ocpi221CredentialsPut**](DefaultApi.md#Ocpi221CredentialsPut) | **Put** /ocpi/2.2.1/credentials | Provides updated credentials object
[**Ocpi221DetailsGet**](DefaultApi.md#Ocpi221DetailsGet) | **Get** /ocpi/2.2.1/details | Get version supported endpoints
[**OcpiCdrsGet**](DefaultApi.md#OcpiCdrsGet) | **Get** /ocpi/cdrs | Get CDRs
[**OcpiCdrsPost**](DefaultApi.md#OcpiCdrsPost) | **Post** /ocpi/cdrs | Create CDR object
[**OcpiChargingprofilePost**](DefaultApi.md#OcpiChargingprofilePost) | **Post** /ocpi/chargingprofile | Send a result of the original request by the eMSP
[**OcpiChargingprofileSessionIdPut**](DefaultApi.md#OcpiChargingprofileSessionIdPut) | **Put** /ocpi/chargingprofile/{session_id} | Update charging profile
[**OcpiChargingprofilesSessionIdDurationUrlGet**](DefaultApi.md#OcpiChargingprofilesSessionIdDurationUrlGet) | **Get** /ocpi/chargingprofiles/{session_id}/{duration}/{url} | Get active charging profile
[**OcpiChargingprofilesSessionIdPut**](DefaultApi.md#OcpiChargingprofilesSessionIdPut) | **Put** /ocpi/chargingprofiles/{session_id} | Create or update charging profile on session
[**OcpiChargingprofilesSessionIdResponseUrlDelete**](DefaultApi.md#OcpiChargingprofilesSessionIdResponseUrlDelete) | **Delete** /ocpi/chargingprofiles/{session_id}/{response_url} | Delete charging profile on the session
[**OcpiClientinfoCountryCodePartyIdGet**](DefaultApi.md#OcpiClientinfoCountryCodePartyIdGet) | **Get** /ocpi/clientinfo/{country_code}/{party_id} | Get a ClientInfo object as it is stored in the connected clients system.
[**OcpiClientinfoCountryCodePartyIdPut**](DefaultApi.md#OcpiClientinfoCountryCodePartyIdPut) | **Put** /ocpi/clientinfo/{country_code}/{party_id} | Create or update client info
[**OcpiCommandsCommandPost**](DefaultApi.md#OcpiCommandsCommandPost) | **Post** /ocpi/commands/{command} | Send a command
[**OcpiCommandsCommandUidPost**](DefaultApi.md#OcpiCommandsCommandUidPost) | **Post** /ocpi/commands/{command}/{uid} | Receive the asynchronous response from the Charge Point.
[**OcpiHubclientinfoGet**](DefaultApi.md#OcpiHubclientinfoGet) | **Get** /ocpi/hubclientinfo | Get hub connected clients
[**OcpiLocationsCountryCodePartyIdLocationIdGet**](DefaultApi.md#OcpiLocationsCountryCodePartyIdLocationIdGet) | **Get** /ocpi/locations/{country_code}/{party_id}/{location_id} | Get location, EVSE or connector
[**OcpiLocationsCountryCodePartyIdLocationIdPatch**](DefaultApi.md#OcpiLocationsCountryCodePartyIdLocationIdPatch) | **Patch** /ocpi/locations/{country_code}/{party_id}/{location_id} | Sends new location object
[**OcpiLocationsCountryCodePartyIdLocationIdPut**](DefaultApi.md#OcpiLocationsCountryCodePartyIdLocationIdPut) | **Put** /ocpi/locations/{country_code}/{party_id}/{location_id} | Sends new location object
[**OcpiLocationsGet**](DefaultApi.md#OcpiLocationsGet) | **Get** /ocpi/locations | Get available locations
[**OcpiLocationsLocationIdGet**](DefaultApi.md#OcpiLocationsLocationIdGet) | **Get** /ocpi/locations/{location_id} | Get location, EVSE or connector
[**OcpiSessionsCountryCodePartyIdSessionIdGet**](DefaultApi.md#OcpiSessionsCountryCodePartyIdSessionIdGet) | **Get** /ocpi/sessions/{country_code}/{party_id}/{session_id} | Get session with session_id
[**OcpiSessionsCountryCodePartyIdSessionIdPatch**](DefaultApi.md#OcpiSessionsCountryCodePartyIdSessionIdPatch) | **Patch** /ocpi/sessions/{country_code}/{party_id}/{session_id} | Update session with session_id
[**OcpiSessionsCountryCodePartyIdSessionIdPut**](DefaultApi.md#OcpiSessionsCountryCodePartyIdSessionIdPut) | **Put** /ocpi/sessions/{country_code}/{party_id}/{session_id} | Update session with session_id
[**OcpiSessionsGet**](DefaultApi.md#OcpiSessionsGet) | **Get** /ocpi/sessions | Get charging sessions
[**OcpiSessionsSessionIdChargingPreferencesPut**](DefaultApi.md#OcpiSessionsSessionIdChargingPreferencesPut) | **Put** /ocpi/sessions/{session_id}/charging_preferences | Update charging session
[**OcpiTariffsCountryCodePartyIdTariffIdDelete**](DefaultApi.md#OcpiTariffsCountryCodePartyIdTariffIdDelete) | **Delete** /ocpi/tariffs/{country_code}/{party_id}/{tariff_id} | Delete tariff
[**OcpiTariffsCountryCodePartyIdTariffIdGet**](DefaultApi.md#OcpiTariffsCountryCodePartyIdTariffIdGet) | **Get** /ocpi/tariffs/{country_code}/{party_id}/{tariff_id} | Get tariff with tariff_id
[**OcpiTariffsCountryCodePartyIdTariffIdPut**](DefaultApi.md#OcpiTariffsCountryCodePartyIdTariffIdPut) | **Put** /ocpi/tariffs/{country_code}/{party_id}/{tariff_id} | Update tariff
[**OcpiTariffsGet**](DefaultApi.md#OcpiTariffsGet) | **Get** /ocpi/tariffs | Get tariffs
[**OcpiTokenUidAuthorizePost**](DefaultApi.md#OcpiTokenUidAuthorizePost) | **Post** /ocpi/{token_uid}/authorize | Authorization
[**OcpiTokensCountryCodePartyIdTokenUidGet**](DefaultApi.md#OcpiTokensCountryCodePartyIdTokenUidGet) | **Get** /ocpi/tokens/{country_code}/{party_id}/{token_uid} | Get the token object
[**OcpiTokensCountryCodePartyIdTokenUidPatch**](DefaultApi.md#OcpiTokensCountryCodePartyIdTokenUidPatch) | **Patch** /ocpi/tokens/{country_code}/{party_id}/{token_uid} | Update token object
[**OcpiTokensCountryCodePartyIdTokenUidPut**](DefaultApi.md#OcpiTokensCountryCodePartyIdTokenUidPut) | **Put** /ocpi/tokens/{country_code}/{party_id}/{token_uid} | Update token object
[**OcpiTokensGet**](DefaultApi.md#OcpiTokensGet) | **Get** /ocpi/tokens | Get tokens
[**OcpiVersionsGet**](DefaultApi.md#OcpiVersionsGet) | **Get** /ocpi/versions | Get supported versions



## Ocpi221CredentialsDelete

> Credentials Ocpi221CredentialsDelete(ctx).Execute()

Informs that credentials are invalid



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ocpi221CredentialsDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ocpi221CredentialsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ocpi221CredentialsDelete`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ocpi221CredentialsDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOcpi221CredentialsDeleteRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ocpi221CredentialsGet

> Credentials Ocpi221CredentialsGet(ctx).Execute()

Get the credentials object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ocpi221CredentialsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ocpi221CredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ocpi221CredentialsGet`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ocpi221CredentialsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOcpi221CredentialsGetRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ocpi221CredentialsPost

> Credentials Ocpi221CredentialsPost(ctx).Credentials(credentials).Execute()

Provides the credentials object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    credentials := *openapiclient.NewCredentials(float32(1000)) // Credentials |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ocpi221CredentialsPost(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ocpi221CredentialsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ocpi221CredentialsPost`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ocpi221CredentialsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpi221CredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ocpi221CredentialsPut

> Credentials Ocpi221CredentialsPut(ctx).Credentials(credentials).Execute()

Provides updated credentials object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    credentials := *openapiclient.NewCredentials(float32(1000)) // Credentials |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ocpi221CredentialsPut(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ocpi221CredentialsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ocpi221CredentialsPut`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ocpi221CredentialsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpi221CredentialsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) |  | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ocpi221DetailsGet

> Details Ocpi221DetailsGet(ctx).Execute()

Get version supported endpoints



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Ocpi221DetailsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ocpi221DetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ocpi221DetailsGet`: Details
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ocpi221DetailsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOcpi221DetailsGetRequest struct via the builder pattern


### Return type

[**Details**](Details.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiCdrsGet

> CdrsResponse OcpiCdrsGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get CDRs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return CDRs that have last_updated after or equal to this Date/Time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return CDRs that have last_updated up to this Date/Time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiCdrsGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiCdrsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiCdrsGet`: CdrsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiCdrsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiCdrsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return CDRs that have last_updated after or equal to this Date/Time (inclusive). | 
 **dateTo** | **string** | Return CDRs that have last_updated up to this Date/Time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**CdrsResponse**](CdrsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiCdrsPost

> CdrResponse OcpiCdrsPost(ctx).Cdr(cdr).Execute()

Create CDR object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    cdr := *openapiclient.NewCdr("CZ", "BEC", "12345", "2015-06-29T21:39:09Z", "2015-06-29T21:39:09Z", *openapiclient.NewCdrCdrToken("CZ", "STK", "123abc", "Type_example", "NL-TST-C12345678-S"), "AuthMethod_example", *openapiclient.NewCdrCdrLocation(), "EUR", *openapiclient.NewPrice(float32(2.5)), float32(15.342), float32(1.973), "2015-06-29T22:39:09Z") // Cdr |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiCdrsPost(context.Background()).Cdr(cdr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiCdrsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiCdrsPost`: CdrResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiCdrsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiCdrsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdr** | [**Cdr**](Cdr.md) |  | 

### Return type

[**CdrResponse**](CdrResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiChargingprofilePost

> GenericResponse OcpiChargingprofilePost(ctx).OcpiChargingprofilePostRequest(ocpiChargingprofilePostRequest).Execute()

Send a result of the original request by the eMSP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    ocpiChargingprofilePostRequest := openapiclient._ocpi_chargingprofile_post_request{ActiveChargingProfileResult: openapiclient.NewActiveChargingProfileResult("Result_example")} // OcpiChargingprofilePostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiChargingprofilePost(context.Background()).OcpiChargingprofilePostRequest(ocpiChargingprofilePostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiChargingprofilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiChargingprofilePost`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiChargingprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiChargingprofilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocpiChargingprofilePostRequest** | [**OcpiChargingprofilePostRequest**](OcpiChargingprofilePostRequest.md) |  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiChargingprofileSessionIdPut

> GenericResponse OcpiChargingprofileSessionIdPut(ctx, sessionId).ActiveChargingProfile(activeChargingProfile).Execute()

Update charging profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sessionId := "sessionId_example" // string | The unique id that identifies the session in the CPO platform.
    activeChargingProfile := *openapiclient.NewActiveChargingProfile("2019-06-23T08:16:02Z", *openapiclient.NewChargingProfile("ChargingRateUnit_example")) // ActiveChargingProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiChargingprofileSessionIdPut(context.Background(), sessionId).ActiveChargingProfile(activeChargingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiChargingprofileSessionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiChargingprofileSessionIdPut`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiChargingprofileSessionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The unique id that identifies the session in the CPO platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiChargingprofileSessionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeChargingProfile** | [**ActiveChargingProfile**](ActiveChargingProfile.md) |  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiChargingprofilesSessionIdDurationUrlGet

> ChargingProfilesResponse OcpiChargingprofilesSessionIdDurationUrlGet(ctx, sessionId, duration, url).Execute()

Get active charging profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sessionId := "sessionId_example" // string | The unique id that identifies the session in the CPO platform.
    duration := int32(56) // int32 | Length of the requested ActiveChargingProfile in seconds Duration in seconds.
    url := "url_example" // string | URL that the ActiveChargingProfileResult POST should be send to. This URL  might contain an unique ID to be able to distinguish between GET ActiveChargingProfile  requests.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiChargingprofilesSessionIdDurationUrlGet(context.Background(), sessionId, duration, url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiChargingprofilesSessionIdDurationUrlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiChargingprofilesSessionIdDurationUrlGet`: ChargingProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiChargingprofilesSessionIdDurationUrlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The unique id that identifies the session in the CPO platform. | 
**duration** | **int32** | Length of the requested ActiveChargingProfile in seconds Duration in seconds. | 
**url** | **string** | URL that the ActiveChargingProfileResult POST should be send to. This URL  might contain an unique ID to be able to distinguish between GET ActiveChargingProfile  requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiChargingprofilesSessionIdDurationUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ChargingProfilesResponse**](ChargingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiChargingprofilesSessionIdPut

> ChargingProfilesResponse OcpiChargingprofilesSessionIdPut(ctx, sessionId).SetChargingProfile(setChargingProfile).Execute()

Create or update charging profile on session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sessionId := "sessionId_example" // string | The unique id that identifies the session in the CPO platform.
    setChargingProfile := *openapiclient.NewSetChargingProfile(*openapiclient.NewChargingProfile("ChargingRateUnit_example"), "https://client.com/12345") // SetChargingProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiChargingprofilesSessionIdPut(context.Background(), sessionId).SetChargingProfile(setChargingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiChargingprofilesSessionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiChargingprofilesSessionIdPut`: ChargingProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiChargingprofilesSessionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The unique id that identifies the session in the CPO platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiChargingprofilesSessionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setChargingProfile** | [**SetChargingProfile**](SetChargingProfile.md) |  | 

### Return type

[**ChargingProfilesResponse**](ChargingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiChargingprofilesSessionIdResponseUrlDelete

> ChargingProfilesResponse OcpiChargingprofilesSessionIdResponseUrlDelete(ctx, sessionId, responseUrl).Execute()

Delete charging profile on the session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sessionId := "sessionId_example" // string | The unique id that identifies the session in the CPO platform.
    responseUrl := "responseUrl_example" // string | URL that the ClearProfileResult POST should be send to. This URL  might contain an unique ID to be able to distinguish between DELETE  ChargingProfile requests.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiChargingprofilesSessionIdResponseUrlDelete(context.Background(), sessionId, responseUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiChargingprofilesSessionIdResponseUrlDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiChargingprofilesSessionIdResponseUrlDelete`: ChargingProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiChargingprofilesSessionIdResponseUrlDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The unique id that identifies the session in the CPO platform. | 
**responseUrl** | **string** | URL that the ClearProfileResult POST should be send to. This URL  might contain an unique ID to be able to distinguish between DELETE  ChargingProfile requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiChargingprofilesSessionIdResponseUrlDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChargingProfilesResponse**](ChargingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiClientinfoCountryCodePartyIdGet

> ClientsInfoResponse OcpiClientinfoCountryCodePartyIdGet(ctx, countryCode, partyId).Execute()

Get a ClientInfo object as it is stored in the connected clients system.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the requested ClientInfo object.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the requested ClientInfo object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiClientinfoCountryCodePartyIdGet(context.Background(), countryCode, partyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiClientinfoCountryCodePartyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiClientinfoCountryCodePartyIdGet`: ClientsInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiClientinfoCountryCodePartyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the requested ClientInfo object. | 
**partyId** | **string** | Party ID (Provider ID) of the requested ClientInfo object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiClientinfoCountryCodePartyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientsInfoResponse**](ClientsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiClientinfoCountryCodePartyIdPut

> ClientsInfoResponse OcpiClientinfoCountryCodePartyIdPut(ctx, countryCode, partyId).ClientInfo(clientInfo).Execute()

Create or update client info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the eMSP sending this PUT request to the CPO system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the eMSP sending this PUT request to the CPO system.
    clientInfo := *openapiclient.NewClientInfo("TNM", "DE", "Role_example", "Status_example", "2019-06-23T08:16:02Z") // ClientInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiClientinfoCountryCodePartyIdPut(context.Background(), countryCode, partyId).ClientInfo(clientInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiClientinfoCountryCodePartyIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiClientinfoCountryCodePartyIdPut`: ClientsInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiClientinfoCountryCodePartyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the eMSP sending this PUT request to the CPO system. | 
**partyId** | **string** | Party ID (Provider ID) of the eMSP sending this PUT request to the CPO system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiClientinfoCountryCodePartyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientInfo** | [**ClientInfo**](ClientInfo.md) |  | 

### Return type

[**ClientsInfoResponse**](ClientsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiCommandsCommandPost

> CommandResponse OcpiCommandsCommandPost(ctx, command).OcpiCommandsCommandPostRequest(ocpiCommandsCommandPostRequest).Execute()

Send a command



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    command := "command_example" // string | Type of command that is requested.
    ocpiCommandsCommandPostRequest := openapiclient._ocpi_commands__command__post_request{CancelReservation: openapiclient.NewCancelReservation("ResponseUrl_example", "ReservationId_example")} // OcpiCommandsCommandPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiCommandsCommandPost(context.Background(), command).OcpiCommandsCommandPostRequest(ocpiCommandsCommandPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiCommandsCommandPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiCommandsCommandPost`: CommandResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiCommandsCommandPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** | Type of command that is requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiCommandsCommandPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ocpiCommandsCommandPostRequest** | [**OcpiCommandsCommandPostRequest**](OcpiCommandsCommandPostRequest.md) |  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiCommandsCommandUidPost

> CommandResponse OcpiCommandsCommandUidPost(ctx, command, uid).CommandResult(commandResult).Execute()

Receive the asynchronous response from the Charge Point.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    command := "command_example" // string | Type of command that is requested.
    uid := "uid_example" // string | 
    commandResult := *openapiclient.NewCommandResult("Result_example") // CommandResult | Result of the command request, from the Charge Point. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiCommandsCommandUidPost(context.Background(), command, uid).CommandResult(commandResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiCommandsCommandUidPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiCommandsCommandUidPost`: CommandResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiCommandsCommandUidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** | Type of command that is requested. | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiCommandsCommandUidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commandResult** | [**CommandResult**](CommandResult.md) | Result of the command request, from the Charge Point. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiHubclientinfoGet

> ClientInfoResponse OcpiHubclientinfoGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get hub connected clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return ClientInfo that have last_updated after or equal to Date/Time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return ClientInfo that have last_updated up to Date/Time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiHubclientinfoGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiHubclientinfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiHubclientinfoGet`: ClientInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiHubclientinfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiHubclientinfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return ClientInfo that have last_updated after or equal to Date/Time (inclusive). | 
 **dateTo** | **string** | Return ClientInfo that have last_updated up to Date/Time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**ClientInfoResponse**](ClientInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiLocationsCountryCodePartyIdLocationIdGet

> OcpiLocationsLocationIdGet200Response OcpiLocationsCountryCodePartyIdLocationIdGet(ctx, countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()

Get location, EVSE or connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO requesting data from the eMSP system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO requesting data from the eMSP system.
    locationId := "locationId_example" // string | Location.id of the Location object to retrieve.
    evseUid := "evseUid_example" // string | Evse.uid, required when requesting an EVSE or Connector object. (optional)
    connectorId := "connectorId_example" // string | Connector.id, required when requesting a Connector object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdGet(context.Background(), countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiLocationsCountryCodePartyIdLocationIdGet`: OcpiLocationsLocationIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO requesting data from the eMSP system. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO requesting data from the eMSP system. | 
**locationId** | **string** | Location.id of the Location object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiLocationsCountryCodePartyIdLocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evseUid** | **string** | Evse.uid, required when requesting an EVSE or Connector object. | 
 **connectorId** | **string** | Connector.id, required when requesting a Connector object. | 

### Return type

[**OcpiLocationsLocationIdGet200Response**](OcpiLocationsLocationIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiLocationsCountryCodePartyIdLocationIdPatch

> OcpiLocationsLocationIdGet200Response OcpiLocationsCountryCodePartyIdLocationIdPatch(ctx, countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()

Sends new location object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO requesting data from the eMSP system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO requesting data from the eMSP system.
    locationId := "locationId_example" // string | Location.id of the Location object to retrieve.
    evseUid := "evseUid_example" // string | Evse.uid, required when requesting an EVSE or Connector object. (optional)
    connectorId := "connectorId_example" // string | Connector.id, required when requesting a Connector object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPatch(context.Background(), countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiLocationsCountryCodePartyIdLocationIdPatch`: OcpiLocationsLocationIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO requesting data from the eMSP system. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO requesting data from the eMSP system. | 
**locationId** | **string** | Location.id of the Location object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiLocationsCountryCodePartyIdLocationIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evseUid** | **string** | Evse.uid, required when requesting an EVSE or Connector object. | 
 **connectorId** | **string** | Connector.id, required when requesting a Connector object. | 

### Return type

[**OcpiLocationsLocationIdGet200Response**](OcpiLocationsLocationIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiLocationsCountryCodePartyIdLocationIdPut

> OcpiLocationsLocationIdGet200Response OcpiLocationsCountryCodePartyIdLocationIdPut(ctx, countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).OcpiLocationsLocationIdGet200Response(ocpiLocationsLocationIdGet200Response).Execute()

Sends new location object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO requesting data from the eMSP system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO requesting data from the eMSP system.
    locationId := "locationId_example" // string | Location.id of the Location object to retrieve.
    evseUid := "evseUid_example" // string | Evse.uid, required when requesting an EVSE or Connector object. (optional)
    connectorId := "connectorId_example" // string | Connector.id, required when requesting a Connector object. (optional)
    ocpiLocationsLocationIdGet200Response := openapiclient._ocpi_locations__location_id__get_200_response{Connector: openapiclient.NewConnector("1", "Standard_example", "Format_example", "PowerType_example", int32(220), int32(16), "2015-03-16T10:10:02Z")} // OcpiLocationsLocationIdGet200Response |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPut(context.Background(), countryCode, partyId, locationId).EvseUid(evseUid).ConnectorId(connectorId).OcpiLocationsLocationIdGet200Response(ocpiLocationsLocationIdGet200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiLocationsCountryCodePartyIdLocationIdPut`: OcpiLocationsLocationIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiLocationsCountryCodePartyIdLocationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO requesting data from the eMSP system. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO requesting data from the eMSP system. | 
**locationId** | **string** | Location.id of the Location object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiLocationsCountryCodePartyIdLocationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evseUid** | **string** | Evse.uid, required when requesting an EVSE or Connector object. | 
 **connectorId** | **string** | Connector.id, required when requesting a Connector object. | 
 **ocpiLocationsLocationIdGet200Response** | [**OcpiLocationsLocationIdGet200Response**](OcpiLocationsLocationIdGet200Response.md) |  | 

### Return type

[**OcpiLocationsLocationIdGet200Response**](OcpiLocationsLocationIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiLocationsGet

> Locations OcpiLocationsGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get available locations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return Locations that have last_updated after or equal to this date time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return Locations that have last_updated up to this date time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiLocationsGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiLocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiLocationsGet`: Locations
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiLocationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return Locations that have last_updated after or equal to this date time (inclusive). | 
 **dateTo** | **string** | Return Locations that have last_updated up to this date time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**Locations**](Locations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiLocationsLocationIdGet

> OcpiLocationsLocationIdGet200Response OcpiLocationsLocationIdGet(ctx, locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()

Get location, EVSE or connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    locationId := "locationId_example" // string | Location.id of the Location object to retrieve.
    evseUid := "evseUid_example" // string | Evse.uid, required when requesting an EVSE or Connector object. (optional)
    connectorId := "connectorId_example" // string | Connector.id, required when requesting a Connector object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiLocationsLocationIdGet(context.Background(), locationId).EvseUid(evseUid).ConnectorId(connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiLocationsLocationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiLocationsLocationIdGet`: OcpiLocationsLocationIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiLocationsLocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | Location.id of the Location object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiLocationsLocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evseUid** | **string** | Evse.uid, required when requesting an EVSE or Connector object. | 
 **connectorId** | **string** | Connector.id, required when requesting a Connector object. | 

### Return type

[**OcpiLocationsLocationIdGet200Response**](OcpiLocationsLocationIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiSessionsCountryCodePartyIdSessionIdGet

> SessionResponse OcpiSessionsCountryCodePartyIdSessionIdGet(ctx, countryCode, partyId, sessionId).Execute()

Get session with session_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing the GET on the eMSP's system
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing the GET on the eMSP's system.
    sessionId := "sessionId_example" // string | Id of the Session object to get from the eMSP's system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdGet(context.Background(), countryCode, partyId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiSessionsCountryCodePartyIdSessionIdGet`: SessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing the GET on the eMSP&#39;s system | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing the GET on the eMSP&#39;s system. | 
**sessionId** | **string** | Id of the Session object to get from the eMSP&#39;s system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiSessionsCountryCodePartyIdSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SessionResponse**](SessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiSessionsCountryCodePartyIdSessionIdPatch

> SessionResponse OcpiSessionsCountryCodePartyIdSessionIdPatch(ctx, countryCode, partyId, sessionId).Execute()

Update session with session_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing this PATCH on the eMSP's system.  This SHALL be the same value as the country_code in the Session object being pushed.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing this PATCH on the eMSP's  system. This SHALL be the same value as the party_id in the Session object being  pushed.
    sessionId := "sessionId_example" // string | Id of the new or updated Session object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPatch(context.Background(), countryCode, partyId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiSessionsCountryCodePartyIdSessionIdPatch`: SessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing this PATCH on the eMSP&#39;s system.  This SHALL be the same value as the country_code in the Session object being pushed. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing this PATCH on the eMSP&#39;s  system. This SHALL be the same value as the party_id in the Session object being  pushed. | 
**sessionId** | **string** | Id of the new or updated Session object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiSessionsCountryCodePartyIdSessionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SessionResponse**](SessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiSessionsCountryCodePartyIdSessionIdPut

> SessionResponse OcpiSessionsCountryCodePartyIdSessionIdPut(ctx, countryCode, partyId, sessionId).Session(session).Execute()

Update session with session_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing this PUT on the eMSP's system.  This SHALL be the same value as the country_code in the Session object being pushed.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing this PUT on the eMSP's  system. This SHALL be the same value as the party_id in the Session object being  pushed.
    sessionId := "sessionId_example" // string | Id of the new or updated Session object.
    session := *openapiclient.NewSession("CZ", "STK", "101", "2020-03-09T10:17:09Z", float32(0.0), *openapiclient.NewCdrCdrToken("CZ", "STK", "123abc", "Type_example", "NL-TST-C12345678-S"), "AuthMethod_example", "LOC1", "3256", "1", "EUR", "Status_example", "2019-06-24T12:39:09Z") // Session |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPut(context.Background(), countryCode, partyId, sessionId).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiSessionsCountryCodePartyIdSessionIdPut`: SessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiSessionsCountryCodePartyIdSessionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing this PUT on the eMSP&#39;s system.  This SHALL be the same value as the country_code in the Session object being pushed. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing this PUT on the eMSP&#39;s  system. This SHALL be the same value as the party_id in the Session object being  pushed. | 
**sessionId** | **string** | Id of the new or updated Session object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiSessionsCountryCodePartyIdSessionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **session** | [**Session**](Session.md) |  | 

### Return type

[**SessionResponse**](SessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiSessionsGet

> SessionsResponse OcpiSessionsGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get charging sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return Sessions that have last_updated after or equal to this date time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return Sessions that have last_updated up to this date time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiSessionsGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiSessionsGet`: SessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiSessionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return Sessions that have last_updated after or equal to this date time (inclusive). | 
 **dateTo** | **string** | Return Sessions that have last_updated up to this date time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**SessionsResponse**](SessionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiSessionsSessionIdChargingPreferencesPut

> ChargingPreferencesResponse OcpiSessionsSessionIdChargingPreferencesPut(ctx, sessionId).ChargingPreferences(chargingPreferences).Execute()

Update charging session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sessionId := "sessionId_example" // string | Session.id of the Session for which the Charging Preferences are to be set.
    chargingPreferences := *openapiclient.NewChargingPreferences("ProfileType_example") // ChargingPreferences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiSessionsSessionIdChargingPreferencesPut(context.Background(), sessionId).ChargingPreferences(chargingPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiSessionsSessionIdChargingPreferencesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiSessionsSessionIdChargingPreferencesPut`: ChargingPreferencesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiSessionsSessionIdChargingPreferencesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session.id of the Session for which the Charging Preferences are to be set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiSessionsSessionIdChargingPreferencesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargingPreferences** | [**ChargingPreferences**](ChargingPreferences.md) |  | 

### Return type

[**ChargingPreferencesResponse**](ChargingPreferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTariffsCountryCodePartyIdTariffIdDelete

> TariffDeleteResponse OcpiTariffsCountryCodePartyIdTariffIdDelete(ctx, countryCode, partyId, tariffId).Execute()

Delete tariff



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing the PUT request on the eMSP's system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing the PUT request on the eMSP's system.
    tariffId := "tariffId_example" // string | Tariff.id of the Tariff object to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdDelete(context.Background(), countryCode, partyId, tariffId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTariffsCountryCodePartyIdTariffIdDelete`: TariffDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing the PUT request on the eMSP&#39;s system. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing the PUT request on the eMSP&#39;s system. | 
**tariffId** | **string** | Tariff.id of the Tariff object to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTariffsCountryCodePartyIdTariffIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TariffDeleteResponse**](TariffDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTariffsCountryCodePartyIdTariffIdGet

> TariffResponse OcpiTariffsCountryCodePartyIdTariffIdGet(ctx, countryCode, partyId, tariffId).Execute()

Get tariff with tariff_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing the GET on the eMSP's system
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing the GET on the eMSP's system.
    tariffId := "tariffId_example" // string | Tariff.id of the Tariff object to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdGet(context.Background(), countryCode, partyId, tariffId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTariffsCountryCodePartyIdTariffIdGet`: TariffResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing the GET on the eMSP&#39;s system | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing the GET on the eMSP&#39;s system. | 
**tariffId** | **string** | Tariff.id of the Tariff object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTariffsCountryCodePartyIdTariffIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TariffResponse**](TariffResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTariffsCountryCodePartyIdTariffIdPut

> TariffResponse OcpiTariffsCountryCodePartyIdTariffIdPut(ctx, countryCode, partyId, tariffId).Tariff(tariff).Execute()

Update tariff



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the CPO performing the PUT request on the eMSP's system.  This SHALL be the same value as the country_code in the Tariff object being pushed.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the CPO performing the PUT request on the eMSP's system. This SHALL be the same value as the party_id in the Tariff object being pushed.
    tariffId := "tariffId_example" // string | Tariff.id of the Tariff object to create or replace.
    tariff := *openapiclient.NewTariff("BE", "BEC", "12", "EUR", "2015-06-29T22:39:09Z") // Tariff |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdPut(context.Background(), countryCode, partyId, tariffId).Tariff(tariff).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTariffsCountryCodePartyIdTariffIdPut`: TariffResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTariffsCountryCodePartyIdTariffIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the CPO performing the PUT request on the eMSP&#39;s system.  This SHALL be the same value as the country_code in the Tariff object being pushed. | 
**partyId** | **string** | Party ID (Provider ID) of the CPO performing the PUT request on the eMSP&#39;s system. This SHALL be the same value as the party_id in the Tariff object being pushed. | 
**tariffId** | **string** | Tariff.id of the Tariff object to create or replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTariffsCountryCodePartyIdTariffIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tariff** | [**Tariff**](Tariff.md) |  | 

### Return type

[**TariffResponse**](TariffResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTariffsGet

> TariffsResponse OcpiTariffsGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get tariffs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return Tariffs that have last_updated after or equal to Date/Time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return Tariffs that have last_updated up to Date/Time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTariffsGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTariffsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTariffsGet`: TariffsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTariffsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTariffsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return Tariffs that have last_updated after or equal to Date/Time (inclusive). | 
 **dateTo** | **string** | Return Tariffs that have last_updated up to Date/Time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**TariffsResponse**](TariffsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTokenUidAuthorizePost

> AuthorizationInfo OcpiTokenUidAuthorizePost(ctx, tokenUid).Type_(type_).LocationReferences(locationReferences).Execute()

Authorization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    tokenUid := "tokenUid_example" // string | Token.uid of the Token for which authorization is requested.
    type_ := "type__example" // string | Token.type of the Token to retrieve. Default if omitted: RFID (optional)
    locationReferences := *openapiclient.NewLocationReferences("LOC1") // LocationReferences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTokenUidAuthorizePost(context.Background(), tokenUid).Type_(type_).LocationReferences(locationReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTokenUidAuthorizePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTokenUidAuthorizePost`: AuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTokenUidAuthorizePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenUid** | **string** | Token.uid of the Token for which authorization is requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTokenUidAuthorizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Token.type of the Token to retrieve. Default if omitted: RFID | 
 **locationReferences** | [**LocationReferences**](LocationReferences.md) |  | 

### Return type

[**AuthorizationInfo**](AuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTokensCountryCodePartyIdTokenUidGet

> TokenResponse OcpiTokensCountryCodePartyIdTokenUidGet(ctx, countryCode, partyId, tokenUid).Type_(type_).Execute()

Get the token object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the eMSP requesting this GET from the CPO system.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the eMSP requesting this GET from the CPO system.
    tokenUid := "tokenUid_example" // string | Token.uid of the Token object to retrieve.
    type_ := "type__example" // string | Token.type of the Token to retrieve. Default if omitted: RFID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTokensCountryCodePartyIdTokenUidGet(context.Background(), countryCode, partyId, tokenUid).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTokensCountryCodePartyIdTokenUidGet`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the eMSP requesting this GET from the CPO system. | 
**partyId** | **string** | Party ID (Provider ID) of the eMSP requesting this GET from the CPO system. | 
**tokenUid** | **string** | Token.uid of the Token object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTokensCountryCodePartyIdTokenUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Token.type of the Token to retrieve. Default if omitted: RFID | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTokensCountryCodePartyIdTokenUidPatch

> TokenResponse OcpiTokensCountryCodePartyIdTokenUidPatch(ctx, countryCode, partyId, tokenUid).Type_(type_).Execute()

Update token object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the eMSP sending this PATCH request to the CPO system.  This SHALL be the same value as the country_code in the Token object being pushed.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the eMSP sending this PATCH request to the CPO  system. This SHALL be the same value as the party_id in the Token object being pushed.
    tokenUid := "tokenUid_example" // string | Token.uid of the Token object to retrieve.
    type_ := "type__example" // string | Token.type of the Token to retrieve. Default if omitted: RFID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPatch(context.Background(), countryCode, partyId, tokenUid).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTokensCountryCodePartyIdTokenUidPatch`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the eMSP sending this PATCH request to the CPO system.  This SHALL be the same value as the country_code in the Token object being pushed. | 
**partyId** | **string** | Party ID (Provider ID) of the eMSP sending this PATCH request to the CPO  system. This SHALL be the same value as the party_id in the Token object being pushed. | 
**tokenUid** | **string** | Token.uid of the Token object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTokensCountryCodePartyIdTokenUidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Token.type of the Token to retrieve. Default if omitted: RFID | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTokensCountryCodePartyIdTokenUidPut

> TokenResponse OcpiTokensCountryCodePartyIdTokenUidPut(ctx, countryCode, partyId, tokenUid).Type_(type_).Token(token).Execute()

Update token object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    countryCode := "countryCode_example" // string | Country code of the eMSP sending this PUT request to the CPO system.  This SHALL be the same value as the country_code in the Token object being pushed.
    partyId := "partyId_example" // string | Party ID (Provider ID) of the eMSP sending this PUT request to the CPO  system. This SHALL be the same value as the party_id in the Token object being pushed.
    tokenUid := "tokenUid_example" // string | Token.uid of the Token object to retrieve.
    type_ := "type__example" // string | Token.type of the Token to retrieve. Default if omitted: RFID (optional)
    token := *openapiclient.NewToken("DE", "TNM", "bdf21bce-fc97-11e8-8eb2-f2801f1b9fd1", "Type_example", "DE8ACC12E46L89", "TheNewMotion", true, "Whitelist_example", "2018-12-10T17:25:10Z") // Token |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPut(context.Background(), countryCode, partyId, tokenUid).Type_(type_).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTokensCountryCodePartyIdTokenUidPut`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTokensCountryCodePartyIdTokenUidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country code of the eMSP sending this PUT request to the CPO system.  This SHALL be the same value as the country_code in the Token object being pushed. | 
**partyId** | **string** | Party ID (Provider ID) of the eMSP sending this PUT request to the CPO  system. This SHALL be the same value as the party_id in the Token object being pushed. | 
**tokenUid** | **string** | Token.uid of the Token object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTokensCountryCodePartyIdTokenUidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Token.type of the Token to retrieve. Default if omitted: RFID | 
 **token** | [**Token**](Token.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiTokensGet

> TokensResponse OcpiTokensGet(ctx).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()

Get tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dateFrom := "dateFrom_example" // string | Return tokens that have last_updated after or equal to this Date/Time (inclusive). (optional)
    dateTo := "dateTo_example" // string | Return tokens that have last_updated up to Date/Time, but not including (exclusive). (optional)
    offset := int32(56) // int32 | The offset of the first object returned. Default is 0. (optional)
    limit := int32(56) // int32 | Maximum number of objects to GET. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiTokensGet(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiTokensGet`: TokensResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOcpiTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Return tokens that have last_updated after or equal to this Date/Time (inclusive). | 
 **dateTo** | **string** | Return tokens that have last_updated up to Date/Time, but not including (exclusive). | 
 **offset** | **int32** | The offset of the first object returned. Default is 0. | 
 **limit** | **int32** | Maximum number of objects to GET. | 

### Return type

[**TokensResponse**](TokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OcpiVersionsGet

> Versions OcpiVersionsGet(ctx).Execute()

Get supported versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OcpiVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OcpiVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OcpiVersionsGet`: Versions
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OcpiVersionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOcpiVersionsGetRequest struct via the builder pattern


### Return type

[**Versions**](Versions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

