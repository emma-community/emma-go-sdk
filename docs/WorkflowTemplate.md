# WorkflowTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the workflow template | 
**Name** | **string** | Name of the workflow template | 
**Description** | Pointer to **string** | Description of the workflow template | [optional] 
**CreatedAt** | **string** | Date and time of the workflow template creation | 
**CreatedByName** | **string** | Name of the user who created the workflow template | 
**CreatedById** | **int32** | ID of the user who created the workflow template | 
**ModifiedAt** | **string** | Date and time when the workflow template was last edited | 
**ModifiedByName** | **string** | Name of the user who last edited the workflow template | 
**ModifiedById** | **int32** | ID of the user who last edited the workflow template | 
**IsDeleted** | **bool** | Indicates whether the workflow template is deleted | 
**Status** | **string** | Status of the workflow template. Can be PUBLISHED or UNPUBLISHED | 
**ResourceType** | **string** | Type of the resource, in this case COMPUTE | 
**ResourceParams** | Pointer to [**[]WorkflowTemplateResourceParamsInner**](WorkflowTemplateResourceParamsInner.md) | Parameters of the resource limitation. The parameters depend on the resource type and include parameters for compute instance. | [optional] 
**Tags** | Pointer to [**[]WorkflowTemplateCreateTagsInner**](WorkflowTemplateCreateTagsInner.md) | List of tags assigned to the workflow template | [optional] 
**ContentType** | **string** | Format of the content field, only Shell is supported for now | 
**Content** | **string** | Content of the workflow template. For Shell contentType, it includes shell commands to execute. | 
**IsContentValid** | Pointer to **bool** | Indicates whether the content of the workflow template is valid. The content is validated against the contentType, e.g. for Shell contentType, the content is validated to be correct shell commands. | [optional] 
**ContentParams** | Pointer to [**[]WorkflowTemplateCreateContentParamsInner**](WorkflowTemplateCreateContentParamsInner.md) | List of parameters extracted from the content. These parameters can be used to fill in the blanks in the content when creating a workflow from the template. | [optional] 

## Methods

### NewWorkflowTemplate

`func NewWorkflowTemplate(id int32, name string, createdAt string, createdByName string, createdById int32, modifiedAt string, modifiedByName string, modifiedById int32, isDeleted bool, status string, resourceType string, contentType string, content string, ) *WorkflowTemplate`

NewWorkflowTemplate instantiates a new WorkflowTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateWithDefaults

`func NewWorkflowTemplateWithDefaults() *WorkflowTemplate`

NewWorkflowTemplateWithDefaults instantiates a new WorkflowTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkflowTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByName

`func (o *WorkflowTemplate) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *WorkflowTemplate) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *WorkflowTemplate) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.


### GetCreatedById

`func (o *WorkflowTemplate) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkflowTemplate) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkflowTemplate) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedAt

`func (o *WorkflowTemplate) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *WorkflowTemplate) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *WorkflowTemplate) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByName

`func (o *WorkflowTemplate) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *WorkflowTemplate) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *WorkflowTemplate) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.


### GetModifiedById

`func (o *WorkflowTemplate) GetModifiedById() int32`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkflowTemplate) GetModifiedByIdOk() (*int32, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkflowTemplate) SetModifiedById(v int32)`

SetModifiedById sets ModifiedById field to given value.


### GetIsDeleted

`func (o *WorkflowTemplate) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkflowTemplate) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkflowTemplate) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetStatus

`func (o *WorkflowTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResourceType

`func (o *WorkflowTemplate) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WorkflowTemplate) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WorkflowTemplate) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceParams

`func (o *WorkflowTemplate) GetResourceParams() []WorkflowTemplateResourceParamsInner`

GetResourceParams returns the ResourceParams field if non-nil, zero value otherwise.

### GetResourceParamsOk

`func (o *WorkflowTemplate) GetResourceParamsOk() (*[]WorkflowTemplateResourceParamsInner, bool)`

GetResourceParamsOk returns a tuple with the ResourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceParams

`func (o *WorkflowTemplate) SetResourceParams(v []WorkflowTemplateResourceParamsInner)`

SetResourceParams sets ResourceParams field to given value.

### HasResourceParams

`func (o *WorkflowTemplate) HasResourceParams() bool`

HasResourceParams returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowTemplate) GetTags() []WorkflowTemplateCreateTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowTemplate) GetTagsOk() (*[]WorkflowTemplateCreateTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowTemplate) SetTags(v []WorkflowTemplateCreateTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContentType

`func (o *WorkflowTemplate) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkflowTemplate) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkflowTemplate) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContent

`func (o *WorkflowTemplate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkflowTemplate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WorkflowTemplate) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsContentValid

`func (o *WorkflowTemplate) GetIsContentValid() bool`

GetIsContentValid returns the IsContentValid field if non-nil, zero value otherwise.

### GetIsContentValidOk

`func (o *WorkflowTemplate) GetIsContentValidOk() (*bool, bool)`

GetIsContentValidOk returns a tuple with the IsContentValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContentValid

`func (o *WorkflowTemplate) SetIsContentValid(v bool)`

SetIsContentValid sets IsContentValid field to given value.

### HasIsContentValid

`func (o *WorkflowTemplate) HasIsContentValid() bool`

HasIsContentValid returns a boolean if a field has been set.

### GetContentParams

`func (o *WorkflowTemplate) GetContentParams() []WorkflowTemplateCreateContentParamsInner`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *WorkflowTemplate) GetContentParamsOk() (*[]WorkflowTemplateCreateContentParamsInner, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *WorkflowTemplate) SetContentParams(v []WorkflowTemplateCreateContentParamsInner)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *WorkflowTemplate) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


