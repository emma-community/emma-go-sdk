# \VolumesConfigurationsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemVolumeConfigs**](VolumesConfigurationsAPI.md#GetSystemVolumeConfigs) | **Get** /v1/system-volumes-configs | List of available configurations for system volumes



## GetSystemVolumeConfigs

> GetSystemVolumeConfigs200Response GetSystemVolumeConfigs(ctx).AttachedToId(attachedToId).VolumeGbMin(volumeGbMin).VolumeType(volumeType).Page(page).Size(size).Execute()

List of available configurations for system volumes



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
	attachedToId := int32(56) // int32 | ID of the linked instance (optional)
	volumeGbMin := int32(250) // int32 | Minimum volume size of the compute instance in gigabytes (optional)
	volumeType := "ssd" // string | Volume type of the compute instance (optional)
	page := int32(0) // int32 | Page number (optional)
	size := int32(100) // int32 | Query size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesConfigurationsAPI.GetSystemVolumeConfigs(context.Background()).AttachedToId(attachedToId).VolumeGbMin(volumeGbMin).VolumeType(volumeType).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesConfigurationsAPI.GetSystemVolumeConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemVolumeConfigs`: GetSystemVolumeConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumesConfigurationsAPI.GetSystemVolumeConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemVolumeConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachedToId** | **int32** | ID of the linked instance | 
 **volumeGbMin** | **int32** | Minimum volume size of the compute instance in gigabytes | 
 **volumeType** | **string** | Volume type of the compute instance | 
 **page** | **int32** | Page number | 
 **size** | **int32** | Query size | 

### Return type

[**GetSystemVolumeConfigs200Response**](GetSystemVolumeConfigs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

