# \VolumesAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeActions**](VolumesAPI.md#VolumeActions) | **Post** /v1/volumes/{volumeId}/actions | Perform actions with a volume



## VolumeActions

> Volume VolumeActions(ctx, volumeId).VolumeActionsRequest(volumeActionsRequest).Execute()

Perform actions with a volume



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
	volumeId := int32(56) // int32 | ID of the volume
	volumeActionsRequest := openapiclient.VolumeActions_request{VolumeEdit: openapiclient.NewVolumeEdit("edit", int32(250))} // VolumeActionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeActions(context.Background(), volumeId).VolumeActionsRequest(volumeActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeActions`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeActionsRequest** | [**VolumeActionsRequest**](VolumeActionsRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

