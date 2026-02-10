# \AcceleratorTypesAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAcceleratorTypes**](AcceleratorTypesAPI.md#GetAcceleratorTypes) | **Get** /v1/accelerator-types | Get list of GPU accelerator types



## GetAcceleratorTypes

> []AcceleratorType GetAcceleratorTypes(ctx).Execute()

Get list of GPU accelerator types



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
	resp, r, err := apiClient.AcceleratorTypesAPI.GetAcceleratorTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorTypesAPI.GetAcceleratorTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAcceleratorTypes`: []AcceleratorType
	fmt.Fprintf(os.Stdout, "Response from `AcceleratorTypesAPI.GetAcceleratorTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcceleratorTypesRequest struct via the builder pattern


### Return type

[**[]AcceleratorType**](AcceleratorType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

