# \SpotInstancesAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpot**](SpotInstancesAPI.md#GetSpot) | **Get** /v1/spot-instances/{spotInstanceId} | Get spot instance by ID
[**GetSpots**](SpotInstancesAPI.md#GetSpots) | **Get** /v1/spot-instances | Get list of spot instances
[**SpotActions**](SpotInstancesAPI.md#SpotActions) | **Post** /v1/spot-instances/{spotInstanceId}/actions | Perform actions with a spot instance
[**SpotCreate**](SpotInstancesAPI.md#SpotCreate) | **Post** /v1/spot-instances | Create spot instance
[**SpotDelete**](SpotInstancesAPI.md#SpotDelete) | **Delete** /v1/spot-instances/{spotInstanceId} | Delete spot instance



## GetSpot

> Vm GetSpot(ctx, spotInstanceId).Execute()

Get spot instance by ID



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
	spotInstanceId := int32(56) // int32 | ID of the spot instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotInstancesAPI.GetSpot(context.Background(), spotInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotInstancesAPI.GetSpot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpot`: Vm
	fmt.Fprintf(os.Stdout, "Response from `SpotInstancesAPI.GetSpot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotInstanceId** | **int32** | ID of the spot instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpots

> []Vm GetSpots(ctx).Execute()

Get list of spot instances



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
	resp, r, err := apiClient.SpotInstancesAPI.GetSpots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotInstancesAPI.GetSpots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpots`: []Vm
	fmt.Fprintf(os.Stdout, "Response from `SpotInstancesAPI.GetSpots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpotsRequest struct via the builder pattern


### Return type

[**[]Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotActions

> Vm SpotActions(ctx, spotInstanceId).SpotActionsRequest(spotActionsRequest).Execute()

Perform actions with a spot instance



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
	spotInstanceId := int32(56) // int32 | ID of the spot instance
	spotActionsRequest := openapiclient.SpotActions_request{SpotReboot: openapiclient.NewSpotReboot("reboot")} // SpotActionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotInstancesAPI.SpotActions(context.Background(), spotInstanceId).SpotActionsRequest(spotActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotInstancesAPI.SpotActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotActions`: Vm
	fmt.Fprintf(os.Stdout, "Response from `SpotInstancesAPI.SpotActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotInstanceId** | **int32** | ID of the spot instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpotActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spotActionsRequest** | [**SpotActionsRequest**](SpotActionsRequest.md) |  | 

### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreate

> Vm SpotCreate(ctx).SpotCreate(spotCreate).Execute()

Create spot instance



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
	spotCreate := *openapiclient.NewSpotCreate("spot-default", "aws-us-west-1", int32(35), "multi-cloud", "shared", int32(4), int32(1), "ssd", int32(16), int32(124), float32(0.002635)) // SpotCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotInstancesAPI.SpotCreate(context.Background()).SpotCreate(spotCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotInstancesAPI.SpotCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreate`: Vm
	fmt.Fprintf(os.Stdout, "Response from `SpotInstancesAPI.SpotCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spotCreate** | [**SpotCreate**](SpotCreate.md) |  | 

### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotDelete

> Vm SpotDelete(ctx, spotInstanceId).Execute()

Delete spot instance



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
	spotInstanceId := int32(56) // int32 | ID of the spot instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotInstancesAPI.SpotDelete(context.Background(), spotInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotInstancesAPI.SpotDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotDelete`: Vm
	fmt.Fprintf(os.Stdout, "Response from `SpotInstancesAPI.SpotDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotInstanceId** | **int32** | ID of the spot instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpotDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

