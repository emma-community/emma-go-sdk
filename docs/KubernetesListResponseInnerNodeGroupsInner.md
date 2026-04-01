# KubernetesListResponseInnerNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the worker node group | [optional] 
**Nodes** | Pointer to [**[]KubernetesListResponseInnerNodeGroupsInnerNodesInner**](KubernetesListResponseInnerNodeGroupsInnerNodesInner.md) | List of the worker nodes | [optional] 

## Methods

### NewKubernetesListResponseInnerNodeGroupsInner

`func NewKubernetesListResponseInnerNodeGroupsInner() *KubernetesListResponseInnerNodeGroupsInner`

NewKubernetesListResponseInnerNodeGroupsInner instantiates a new KubernetesListResponseInnerNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesListResponseInnerNodeGroupsInnerWithDefaults

`func NewKubernetesListResponseInnerNodeGroupsInnerWithDefaults() *KubernetesListResponseInnerNodeGroupsInner`

NewKubernetesListResponseInnerNodeGroupsInnerWithDefaults instantiates a new KubernetesListResponseInnerNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesListResponseInnerNodeGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesListResponseInnerNodeGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesListResponseInnerNodeGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesListResponseInnerNodeGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesListResponseInnerNodeGroupsInner) GetNodes() []KubernetesListResponseInnerNodeGroupsInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesListResponseInnerNodeGroupsInner) GetNodesOk() (*[]KubernetesListResponseInnerNodeGroupsInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesListResponseInnerNodeGroupsInner) SetNodes(v []KubernetesListResponseInnerNodeGroupsInnerNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesListResponseInnerNodeGroupsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


