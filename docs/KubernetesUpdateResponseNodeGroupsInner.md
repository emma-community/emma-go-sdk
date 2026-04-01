# KubernetesUpdateResponseNodeGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Node group name | 
**Nodes** | [**[]KubernetesUpdateResponseNodeGroupsInnerNodesInner**](KubernetesUpdateResponseNodeGroupsInnerNodesInner.md) | List of nodes in the node group | 

## Methods

### NewKubernetesUpdateResponseNodeGroupsInner

`func NewKubernetesUpdateResponseNodeGroupsInner(name string, nodes []KubernetesUpdateResponseNodeGroupsInnerNodesInner, ) *KubernetesUpdateResponseNodeGroupsInner`

NewKubernetesUpdateResponseNodeGroupsInner instantiates a new KubernetesUpdateResponseNodeGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdateResponseNodeGroupsInnerWithDefaults

`func NewKubernetesUpdateResponseNodeGroupsInnerWithDefaults() *KubernetesUpdateResponseNodeGroupsInner`

NewKubernetesUpdateResponseNodeGroupsInnerWithDefaults instantiates a new KubernetesUpdateResponseNodeGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesUpdateResponseNodeGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesUpdateResponseNodeGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesUpdateResponseNodeGroupsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNodes

`func (o *KubernetesUpdateResponseNodeGroupsInner) GetNodes() []KubernetesUpdateResponseNodeGroupsInnerNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesUpdateResponseNodeGroupsInner) GetNodesOk() (*[]KubernetesUpdateResponseNodeGroupsInnerNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesUpdateResponseNodeGroupsInner) SetNodes(v []KubernetesUpdateResponseNodeGroupsInnerNodesInner)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


