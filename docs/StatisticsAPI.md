# \StatisticsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatisticalData**](StatisticsAPI.md#GetStatisticalData) | **Post** /v1/statistics/retrieve | Extract data from the DWH



## GetStatisticalData

> GetStatisticalData200Response GetStatisticalData(ctx).GetStatisticalDataRequest(getStatisticalDataRequest).Execute()

Extract data from the DWH



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
	getStatisticalDataRequest := openapiclient.GetStatisticalData_request{KubernetesClusterChangingMetricsQuery: openapiclient.NewKubernetesClusterChangingMetricsQuery("k8s_dashboard_changing_metrics", int32(11211), *openapiclient.NewKubernetesClusterChangingMetricsQueryFilters("cluster_id", "10650", "node", []string{"cluster_pods_count_pending"}, "60M", "true", "CustomFilterAvgCpuRule_example", float32(50), "less than", float32(50), "greater than", float32(50), []string{"node_1"}))} // GetStatisticalDataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetStatisticalData(context.Background()).GetStatisticalDataRequest(getStatisticalDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetStatisticalData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticalData`: GetStatisticalData200Response
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetStatisticalData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticalDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStatisticalDataRequest** | [**GetStatisticalDataRequest**](GetStatisticalDataRequest.md) |  | 

### Return type

[**GetStatisticalData200Response**](GetStatisticalData200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

