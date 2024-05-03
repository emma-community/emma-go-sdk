# VmTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **interface{}** | Action with a virtual machine | 
**DataCenterId** | **interface{}** | Provider&#39;s data center ID | 
**IsKeepOriginalInstance** | Pointer to **interface{}** | Keep original instance | [optional] [default to true]

## Methods

### NewVmTransfer

`func NewVmTransfer(action interface{}, dataCenterId interface{}, ) *VmTransfer`

NewVmTransfer instantiates a new VmTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmTransferWithDefaults

`func NewVmTransferWithDefaults() *VmTransfer`

NewVmTransferWithDefaults instantiates a new VmTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmTransfer) GetAction() interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmTransfer) GetActionOk() (*interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmTransfer) SetAction(v interface{})`

SetAction sets Action field to given value.


### SetActionNil

`func (o *VmTransfer) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *VmTransfer) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDataCenterId

`func (o *VmTransfer) GetDataCenterId() interface{}`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmTransfer) GetDataCenterIdOk() (*interface{}, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmTransfer) SetDataCenterId(v interface{})`

SetDataCenterId sets DataCenterId field to given value.


### SetDataCenterIdNil

`func (o *VmTransfer) SetDataCenterIdNil(b bool)`

 SetDataCenterIdNil sets the value for DataCenterId to be an explicit nil

### UnsetDataCenterId
`func (o *VmTransfer) UnsetDataCenterId()`

UnsetDataCenterId ensures that no value is present for DataCenterId, not even an explicit nil
### GetIsKeepOriginalInstance

`func (o *VmTransfer) GetIsKeepOriginalInstance() interface{}`

GetIsKeepOriginalInstance returns the IsKeepOriginalInstance field if non-nil, zero value otherwise.

### GetIsKeepOriginalInstanceOk

`func (o *VmTransfer) GetIsKeepOriginalInstanceOk() (*interface{}, bool)`

GetIsKeepOriginalInstanceOk returns a tuple with the IsKeepOriginalInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepOriginalInstance

`func (o *VmTransfer) SetIsKeepOriginalInstance(v interface{})`

SetIsKeepOriginalInstance sets IsKeepOriginalInstance field to given value.

### HasIsKeepOriginalInstance

`func (o *VmTransfer) HasIsKeepOriginalInstance() bool`

HasIsKeepOriginalInstance returns a boolean if a field has been set.

### SetIsKeepOriginalInstanceNil

`func (o *VmTransfer) SetIsKeepOriginalInstanceNil(b bool)`

 SetIsKeepOriginalInstanceNil sets the value for IsKeepOriginalInstance to be an explicit nil

### UnsetIsKeepOriginalInstance
`func (o *VmTransfer) UnsetIsKeepOriginalInstance()`

UnsetIsKeepOriginalInstance ensures that no value is present for IsKeepOriginalInstance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


