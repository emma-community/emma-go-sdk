# WorkflowTemplateCreateContentParamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the content parameter, e.g. command1, command2. | [optional] 
**DefaultValue** | Pointer to **string** | Value of the content parameter. | [optional] 
**Mandatory** | Pointer to **bool** | Indicates whether the content parameter is mandatory to fill in when creating | [optional] 

## Methods

### NewWorkflowTemplateCreateContentParamsInner

`func NewWorkflowTemplateCreateContentParamsInner() *WorkflowTemplateCreateContentParamsInner`

NewWorkflowTemplateCreateContentParamsInner instantiates a new WorkflowTemplateCreateContentParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateCreateContentParamsInnerWithDefaults

`func NewWorkflowTemplateCreateContentParamsInnerWithDefaults() *WorkflowTemplateCreateContentParamsInner`

NewWorkflowTemplateCreateContentParamsInnerWithDefaults instantiates a new WorkflowTemplateCreateContentParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowTemplateCreateContentParamsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateCreateContentParamsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateCreateContentParamsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTemplateCreateContentParamsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultValue

`func (o *WorkflowTemplateCreateContentParamsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *WorkflowTemplateCreateContentParamsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *WorkflowTemplateCreateContentParamsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *WorkflowTemplateCreateContentParamsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetMandatory

`func (o *WorkflowTemplateCreateContentParamsInner) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *WorkflowTemplateCreateContentParamsInner) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *WorkflowTemplateCreateContentParamsInner) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *WorkflowTemplateCreateContentParamsInner) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


