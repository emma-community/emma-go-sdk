# VmClone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **interface{}** | Action with a virtual machine | 
**Name** | **interface{}** | Virtual machine name | 

## Methods

### NewVmClone

`func NewVmClone(action interface{}, name interface{}, ) *VmClone`

NewVmClone instantiates a new VmClone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCloneWithDefaults

`func NewVmCloneWithDefaults() *VmClone`

NewVmCloneWithDefaults instantiates a new VmClone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VmClone) GetAction() interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VmClone) GetActionOk() (*interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VmClone) SetAction(v interface{})`

SetAction sets Action field to given value.


### SetActionNil

`func (o *VmClone) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *VmClone) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetName

`func (o *VmClone) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmClone) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmClone) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *VmClone) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmClone) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


