# KubernetesDeleteNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node group name | [optional] 
**Nodes** | Pointer to [**[]KubernetesDeleteNodeGroupsInnerNodesInner**](KubernetesDeleteNodeGroupsInnerNodesInner.md) | List of nodes in the node group | [optional] 

## Methods

### NewKubernetesDeleteNodeGroupsInner

`func NewKubernetesDeleteNodeGroupsInner() *KubernetesDeleteNodeGroupsInner`

NewKubernetesDeleteNodeGroupsInner instantiates a new KubernetesDeleteNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeleteNodeGroupsInnerWithDefaults

`func NewKubernetesDeleteNodeGroupsInnerWithDefaults() *KubernetesDeleteNodeGroupsInner`

NewKubernetesDeleteNodeGroupsInnerWithDefaults instantiates a new KubernetesDeleteNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesDeleteNodeGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesDeleteNodeGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesDeleteNodeGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesDeleteNodeGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesDeleteNodeGroupsInner) GetNodes() []KubernetesDeleteNodeGroupsInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesDeleteNodeGroupsInner) GetNodesOk() (*[]KubernetesDeleteNodeGroupsInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesDeleteNodeGroupsInner) SetNodes(v []KubernetesDeleteNodeGroupsInnerNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesDeleteNodeGroupsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


