# \ComputeInstancesConfigurationsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpotConfigs**](ComputeInstancesConfigurationsAPI.md#GetSpotConfigs) | **Get** /v1/spots-configs | Search available configurations for spot instance creation
[**GetVmConfigs**](ComputeInstancesConfigurationsAPI.md#GetVmConfigs) | **Get** /v1/vms-configs | Search available configurations for virtual machine creation



## GetSpotConfigs

> GetVmConfigs200Response GetSpotConfigs(ctx).ProviderId(providerId).LocationId(locationId).DataCenterId(dataCenterId).CloudNetworkType(cloudNetworkType).VCpuType(vCpuType).VCpu(vCpu).VCpuMin(vCpuMin).VCpuMax(vCpuMax).RamGb(ramGb).RamGbMin(ramGbMin).RamGbMax(ramGbMax).VolumeGb(volumeGb).VolumeGbMin(volumeGbMin).VolumeGbMax(volumeGbMax).VolumeType(volumeType).PriceMin(priceMin).PriceMax(priceMax).Page(page).Size(size).Execute()

Search available configurations for spot instance creation

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
	providerId := int32(5) // int32 | Provider ID (optional)
	locationId := int32(6) // int32 | Location ID (optional)
	dataCenterId := "aws-us-west-1" // string | Data center ID (optional)
	cloudNetworkType := "multi-cloud" // string | Cloud network type (optional)
	vCpuType := "Standard" // string | Compute instance vCPU type (optional)
	vCpu := int32(4) // int32 | Compute instance vCPU (optional)
	vCpuMin := int32(1) // int32 | Compute instance vCPU minimum (optional)
	vCpuMax := int32(8) // int32 | Compute instance vCPU maximum (optional)
	ramGb := int32(16) // int32 | Compute instance RAM (GB) (optional)
	ramGbMin := int32(8) // int32 | Compute instance RAM (GB) minimum (optional)
	ramGbMax := int32(32) // int32 | Compute instance RAM (GB) maximum (optional)
	volumeGb := int32(500) // int32 | Compute instance volume (GB) (optional)
	volumeGbMin := int32(250) // int32 | Compute instance volume minimum (GB) (optional)
	volumeGbMax := int32(1000) // int32 | Compute instance volume maximun (GB) (optional)
	volumeType := "ssd" // string | Compute instance volume type (optional)
	priceMin := float32(8.14) // float32 | Instance price minimum (optional)
	priceMax := float32(8.14) // float32 | Instance price maximum (optional)
	page := int32(0) // int32 | Page number (optional)
	size := int32(100) // int32 | Query size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeInstancesConfigurationsAPI.GetSpotConfigs(context.Background()).ProviderId(providerId).LocationId(locationId).DataCenterId(dataCenterId).CloudNetworkType(cloudNetworkType).VCpuType(vCpuType).VCpu(vCpu).VCpuMin(vCpuMin).VCpuMax(vCpuMax).RamGb(ramGb).RamGbMin(ramGbMin).RamGbMax(ramGbMax).VolumeGb(volumeGb).VolumeGbMin(volumeGbMin).VolumeGbMax(volumeGbMax).VolumeType(volumeType).PriceMin(priceMin).PriceMax(priceMax).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeInstancesConfigurationsAPI.GetSpotConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpotConfigs`: GetVmConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `ComputeInstancesConfigurationsAPI.GetSpotConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpotConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **int32** | Provider ID | 
 **locationId** | **int32** | Location ID | 
 **dataCenterId** | **string** | Data center ID | 
 **cloudNetworkType** | **string** | Cloud network type | 
 **vCpuType** | **string** | Compute instance vCPU type | 
 **vCpu** | **int32** | Compute instance vCPU | 
 **vCpuMin** | **int32** | Compute instance vCPU minimum | 
 **vCpuMax** | **int32** | Compute instance vCPU maximum | 
 **ramGb** | **int32** | Compute instance RAM (GB) | 
 **ramGbMin** | **int32** | Compute instance RAM (GB) minimum | 
 **ramGbMax** | **int32** | Compute instance RAM (GB) maximum | 
 **volumeGb** | **int32** | Compute instance volume (GB) | 
 **volumeGbMin** | **int32** | Compute instance volume minimum (GB) | 
 **volumeGbMax** | **int32** | Compute instance volume maximun (GB) | 
 **volumeType** | **string** | Compute instance volume type | 
 **priceMin** | **float32** | Instance price minimum | 
 **priceMax** | **float32** | Instance price maximum | 
 **page** | **int32** | Page number | 
 **size** | **int32** | Query size | 

### Return type

[**GetVmConfigs200Response**](GetVmConfigs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmConfigs

> GetVmConfigs200Response GetVmConfigs(ctx).ProviderId(providerId).LocationId(locationId).DataCenterId(dataCenterId).CloudNetworkType(cloudNetworkType).VCpuType(vCpuType).VCpu(vCpu).VCpuMin(vCpuMin).VCpuMax(vCpuMax).RamGb(ramGb).RamGbMin(ramGbMin).RamGbMax(ramGbMax).VolumeGb(volumeGb).VolumeGbMin(volumeGbMin).VolumeGbMax(volumeGbMax).VolumeType(volumeType).PriceMin(priceMin).PriceMax(priceMax).Page(page).Size(size).Execute()

Search available configurations for virtual machine creation

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
	providerId := int32(5) // int32 | Provider ID (optional)
	locationId := int32(6) // int32 | Location ID (optional)
	dataCenterId := "aws-us-west-1" // string | Data center ID (optional)
	cloudNetworkType := "multi-cloud" // string | Cloud network type (optional)
	vCpuType := "Standard" // string | Compute instance vCPU type (optional)
	vCpu := int32(4) // int32 | Compute instance vCPU (optional)
	vCpuMin := int32(1) // int32 | Compute instance vCPU minimum (optional)
	vCpuMax := int32(8) // int32 | Compute instance vCPU maximum (optional)
	ramGb := int32(16) // int32 | Compute instance RAM (GB) (optional)
	ramGbMin := int32(8) // int32 | Compute instance RAM (GB) minimum (optional)
	ramGbMax := int32(32) // int32 | Compute instance RAM (GB) maximum (optional)
	volumeGb := int32(500) // int32 | Compute instance volume (GB) (optional)
	volumeGbMin := int32(250) // int32 | Compute instance volume minimum (GB) (optional)
	volumeGbMax := int32(1000) // int32 | Compute instance volume maximun (GB) (optional)
	volumeType := "ssd" // string | Compute instance volume type (optional)
	priceMin := float32(8.14) // float32 | Instance price minimum (optional)
	priceMax := float32(8.14) // float32 | Instance price maximum (optional)
	page := int32(0) // int32 | Page number (optional)
	size := int32(100) // int32 | Query size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeInstancesConfigurationsAPI.GetVmConfigs(context.Background()).ProviderId(providerId).LocationId(locationId).DataCenterId(dataCenterId).CloudNetworkType(cloudNetworkType).VCpuType(vCpuType).VCpu(vCpu).VCpuMin(vCpuMin).VCpuMax(vCpuMax).RamGb(ramGb).RamGbMin(ramGbMin).RamGbMax(ramGbMax).VolumeGb(volumeGb).VolumeGbMin(volumeGbMin).VolumeGbMax(volumeGbMax).VolumeType(volumeType).PriceMin(priceMin).PriceMax(priceMax).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeInstancesConfigurationsAPI.GetVmConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVmConfigs`: GetVmConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `ComputeInstancesConfigurationsAPI.GetVmConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVmConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **int32** | Provider ID | 
 **locationId** | **int32** | Location ID | 
 **dataCenterId** | **string** | Data center ID | 
 **cloudNetworkType** | **string** | Cloud network type | 
 **vCpuType** | **string** | Compute instance vCPU type | 
 **vCpu** | **int32** | Compute instance vCPU | 
 **vCpuMin** | **int32** | Compute instance vCPU minimum | 
 **vCpuMax** | **int32** | Compute instance vCPU maximum | 
 **ramGb** | **int32** | Compute instance RAM (GB) | 
 **ramGbMin** | **int32** | Compute instance RAM (GB) minimum | 
 **ramGbMax** | **int32** | Compute instance RAM (GB) maximum | 
 **volumeGb** | **int32** | Compute instance volume (GB) | 
 **volumeGbMin** | **int32** | Compute instance volume minimum (GB) | 
 **volumeGbMax** | **int32** | Compute instance volume maximun (GB) | 
 **volumeType** | **string** | Compute instance volume type | 
 **priceMin** | **float32** | Instance price minimum | 
 **priceMax** | **float32** | Instance price maximum | 
 **page** | **int32** | Page number | 
 **size** | **int32** | Query size | 

### Return type

[**GetVmConfigs200Response**](GetVmConfigs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

