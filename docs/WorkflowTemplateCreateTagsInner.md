# WorkflowTemplateCreateTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | Pointer to **string** | ID of the tag | [optional] 
**Key** | Pointer to **string** | Name of the tag, e.g. environment, department, project. | [optional] 
**Value** | Pointer to **string** | Value of the tag, e.g. production, staging, development for environment tag; sales, marketing for department tag; alpha, beta for project tag. | [optional] 

## Methods

### NewWorkflowTemplateCreateTagsInner

`func NewWorkflowTemplateCreateTagsInner() *WorkflowTemplateCreateTagsInner`

NewWorkflowTemplateCreateTagsInner instantiates a new WorkflowTemplateCreateTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateCreateTagsInnerWithDefaults

`func NewWorkflowTemplateCreateTagsInnerWithDefaults() *WorkflowTemplateCreateTagsInner`

NewWorkflowTemplateCreateTagsInnerWithDefaults instantiates a new WorkflowTemplateCreateTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *WorkflowTemplateCreateTagsInner) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *WorkflowTemplateCreateTagsInner) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *WorkflowTemplateCreateTagsInner) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *WorkflowTemplateCreateTagsInner) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetKey

`func (o *WorkflowTemplateCreateTagsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowTemplateCreateTagsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowTemplateCreateTagsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowTemplateCreateTagsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowTemplateCreateTagsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowTemplateCreateTagsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowTemplateCreateTagsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowTemplateCreateTagsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


