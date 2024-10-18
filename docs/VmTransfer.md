# VmTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with a virtual machine | 
**DataCenterId** | **string** | ID of the provider&#39;s data center | 
**IsKeepOriginalInstance** | Pointer to **bool** | Indicate if it is necessary to keep the original instance | [optional] [default to true]

## Methods

### NewVmTransfer

`func NewVmTransfer(action string, dataCenterId string, ) *VmTransfer`

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

`func (o *VmTransfer) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmTransfer) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmTransfer) SetAction(v string)`

SetAction sets Action field to given value.


### GetDataCenterId

`func (o *VmTransfer) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *VmTransfer) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *VmTransfer) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetIsKeepOriginalInstance

`func (o *VmTransfer) GetIsKeepOriginalInstance() bool`

GetIsKeepOriginalInstance returns the IsKeepOriginalInstance field if non-nil, zero value otherwise.

### GetIsKeepOriginalInstanceOk

`func (o *VmTransfer) GetIsKeepOriginalInstanceOk() (*bool, bool)`

GetIsKeepOriginalInstanceOk returns a tuple with the IsKeepOriginalInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepOriginalInstance

`func (o *VmTransfer) SetIsKeepOriginalInstance(v bool)`

SetIsKeepOriginalInstance sets IsKeepOriginalInstance field to given value.

### HasIsKeepOriginalInstance

`func (o *VmTransfer) HasIsKeepOriginalInstance() bool`

HasIsKeepOriginalInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


