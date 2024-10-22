# KubernetesClusterObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** | Type of kubernetes cluster object | [optional] 
**ObjectName** | Pointer to **string** | Name of the object for which the user requests a list of metrics | [optional] 

## Methods

### NewKubernetesClusterObjectsResponse

`func NewKubernetesClusterObjectsResponse() *KubernetesClusterObjectsResponse`

NewKubernetesClusterObjectsResponse instantiates a new KubernetesClusterObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterObjectsResponseWithDefaults

`func NewKubernetesClusterObjectsResponseWithDefaults() *KubernetesClusterObjectsResponse`

NewKubernetesClusterObjectsResponseWithDefaults instantiates a new KubernetesClusterObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesClusterObjectsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterObjectsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterObjectsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KubernetesClusterObjectsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectName

`func (o *KubernetesClusterObjectsResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesClusterObjectsResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesClusterObjectsResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *KubernetesClusterObjectsResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


