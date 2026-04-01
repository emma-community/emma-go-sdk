# \NetworkAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultiCloudNetworks**](NetworkAPI.md#GetMultiCloudNetworks) | **Get** /v1/multi-cloud-networks | Get the Direct Connect multi-cloud network configuration across the EMEA, AMER, and APAC macro regions.
[**GetProjectConnectivityCenterById**](NetworkAPI.md#GetProjectConnectivityCenterById) | **Get** /v1/connectivity-centers/{id} | Get connectivity center by ID
[**ListConnectivityCenters**](NetworkAPI.md#ListConnectivityCenters) | **Get** /v1/connectivity-centers | List connectivity centers in direct connect network
[**PostMultiCloudNetworkAction**](NetworkAPI.md#PostMultiCloudNetworkAction) | **Post** /v1/multi-cloud-networks/actions | Perform an action on multi-cloud network resources (enable/disable).



## GetMultiCloudNetworks

> MultiCloudNetwork GetMultiCloudNetworks(ctx).ProjectId(projectId).Execute()

Get the Direct Connect multi-cloud network configuration across the EMEA, AMER, and APAC macro regions.



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
	resp, r, err := apiClient.NetworkAPI.GetMultiCloudNetworks(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetMultiCloudNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiCloudNetworks`: MultiCloudNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetMultiCloudNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiCloudNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**MultiCloudNetwork**](MultiCloudNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectConnectivityCenterById

> ConnectivityCenter GetProjectConnectivityCenterById(ctx, id).Execute()

Get connectivity center by ID



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
	id := int32(1) // int32 | Connectivity center ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetProjectConnectivityCenterById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetProjectConnectivityCenterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectConnectivityCenterById`: ConnectivityCenter
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetProjectConnectivityCenterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Connectivity center ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectConnectivityCenterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectivityCenter**](ConnectivityCenter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectivityCenters

> []ConnectivityCenter ListConnectivityCenters(ctx).Location(location).LocationNe(locationNe).LocationLike(locationLike).LocationNlike(locationNlike).MacroRegion(macroRegion).NetworkType(networkType).NetworkTypeNe(networkTypeNe).Execute()

List connectivity centers in direct connect network



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
	location := "Frankfurt" // string | Filter by location using equality. Can be repeated to apply OR logic. Example: `?location=Frankfurt&location=Paris`  (optional)
	locationNe := "London" // string | Filter by location using not equal operator. (optional)
	locationLike := "Frank" // string | Filter by location using contains operator. (optional)
	locationNlike := "don" // string | Filter by location using does not contain operator. (optional)
	macroRegion := "emea" // string | Filter by macro region using equality. Can be repeated to apply OR logic.  Example: `?macroRegion=emea&macroRegion=amer`  (optional)
	networkType := "direct-connect" // string | Filter by network type using equality. (optional)
	networkTypeNe := "direct-connect" // string | Filter by network type using not equal operator. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.ListConnectivityCenters(context.Background()).Location(location).LocationNe(locationNe).LocationLike(locationLike).LocationNlike(locationNlike).MacroRegion(macroRegion).NetworkType(networkType).NetworkTypeNe(networkTypeNe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.ListConnectivityCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectivityCenters`: []ConnectivityCenter
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.ListConnectivityCenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectivityCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | Filter by location using equality. Can be repeated to apply OR logic. Example: &#x60;?location&#x3D;Frankfurt&amp;location&#x3D;Paris&#x60;  | 
 **locationNe** | **string** | Filter by location using not equal operator. | 
 **locationLike** | **string** | Filter by location using contains operator. | 
 **locationNlike** | **string** | Filter by location using does not contain operator. | 
 **macroRegion** | **string** | Filter by macro region using equality. Can be repeated to apply OR logic.  Example: &#x60;?macroRegion&#x3D;emea&amp;macroRegion&#x3D;amer&#x60;  | 
 **networkType** | **string** | Filter by network type using equality. | 
 **networkTypeNe** | **string** | Filter by network type using not equal operator. | 

### Return type

[**[]ConnectivityCenter**](ConnectivityCenter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMultiCloudNetworkAction

> []MultiCloudNetwork PostMultiCloudNetworkAction(ctx).PostMultiCloudNetworkActionRequest(postMultiCloudNetworkActionRequest).ProjectId(projectId).Execute()

Perform an action on multi-cloud network resources (enable/disable).



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
	postMultiCloudNetworkActionRequest := openapiclient.PostMultiCloudNetworkAction_request{McnDirectCloudConnectAction: openapiclient.NewMcnDirectCloudConnectAction("direct-connect", "emea", "aws", "enable-cloud-connect")} // PostMultiCloudNetworkActionRequest | 
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.PostMultiCloudNetworkAction(context.Background()).PostMultiCloudNetworkActionRequest(postMultiCloudNetworkActionRequest).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.PostMultiCloudNetworkAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultiCloudNetworkAction`: []MultiCloudNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.PostMultiCloudNetworkAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultiCloudNetworkActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postMultiCloudNetworkActionRequest** | [**PostMultiCloudNetworkActionRequest**](PostMultiCloudNetworkActionRequest.md) |  | 
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]MultiCloudNetwork**](MultiCloudNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

