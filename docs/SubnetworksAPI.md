# \SubnetworksAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SubnetworksGet**](SubnetworksAPI.md#V1SubnetworksGet) | **Get** /v1/subnetworks | Get list of subnetworks
[**V1SubnetworksPost**](SubnetworksAPI.md#V1SubnetworksPost) | **Post** /v1/subnetworks | Create subnetwork
[**V1SubnetworksSubnetworkIdDelete**](SubnetworksAPI.md#V1SubnetworksSubnetworkIdDelete) | **Delete** /v1/subnetworks/{subnetworkId} | Delete subnetwork
[**V1SubnetworksSubnetworkIdGet**](SubnetworksAPI.md#V1SubnetworksSubnetworkIdGet) | **Get** /v1/subnetworks/{subnetworkId} | Get subnetwork by ID
[**V1SubnetworksSubnetworkIdPut**](SubnetworksAPI.md#V1SubnetworksSubnetworkIdPut) | **Put** /v1/subnetworks/{subnetworkId} | Update subnetwork



## V1SubnetworksGet

> []Subnetwork V1SubnetworksGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.V1SubnetworksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.V1SubnetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubnetworksGet`: []Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.V1SubnetworksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubnetworksGetRequest struct via the builder pattern


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


## V1SubnetworksPost

> Subnetwork V1SubnetworksPost(ctx).SubnetworkCreate(subnetworkCreate).Execute()

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
	resp, r, err := apiClient.SubnetworksAPI.V1SubnetworksPost(context.Background()).SubnetworkCreate(subnetworkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.V1SubnetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubnetworksPost`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.V1SubnetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SubnetworksPostRequest struct via the builder pattern


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


## V1SubnetworksSubnetworkIdDelete

> Subnetwork V1SubnetworksSubnetworkIdDelete(ctx, subnetworkId).Execute()

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
	resp, r, err := apiClient.SubnetworksAPI.V1SubnetworksSubnetworkIdDelete(context.Background(), subnetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.V1SubnetworksSubnetworkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubnetworksSubnetworkIdDelete`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.V1SubnetworksSubnetworkIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubnetworksSubnetworkIdDeleteRequest struct via the builder pattern


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


## V1SubnetworksSubnetworkIdGet

> Subnetwork V1SubnetworksSubnetworkIdGet(ctx, subnetworkId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetworksAPI.V1SubnetworksSubnetworkIdGet(context.Background(), subnetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.V1SubnetworksSubnetworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubnetworksSubnetworkIdGet`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.V1SubnetworksSubnetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubnetworksSubnetworkIdGetRequest struct via the builder pattern


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


## V1SubnetworksSubnetworkIdPut

> Subnetwork V1SubnetworksSubnetworkIdPut(ctx, subnetworkId).SubnetworkEdit(subnetworkEdit).Execute()

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
	resp, r, err := apiClient.SubnetworksAPI.V1SubnetworksSubnetworkIdPut(context.Background(), subnetworkId).SubnetworkEdit(subnetworkEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetworksAPI.V1SubnetworksSubnetworkIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubnetworksSubnetworkIdPut`: Subnetwork
	fmt.Fprintf(os.Stdout, "Response from `SubnetworksAPI.V1SubnetworksSubnetworkIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetworkId** | **string** | Subnetwork ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubnetworksSubnetworkIdPutRequest struct via the builder pattern


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

