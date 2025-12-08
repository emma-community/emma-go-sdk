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

> Kubernetes CreateKubernetesCluster(ctx).KubernetesCreate(kubernetesCreate).Execute()

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
	kubernetesCreate := *openapiclient.NewKubernetesCreate("default-k8s-managed-cluster-name", "eu", "K8sConnectionType_example", []openapiclient.KubernetesCreateWorkerNodesInner{*openapiclient.NewKubernetesCreateWorkerNodesInner("default-name", "europe-west1-d", "shared", int32(2), int32(4), "ssd", int32(16))}) // KubernetesCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.CreateKubernetesCluster(context.Background()).KubernetesCreate(kubernetesCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.CreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesCluster`: Kubernetes
	fmt.Fprintf(os.Stdout, "Response from `KubernetesClustersAPI.CreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCreate** | [**KubernetesCreate**](KubernetesCreate.md) |  | 

### Return type

[**Kubernetes**](Kubernetes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubernetesCluster

> Kubernetes DeleteKubernetesCluster(ctx, kubernetesId).Execute()

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
	// response from `DeleteKubernetesCluster`: Kubernetes
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

[**Kubernetes**](Kubernetes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditKubernetesCluster

> Kubernetes EditKubernetesCluster(ctx, kubernetesId).KubernetesUpdate(kubernetesUpdate).Execute()

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
	kubernetesUpdate := *openapiclient.NewKubernetesUpdate([]openapiclient.KubernetesUpdateWorkerNodesInner{*openapiclient.NewKubernetesUpdateWorkerNodesInner("default-name", "europe-west1-d", "shared", int32(2), int32(4), "ssd", int32(16))}) // KubernetesUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesClustersAPI.EditKubernetesCluster(context.Background(), kubernetesId).KubernetesUpdate(kubernetesUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClustersAPI.EditKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditKubernetesCluster`: Kubernetes
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

 **kubernetesUpdate** | [**KubernetesUpdate**](KubernetesUpdate.md) |  | 

### Return type

[**Kubernetes**](Kubernetes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesCluster

> Kubernetes GetKubernetesCluster(ctx, kubernetesId).ProjectId(projectId).Execute()

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
	// response from `GetKubernetesCluster`: Kubernetes
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

[**Kubernetes**](Kubernetes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusters

> []Kubernetes GetKubernetesClusters(ctx).ProjectId(projectId).Execute()

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
	// response from `GetKubernetesClusters`: []Kubernetes
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

[**[]Kubernetes**](Kubernetes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

