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
	providerId := int32(5) // int32 | ID of the cloud provider (optional)
	locationId := int32(6) // int32 | ID of the geographic location (optional)
	dataCenterId := "aws-us-west-1" // string | ID of the cloud provider's data center (optional)
	cloudNetworkType := "multi-cloud" // string | Type of cloud network (optional)
	vCpuType := "Standard" // string | Type of vCPUs for the compute instance (optional)
	vCpu := int32(4) // int32 | virtual Central Processing Units (vCPUs) for the compute instance (optional)
	vCpuMin := int32(1) // int32 | Minimum number of vCPUs for the compute instance (optional)
	vCpuMax := int32(8) // int32 | Maximum number of vCPUs for the compute instance (optional)
	ramGb := float32(16) // float32 | RAM of the compute instance in gigabytes (optional)
	ramGbMin := float32(8) // float32 | Minimum RAM of the compute instance in gigabytes (optional)
	ramGbMax := float32(32) // float32 | Maximum RAM of the compute instance in gigabytes (optional)
	volumeGb := int32(500) // int32 | Volume size of the compute instance in gigabytes (optional)
	volumeGbMin := int32(250) // int32 | Minimum volume size of the compute instance in gigabytes (optional)
	volumeGbMax := int32(1000) // int32 | Maximum volume size of the compute instance in gigabytes (optional)
	volumeType := "ssd" // string | Volume type of the compute instance (optional)
	priceMin := float32(8.14) // float32 | Minimum price of the compute instance (optional)
	priceMax := float32(8.14) // float32 | Maximum price of the compute instance (optional)
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
 **providerId** | **int32** | ID of the cloud provider | 
 **locationId** | **int32** | ID of the geographic location | 
 **dataCenterId** | **string** | ID of the cloud provider&#39;s data center | 
 **cloudNetworkType** | **string** | Type of cloud network | 
 **vCpuType** | **string** | Type of vCPUs for the compute instance | 
 **vCpu** | **int32** | virtual Central Processing Units (vCPUs) for the compute instance | 
 **vCpuMin** | **int32** | Minimum number of vCPUs for the compute instance | 
 **vCpuMax** | **int32** | Maximum number of vCPUs for the compute instance | 
 **ramGb** | **float32** | RAM of the compute instance in gigabytes | 
 **ramGbMin** | **float32** | Minimum RAM of the compute instance in gigabytes | 
 **ramGbMax** | **float32** | Maximum RAM of the compute instance in gigabytes | 
 **volumeGb** | **int32** | Volume size of the compute instance in gigabytes | 
 **volumeGbMin** | **int32** | Minimum volume size of the compute instance in gigabytes | 
 **volumeGbMax** | **int32** | Maximum volume size of the compute instance in gigabytes | 
 **volumeType** | **string** | Volume type of the compute instance | 
 **priceMin** | **float32** | Minimum price of the compute instance | 
 **priceMax** | **float32** | Maximum price of the compute instance | 
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
	providerId := int32(5) // int32 | ID of the cloud provider (optional)
	locationId := int32(6) // int32 | ID of the geographic location (optional)
	dataCenterId := "aws-us-west-1" // string | ID of the cloud provider's data center (optional)
	cloudNetworkType := "multi-cloud" // string | Type of cloud network (optional)
	vCpuType := "Standard" // string | Type of vCPUs for the compute instance (optional)
	vCpu := int32(4) // int32 | virtual Central Processing Units (vCPUs) for the compute instance (optional)
	vCpuMin := int32(1) // int32 | Minimum number of vCPUs for the compute instance (optional)
	vCpuMax := int32(8) // int32 | Maximum number of vCPUs for the compute instance (optional)
	ramGb := float32(16) // float32 | RAM of the compute instance in gigabytes (optional)
	ramGbMin := float32(8) // float32 | Minimum RAM of the compute instance in gigabytes (optional)
	ramGbMax := float32(32) // float32 | Maximum RAM of the compute instance in gigabytes (optional)
	volumeGb := int32(500) // int32 | Volume size of the compute instance in gigabytes (optional)
	volumeGbMin := int32(250) // int32 | Minimum volume size of the compute instance in gigabytes (optional)
	volumeGbMax := int32(1000) // int32 | Maximum volume size of the compute instance in gigabytes (optional)
	volumeType := "ssd" // string | Volume type of the compute instance (optional)
	priceMin := float32(8.14) // float32 | Minimum price of the compute instance (optional)
	priceMax := float32(8.14) // float32 | Maximum price of the compute instance (optional)
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
 **providerId** | **int32** | ID of the cloud provider | 
 **locationId** | **int32** | ID of the geographic location | 
 **dataCenterId** | **string** | ID of the cloud provider&#39;s data center | 
 **cloudNetworkType** | **string** | Type of cloud network | 
 **vCpuType** | **string** | Type of vCPUs for the compute instance | 
 **vCpu** | **int32** | virtual Central Processing Units (vCPUs) for the compute instance | 
 **vCpuMin** | **int32** | Minimum number of vCPUs for the compute instance | 
 **vCpuMax** | **int32** | Maximum number of vCPUs for the compute instance | 
 **ramGb** | **float32** | RAM of the compute instance in gigabytes | 
 **ramGbMin** | **float32** | Minimum RAM of the compute instance in gigabytes | 
 **ramGbMax** | **float32** | Maximum RAM of the compute instance in gigabytes | 
 **volumeGb** | **int32** | Volume size of the compute instance in gigabytes | 
 **volumeGbMin** | **int32** | Minimum volume size of the compute instance in gigabytes | 
 **volumeGbMax** | **int32** | Maximum volume size of the compute instance in gigabytes | 
 **volumeType** | **string** | Volume type of the compute instance | 
 **priceMin** | **float32** | Minimum price of the compute instance | 
 **priceMax** | **float32** | Maximum price of the compute instance | 
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

