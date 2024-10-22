# KubernetesNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the worker node group | [optional] 
**Nodes** | Pointer to [**[]KubernetesNodeGroupsInnerNodesInner**](KubernetesNodeGroupsInnerNodesInner.md) | List of the worker nodes | [optional] 

## Methods

### NewKubernetesNodeGroupsInner

`func NewKubernetesNodeGroupsInner() *KubernetesNodeGroupsInner`

NewKubernetesNodeGroupsInner instantiates a new KubernetesNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeGroupsInnerWithDefaults

`func NewKubernetesNodeGroupsInnerWithDefaults() *KubernetesNodeGroupsInner`

NewKubernetesNodeGroupsInnerWithDefaults instantiates a new KubernetesNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNodeGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodeGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodeGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNodeGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesNodeGroupsInner) GetNodes() []KubernetesNodeGroupsInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesNodeGroupsInner) GetNodesOk() (*[]KubernetesNodeGroupsInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesNodeGroupsInner) SetNodes(v []KubernetesNodeGroupsInnerNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesNodeGroupsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


