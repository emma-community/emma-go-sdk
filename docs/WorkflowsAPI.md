# \WorkflowsAPI

All URIs are relative to *https://api.emma.ms/external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowTemplate**](WorkflowsAPI.md#CreateWorkflowTemplate) | **Post** /v1/workflow-templates | Create a workflow template
[**DeleteWorkflowTemplate**](WorkflowsAPI.md#DeleteWorkflowTemplate) | **Delete** /v1/workflow-templates/{id} | Delete a workflow template
[**GetWorkflowTemplate**](WorkflowsAPI.md#GetWorkflowTemplate) | **Get** /v1/workflow-templates/{id} | Get a workflow template by ID
[**GetWorkflowTemplates**](WorkflowsAPI.md#GetWorkflowTemplates) | **Get** /v1/workflow-templates | List of available workflow templates
[**UpdateWorkflowTemplate**](WorkflowsAPI.md#UpdateWorkflowTemplate) | **Put** /v1/workflow-templates/{id} | Update a workflow template



## CreateWorkflowTemplate

> WorkflowTemplate CreateWorkflowTemplate(ctx).WorkflowTemplateCreate(workflowTemplateCreate).Execute()

Create a workflow template



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
	workflowTemplateCreate := *openapiclient.NewWorkflowTemplateCreate("Example Workflow Template", "Shell", "echo Hello, World!", "PUBLISHED", "COMPUTE") // WorkflowTemplateCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.CreateWorkflowTemplate(context.Background()).WorkflowTemplateCreate(workflowTemplateCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateWorkflowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowTemplate`: WorkflowTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTemplateCreate** | [**WorkflowTemplateCreate**](WorkflowTemplateCreate.md) |  | 

### Return type

[**WorkflowTemplate**](WorkflowTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowTemplate

> DeleteWorkflowTemplate(ctx, id).Execute()

Delete a workflow template



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
	id := int32(56) // int32 | ID of the workflow template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.DeleteWorkflowTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.DeleteWorkflowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the workflow template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTemplate

> WorkflowTemplate GetWorkflowTemplate(ctx, id).Execute()

Get a workflow template by ID



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
	id := int32(56) // int32 | ID of the workflow template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowTemplate`: WorkflowTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the workflow template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTemplate**](WorkflowTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTemplates

> GetWorkflowTemplates200Response GetWorkflowTemplates(ctx).Page(page).Size(size).NameLike(nameLike).FilterTag(filterTag).FilterTagNe(filterTagNe).Status(status).StatusNe(statusNe).Execute()

List of available workflow templates



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
	page := int32(0) // int32 | Page number
	size := int32(100) // int32 | Query size
	nameLike := "Deploy" // string | Filter workflow templates by name using partial match (optional)
	filterTag := "env:prod" // string | Filter workflow templates by tag key and value. For example, __filter.tag=env:prod will return workflow templates with tag key \"env\" and value \"prod\", __filter.tag=env:* will return workflow templates with tag key \"env\" and any value (optional)
	filterTagNe := "env:prod" // string | Exclude workflow templates by tag key and value. For example, __filter.tag.ne=env:prod will exclude workflow templates with tag key \"env\" and value \"prod\", __filter.tag.ne=env:* will exclude workflow templates with tag key \"env\" and any value (optional)
	status := "PUBLISHED" // string | Filter workflow templates by status (optional)
	statusNe := "UNPUBLISHED" // string | Exclude workflow templates by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowTemplates(context.Background()).Page(page).Size(size).NameLike(nameLike).FilterTag(filterTag).FilterTagNe(filterTagNe).Status(status).StatusNe(statusNe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowTemplates`: GetWorkflowTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | 
 **size** | **int32** | Query size | 
 **nameLike** | **string** | Filter workflow templates by name using partial match | 
 **filterTag** | **string** | Filter workflow templates by tag key and value. For example, __filter.tag&#x3D;env:prod will return workflow templates with tag key \&quot;env\&quot; and value \&quot;prod\&quot;, __filter.tag&#x3D;env:* will return workflow templates with tag key \&quot;env\&quot; and any value | 
 **filterTagNe** | **string** | Exclude workflow templates by tag key and value. For example, __filter.tag.ne&#x3D;env:prod will exclude workflow templates with tag key \&quot;env\&quot; and value \&quot;prod\&quot;, __filter.tag.ne&#x3D;env:* will exclude workflow templates with tag key \&quot;env\&quot; and any value | 
 **status** | **string** | Filter workflow templates by status | 
 **statusNe** | **string** | Exclude workflow templates by status | 

### Return type

[**GetWorkflowTemplates200Response**](GetWorkflowTemplates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTemplate

> WorkflowTemplate UpdateWorkflowTemplate(ctx, id).WorkflowTemplateCreate(workflowTemplateCreate).Execute()

Update a workflow template



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
	id := int32(56) // int32 | ID of the workflow template
	workflowTemplateCreate := *openapiclient.NewWorkflowTemplateCreate("Example Workflow Template", "Shell", "echo Hello, World!", "PUBLISHED", "COMPUTE") // WorkflowTemplateCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.UpdateWorkflowTemplate(context.Background(), id).WorkflowTemplateCreate(workflowTemplateCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.UpdateWorkflowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowTemplate`: WorkflowTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.UpdateWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the workflow template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowTemplateCreate** | [**WorkflowTemplateCreate**](WorkflowTemplateCreate.md) |  | 

### Return type

[**WorkflowTemplate**](WorkflowTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

