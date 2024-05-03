# \VirtualMachinesAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVm**](VirtualMachinesAPI.md#GetVm) | **Get** /v1/vms/{vmId} | Get virtual machine by id
[**GetVms**](VirtualMachinesAPI.md#GetVms) | **Get** /v1/vms | Get list of virtual machines
[**VmActions**](VirtualMachinesAPI.md#VmActions) | **Post** /v1/vms/{vmId}/actions | Perform actions with a virtual machine
[**VmCreate**](VirtualMachinesAPI.md#VmCreate) | **Post** /v1/vms | Create virtual machine
[**VmDelete**](VirtualMachinesAPI.md#VmDelete) | **Delete** /v1/vms/{vmId} | Delete virtual machine



## GetVm

> Vm GetVm(ctx, vmId).Execute()

Get virtual machine by id

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
	vmId := int32(56) // int32 | Virtual machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVm(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVm`: Vm
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** | Virtual machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVms

> []Vm GetVms(ctx).Execute()

Get list of virtual machines

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVms`: []Vm
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmsRequest struct via the builder pattern


### Return type

[**[]Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmActions

> interface{} VmActions(ctx, vmId).VmActionsRequest(vmActionsRequest).Execute()

Perform actions with a virtual machine

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
	vmId := int32(56) // int32 | Virtual machine ID
	vmActionsRequest := openapiclient.VmActions_request{VmClone: openapiclient.NewVmClone(interface{}(clone), interface{}(vm-default-name))} // VmActionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.VmActions(context.Background(), vmId).VmActionsRequest(vmActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.VmActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmActions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.VmActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** | Virtual machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmActionsRequest** | [**VmActionsRequest**](VmActionsRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmCreate

> Vm VmCreate(ctx).VmCreate(vmCreate).Execute()

Create virtual machine

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
	vmCreate := *openapiclient.NewVmCreate(interface{}(vm-test1), interface{}(aws-us-west-1), interface{}(5), interface{}(multi-cloud), interface{}(shared), interface{}(2), interface{}(1), interface{}(ssd), interface{}(16), interface{}(124)) // VmCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.VmCreate(context.Background()).VmCreate(vmCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.VmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmCreate`: Vm
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.VmCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmCreate** | [**VmCreate**](VmCreate.md) |  | 

### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmDelete

> Vm VmDelete(ctx, vmId).Execute()

Delete virtual machine

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
	vmId := int32(56) // int32 | Virtual machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.VmDelete(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.VmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmDelete`: Vm
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.VmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** | Virtual machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vm**](Vm.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

