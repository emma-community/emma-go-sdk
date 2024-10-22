# ForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code | 
**Message** | **string** | Error message | 

## Methods

### NewForbiddenError

`func NewForbiddenError(code string, message string, ) *ForbiddenError`

NewForbiddenError instantiates a new ForbiddenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenErrorWithDefaults

`func NewForbiddenErrorWithDefaults() *ForbiddenError`

NewForbiddenErrorWithDefaults instantiates a new ForbiddenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ForbiddenError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ForbiddenError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ForbiddenError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ForbiddenError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ForbiddenError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ForbiddenError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


