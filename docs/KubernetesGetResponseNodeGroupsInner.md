# KubernetesGetResponseNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the worker node group | [optional] 
**Nodes** | Pointer to [**[]KubernetesGetResponseNodeGroupsInnerNodesInner**](KubernetesGetResponseNodeGroupsInnerNodesInner.md) | List of the worker nodes | [optional] 

## Methods

### NewKubernetesGetResponseNodeGroupsInner

`func NewKubernetesGetResponseNodeGroupsInner() *KubernetesGetResponseNodeGroupsInner`

NewKubernetesGetResponseNodeGroupsInner instantiates a new KubernetesGetResponseNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesGetResponseNodeGroupsInnerWithDefaults

`func NewKubernetesGetResponseNodeGroupsInnerWithDefaults() *KubernetesGetResponseNodeGroupsInner`

NewKubernetesGetResponseNodeGroupsInnerWithDefaults instantiates a new KubernetesGetResponseNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesGetResponseNodeGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesGetResponseNodeGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesGetResponseNodeGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesGetResponseNodeGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesGetResponseNodeGroupsInner) GetNodes() []KubernetesGetResponseNodeGroupsInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesGetResponseNodeGroupsInner) GetNodesOk() (*[]KubernetesGetResponseNodeGroupsInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesGetResponseNodeGroupsInner) SetNodes(v []KubernetesGetResponseNodeGroupsInnerNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesGetResponseNodeGroupsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


