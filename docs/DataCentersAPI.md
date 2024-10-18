# \DataCentersAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataCenter**](DataCentersAPI.md#GetDataCenter) | **Get** /v1/data-centers/{dataCenterId} | Get data center by ID
[**GetDataCenters**](DataCentersAPI.md#GetDataCenters) | **Get** /v1/data-centers | Get list of data centers



## GetDataCenter

> DataCenter GetDataCenter(ctx, dataCenterId).Execute()

Get data center by ID



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
	dataCenterId := "aws-us-west-1" // string | ID of the cloud provider's data center

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataCentersAPI.GetDataCenter(context.Background(), dataCenterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCentersAPI.GetDataCenter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataCenter`: DataCenter
	fmt.Fprintf(os.Stdout, "Response from `DataCentersAPI.GetDataCenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataCenterId** | **string** | ID of the cloud provider&#39;s data center | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataCenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataCenter**](DataCenter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataCenters

> []DataCenter GetDataCenters(ctx).DataCenterName(dataCenterName).LocationId(locationId).ProviderName(providerName).Execute()

Get list of data centers



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
	dataCenterName := "eu-north-1" // string | Name of the cloud provider's data center (optional)
	locationId := int32(6) // int32 | ID of the geographic location (optional)
	providerName := "GCP" // string | Name of the cloud provider (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataCentersAPI.GetDataCenters(context.Background()).DataCenterName(dataCenterName).LocationId(locationId).ProviderName(providerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCentersAPI.GetDataCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataCenters`: []DataCenter
	fmt.Fprintf(os.Stdout, "Response from `DataCentersAPI.GetDataCenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataCenterName** | **string** | Name of the cloud provider&#39;s data center | 
 **locationId** | **int32** | ID of the geographic location | 
 **providerName** | **string** | Name of the cloud provider | 

### Return type

[**[]DataCenter**](DataCenter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

