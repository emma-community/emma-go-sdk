# WorkflowTemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the workflow template | 
**Description** | Pointer to **string** | Description of the workflow template | [optional] 
**ContentType** | **string** | Format of the content field, only Shell is supported for now | 
**Content** | **string** | Content of the workflow template. For Shell contentType, it includes shell commands to execute. | 
**Status** | **string** | Status of the workflow template. Can be PUBLISHED or UNPUBLISHED | 
**ResourceType** | **string** | Type of the resource, in this case COMPUTE | 
**ResourceParams** | Pointer to [**[]WorkflowTemplateCreateResourceParamsInner**](WorkflowTemplateCreateResourceParamsInner.md) | Parameters of the resource limitation. The parameters depend on the resource type and include parameters for compute instance. | [optional] 
**Tags** | Pointer to [**[]WorkflowTemplateCreateTagsInner**](WorkflowTemplateCreateTagsInner.md) | List of tags assigned to the workflow template | [optional] 
**ContentParams** | Pointer to [**[]WorkflowTemplateCreateContentParamsInner**](WorkflowTemplateCreateContentParamsInner.md) | List of parameters extracted from the content. These parameters can be used to fill in the blanks in the content when creating a workflow from the template. | [optional] 

## Methods

### NewWorkflowTemplateCreate

`func NewWorkflowTemplateCreate(name string, contentType string, content string, status string, resourceType string, ) *WorkflowTemplateCreate`

NewWorkflowTemplateCreate instantiates a new WorkflowTemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateCreateWithDefaults

`func NewWorkflowTemplateCreateWithDefaults() *WorkflowTemplateCreate`

NewWorkflowTemplateCreateWithDefaults instantiates a new WorkflowTemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowTemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowTemplateCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTemplateCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTemplateCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTemplateCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentType

`func (o *WorkflowTemplateCreate) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkflowTemplateCreate) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkflowTemplateCreate) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContent

`func (o *WorkflowTemplateCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkflowTemplateCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WorkflowTemplateCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetStatus

`func (o *WorkflowTemplateCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTemplateCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTemplateCreate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResourceType

`func (o *WorkflowTemplateCreate) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WorkflowTemplateCreate) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WorkflowTemplateCreate) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceParams

`func (o *WorkflowTemplateCreate) GetResourceParams() []WorkflowTemplateCreateResourceParamsInner`

GetResourceParams returns the ResourceParams field if non-nil, zero value otherwise.

### GetResourceParamsOk

`func (o *WorkflowTemplateCreate) GetResourceParamsOk() (*[]WorkflowTemplateCreateResourceParamsInner, bool)`

GetResourceParamsOk returns a tuple with the ResourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceParams

`func (o *WorkflowTemplateCreate) SetResourceParams(v []WorkflowTemplateCreateResourceParamsInner)`

SetResourceParams sets ResourceParams field to given value.

### HasResourceParams

`func (o *WorkflowTemplateCreate) HasResourceParams() bool`

HasResourceParams returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowTemplateCreate) GetTags() []WorkflowTemplateCreateTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowTemplateCreate) GetTagsOk() (*[]WorkflowTemplateCreateTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowTemplateCreate) SetTags(v []WorkflowTemplateCreateTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowTemplateCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContentParams

`func (o *WorkflowTemplateCreate) GetContentParams() []WorkflowTemplateCreateContentParamsInner`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *WorkflowTemplateCreate) GetContentParamsOk() (*[]WorkflowTemplateCreateContentParamsInner, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *WorkflowTemplateCreate) SetContentParams(v []WorkflowTemplateCreateContentParamsInner)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *WorkflowTemplateCreate) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


