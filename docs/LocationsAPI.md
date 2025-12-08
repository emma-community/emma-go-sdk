# \LocationsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocation**](LocationsAPI.md#GetLocation) | **Get** /v1/locations/{locationId} | Get location by ID
[**GetLocations**](LocationsAPI.md#GetLocations) | **Get** /v1/locations | Get list of locations



## GetLocation

> Location GetLocation(ctx, locationId).ProjectId(projectId).Execute()

Get location by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	locationId := int32(56) // int32 | ID of the geographic location
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.GetLocation(context.Background(), locationId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.GetLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocation`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **int32** | ID of the geographic location | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**Location**](Location.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocations

> []Location GetLocations(ctx).Name(name).ProjectId(projectId).Execute()

Get list of locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emma-community/emma-go-sdk"
)

func main() {
	name := "Stockholm" // string | Name of the geographic location (optional)
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.GetLocations(context.Background()).Name(name).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.GetLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocations`: []Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the geographic location | 
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]Location**](Location.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

