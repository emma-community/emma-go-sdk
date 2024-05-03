/*
Public EMMA API

This <b>Infrastructure</b> API is for managing the cloud infrastructure within a project.  To access the API, enter your project, navigate to <b>Settings</b> > <b>Service Apps</b>, and create a service application. Select the access level: <b>Read</b>, <b>Operate</b>, or <b>Manage</b>.  After creating the service application, copy the <b>Client ID</b> and <b>Client Secret</b>. Send an API request to the endpoint <b>/issue-token</b> as specified in the <b>Authentication</b> section of the API documentation. You will receive access and refresh tokens in the response.  The Bearer access token is a text string, included in the request header:  -H \"Authorization: Bearer {token}\"  Use this token for API requests.  The access token will expire in 5 minutes, after which it must be refreshed using the refresh token.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the SshKeyGenerated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshKeyGenerated{}

// SshKeyGenerated struct for SshKeyGenerated
type SshKeyGenerated struct {
	Id            interface{} `json:"id,omitempty"`
	Name          interface{} `json:"name,omitempty"`
	Key           interface{} `json:"key,omitempty"`
	Fingerprint   interface{} `json:"fingerprint,omitempty"`
	KeyType       interface{} `json:"keyType,omitempty"`
	PrivateKey    interface{} `json:"privateKey,omitempty"`
	CreatedAt     interface{} `json:"createdAt,omitempty"`
	CreatedByName interface{} `json:"createdByName,omitempty"`
	CreatedById   interface{} `json:"createdById,omitempty"`
}

// NewSshKeyGenerated instantiates a new SshKeyGenerated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyGenerated() *SshKeyGenerated {
	this := SshKeyGenerated{}
	return &this
}

// NewSshKeyGeneratedWithDefaults instantiates a new SshKeyGenerated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyGeneratedWithDefaults() *SshKeyGenerated {
	this := SshKeyGenerated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *SshKeyGenerated) SetId(v interface{}) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *SshKeyGenerated) SetName(v interface{}) {
	o.Name = v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *SshKeyGenerated) SetKey(v interface{}) {
	o.Key = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetFingerprint() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetFingerprintOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return &o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given interface{} and assigns it to the Fingerprint field.
func (o *SshKeyGenerated) SetFingerprint(v interface{}) {
	o.Fingerprint = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetKeyType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetKeyTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return &o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given interface{} and assigns it to the KeyType field.
func (o *SshKeyGenerated) SetKeyType(v interface{}) {
	o.KeyType = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetPrivateKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetPrivateKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return &o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given interface{} and assigns it to the PrivateKey field.
func (o *SshKeyGenerated) SetPrivateKey(v interface{}) {
	o.PrivateKey = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *SshKeyGenerated) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetCreatedByName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetCreatedByNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedByName) {
		return nil, false
	}
	return &o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedByName() bool {
	if o != nil && !IsNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given interface{} and assigns it to the CreatedByName field.
func (o *SshKeyGenerated) SetCreatedByName(v interface{}) {
	o.CreatedByName = v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshKeyGenerated) GetCreatedById() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshKeyGenerated) GetCreatedByIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return &o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given interface{} and assigns it to the CreatedById field.
func (o *SshKeyGenerated) SetCreatedById(v interface{}) {
	o.CreatedById = v
}

func (o SshKeyGenerated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyGenerated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.KeyType != nil {
		toSerialize["keyType"] = o.KeyType
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedByName != nil {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if o.CreatedById != nil {
		toSerialize["createdById"] = o.CreatedById
	}
	return toSerialize, nil
}

type NullableSshKeyGenerated struct {
	value *SshKeyGenerated
	isSet bool
}

func (v NullableSshKeyGenerated) Get() *SshKeyGenerated {
	return v.value
}

func (v *NullableSshKeyGenerated) Set(val *SshKeyGenerated) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyGenerated) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyGenerated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyGenerated(val *SshKeyGenerated) *NullableSshKeyGenerated {
	return &NullableSshKeyGenerated{value: val, isSet: true}
}

func (v NullableSshKeyGenerated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyGenerated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
