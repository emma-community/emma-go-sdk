# \SecurityGroupsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityGroup**](SecurityGroupsAPI.md#GetSecurityGroup) | **Get** /v1/security-groups/{securityGroupId} | Get security group by ID
[**GetSecurityGroups**](SecurityGroupsAPI.md#GetSecurityGroups) | **Get** /v1/security-groups | Get list of security groups
[**SecurityGroupCreate**](SecurityGroupsAPI.md#SecurityGroupCreate) | **Post** /v1/security-groups | Create security group
[**SecurityGroupDelete**](SecurityGroupsAPI.md#SecurityGroupDelete) | **Delete** /v1/security-groups/{securityGroupId} | Delete security group
[**SecurityGroupInstanceAdd**](SecurityGroupsAPI.md#SecurityGroupInstanceAdd) | **Post** /v1/security-groups/{securityGroupId}/instances | Add instance to security group
[**SecurityGroupInstances**](SecurityGroupsAPI.md#SecurityGroupInstances) | **Get** /v1/security-groups/{securityGroupId}/instances | Get instances in security group
[**SecurityGroupUpdate**](SecurityGroupsAPI.md#SecurityGroupUpdate) | **Put** /v1/security-groups/{securityGroupId} | Update security group



## GetSecurityGroup

> SecurityGroup GetSecurityGroup(ctx, securityGroupId).Execute()

Get security group by ID

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
	securityGroupId := int32(56) // int32 | Security group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.GetSecurityGroup(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.GetSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **int32** | Security group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroups

> []SecurityGroup GetSecurityGroups(ctx).Execute()

Get list of security groups

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
	resp, r, err := apiClient.SecurityGroupsAPI.GetSecurityGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.GetSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroups`: []SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.GetSecurityGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupsRequest struct via the builder pattern


### Return type

[**[]SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGroupCreate

> SecurityGroup SecurityGroupCreate(ctx).SecurityGroupRequest(securityGroupRequest).Execute()

Create security group

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
	securityGroupRequest := *openapiclient.NewSecurityGroupRequest("new_default", []openapiclient.SecurityGroupRuleRequest{*openapiclient.NewSecurityGroupRuleRequest("INBOUND", "TCP", "1-30321", "0.0.0.0/0")}) // SecurityGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.SecurityGroupCreate(context.Background()).SecurityGroupRequest(securityGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.SecurityGroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityGroupCreate`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.SecurityGroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupRequest** | [**SecurityGroupRequest**](SecurityGroupRequest.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGroupDelete

> SecurityGroup SecurityGroupDelete(ctx, securityGroupId).Execute()

Delete security group

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
	securityGroupId := int32(56) // int32 | Security group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.SecurityGroupDelete(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.SecurityGroupDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityGroupDelete`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.SecurityGroupDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **int32** | Security group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGroupDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGroupInstanceAdd

> Vm SecurityGroupInstanceAdd(ctx, securityGroupId).SecurityGroupInstanceAdd(securityGroupInstanceAdd).Execute()

Add instance to security group

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
	securityGroupId := int32(56) // int32 | Security group id
	securityGroupInstanceAdd := *openapiclient.NewSecurityGroupInstanceAdd() // SecurityGroupInstanceAdd |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.SecurityGroupInstanceAdd(context.Background(), securityGroupId).SecurityGroupInstanceAdd(securityGroupInstanceAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.SecurityGroupInstanceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityGroupInstanceAdd`: Vm
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.SecurityGroupInstanceAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **int32** | Security group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGroupInstanceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityGroupInstanceAdd** | [**SecurityGroupInstanceAdd**](SecurityGroupInstanceAdd.md) |  | 

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


## SecurityGroupInstances

> []Vm SecurityGroupInstances(ctx, securityGroupId).Execute()

Get instances in security group

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
	securityGroupId := int32(56) // int32 | Security group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.SecurityGroupInstances(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.SecurityGroupInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityGroupInstances`: []Vm
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.SecurityGroupInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **int32** | Security group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGroupInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SecurityGroupUpdate

> SecurityGroup SecurityGroupUpdate(ctx, securityGroupId).SecurityGroupRequest(securityGroupRequest).Execute()

Update security group

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
	securityGroupId := int32(56) // int32 | Security group id
	securityGroupRequest := *openapiclient.NewSecurityGroupRequest("new_default", []openapiclient.SecurityGroupRuleRequest{*openapiclient.NewSecurityGroupRuleRequest("INBOUND", "TCP", "1-30321", "0.0.0.0/0")}) // SecurityGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.SecurityGroupUpdate(context.Background(), securityGroupId).SecurityGroupRequest(securityGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.SecurityGroupUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityGroupUpdate`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.SecurityGroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **int32** | Security group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityGroupRequest** | [**SecurityGroupRequest**](SecurityGroupRequest.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

