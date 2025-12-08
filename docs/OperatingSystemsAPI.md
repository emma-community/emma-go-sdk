# \OperatingSystemsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOperatingSystem**](OperatingSystemsAPI.md#GetOperatingSystem) | **Get** /v1/operating-systems/{operatingSystemId} | Get operating system by ID
[**GetOperatingSystems**](OperatingSystemsAPI.md#GetOperatingSystems) | **Get** /v1/operating-systems | Get list of operating systems



## GetOperatingSystem

> OperatingSystem GetOperatingSystem(ctx, operatingSystemId).ProjectId(projectId).Execute()

Get operating system by ID



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
	operatingSystemId := int32(56) // int32 | ID of the operating system
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperatingSystemsAPI.GetOperatingSystem(context.Background(), operatingSystemId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperatingSystemsAPI.GetOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystem`: OperatingSystem
	fmt.Fprintf(os.Stdout, "Response from `OperatingSystemsAPI.GetOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **int32** | ID of the operating system | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**OperatingSystem**](OperatingSystem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystems

> []OperatingSystem GetOperatingSystems(ctx).Type_(type_).Architecture(architecture).Version(version).ProjectId(projectId).Execute()

Get list of operating systems



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
	type_ := "Ubuntu" // string | Type of the operating system (optional)
	architecture := "x86-64" // string | Architecture of the operating system (optional)
	version := "18.04" // string | Version of the operating system (optional)
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperatingSystemsAPI.GetOperatingSystems(context.Background()).Type_(type_).Architecture(architecture).Version(version).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperatingSystemsAPI.GetOperatingSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystems`: []OperatingSystem
	fmt.Fprintf(os.Stdout, "Response from `OperatingSystemsAPI.GetOperatingSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of the operating system | 
 **architecture** | **string** | Architecture of the operating system | 
 **version** | **string** | Version of the operating system | 
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]OperatingSystem**](OperatingSystem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

