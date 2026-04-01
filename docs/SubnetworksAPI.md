# \SubnetworksAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubnetwork**](SubnetworksAPI.md#GetSubnetwork) | **Get** /v1/subnetworks/{subnetworkId} | Get subnetwork by ID
[**GetSubnetworks**](SubnetworksAPI.md#GetSubnetworks) | **Get** /v1/subnetworks | Get list of subnetworks
[**SubnetworkCreate**](SubnetworksAPI.md#SubnetworkCreate) | **Post** /v1/subnetworks | Create subnetwork
[**SubnetworkDelete**](SubnetworksAPI.md#SubnetworkDelete) | **Delete** /v1/subnetworks/{subnetworkId} | Delete subnetwork
[**SubnetworkUpdate**](SubnetworksAPI.md#SubnetworkUpdate) | **Put** /v1/subnetworks/{subnetworkId} | Update subnetwork



## GetSubnetwork

> Subnetwork GetSubnetwork(ctx, subnetworkId).ProjectId(projectId).Execute()

Get subnetwork by ID



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
	subnetworkId := "subnetworkId_example" // string | Subnetwork ID
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.GetSubnetwork(context.Background(), subnetworkId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.GetSubnetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnetwork`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.GetSubnetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**Subnetwork**](Subnetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnetworks

> []Subnetwork GetSubnetworks(ctx).ProjectId(projectId).Execute()

Get list of subnetworks



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
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.GetSubnetworks(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.GetSubnetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnetworks`: []Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.GetSubnetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]Subnetwork**](Subnetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubnetworkCreate

> Subnetwork SubnetworkCreate(ctx).SubnetworkCreate(subnetworkCreate).Execute()

Create subnetwork



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
	subnetworkCreate := *openapiclient.NewSubnetworkCreate("gcore-frankfurt", int32(24)) // SubnetworkCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.SubnetworkCreate(context.Background()).SubnetworkCreate(subnetworkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.SubnetworkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetworkCreate`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.SubnetworkCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubnetworkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetworkCreate** | [**SubnetworkCreate**](SubnetworkCreate.md) |  | 

### Return type

[**Subnetwork**](Subnetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubnetworkDelete

> Subnetwork SubnetworkDelete(ctx, subnetworkId).Execute()

Delete subnetwork



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
	subnetworkId := "subnetworkId_example" // string | Subnetwork ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.SubnetworkDelete(context.Background(), subnetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.SubnetworkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetworkDelete`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.SubnetworkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetworkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subnetwork**](Subnetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubnetworkUpdate

> Subnetwork SubnetworkUpdate(ctx, subnetworkId).SubnetworkEdit(subnetworkEdit).Execute()

Update subnetwork



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
	subnetworkId := "subnetworkId_example" // string | Subnetwork ID
	subnetworkEdit := *openapiclient.NewSubnetworkEdit() // SubnetworkEdit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.SubnetworkUpdate(context.Background(), subnetworkId).SubnetworkEdit(subnetworkEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.SubnetworkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubnetworkUpdate`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.SubnetworkUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetworkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetworkEdit** | [**SubnetworkEdit**](SubnetworkEdit.md) |  | 

### Return type

[**Subnetwork**](Subnetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

