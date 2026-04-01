# \KubernetesClustersAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKubernetesCluster**](KubernetesClustersAPI.md#CreateKubernetesCluster) | **Post** /v1/kubernetes | Create Kubernetes cluster
[**DeleteKubernetesCluster**](KubernetesClustersAPI.md#DeleteKubernetesCluster) | **Delete** /v1/kubernetes/{kubernetesId} | Delete Kubernetes cluster
[**EditKubernetesCluster**](KubernetesClustersAPI.md#EditKubernetesCluster) | **Put** /v1/kubernetes/{kubernetesId} | Edit Kubernetes cluster
[**GetKubernetesCluster**](KubernetesClustersAPI.md#GetKubernetesCluster) | **Get** /v1/kubernetes/{kubernetesId} | Get Kubernetes cluster by id
[**GetKubernetesClusters**](KubernetesClustersAPI.md#GetKubernetesClusters) | **Get** /v1/kubernetes | Get list of Kubernetes clusters



## CreateKubernetesCluster

> KubernetesCreateResponse CreateKubernetesCluster(ctx).KubernetesCreateRequest(kubernetesCreateRequest).Execute()

Create Kubernetes cluster



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
	kubernetesCreateRequest := *openapiclient.NewKubernetesCreateRequest("eu", "default-k8s-managed-cluster-name", "internet_connect", []openapiclient.KubernetesCreateRequestWorkerNodesInner{*openapiclient.NewKubernetesCreateRequestWorkerNodesInner("vm-automation-co2ep", "gcp-europe-west4-a", "standard", int32(4), int32(16), int32(32), "ssd")}) // KubernetesCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.CreateKubernetesCluster(context.Background()).KubernetesCreateRequest(kubernetesCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.CreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesCluster`: KubernetesCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.CreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCreateRequest** | [**KubernetesCreateRequest**](KubernetesCreateRequest.md) |  | 

### Return type

[**KubernetesCreateResponse**](KubernetesCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubernetesCluster

> KubernetesDelete DeleteKubernetesCluster(ctx, kubernetesId).Execute()

Delete Kubernetes cluster



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
	kubernetesId := int32(56) // int32 | ID of the Kubernetes cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.DeleteKubernetesCluster(context.Background(), kubernetesId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.DeleteKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKubernetesCluster`: KubernetesDelete
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.DeleteKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesId** | **int32** | ID of the Kubernetes cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesDelete**](KubernetesDelete.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditKubernetesCluster

> KubernetesUpdateResponse EditKubernetesCluster(ctx, kubernetesId).KubernetesUpdateRequest(kubernetesUpdateRequest).Execute()

Edit Kubernetes cluster



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
	kubernetesId := int32(56) // int32 | ID of the Kubernetes cluster
	kubernetesUpdateRequest := *openapiclient.NewKubernetesUpdateRequest([]openapiclient.KubernetesUpdateRequestWorkerNodesInner{*openapiclient.NewKubernetesUpdateRequestWorkerNodesInner("lsydyf-node", "gcp-europe-west4-a", "standard", int32(4), int32(16), "ssd", int32(32))}) // KubernetesUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.EditKubernetesCluster(context.Background(), kubernetesId).KubernetesUpdateRequest(kubernetesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.EditKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditKubernetesCluster`: KubernetesUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.EditKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesId** | **int32** | ID of the Kubernetes cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesUpdateRequest** | [**KubernetesUpdateRequest**](KubernetesUpdateRequest.md) |  | 

### Return type

[**KubernetesUpdateResponse**](KubernetesUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesCluster

> KubernetesGetResponse GetKubernetesCluster(ctx, kubernetesId).ProjectId(projectId).Execute()

Get Kubernetes cluster by id



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
	kubernetesId := int32(56) // int32 | ID of the Kubernetes cluster
	projectId := int32(56) // int32 | Unused, created for future API extention. Will be ignored if used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.GetKubernetesCluster(context.Background(), kubernetesId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.GetKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesCluster`: KubernetesGetResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.GetKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesId** | **int32** | ID of the Kubernetes cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**KubernetesGetResponse**](KubernetesGetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusters

> []KubernetesListResponseInner GetKubernetesClusters(ctx).ProjectId(projectId).Execute()

Get list of Kubernetes clusters



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
	resp, r, err := apiClient.KubernetesClustersAPI.GetKubernetesClusters(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.GetKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesClusters`: []KubernetesListResponseInner
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.GetKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** | Unused, created for future API extention. Will be ignored if used. | 

### Return type

[**[]KubernetesListResponseInner**](KubernetesListResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

