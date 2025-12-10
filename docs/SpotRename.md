# SpotRename

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action with the spot instance - rename | 
**Name** | **string** | A new name of the spot instance | 

## Methods

### NewSpotRename

`func NewSpotRename(action string, name string, ) *SpotRename`

NewSpotRename instantiates a new SpotRename object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotRenameWithDefaults

`func NewSpotRenameWithDefaults() *SpotRename`

NewSpotRenameWithDefaults instantiates a new SpotRename object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SpotRename) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SpotRename) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SpotRename) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *SpotRename) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpotRename) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpotRename) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


