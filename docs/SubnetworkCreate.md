# SubnetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Subnetwork name | [optional] 
**DataCenterId** | **string** | ID of the cloud provider&#39;s data center | 
**SubnetworkPrefix** | Pointer to **string** | The prefix for the subnetwork IP address range | [optional] 
**SubnetworkSize** | **int32** | The net mask size for the subnetwork | 

## Methods

### NewSubnetworkCreate

`func NewSubnetworkCreate(dataCenterId string, subnetworkSize int32, ) *SubnetworkCreate`

NewSubnetworkCreate instantiates a new SubnetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetworkCreateWithDefaults

`func NewSubnetworkCreateWithDefaults() *SubnetworkCreate`

NewSubnetworkCreateWithDefaults instantiates a new SubnetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubnetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubnetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubnetworkCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubnetworkCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *SubnetworkCreate) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *SubnetworkCreate) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *SubnetworkCreate) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.


### GetSubnetworkPrefix

`func (o *SubnetworkCreate) GetSubnetworkPrefix() string`

GetSubnetworkPrefix returns the SubnetworkPrefix field if non-nil, zero value otherwise.

### GetSubnetworkPrefixOk

`func (o *SubnetworkCreate) GetSubnetworkPrefixOk() (*string, bool)`

GetSubnetworkPrefixOk returns a tuple with the SubnetworkPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkPrefix

`func (o *SubnetworkCreate) SetSubnetworkPrefix(v string)`

SetSubnetworkPrefix sets SubnetworkPrefix field to given value.

### HasSubnetworkPrefix

`func (o *SubnetworkCreate) HasSubnetworkPrefix() bool`

HasSubnetworkPrefix returns a boolean if a field has been set.

### GetSubnetworkSize

`func (o *SubnetworkCreate) GetSubnetworkSize() int32`

GetSubnetworkSize returns the SubnetworkSize field if non-nil, zero value otherwise.

### GetSubnetworkSizeOk

`func (o *SubnetworkCreate) GetSubnetworkSizeOk() (*int32, bool)`

GetSubnetworkSizeOk returns a tuple with the SubnetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkSize

`func (o *SubnetworkCreate) SetSubnetworkSize(v int32)`

SetSubnetworkSize sets SubnetworkSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


