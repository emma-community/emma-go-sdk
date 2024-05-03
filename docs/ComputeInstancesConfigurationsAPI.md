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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	providerId := TODO // interface{} | Provider ID (optional)
	locationId := int32(6) // int32 | Location ID (optional)
	dataCenterId := TODO // interface{} | Data center ID (optional)
	cloudNetworkType := TODO // interface{} | Cloud network type (optional)
	vCpuType := TODO // interface{} | Compute instance vCPU type (optional)
	vCpu := TODO // interface{} | Compute instance vCPU (optional)
	vCpuMin := TODO // interface{} | Compute instance vCPU minimum (optional)
	vCpuMax := TODO // interface{} | Compute instance vCPU maximum (optional)
	ramGb := TODO // interface{} | Compute instance RAM (GB) (optional)
	ramGbMin := TODO // interface{} | Compute instance RAM (GB) minimum (optional)
	ramGbMax := TODO // interface{} | Compute instance RAM (GB) maximum (optional)
	volumeGb := TODO // interface{} | Compute instance volume (GB) (optional)
	volumeGbMin := TODO // interface{} | Compute instance volume minimum (GB) (optional)
	volumeGbMax := TODO // interface{} | Compute instance volume maximun (GB) (optional)
	volumeType := TODO // interface{} | Compute instance volume type (optional)
	priceMin := TODO // interface{} | Instance price minimum (optional)
	priceMax := TODO // interface{} | Instance price maximum (optional)
	page := TODO // interface{} | Page number (optional)
	size := TODO // interface{} | Query size (optional)

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
 **providerId** | [**interface{}**](interface{}.md) | Provider ID | 
 **locationId** | **int32** | Location ID | 
 **dataCenterId** | [**interface{}**](interface{}.md) | Data center ID | 
 **cloudNetworkType** | [**interface{}**](interface{}.md) | Cloud network type | 
 **vCpuType** | [**interface{}**](interface{}.md) | Compute instance vCPU type | 
 **vCpu** | [**interface{}**](interface{}.md) | Compute instance vCPU | 
 **vCpuMin** | [**interface{}**](interface{}.md) | Compute instance vCPU minimum | 
 **vCpuMax** | [**interface{}**](interface{}.md) | Compute instance vCPU maximum | 
 **ramGb** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) | 
 **ramGbMin** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) minimum | 
 **ramGbMax** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) maximum | 
 **volumeGb** | [**interface{}**](interface{}.md) | Compute instance volume (GB) | 
 **volumeGbMin** | [**interface{}**](interface{}.md) | Compute instance volume minimum (GB) | 
 **volumeGbMax** | [**interface{}**](interface{}.md) | Compute instance volume maximun (GB) | 
 **volumeType** | [**interface{}**](interface{}.md) | Compute instance volume type | 
 **priceMin** | [**interface{}**](interface{}.md) | Instance price minimum | 
 **priceMax** | [**interface{}**](interface{}.md) | Instance price maximum | 
 **page** | [**interface{}**](interface{}.md) | Page number | 
 **size** | [**interface{}**](interface{}.md) | Query size | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	providerId := TODO // interface{} | Provider ID (optional)
	locationId := int32(6) // int32 | Location ID (optional)
	dataCenterId := TODO // interface{} | Data center ID (optional)
	cloudNetworkType := TODO // interface{} | Cloud network type (optional)
	vCpuType := TODO // interface{} | Compute instance vCPU type (optional)
	vCpu := TODO // interface{} | Compute instance vCPU (optional)
	vCpuMin := TODO // interface{} | Compute instance vCPU minimum (optional)
	vCpuMax := TODO // interface{} | Compute instance vCPU maximum (optional)
	ramGb := TODO // interface{} | Compute instance RAM (GB) (optional)
	ramGbMin := TODO // interface{} | Compute instance RAM (GB) minimum (optional)
	ramGbMax := TODO // interface{} | Compute instance RAM (GB) maximum (optional)
	volumeGb := TODO // interface{} | Compute instance volume (GB) (optional)
	volumeGbMin := TODO // interface{} | Compute instance volume minimum (GB) (optional)
	volumeGbMax := TODO // interface{} | Compute instance volume maximun (GB) (optional)
	volumeType := TODO // interface{} | Compute instance volume type (optional)
	priceMin := TODO // interface{} | Instance price minimum (optional)
	priceMax := TODO // interface{} | Instance price maximum (optional)
	page := TODO // interface{} | Page number (optional)
	size := TODO // interface{} | Query size (optional)

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
 **providerId** | [**interface{}**](interface{}.md) | Provider ID | 
 **locationId** | **int32** | Location ID | 
 **dataCenterId** | [**interface{}**](interface{}.md) | Data center ID | 
 **cloudNetworkType** | [**interface{}**](interface{}.md) | Cloud network type | 
 **vCpuType** | [**interface{}**](interface{}.md) | Compute instance vCPU type | 
 **vCpu** | [**interface{}**](interface{}.md) | Compute instance vCPU | 
 **vCpuMin** | [**interface{}**](interface{}.md) | Compute instance vCPU minimum | 
 **vCpuMax** | [**interface{}**](interface{}.md) | Compute instance vCPU maximum | 
 **ramGb** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) | 
 **ramGbMin** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) minimum | 
 **ramGbMax** | [**interface{}**](interface{}.md) | Compute instance RAM (GB) maximum | 
 **volumeGb** | [**interface{}**](interface{}.md) | Compute instance volume (GB) | 
 **volumeGbMin** | [**interface{}**](interface{}.md) | Compute instance volume minimum (GB) | 
 **volumeGbMax** | [**interface{}**](interface{}.md) | Compute instance volume maximun (GB) | 
 **volumeType** | [**interface{}**](interface{}.md) | Compute instance volume type | 
 **priceMin** | [**interface{}**](interface{}.md) | Instance price minimum | 
 **priceMax** | [**interface{}**](interface{}.md) | Instance price maximum | 
 **page** | [**interface{}**](interface{}.md) | Page number | 
 **size** | [**interface{}**](interface{}.md) | Query size | 

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

