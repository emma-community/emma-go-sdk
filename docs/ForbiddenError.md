# ForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **interface{}** |  | 
**Message** | **interface{}** |  | 

## Methods

### NewForbiddenError

`func NewForbiddenError(code interface{}, message interface{}, ) *ForbiddenError`

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

`func (o *ForbiddenError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ForbiddenError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ForbiddenError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ForbiddenError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ForbiddenError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *ForbiddenError) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ForbiddenError) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ForbiddenError) SetMessage(v interface{})`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ForbiddenError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ForbiddenError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


