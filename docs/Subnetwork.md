# Subnetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Subnetwork id | [optional] 
**Name** | Pointer to **string** | Subnetwork name | [optional] 
**DataCenterId** | Pointer to **string** | ID of the cloud provider&#39;s data center | [optional] 
**SubnetworkPrefix** | Pointer to **string** | The prefix for the subnetwork IP address range | [optional] 
**SubnetworkSize** | Pointer to **int32** | The net mask size for the subnetwork | [optional] 
**Resources** | Pointer to [**SubnetworkResources**](SubnetworkResources.md) |  | [optional] 

## Methods

### NewSubnetwork

`func NewSubnetwork() *Subnetwork`

NewSubnetwork instantiates a new Subnetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetworkWithDefaults

`func NewSubnetworkWithDefaults() *Subnetwork`

NewSubnetworkWithDefaults instantiates a new Subnetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subnetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subnetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Subnetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subnetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataCenterId

`func (o *Subnetwork) GetDataCenterId() string`

GetDataCenterId returns the DataCenterId field if non-nil, zero value otherwise.

### GetDataCenterIdOk

`func (o *Subnetwork) GetDataCenterIdOk() (*string, bool)`

GetDataCenterIdOk returns a tuple with the DataCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterId

`func (o *Subnetwork) SetDataCenterId(v string)`

SetDataCenterId sets DataCenterId field to given value.

### HasDataCenterId

`func (o *Subnetwork) HasDataCenterId() bool`

HasDataCenterId returns a boolean if a field has been set.

### GetSubnetworkPrefix

`func (o *Subnetwork) GetSubnetworkPrefix() string`

GetSubnetworkPrefix returns the SubnetworkPrefix field if non-nil, zero value otherwise.

### GetSubnetworkPrefixOk

`func (o *Subnetwork) GetSubnetworkPrefixOk() (*string, bool)`

GetSubnetworkPrefixOk returns a tuple with the SubnetworkPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkPrefix

`func (o *Subnetwork) SetSubnetworkPrefix(v string)`

SetSubnetworkPrefix sets SubnetworkPrefix field to given value.

### HasSubnetworkPrefix

`func (o *Subnetwork) HasSubnetworkPrefix() bool`

HasSubnetworkPrefix returns a boolean if a field has been set.

### GetSubnetworkSize

`func (o *Subnetwork) GetSubnetworkSize() int32`

GetSubnetworkSize returns the SubnetworkSize field if non-nil, zero value otherwise.

### GetSubnetworkSizeOk

`func (o *Subnetwork) GetSubnetworkSizeOk() (*int32, bool)`

GetSubnetworkSizeOk returns a tuple with the SubnetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkSize

`func (o *Subnetwork) SetSubnetworkSize(v int32)`

SetSubnetworkSize sets SubnetworkSize field to given value.

### HasSubnetworkSize

`func (o *Subnetwork) HasSubnetworkSize() bool`

HasSubnetworkSize returns a boolean if a field has been set.

### GetResources

`func (o *Subnetwork) GetResources() SubnetworkResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Subnetwork) GetResourcesOk() (*SubnetworkResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Subnetwork) SetResources(v SubnetworkResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Subnetwork) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


