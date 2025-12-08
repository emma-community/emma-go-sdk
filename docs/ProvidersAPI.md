# \ProvidersAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvider**](ProvidersAPI.md#GetProvider) | **Get** /v1/providers/{providerId} | Get cloud provider by ID
[**GetProviders**](ProvidersAPI.md#GetProviders) | **Get** /v1/providers | Get list of cloud providers



## GetProvider

> Provider GetProvider(ctx, providerId).ProjectId(projectId).Execute()

Get cloud provider by ID



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
	providerId := int32(56) // int32 | ID of the cloud provider
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.GetProvider(context.Background(), providerId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | ID of the cloud provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**Provider**](Provider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviders

> []Provider GetProviders(ctx).ProviderName(providerName).ProjectId(projectId).Execute()

Get list of cloud providers



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
	providerName := "GCP" // string | Name of the cloud provider (optional)
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.GetProviders(context.Background()).ProviderName(providerName).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.GetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviders`: []Provider
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** | Name of the cloud provider | 
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]Provider**](Provider.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

