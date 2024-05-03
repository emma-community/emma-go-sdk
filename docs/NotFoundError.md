# NotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewNotFoundError

`func NewNotFoundError(code string, message string, ) *NotFoundError`

NewNotFoundError instantiates a new NotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorWithDefaults

`func NewNotFoundErrorWithDefaults() *NotFoundError`

NewNotFoundErrorWithDefaults instantiates a new NotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NotFoundError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *NotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


