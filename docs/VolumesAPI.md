# \VolumesAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolume**](VolumesAPI.md#GetVolume) | **Get** /v1/volumes/{volumeId} | Get volume by id
[**GetVolumes**](VolumesAPI.md#GetVolumes) | **Get** /v1/volumes | Get list of volumes
[**VolumeActions**](VolumesAPI.md#VolumeActions) | **Post** /v1/volumes/{volumeId}/actions | Perform actions with a volume
[**VolumeCreate**](VolumesAPI.md#VolumeCreate) | **Post** /v1/volumes | Create volume
[**VolumeDelete**](VolumesAPI.md#VolumeDelete) | **Delete** /v1/volumes/{volumeId} | Delete volume



## GetVolume

> Volume GetVolume(ctx, volumeId).Execute()

Get volume by id



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.GetVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumes

> []Volume GetVolumes(ctx).Execute()

Get list of volumes



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
	resp, r, err := apiClient.VolumesAPI.GetVolumes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumes`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolumes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumesRequest struct via the builder pattern


### Return type

[**[]Volume**](Volume.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeActions

> Volume VolumeActions(ctx, volumeId).VolumeEdit(volumeEdit).Execute()

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
	volumeEdit := *openapiclient.NewVolumeEdit("edit", int32(250)) // VolumeEdit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeActions(context.Background(), volumeId).VolumeEdit(volumeEdit).Execute()
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

 **volumeEdit** | [**VolumeEdit**](VolumeEdit.md) |  | 

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


## VolumeCreate

> Volume VolumeCreate(ctx).VolumeCreate(volumeCreate).Execute()

Create volume



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
	volumeCreate := *openapiclient.NewVolumeCreate("aws-us-west-1", int32(500), "SSD") // VolumeCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeCreate(context.Background()).VolumeCreate(volumeCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeCreate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeCreate** | [**VolumeCreate**](VolumeCreate.md) |  | 

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


## VolumeDelete

> Volume VolumeDelete(ctx, volumeId).Execute()

Delete volume



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeDelete(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeDelete`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

